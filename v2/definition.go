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

import ()

// CLASS ACCESS

// Reference

var definitionClass = &definitionClass_{
	// TBA - Assign constant values.
}

// Function

func Definition() DefinitionClassLike {
	return definitionClass
}

// CLASS METHODS

// Target

type definitionClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *definitionClass_) MakeWithAttributes(symbol_ string, expression_ ExpressionLike) DefinitionLike {
	var result_ = &definition_{
		symbol_:     symbol_,
		expression_: expression_,
	}
	return result_
}

// Functions

// INSTANCE METHODS

// Target

type definition_ struct {
	expression_ ExpressionLike
	symbol_     string
}

// Attributes

func (v *definition_) GetExpression() ExpressionLike {
	return v.expression_
}

func (v *definition_) GetSymbol() string {
	return v.symbol_
}

// Public

// Private
