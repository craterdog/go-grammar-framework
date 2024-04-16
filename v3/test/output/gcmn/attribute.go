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

var attributeClass = &attributeClass_{
	// This class has no private constants to initialize.
}

// Function

func Attribute() AttributeClassLike {
	return attributeClass
}

// CLASS METHODS

// Target

type attributeClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *attributeClass_) MakeWithAttributes(
	identifier string,
	parameter ParameterLike,
	abstraction AbstractionLike,
) AttributeLike {
	return &attribute_{
		identifier_: identifier,
		parameter_: parameter,
		abstraction_: abstraction,
	}
}

// Functions

// INSTANCE METHODS

// Target

type attribute_ struct {
	identifier_ string
	parameter_ ParameterLike
	abstraction_ AbstractionLike
}

// Attributes

func (v *attribute_) GetIdentifier() string {
	return v.identifier_
}

func (v *attribute_) GetParameter() ParameterLike {
	return v.parameter_
}

func (v *attribute_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Public

// Private
