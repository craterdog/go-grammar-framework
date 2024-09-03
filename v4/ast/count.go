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

var countClass = &countClass_{
	// Initialize class constants.
}

// Function

func Count() CountClassLike {
	return countClass
}

// CLASS METHODS

// Target

type countClass_ struct {
	// Define class constants.
}

// Constructors

func (c *countClass_) Make(numbers abs.Sequential[string]) CountLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(numbers):
		panic("The numbers attribute is required by this class.")
	default:
		return &count_{
			// Initialize instance attributes.
			class_:   c,
			numbers_: numbers,
		}
	}
}

// INSTANCE METHODS

// Target

type count_ struct {
	// Define instance attributes.
	class_   CountClassLike
	numbers_ abs.Sequential[string]
}

// Attributes

func (v *count_) GetClass() CountClassLike {
	return v.class_
}

func (v *count_) GetNumbers() abs.Sequential[string] {
	return v.numbers_
}

// Private
