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

import ()

// CLASS ACCESS

// Reference

var extentClass = &extentClass_{
	// Initialize class constants.
}

// Function

func Extent() ExtentClassLike {
	return extentClass
}

// CLASS METHODS

// Target

type extentClass_ struct {
	// Define class constants.
}

// Constructors

func (c *extentClass_) Make(rune_ string) ExtentLike {
	return &extent_{
		// Initialize instance attributes.
		class_: c,
		rune_:  rune_,
	}
}

// INSTANCE METHODS

// Target

type extent_ struct {
	// Define instance attributes.
	class_ ExtentClassLike
	rune_  string
}

// Attributes

func (v *extent_) GetClass() ExtentClassLike {
	return v.class_
}

func (v *extent_) GetRune() string {
	return v.rune_
}

// Private
