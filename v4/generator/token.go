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
	col "github.com/craterdog/go-collection-framework/v4"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
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
		class_:    c,
		analyzer_: gra.Analyzer().Make(),
	}
}

// INSTANCE METHODS

// Target

type token_ struct {
	// Define the instance attributes.
	class_    TokenClassLike
	analyzer_ gra.AnalyzerLike
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
	v.analyzer_.AnalyzeSyntax(syntax)
	var notice = v.analyzer_.GetNotice()
	var template = v.getTemplate("tokenClass")
	implementation = replaceAll(template, "notice", notice)
	return implementation
}

// Private

func (v *token_) getTemplate(name string) string {
	var template = tokenTemplates_.GetValue(name)
	return template
}

// PRIVATE GLOBALS

// Constants

var tokenTemplates_ = col.Catalog[string, string](
	map[string]string{
		"tokenClass": `<Notice>

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
`,
	},
)
