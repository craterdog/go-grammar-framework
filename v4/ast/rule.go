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
	optionalComment string,
	uppercase string,
	delimiter string,
	definition DefinitionLike,
	newlines abs.Sequential[string],
) RuleLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(uppercase):
		panic("The uppercase attribute is required by this class.")
	case col.IsUndefined(delimiter):
		panic("The delimiter attribute is required by this class.")
	case col.IsUndefined(definition):
		panic("The definition attribute is required by this class.")
	case col.IsUndefined(newlines):
		panic("The newlines attribute is required by this class.")
	default:
		return &rule_{
			// Initialize instance attributes.
			class_:           c,
			optionalComment_: optionalComment,
			uppercase_:       uppercase,
			delimiter_:       delimiter,
			definition_:      definition,
			newlines_:        newlines,
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
	delimiter_       string
	definition_      DefinitionLike
	newlines_        abs.Sequential[string]
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

func (v *rule_) GetDelimiter() string {
	return v.delimiter_
}

func (v *rule_) GetDefinition() DefinitionLike {
	return v.definition_
}

func (v *rule_) GetNewlines() abs.Sequential[string] {
	return v.newlines_
}

// Private
