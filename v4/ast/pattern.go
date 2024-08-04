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

var patternClass = &patternClass_{
	// Initialize class constants.
}

// Function

func Pattern() PatternClassLike {
	return patternClass
}

// CLASS METHODS

// Target

type patternClass_ struct {
	// Define class constants.
}

// Constructors

func (c *patternClass_) Make(
	part PartLike,
	optionalSupplement SupplementLike,
) PatternLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(part):
		panic("The part attribute is required by this class.")
	default:
		return &pattern_{
			// Initialize instance attributes.
			class_:              c,
			part_:               part,
			optionalSupplement_: optionalSupplement,
		}
	}
}

// INSTANCE METHODS

// Target

type pattern_ struct {
	// Define instance attributes.
	class_              PatternClassLike
	part_               PartLike
	optionalSupplement_ SupplementLike
}

// Attributes

func (v *pattern_) GetClass() PatternClassLike {
	return v.class_
}

func (v *pattern_) GetPart() PartLike {
	return v.part_
}

func (v *pattern_) GetOptionalSupplement() SupplementLike {
	return v.optionalSupplement_
}

// Private
