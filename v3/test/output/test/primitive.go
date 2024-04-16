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

package test

import ()

// CLASS ACCESS

// Reference

var primitiveClass = &primitiveClass_{
	// This class has no private constants to initialize.
}

// Function

func Primitive() PrimitiveClassLike {
	return primitiveClass
}

// CLASS METHODS

// Target

type primitiveClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *primitiveClass_) MakeWithCharacter(character string) PrimitiveLike {
	return &primitive_{
		character_: character,
	}
}

func (c *primitiveClass_) MakeWithText(text string) PrimitiveLike {
	return &primitive_{
		text_: text,
	}
}

func (c *primitiveClass_) MakeWithInteger(integer string) PrimitiveLike {
	return &primitive_{
		integer_: integer,
	}
}

func (c *primitiveClass_) MakeWithAnything(anything string) PrimitiveLike {
	return &primitive_{
		anything_: anything,
	}
}

// Functions

// INSTANCE METHODS

// Target

type primitive_ struct {
	character_ string
	text_ string
	integer_ string
	anything_ string
}

// Attributes

func (v *primitive_) GetCharacter() string {
	return v.character_
}

func (v *primitive_) GetText() string {
	return v.text_
}

func (v *primitive_) GetInteger() string {
	return v.integer_
}

func (v *primitive_) GetAnything() string {
	return v.anything_
}

// Public

// Private
