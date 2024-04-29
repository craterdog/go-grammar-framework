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

import ()

// CLASS ACCESS

// Reference

var resultClass = &resultClass_{
	// Any private class constants should be initialized here.
}

// Function

func Result() ResultClassLike {
	return resultClass
}

// CLASS METHODS

// Target

type resultClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *resultClass_) MakeWithAbstraction(abstraction AbstractionLike) ResultLike {
	return &result_{
		abstraction_: abstraction,
	}
}

func (c *resultClass_) MakeWithParameters(parameters ParametersLike) ResultLike {
	return &result_{
		parameters_: parameters,
	}
}

// Functions

// INSTANCE METHODS

// Target

type result_ struct {
	class_ ResultClassLike
	abstraction_ AbstractionLike
	parameters_ ParametersLike
}

// Attributes

func (v *result_) GetClass() ResultClassLike {
	return v.class_
}

func (v *result_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *result_) GetParameters() ParametersLike {
	return v.parameters_
}

// Public

// Private
