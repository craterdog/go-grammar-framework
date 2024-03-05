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

import ()

// CLASS ACCESS

// Reference

var predicateClass = &predicateClass_{
	// TBA - Assign constant values.
}

// Function

func Predicate() PredicateClassLike {
	return predicateClass
}

// CLASS METHODS

// Target

type predicateClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *predicateClass_) MakeWithAttributes(assertion_ AssertionLike, inverted_ bool) PredicateLike {
	var result_ = &predicate_{
		assertion_: assertion_,
		inverted_: inverted_,
	}
	return result_
}

// Functions

// INSTANCE METHODS

// Target

type predicate_ struct {
	assertion_ AssertionLike
	inverted_ bool
}

// Attributes

func (v *predicate_) GetAssertion() AssertionLike {
	return v.assertion_
}

func (v *predicate_) GetInverted() bool {
	return v.inverted_
}

// Public

func (v *predicate_) IsInverted() bool {
	return v.inverted_
}

// Private
