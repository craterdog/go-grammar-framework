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

var selectiveClass = &selectiveClass_{
	// Initialize class constants.
}

// Function

func Selective() SelectiveClassLike {
	return selectiveClass
}

// CLASS METHODS

// Target

type selectiveClass_ struct {
	// Define class constants.
}

// Constructors

func (c *selectiveClass_) Make(alternatives abs.Sequential[AlternativeLike]) SelectiveLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(alternatives):
		panic("The alternatives attribute is required by this class.")
	default:
		return &selective_{
			// Initialize instance attributes.
			class_:        c,
			alternatives_: alternatives,
		}
	}
}

// INSTANCE METHODS

// Target

type selective_ struct {
	// Define instance attributes.
	class_        SelectiveClassLike
	alternatives_ abs.Sequential[AlternativeLike]
}

// Attributes

func (v *selective_) GetClass() SelectiveClassLike {
	return v.class_
}

func (v *selective_) GetAlternatives() abs.Sequential[AlternativeLike] {
	return v.alternatives_
}

// Private
