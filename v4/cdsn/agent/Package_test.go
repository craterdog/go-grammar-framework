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
	tes "testing"
)

func TestLifecycle(t *tes.T) {
	var generator = age.Generator().Make()
	var name = "example"

	// Generate a new syntax with a default copyright.
	var copyright string
	var syntax = generator.CreateSyntax(name, copyright)

	// Validate the syntax.
	var validator = age.Validator().Make()
	validator.ValidateSyntax(syntax)

	// Format the syntax.
	var formatter = age.Formatter().Make()
	var source = formatter.FormatSyntax(syntax)

	// Parse the source code for the syntax.
	var parser = age.Parser().Make()
	syntax = parser.ParseSource(source)

	// Generate the AST model for the syntax.
	generator.GenerateAST(syntax)

	// Generate the agent model for the syntax.
	var model = generator.GenerateAgent(syntax)

	// Generate the formatter class for the syntax.
	source = generator.GenerateFormatter(model)
	fmt.Printf("FORMATTER CLASS: %v\n", source)

	// Generate the parser class for the syntax.
	source = generator.GenerateParser(model)
	fmt.Printf("PARSER CLASS: %v\n", source)

	// Generate the scanner class for the syntax.
	source = generator.GenerateScanner(model)
	fmt.Printf("SCANNER CLASS: %v\n", source)

	// Generate the token class for the syntax.
	source = generator.GenerateToken(model)
	fmt.Printf("TOKEN CLASS: %v\n", source)

	// Generate the validator class for the syntax.
	source = generator.GenerateValidator(model)
	fmt.Printf("VALIDATOR CLASS: %v\n", source)
}
