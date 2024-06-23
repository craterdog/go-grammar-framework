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

func (c *ruleClass_) MakeWithAttributes(
	comment string,
	uppercase string,
	expression ExpressionLike,
) RuleLike {
	return &rule_{
		// Initialize instance attributes.
		class_:      c,
		comment_:    comment,
		uppercase_:  uppercase,
		expression_: expression,
	}
}

// INSTANCE METHODS

// Target

type rule_ struct {
	// Define instance attributes.
	class_      RuleClassLike
	comment_    string
	uppercase_  string
	expression_ ExpressionLike
}

// Attributes

func (v *rule_) GetClass() RuleClassLike {
	return v.class_
}

func (v *rule_) GetComment() string {
	return v.comment_
}

func (v *rule_) GetUppercase() string {
	return v.uppercase_
}

func (v *rule_) GetExpression() ExpressionLike {
	return v.expression_
}

// Private
