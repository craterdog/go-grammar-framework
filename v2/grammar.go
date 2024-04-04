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

func (c *grammarClass_) MakeWithAttributes(headers col.ListLike[HeaderLike], definitions col.ListLike[DefinitionLike]) GrammarLike {
	return &grammar_{
		headers_:     headers,
		definitions_: definitions,
	}
}

// Functions

// INSTANCE METHODS

// Target

type grammar_ struct {
	headers_     col.ListLike[HeaderLike]
	definitions_ col.ListLike[DefinitionLike]
}

// Attributes

func (v *grammar_) GetHeaders() col.ListLike[HeaderLike] {
	return v.headers_
}

func (v *grammar_) GetDefinitions() col.ListLike[DefinitionLike] {
	return v.definitions_
}

// Public

// Private
