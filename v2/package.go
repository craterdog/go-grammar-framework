/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

/*
Package grammars provides a parser and formatter for language grammars defined
using Crater Dog Syntax Notation™ (CDSN).  The parser performs validation on the
resulting parse tree.  The formatter takes a validated parse tree and generates
the corresponding CDSN source using the canonical format.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-grammar-framework/wiki

This package follows the Crater Dog Technologies™ (craterdog) Go Coding
Conventions located here:
  - https://github.com/craterdog/go-class-framework/wiki

Additional implementations of the classes provided by this package can be
developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types; and the class implementations only depend
on interfaces, not on each other.
*/
package grammars

import (
	cla "github.com/craterdog/go-class-framework/v2"
	col "github.com/craterdog/go-collection-framework/v3"
)

// TYPES

// Specializations

/*
TokenType is a specialized type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	CharacterToken
	CommentToken
	DelimiterToken
	EOFToken
	EOLToken
	IntrinsicToken
	LiteralToken
	NameToken
	NoteToken
	NumberToken
	SpaceToken
	SymbolToken
)

// INTERFACES

// Classes

/*
AlternativeClassLike defines the set of class constants, constructors and
functions that must be supported by all alternative-class-like classes.
*/
type AlternativeClassLike interface {
	// Constructors
	MakeWithAttributes(factors_ col.Sequential[FactorLike], note_ string) AlternativeLike
}

/*
AssertionClassLike defines the set of class constants, constructors and
functions that must be supported by all assertion-class-like classes.
*/
type AssertionClassLike interface {
	// Constructors
	MakeWithElement(element_ ElementLike) AssertionLike
	MakeWithGlyph(glyph_ GlyphLike) AssertionLike
	MakeWithPrecedence(precedence_ PrecedenceLike) AssertionLike
}

/*
CardinalityClassLike defines the set of class constants, constructors and
functions that must be supported by all cardinality-class-like classes.
*/
type CardinalityClassLike interface {
	// Constructors
	MakeWithConstraint(constraint_ ConstraintLike) CardinalityLike
}

/*
ConstraintClassLike defines the set of class constants, constructors and
functions that must be supported by all constraint-class-like classes.
*/
type ConstraintClassLike interface {
	// Constructors
	MakeWithRange(first_ string, last_ string) ConstraintLike
}

/*
DefinitionClassLike defines the set of class constants, constructors and
functions that must be supported by all definition-class-like classes.
*/
type DefinitionClassLike interface {
	// Constructors
	MakeWithAttributes(symbol_ string, expression_ ExpressionLike) DefinitionLike
}

/*
ElementClassLike defines the set of class constants, constructors and functions
that must be supported by all element-class-like classes.
*/
type ElementClassLike interface {
	// Constructors
	MakeWithIntrinsic(intrinsic_ string) ElementLike
	MakeWithLiteral(literal_ string) ElementLike
	MakeWithName(name_ string) ElementLike
}

/*
ExpressionClassLike defines the set of class constants, constructors and
functions that must be supported by all expression-class-like classes.
*/
type ExpressionClassLike interface {
	// Constructors
	MakeWithAttributes(alternatives_ col.Sequential[AlternativeLike], multilined_ bool) ExpressionLike
}

/*
FactorClassLike defines the set of class constants, constructors and functions
that must be supported by all factor-class-like classes.
*/
type FactorClassLike interface {
	// Constructors
	MakeWithAttributes(predicate_ PredicateLike, cardinality_ CardinalityLike) FactorLike
}

/*
FormatterClassLike defines the set of class constants, constructors and
functions that must be supported by all formatter-class-like classes.
*/
type FormatterClassLike interface {
	// Constructors
	Make() FormatterLike
}

/*
GeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all generator-class-like classes.
*/
type GeneratorClassLike interface {
	// Constructors
	Make() GeneratorLike
}

/*
GlyphClassLike defines the set of class constants, constructors and functions
that must be supported by all glyph-class-like classes.
*/
type GlyphClassLike interface {
	// Constructors
	MakeWithRange(first_ string, last_ string) GlyphLike
}

/*
GrammarClassLike defines the set of class constants, constructors and functions
that must be supported by all grammar-class-like classes.
*/
type GrammarClassLike interface {
	// Constructors
	MakeWithStatements(statements_ col.Sequential[StatementLike]) GrammarLike
}

/*
ParserClassLike defines the set of class constants, constructors and functions
that must be supported by all parser-class-like classes.
*/
type ParserClassLike interface {
	// Constructors
	Make() ParserLike
}

/*
PrecedenceClassLike defines the set of class constants, constructors and
functions that must be supported by all precedence-class-like classes.
*/
type PrecedenceClassLike interface {
	// Constructors
	MakeWithExpression(expression_ ExpressionLike) PrecedenceLike
}

/*
PredicateClassLike defines the set of class constants, constructors and
functions that must be supported by all predicate-class-like classes.
*/
type PredicateClassLike interface {
	// Constructors
	MakeWithAttributes(assertion_ AssertionLike, inverted_ bool) PredicateLike
}

/*
ScannerClassLike defines the set of class constants, constructors and functions
that must be supported by all scanner-class-like classes.
*/
type ScannerClassLike interface {
	// Constructors
	Make(source string, tokens col.QueueLike[TokenLike]) ScannerLike

	// Functions
	MatchToken(tokenType TokenType, text string) col.ListLike[string]
}

/*
StatementClassLike defines the set of class constants, constructors and
functions that must be supported by all statement-class-like classes.
*/
type StatementClassLike interface {
	// Constructors
	MakeWithComment(comment_ string) StatementLike
	MakeWithDefinition(definition_ DefinitionLike) StatementLike
}

/*
TokenClassLike defines the set of class constants, constructors and functions
that must be supported by all token-class-like classes.
*/
type TokenClassLike interface {
	// Constructors
	MakeWithAttributes(
		line_ int,
		position_ int,
		type_ TokenType,
		value_ string,
	) TokenLike

	// Functions
	AsString(tokenType TokenType) string
}

/*
ValidatorClassLike defines the set of class constants, constructors and
functions that must be supported by all validator-class-like classes.
*/
type ValidatorClassLike interface {
	// Constructors
	Make() ValidatorLike
}

// Instances

/*
AlternativeLike defines the set of aspects and methods that must be supported by
all alternative-like instances.
*/
type AlternativeLike interface {
	// Attributes
	GetFactors() col.Sequential[FactorLike]
	GetNote() string
}

/*
AssertionLike defines the set of aspects and methods that must be supported by
all assertion-like instances.
*/
type AssertionLike interface {
	// Attributes
	GetElement() ElementLike
	GetGlyph() GlyphLike
	GetPrecedence() PrecedenceLike
}

/*
CardinalityLike defines the set of aspects and methods that must be supported by
all cardinality-like instances.
*/
type CardinalityLike interface {
	// Attributes
	GetConstraint() ConstraintLike
}

/*
ConstraintLike defines the set of aspects and methods that must be supported by
all constraint-like instances.
*/
type ConstraintLike interface {
	// Attributes
	GetFirst() string
	GetLast() string
}

/*
DefinitionLike defines the set of aspects and methods that must be supported by
all definition-like instances.
*/
type DefinitionLike interface {
	// Attributes
	GetExpression() ExpressionLike
	GetSymbol() string
}

/*
ElementLike defines the set of aspects and methods that must be supported by all
element-like instances.
*/
type ElementLike interface {
	// Attributes
	GetIntrinsic() string
	GetLiteral() string
	GetName() string
}

/*
ExpressionLike defines the set of aspects and methods that must be supported by
all expression-like instances.
*/
type ExpressionLike interface {
	// Attributes
	GetAlternatives() col.Sequential[AlternativeLike]
	GetMultilined() bool

	// Methods
	IsMultilined() bool
}

/*
FactorLike defines the set of aspects and methods that must be supported by all
factor-like instances.
*/
type FactorLike interface {
	// Attributes
	GetCardinality() CardinalityLike
	GetPredicate() PredicateLike
}

/*
FormatterLike defines the set of aspects and methods that must be supported by
all formatter-like instances.
*/
type FormatterLike interface {
	// Methods
	FormatDefinition(definition DefinitionLike) string
	FormatGrammar(grammar GrammarLike) string
}

/*
GeneratorLike defines the set of aspects and methods that must be supported by
all generator-like instances.
*/
type GeneratorLike interface {
	// Methods
	GeneratePackage(
		name string,
		license string,
		comment string,
		grammar GrammarLike,
	) cla.GoPNLike
}

/*
GlyphLike defines the set of aspects and methods that must be supported by all
glyph-like instances.
*/
type GlyphLike interface {
	// Attributes
	GetFirst() string
	GetLast() string
}

/*
GrammarLike defines the set of aspects and methods that must be supported by all
grammar-like instances.
*/
type GrammarLike interface {
	// Attributes
	GetStatements() col.Sequential[StatementLike]
}

/*
ParserLike defines the set of aspects and methods that must be supported by all
parser-like instances.
*/
type ParserLike interface {
	// Methods
	ParseSource(source string) GrammarLike
}

/*
PrecedenceLike defines the set of aspects and methods that must be supported by
all precedence-like instances.
*/
type PrecedenceLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
PredicateLike defines the set of aspects and methods that must be supported by
all predicate-like instances.
*/
type PredicateLike interface {
	// Attributes
	GetAssertion() AssertionLike
	GetInverted() bool

	// Methods
	IsInverted() bool
}

/*
ScannerLike defines the set of aspects and methods that must be supported by all
scanner-like instances.
*/
type ScannerLike interface {
}

/*
StatementLike defines the set of aspects and methods that must be supported by
all statement-like instances.
*/
type StatementLike interface {
	// Attributes
	GetComment() string
	GetDefinition() DefinitionLike
}

/*
TokenLike defines the set of aspects and methods that must be supported by all
token-like instances.
*/
type TokenLike interface {
	// Attributes
	GetLine() int
	GetPosition() int
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike defines the set of aspects and methods that must be supported by
all validator-like instances.
*/
type ValidatorLike interface {
	// Methods
	ValidateGrammar(grammar GrammarLike)
}
