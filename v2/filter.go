/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package grammars

import ()

// CLASS ACCESS

// Reference

var filterClass = &filterClass_{
	// TBA - Assign constant values.
}

// Function

func Filter() FilterClassLike {
	return filterClass
}

// CLASS METHODS

// Target

type filterClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *filterClass_) MakeWithGlyph(
	glyph GlyphLike,
	inverted bool,
) FilterLike {
	return &filter_{
		glyph_:    glyph,
		inverted_: inverted,
	}
}

func (c *filterClass_) MakeWithIntrinsic(
	intrinsic string,
	inverted bool,
) FilterLike {
	return &filter_{
		intrinsic_: intrinsic,
		inverted_:  inverted,
	}
}

// Functions

// INSTANCE METHODS

// Target

type filter_ struct {
	inverted_  bool
	intrinsic_ string
	glyph_     GlyphLike
}

// Attributes

func (v *filter_) IsInverted() bool {
	return v.inverted_
}

func (v *filter_) GetIntrinsic() string {
	return v.intrinsic_
}

func (v *filter_) GetGlyph() GlyphLike {
	return v.glyph_
}

// Public

// Private
