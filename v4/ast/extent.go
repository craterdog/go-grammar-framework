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

func (c *extentClass_) Make(
	delimiter string,
	runic string,
) ExtentLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(delimiter):
		panic("The delimiter attribute is required by this class.")
	case col.IsUndefined(runic):
		panic("The runic attribute is required by this class.")
	default:
		return &extent_{
			// Initialize instance attributes.
			class_:     c,
			delimiter_: delimiter,
			runic_:     runic,
		}
	}
}

// INSTANCE METHODS

// Target

type extent_ struct {
	// Define instance attributes.
	class_     ExtentClassLike
	delimiter_ string
	runic_     string
}

// Attributes

func (v *extent_) GetClass() ExtentClassLike {
	return v.class_
}

func (v *extent_) GetDelimiter() string {
	return v.delimiter_
}

func (v *extent_) GetRunic() string {
	return v.runic_
}

// Private
