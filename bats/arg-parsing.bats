#!/usr/bin/env bats
load $BATS_TEST_DIRNAME/helper/common.bash

setup() {
    setup_common
}

teardown() {
    teardown_common
}

@test "dolt supports Nix style argument parsing" {
    dolt checkout -b this-should-work
    run dolt branch
    [ $status -eq 0 ]
    [[ "$output" =~ "this-should-work" ]] || false
    dolt checkout master
    dolt branch -d this-should-work

    dolt checkout -b "this-should-work"
    run dolt branch
    [ $status -eq 0 ]
    [[ "$output" =~ "this-should-work" ]] || false
    dolt checkout master
    dolt branch -d "this-should-work"

    dolt checkout --b "this-should-work"
    run dolt branch
    [ $status -eq 0 ]
    [[ "$output" =~ "this-should-work" ]] || false
    dolt checkout master
    dolt branch --d "this-should-work"

    run dolt checkout -bthis-should-work
    [ $status -eq 0 ]
    run dolt branch
    [ $status -eq 0 ] || echo $output
    [[ "$output" =~ "this-should-work" ]] || false
    dolt checkout master
    dolt branch -dthis-should-work
}

@test "dolt supports chaining of modal arguments" {
    dolt sql -q "create table test(pk int, primary key (pk))"
    dolt table import -fc test `batshelper 1pk5col-ints.csv`
}

@test "dolt checkout with empty string returns error" {
    run dolt checkout ""
    [[ "$output" =~ "error: cannot checkout empty string" ]] || false
    [ $status -ne 0 ]

    run dolt checkout -b ""
    [[ "$output" =~ "error: cannot checkout empty string" ]] || false
    [ $status -ne 0 ]
}
