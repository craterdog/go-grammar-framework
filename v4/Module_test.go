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

package module_test

import (
	fmt "fmt"
	gra "github.com/craterdog/go-grammar-framework/v4"
	mod "github.com/craterdog/go-model-framework/v4"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

const module = "github.com/craterdog/go-grammar-framework"

const syntaxFile = "Syntax.cdsn"

func TestRoundTrips(t *tes.T) {
	var bytes, err = osx.ReadFile(syntaxFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = gra.Parser()
	var syntax = parser.ParseSource(source)
	var formatter = gra.Formatter()
	var actual = formatter.FormatSyntax(syntax)
	ass.Equal(t, actual, source)
	var validator = gra.Validator()
	validator.ValidateSyntax(syntax)
}

func TestModelGeneration(t *tes.T) {
	var bytes, err = osx.ReadFile(syntaxFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = gra.Parser()
	var syntax = parser.ParseSource(source)
	var generator = gra.Generator()
	var model = generator.GenerateAST(module, syntax)
	var formatter = mod.Formatter()
	source = formatter.FormatModel(model)
	bytes = []byte(source)
	var filename = "ast/Package.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func TestLifecycle(t *tes.T) {
	var generator = gra.Generator()
	var module = "github.com/craterdog/go-grammar-framework/v4"
	var name = "example"

	// Generate a new syntax with a default copyright.
	var copyright string
	var syntax = generator.CreateSyntax(name, copyright)

	// Validate the syntax.
	var validator = gra.Validator()
	validator.ValidateSyntax(syntax)

	// Format the syntax.
	var formatter = gra.Formatter()
	var source = formatter.FormatSyntax(syntax)
	fmt.Println("Syntax:")
	fmt.Println(source)

	// Parse the source code for the syntax.
	var parser = gra.Parser()
	syntax = parser.ParseSource(source)

	// Generate the AST model for the syntax.
	var model = generator.GenerateAST(module, syntax)
	var formatter2 = mod.Formatter()
	source = formatter2.FormatModel(model)
	fmt.Println("AST Model:")
	fmt.Println(source)

	// Generate the agent model for the syntax.
	model = generator.GenerateAgent(module, syntax)
	source = formatter2.FormatModel(model)
	fmt.Println("Agent Model:")
	fmt.Println(source)

	// Generate the formatter class for the syntax.
	generator.GenerateFormatter(module, syntax)

	// Generate the parser class for the syntax.
	generator.GenerateParser(module, syntax)

	// Generate the scanner class for the syntax.
	source = generator.GenerateScanner(module, syntax)
	fmt.Println("Scanner Class:")
	fmt.Println(source)

	// Generate the token class for the syntax.
	generator.GenerateToken(module, syntax)

	// Generate the validator class for the syntax.
	generator.GenerateValidator(module, syntax)
}
