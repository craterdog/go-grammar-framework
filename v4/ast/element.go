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

import ()

// CLASS ACCESS

// Reference

var elementClass = &elementClass_{
	// Initialize class constants.
}

// Function

func Element() ElementClassLike {
	return elementClass
}

// CLASS METHODS

// Target

type elementClass_ struct {
	// Define class constants.
	// This class has no private constants.

}

// Constructors

func (c *elementClass_) MakeWithLiteral(literal string) ElementLike {
	return &element_{
		// Initialize instance attributes.
		class_:   c,
		literal_: literal,
	}
}

func (c *elementClass_) MakeWithName(name string) ElementLike {
	return &element_{
		// Initialize instance attributes.
		class_: c,
		name_:  name,
	}
}

// INSTANCE METHODS

// Target

type element_ struct {
	// Define instance attributes.
	class_   ElementClassLike
	literal_ string
	name_    string
}

// Attributes

func (v *element_) GetClass() ElementClassLike {
	return v.class_
}

func (v *element_) GetLiteral() string {
	return v.literal_
}

func (v *element_) GetName() string {
	return v.name_
}

// Private
