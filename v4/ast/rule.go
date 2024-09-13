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
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
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
	uppercase string,
	definition DefinitionLike,
	newlines abs.Sequential[string],
) RuleLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(uppercase):
		panic("The uppercase attribute is required by this class.")
	case col.IsUndefined(definition):
		panic("The definition attribute is required by this class.")
	case col.IsUndefined(newlines):
		panic("The newlines attribute is required by this class.")
	default:
		return &rule_{
			// Initialize instance attributes.
			class_:      c,
			uppercase_:  uppercase,
			definition_: definition,
			newlines_:   newlines,
		}
	}
}

// INSTANCE METHODS

// Target

type rule_ struct {
	// Define instance attributes.
	class_      RuleClassLike
	uppercase_  string
	definition_ DefinitionLike
	newlines_   abs.Sequential[string]
}

// Attributes

func (v *rule_) GetClass() RuleClassLike {
	return v.class_
}

func (v *rule_) GetUppercase() string {
	return v.uppercase_
}

func (v *rule_) GetDefinition() DefinitionLike {
	return v.definition_
}

func (v *rule_) GetNewlines() abs.Sequential[string] {
	return v.newlines_
}

// Private
