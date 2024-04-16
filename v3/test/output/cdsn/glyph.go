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

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var glyphClass = &glyphClass_{
	// This class has no private constants to initialize.
}

// Function

func Glyph() GlyphClassLike {
	return glyphClass
}

// CLASS METHODS

// Target

type glyphClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *glyphClass_) MakeWithCharacters(characters col.ListLike[string]) GlyphLike {
	return &glyph_{
		characters_: characters,
	}
}

// Functions

// INSTANCE METHODS

// Target

type glyph_ struct {
	characters_ col.ListLike[string]
}

// Attributes

func (v *glyph_) GetCharacters() col.ListLike[string] {
	return v.characters_
}

// Public

// Private
