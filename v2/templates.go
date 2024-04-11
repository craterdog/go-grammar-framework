/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package grammars

const packageTemplate_ = `
<Notice>
/*
Package "<packagename>" provides...

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
package <packagename>

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
)

// Classes

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
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructors
	Make() ParserLike
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
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Methods
	Format<Class>(<class> <Class>Like) string
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Methods
	ParseSource(source string) <Class>Like
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
	Validate<Class>(<class> <Class>Like)
}
`

const classCommentTemplate_ = `
/*
<ClassName>ClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete <class-name>-like class.
*/
`

const instanceCommentTemplate_ = `
/*
<ClassName>Like is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete <class-name>-like class.
*/
`

const grammarTemplate_ = `
!>
................................................................................
<Copyright>
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
<!

!>
<Notation>
This document contains a formal definition of...

A language grammar consists of a set of rule definitions and token definitions.

The following intrinsic character types are context specific:
 * ANY - Any language specific character.
 * LOWER - Any language specific lowercase character.
 * UPPER - Any language specific uppercase character.
 * DIGIT - Any language specific digit.
 * ESCAPE - Any environment specific escape sequence.
 * CONTROL - Any environment specific (non-printable) control character.
 * EOL - The environment specific end-of-line character.
 * EOF - The environment specific end-of-file marker (pseudo character).

A predicate may be constrained by any of the following cardinalities:
 * predicate{M} - Exactly M instances of the specified predicate.
 * predicate{M..} - M or more instances of the specified predicate.
 * predicate{M..N} - M to N instances of the specified predicate.
 * predicate? - Zero or one instances of the specified predicate.
 * predicate* - Zero or more instances of the specified predicate.
 * predicate+ - One or more instances of the specified predicate.

An inversion "~" within a definition may only be applied to an intrinsic
character type or a glyph range.
<!

!>
RULE DEFINITIONS
The following rules are used by the parser when parsing the stream of tokens
generated by the scanner.  Each rule name begins with an uppercase letter.  The
rule definitions may specify the names of tokens or other rules and are matched
by the parser in the order listed.  A rule definition may also be directly or
indirectly recursive.  The sequence of factors within in a rule definition may
be separated by spaces which are ignored by the parser.
<!
Source: Rule EOF  ! Terminated with an end-of-file marker.

Rule: token+

!>
TOKEN DEFINITIONS
The following token definitions are used by the scanner to generate the stream
of tokens that are processed by the parser.  Each token name begins with a
lowercase letter.  Unlike with rule definitions, a token definition cannot
specify the name of a rule within its definition but it can specify the name of
other tokens.  Token definitions cannot be recursive and the scanning of tokens
is NOT greedy.  Any spaces within a token definition are NOT ignored.
<!
token: UPPER (LOWER | UPPER)*

`
