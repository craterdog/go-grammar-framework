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
	// This class does not initialize any constants.
}

// Function

func Alternative() AlternativeClassLike {
	return alternativeClass
}

// CLASS METHODS

// Target

type alternativeClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *alternativeClass_) Make(
	factors col.Sequential[FactorLike],
	note string,
) AlternativeLike {
	var alternative = &alternative_{
		// This class does not initialize any attributes.
	}
	alternative.SetFactors(factors)
	alternative.SetNote(note)
	return alternative
}

// INSTANCE METHODS

// Target

type alternative_ struct {
	factors col.Sequential[FactorLike]
	note    string
}

// Public

func (v *alternative_) GetFactors() col.Sequential[FactorLike] {
	return v.factors
}

func (v *alternative_) GetNote() string {
	return v.note
}

func (v *alternative_) SetFactors(factors col.Sequential[FactorLike]) {
	if factors == nil || factors.IsEmpty() {
		panic("An alternative must have at least one factor.")
	}
	v.factors = factors
}

func (v *alternative_) SetNote(note string) {
	v.note = note
}
