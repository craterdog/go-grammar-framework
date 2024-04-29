/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent_test

import (
	fmt "fmt"
	age "github.com/craterdog/go-grammar-framework/v4/cdsn/agent"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

func TestRoundtrips(t *tes.T) {
	var directories, err = osx.ReadDir(testDirectory)
	if err != nil {
		panic(err)
	}

	for _, directory := range directories {
		if sts.HasPrefix(directory.Name(), "go.") {
			continue // This is not a directory.
		}
		var directoryName = testDirectory + directory.Name() + "/"
		var syntaxFile = directoryName + "Syntax.cdsn"
		fmt.Println(syntaxFile)
		var parser = age.Parser().Make()
		var validator = age.Validator().Make()
		var formatter = age.Formatter().Make()
		var bytes, err = osx.ReadFile(syntaxFile)
		if err != nil {
			panic(err)
		}
		var expected = string(bytes)
		var syntax = parser.ParseSource(expected)
		validator.ValidateSyntax(syntax)
		var actual = formatter.FormatSyntax(syntax)
		ass.Equal(t, expected, actual)
	}
}

const header = `!>
HEADER
<!

`

func TestRuleInTokenDefinition(t *tes.T) {
	var parser = age.Parser().Make()
	var validator = age.Validator().Make()
	var source = header + `bad: Rule
Rule: "bad"
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"The definition for bad is invalid:\nA token definition cannot contain a rule name.\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateSyntax(parser.ParseSource(source))
}

func TestDoubleInversion(t *tes.T) {
	var parser = age.Parser().Make()
	var validator = age.Validator().Make()
	var source = header + `Bad: ~~[CONTROL]
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"An unexpected token was received by the parser: Token [type: delimiter, line: 5, position: 7]: \"~\"\n\x1b[36m0004: \n0005: Bad: ~~[CONTROL]\n \x1b[32m>>>────────⌃\x1b[36m\n0006: \n\x1b[0m\nWas expecting '[' from:\n  \x1b[32mFilter: \x1b[33m\"~\"? Atom\x1b[0m\n\n  \x1b[32mAtom: \x1b[33mGlyph | intrinsic\x1b[0m\n\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateSyntax(parser.ParseSource(source))
}

func TestInvertedString(t *tes.T) {
	var parser = age.Parser().Make()
	var validator = age.Validator().Make()
	var source = header + `Bad: ~"ow"
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"An unexpected token was received by the parser: Token [type: literal, line: 5, position: 7]: \"\\\"ow\\\"\"\n\x1b[36m0004: \n0005: Bad: ~\"ow\"\n \x1b[32m>>>────────⌃\x1b[36m\n0006: \n\x1b[0m\nWas expecting '[' from:\n  \x1b[32mFilter: \x1b[33m\"~\"? Atom\x1b[0m\n\n  \x1b[32mAtom: \x1b[33mGlyph | intrinsic\x1b[0m\n\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateSyntax(parser.ParseSource(source))
}

func TestInvertedRule(t *tes.T) {
	var parser = age.Parser().Make()
	var validator = age.Validator().Make()
	var source = header + `bad: ~rule
rule: "rule"
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"An unexpected token was received by the parser: Token [type: name, line: 5, position: 7]: \"rule\"\n\x1b[36m0004: \n0005: bad: ~rule\n \x1b[32m>>>────────⌃\x1b[36m\n0006: rule: \"rule\"\n\x1b[0m\nWas expecting '[' from:\n  \x1b[32mFilter: \x1b[33m\"~\"? Atom\x1b[0m\n\n  \x1b[32mAtom: \x1b[33mGlyph | intrinsic\x1b[0m\n\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateSyntax(parser.ParseSource(source))
}

func TestMissingRule(t *tes.T) {
	var parser = age.Parser().Make()
	var validator = age.Validator().Make()
	var source = header + `bad: rule
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"The syntax is missing a definition for the symbol: rule\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateSyntax(parser.ParseSource(source))
}

func TestDuplicateRule(t *tes.T) {
	var parser = age.Parser().Make()
	var validator = age.Validator().Make()
	var source = header + `bad: "bad"
bad: "worse"
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"The definition for bad is invalid:\nThe name bad is defined more than once.\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateSyntax(parser.ParseSource(source))
}

func TestNestedFilters(t *tes.T) {
	var parser = age.Parser().Make()
	var validator = age.Validator().Make()
	var source = header + `Bad: ~(Worse | ~Bad)
Worse: CONTROL
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"An unexpected token was received by the parser: Token [type: delimiter, line: 5, position: 7]: \"(\"\n\x1b[36m0004: \n0005: Bad: ~(Worse | ~Bad)\n \x1b[32m>>>────────⌃\x1b[36m\n0006: Worse: CONTROL\n\x1b[0m\nWas expecting '[' from:\n  \x1b[32mFilter: \x1b[33m\"~\"? Atom\x1b[0m\n\n  \x1b[32mAtom: \x1b[33mGlyph | intrinsic\x1b[0m\n\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateSyntax(parser.ParseSource(source))
}
