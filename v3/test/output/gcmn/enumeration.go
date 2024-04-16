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

package gcmn

import ()

// CLASS ACCESS

// Reference

var enumerationClass = &enumerationClass_{
	// This class has no private constants to initialize.
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
	identifier string,
) EnumerationLike {
	return &enumeration_{
		parameter_: parameter,
		identifier_: identifier,
	}
}

// Functions

// INSTANCE METHODS

// Target

type enumeration_ struct {
	parameter_ ParameterLike
	identifier_ string
}

// Attributes

func (v *enumeration_) GetParameter() ParameterLike {
	return v.parameter_
}

func (v *enumeration_) GetIdentifier() string {
	return v.identifier_
}

// Public

// Private
