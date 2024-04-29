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

var valuesClass = &valuesClass_{
	// Any private class constants should be initialized here.
}

// Function

func Values() ValuesClassLike {
	return valuesClass
}

// CLASS METHODS

// Target

type valuesClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *valuesClass_) MakeWithValues(values col.ListLike[ValueLike]) ValuesLike {
	return &values_{
		values_: values,
	}
}

// Functions

// INSTANCE METHODS

// Target

type values_ struct {
	class_ ValuesClassLike
	values_ col.ListLike[ValueLike]
}

// Attributes

func (v *values_) GetClass() ValuesClassLike {
	return v.class_
}

func (v *values_) GetValues() col.ListLike[ValueLike] {
	return v.values_
}

// Public

// Private
