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
	mod "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var elementClass = &elementClass_{
	// Initialize class constants.
}

// Function

func Element() ElementClassLike {
	return elementClass
}

// CLASS METHODS

// Target

type elementClass_ struct {
	// Define class constants.
}

// Constructors

func (c *elementClass_) Make(any_ any) ElementLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(any_):
		panic("The any_ attribute is required for each Element.")
	default:
		return &element_{
			// Initialize instance attributes.
			class_: c,
			any_:   any_,
		}
	}
}

// INSTANCE METHODS

// Target

type element_ struct {
	// Define instance attributes.
	class_ ElementClassLike
	any_   any
}

// Attributes

func (v *element_) GetClass() ElementClassLike {
	return v.class_
}

func (v *element_) GetAny() any {
	return v.any_
}

// Private
