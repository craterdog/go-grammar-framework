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

var multilineClass = &multilineClass_{
	// Initialize class constants.
}

// Function

func Multiline() MultilineClassLike {
	return multilineClass
}

// CLASS METHODS

// Target

type multilineClass_ struct {
	// Define class constants.
}

// Constructors

func (c *multilineClass_) MakeWithLines(lines col.ListLike[LineLike]) MultilineLike {
	return &multiline_{
		// Initialize instance attributes.
		class_: c,
		lines_: lines,
	}
}

// INSTANCE METHODS

// Target

type multiline_ struct {
	// Define instance attributes.
	class_ MultilineClassLike
	lines_ col.ListLike[LineLike]
}

// Attributes

func (v *multiline_) GetClass() MultilineClassLike {
	return v.class_
}

func (v *multiline_) GetLines() col.ListLike[LineLike] {
	return v.lines_
}

// Private
