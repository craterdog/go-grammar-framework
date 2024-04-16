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

var functionalsClass = &functionalsClass_{
	// This class has no private constants to initialize.
}

// Function

func Functionals() FunctionalsClassLike {
	return functionalsClass
}

// CLASS METHODS

// Target

type functionalsClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *functionalsClass_) MakeWithFunctionals(functionals col.ListLike[FunctionalLike]) FunctionalsLike {
	return &functionals_{
		functionals_: functionals,
	}
}

// Functions

// INSTANCE METHODS

// Target

type functionals_ struct {
	functionals_ col.ListLike[FunctionalLike]
}

// Attributes

func (v *functionals_) GetFunctionals() col.ListLike[FunctionalLike] {
	return v.functionals_
}

// Public

// Private
