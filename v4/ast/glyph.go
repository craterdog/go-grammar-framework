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

import ()

// CLASS ACCESS

// Reference

var glyphClass = &glyphClass_{
	// Initialize class constants.
}

// Function

func Glyph() GlyphClassLike {
	return glyphClass
}

// CLASS METHODS

// Target

type glyphClass_ struct {
	// Define class constants.
	// This class has no private constants.

}

// Constructors

func (c *glyphClass_) MakeWithAttributes(
	first string,
	last string,
) GlyphLike {
	return &glyph_{
		// Initialize instance attributes.
		class_: c,
		first_: first,
		last_:  last,
	}
}

// INSTANCE METHODS

// Target

type glyph_ struct {
	// Define instance attributes.
	class_ GlyphClassLike
	first_ string
	last_  string
}

// Attributes

func (v *glyph_) GetClass() GlyphClassLike {
	return v.class_
}

func (v *glyph_) GetFirst() string {
	return v.first_
}

func (v *glyph_) GetLast() string {
	return v.last_
}

// Private
