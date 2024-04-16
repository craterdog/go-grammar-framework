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

package cdcn

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

func (c *primitiveClass_) MakeWithBoolean(boolean string) PrimitiveLike {
	return &primitive_{
		boolean_: boolean,
	}
}

func (c *primitiveClass_) MakeWithComplex(complex_ string) PrimitiveLike {
	return &primitive_{
		complex_: complex_,
	}
}

func (c *primitiveClass_) MakeWithFloat(float string) PrimitiveLike {
	return &primitive_{
		float_: float,
	}
}

func (c *primitiveClass_) MakeWithHexadecimal(hexadecimal string) PrimitiveLike {
	return &primitive_{
		hexadecimal_: hexadecimal,
	}
}

func (c *primitiveClass_) MakeWithInteger(integer string) PrimitiveLike {
	return &primitive_{
		integer_: integer,
	}
}

func (c *primitiveClass_) MakeWithNil(nil_ string) PrimitiveLike {
	return &primitive_{
		nil_: nil_,
	}
}

func (c *primitiveClass_) MakeWithRune(rune_ string) PrimitiveLike {
	return &primitive_{
		rune_: rune_,
	}
}

func (c *primitiveClass_) MakeWithString(string_ string) PrimitiveLike {
	return &primitive_{
		string_: string_,
	}
}

// Functions

// INSTANCE METHODS

// Target

type primitive_ struct {
	boolean_ string
	complex_ string
	float_ string
	hexadecimal_ string
	integer_ string
	nil_ string
	rune_ string
	string_ string
}

// Attributes

func (v *primitive_) GetBoolean() string {
	return v.boolean_
}

func (v *primitive_) GetComplex() string {
	return v.complex_
}

func (v *primitive_) GetFloat() string {
	return v.float_
}

func (v *primitive_) GetHexadecimal() string {
	return v.hexadecimal_
}

func (v *primitive_) GetInteger() string {
	return v.integer_
}

func (v *primitive_) GetNil() string {
	return v.nil_
}

func (v *primitive_) GetRune() string {
	return v.rune_
}

func (v *primitive_) GetString() string {
	return v.string_
}

// Public

// Private
