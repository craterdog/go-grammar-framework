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

import (
	fmt "fmt"
)

// CLASS ACCESS

// Reference

var definitionClass = &definitionClass_{
	// This class does not initialize any constants.
}

// Function

func Definition() DefinitionClassLike {
	return definitionClass
}

// CLASS METHODS

// Target

type definitionClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *definitionClass_) Make(
	symbol string,
	expression ExpressionLike,
) DefinitionLike {
	var definition = &definition_{
		// This class does not initialize any attributes.
	}
	definition.SetSymbol(symbol)
	definition.SetExpression(expression)
	return definition
}

// INSTANCE METHODS

// Target

type definition_ struct {
	expression ExpressionLike
	symbol     string
}

// Public

func (v *definition_) GetExpression() ExpressionLike {
	return v.expression
}

func (v *definition_) GetSymbol() string {
	return v.symbol
}

func (v *definition_) SetExpression(expression ExpressionLike) {
	if expression == nil {
		panic("An expression cannot be nil.")
	}
	v.expression = expression
}

func (v *definition_) SetSymbol(symbol string) {
	if len(symbol) < 2 {
		var message = fmt.Sprintf("An invalid symbol was found:\n    %v\n", symbol)
		panic(message)
	}
	v.symbol = symbol
}
