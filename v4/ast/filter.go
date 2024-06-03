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

var filterClass = &filterClass_{
	// Initialize class constants.
}

// Function

func Filter() FilterClassLike {
	return filterClass
}

// CLASS METHODS

// Target

type filterClass_ struct {
	// Define class constants.
	// This class has no private constants.

}

// Constructors

func (c *filterClass_) MakeWithAttributes(
	inverted bool,
	atoms col.ListLike[AtomLike],
) FilterLike {
	return &filter_{
		// Initialize instance attributes.
		class_:    c,
		inverted_: inverted,
		atoms_:    atoms,
	}
}

// INSTANCE METHODS

// Target

type filter_ struct {
	// Define instance attributes.
	class_    FilterClassLike
	inverted_ bool
	atoms_    col.ListLike[AtomLike]
}

// Attributes

func (v *filter_) GetClass() FilterClassLike {
	return v.class_
}

func (v *filter_) IsInverted() bool {
	return v.inverted_
}

func (v *filter_) GetAtoms() col.ListLike[AtomLike] {
	return v.atoms_
}

// Private
