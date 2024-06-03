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
	// Initialize class constants.
}

// Function

func Precedence() PrecedenceClassLike {
	return precedenceClass
}

// CLASS METHODS

// Target

type precedenceClass_ struct {
	// Define class constants.
}

// Constructors

func (c *precedenceClass_) MakeWithExpression(expression ExpressionLike) PrecedenceLike {
	return &precedence_{
		// Initialize instance attributes.
		class_:      c,
		expression_: expression,
	}
}

// INSTANCE METHODS

// Target

type precedence_ struct {
	// Define instance attributes.
	class_      PrecedenceClassLike
	expression_ ExpressionLike
}

// Attributes

func (v *precedence_) GetClass() PrecedenceClassLike {
	return v.class_
}

func (v *precedence_) GetExpression() ExpressionLike {
	return v.expression_
}

// Private
