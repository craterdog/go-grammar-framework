/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammar

import (
	fmt "fmt"
)

// CLASS ACCESS

// Reference

var tokenClass = &tokenClass_{
	strings: map[TokenType]string{
		ErrorToken:     "Error",
		CharacterToken: "Character",
		CommentToken:   "Comment",
		DelimiterToken: "Delimiter",
		EOFToken:       "EOF",
		EOLToken:       "EOL",
		IntrinsicToken: "Intrinsic",
		LiteralToken:   "Literal",
		NameToken:      "Name",
		NoteToken:      "Note",
		NumberToken:    "Number",
		SpaceToken:     "Space",
		SymbolToken:    "Symbol",
	},
}

// Function

func Token() TokenClassLike {
	return tokenClass
}

// CLASS METHODS

// Target

type tokenClass_ struct {
	strings map[TokenType]string
}

// Constructors

func (c *tokenClass_) Make(
	line int,
	position int,
	tokenType TokenType,
	tokenValue string,
) TokenLike {
	var token = &token_{
		line:       line,
		position:   position,
		tokenType:  tokenType,
		tokenValue: tokenValue,
	}
	return token
}

// Functions

func (c *tokenClass_) AsString(tokenType TokenType) string {
	return c.strings[tokenType]
}

// INSTANCE METHODS

// Target

type token_ struct {
	line       int // The line number of the token in the source string.
	position   int // The position in the line of the first rune of the token.
	tokenType  TokenType
	tokenValue string
}

// Stringer

func (v *token_) String() string {
	var s = fmt.Sprintf("%q", v.tokenValue)
	if len(s) > 40 {
		s = fmt.Sprintf("%.40q...", v.tokenValue)
	}
	return fmt.Sprintf("Token [type: %s, line: %d, position: %d]: %s",
		Token().AsString(v.tokenType),
		v.line,
		v.position,
		s,
	)
}

// Public

func (v *token_) GetLine() int {
	return v.line
}

func (v *token_) GetPosition() int {
	return v.position
}

func (v *token_) GetType() TokenType {
	return v.tokenType
}

func (v *token_) GetValue() string {
	return v.tokenValue
}
