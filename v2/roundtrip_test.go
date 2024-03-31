/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package grammars_test

import (
	fmt "fmt"
	gra "github.com/craterdog/go-grammar-framework/v2"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

const testDirectory = "./test/"

func TestRoundtrips(t *tes.T) {
	var files, err = osx.ReadDir(testDirectory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var parser = gra.Parser().Make()
		var validator = gra.Validator().Make()
		var formatter = gra.Formatter().Make()
		var filename = testDirectory + file.Name()
		if sts.HasSuffix(filename, ".cdsn") {
			fmt.Println(filename)
			var bytes, err = osx.ReadFile(filename)
			if err != nil {
				panic(err)
			}
			var expected = string(bytes)
			var grammar = parser.ParseSource(expected)
			validator.ValidateGrammar(grammar)
			var actual = formatter.FormatGrammar(grammar)
			ass.Equal(t, expected, actual)
		}
	}
}

const header = `!>
HEADER
<!

`

func TestRuleInTokenDefinition(t *tes.T) {
	var parser = gra.Parser().Make()
	var validator = gra.Validator().Make()
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

	validator.ValidateGrammar(parser.ParseSource(source))
}

func TestDoubleInversion(t *tes.T) {
	var parser = gra.Parser().Make()
	var validator = gra.Validator().Make()
	var source = header + `Bad: ~~CONTROL
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"An unexpected token was received by the parser: Token [type: delimiter, line: 5, position: 7]: \"~\"\n\x1b[36m0004: \n0005: Bad: ~~CONTROL\n \x1b[32m>>>────────⌃\x1b[36m\n0006: \n\x1b[0m\nWas expecting 'Filter' from:\n  \x1b[32mInversion: \x1b[33m\"~\"? Filter\x1b[0m\n\n  \x1b[32mFilter: \x1b[33mintrinsic | Glyph\x1b[0m\n\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateGrammar(parser.ParseSource(source))
}

func TestInvertedString(t *tes.T) {
	var parser = gra.Parser().Make()
	var validator = gra.Validator().Make()
	var source = header + `Bad: ~"ow"
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"An unexpected token was received by the parser: Token [type: literal, line: 5, position: 7]: \"\\\"ow\\\"\"\n\x1b[36m0004: \n0005: Bad: ~\"ow\"\n \x1b[32m>>>────────⌃\x1b[36m\n0006: \n\x1b[0m\nWas expecting 'Filter' from:\n  \x1b[32mInversion: \x1b[33m\"~\"? Filter\x1b[0m\n\n  \x1b[32mFilter: \x1b[33mintrinsic | Glyph\x1b[0m\n\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateGrammar(parser.ParseSource(source))
}

func TestInvertedRule(t *tes.T) {
	var parser = gra.Parser().Make()
	var validator = gra.Validator().Make()
	var source = header + `bad: ~rule
rule: "rule"
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"An unexpected token was received by the parser: Token [type: name, line: 5, position: 7]: \"rule\"\n\x1b[36m0004: \n0005: bad: ~rule\n \x1b[32m>>>────────⌃\x1b[36m\n0006: rule: \"rule\"\n\x1b[0m\nWas expecting 'Filter' from:\n  \x1b[32mInversion: \x1b[33m\"~\"? Filter\x1b[0m\n\n  \x1b[32mFilter: \x1b[33mintrinsic | Glyph\x1b[0m\n\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateGrammar(parser.ParseSource(source))
}

func TestMissingRule(t *tes.T) {
	var parser = gra.Parser().Make()
	var validator = gra.Validator().Make()
	var source = header + `bad: rule
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"The grammar is missing a definition for the symbol: rule\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateGrammar(parser.ParseSource(source))
}

func TestDuplicateRule(t *tes.T) {
	var parser = gra.Parser().Make()
	var validator = gra.Validator().Make()
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

	validator.ValidateGrammar(parser.ParseSource(source))
}

func TestNestedInversions(t *tes.T) {
	var parser = gra.Parser().Make()
	var validator = gra.Validator().Make()
	var source = header + `Bad: ~(Worse | ~Bad)
Worse: CONTROL
`
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(
				t,
				"An unexpected token was received by the parser: Token [type: delimiter, line: 5, position: 7]: \"(\"\n\x1b[36m0004: \n0005: Bad: ~(Worse | ~Bad)\n \x1b[32m>>>────────⌃\x1b[36m\n0006: Worse: CONTROL\n\x1b[0m\nWas expecting 'Filter' from:\n  \x1b[32mInversion: \x1b[33m\"~\"? Filter\x1b[0m\n\n  \x1b[32mFilter: \x1b[33mintrinsic | Glyph\x1b[0m\n\n",
				e,
			)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()

	validator.ValidateGrammar(parser.ParseSource(source))
}
