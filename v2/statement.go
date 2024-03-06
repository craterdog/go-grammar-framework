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

var statementClass = &statementClass_{
	// TBA - Assign constant values.
}

// Function

func Statement() StatementClassLike {
	return statementClass
}

// CLASS METHODS

// Target

type statementClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *statementClass_) MakeWithComment(comment_ string) StatementLike {
	return &statement_{
		comment_: comment_,
	}
}

func (c *statementClass_) MakeWithDefinition(definition_ DefinitionLike) StatementLike {
	return &statement_{
		definition_: definition_,
	}
}

// Functions

// INSTANCE METHODS

// Target

type statement_ struct {
	comment_    string
	definition_ DefinitionLike
}

// Attributes

func (v *statement_) GetComment() string {
	return v.comment_
}

func (v *statement_) GetDefinition() DefinitionLike {
	return v.definition_
}

// Public

// Private
