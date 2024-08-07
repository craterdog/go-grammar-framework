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

import (
	col "github.com/craterdog/go-collection-framework/v4"
)

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

func (c *predicateClass_) Make(
	identifier IdentifierLike,
	optionalCardinality CardinalityLike,
) PredicateLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(identifier):
		panic("The identifier attribute is required by this class.")
	default:
		return &predicate_{
			// Initialize instance attributes.
			class_:               c,
			identifier_:          identifier,
			optionalCardinality_: optionalCardinality,
		}
	}
}

// INSTANCE METHODS

// Target

type predicate_ struct {
	// Define instance attributes.
	class_               PredicateClassLike
	identifier_          IdentifierLike
	optionalCardinality_ CardinalityLike
}

// Attributes

func (v *predicate_) GetClass() PredicateClassLike {
	return v.class_
}

func (v *predicate_) GetIdentifier() IdentifierLike {
	return v.identifier_
}

func (v *predicate_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// Private
