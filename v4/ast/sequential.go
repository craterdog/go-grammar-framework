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

var sequentialClass = &sequentialClass_{
	// Initialize class constants.
}

// Function

func Sequential() SequentialClassLike {
	return sequentialClass
}

// CLASS METHODS

// Target

type sequentialClass_ struct {
	// Define class constants.
}

// Constructors

func (c *sequentialClass_) Make(parts abs.Sequential[PartLike]) SequentialLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(parts):
		panic("The parts attribute is required by this class.")
	default:
		return &sequential_{
			// Initialize instance attributes.
			class_: c,
			parts_: parts,
		}
	}
}

// INSTANCE METHODS

// Target

type sequential_ struct {
	// Define instance attributes.
	class_ SequentialClassLike
	parts_ abs.Sequential[PartLike]
}

// Attributes

func (v *sequential_) GetClass() SequentialClassLike {
	return v.class_
}

func (v *sequential_) GetParts() abs.Sequential[PartLike] {
	return v.parts_
}

// Private
