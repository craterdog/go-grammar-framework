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

var characterClass = &characterClass_{
	// Initialize class constants.
}

// Function

func Character() CharacterClassLike {
	return characterClass
}

// CLASS METHODS

// Target

type characterClass_ struct {
	// Define class constants.
}

// Constructors

func (c *characterClass_) MakeWithBounded(bounded BoundedLike) CharacterLike {
	return &character_{
		// Initialize instance attributes.
		class_: c,
		any_:   bounded,
	}
}

func (c *characterClass_) MakeWithIntrinsic(intrinsic string) CharacterLike {
	return &character_{
		// Initialize instance attributes.
		class_: c,
		any_:   intrinsic,
	}
}

// INSTANCE METHODS

// Target

type character_ struct {
	// Define instance attributes.
	class_ CharacterClassLike
	any_   any
}

// Attributes

func (v *character_) GetClass() CharacterClassLike {
	return v.class_
}

func (v *character_) GetAny() any {
	return v.any_
}

// Private
