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

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var exampleClass = &exampleClass_{
	// This class has no private constants to initialize.
}

// Function

func Example() ExampleClassLike {
	return exampleClass
}

// CLASS METHODS

// Target

type exampleClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *exampleClass_) MakeWithDefault(default_ DefaultLike) ExampleLike {
	return &example_{
		default_: default_,
	}
}

func (c *exampleClass_) MakeWithPrimitive(primitive PrimitiveLike) ExampleLike {
	return &example_{
		primitive_: primitive,
	}
}

func (c *exampleClass_) MakeWithLists(lists col.ListLike[ListLike]) ExampleLike {
	return &example_{
		lists_: lists,
	}
}

// Functions

// INSTANCE METHODS

// Target

type example_ struct {
	default_ DefaultLike
	primitive_ PrimitiveLike
	lists_ col.ListLike[ListLike]
}

// Attributes

func (v *example_) GetDefault() DefaultLike {
	return v.default_
}

func (v *example_) GetPrimitive() PrimitiveLike {
	return v.primitive_
}

func (v *example_) GetLists() col.ListLike[ListLike] {
	return v.lists_
}

// Public

// Private
