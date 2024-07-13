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
	mod "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var minimumClass = &minimumClass_{
	// Initialize class constants.
}

// Function

func Minimum() MinimumClassLike {
	return minimumClass
}

// CLASS METHODS

// Target

type minimumClass_ struct {
	// Define class constants.
}

// Constructors

func (c *minimumClass_) Make(number string) MinimumLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(number):
		panic("The number attribute is required for each Minimum.")
	default:
		return &minimum_{
			// Initialize instance attributes.
			class_:  c,
			number_: number,
		}
	}
}

// INSTANCE METHODS

// Target

type minimum_ struct {
	// Define instance attributes.
	class_  MinimumClassLike
	number_ string
}

// Attributes

func (v *minimum_) GetClass() MinimumClassLike {
	return v.class_
}

func (v *minimum_) GetNumber() string {
	return v.number_
}

// Private
