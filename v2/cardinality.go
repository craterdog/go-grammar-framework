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

// CLASS ACCESS

// Reference

var cardinalityClass = &cardinalityClass_{
	// This class does not initialize any constants.
}

// Function

func Cardinality() CardinalityClassLike {
	return cardinalityClass
}

// CLASS METHODS

// Target

type cardinalityClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *cardinalityClass_) Make(constraint ConstraintLike) CardinalityLike {
	var cardinality = &cardinality_{
		// This class does not initialize any attributes.
	}
	cardinality.SetConstraint(constraint)
	return cardinality
}

// INSTANCE METHODS

// Target

type cardinality_ struct {
	constraint ConstraintLike
}

// Public

func (v *cardinality_) GetConstraint() ConstraintLike {
	return v.constraint
}

func (v *cardinality_) SetConstraint(constraint ConstraintLike) {
	if constraint == nil {
		panic("A constraint must not be nil.")
	}
	v.constraint = constraint
}
