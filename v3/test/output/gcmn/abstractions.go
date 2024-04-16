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

var abstractionsClass = &abstractionsClass_{
	// This class has no private constants to initialize.
}

// Function

func Abstractions() AbstractionsClassLike {
	return abstractionsClass
}

// CLASS METHODS

// Target

type abstractionsClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *abstractionsClass_) MakeWithAbstractions(abstractions col.ListLike[AbstractionLike]) AbstractionsLike {
	return &abstractions_{
		abstractions_: abstractions,
	}
}

// Functions

// INSTANCE METHODS

// Target

type abstractions_ struct {
	abstractions_ col.ListLike[AbstractionLike]
}

// Attributes

func (v *abstractions_) GetAbstractions() col.ListLike[AbstractionLike] {
	return v.abstractions_
}

// Public

// Private
