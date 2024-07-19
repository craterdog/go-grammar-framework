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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
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
	parts abs.Sequential[PartLike],
	alternatives abs.Sequential[AlternativeLike],
) PatternLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(parts):
		panic("The parts attribute is required by this class.")
	case col.IsUndefined(alternatives):
		panic("The alternatives attribute is required by this class.")
	default:
		return &pattern_{
			// Initialize instance attributes.
			class_:        c,
			parts_:        parts,
			alternatives_: alternatives,
		}
	}
}

// INSTANCE METHODS

// Target

type pattern_ struct {
	// Define instance attributes.
	class_        PatternClassLike
	parts_        abs.Sequential[PartLike]
	alternatives_ abs.Sequential[AlternativeLike]
}

// Attributes

func (v *pattern_) GetClass() PatternClassLike {
	return v.class_
}

func (v *pattern_) GetParts() abs.Sequential[PartLike] {
	return v.parts_
}

func (v *pattern_) GetAlternatives() abs.Sequential[AlternativeLike] {
	return v.alternatives_
}

// Private
