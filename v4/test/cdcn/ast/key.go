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

var keyClass = &keyClass_{
	// Any private class constants should be initialized here.
}

// Function

func Key() KeyClassLike {
	return keyClass
}

// CLASS METHODS

// Target

type keyClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *keyClass_) MakeWithPrimitive(primitive PrimitiveLike) KeyLike {
	return &key_{
		primitive_: primitive,
	}
}

// Functions

// INSTANCE METHODS

// Target

type key_ struct {
	class_ KeyClassLike
	primitive_ PrimitiveLike
}

// Attributes

func (v *key_) GetClass() KeyClassLike {
	return v.class_
}

func (v *key_) GetPrimitive() PrimitiveLike {
	return v.primitive_
}

// Public

// Private
