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
	var result_ = &statement_{
		comment_: comment_,
	}
	return result_
}

func (c *statementClass_) MakeWithDefinition(definition_ DefinitionLike) StatementLike {
	var result_ = &statement_{
		definition_: definition_,
	}
	return result_
}

// Functions

// INSTANCE METHODS

// Target

type statement_ struct {
	comment_ string
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
