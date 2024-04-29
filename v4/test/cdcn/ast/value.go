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

var valueClass = &valueClass_{
	// Any private class constants should be initialized here.
}

// Function

func Value() ValueClassLike {
	return valueClass
}

// CLASS METHODS

// Target

type valueClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *valueClass_) MakeWithPrimitive(primitive PrimitiveLike) ValueLike {
	return &value_{
		primitive_: primitive,
	}
}

func (c *valueClass_) MakeWithCollection(collection CollectionLike) ValueLike {
	return &value_{
		collection_: collection,
	}
}

// Functions

// INSTANCE METHODS

// Target

type value_ struct {
	class_ ValueClassLike
	primitive_ PrimitiveLike
	collection_ CollectionLike
}

// Attributes

func (v *value_) GetClass() ValueClassLike {
	return v.class_
}

func (v *value_) GetPrimitive() PrimitiveLike {
	return v.primitive_
}

func (v *value_) GetCollection() CollectionLike {
	return v.collection_
}

// Public

// Private
