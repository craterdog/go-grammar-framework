/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammar

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

func (c *alternativeClass_) MakeWithAttributes(factors_ col.Sequential[FactorLike], note_ string) AlternativeLike {
	var result_ = &alternative_{
		factors_: factors_,
		note_: note_,
	}
	return result_
}

// Functions

// INSTANCE METHODS

// Target

type alternative_ struct {
	factors_ col.Sequential[FactorLike]
	note_ string
}

// Attributes

func (v *alternative_) GetFactors() col.Sequential[FactorLike] {
	return v.factors_
}

func (v *alternative_) GetNote() string {
	return v.note_
}

// Public

// Private
