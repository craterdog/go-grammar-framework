/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammars

import ()

// CLASS ACCESS

// Reference

var glyphClass = &glyphClass_{
	// TBA - Assign constant values.
}

// Function

func Glyph() GlyphClassLike {
	return glyphClass
}

// CLASS METHODS

// Target

type glyphClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *glyphClass_) MakeWithRange(first_ string, last_ string) GlyphLike {
	var result_ = &glyph_{
		first_: first_,
		last_:  last_,
	}
	return result_
}

// Functions

// INSTANCE METHODS

// Target

type glyph_ struct {
	first_ string
	last_  string
}

// Attributes

func (v *glyph_) GetFirst() string {
	return v.first_
}

func (v *glyph_) GetLast() string {
	return v.last_
}

// Public

// Private
