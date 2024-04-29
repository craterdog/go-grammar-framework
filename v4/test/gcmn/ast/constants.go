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

var constantsClass = &constantsClass_{
	// Any private class constants should be initialized here.
}

// Function

func Constants() ConstantsClassLike {
	return constantsClass
}

// CLASS METHODS

// Target

type constantsClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *constantsClass_) MakeWithConstants(constants col.ListLike[ConstantLike]) ConstantsLike {
	return &constants_{
		constants_: constants,
	}
}

// Functions

// INSTANCE METHODS

// Target

type constants_ struct {
	class_ ConstantsClassLike
	constants_ col.ListLike[ConstantLike]
}

// Attributes

func (v *constants_) GetClass() ConstantsClassLike {
	return v.class_
}

func (v *constants_) GetConstants() col.ListLike[ConstantLike] {
	return v.constants_
}

// Public

// Private
