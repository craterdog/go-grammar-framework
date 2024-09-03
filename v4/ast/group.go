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

var groupClass = &groupClass_{
	// Initialize class constants.
}

// Function

func Group() GroupClassLike {
	return groupClass
}

// CLASS METHODS

// Target

type groupClass_ struct {
	// Define class constants.
}

// Constructors

func (c *groupClass_) Make(pattern PatternLike) GroupLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(pattern):
		panic("The pattern attribute is required by this class.")
	default:
		return &group_{
			// Initialize instance attributes.
			class_:   c,
			pattern_: pattern,
		}
	}
}

// INSTANCE METHODS

// Target

type group_ struct {
	// Define instance attributes.
	class_   GroupClassLike
	pattern_ PatternLike
}

// Attributes

func (v *group_) GetClass() GroupClassLike {
	return v.class_
}

func (v *group_) GetPattern() PatternLike {
	return v.pattern_
}

// Private
