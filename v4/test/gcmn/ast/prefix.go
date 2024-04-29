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

var prefixClass = &prefixClass_{
	// Any private class constants should be initialized here.
}

// Function

func Prefix() PrefixClassLike {
	return prefixClass
}

// CLASS METHODS

// Target

type prefixClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *prefixClass_) MakeWithIdentifier(identifier string) PrefixLike {
	return &prefix_{
		identifier_: identifier,
	}
}

// Functions

// INSTANCE METHODS

// Target

type prefix_ struct {
	class_ PrefixClassLike
	identifier_ string
}

// Attributes

func (v *prefix_) GetClass() PrefixClassLike {
	return v.class_
}

func (v *prefix_) GetIdentifier() string {
	return v.identifier_
}

// Public

// Private
