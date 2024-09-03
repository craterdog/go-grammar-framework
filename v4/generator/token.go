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

package generator

// CLASS ACCESS

import (
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
)

// Reference

var tokenClass = &tokenClass_{
	// Initialize the class constants.
}

// Function

func Token() TokenClassLike {
	return tokenClass
}

// CLASS METHODS

// Target

type tokenClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *tokenClass_) Make() TokenLike {
	return &token_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type token_ struct {
	// Define the instance attributes.
	class_ TokenClassLike
}

// Attributes

func (v *token_) GetClass() TokenClassLike {
	return v.class_
}

// Public

func (v *token_) GenerateTokenClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	var notice = v.generateNotice(syntax)
	implementation = replaceAll(tokenTemplate_, "notice", notice)
	return implementation
}

// Private

func (v *token_) generateNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

const tokenTemplate_ = `/*<Notice>*/

package grammar

// CLASS ACCESS

// Reference

var tokenClass = &tokenClass_{
	// Initialize the class constants.
}

// Function

func Token() TokenClassLike {
	return tokenClass
}

// CLASS METHODS

// Target

type tokenClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *tokenClass_) Make(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) TokenLike {
	return &token_{
		// Initialize the instance attributes.
		class_:    c,
		line_:     line,
		position_: position,
		type_:     type_,
		value_:    value,
	}
}

// INSTANCE METHODS

// Target

type token_ struct {
	// Define the instance attributes.
	class_    TokenClassLike
	line_     uint
	position_ uint
	type_     TokenType
	value_    string
}

// Attributes

func (v *token_) GetClass() TokenClassLike {
	return v.class_
}

func (v *token_) GetLine() uint {
	return v.line_
}

func (v *token_) GetPosition() uint {
	return v.position_
}

func (v *token_) GetType() TokenType {
	return v.type_
}

func (v *token_) GetValue() string {
	return v.value_
}
`
