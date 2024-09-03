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
}

// Constructors

func (c *filterClass_) Make(
	optionalExcluded string,
	characters abs.Sequential[CharacterLike],
) FilterLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(characters):
		panic("The characters attribute is required by this class.")
	default:
		return &filter_{
			// Initialize instance attributes.
			class_:            c,
			optionalExcluded_: optionalExcluded,
			characters_:       characters,
		}
	}
}

// INSTANCE METHODS

// Target

type filter_ struct {
	// Define instance attributes.
	class_            FilterClassLike
	optionalExcluded_ string
	characters_       abs.Sequential[CharacterLike]
}

// Attributes

func (v *filter_) GetClass() FilterClassLike {
	return v.class_
}

func (v *filter_) GetOptionalExcluded() string {
	return v.optionalExcluded_
}

func (v *filter_) GetCharacters() abs.Sequential[CharacterLike] {
	return v.characters_
}

// Private
