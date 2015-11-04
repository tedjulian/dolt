// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_optional_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("OptionalStruct",
			[]types.Field{
				types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), true},
				types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), true},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_optional_CachedRef = types.RegisterPackage(&p)
}

// OptionalStruct

type OptionalStruct struct {
	_s          string
	__optionals bool
	_b          bool
	__optionalb bool

	ref *ref.Ref
}

func NewOptionalStruct() OptionalStruct {
	return OptionalStruct{

		ref: &ref.Ref{},
	}
}

type OptionalStructDef struct {
	S string
	B bool
}

func (def OptionalStructDef) New() OptionalStruct {
	return OptionalStruct{
		_s:          def.S,
		__optionals: true,
		_b:          def.B,
		__optionalb: true,
		ref:         &ref.Ref{},
	}
}

func (s OptionalStruct) Def() (d OptionalStructDef) {
	if s.__optionals {
		d.S = s._s
	}
	if s.__optionalb {
		d.B = s._b
	}
	return
}

var __typeRefForOptionalStruct types.TypeRef
var __typeDefForOptionalStruct types.TypeRef

func (m OptionalStruct) TypeRef() types.TypeRef {
	return __typeRefForOptionalStruct
}

func init() {
	__typeRefForOptionalStruct = types.MakeTypeRef(__genPackageInFile_struct_optional_CachedRef, 0)
	__typeDefForOptionalStruct = types.MakeStructTypeRef("OptionalStruct",
		[]types.Field{
			types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), true},
			types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), true},
		},
		types.Choices{},
	)
	types.RegisterStructBuilder(__typeRefForOptionalStruct, builderForOptionalStruct)
}

func (s OptionalStruct) InternalImplementation() types.Struct {
	// TODO: Remove this
	m := map[string]types.Value{}
	if s.__optionals {
		m["s"] = types.NewString(s._s)
	}
	if s.__optionalb {
		m["b"] = types.Bool(s._b)
	}
	return types.NewStruct(__typeRefForOptionalStruct, __typeDefForOptionalStruct, m)
}

func builderForOptionalStruct() chan types.Value {
	c := make(chan types.Value)
	s := OptionalStruct{ref: &ref.Ref{}}
	go func() {
		s.__optionals = bool((<-c).(types.Bool))
		if s.__optionals {
			s._s = (<-c).(types.String).String()
		}
		s.__optionalb = bool((<-c).(types.Bool))
		if s.__optionalb {
			s._b = bool((<-c).(types.Bool))
		}

		c <- s
	}()
	return c
}

func (s OptionalStruct) Equals(other types.Value) bool {
	return other != nil && __typeRefForOptionalStruct.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s OptionalStruct) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s OptionalStruct) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForOptionalStruct.Chunks()...)
	return
}

func (s OptionalStruct) S() (v string, ok bool) {
	if s.__optionals {
		return s._s, true
	}
	return
}

func (s OptionalStruct) SetS(val string) OptionalStruct {
	s.__optionals = true
	s._s = val
	s.ref = &ref.Ref{}
	return s
}

func (s OptionalStruct) B() (v bool, ok bool) {
	if s.__optionalb {
		return s._b, true
	}
	return
}

func (s OptionalStruct) SetB(val bool) OptionalStruct {
	s.__optionalb = true
	s._b = val
	s.ref = &ref.Ref{}
	return s
}
