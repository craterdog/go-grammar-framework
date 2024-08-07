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
	delimiter string,
	characters abs.Sequential[CharacterLike],
	delimiter2 string,
) FilteredLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(delimiter):
		panic("The delimiter attribute is required by this class.")
	case col.IsUndefined(characters):
		panic("The characters attribute is required by this class.")
	case col.IsUndefined(delimiter2):
		panic("The delimiter2 attribute is required by this class.")
	default:
		return &filtered_{
			// Initialize instance attributes.
			class_:            c,
			optionalNegation_: optionalNegation,
			delimiter_:        delimiter,
			characters_:       characters,
			delimiter2_:       delimiter2,
		}
	}
}

// INSTANCE METHODS

// Target

type filtered_ struct {
	// Define instance attributes.
	class_            FilteredClassLike
	optionalNegation_ string
	delimiter_        string
	characters_       abs.Sequential[CharacterLike]
	delimiter2_       string
}

// Attributes

func (v *filtered_) GetClass() FilteredClassLike {
	return v.class_
}

func (v *filtered_) GetOptionalNegation() string {
	return v.optionalNegation_
}

func (v *filtered_) GetDelimiter() string {
	return v.delimiter_
}

func (v *filtered_) GetCharacters() abs.Sequential[CharacterLike] {
	return v.characters_
}

func (v *filtered_) GetDelimiter2() string {
	return v.delimiter2_
}

// Private
