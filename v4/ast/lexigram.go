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

import ()

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

func (c *lexigramClass_) MakeWithAttributes(
	comment string,
	lowercase string,
	pattern PatternLike,
	note string,
) LexigramLike {
	return &lexigram_{
		// Initialize instance attributes.
		class_:     c,
		comment_:   comment,
		lowercase_: lowercase,
		pattern_:   pattern,
		note_:      note,
	}
}

// INSTANCE METHODS

// Target

type lexigram_ struct {
	// Define instance attributes.
	class_     LexigramClassLike
	comment_   string
	lowercase_ string
	pattern_   PatternLike
	note_      string
}

// Attributes

func (v *lexigram_) GetClass() LexigramClassLike {
	return v.class_
}

func (v *lexigram_) GetComment() string {
	return v.comment_
}

func (v *lexigram_) GetLowercase() string {
	return v.lowercase_
}

func (v *lexigram_) GetPattern() PatternLike {
	return v.pattern_
}

func (v *lexigram_) GetNote() string {
	return v.note_
}

// Private
