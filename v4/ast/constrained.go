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

import (
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var constrainedClass = &constrainedClass_{
	// Initialize class constants.
}

// Function

func Constrained() ConstrainedClassLike {
	return constrainedClass
}

// CLASS METHODS

// Target

type constrainedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constrainedClass_) Make(
	number string,
	optionalLimit LimitLike,
) ConstrainedLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(number):
		panic("The number attribute is required by this class.")
	default:
		return &constrained_{
			// Initialize instance attributes.
			class_:         c,
			number_:        number,
			optionalLimit_: optionalLimit,
		}
	}
}

// INSTANCE METHODS

// Target

type constrained_ struct {
	// Define instance attributes.
	class_         ConstrainedClassLike
	number_        string
	optionalLimit_ LimitLike
}

// Attributes

func (v *constrained_) GetClass() ConstrainedClassLike {
	return v.class_
}

func (v *constrained_) GetNumber() string {
	return v.number_
}

func (v *constrained_) GetOptionalLimit() LimitLike {
	return v.optionalLimit_
}

// Private
