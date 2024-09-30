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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
)

// Classes

/*
AnalyzerClassLike defines the set of class constants, constructors and
functions that must be supported by all analyzer-class-like classes.
*/
type AnalyzerClassLike interface {
	// Constructor
	Make() AnalyzerLike
}

/*
FormatterClassLike defines the set of class constants, constructors and
functions that must be supported by all formatter-class-like classes.
*/
type FormatterClassLike interface {
	// Constructor
	Make() FormatterLike
}

/*
GrammarClassLike defines the set of class constants, constructors and
functions that must be supported by all grammar-class-like classes.
*/
type GrammarClassLike interface {
	// Constructor
	Make() GrammarLike
}

/*
AstClassLike defines the set of class constants, constructors and
functions that must be supported by all ast-class-like classes.
*/
type AstClassLike interface {
	// Constructor
	Make() AstLike
}

/*
ParserClassLike defines the set of class constants, constructors and
functions that must be supported by all parser-class-like classes.
*/
type ParserClassLike interface {
	// Constructor
	Make() ParserLike
}

/*
ProcessorClassLike defines the set of class constants, constructors and
functions that must be supported by all processor-class-like classes.
*/
type ProcessorClassLike interface {
	// Constructor
	Make() ProcessorLike
}

/*
ScannerClassLike defines the set of class constants, constructors and
functions that must be supported by all scanner-class-like classes.
*/
type ScannerClassLike interface {
	// Constructor
	Make() ScannerLike
}

/*
SyntaxClassLike defines the set of class constants, constructors and
functions that must be supported by all syntax-class-like classes.
*/
type SyntaxClassLike interface {
	// Constructor
	Make() SyntaxLike
}

/*
TokenClassLike defines the set of class constants, constructors and
functions that must be supported by all token-class-like classes.
*/
type TokenClassLike interface {
	// Constructor
	Make() TokenLike
}

/*
ValidatorClassLike defines the set of class constants, constructors and
functions that must be supported by all validator-class-like classes.
*/
type ValidatorClassLike interface {
	// Constructor
	Make() ValidatorLike
}

/*
VisitorClassLike defines the set of class constants, constructors and
functions that must be supported by all visitor-class-like classes.
*/
type VisitorClassLike interface {
	// Constructor
	Make() VisitorLike
}

// Instances

/*
AnalyzerLike defines the set of aspects and methods that must be supported by
all analyzer-like instances.
*/
type AnalyzerLike interface {
	// Public
	GetClass() AnalyzerClassLike
	AnalyzeSyntax(syntax ast.SyntaxLike)
	GetExpressions() abs.Sequential[abs.AssociationLike[string, string]]
	GetIdentifiers(ruleName string) abs.Sequential[ast.IdentifierLike]
	GetNotice() string
	GetReferences(ruleName string) abs.Sequential[ast.ReferenceLike]
	GetRuleNames() abs.Sequential[string]
	GetSyntaxMap() string
	GetSyntaxName() string
	GetTerms(ruleName string) abs.Sequential[ast.TermLike]
	GetTokenNames() abs.Sequential[string]
	IsDelimited(ruleName string) bool
	IsPlural(name string) bool

	// Aspect
	gra.Methodical
}

/*
FormatterLike defines the set of aspects and methods that must be supported by
all formatter-like instances.
*/
type FormatterLike interface {
	// Public
	GetClass() FormatterClassLike
	GenerateFormatterClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
GrammarLike defines the set of aspects and methods that must be supported by
all grammar-like instances.
*/
type GrammarLike interface {
	// Public
	GetClass() GrammarClassLike
	GenerateGrammarModel(
		module string,
		wiki string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
AstLike defines the set of aspects and methods that must be supported by
all ast-like instances.
*/
type AstLike interface {
	// Public
	GetClass() AstClassLike
	GenerateAstModel(
		wiki string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
ParserLike defines the set of aspects and methods that must be supported by
all parser-like instances.
*/
type ParserLike interface {
	// Public
	GetClass() ParserClassLike
	GenerateParserClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
ProcessorLike defines the set of aspects and methods that must be supported by
all processor-like instances.
*/
type ProcessorLike interface {
	// Public
	GetClass() ProcessorClassLike
	GenerateProcessorClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
ScannerLike defines the set of aspects and methods that must be supported by
all scanner-like instances.
*/
type ScannerLike interface {
	// Public
	GetClass() ScannerClassLike
	GenerateScannerClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
SyntaxLike defines the set of aspects and methods that must be supported by
all syntax-like instances.
*/
type SyntaxLike interface {
	// Public
	GetClass() SyntaxClassLike
	GenerateSyntaxNotation(
		syntax string,
		copyright string,
	) (
		implementation string,
	)
}

/*
TokenLike defines the set of aspects and methods that must be supported by
all token-like instances.
*/
type TokenLike interface {
	// Public
	GetClass() TokenClassLike
	GenerateTokenClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
ValidatorLike defines the set of aspects and methods that must be supported by
all validator-like instances.
*/
type ValidatorLike interface {
	// Public
	GetClass() ValidatorClassLike
	GenerateValidatorClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
VisitorLike defines the set of aspects and methods that must be supported by
all visitor-like instances.
*/
type VisitorLike interface {
	// Public
	GetClass() VisitorClassLike
	GenerateVisitorClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}
