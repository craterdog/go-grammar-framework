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

var boundedClass = &boundedClass_{
	// Initialize class constants.
}

// Function

func Bounded() BoundedClassLike {
	return boundedClass
}

// CLASS METHODS

// Target

type boundedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *boundedClass_) Make(
	glyph string,
	optionalExtent ExtentLike,
) BoundedLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(glyph):
		panic("The glyph attribute is required by this class.")
	default:
		return &bounded_{
			// Initialize instance attributes.
			class_:          c,
			glyph_:          glyph,
			optionalExtent_: optionalExtent,
		}
	}
}

// INSTANCE METHODS

// Target

type bounded_ struct {
	// Define instance attributes.
	class_          BoundedClassLike
	glyph_          string
	optionalExtent_ ExtentLike
}

// Attributes

func (v *bounded_) GetClass() BoundedClassLike {
	return v.class_
}

func (v *bounded_) GetGlyph() string {
	return v.glyph_
}

func (v *bounded_) GetOptionalExtent() ExtentLike {
	return v.optionalExtent_
}

// Private
