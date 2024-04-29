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

var precedenceClass = &precedenceClass_{
	// Any private class constants should be initialized here.
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
	class_ PrecedenceClassLike
	expression_ ExpressionLike
}

// Attributes

func (v *precedence_) GetClass() PrecedenceClassLike {
	return v.class_
}

func (v *precedence_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
