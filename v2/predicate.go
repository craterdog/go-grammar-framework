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

func (c *predicateClass_) MakeWithElement(element ElementLike) PredicateLike {
	return &predicate_{
		element_: element,
	}
}

func (c *predicateClass_) MakeWithInversion(inversion InversionLike) PredicateLike {
	return &predicate_{
		inversion_: inversion,
	}
}

func (c *predicateClass_) MakeWithPrecedence(precedence PrecedenceLike) PredicateLike {
	return &predicate_{
		precedence_: precedence,
	}
}

// Functions

// INSTANCE METHODS

// Target

type predicate_ struct {
	element_    ElementLike
	inversion_  InversionLike
	precedence_ PrecedenceLike
}

// Attributes

func (v *predicate_) GetElement() ElementLike {
	return v.element_
}

func (v *predicate_) GetInversion() InversionLike {
	return v.inversion_
}

func (v *predicate_) GetPrecedence() PrecedenceLike {
	return v.precedence_
}

// Public

// Private
