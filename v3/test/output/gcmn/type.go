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

var typeClass = &typeClass_{
	// This class has no private constants to initialize.
}

// Function

func Type() TypeClassLike {
	return typeClass
}

// CLASS METHODS

// Target

type typeClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *typeClass_) MakeWithAttributes(
	declaration DeclarationLike,
	abstraction AbstractionLike,
	enumeration EnumerationLike,
) TypeLike {
	return &type_{
		declaration_: declaration,
		abstraction_: abstraction,
		enumeration_: enumeration,
	}
}

// Functions

// INSTANCE METHODS

// Target

type type_ struct {
	declaration_ DeclarationLike
	abstraction_ AbstractionLike
	enumeration_ EnumerationLike
}

// Attributes

func (v *type_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *type_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *type_) GetEnumeration() EnumerationLike {
	return v.enumeration_
}

// Public

// Private
