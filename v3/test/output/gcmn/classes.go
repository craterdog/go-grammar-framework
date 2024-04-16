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

var classesClass = &classesClass_{
	// This class has no private constants to initialize.
}

// Function

func Classes() ClassesClassLike {
	return classesClass
}

// CLASS METHODS

// Target

type classesClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *classesClass_) MakeWithClasses(classes col.ListLike[ClassLike]) ClassesLike {
	return &classes_{
		classes_: classes,
	}
}

// Functions

// INSTANCE METHODS

// Target

type classes_ struct {
	classes_ col.ListLike[ClassLike]
}

// Attributes

func (v *classes_) GetClasses() col.ListLike[ClassLike] {
	return v.classes_
}

// Public

// Private
