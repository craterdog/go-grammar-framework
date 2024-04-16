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

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var constraintClass = &constraintClass_{
	// This class has no private constants to initialize.
}

// Function

func Constraint() ConstraintClassLike {
	return constraintClass
}

// CLASS METHODS

// Target

type constraintClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *constraintClass_) MakeWithNumbers(numbers col.ListLike[string]) ConstraintLike {
	return &constraint_{
		numbers_: numbers,
	}
}

// Functions

// INSTANCE METHODS

// Target

type constraint_ struct {
	numbers_ col.ListLike[string]
}

// Attributes

func (v *constraint_) GetNumbers() col.ListLike[string] {
	return v.numbers_
}

// Public

// Private
