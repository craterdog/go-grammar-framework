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
	mod "github.com/craterdog/go-collection-framework/v4"
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var filteredClass = &filteredClass_{
	// Initialize class constants.
}

// Function

func Filtered() FilteredClassLike {
	return filteredClass
}

// CLASS METHODS

// Target

type filteredClass_ struct {
	// Define class constants.
}

// Constructors

func (c *filteredClass_) Make(
	optionalNegation string,
	characters col.Sequential[CharacterLike],
) FilteredLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(characters):
		panic("The characters attribute is required for each Filtered.")
	default:
		return &filtered_{
			// Initialize instance attributes.
			class_:            c,
			optionalNegation_: optionalNegation,
			characters_:       characters,
		}
	}
}

// INSTANCE METHODS

// Target

type filtered_ struct {
	// Define instance attributes.
	class_            FilteredClassLike
	optionalNegation_ string
	characters_       col.Sequential[CharacterLike]
}

// Attributes

func (v *filtered_) GetClass() FilteredClassLike {
	return v.class_
}

func (v *filtered_) GetOptionalNegation() string {
	return v.optionalNegation_
}

func (v *filtered_) GetCharacters() col.Sequential[CharacterLike] {
	return v.characters_
}

// Private
