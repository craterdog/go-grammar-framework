/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package grammars

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// CLASS ACCESS

// Reference

var alternativeClass = &alternativeClass_{
	// TBA - Assign constant values.
}

// Function

func Alternative() AlternativeClassLike {
	return alternativeClass
}

// CLASS METHODS

// Target

type alternativeClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *alternativeClass_) MakeWithFactors(factors col.ListLike[FactorLike]) AlternativeLike {
	return &alternative_{
		factors_: factors,
	}
}

// Functions

// INSTANCE METHODS

// Target

type alternative_ struct {
	factors_ col.ListLike[FactorLike]
}

// Attributes

func (v *alternative_) GetFactors() col.ListLike[FactorLike] {
	return v.factors_
}

// Public

// Private
