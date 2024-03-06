/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammars

import ()

// CLASS ACCESS

// Reference

var factorClass = &factorClass_{
	// TBA - Assign constant values.
}

// Function

func Factor() FactorClassLike {
	return factorClass
}

// CLASS METHODS

// Target

type factorClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *factorClass_) MakeWithAttributes(predicate_ PredicateLike, cardinality_ CardinalityLike) FactorLike {
	return &factor_{
		predicate_:   predicate_,
		cardinality_: cardinality_,
	}
}

// Functions

// INSTANCE METHODS

// Target

type factor_ struct {
	cardinality_ CardinalityLike
	predicate_   PredicateLike
}

// Attributes

func (v *factor_) GetCardinality() CardinalityLike {
	return v.cardinality_
}

func (v *factor_) GetPredicate() PredicateLike {
	return v.predicate_
}

// Public

// Private
