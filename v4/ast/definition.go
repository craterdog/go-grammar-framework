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

var definitionClass = &definitionClass_{
	// Initialize class constants.
}

// Function

func Definition() DefinitionClassLike {
	return definitionClass
}

// CLASS METHODS

// Target

type definitionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *definitionClass_) Make(any_ any) DefinitionLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(any_):
		panic("The any_ attribute is required for each Definition.")
	default:
		return &definition_{
			// Initialize instance attributes.
			class_: c,
			any_:   any_,
		}
	}
}

// INSTANCE METHODS

// Target

type definition_ struct {
	// Define instance attributes.
	class_ DefinitionClassLike
	any_   any
}

// Attributes

func (v *definition_) GetClass() DefinitionClassLike {
	return v.class_
}

func (v *definition_) GetAny() any {
	return v.any_
}

// Private
