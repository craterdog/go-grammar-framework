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
	col "github.com/craterdog/go-collection-framework/v3"
)

// CLASS ACCESS

// Reference

var grammarClass = &grammarClass_{
	// This class does not initialize any constants.
}

// Function

func Grammar() GrammarClassLike {
	return grammarClass
}

// CLASS METHODS

// Target

type grammarClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *grammarClass_) Make(statements col.Sequential[StatementLike]) GrammarLike {
	var grammar = &grammar_{
		// This class does not initialize any attributes.
	}
	grammar.SetStatements(statements)
	return grammar
}

// INSTANCE METHODS

// Target

type grammar_ struct {
	statements col.Sequential[StatementLike]
}

// Public

func (v *grammar_) GetStatements() col.Sequential[StatementLike] {
	return v.statements
}

func (v *grammar_) SetStatements(statements col.Sequential[StatementLike]) {
	if statements == nil || statements.IsEmpty() {
		panic("An grammar must have at least one statement.")
	}
	v.statements = statements
}
