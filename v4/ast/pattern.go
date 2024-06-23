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
	col "github.com/craterdog/go-collection-framework/v4/collection"
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

func (c *patternClass_) MakeWithAttributes(
	parts col.ListLike[PartLike],
	alternatives col.ListLike[AlternativeLike],
) PatternLike {
	return &pattern_{
		// Initialize instance attributes.
		class_:        c,
		parts_:        parts,
		alternatives_: alternatives,
	}
}

// INSTANCE METHODS

// Target

type pattern_ struct {
	// Define instance attributes.
	class_        PatternClassLike
	parts_        col.ListLike[PartLike]
	alternatives_ col.ListLike[AlternativeLike]
}

// Attributes

func (v *pattern_) GetClass() PatternClassLike {
	return v.class_
}

func (v *pattern_) GetParts() col.ListLike[PartLike] {
	return v.parts_
}

func (v *pattern_) GetAlternatives() col.ListLike[AlternativeLike] {
	return v.alternatives_
}

// Private
