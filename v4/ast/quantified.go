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
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var quantifiedClass = &quantifiedClass_{
	// Initialize class constants.
}

// Function

func Quantified() QuantifiedClassLike {
	return quantifiedClass
}

// CLASS METHODS

// Target

type quantifiedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *quantifiedClass_) Make(
	delimiter string,
	number string,
	optionalLimit LimitLike,
	delimiter2 string,
) QuantifiedLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(delimiter):
		panic("The delimiter attribute is required by this class.")
	case col.IsUndefined(number):
		panic("The number attribute is required by this class.")
	case col.IsUndefined(delimiter2):
		panic("The delimiter2 attribute is required by this class.")
	default:
		return &quantified_{
			// Initialize instance attributes.
			class_:         c,
			delimiter_:     delimiter,
			number_:        number,
			optionalLimit_: optionalLimit,
			delimiter2_:    delimiter2,
		}
	}
}

// INSTANCE METHODS

// Target

type quantified_ struct {
	// Define instance attributes.
	class_         QuantifiedClassLike
	delimiter_     string
	number_        string
	optionalLimit_ LimitLike
	delimiter2_    string
}

// Attributes

func (v *quantified_) GetClass() QuantifiedClassLike {
	return v.class_
}

func (v *quantified_) GetDelimiter() string {
	return v.delimiter_
}

func (v *quantified_) GetNumber() string {
	return v.number_
}

func (v *quantified_) GetOptionalLimit() LimitLike {
	return v.optionalLimit_
}

func (v *quantified_) GetDelimiter2() string {
	return v.delimiter2_
}

// Private
