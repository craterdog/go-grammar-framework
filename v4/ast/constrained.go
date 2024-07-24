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
	separator string,
	number string,
	optionalLimit LimitLike,
	separator2 string,
) ConstrainedLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(separator):
		panic("The separator attribute is required by this class.")
	case col.IsUndefined(number):
		panic("The number attribute is required by this class.")
	case col.IsUndefined(separator2):
		panic("The separator2 attribute is required by this class.")
	default:
		return &constrained_{
			// Initialize instance attributes.
			class_:         c,
			separator_:     separator,
			number_:        number,
			optionalLimit_: optionalLimit,
			separator2_:    separator2,
		}
	}
}

// INSTANCE METHODS

// Target

type constrained_ struct {
	// Define instance attributes.
	class_         ConstrainedClassLike
	separator_     string
	number_        string
	optionalLimit_ LimitLike
	separator2_    string
}

// Attributes

func (v *constrained_) GetClass() ConstrainedClassLike {
	return v.class_
}

func (v *constrained_) GetSeparator() string {
	return v.separator_
}

func (v *constrained_) GetNumber() string {
	return v.number_
}

func (v *constrained_) GetOptionalLimit() LimitLike {
	return v.optionalLimit_
}

func (v *constrained_) GetSeparator2() string {
	return v.separator2_
}

// Private
