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

import ()

// CLASS ACCESS

// Reference

var factorClass = &factorClass_{
	// Initialize class constants.
}

// Function

func Factor() FactorClassLike {
	return factorClass
}

// CLASS METHODS

// Target

type factorClass_ struct {
	// Define class constants.
	// This class has no private constants.

}

// Constructors

func (c *factorClass_) MakeWithAttributes(
	predicate PredicateLike,
	cardinality CardinalityLike,
) FactorLike {
	return &factor_{
		// Initialize instance attributes.
		class_:       c,
		predicate_:   predicate,
		cardinality_: cardinality,
	}
}

// INSTANCE METHODS

// Target

type factor_ struct {
	// Define instance attributes.
	class_       FactorClassLike
	predicate_   PredicateLike
	cardinality_ CardinalityLike
}

// Attributes

func (v *factor_) GetClass() FactorClassLike {
	return v.class_
}

func (v *factor_) GetPredicate() PredicateLike {
	return v.predicate_
}

func (v *factor_) GetCardinality() CardinalityLike {
	return v.cardinality_
}

// Private
