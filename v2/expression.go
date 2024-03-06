/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammars

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// CLASS ACCESS

// Reference

var expressionClass = &expressionClass_{
	// TBA - Assign constant values.
}

// Function

func Expression() ExpressionClassLike {
	return expressionClass
}

// CLASS METHODS

// Target

type expressionClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *expressionClass_) MakeWithAttributes(alternatives_ col.Sequential[AlternativeLike], multilined_ bool) ExpressionLike {
	return &expression_{
		alternatives_: alternatives_,
		multilined_:   multilined_,
	}
}

// Functions

// INSTANCE METHODS

// Target

type expression_ struct {
	alternatives_ col.Sequential[AlternativeLike]
	multilined_   bool
}

// Attributes

func (v *expression_) GetAlternatives() col.Sequential[AlternativeLike] {
	return v.alternatives_
}

func (v *expression_) IsMultilined() bool {
	return v.multilined_
}

// Public

// Private
