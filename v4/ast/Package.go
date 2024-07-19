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
Package "ast" provides the abstract syntax tree (AST) classes for this module.
Each AST class manages the attributes associated with the rule definition found
in the syntax grammar with the same rule name as the class.

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
package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// Classes

/*
AlternativeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete alternative-like class.
*/
type AlternativeClassLike interface {
	// Constructors
	Make(parts abs.Sequential[PartLike]) AlternativeLike
}

/*
BoundedClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete bounded-like class.
*/
type BoundedClassLike interface {
	// Constructors
	Make(
		rune_ string,
		optionalExtent ExtentLike,
	) BoundedLike
}

/*
CardinalityClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete cardinality-like class.
*/
type CardinalityClassLike interface {
	// Constructors
	Make(any_ any) CardinalityLike
}

/*
CharacterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete character-like class.
*/
type CharacterClassLike interface {
	// Constructors
	Make(any_ any) CharacterLike
}

/*
ConstrainedClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constrained-like class.
*/
type ConstrainedClassLike interface {
	// Constructors
	Make(
		number string,
		optionalLimit LimitLike,
	) ConstrainedLike
}

/*
DefinitionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete definition-like class.
*/
type DefinitionClassLike interface {
	// Constructors
	Make(any_ any) DefinitionLike
}

/*
ElementClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete element-like class.
*/
type ElementClassLike interface {
	// Constructors
	Make(any_ any) ElementLike
}

/*
ExpressionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete expression-like class.
*/
type ExpressionClassLike interface {
	// Constructors
	Make(
		optionalComment string,
		lowercase string,
		pattern PatternLike,
		optionalNote string,
	) ExpressionLike
}

/*
ExtentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete extent-like class.
*/
type ExtentClassLike interface {
	// Constructors
	Make(rune_ string) ExtentLike
}

/*
FactorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete factor-like class.
*/
type FactorClassLike interface {
	// Constructors
	Make(
		predicate PredicateLike,
		optionalCardinality CardinalityLike,
	) FactorLike
}

/*
FilteredClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete filtered-like class.
*/
type FilteredClassLike interface {
	// Constructors
	Make(
		optionalNegation string,
		characters abs.Sequential[CharacterLike],
	) FilteredLike
}

/*
GroupedClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete grouped-like class.
*/
type GroupedClassLike interface {
	// Constructors
	Make(pattern PatternLike) GroupedLike
}

/*
HeaderClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete header-like class.
*/
type HeaderClassLike interface {
	// Constructors
	Make(comment string) HeaderLike
}

/*
IdentifierClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete identifier-like class.
*/
type IdentifierClassLike interface {
	// Constructors
	Make(any_ any) IdentifierLike
}

/*
InlinedClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete inlined-like class.
*/
type InlinedClassLike interface {
	// Constructors
	Make(
		factors abs.Sequential[FactorLike],
		optionalNote string,
	) InlinedLike
}

/*
LimitClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete limit-like class.
*/
type LimitClassLike interface {
	// Constructors
	Make(optionalNumber string) LimitLike
}

/*
LineClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete line-like class.
*/
type LineClassLike interface {
	// Constructors
	Make(
		identifier IdentifierLike,
		optionalNote string,
	) LineLike
}

/*
MultilinedClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete multilined-like class.
*/
type MultilinedClassLike interface {
	// Constructors
	Make(lines abs.Sequential[LineLike]) MultilinedLike
}

/*
PartClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete part-like class.
*/
type PartClassLike interface {
	// Constructors
	Make(
		element ElementLike,
		optionalCardinality CardinalityLike,
	) PartLike
}

/*
PatternClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete pattern-like class.
*/
type PatternClassLike interface {
	// Constructors
	Make(
		parts abs.Sequential[PartLike],
		alternatives abs.Sequential[AlternativeLike],
	) PatternLike
}

/*
PredicateClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete predicate-like class.
*/
type PredicateClassLike interface {
	// Constructors
	Make(any_ any) PredicateLike
}

/*
RuleClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete rule-like class.
*/
type RuleClassLike interface {
	// Constructors
	Make(
		optionalComment string,
		uppercase string,
		definition DefinitionLike,
	) RuleLike
}

/*
StringClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete string-like class.
*/
type StringClassLike interface {
	// Constructors
	Make(any_ any) StringLike
}

/*
SyntaxClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete syntax-like class.
*/
type SyntaxClassLike interface {
	// Constructors
	Make(
		headers abs.Sequential[HeaderLike],
		rules abs.Sequential[RuleLike],
		expressions abs.Sequential[ExpressionLike],
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
	GetParts() abs.Sequential[PartLike]
}

/*
BoundedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete bounded-like class.
*/
type BoundedLike interface {
	// Attributes
	GetClass() BoundedClassLike
	GetRune() string
	GetOptionalExtent() ExtentLike
}

/*
CardinalityLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete cardinality-like class.
*/
type CardinalityLike interface {
	// Attributes
	GetClass() CardinalityClassLike
	GetAny() any
}

/*
CharacterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete character-like class.
*/
type CharacterLike interface {
	// Attributes
	GetClass() CharacterClassLike
	GetAny() any
}

/*
ConstrainedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constrained-like class.
*/
type ConstrainedLike interface {
	// Attributes
	GetClass() ConstrainedClassLike
	GetNumber() string
	GetOptionalLimit() LimitLike
}

/*
DefinitionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete definition-like class.
*/
type DefinitionLike interface {
	// Attributes
	GetClass() DefinitionClassLike
	GetAny() any
}

/*
ElementLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete element-like class.
*/
type ElementLike interface {
	// Attributes
	GetClass() ElementClassLike
	GetAny() any
}

/*
ExpressionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete expression-like class.
*/
type ExpressionLike interface {
	// Attributes
	GetClass() ExpressionClassLike
	GetOptionalComment() string
	GetLowercase() string
	GetPattern() PatternLike
	GetOptionalNote() string
}

/*
ExtentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete extent-like class.
*/
type ExtentLike interface {
	// Attributes
	GetClass() ExtentClassLike
	GetRune() string
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
	GetOptionalCardinality() CardinalityLike
}

/*
FilteredLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete filtered-like class.
*/
type FilteredLike interface {
	// Attributes
	GetClass() FilteredClassLike
	GetOptionalNegation() string
	GetCharacters() abs.Sequential[CharacterLike]
}

/*
GroupedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete grouped-like class.
*/
type GroupedLike interface {
	// Attributes
	GetClass() GroupedClassLike
	GetPattern() PatternLike
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
IdentifierLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete identifier-like class.
*/
type IdentifierLike interface {
	// Attributes
	GetClass() IdentifierClassLike
	GetAny() any
}

/*
InlinedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete inlined-like class.
*/
type InlinedLike interface {
	// Attributes
	GetClass() InlinedClassLike
	GetFactors() abs.Sequential[FactorLike]
	GetOptionalNote() string
}

/*
LimitLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete limit-like class.
*/
type LimitLike interface {
	// Attributes
	GetClass() LimitClassLike
	GetOptionalNumber() string
}

/*
LineLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete line-like class.
*/
type LineLike interface {
	// Attributes
	GetClass() LineClassLike
	GetIdentifier() IdentifierLike
	GetOptionalNote() string
}

/*
MultilinedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete multilined-like class.
*/
type MultilinedLike interface {
	// Attributes
	GetClass() MultilinedClassLike
	GetLines() abs.Sequential[LineLike]
}

/*
PartLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete part-like class.
*/
type PartLike interface {
	// Attributes
	GetClass() PartClassLike
	GetElement() ElementLike
	GetOptionalCardinality() CardinalityLike
}

/*
PatternLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete pattern-like class.
*/
type PatternLike interface {
	// Attributes
	GetClass() PatternClassLike
	GetParts() abs.Sequential[PartLike]
	GetAlternatives() abs.Sequential[AlternativeLike]
}

/*
PredicateLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete predicate-like class.
*/
type PredicateLike interface {
	// Attributes
	GetClass() PredicateClassLike
	GetAny() any
}

/*
RuleLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete rule-like class.
*/
type RuleLike interface {
	// Attributes
	GetClass() RuleClassLike
	GetOptionalComment() string
	GetUppercase() string
	GetDefinition() DefinitionLike
}

/*
StringLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete string-like class.
*/
type StringLike interface {
	// Attributes
	GetClass() StringClassLike
	GetAny() any
}

/*
SyntaxLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete syntax-like class.
*/
type SyntaxLike interface {
	// Attributes
	GetClass() SyntaxClassLike
	GetHeaders() abs.Sequential[HeaderLike]
	GetRules() abs.Sequential[RuleLike]
	GetExpressions() abs.Sequential[ExpressionLike]
}
