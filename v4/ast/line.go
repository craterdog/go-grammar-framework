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
)

// CLASS ACCESS

// Reference

var lineClass = &lineClass_{
	// Initialize class constants.
}

// Function

func Line() LineClassLike {
	return lineClass
}

// CLASS METHODS

// Target

type lineClass_ struct {
	// Define class constants.
}

// Constructors

func (c *lineClass_) Make(
	newline string,
	identifier IdentifierLike,
	optionalNote string,
) LineLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(newline):
		panic("The newline attribute is required by this class.")
	case col.IsUndefined(identifier):
		panic("The identifier attribute is required by this class.")
	default:
		return &line_{
			// Initialize instance attributes.
			class_:        c,
			newline_:      newline,
			identifier_:   identifier,
			optionalNote_: optionalNote,
		}
	}
}

// INSTANCE METHODS

// Target

type line_ struct {
	// Define instance attributes.
	class_        LineClassLike
	newline_      string
	identifier_   IdentifierLike
	optionalNote_ string
}

// Attributes

func (v *line_) GetClass() LineClassLike {
	return v.class_
}

func (v *line_) GetNewline() string {
	return v.newline_
}

func (v *line_) GetIdentifier() IdentifierLike {
	return v.identifier_
}

func (v *line_) GetOptionalNote() string {
	return v.optionalNote_
}

// Private
