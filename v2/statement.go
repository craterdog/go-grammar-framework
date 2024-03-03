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

var statementClass = &statementClass_{
	// This class does not initialize any constants.
}

// Function

func Statement() StatementClassLike {
	return statementClass
}

// CLASS METHODS

// Target

type statementClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *statementClass_) MakeFromComment(comment string) StatementLike {
	var statement = &statement_{
		// This class does not initialize any attributes.
	}
	statement.SetComment(comment)
	return statement
}

func (c *statementClass_) MakeFromDefinition(definition DefinitionLike) StatementLike {
	var statement = &statement_{
		// This class does not initialize any attributes.
	}
	statement.SetDefinition(definition)
	return statement
}

// INSTANCE METHODS

// Target

type statement_ struct {
	comment    string
	definition DefinitionLike
}

// Public

func (v *statement_) GetComment() string {
	return v.comment
}

func (v *statement_) GetDefinition() DefinitionLike {
	return v.definition
}

func (v *statement_) SetComment(comment string) {
	if len(comment) < 4 {
		var message = fmt.Sprintf(
			"Attempted to set an invalid comment:\n%v\n",
			comment,
		)
		panic(message)
	}
	v.comment = comment
	v.definition = nil
}

func (v *statement_) SetDefinition(definition DefinitionLike) {
	if definition == nil {
		panic("A definition must not be nil.")
	}
	v.comment = ""
	v.definition = definition
}
