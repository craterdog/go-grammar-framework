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

var referenceClass = &referenceClass_{
	// Initialize class constants.
}

// Function

func Reference() ReferenceClassLike {
	return referenceClass
}

// CLASS METHODS

// Target

type referenceClass_ struct {
	// Define class constants.
}

// Constructors

func (c *referenceClass_) Make(
	identifier IdentifierLike,
	optionalCardinality CardinalityLike,
) ReferenceLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(identifier):
		panic("The identifier attribute is required by this class.")
	default:
		return &reference_{
			// Initialize instance attributes.
			class_:               c,
			identifier_:          identifier,
			optionalCardinality_: optionalCardinality,
		}
	}
}

// INSTANCE METHODS

// Target

type reference_ struct {
	// Define instance attributes.
	class_               ReferenceClassLike
	identifier_          IdentifierLike
	optionalCardinality_ CardinalityLike
}

// Attributes

func (v *reference_) GetClass() ReferenceClassLike {
	return v.class_
}

func (v *reference_) GetIdentifier() IdentifierLike {
	return v.identifier_
}

func (v *reference_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// Private
