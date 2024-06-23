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

var predicateClass = &predicateClass_{
	// Initialize class constants.
}

// Function

func Predicate() PredicateClassLike {
	return predicateClass
}

// CLASS METHODS

// Target

type predicateClass_ struct {
	// Define class constants.
}

// Constructors

func (c *predicateClass_) MakeWithLowercase(lowercase string) PredicateLike {
	return &predicate_{
		// Initialize instance attributes.
		class_: c,
		any_:   lowercase,
	}
}

func (c *predicateClass_) MakeWithUppercase(uppercase string) PredicateLike {
	return &predicate_{
		// Initialize instance attributes.
		class_: c,
		any_:   uppercase,
	}
}

func (c *predicateClass_) MakeWithIntrinsic(intrinsic string) PredicateLike {
	return &predicate_{
		// Initialize instance attributes.
		class_: c,
		any_:   intrinsic,
	}
}

func (c *predicateClass_) MakeWithLiteral(literal string) PredicateLike {
	return &predicate_{
		// Initialize instance attributes.
		class_: c,
		any_:   literal,
	}
}

// INSTANCE METHODS

// Target

type predicate_ struct {
	// Define instance attributes.
	class_ PredicateClassLike
	any_   any
}

// Attributes

func (v *predicate_) GetClass() PredicateClassLike {
	return v.class_
}

func (v *predicate_) GetAny() any {
	return v.any_
}

// Private
