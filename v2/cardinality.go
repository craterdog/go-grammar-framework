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

var cardinalityClass = &cardinalityClass_{
	// TBA - Assign constant values.
}

// Function

func Cardinality() CardinalityClassLike {
	return cardinalityClass
}

// CLASS METHODS

// Target

type cardinalityClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *cardinalityClass_) MakeWithAttributes(constraint_ ConstraintLike) CardinalityLike {
	return &cardinality_{
		constraint_: constraint_,
	}
}

// Functions

// INSTANCE METHODS

// Target

type cardinality_ struct {
	constraint_ ConstraintLike
}

// Attributes

func (v *cardinality_) GetConstraint() ConstraintLike {
	return v.constraint_
}

// Public

// Private
