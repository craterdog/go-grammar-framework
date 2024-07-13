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
	mod "github.com/craterdog/go-collection-framework/v4"
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
	minimum MinimumLike,
	optionalMaximum MaximumLike,
) ConstrainedLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(minimum):
		panic("The minimum attribute is required for each Constrained.")
	default:
		return &constrained_{
			// Initialize instance attributes.
			class_:           c,
			minimum_:         minimum,
			optionalMaximum_: optionalMaximum,
		}
	}
}

// INSTANCE METHODS

// Target

type constrained_ struct {
	// Define instance attributes.
	class_           ConstrainedClassLike
	minimum_         MinimumLike
	optionalMaximum_ MaximumLike
}

// Attributes

func (v *constrained_) GetClass() ConstrainedClassLike {
	return v.class_
}

func (v *constrained_) GetMinimum() MinimumLike {
	return v.minimum_
}

func (v *constrained_) GetOptionalMaximum() MaximumLike {
	return v.optionalMaximum_
}

// Private
