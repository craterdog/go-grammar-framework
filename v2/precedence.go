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

import ()

// CLASS ACCESS

// Reference

var precedenceClass = &precedenceClass_{
	// TBA - Assign constant values.
}

// Function

func Precedence() PrecedenceClassLike {
	return precedenceClass
}

// CLASS METHODS

// Target

type precedenceClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *precedenceClass_) MakeWithAttributes(expression ExpressionLike) PrecedenceLike {
	return &precedence_{
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type precedence_ struct {
	expression_ ExpressionLike
}

// Attributes

func (v *precedence_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
