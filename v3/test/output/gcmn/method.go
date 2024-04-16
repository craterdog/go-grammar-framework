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

var methodClass = &methodClass_{
	// This class has no private constants to initialize.
}

// Function

func Method() MethodClassLike {
	return methodClass
}

// CLASS METHODS

// Target

type methodClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *methodClass_) MakeWithAttributes(
	identifier string,
	parameters ParametersLike,
	result ResultLike,
) MethodLike {
	return &method_{
		identifier_: identifier,
		parameters_: parameters,
		result_: result,
	}
}

// Functions

// INSTANCE METHODS

// Target

type method_ struct {
	identifier_ string
	parameters_ ParametersLike
	result_ ResultLike
}

// Attributes

func (v *method_) GetIdentifier() string {
	return v.identifier_
}

func (v *method_) GetParameters() ParametersLike {
	return v.parameters_
}

func (v *method_) GetResult() ResultLike {
	return v.result_
}

// Public

// Private
