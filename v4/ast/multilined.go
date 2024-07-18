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

var multilinedClass = &multilinedClass_{
	// Initialize class constants.
}

// Function

func Multilined() MultilinedClassLike {
	return multilinedClass
}

// CLASS METHODS

// Target

type multilinedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *multilinedClass_) Make(lines abs.Sequential[LineLike]) MultilinedLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(lines):
		panic("The lines attribute is required for each Multilined.")
	default:
		return &multilined_{
			// Initialize instance attributes.
			class_: c,
			lines_: lines,
		}
	}
}

// INSTANCE METHODS

// Target

type multilined_ struct {
	// Define instance attributes.
	class_ MultilinedClassLike
	lines_ abs.Sequential[LineLike]
}

// Attributes

func (v *multilined_) GetClass() MultilinedClassLike {
	return v.class_
}

func (v *multilined_) GetLines() abs.Sequential[LineLike] {
	return v.lines_
}

// Private
