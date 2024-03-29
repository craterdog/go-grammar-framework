/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

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

func (c *definitionClass_) MakeWithAttributes(
	comment string,
	name string,
	expression ExpressionLike,
) DefinitionLike {
	return &definition_{
		comment_:    comment,
		name_:       name,
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type definition_ struct {
	comment_    string
	name_       string
	expression_ ExpressionLike
}

// Attributes

func (v *definition_) GetComment() string {
	return v.comment_
}

func (v *definition_) GetName() string {
	return v.name_
}

func (v *definition_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
