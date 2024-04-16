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

package cdsn

import ()

// CLASS ACCESS

// Reference

var factorClass = &factorClass_{
	// This class has no private constants to initialize.
}

// Function

func Factor() FactorClassLike {
	return factorClass
}

// CLASS METHODS

// Target

type factorClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *factorClass_) MakeWithAttributes(
	predicate PredicateLike,
	cardinality CardinalityLike,
) FactorLike {
	return &factor_{
		predicate_: predicate,
		cardinality_: cardinality,
	}
}

// Functions

// INSTANCE METHODS

// Target

type factor_ struct {
	predicate_ PredicateLike
	cardinality_ CardinalityLike
}

// Attributes

func (v *factor_) GetPredicate() PredicateLike {
	return v.predicate_
}

func (v *factor_) GetCardinality() CardinalityLike {
	return v.cardinality_
}

// Public

// Private
