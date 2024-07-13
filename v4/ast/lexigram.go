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
)

// CLASS ACCESS

// Reference

var lexigramClass = &lexigramClass_{
	// Initialize class constants.
}

// Function

func Lexigram() LexigramClassLike {
	return lexigramClass
}

// CLASS METHODS

// Target

type lexigramClass_ struct {
	// Define class constants.
}

// Constructors

func (c *lexigramClass_) Make(
	optionalComment string,
	lowercase string,
	pattern PatternLike,
	optionalNote string,
) LexigramLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(lowercase):
		panic("The lowercase attribute is required for each Lexigram.")
	case mod.IsUndefined(pattern):
		panic("The pattern attribute is required for each Lexigram.")
	default:
		return &lexigram_{
			// Initialize instance attributes.
			class_:           c,
			optionalComment_: optionalComment,
			lowercase_:       lowercase,
			pattern_:         pattern,
			optionalNote_:    optionalNote,
		}
	}
}

// INSTANCE METHODS

// Target

type lexigram_ struct {
	// Define instance attributes.
	class_           LexigramClassLike
	optionalComment_ string
	lowercase_       string
	pattern_         PatternLike
	optionalNote_    string
}

// Attributes

func (v *lexigram_) GetClass() LexigramClassLike {
	return v.class_
}

func (v *lexigram_) GetOptionalComment() string {
	return v.optionalComment_
}

func (v *lexigram_) GetLowercase() string {
	return v.lowercase_
}

func (v *lexigram_) GetPattern() PatternLike {
	return v.pattern_
}

func (v *lexigram_) GetOptionalNote() string {
	return v.optionalNote_
}

// Private
