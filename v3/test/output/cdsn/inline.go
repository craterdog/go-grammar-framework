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

var inlineClass = &inlineClass_{
	// This class has no private constants to initialize.
}

// Function

func Inline() InlineClassLike {
	return inlineClass
}

// CLASS METHODS

// Target

type inlineClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *inlineClass_) MakeWithAttributes(
	alternatives col.ListLike[AlternativeLike],
	note string,
) InlineLike {
	return &inline_{
		alternatives_: alternatives,
		note_: note,
	}
}

// Functions

// INSTANCE METHODS

// Target

type inline_ struct {
	alternatives_ col.ListLike[AlternativeLike]
	note_ string
}

// Attributes

func (v *inline_) GetAlternatives() col.ListLike[AlternativeLike] {
	return v.alternatives_
}

func (v *inline_) GetNote() string {
	return v.note_
}

// Public

// Private
