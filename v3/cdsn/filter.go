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

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var filterClass = &filterClass_{
	// TBA - Assign constant values.
}

// Function

func Filter() FilterClassLike {
	return filterClass
}

// CLASS METHODS

// Target

type filterClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *filterClass_) MakeWithAttributes(
	inverted bool,
	atoms col.ListLike[AtomLike],
) FilterLike {
	return &filter_{
		inverted_: inverted,
		atoms_:    atoms,
	}
}

// Functions

// INSTANCE METHODS

// Target

type filter_ struct {
	inverted_ bool
	atoms_    col.ListLike[AtomLike]
}

// Attributes

func (v *filter_) IsInverted() bool {
	return v.inverted_
}

func (v *filter_) GetAtoms() col.ListLike[AtomLike] {
	return v.atoms_
}

// Public

// Private
