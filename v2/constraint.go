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

var constraintClass = &constraintClass_{
	// This class does not initialize any constants.
}

// Function

func Constraint() ConstraintClassLike {
	return constraintClass
}

// CLASS METHODS

// Target

type constraintClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *constraintClass_) Make(first, last string) ConstraintLike {
	var constraint = &constraint_{
		// This class does not initialize any attributes.
	}
	constraint.SetFirst(first)
	constraint.SetLast(last)
	return constraint
}

// INSTANCE METHODS

// Target

type constraint_ struct {
	first string
	last  string
}

// Public

func (v *constraint_) GetFirst() string {
	return v.first
}

func (v *constraint_) GetLast() string {
	return v.last
}

func (v *constraint_) SetFirst(first string) {
	if len(first) < 1 {
		panic("A constraint requires a first number.")
	}
	v.first = first
}

func (v *constraint_) SetLast(last string) {
	v.last = last
}
