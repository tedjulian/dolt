// Copyright 2019 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/dolthub/dolt/go/cmd/dolt/cli"
	"github.com/dolthub/dolt/go/cmd/dolt/errhand"
	"github.com/dolthub/dolt/go/libraries/doltcore/diff"
	"github.com/dolthub/dolt/go/libraries/doltcore/env"
	"github.com/dolthub/dolt/go/libraries/doltcore/env/actions"
	"github.com/dolthub/dolt/go/libraries/utils/argparser"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
)

const (
	SoftResetParam = "soft"
	HardResetParam = "hard"
)

var resetDocContent = cli.CommandDocumentationContent{
	ShortDesc: "Resets staged or working tables to HEAD or a specified commit",
	LongDesc: "{{.EmphasisLeft}}dolt reset <tables>...{{.EmphasisRight}}" +
		"\n\n" +
		"The default form resets the values for all staged {{.LessThan}}tables{{.GreaterThan}} to their values at {{.EmphasisLeft}}HEAD{{.EmphasisRight}}. " +
		"It does not affect the working tree or the current branch." +
		"\n\n" +
		"This means that {{.EmphasisLeft}}dolt reset <tables>{{.EmphasisRight}} is the opposite of {{.EmphasisLeft}}dolt add <tables>{{.EmphasisRight}}." +
		"\n\n" +
		"After running {{.EmphasisLeft}}dolt reset <tables>{{.EmphasisRight}} to update the staged tables, you can use {{.EmphasisLeft}}dolt checkout{{.EmphasisRight}} to check the contents out of the staged tables to the working tables." +
		"\n\n" +
		"{{.EmphasisLeft}}dolt reset [--hard | --soft] <revision>{{.EmphasisRight}}" +
		"\n\n" +
		"This form resets all tables to values in the specified revision (i.e. commit, tag, working set). " +
		"The --soft option resets HEAD to a revision without changing the current working set. " +
		" The --hard option resets all three HEADs to a revision, deleting all uncommitted changes in the current working set." +
		"\n\n" +
		"{{.EmphasisLeft}}dolt reset .{{.EmphasisRight}}" +
		"\n\n" +
		"This form resets {{.EmphasisLeft}}all{{.EmphasisRight}} staged tables to their values at HEAD. It is the opposite of {{.EmphasisLeft}}dolt add .{{.EmphasisRight}}",
	Synopsis: []string{
		"{{.LessThan}}tables{{.GreaterThan}}...",
		"[--hard | --soft] {{.LessThan}}revision{{.GreaterThan}}",
	},
}

type ResetCmd struct{}

// Name is returns the name of the Dolt cli command. This is what is used on the command line to invoke the command
func (cmd ResetCmd) Name() string {
	return "reset"
}

// Description returns a description of the command
func (cmd ResetCmd) Description() string {
	return "Remove table changes from the list of staged table changes."
}

func (cmd ResetCmd) Docs() *cli.CommandDocumentation {
	ap := cli.CreateResetArgParser()
	return cli.NewCommandDocumentation(resetDocContent, ap)
}

func (cmd ResetCmd) ArgParser() *argparser.ArgParser {
	return cli.CreateResetArgParser()
}

func (cmd ResetCmd) RequiresRepo() bool {
	return false
}

// Exec executes the command
func (cmd ResetCmd) Exec(ctx context.Context, commandStr string, args []string, dEnv *env.DoltEnv, cliCtx cli.CliContext) int {
	ap := cli.CreateResetArgParser()
	help, usage := cli.HelpAndUsagePrinters(cli.CommandDocsForCommandString(commandStr, resetDocContent, ap))
	apr := cli.ParseArgsOrDie(ap, args, help)

	queryist, sqlCtx, closeFunc, err := cliCtx.QueryEngine(ctx)
	if err != nil {
		cli.Println(err.Error())
		return 1
	}
	if closeFunc != nil {
		defer closeFunc()
	}

	if apr.ContainsAll(HardResetParam, SoftResetParam) {
		verr := errhand.BuildDError("error: --%s and --%s are mutually exclusive options.", HardResetParam, SoftResetParam).Build()
		return HandleVErrAndExitCode(verr, usage)
	} else if apr.Contains(HardResetParam) {
		if apr.NArg() > 1 {
			return handleResetError(fmt.Errorf("--hard supports at most one additional param"), usage)
		}
	}

	// process query through prepared statement to prevent sql injection
	query, params := constructParametrizedDoltResetQuery(apr)
	interpolatedQuery, err := dbr.InterpolateForDialect(query, params, dialect.MySQL)
	if err != nil {
		cli.Println(err.Error())
		return 1
	}
	_, _, err = queryist.Query(sqlCtx, interpolatedQuery)
	if err != nil {
		cli.Println(err.Error())
		return 1
	}

	printNotStaged(sqlCtx, queryist)

	return 0
}

// constructParametrizedDoltResetQuery generates the sql query necessary to call the DOLT_RESET() stored procedure with placeholders
// for arg input. Also returns a list of the inputs in the order in which they appear in the query.
func constructParametrizedDoltResetQuery(apr *argparser.ArgParseResults) (string, []interface{}) {
	var params []interface{}
	var param bool

	var buffer bytes.Buffer
	var first bool
	first = true
	buffer.WriteString("CALL DOLT_RESET(")

	writeToBuffer := func(s string) {
		if !first {
			buffer.WriteString(", ")
		}
		if !param {
			buffer.WriteString("'")
		}
		buffer.WriteString(s)
		if !param {
			buffer.WriteString("'")
		}
		first = false
		param = false
	}

	if apr.Contains(HardResetParam) {
		writeToBuffer("--hard")
		if apr.NArg() == 1 {
			param = true
			writeToBuffer("?")
			params = append(params, apr.Arg(0))
		}
	} else if apr.Contains(SoftResetParam) {
		writeToBuffer("--soft")
		if apr.NArg() > 0 {
			for _, input := range apr.Args {
				param = true
				writeToBuffer("?")
				params = append(params, input)
			}
		}
	} else {
		for _, input := range apr.Args {
			if strings.ToLower(input) == "head" && apr.NArg() == 1 {
				buffer.Reset()
				buffer.WriteString("CALL DOLT_RESET(")
				break
			}
			param = true
			writeToBuffer("?")
			params = append(params, input)
		}
	}

	buffer.WriteString(")")
	return buffer.String(), params
}

var tblDiffTypeToShortLabel = map[diff.TableDiffType]string{
	diff.ModifiedTable: "M",
	diff.RemovedTable:  "D",
	diff.AddedTable:    "N",
}

func printNotStaged(sqlCtx *sql.Context, queryist cli.Queryist) {
	// Printing here is best effort.  Fail silently
	schema, rowIter, err := queryist.Query(sqlCtx, "select * from dolt_status where staged = false")
	if err != nil {
		return
	}
	rows, err := sql.RowIterToRows(sqlCtx, schema, rowIter)
	if err != nil {
		return
	}
	if rows == nil {
		return
	}

	removeModified := 0
	for _, row := range rows {
		if row[2] != "new table" {
			removeModified++
		}
	}

	if removeModified > 0 {
		cli.Println("Unstaged changes after reset:")

		var lines []string
		for _, row := range rows {
			if row[2] == "new table" {
				//  per Git, unstaged new tables are untracked
				continue
			} else if row[2] == "deleted" {
				lines = append(lines, fmt.Sprintf("%s\t%s", tblDiffTypeToShortLabel[diff.RemovedTable], row[0]))
			} else if row[2] == "renamed" {
				// per Git, unstaged renames are shown as drop + add
				lines = append(lines, fmt.Sprintf("%s\t%s", tblDiffTypeToShortLabel[diff.RemovedTable], row[0]))
			} else {
				lines = append(lines, fmt.Sprintf("%s\t%s", tblDiffTypeToShortLabel[diff.ModifiedTable], row[0]))
			}
		}
		cli.Println(strings.Join(lines, "\n"))
	}
}

func handleResetError(err error, usage cli.UsagePrinter) int {
	if actions.IsTblNotExist(err) {
		tbls := actions.GetTablesForError(err)

		// In case the ref does not exist.
		bdr := errhand.BuildDError("Invalid Ref or Table:")
		if len(tbls) > 1 {
			bdr = errhand.BuildDError("Invalid Table(s):")
		}

		for _, tbl := range tbls {
			bdr.AddDetails("\t" + tbl)
		}

		return HandleVErrAndExitCode(bdr.Build(), usage)
	}

	var verr errhand.VerboseError = nil
	if err != nil {
		verr = errhand.BuildDError("error: Failed to reset changes.").AddCause(err).Build()
	}

	return HandleVErrAndExitCode(verr, usage)
}
