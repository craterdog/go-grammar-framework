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

import ()

// CLASS ACCESS

// Reference

var constraintClass = &constraintClass_{
	// TBA - Assign constant values.
}

// Function

func Constraint() ConstraintClassLike {
	return constraintClass
}

// CLASS METHODS

// Target

type constraintClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *constraintClass_) MakeWithAttributes(first string, last string) ConstraintLike {
	return &constraint_{
		first_: first,
		last_:  last,
	}
}

// Functions

// INSTANCE METHODS

// Target

type constraint_ struct {
	first_ string
	last_  string
}

// Attributes

func (v *constraint_) GetFirst() string {
	return v.first_
}

func (v *constraint_) GetLast() string {
	return v.last_
}

// Public

// Private
