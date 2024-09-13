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

var explicitClass = &explicitClass_{
	// Initialize class constants.
}

// Function

func Explicit() ExplicitClassLike {
	return explicitClass
}

// CLASS METHODS

// Target

type explicitClass_ struct {
	// Define class constants.
}

// Constructors

func (c *explicitClass_) Make(
	glyph string,
	optionalExtent ExtentLike,
) ExplicitLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(glyph):
		panic("The glyph attribute is required by this class.")
	default:
		return &explicit_{
			// Initialize instance attributes.
			class_:          c,
			glyph_:          glyph,
			optionalExtent_: optionalExtent,
		}
	}
}

// INSTANCE METHODS

// Target

type explicit_ struct {
	// Define instance attributes.
	class_          ExplicitClassLike
	glyph_          string
	optionalExtent_ ExtentLike
}

// Attributes

func (v *explicit_) GetClass() ExplicitClassLike {
	return v.class_
}

func (v *explicit_) GetGlyph() string {
	return v.glyph_
}

func (v *explicit_) GetOptionalExtent() ExtentLike {
	return v.optionalExtent_
}

// Private
