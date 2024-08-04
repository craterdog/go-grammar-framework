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

var textualClass = &textualClass_{
	// Initialize class constants.
}

// Function

func Textual() TextualClassLike {
	return textualClass
}

// CLASS METHODS

// Target

type textualClass_ struct {
	// Define class constants.
}

// Constructors

func (c *textualClass_) Make(any_ any) TextualLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(any_):
		panic("The any attribute is required by this class.")
	default:
		return &textual_{
			// Initialize instance attributes.
			class_: c,
			any_:   any_,
		}
	}
}

// INSTANCE METHODS

// Target

type textual_ struct {
	// Define instance attributes.
	class_ TextualClassLike
	any_   any
}

// Attributes

func (v *textual_) GetClass() TextualClassLike {
	return v.class_
}

func (v *textual_) GetAny() any {
	return v.any_
}

// Private
