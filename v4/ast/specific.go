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

var specificClass = &specificClass_{
	// Initialize class constants.
}

// Function

func Specific() SpecificClassLike {
	return specificClass
}

// CLASS METHODS

// Target

type specificClass_ struct {
	// Define class constants.
}

// Constructors

func (c *specificClass_) Make(runics abs.Sequential[string]) SpecificLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(runics):
		panic("The runics attribute is required by this class.")
	default:
		return &specific_{
			// Initialize instance attributes.
			class_:  c,
			runics_: runics,
		}
	}
}

// INSTANCE METHODS

// Target

type specific_ struct {
	// Define instance attributes.
	class_  SpecificClassLike
	runics_ abs.Sequential[string]
}

// Attributes

func (v *specific_) GetClass() SpecificClassLike {
	return v.class_
}

func (v *specific_) GetRunics() abs.Sequential[string] {
	return v.runics_
}

// Private
