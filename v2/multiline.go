/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package grammars

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// CLASS ACCESS

// Reference

var multilineClass = &multilineClass_{
	// TBA - Assign constant values.
}

// Function

func Multiline() MultilineClassLike {
	return multilineClass
}

// CLASS METHODS

// Target

type multilineClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *multilineClass_) MakeWithAttributes(lines col.ListLike[LineLike]) MultilineLike {
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
