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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var bracketClass = &bracketClass_{
	// Initialize class constants.
}

// Function

func Bracket() BracketClassLike {
	return bracketClass
}

// CLASS METHODS

// Target

type bracketClass_ struct {
	// Define class constants.
}

// Constructors

func (c *bracketClass_) Make(
	factors abs.Sequential[FactorLike],
	cardinality CardinalityLike,
) BracketLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(factors):
		panic("The factors attribute is required by this class.")
	case col.IsUndefined(cardinality):
		panic("The cardinality attribute is required by this class.")
	default:
		return &bracket_{
			// Initialize instance attributes.
			class_:       c,
			factors_:     factors,
			cardinality_: cardinality,
		}
	}
}

// INSTANCE METHODS

// Target

type bracket_ struct {
	// Define instance attributes.
	class_       BracketClassLike
	factors_     abs.Sequential[FactorLike]
	cardinality_ CardinalityLike
}

// Attributes

func (v *bracket_) GetClass() BracketClassLike {
	return v.class_
}

func (v *bracket_) GetFactors() abs.Sequential[FactorLike] {
	return v.factors_
}

func (v *bracket_) GetCardinality() CardinalityLike {
	return v.cardinality_
}

// Private
