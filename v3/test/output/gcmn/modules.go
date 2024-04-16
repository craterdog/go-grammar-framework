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

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var modulesClass = &modulesClass_{
	// This class has no private constants to initialize.
}

// Function

func Modules() ModulesClassLike {
	return modulesClass
}

// CLASS METHODS

// Target

type modulesClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *modulesClass_) MakeWithModules(modules col.ListLike[ModuleLike]) ModulesLike {
	return &modules_{
		modules_: modules,
	}
}

// Functions

// INSTANCE METHODS

// Target

type modules_ struct {
	modules_ col.ListLike[ModuleLike]
}

// Attributes

func (v *modules_) GetModules() col.ListLike[ModuleLike] {
	return v.modules_
}

// Public

// Private
