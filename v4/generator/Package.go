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

/*
Package "generator" provides a template-based code generator that can generate
the class model packages for both the abstract syntax tree (AST) and the
language grammar tools for processing any language defined using Crater Dog
Syntax Notation™ (CDSN).

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-grammar-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package generator

import (
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	mod "github.com/craterdog/go-model-framework/v4"
)

// Classes

/*
GeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all generator-class-like classes.
*/
type GeneratorClassLike interface {
	// Constructors
	Make() GeneratorLike
}

// Instances

/*
GeneratorLike defines the set of aspects and methods that must be supported by
all generator-like instances.
*/
type GeneratorLike interface {
	// Attributes
	GetClass() GeneratorClassLike

	// Methods
	CreateSyntax(
		name string,
		copyright string,
	) ast.SyntaxLike
	GenerateAst(
		module string,
		syntax ast.SyntaxLike,
	) mod.ModelLike
	GenerateGrammar(
		module string,
		syntax ast.SyntaxLike,
	) mod.ModelLike
	GenerateFormatter(
		module string,
		syntax ast.SyntaxLike,
	) string
	GenerateParser(
		module string,
		syntax ast.SyntaxLike,
	) string
	GenerateScanner(
		module string,
		syntax ast.SyntaxLike,
	) string
	GenerateToken(
		module string,
		syntax ast.SyntaxLike,
	) string
	GenerateValidator(
		module string,
		syntax ast.SyntaxLike,
	) string
}
