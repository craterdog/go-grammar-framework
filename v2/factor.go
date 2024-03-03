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

var factorClass = &factorClass_{
	// This class does not initialize any constants.
}

// Function

func Factor() FactorClassLike {
	return factorClass
}

// CLASS METHODS

// Target

type factorClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *factorClass_) Make(
	predicate PredicateLike,
	cardinality CardinalityLike,
) FactorLike {
	var factor = &factor_{
		// This class does not initialize any attributes.
	}
	factor.SetPredicate(predicate)
	factor.SetCardinality(cardinality)
	return factor
}

// INSTANCE METHODS

// Target

type factor_ struct {
	cardinality CardinalityLike
	predicate   PredicateLike
}

// Public

func (v *factor_) GetCardinality() CardinalityLike {
	return v.cardinality
}

func (v *factor_) GetPredicate() PredicateLike {
	return v.predicate
}

func (v *factor_) SetCardinality(cardinality CardinalityLike) {
	v.cardinality = cardinality
}

func (v *factor_) SetPredicate(predicate PredicateLike) {
	if predicate == nil {
		panic("A predicate within a factor cannot be nil.")
	}
	v.predicate = predicate
}
