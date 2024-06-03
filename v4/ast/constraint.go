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

var constraintClass = &constraintClass_{
	// Initialize class constants.
}

// Function

func Constraint() ConstraintClassLike {
	return constraintClass
}

// CLASS METHODS

// Target

type constraintClass_ struct {
	// Define class constants.
	// This class has no private constants.

}

// Constructors

func (c *constraintClass_) MakeWithAttributes(
	first string,
	last string,
) ConstraintLike {
	return &constraint_{
		// Initialize instance attributes.
		class_: c,
		first_: first,
		last_:  last,
	}
}

// INSTANCE METHODS

// Target

type constraint_ struct {
	// Define instance attributes.
	class_ ConstraintClassLike
	first_ string
	last_  string
}

// Attributes

func (v *constraint_) GetClass() ConstraintClassLike {
	return v.class_
}

func (v *constraint_) GetFirst() string {
	return v.first_
}

func (v *constraint_) GetLast() string {
	return v.last_
}

// Private
