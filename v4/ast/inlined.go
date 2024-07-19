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

var inlinedClass = &inlinedClass_{
	// Initialize class constants.
}

// Function

func Inlined() InlinedClassLike {
	return inlinedClass
}

// CLASS METHODS

// Target

type inlinedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *inlinedClass_) Make(
	factors abs.Sequential[FactorLike],
	optionalNote string,
) InlinedLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(factors):
		panic("The factors attribute is required by this class.")
	default:
		return &inlined_{
			// Initialize instance attributes.
			class_:        c,
			factors_:      factors,
			optionalNote_: optionalNote,
		}
	}
}

// INSTANCE METHODS

// Target

type inlined_ struct {
	// Define instance attributes.
	class_        InlinedClassLike
	factors_      abs.Sequential[FactorLike]
	optionalNote_ string
}

// Attributes

func (v *inlined_) GetClass() InlinedClassLike {
	return v.class_
}

func (v *inlined_) GetFactors() abs.Sequential[FactorLike] {
	return v.factors_
}

func (v *inlined_) GetOptionalNote() string {
	return v.optionalNote_
}

// Private
