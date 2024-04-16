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

package cdsn

import ()

// CLASS ACCESS

// Reference

var atomClass = &atomClass_{
	// This class has no private constants to initialize.
}

// Function

func Atom() AtomClassLike {
	return atomClass
}

// CLASS METHODS

// Target

type atomClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *atomClass_) MakeWithGlyph(glyph GlyphLike) AtomLike {
	return &atom_{
		glyph_: glyph,
	}
}

func (c *atomClass_) MakeWithIntrinsic(intrinsic string) AtomLike {
	return &atom_{
		intrinsic_: intrinsic,
	}
}

// Functions

// INSTANCE METHODS

// Target

type atom_ struct {
	glyph_ GlyphLike
	intrinsic_ string
}

// Attributes

func (v *atom_) GetGlyph() GlyphLike {
	return v.glyph_
}

func (v *atom_) GetIntrinsic() string {
	return v.intrinsic_
}

// Public

// Private
