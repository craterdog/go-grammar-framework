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

package cdsn

import ()

// CLASS ACCESS

// Reference

var precedenceClass = &precedenceClass_{
	// This class has no private constants to initialize.
}

// Function

func Precedence() PrecedenceClassLike {
	return precedenceClass
}

// CLASS METHODS

// Target

type precedenceClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *precedenceClass_) MakeWithExpression(expression ExpressionLike) PrecedenceLike {
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
