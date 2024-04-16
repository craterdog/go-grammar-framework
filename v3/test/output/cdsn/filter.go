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
	// This class has no private constants to initialize.
}

// Function

func Filter() FilterClassLike {
	return filterClass
}

// CLASS METHODS

// Target

type filterClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *filterClass_) MakeWithAtoms(atoms col.ListLike[AtomLike]) FilterLike {
	return &filter_{
		atoms_: atoms,
	}
}

// Functions

// INSTANCE METHODS

// Target

type filter_ struct {
	atoms_ col.ListLike[AtomLike]
}

// Attributes

func (v *filter_) GetAtoms() col.ListLike[AtomLike] {
	return v.atoms_
}

// Public

// Private
