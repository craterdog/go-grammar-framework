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

var groupedClass = &groupedClass_{
	// Initialize class constants.
}

// Function

func Grouped() GroupedClassLike {
	return groupedClass
}

// CLASS METHODS

// Target

type groupedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *groupedClass_) Make(
	reserved string,
	pattern PatternLike,
	reserved2 string,
) GroupedLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(reserved):
		panic("The reserved attribute is required by this class.")
	case col.IsUndefined(pattern):
		panic("The pattern attribute is required by this class.")
	case col.IsUndefined(reserved2):
		panic("The reserved2 attribute is required by this class.")
	default:
		return &grouped_{
			// Initialize instance attributes.
			class_:     c,
			reserved_:  reserved,
			pattern_:   pattern,
			reserved2_: reserved2,
		}
	}
}

// INSTANCE METHODS

// Target

type grouped_ struct {
	// Define instance attributes.
	class_     GroupedClassLike
	reserved_  string
	pattern_   PatternLike
	reserved2_ string
}

// Attributes

func (v *grouped_) GetClass() GroupedClassLike {
	return v.class_
}

func (v *grouped_) GetReserved() string {
	return v.reserved_
}

func (v *grouped_) GetPattern() PatternLike {
	return v.pattern_
}

func (v *grouped_) GetReserved2() string {
	return v.reserved2_
}

// Private
