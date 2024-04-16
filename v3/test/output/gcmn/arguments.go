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

var argumentsClass = &argumentsClass_{
	// This class has no private constants to initialize.
}

// Function

func Arguments() ArgumentsClassLike {
	return argumentsClass
}

// CLASS METHODS

// Target

type argumentsClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *argumentsClass_) MakeWithAbstractions(abstractions col.ListLike[AbstractionLike]) ArgumentsLike {
	return &arguments_{
		abstractions_: abstractions,
	}
}

// Functions

// INSTANCE METHODS

// Target

type arguments_ struct {
	abstractions_ col.ListLike[AbstractionLike]
}

// Attributes

func (v *arguments_) GetAbstractions() col.ListLike[AbstractionLike] {
	return v.abstractions_
}

// Public

// Private
