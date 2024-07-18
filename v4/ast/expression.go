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

package ast

import (
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var expressionClass = &expressionClass_{
	// Initialize class constants.
}

// Function

func Expression() ExpressionClassLike {
	return expressionClass
}

// CLASS METHODS

// Target

type expressionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *expressionClass_) Make(
	optionalComment string,
	lowercase string,
	pattern PatternLike,
	optionalNote string,
) ExpressionLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(lowercase):
		panic("The lowercase attribute is required for each Expression.")
	case col.IsUndefined(pattern):
		panic("The pattern attribute is required for each Expression.")
	default:
		return &expression_{
			// Initialize instance attributes.
			class_:           c,
			optionalComment_: optionalComment,
			lowercase_:       lowercase,
			pattern_:         pattern,
			optionalNote_:    optionalNote,
		}
	}
}

// INSTANCE METHODS

// Target

type expression_ struct {
	// Define instance attributes.
	class_           ExpressionClassLike
	optionalComment_ string
	lowercase_       string
	pattern_         PatternLike
	optionalNote_    string
}

// Attributes

func (v *expression_) GetClass() ExpressionClassLike {
	return v.class_
}

func (v *expression_) GetOptionalComment() string {
	return v.optionalComment_
}

func (v *expression_) GetLowercase() string {
	return v.lowercase_
}

func (v *expression_) GetPattern() PatternLike {
	return v.pattern_
}

func (v *expression_) GetOptionalNote() string {
	return v.optionalNote_
}

// Private
