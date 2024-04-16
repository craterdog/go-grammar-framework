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

var expressionClass = &expressionClass_{
	// This class has no private constants to initialize.
}

// Function

func Expression() ExpressionClassLike {
	return expressionClass
}

// CLASS METHODS

// Target

type expressionClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *expressionClass_) MakeWithInline(inline InlineLike) ExpressionLike {
	return &expression_{
		inline_: inline,
	}
}

func (c *expressionClass_) MakeWithMultiline(multiline MultilineLike) ExpressionLike {
	return &expression_{
		multiline_: multiline,
	}
}

// Functions

// INSTANCE METHODS

// Target

type expression_ struct {
	inline_ InlineLike
	multiline_ MultilineLike
}

// Attributes

func (v *expression_) GetInline() InlineLike {
	return v.inline_
}

func (v *expression_) GetMultiline() MultilineLike {
	return v.multiline_
}

// Public

// Private
