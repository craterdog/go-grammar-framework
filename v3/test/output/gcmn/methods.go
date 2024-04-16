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

var methodsClass = &methodsClass_{
	// This class has no private constants to initialize.
}

// Function

func Methods() MethodsClassLike {
	return methodsClass
}

// CLASS METHODS

// Target

type methodsClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *methodsClass_) MakeWithMethods(methods col.ListLike[MethodLike]) MethodsLike {
	return &methods_{
		methods_: methods,
	}
}

// Functions

// INSTANCE METHODS

// Target

type methods_ struct {
	methods_ col.ListLike[MethodLike]
}

// Attributes

func (v *methods_) GetMethods() col.ListLike[MethodLike] {
	return v.methods_
}

// Public

// Private
