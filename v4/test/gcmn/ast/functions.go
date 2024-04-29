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

var functionsClass = &functionsClass_{
	// Any private class constants should be initialized here.
}

// Function

func Functions() FunctionsClassLike {
	return functionsClass
}

// CLASS METHODS

// Target

type functionsClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *functionsClass_) MakeWithFunctions(functions col.ListLike[FunctionLike]) FunctionsLike {
	return &functions_{
		functions_: functions,
	}
}

// Functions

// INSTANCE METHODS

// Target

type functions_ struct {
	class_ FunctionsClassLike
	functions_ col.ListLike[FunctionLike]
}

// Attributes

func (v *functions_) GetClass() FunctionsClassLike {
	return v.class_
}

func (v *functions_) GetFunctions() col.ListLike[FunctionLike] {
	return v.functions_
}

// Public

// Private
