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

var assertionClass = &assertionClass_{
	// TBA - Assign constant values.
}

// Function

func Assertion() AssertionClassLike {
	return assertionClass
}

// CLASS METHODS

// Target

type assertionClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *assertionClass_) MakeWithElement(element ElementLike) AssertionLike {
	return &assertion_{
		element_: element,
	}
}

func (c *assertionClass_) MakeWithGlyph(glyph GlyphLike) AssertionLike {
	return &assertion_{
		glyph_: glyph,
	}
}

func (c *assertionClass_) MakeWithPrecedence(precedence PrecedenceLike) AssertionLike {
	return &assertion_{
		precedence_: precedence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type assertion_ struct {
	element_    ElementLike
	glyph_      GlyphLike
	precedence_ PrecedenceLike
}

// Attributes

func (v *assertion_) GetElement() ElementLike {
	return v.element_
}

func (v *assertion_) GetGlyph() GlyphLike {
	return v.glyph_
}

func (v *assertion_) GetPrecedence() PrecedenceLike {
	return v.precedence_
}

// Public

// Private
