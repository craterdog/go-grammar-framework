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

var expressionClass = &expressionClass_{
	// Initialize class constants.
}

// Function

func Expression() ExpressionClassLike {
	return expressionClass
}

// CLASS METHODS

// Target

type expressionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *expressionClass_) Make(any_ any) ExpressionLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(any_):
		panic("The any_ attribute is required for each Expression.")
	default:
		return &expression_{
			// Initialize instance attributes.
			class_: c,
			any_:   any_,
		}
	}
}

// INSTANCE METHODS

// Target

type expression_ struct {
	// Define instance attributes.
	class_ ExpressionClassLike
	any_   any
}

// Attributes

func (v *expression_) GetClass() ExpressionClassLike {
	return v.class_
}

func (v *expression_) GetAny() any {
	return v.any_
}

// Private
