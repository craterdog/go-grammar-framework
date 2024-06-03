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

var expressionClass = &expressionClass_{
	// Initialize class constants.
}

// Function

func Expression() ExpressionClassLike {
	return expressionClass
}

// CLASS METHODS

// Target

type expressionClass_ struct {
	// Define class constants.
	// This class has no private constants.

}

// Constructors

func (c *expressionClass_) MakeWithInline(inline InlineLike) ExpressionLike {
	return &expression_{
		// Initialize instance attributes.
		class_:  c,
		inline_: inline,
	}
}

func (c *expressionClass_) MakeWithMultiline(multiline MultilineLike) ExpressionLike {
	return &expression_{
		// Initialize instance attributes.
		class_:     c,
		multiline_: multiline,
	}
}

// INSTANCE METHODS

// Target

type expression_ struct {
	// Define instance attributes.
	class_     ExpressionClassLike
	inline_    InlineLike
	multiline_ MultilineLike
}

// Attributes

func (v *expression_) GetClass() ExpressionClassLike {
	return v.class_
}

func (v *expression_) GetInline() InlineLike {
	return v.inline_
}

func (v *expression_) GetMultiline() MultilineLike {
	return v.multiline_
}

// Private
