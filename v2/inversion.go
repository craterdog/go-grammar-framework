/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package grammars

import ()

// CLASS ACCESS

// Reference

var inversionClass = &inversionClass_{
	// TBA - Assign constant values.
}

// Function

func Inversion() InversionClassLike {
	return inversionClass
}

// CLASS METHODS

// Target

type inversionClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *inversionClass_) MakeWithAttributes(
	inverted bool,
	filter FilterLike,
) InversionLike {
	return &inversion_{
		inverted_: inverted,
		filter_:   filter,
	}
}

// Functions

// INSTANCE METHODS

// Target

type inversion_ struct {
	inverted_ bool
	filter_   FilterLike
}

// Attributes

func (v *inversion_) IsInverted() bool {
	return v.inverted_
}

func (v *inversion_) GetFilter() FilterLike {
	return v.filter_
}

// Public

// Private
