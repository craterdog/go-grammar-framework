/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammar

// CLASS ACCESS

// Reference

var precedenceClass = &precedenceClass_{
	// This class does not initialize any constants.
}

// Function

func Precedence() PrecedenceClassLike {
	return precedenceClass
}

// CLASS METHODS

// Target

type precedenceClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *precedenceClass_) Make(
	expression ExpressionLike,
) PrecedenceLike {
	var precedence = &precedence_{
		// This class does not initialize any attributes.
	}
	precedence.SetExpression(expression)
	return precedence
}

// INSTANCE METHODS

// Target

type precedence_ struct {
	expression ExpressionLike
}

// Public

func (v *precedence_) GetExpression() ExpressionLike {
	return v.expression
}

func (v *precedence_) SetExpression(expression ExpressionLike) {
	if expression == nil {
		panic("The expression within a precedence cannot be nil.")
	}
	v.expression = expression
}
