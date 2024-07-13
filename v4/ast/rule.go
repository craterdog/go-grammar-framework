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

var ruleClass = &ruleClass_{
	// Initialize class constants.
}

// Function

func Rule() RuleClassLike {
	return ruleClass
}

// CLASS METHODS

// Target

type ruleClass_ struct {
	// Define class constants.
}

// Constructors

func (c *ruleClass_) Make(
	optionalComment string,
	uppercase string,
	expression ExpressionLike,
) RuleLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(uppercase):
		panic("The uppercase attribute is required for each Rule.")
	case mod.IsUndefined(expression):
		panic("The expression attribute is required for each Rule.")
	default:
		return &rule_{
			// Initialize instance attributes.
			class_:           c,
			optionalComment_: optionalComment,
			uppercase_:       uppercase,
			expression_:      expression,
		}
	}
}

// INSTANCE METHODS

// Target

type rule_ struct {
	// Define instance attributes.
	class_           RuleClassLike
	optionalComment_ string
	uppercase_       string
	expression_      ExpressionLike
}

// Attributes

func (v *rule_) GetClass() RuleClassLike {
	return v.class_
}

func (v *rule_) GetOptionalComment() string {
	return v.optionalComment_
}

func (v *rule_) GetUppercase() string {
	return v.uppercase_
}

func (v *rule_) GetExpression() ExpressionLike {
	return v.expression_
}

// Private
