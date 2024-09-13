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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var optionClass = &optionClass_{
	// Initialize class constants.
}

// Function

func Option() OptionClassLike {
	return optionClass
}

// CLASS METHODS

// Target

type optionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *optionClass_) Make(repetitions abs.Sequential[RepetitionLike]) OptionLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(repetitions):
		panic("The repetitions attribute is required by this class.")
	default:
		return &option_{
			// Initialize instance attributes.
			class_:       c,
			repetitions_: repetitions,
		}
	}
}

// INSTANCE METHODS

// Target

type option_ struct {
	// Define instance attributes.
	class_       OptionClassLike
	repetitions_ abs.Sequential[RepetitionLike]
}

// Attributes

func (v *option_) GetClass() OptionClassLike {
	return v.class_
}

func (v *option_) GetRepetitions() abs.Sequential[RepetitionLike] {
	return v.repetitions_
}

// Private
