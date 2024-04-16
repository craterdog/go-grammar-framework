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

package cdcn

import ()

// CLASS ACCESS

// Reference

var associationClass = &associationClass_{
	// This class has no private constants to initialize.
}

// Function

func Association() AssociationClassLike {
	return associationClass
}

// CLASS METHODS

// Target

type associationClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *associationClass_) MakeWithAttributes(
	key KeyLike,
	value ValueLike,
) AssociationLike {
	return &association_{
		key_: key,
		value_: value,
	}
}

// Functions

// INSTANCE METHODS

// Target

type association_ struct {
	key_ KeyLike
	value_ ValueLike
}

// Attributes

func (v *association_) GetKey() KeyLike {
	return v.key_
}

func (v *association_) GetValue() ValueLike {
	return v.value_
}

// Public

// Private
