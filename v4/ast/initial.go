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

var initialClass = &initialClass_{
	// Initialize class constants.
}

// Function

func Initial() InitialClassLike {
	return initialClass
}

// CLASS METHODS

// Target

type initialClass_ struct {
	// Define class constants.
}

// Constructors

func (c *initialClass_) MakeWithRune(rune_ string) InitialLike {
	return &initial_{
		// Initialize instance attributes.
		class_: c,
		rune_:  rune_,
	}
}

// INSTANCE METHODS

// Target

type initial_ struct {
	// Define instance attributes.
	class_ InitialClassLike
	rune_  string
}

// Attributes

func (v *initial_) GetClass() InitialClassLike {
	return v.class_
}

func (v *initial_) GetRune() string {
	return v.rune_
}

// Private
