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
	Make(repetitions abs.Sequential[RepetitionLike]) AlternativeLike
}

/*
BracketClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete bracket-like class.
*/
type BracketClassLike interface {
	// Constructors
	Make(
		factors abs.Sequential[FactorLike],
		cardinality CardinalityLike,
	) BracketLike
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
ConstraintClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constraint-like class.
*/
type ConstraintClassLike interface {
	// Constructors
	Make(any_ any) ConstraintLike
}

/*
CountClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete count-like class.
*/
type CountClassLike interface {
	// Constructors
	Make(numbers abs.Sequential[string]) CountLike
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
		newlines abs.Sequential[string],
	) ExpressionLike
}

/*
FactorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete factor-like class.
*/
type FactorClassLike interface {
	// Constructors
	Make(any_ any) FactorLike
}

/*
FilterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete filter-like class.
*/
type FilterClassLike interface {
	// Constructors
	Make(
		optionalExcluded string,
		characters abs.Sequential[CharacterLike],
	) FilterLike
}

/*
GroupClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete group-like class.
*/
type GroupClassLike interface {
	// Constructors
	Make(pattern PatternLike) GroupLike
}

/*
HeaderClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete header-like class.
*/
type HeaderClassLike interface {
	// Constructors
	Make(
		comment string,
		newline string,
	) HeaderLike
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
InlineClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete inline-like class.
*/
type InlineClassLike interface {
	// Constructors
	Make(
		terms abs.Sequential[TermLike],
		optionalNote string,
	) InlineLike
}

/*
LineClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete line-like class.
*/
type LineClassLike interface {
	// Constructors
	Make(
		newline string,
		identifier IdentifierLike,
		optionalNote string,
	) LineLike
}

/*
MultilineClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete multiline-like class.
*/
type MultilineClassLike interface {
	// Constructors
	Make(lines abs.Sequential[LineLike]) MultilineLike
}

/*
PatternClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete pattern-like class.
*/
type PatternClassLike interface {
	// Constructors
	Make(alternatives abs.Sequential[AlternativeLike]) PatternLike
}

/*
ReferenceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete reference-like class.
*/
type ReferenceClassLike interface {
	// Constructors
	Make(
		identifier IdentifierLike,
		optionalCardinality CardinalityLike,
	) ReferenceLike
}

/*
RepetitionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete repetition-like class.
*/
type RepetitionClassLike interface {
	// Constructors
	Make(
		element ElementLike,
		optionalCardinality CardinalityLike,
	) RepetitionLike
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
		newlines abs.Sequential[string],
	) RuleLike
}

/*
SpecificClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete specific-like class.
*/
type SpecificClassLike interface {
	// Constructors
	Make(runics abs.Sequential[string]) SpecificLike
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

/*
TermClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete term-like class.
*/
type TermClassLike interface {
	// Constructors
	Make(any_ any) TermLike
}

/*
TextClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete text-like class.
*/
type TextClassLike interface {
	// Constructors
	Make(any_ any) TextLike
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
	GetRepetitions() abs.Sequential[RepetitionLike]
}

/*
BracketLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete bracket-like class.
*/
type BracketLike interface {
	// Attributes
	GetClass() BracketClassLike
	GetFactors() abs.Sequential[FactorLike]
	GetCardinality() CardinalityLike
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
ConstraintLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constraint-like class.
*/
type ConstraintLike interface {
	// Attributes
	GetClass() ConstraintClassLike
	GetAny() any
}

/*
CountLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete count-like class.
*/
type CountLike interface {
	// Attributes
	GetClass() CountClassLike
	GetNumbers() abs.Sequential[string]
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
	GetNewlines() abs.Sequential[string]
}

/*
FactorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete factor-like class.
*/
type FactorLike interface {
	// Attributes
	GetClass() FactorClassLike
	GetAny() any
}

/*
FilterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete filter-like class.
*/
type FilterLike interface {
	// Attributes
	GetClass() FilterClassLike
	GetOptionalExcluded() string
	GetCharacters() abs.Sequential[CharacterLike]
}

/*
GroupLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete group-like class.
*/
type GroupLike interface {
	// Attributes
	GetClass() GroupClassLike
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
	GetNewline() string
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
InlineLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete inline-like class.
*/
type InlineLike interface {
	// Attributes
	GetClass() InlineClassLike
	GetTerms() abs.Sequential[TermLike]
	GetOptionalNote() string
}

/*
LineLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete line-like class.
*/
type LineLike interface {
	// Attributes
	GetClass() LineClassLike
	GetNewline() string
	GetIdentifier() IdentifierLike
	GetOptionalNote() string
}

/*
MultilineLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete multiline-like class.
*/
type MultilineLike interface {
	// Attributes
	GetClass() MultilineClassLike
	GetLines() abs.Sequential[LineLike]
}

/*
PatternLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete pattern-like class.
*/
type PatternLike interface {
	// Attributes
	GetClass() PatternClassLike
	GetAlternatives() abs.Sequential[AlternativeLike]
}

/*
ReferenceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete reference-like class.
*/
type ReferenceLike interface {
	// Attributes
	GetClass() ReferenceClassLike
	GetIdentifier() IdentifierLike
	GetOptionalCardinality() CardinalityLike
}

/*
RepetitionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete repetition-like class.
*/
type RepetitionLike interface {
	// Attributes
	GetClass() RepetitionClassLike
	GetElement() ElementLike
	GetOptionalCardinality() CardinalityLike
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
	GetNewlines() abs.Sequential[string]
}

/*
SpecificLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete specific-like class.
*/
type SpecificLike interface {
	// Attributes
	GetClass() SpecificClassLike
	GetRunics() abs.Sequential[string]
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

/*
TermLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete term-like class.
*/
type TermLike interface {
	// Attributes
	GetClass() TermClassLike
	GetAny() any
}

/*
TextLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete text-like class.
*/
type TextLike interface {
	// Attributes
	GetClass() TextClassLike
	GetAny() any
}
