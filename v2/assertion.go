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

var assertionClass = &assertionClass_{
	// This class does not initialize any constants.
}

// Function

func Assertion() AssertionClassLike {
	return assertionClass
}

// CLASS METHODS

// Target

type assertionClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *assertionClass_) MakeFromElement(element ElementLike) AssertionLike {
	var assertion = &assertion_{
		// This class does not initialize any attributes.
	}
	assertion.SetElement(element)
	return assertion
}

func (c *assertionClass_) MakeFromGlyph(glyph GlyphLike) AssertionLike {
	var assertion = &assertion_{
		// This class does not initialize any attributes.
	}
	assertion.SetGlyph(glyph)
	return assertion
}

func (c *assertionClass_) MakeFromPrecedence(precedence PrecedenceLike) AssertionLike {
	var assertion = &assertion_{
		// This class does not initialize any attributes.
	}
	assertion.SetPrecedence(precedence)
	return assertion
}

// INSTANCE METHODS

// Target

type assertion_ struct {
	element    ElementLike
	glyph      GlyphLike
	precedence PrecedenceLike
}

// Public

func (v *assertion_) GetElement() ElementLike {
	return v.element
}

func (v *assertion_) GetGlyph() GlyphLike {
	return v.glyph
}

func (v *assertion_) GetPrecedence() PrecedenceLike {
	return v.precedence
}

func (v *assertion_) SetElement(element ElementLike) {
	if element == nil {
		panic("An element must not be nil.")
	}
	v.element = element
	v.glyph = nil
	v.precedence = nil
}

func (v *assertion_) SetGlyph(glyph GlyphLike) {
	if glyph == nil {
		panic("A glyph must not be nil.")
	}
	v.element = nil
	v.glyph = glyph
	v.precedence = nil
}

func (v *assertion_) SetPrecedence(precedence PrecedenceLike) {
	if precedence == nil {
		panic("A precedence must not be nil.")
	}
	v.element = nil
	v.glyph = nil
	v.precedence = precedence
}
