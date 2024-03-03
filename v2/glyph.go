/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammar

// CLASS ACCESS

// Reference

var glyphClass = &glyphClass_{
	// This class does not initialize any constants.
}

// Function

func Glyph() GlyphClassLike {
	return glyphClass
}

// CLASS METHODS

// Target

type glyphClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *glyphClass_) Make(first, last string) GlyphLike {
	var glyph = &glyph_{
		// This class does not initialize any attributes.
	}
	glyph.SetFirst(first)
	glyph.SetLast(last)
	return glyph
}

// INSTANCE METHODS

// Target

type glyph_ struct {
	first string
	last  string
}

// Public

func (v *glyph_) GetFirst() string {
	return v.first
}

func (v *glyph_) GetLast() string {
	return v.last
}

func (v *glyph_) SetFirst(first string) {
	if len(first) < 1 {
		panic("A glyph requires a first character.")
	}
	v.first = first
}

func (v *glyph_) SetLast(last string) {
	v.last = last
}
