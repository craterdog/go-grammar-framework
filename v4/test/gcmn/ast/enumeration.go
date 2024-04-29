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
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var enumerationClass = &enumerationClass_{
	// Any private class constants should be initialized here.
}

// Function

func Enumeration() EnumerationClassLike {
	return enumerationClass
}

// CLASS METHODS

// Target

type enumerationClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *enumerationClass_) MakeWithAttributes(
	parameter ParameterLike,
	identifiers col.ListLike[string],
) EnumerationLike {
	return &enumeration_{
		parameter_: parameter,
		identifiers_: identifiers,
	}
}

// Functions

// INSTANCE METHODS

// Target

type enumeration_ struct {
	class_ EnumerationClassLike
	parameter_ ParameterLike
	identifiers_ col.ListLike[string]
}

// Attributes

func (v *enumeration_) GetClass() EnumerationClassLike {
	return v.class_
}

func (v *enumeration_) GetParameter() ParameterLike {
	return v.parameter_
}

func (v *enumeration_) GetIdentifiers() col.ListLike[string] {
	return v.identifiers_
}

// Public

// Private
