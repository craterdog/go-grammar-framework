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

package gcmn

import ()

// CLASS ACCESS

// Reference

var moduleClass = &moduleClass_{
	// This class has no private constants to initialize.
}

// Function

func Module() ModuleClassLike {
	return moduleClass
}

// CLASS METHODS

// Target

type moduleClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *moduleClass_) MakeWithAttributes(
	identifier string,
	text string,
) ModuleLike {
	return &module_{
		identifier_: identifier,
		text_: text,
	}
}

// Functions

// INSTANCE METHODS

// Target

type module_ struct {
	identifier_ string
	text_ string
}

// Attributes

func (v *module_) GetIdentifier() string {
	return v.identifier_
}

func (v *module_) GetText() string {
	return v.text_
}

// Public

// Private
