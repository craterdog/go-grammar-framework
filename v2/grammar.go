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

var grammarClass = &grammarClass_{
	// TBA - Assign constant values.
}

// Function

func Grammar() GrammarClassLike {
	return grammarClass
}

// CLASS METHODS

// Target

type grammarClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *grammarClass_) MakeWithAttributes(comment string, statements col.Sequential[StatementLike]) GrammarLike {
	return &grammar_{
		comment_:    comment,
		statements_: statements,
	}
}

// Functions

// INSTANCE METHODS

// Target

type grammar_ struct {
	comment_    string
	statements_ col.Sequential[StatementLike]
}

// Attributes

func (v *grammar_) GetComment() string {
	return v.comment_
}

func (v *grammar_) GetStatements() col.Sequential[StatementLike] {
	return v.statements_
}

// Public

// Private
