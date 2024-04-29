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

var constructorsClass = &constructorsClass_{
	// Any private class constants should be initialized here.
}

// Function

func Constructors() ConstructorsClassLike {
	return constructorsClass
}

// CLASS METHODS

// Target

type constructorsClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *constructorsClass_) MakeWithConstructors(constructors col.ListLike[ConstructorLike]) ConstructorsLike {
	return &constructors_{
		constructors_: constructors,
	}
}

// Functions

// INSTANCE METHODS

// Target

type constructors_ struct {
	class_ ConstructorsClassLike
	constructors_ col.ListLike[ConstructorLike]
}

// Attributes

func (v *constructors_) GetClass() ConstructorsClassLike {
	return v.class_
}

func (v *constructors_) GetConstructors() col.ListLike[ConstructorLike] {
	return v.constructors_
}

// Public

// Private
