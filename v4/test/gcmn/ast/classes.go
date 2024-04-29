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

var classesClass = &classesClass_{
	// Any private class constants should be initialized here.
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
	class_ ClassesClassLike
	classes_ col.ListLike[ClassLike]
}

// Attributes

func (v *classes_) GetClass() ClassesClassLike {
	return v.class_
}

func (v *classes_) GetClasses() col.ListLike[ClassLike] {
	return v.classes_
}

// Public

// Private
