/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package grammars

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// CLASS ACCESS

// Reference

var inlineClass = &inlineClass_{
	// TBA - Assign constant values.
}

// Function

func Inline() InlineClassLike {
	return inlineClass
}

// CLASS METHODS

// Target

type inlineClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *inlineClass_) MakeWithAttributes(alternatives col.Sequential[AlternativeLike], note string) InlineLike {
	return &inline_{
		alternatives_: alternatives,
		note_:         note,
	}
}

// Functions

// INSTANCE METHODS

// Target

type inline_ struct {
	alternatives_ col.Sequential[AlternativeLike]
	note_         string
}

// Attributes

func (v *inline_) GetAlternatives() col.Sequential[AlternativeLike] {
	return v.alternatives_
}

func (v *inline_) GetNote() string {
	return v.note_
}

// Public

// Private
