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
Package "ast" provides a parser and formatter for language syntaxes defined
using Crater Dog Syntax Notation™ (CDSN).  The parser performs validation on the
resulting parse tree.  The formatter takes a validated parse tree and generates
the corresponding CDSN source using the canonical format.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-grammar-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	col "github.com/craterdog/go-collection-framework/v4"
)

// Classes

/*
AlternativeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete alternative-like class.
*/
type AlternativeClassLike interface {
	// Constructors
	MakeWithFactors(factors col.ListLike[FactorLike]) AlternativeLike
}

/*
AtomClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete atom-like class.
*/
type AtomClassLike interface {
	// Constructors
	MakeWithGlyph(glyph GlyphLike) AtomLike
	MakeWithIntrinsic(intrinsic string) AtomLike
}

/*
CardinalityClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete cardinality-like class.
*/
type CardinalityClassLike interface {
	// Constructors
	MakeWithConstraint(constraint ConstraintLike) CardinalityLike
}

/*
ConstraintClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constraint-like class.
*/
type ConstraintClassLike interface {
	// Constructors
	MakeWithAttributes(
		first string,
		last string,
	) ConstraintLike
}

/*
DefinitionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete definition-like class.
*/
type DefinitionClassLike interface {
	// Constructors
	MakeWithAttributes(
		comment string,
		name string,
		expression ExpressionLike,
	) DefinitionLike
}

/*
ElementClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete element-like class.
*/
type ElementClassLike interface {
	// Constructors
	MakeWithLiteral(literal string) ElementLike
	MakeWithName(name string) ElementLike
}

/*
ExpressionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete expression-like class.
*/
type ExpressionClassLike interface {
	// Constructors
	MakeWithInline(inline InlineLike) ExpressionLike
	MakeWithMultiline(multiline MultilineLike) ExpressionLike
}

/*
FactorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete factor-like class.
*/
type FactorClassLike interface {
	// Constructors
	MakeWithAttributes(
		predicate PredicateLike,
		cardinality CardinalityLike,
	) FactorLike
}

/*
FilterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete filter-like class.
*/
type FilterClassLike interface {
	// Constructors
	MakeWithAttributes(
		inverted bool,
		atoms col.ListLike[AtomLike],
	) FilterLike
}

/*
GlyphClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete glyph-like class.
*/
type GlyphClassLike interface {
	// Constructors
	MakeWithAttributes(
		first string,
		last string,
	) GlyphLike
}

/*
HeaderClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete header-like class.
*/
type HeaderClassLike interface {
	// Constructors
	MakeWithComment(comment string) HeaderLike
}

/*
InlineClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete inline-like class.
*/
type InlineClassLike interface {
	// Constructors
	MakeWithAttributes(
		alternatives col.ListLike[AlternativeLike],
		note string,
	) InlineLike
}

/*
LineClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete line-like class.
*/
type LineClassLike interface {
	// Constructors
	MakeWithAttributes(
		alternative AlternativeLike,
		note string,
	) LineLike
}

/*
MultilineClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete multiline-like class.
*/
type MultilineClassLike interface {
	// Constructors
	MakeWithLines(lines col.ListLike[LineLike]) MultilineLike
}

/*
PrecedenceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete precedence-like class.
*/
type PrecedenceClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) PrecedenceLike
}

/*
PredicateClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete predicate-like class.
*/
type PredicateClassLike interface {
	// Constructors
	MakeWithAtom(atom AtomLike) PredicateLike
	MakeWithElement(element ElementLike) PredicateLike
	MakeWithFilter(filter FilterLike) PredicateLike
	MakeWithPrecedence(precedence PrecedenceLike) PredicateLike
}

/*
SyntaxClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete syntax-like class.
*/
type SyntaxClassLike interface {
	// Constructors
	MakeWithAttributes(
		headers col.ListLike[HeaderLike],
		definitions col.ListLike[DefinitionLike],
	) SyntaxLike
}

// Instances

/*
AlternativeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete alternative-like class.
*/
type AlternativeLike interface {
	// Attributes
	GetClass() AlternativeClassLike
	GetFactors() col.ListLike[FactorLike]
}

/*
AtomLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete atom-like class.
*/
type AtomLike interface {
	// Attributes
	GetClass() AtomClassLike
	GetGlyph() GlyphLike
	GetIntrinsic() string
}

/*
CardinalityLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete cardinality-like class.
*/
type CardinalityLike interface {
	// Attributes
	GetClass() CardinalityClassLike
	GetConstraint() ConstraintLike
}

/*
ConstraintLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constraint-like class.
*/
type ConstraintLike interface {
	// Attributes
	GetClass() ConstraintClassLike
	GetFirst() string
	GetLast() string
}

/*
DefinitionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete definition-like class.
*/
type DefinitionLike interface {
	// Attributes
	GetClass() DefinitionClassLike
	GetComment() string
	GetName() string
	GetExpression() ExpressionLike
}

/*
ElementLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete element-like class.
*/
type ElementLike interface {
	// Attributes
	GetClass() ElementClassLike
	GetLiteral() string
	GetName() string
}

/*
ExpressionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete expression-like class.
*/
type ExpressionLike interface {
	// Attributes
	GetClass() ExpressionClassLike
	GetInline() InlineLike
	GetMultiline() MultilineLike
}

/*
FactorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete factor-like class.
*/
type FactorLike interface {
	// Attributes
	GetClass() FactorClassLike
	GetPredicate() PredicateLike
	GetCardinality() CardinalityLike
}

/*
FilterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete filter-like class.
*/
type FilterLike interface {
	// Attributes
	GetClass() FilterClassLike
	IsInverted() bool
	GetAtoms() col.ListLike[AtomLike]
}

/*
GlyphLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete glyph-like class.
*/
type GlyphLike interface {
	// Attributes
	GetClass() GlyphClassLike
	GetFirst() string
	GetLast() string
}

/*
HeaderLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete header-like class.
*/
type HeaderLike interface {
	// Attributes
	GetClass() HeaderClassLike
	GetComment() string
}

/*
InlineLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete inline-like class.
*/
type InlineLike interface {
	// Attributes
	GetClass() InlineClassLike
	GetAlternatives() col.ListLike[AlternativeLike]
	GetNote() string
}

/*
LineLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete line-like class.
*/
type LineLike interface {
	// Attributes
	GetClass() LineClassLike
	GetAlternative() AlternativeLike
	GetNote() string
}

/*
MultilineLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete multiline-like class.
*/
type MultilineLike interface {
	// Attributes
	GetClass() MultilineClassLike
	GetLines() col.ListLike[LineLike]
}

/*
PrecedenceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete precedence-like class.
*/
type PrecedenceLike interface {
	// Attributes
	GetClass() PrecedenceClassLike
	GetExpression() ExpressionLike
}

/*
PredicateLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete predicate-like class.
*/
type PredicateLike interface {
	// Attributes
	GetClass() PredicateClassLike
	GetAtom() AtomLike
	GetElement() ElementLike
	GetFilter() FilterLike
	GetPrecedence() PrecedenceLike
}

/*
SyntaxLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete syntax-like class.
*/
type SyntaxLike interface {
	// Attributes
	GetClass() SyntaxClassLike
	GetHeaders() col.ListLike[HeaderLike]
	GetDefinitions() col.ListLike[DefinitionLike]
}
