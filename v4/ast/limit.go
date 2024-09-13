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

// CLASS ACCESS

// Reference

var limitClass = &limitClass_{
	// Initialize class constants.
}

// Function

func Limit() LimitClassLike {
	return limitClass
}

// CLASS METHODS

// Target

type limitClass_ struct {
	// Define class constants.
}

// Constructors

func (c *limitClass_) Make(optionalNumber string) LimitLike {
	// Validate the arguments.
	switch {
	default:
		return &limit_{
			// Initialize instance attributes.
			class_:          c,
			optionalNumber_: optionalNumber,
		}
	}
}

// INSTANCE METHODS

// Target

type limit_ struct {
	// Define instance attributes.
	class_          LimitClassLike
	optionalNumber_ string
}

// Attributes

func (v *limit_) GetClass() LimitClassLike {
	return v.class_
}

func (v *limit_) GetOptionalNumber() string {
	return v.optionalNumber_
}

// Private
