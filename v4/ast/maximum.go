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

var maximumClass = &maximumClass_{
	// Initialize class constants.
}

// Function

func Maximum() MaximumClassLike {
	return maximumClass
}

// CLASS METHODS

// Target

type maximumClass_ struct {
	// Define class constants.
}

// Constructors

func (c *maximumClass_) Make(number string) MaximumLike {
	return &maximum_{
		// Initialize instance attributes.
		class_:  c,
		number_: number,
	}
}

// INSTANCE METHODS

// Target

type maximum_ struct {
	// Define instance attributes.
	class_  MaximumClassLike
	number_ string
}

// Attributes

func (v *maximum_) GetClass() MaximumClassLike {
	return v.class_
}

func (v *maximum_) GetNumber() string {
	return v.number_
}

// Private
