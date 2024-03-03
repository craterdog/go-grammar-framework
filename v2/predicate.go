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

var predicateClass = &predicateClass_{
	// This class does not initialize any constants.
}

// Function

func Predicate() PredicateClassLike {
	return predicateClass
}

// CLASS METHODS

// Target

type predicateClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *predicateClass_) Make(
	assertion AssertionLike,
	isInverted bool,
) PredicateLike {
	var predicate = &predicate_{
		// This class does not initialize any attributes.
	}
	predicate.SetAssertion(assertion)
	predicate.SetInverted(isInverted)
	return predicate
}

// INSTANCE METHODS

// Target

type predicate_ struct {
	assertion  AssertionLike
	isInverted bool
}

// Public

func (v *predicate_) GetAssertion() AssertionLike {
	return v.assertion
}

func (v *predicate_) IsInverted() bool {
	return v.isInverted
}

func (v *predicate_) SetAssertion(assertion AssertionLike) {
	if assertion == nil {
		panic("An assertion must not be nil.")
	}
	v.assertion = assertion
}

func (v *predicate_) SetInverted(isInverted bool) {
	v.isInverted = isInverted
}
