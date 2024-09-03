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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var inlineClass = &inlineClass_{
	// Initialize class constants.
}

// Function

func Inline() InlineClassLike {
	return inlineClass
}

// CLASS METHODS

// Target

type inlineClass_ struct {
	// Define class constants.
}

// Constructors

func (c *inlineClass_) Make(
	terms abs.Sequential[TermLike],
	optionalNote string,
) InlineLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(terms):
		panic("The terms attribute is required by this class.")
	default:
		return &inline_{
			// Initialize instance attributes.
			class_:        c,
			terms_:        terms,
			optionalNote_: optionalNote,
		}
	}
}

// INSTANCE METHODS

// Target

type inline_ struct {
	// Define instance attributes.
	class_        InlineClassLike
	terms_        abs.Sequential[TermLike]
	optionalNote_ string
}

// Attributes

func (v *inline_) GetClass() InlineClassLike {
	return v.class_
}

func (v *inline_) GetTerms() abs.Sequential[TermLike] {
	return v.terms_
}

func (v *inline_) GetOptionalNote() string {
	return v.optionalNote_
}

// Private
