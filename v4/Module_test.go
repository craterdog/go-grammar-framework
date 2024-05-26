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
	tes "testing"
)

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

	// Parse the source code for the syntax.
	var parser = gra.Parser()
	syntax = parser.ParseSource(source)

	// Generate the AST model for the syntax.
	generator.GenerateAST(module, syntax)

	// Generate the agent model for the syntax.
	var model = generator.GenerateAgent(module, syntax)

	// Generate the formatter class for the syntax.
	source = generator.GenerateFormatter(module, syntax, model)
	fmt.Printf("FORMATTER CLASS: %v\n", source)

	// Generate the parser class for the syntax.
	source = generator.GenerateParser(module, syntax, model)
	fmt.Printf("PARSER CLASS: %v\n", source)

	// Generate the scanner class for the syntax.
	source = generator.GenerateScanner(module, syntax, model)
	fmt.Printf("SCANNER CLASS: %v\n", source)

	// Generate the token class for the syntax.
	source = generator.GenerateToken(module, syntax, model)
	fmt.Printf("TOKEN CLASS: %v\n", source)

	// Generate the validator class for the syntax.
	source = generator.GenerateValidator(module, syntax, model)
	fmt.Printf("VALIDATOR CLASS: %v\n", source)
}
