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
	separator string,
	pattern PatternLike,
	separator2 string,
) GroupedLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(separator):
		panic("The separator attribute is required by this class.")
	case col.IsUndefined(pattern):
		panic("The pattern attribute is required by this class.")
	case col.IsUndefined(separator2):
		panic("The separator2 attribute is required by this class.")
	default:
		return &grouped_{
			// Initialize instance attributes.
			class_:      c,
			separator_:  separator,
			pattern_:    pattern,
			separator2_: separator2,
		}
	}
}

// INSTANCE METHODS

// Target

type grouped_ struct {
	// Define instance attributes.
	class_      GroupedClassLike
	separator_  string
	pattern_    PatternLike
	separator2_ string
}

// Attributes

func (v *grouped_) GetClass() GroupedClassLike {
	return v.class_
}

func (v *grouped_) GetSeparator() string {
	return v.separator_
}

func (v *grouped_) GetPattern() PatternLike {
	return v.pattern_
}

func (v *grouped_) GetSeparator2() string {
	return v.separator2_
}

// Private
