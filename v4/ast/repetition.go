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

var repetitionClass = &repetitionClass_{
	// Initialize class constants.
}

// Function

func Repetition() RepetitionClassLike {
	return repetitionClass
}

// CLASS METHODS

// Target

type repetitionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *repetitionClass_) Make(
	element ElementLike,
	optionalCardinality CardinalityLike,
) RepetitionLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(element):
		panic("The element attribute is required by this class.")
	default:
		return &repetition_{
			// Initialize instance attributes.
			class_:               c,
			element_:             element,
			optionalCardinality_: optionalCardinality,
		}
	}
}

// INSTANCE METHODS

// Target

type repetition_ struct {
	// Define instance attributes.
	class_               RepetitionClassLike
	element_             ElementLike
	optionalCardinality_ CardinalityLike
}

// Attributes

func (v *repetition_) GetClass() RepetitionClassLike {
	return v.class_
}

func (v *repetition_) GetElement() ElementLike {
	return v.element_
}

func (v *repetition_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// Private
