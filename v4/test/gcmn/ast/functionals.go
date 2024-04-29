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

import (
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var functionalsClass = &functionalsClass_{
	// Any private class constants should be initialized here.
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
	class_ FunctionalsClassLike
	functionals_ col.ListLike[FunctionalLike]
}

// Attributes

func (v *functionals_) GetClass() FunctionalsClassLike {
	return v.class_
}

func (v *functionals_) GetFunctionals() col.ListLike[FunctionalLike] {
	return v.functionals_
}

// Public

// Private
