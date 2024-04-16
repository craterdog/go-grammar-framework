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

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var parametersClass = &parametersClass_{
	// This class has no private constants to initialize.
}

// Function

func Parameters() ParametersClassLike {
	return parametersClass
}

// CLASS METHODS

// Target

type parametersClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *parametersClass_) MakeWithParameters(parameters col.ListLike[ParameterLike]) ParametersLike {
	return &parameters_{
		parameters_: parameters,
	}
}

// Functions

// INSTANCE METHODS

// Target

type parameters_ struct {
	parameters_ col.ListLike[ParameterLike]
}

// Attributes

func (v *parameters_) GetParameters() col.ListLike[ParameterLike] {
	return v.parameters_
}

// Public

// Private
