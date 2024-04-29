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

var typesClass = &typesClass_{
	// Any private class constants should be initialized here.
}

// Function

func Types() TypesClassLike {
	return typesClass
}

// CLASS METHODS

// Target

type typesClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *typesClass_) MakeWithTypes(types col.ListLike[TypeLike]) TypesLike {
	return &types_{
		types_: types,
	}
}

// Functions

// INSTANCE METHODS

// Target

type types_ struct {
	class_ TypesClassLike
	types_ col.ListLike[TypeLike]
}

// Attributes

func (v *types_) GetClass() TypesClassLike {
	return v.class_
}

func (v *types_) GetTypes() col.ListLike[TypeLike] {
	return v.types_
}

// Public

// Private
