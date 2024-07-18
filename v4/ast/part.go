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

var partClass = &partClass_{
	// Initialize class constants.
}

// Function

func Part() PartClassLike {
	return partClass
}

// CLASS METHODS

// Target

type partClass_ struct {
	// Define class constants.
}

// Constructors

func (c *partClass_) Make(
	element ElementLike,
	optionalCardinality CardinalityLike,
) PartLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(element):
		panic("The element attribute is required for each Part.")
	default:
		return &part_{
			// Initialize instance attributes.
			class_:               c,
			element_:             element,
			optionalCardinality_: optionalCardinality,
		}
	}
}

// INSTANCE METHODS

// Target

type part_ struct {
	// Define instance attributes.
	class_               PartClassLike
	element_             ElementLike
	optionalCardinality_ CardinalityLike
}

// Attributes

func (v *part_) GetClass() PartClassLike {
	return v.class_
}

func (v *part_) GetElement() ElementLike {
	return v.element_
}

func (v *part_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// Private
