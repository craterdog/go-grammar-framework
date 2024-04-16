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
Package "test" provides...

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
package test

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	AnythingToken
	CharacterToken
	DelimiterToken
	EOFToken
	EOLToken
	IntegerToken
	SpaceToken
	TextToken
)

// Classes

/*
DefaultClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete default-like class.
*/
type DefaultClassLike interface {
	// Constructors
	Make() DefaultLike
}

/*
ExampleClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete example-like class.
*/
type ExampleClassLike interface {
	// Constructors
	MakeWithDefault(default_ DefaultLike) ExampleLike
	MakeWithPrimitive(primitive PrimitiveLike) ExampleLike
	MakeWithLists(lists col.ListLike[ListLike]) ExampleLike
}

/*
FormatterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructors
	Make() FormatterLike
}

/*
ListClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete list-like class.
*/
type ListClassLike interface {
	// Constructors
	MakeWithExamples(examples col.ListLike[ExampleLike]) ListLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructors
	Make() ParserLike
}

/*
PrimitiveClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete primitive-like class.
*/
type PrimitiveClassLike interface {
	// Constructors
	MakeWithCharacter(character string) PrimitiveLike
	MakeWithText(text string) PrimitiveLike
	MakeWithInteger(integer string) PrimitiveLike
	MakeWithAnything(anything string) PrimitiveLike
}

/*
ScannerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

MatchToken() a list of strings representing any matches found in the specified
text of the specified token type using the regular expression defined for that
token type.  If the regular expression contains submatch patterns the matching
substrings are returned as additional values in the list.
*/
type ScannerClassLike interface {
	// Constructors
	Make(
		source string,
		tokens col.QueueLike[TokenLike],
	) ScannerLike

	// Functions
	FormatToken(token TokenLike) string
	MatchToken(
		type_ TokenType,
		text string,
	) col.ListLike[string]
}

/*
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructors
	MakeWithAttributes(
		line int,
		position int,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructors
	Make() ValidatorLike
}

// Instances

/*
DefaultLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete default-like class.
*/
type DefaultLike interface {
}

/*
ExampleLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete example-like class.
*/
type ExampleLike interface {
	// Attributes
	GetDefault() DefaultLike
	GetPrimitive() PrimitiveLike
	GetLists() col.ListLike[ListLike]
}

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Methods
	FormatExample(example ExampleLike) string
}

/*
ListLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete list-like class.
*/
type ListLike interface {
	// Attributes
	GetExamples() col.ListLike[ExampleLike]
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Methods
	ParseSource(source string) ExampleLike
}

/*
PrimitiveLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete primitive-like class.
*/
type PrimitiveLike interface {
	// Attributes
	GetCharacter() string
	GetText() string
	GetInteger() string
	GetAnything() string
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Attributes
	GetLine() int
	GetPosition() int
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Methods
	ValidateExample(example ExampleLike)
}
