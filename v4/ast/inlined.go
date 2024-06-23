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
	col "github.com/craterdog/go-collection-framework/v4/collection"
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

func (c *inlinedClass_) MakeWithAttributes(
	factors col.ListLike[FactorLike],
	note string,
) InlinedLike {
	return &inlined_{
		// Initialize instance attributes.
		class_:   c,
		factors_: factors,
		note_:    note,
	}
}

// INSTANCE METHODS

// Target

type inlined_ struct {
	// Define instance attributes.
	class_   InlinedClassLike
	factors_ col.ListLike[FactorLike]
	note_    string
}

// Attributes

func (v *inlined_) GetClass() InlinedClassLike {
	return v.class_
}

func (v *inlined_) GetFactors() col.ListLike[FactorLike] {
	return v.factors_
}

func (v *inlined_) GetNote() string {
	return v.note_
}

// Private
