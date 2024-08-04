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

var supplementClass = &supplementClass_{
	// Initialize class constants.
}

// Function

func Supplement() SupplementClassLike {
	return supplementClass
}

// CLASS METHODS

// Target

type supplementClass_ struct {
	// Define class constants.
}

// Constructors

func (c *supplementClass_) Make(any_ any) SupplementLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(any_):
		panic("The any attribute is required by this class.")
	default:
		return &supplement_{
			// Initialize instance attributes.
			class_: c,
			any_:   any_,
		}
	}
}

// INSTANCE METHODS

// Target

type supplement_ struct {
	// Define instance attributes.
	class_ SupplementClassLike
	any_   any
}

// Attributes

func (v *supplement_) GetClass() SupplementClassLike {
	return v.class_
}

func (v *supplement_) GetAny() any {
	return v.any_
}

// Private
