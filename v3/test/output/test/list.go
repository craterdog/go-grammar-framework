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

package test

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var listClass = &listClass_{
	// This class has no private constants to initialize.
}

// Function

func List() ListClassLike {
	return listClass
}

// CLASS METHODS

// Target

type listClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *listClass_) MakeWithExamples(examples col.ListLike[ExampleLike]) ListLike {
	return &list_{
		examples_: examples,
	}
}

// Functions

// INSTANCE METHODS

// Target

type list_ struct {
	examples_ col.ListLike[ExampleLike]
}

// Attributes

func (v *list_) GetExamples() col.ListLike[ExampleLike] {
	return v.examples_
}

// Public

// Private
