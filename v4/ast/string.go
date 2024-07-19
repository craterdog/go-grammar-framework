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

var stringClass = &stringClass_{
	// Initialize class constants.
}

// Function

func String() StringClassLike {
	return stringClass
}

// CLASS METHODS

// Target

type stringClass_ struct {
	// Define class constants.
}

// Constructors

func (c *stringClass_) Make(any_ any) StringLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(any_):
		panic("The any attribute is required by this class.")
	default:
		return &string_{
			// Initialize instance attributes.
			class_: c,
			any_:   any_,
		}
	}
}

// INSTANCE METHODS

// Target

type string_ struct {
	// Define instance attributes.
	class_ StringClassLike
	any_   any
}

// Attributes

func (v *string_) GetClass() StringClassLike {
	return v.class_
}

func (v *string_) GetAny() any {
	return v.any_
}

// Private
