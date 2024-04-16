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

package test

import ()

// CLASS ACCESS

// Reference

var defaultClass = &defaultClass_{
	// This class has no private constants to initialize.
}

// Function

func Default() DefaultClassLike {
	return defaultClass
}

// CLASS METHODS

// Target

type defaultClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *defaultClass_) Make() DefaultLike {
	return &default_{}
}

// Functions

// INSTANCE METHODS

// Target

type default_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Public

// Private
