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

package cdsn

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var multilineClass = &multilineClass_{
	// This class has no private constants to initialize.
}

// Function

func Multiline() MultilineClassLike {
	return multilineClass
}

// CLASS METHODS

// Target

type multilineClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *multilineClass_) MakeWithLines(lines col.ListLike[LineLike]) MultilineLike {
	return &multiline_{
		lines_: lines,
	}
}

// Functions

// INSTANCE METHODS

// Target

type multiline_ struct {
	lines_ col.ListLike[LineLike]
}

// Attributes

func (v *multiline_) GetLines() col.ListLike[LineLike] {
	return v.lines_
}

// Public

// Private
