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

var predicateClass = &predicateClass_{
	// This class has no private constants to initialize.
}

// Function

func Predicate() PredicateClassLike {
	return predicateClass
}

// CLASS METHODS

// Target

type predicateClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *predicateClass_) MakeWithAtom(atom AtomLike) PredicateLike {
	return &predicate_{
		atom_: atom,
	}
}

func (c *predicateClass_) MakeWithElement(element ElementLike) PredicateLike {
	return &predicate_{
		element_: element,
	}
}

func (c *predicateClass_) MakeWithFilter(filter FilterLike) PredicateLike {
	return &predicate_{
		filter_: filter,
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
	atom_ AtomLike
	element_ ElementLike
	filter_ FilterLike
	precedence_ PrecedenceLike
}

// Attributes

func (v *predicate_) GetAtom() AtomLike {
	return v.atom_
}

func (v *predicate_) GetElement() ElementLike {
	return v.element_
}

func (v *predicate_) GetFilter() FilterLike {
	return v.filter_
}

func (v *predicate_) GetPrecedence() PrecedenceLike {
	return v.precedence_
}

// Public

// Private
