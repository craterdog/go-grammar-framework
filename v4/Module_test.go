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
	var syntax = gra.ParseSource(source)
	var actual = gra.FormatSyntax(syntax)
	ass.Equal(t, actual, source)
	gra.ValidateSyntax(syntax)
}

func TestModelGeneration(t *tes.T) {
	// Parse the Syntax.cdsn file.
	var bytes, err = osx.ReadFile(syntaxFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var syntax = gra.ParseSource(source)

	// Generate the AST model.
	source = gra.GenerateAstModel(module, wiki, syntax)
	bytes = []byte(source)
	var filename = "ast/Package.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the AST classes.
	var parser = mod.Parser()
	var model = parser.ParseSource(source)
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

	// Generate the grammar model.
	source = gra.GenerateGrammarModel(module, wiki, syntax)
	bytes = []byte(source)
	filename = "grammar/Package.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the processor class for the grammar.
	source = gra.GenerateProcessorClass(module, syntax)
	bytes = []byte(source)
	filename = "grammar/processor.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the visitor class for the grammar.
	source = gra.GenerateVisitorClass(module, syntax)
	bytes = []byte(source)
	filename = "grammar/visitor.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the token class for the grammar.
	source = gra.GenerateTokenClass(module, syntax)
	bytes = []byte(source)
	filename = "grammar/token.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the scanner class for the grammar.
	source = gra.GenerateScannerClass(module, syntax)
	bytes = []byte(source)
	filename = "grammar/scanner.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the parser class for the grammar.
	source = gra.GenerateParserClass(module, syntax)
	bytes = []byte(source)
	filename = "grammar/parser.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the validator class for the grammar.
	source = gra.GenerateValidatorClass(module, syntax)
	bytes = []byte(source)
	filename = "grammar/validator.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the formatter class for the grammar.
	source = gra.GenerateFormatterClass(module, syntax)
	bytes = []byte(source)
	filename = "grammar/formatter.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func TestLifecycle(t *tes.T) {
	var name = "example"

	// Generate the source code for a new syntax with a default copyright.
	var copyright string
	var source = gra.GenerateSyntaxNotation(name, copyright)

	// Parse the source code for the syntax.
	var syntax = gra.ParseSource(source)

	// Validate the syntax.
	gra.ValidateSyntax(syntax)

	// Format the syntax.
	gra.FormatSyntax(syntax)

	// Generate the AST model for the syntax.
	gra.GenerateAstModel(module, wiki, syntax)

	// Generate the language grammar model for the syntax.
	gra.GenerateGrammarModel(module, wiki, syntax)

	// Generate the formatter class for the syntax.
	gra.GenerateFormatterClass(module, syntax)

	// Generate the parser class for the syntax.
	gra.GenerateParserClass(module, syntax)

	// Generate the scanner class for the syntax.
	gra.GenerateScannerClass(module, syntax)

	// Generate the token class for the syntax.
	gra.GenerateTokenClass(module, syntax)

	// Generate the validator class for the syntax.
	gra.GenerateValidatorClass(module, syntax)
}
