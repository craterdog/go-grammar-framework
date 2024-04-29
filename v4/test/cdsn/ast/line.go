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

var lineClass = &lineClass_{
	// Any private class constants should be initialized here.
}

// Function

func Line() LineClassLike {
	return lineClass
}

// CLASS METHODS

// Target

type lineClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *lineClass_) MakeWithAttributes(
	alternative AlternativeLike,
	note string,
) LineLike {
	return &line_{
		alternative_: alternative,
		note_: note,
	}
}

// Functions

// INSTANCE METHODS

// Target

type line_ struct {
	class_ LineClassLike
	alternative_ AlternativeLike
	note_ string
}

// Attributes

func (v *line_) GetClass() LineClassLike {
	return v.class_
}

func (v *line_) GetAlternative() AlternativeLike {
	return v.alternative_
}

func (v *line_) GetNote() string {
	return v.note_
}

// Public

// Private
