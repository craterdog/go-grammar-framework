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
Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

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
package grammar

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
)

// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	CommentToken
	DelimiterToken
	ExcludedToken
	GlyphToken
	IntrinsicToken
	LiteralToken
	LowercaseToken
	NewlineToken
	NoteToken
	NumberToken
	OptionalToken
	RepeatedToken
	SpaceToken
	UppercaseToken
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
ProcessorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructors
	Make() ProcessorLike
}

/*
ScannerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructors
	Make(
		source string,
		tokens abs.QueueLike[TokenLike],
	) ScannerLike

	// Functions
	FormatToken(token TokenLike) string
	FormatType(tokenType TokenType) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructors
	Make(
		line uint,
		position uint,
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

/*
VisitorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructors
	Make(processor Methodical) VisitorLike
}

// Instances

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Attributes
	GetClass() FormatterClassLike

	// Abstractions
	Methodical

	// Methods
	FormatSyntax(syntax ast.SyntaxLike) string
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Attributes
	GetClass() ParserClassLike

	// Methods
	ParseSource(source string) ast.SyntaxLike
}

/*
ProcessorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Attributes
	GetClass() ProcessorClassLike

	// Abstractions
	Methodical
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Attributes
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Attributes
	GetClass() TokenClassLike
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Attributes
	GetClass() ValidatorClassLike

	// Abstractions
	Methodical

	// Methods
	ValidateSyntax(syntax ast.SyntaxLike)
}

/*
VisitorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Attributes
	GetClass() VisitorClassLike

	// Methods
	VisitSyntax(syntax ast.SyntaxLike)
}

// Aspects

/*
Methodical defines the set of method signatures that must be supported
by all methodical processors.
*/
type Methodical interface {
	ProcessComment(comment string)
	ProcessExcluded(excluded string)
	ProcessGlyph(glyph string)
	ProcessIntrinsic(intrinsic string)
	ProcessLiteral(literal string)
	ProcessLowercase(lowercase string)
	ProcessNewline(
		newline string,
		index uint,
		size uint,
	)
	ProcessNote(note string)
	ProcessNumber(number string)
	ProcessOptional(optional string)
	ProcessRepeated(repeated string)
	ProcessUppercase(uppercase string)
	PreprocessAlternative(
		alternative ast.AlternativeLike,
		index uint,
		size uint,
	)
	PostprocessAlternative(
		alternative ast.AlternativeLike,
		index uint,
		size uint,
	)
	PreprocessCardinality(cardinality ast.CardinalityLike)
	PostprocessCardinality(cardinality ast.CardinalityLike)
	PreprocessCharacter(
		character ast.CharacterLike,
		index uint,
		size uint,
	)
	PostprocessCharacter(
		character ast.CharacterLike,
		index uint,
		size uint,
	)
	PreprocessConstrained(constrained ast.ConstrainedLike)
	PostprocessConstrained(constrained ast.ConstrainedLike)
	PreprocessDefinition(definition ast.DefinitionLike)
	PostprocessDefinition(definition ast.DefinitionLike)
	PreprocessElement(element ast.ElementLike)
	PostprocessElement(element ast.ElementLike)
	PreprocessExplicit(explicit ast.ExplicitLike)
	PostprocessExplicit(explicit ast.ExplicitLike)
	PreprocessExpression(
		expression ast.ExpressionLike,
		index uint,
		size uint,
	)
	PostprocessExpression(
		expression ast.ExpressionLike,
		index uint,
		size uint,
	)
	PreprocessExtent(extent ast.ExtentLike)
	PostprocessExtent(extent ast.ExtentLike)
	PreprocessFilter(filter ast.FilterLike)
	PostprocessFilter(filter ast.FilterLike)
	PreprocessGroup(group ast.GroupLike)
	PostprocessGroup(group ast.GroupLike)
	PreprocessIdentifier(identifier ast.IdentifierLike)
	PostprocessIdentifier(identifier ast.IdentifierLike)
	PreprocessInline(inline ast.InlineLike)
	PostprocessInline(inline ast.InlineLike)
	PreprocessLimit(limit ast.LimitLike)
	PostprocessLimit(limit ast.LimitLike)
	PreprocessLine(
		line ast.LineLike,
		index uint,
		size uint,
	)
	PostprocessLine(
		line ast.LineLike,
		index uint,
		size uint,
	)
	PreprocessMultiline(multiline ast.MultilineLike)
	PostprocessMultiline(multiline ast.MultilineLike)
	PreprocessNotice(notice ast.NoticeLike)
	PostprocessNotice(notice ast.NoticeLike)
	PreprocessOption(option ast.OptionLike)
	PostprocessOption(option ast.OptionLike)
	PreprocessPattern(pattern ast.PatternLike)
	PostprocessPattern(pattern ast.PatternLike)
	PreprocessQuantified(quantified ast.QuantifiedLike)
	PostprocessQuantified(quantified ast.QuantifiedLike)
	PreprocessReference(reference ast.ReferenceLike)
	PostprocessReference(reference ast.ReferenceLike)
	PreprocessRepetition(
		repetition ast.RepetitionLike,
		index uint,
		size uint,
	)
	PostprocessRepetition(
		repetition ast.RepetitionLike,
		index uint,
		size uint,
	)
	PreprocessRule(
		rule ast.RuleLike,
		index uint,
		size uint,
	)
	PostprocessRule(
		rule ast.RuleLike,
		index uint,
		size uint,
	)
	PreprocessSyntax(syntax ast.SyntaxLike)
	PostprocessSyntax(syntax ast.SyntaxLike)
	PreprocessTerm(
		term ast.TermLike,
		index uint,
		size uint,
	)
	PostprocessTerm(
		term ast.TermLike,
		index uint,
		size uint,
	)
	PreprocessText(text ast.TextLike)
	PostprocessText(text ast.TextLike)
}
