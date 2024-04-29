/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
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

var rootClass = &rootClass_{
	// Any private class constants should be initialized here.
}

// Function

func Root() RootClassLike {
	return rootClass
}

// CLASS METHODS

// Target

type rootClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *rootClass_) MakeWithLeafs(leafs col.ListLike[string]) RootLike {
	return &root_{
		leafs_: leafs,
	}
}

// Functions

// INSTANCE METHODS

// Target

type root_ struct {
	class_ RootClassLike
	leafs_ col.ListLike[string]
}

// Attributes

func (v *root_) GetClass() RootClassLike {
	return v.class_
}

func (v *root_) GetLeafs() col.ListLike[string] {
	return v.leafs_
}

// Public

// Private
