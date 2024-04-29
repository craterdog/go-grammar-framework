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
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var modulesClass = &modulesClass_{
	// Any private class constants should be initialized here.
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
	class_ ModulesClassLike
	modules_ col.ListLike[ModuleLike]
}

// Attributes

func (v *modules_) GetClass() ModulesClassLike {
	return v.class_
}

func (v *modules_) GetModules() col.ListLike[ModuleLike] {
	return v.modules_
}

// Public

// Private
