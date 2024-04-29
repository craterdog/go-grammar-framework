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

var attributesClass = &attributesClass_{
	// Any private class constants should be initialized here.
}

// Function

func Attributes() AttributesClassLike {
	return attributesClass
}

// CLASS METHODS

// Target

type attributesClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *attributesClass_) MakeWithAttributes(attributes col.ListLike[AttributeLike]) AttributesLike {
	return &attributes_{
		attributes_: attributes,
	}
}

// Functions

// INSTANCE METHODS

// Target

type attributes_ struct {
	class_ AttributesClassLike
	attributes_ col.ListLike[AttributeLike]
}

// Attributes

func (v *attributes_) GetClass() AttributesClassLike {
	return v.class_
}

func (v *attributes_) GetAttributes() col.ListLike[AttributeLike] {
	return v.attributes_
}

// Public

// Private
