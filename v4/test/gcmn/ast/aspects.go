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

var aspectsClass = &aspectsClass_{
	// Any private class constants should be initialized here.
}

// Function

func Aspects() AspectsClassLike {
	return aspectsClass
}

// CLASS METHODS

// Target

type aspectsClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *aspectsClass_) MakeWithAspects(aspects col.ListLike[AspectLike]) AspectsLike {
	return &aspects_{
		aspects_: aspects,
	}
}

// Functions

// INSTANCE METHODS

// Target

type aspects_ struct {
	class_ AspectsClassLike
	aspects_ col.ListLike[AspectLike]
}

// Attributes

func (v *aspects_) GetClass() AspectsClassLike {
	return v.class_
}

func (v *aspects_) GetAspects() col.ListLike[AspectLike] {
	return v.aspects_
}

// Public

// Private
