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

package grammar_test

import (
	fmt "fmt"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var filenames = []string{
	"../Syntax.cdsn",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, filename := range filenames {
		fmt.Printf("   %v\n", filename)
		// Read in the syntax notation file.
		var bytes, err = osx.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		var source = string(bytes)

		// Parse the source code for the syntax notation.
		var parser = gra.Parser().Make()
		var syntax = parser.ParseSource(source)

		// Validate the syntax notation.
		var validator = gra.Validator().Make()
		validator.ValidateSyntax(syntax)

		// Format the syntax notation.
		var formatter = gra.Formatter().Make()
		var actual = formatter.FormatSyntax(syntax)
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
