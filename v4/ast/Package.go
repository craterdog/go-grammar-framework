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
Package "ast" provides...

For detailed documentation on this package refer to the wiki:
  - <wiki>

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
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// Classes

/*
AlternativeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete alternative-like class.
*/
type AlternativeClassLike interface {
	// Constructors
	Make(parts col.ListLike[PartLike]) AlternativeLike
}

/*
BoundedClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete bounded-like class.
*/
type BoundedClassLike interface {
	// Constructors
	Make(
		initial InitialLike,
		extent ExtentLike,
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
		minimum MinimumLike,
		maximum MaximumLike,
	) ConstrainedLike
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
	Make(any_ any) ExpressionLike
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
		cardinality CardinalityLike,
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
		negation string,
		characters col.ListLike[CharacterLike],
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
InitialClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete initial-like class.
*/
type InitialClassLike interface {
	// Constructors
	Make(rune_ string) InitialLike
}

/*
InlinedClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete inlined-like class.
*/
type InlinedClassLike interface {
	// Constructors
	Make(
		factors col.ListLike[FactorLike],
		note string,
	) InlinedLike
}

/*
LexigramClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete lexigram-like class.
*/
type LexigramClassLike interface {
	// Constructors
	Make(
		comment string,
		lowercase string,
		pattern PatternLike,
		note string,
	) LexigramLike
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
		note string,
	) LineLike
}

/*
MaximumClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete maximum-like class.
*/
type MaximumClassLike interface {
	// Constructors
	Make(number string) MaximumLike
}

/*
MinimumClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete minimum-like class.
*/
type MinimumClassLike interface {
	// Constructors
	Make(number string) MinimumLike
}

/*
MultilinedClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete multilined-like class.
*/
type MultilinedClassLike interface {
	// Constructors
	Make(lines col.ListLike[LineLike]) MultilinedLike
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
		cardinality CardinalityLike,
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
		parts col.ListLike[PartLike],
		alternatives col.ListLike[AlternativeLike],
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
		comment string,
		uppercase string,
		expression ExpressionLike,
	) RuleLike
}

/*
SyntaxClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete syntax-like class.
*/
type SyntaxClassLike interface {
	// Constructors
	Make(
		headers col.ListLike[HeaderLike],
		rules col.ListLike[RuleLike],
		lexigrams col.ListLike[LexigramLike],
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
	GetParts() col.ListLike[PartLike]
}

/*
BoundedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete bounded-like class.
*/
type BoundedLike interface {
	// Attributes
	GetClass() BoundedClassLike
	GetInitial() InitialLike
	GetExtent() ExtentLike
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
	GetMinimum() MinimumLike
	GetMaximum() MaximumLike
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
	GetAny() any
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
	GetCardinality() CardinalityLike
}

/*
FilteredLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete filtered-like class.
*/
type FilteredLike interface {
	// Attributes
	GetClass() FilteredClassLike
	GetNegation() string
	GetCharacters() col.ListLike[CharacterLike]
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
InitialLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete initial-like class.
*/
type InitialLike interface {
	// Attributes
	GetClass() InitialClassLike
	GetRune() string
}

/*
InlinedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete inlined-like class.
*/
type InlinedLike interface {
	// Attributes
	GetClass() InlinedClassLike
	GetFactors() col.ListLike[FactorLike]
	GetNote() string
}

/*
LexigramLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete lexigram-like class.
*/
type LexigramLike interface {
	// Attributes
	GetClass() LexigramClassLike
	GetComment() string
	GetLowercase() string
	GetPattern() PatternLike
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
	GetIdentifier() IdentifierLike
	GetNote() string
}

/*
MaximumLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete maximum-like class.
*/
type MaximumLike interface {
	// Attributes
	GetClass() MaximumClassLike
	GetNumber() string
}

/*
MinimumLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete minimum-like class.
*/
type MinimumLike interface {
	// Attributes
	GetClass() MinimumClassLike
	GetNumber() string
}

/*
MultilinedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete multilined-like class.
*/
type MultilinedLike interface {
	// Attributes
	GetClass() MultilinedClassLike
	GetLines() col.ListLike[LineLike]
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
	GetCardinality() CardinalityLike
}

/*
PatternLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete pattern-like class.
*/
type PatternLike interface {
	// Attributes
	GetClass() PatternClassLike
	GetParts() col.ListLike[PartLike]
	GetAlternatives() col.ListLike[AlternativeLike]
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
	GetComment() string
	GetUppercase() string
	GetExpression() ExpressionLike
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
	GetRules() col.ListLike[RuleLike]
	GetLexigrams() col.ListLike[LexigramLike]
}
