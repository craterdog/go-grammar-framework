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
	sts "strings"
	tes "testing"
)

const (
	module = "github.com/craterdog/go-grammar-framework/v4"
	wiki   = "github.com/craterdog/go-grammar-framework/wiki"
)

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
	var formatter = mod.Formatter()

	var model = generator.GenerateAst(module, wiki, syntax)
	source = formatter.FormatModel(model)
	bytes = []byte(source)
	var filename = "ast/Package.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	var classes = model.GetClasses().GetClasses().GetIterator()
	for classes.HasNext() {
		var class = classes.GetNext()
		var name = sts.ToLower(sts.TrimSuffix(
			class.GetDeclaration().GetName(),
			"ClassLike",
		))
		var generator = mod.Generator()
		source = generator.GenerateClass(model, name)
		bytes = []byte(source)
		var filename = "ast/" + name + ".go"
		var err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}

	model = generator.GenerateGrammar(module, wiki, syntax)
	source = formatter.FormatModel(model)
	bytes = []byte(source)
	filename = "grammar/Package.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the scanner class for the syntax.
	source = generator.GenerateScanner(module, wiki, syntax)
	bytes = []byte(source)
	filename = "grammar/scanner.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the token class for the syntax.
	source = generator.GenerateToken(module, wiki, syntax)
	bytes = []byte(source)
	filename = "grammar/token.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func TestLifecycle(t *tes.T) {
	var generator = gra.Generator()
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
	var model = generator.GenerateAst(module, wiki, syntax)
	var formatter2 = mod.Formatter()
	source = formatter2.FormatModel(model)
	fmt.Println("AST Model:")
	fmt.Println(source)

	// Generate the language grammar model for the syntax.
	model = generator.GenerateGrammar(module, wiki, syntax)
	source = formatter2.FormatModel(model)
	fmt.Println("Grammar Model:")
	fmt.Println(source)

	// Generate the formatter class for the syntax.
	generator.GenerateFormatter(module, wiki, syntax)

	// Generate the parser class for the syntax.
	generator.GenerateParser(module, wiki, syntax)

	// Generate the scanner class for the syntax.
	source = generator.GenerateScanner(module, wiki, syntax)
	fmt.Println("Scanner Class:")
	fmt.Println(source)

	// Generate the token class for the syntax.
	generator.GenerateToken(module, wiki, syntax)

	// Generate the validator class for the syntax.
	generator.GenerateValidator(module, wiki, syntax)
}
