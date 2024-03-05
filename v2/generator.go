/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammars

import (
	cla "github.com/craterdog/go-class-framework/v2"
)

// CLASS ACCESS

// Reference

var generatorClass = &generatorClass_{
	// TBA - Assign constant values.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	var result_ = &generator_{}
	return result_
}

// Functions

// INSTANCE METHODS

// Target

type generator_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Public

func (v *generator_) GeneratePackage(
	name string,
	license string,
	comment string,
	grammar GrammarLike,
) cla.GoPNLike {
	var result_ cla.GoPNLike
	// TBA - Implement the method.
	return result_
}

// Private
