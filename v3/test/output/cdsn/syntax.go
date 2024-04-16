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

package cdsn

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var syntaxClass = &syntaxClass_{
	// This class has no private constants to initialize.
}

// Function

func Syntax() SyntaxClassLike {
	return syntaxClass
}

// CLASS METHODS

// Target

type syntaxClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *syntaxClass_) MakeWithAttributes(
	headers col.ListLike[HeaderLike],
	definitions col.ListLike[DefinitionLike],
) SyntaxLike {
	return &syntax_{
		headers_: headers,
		definitions_: definitions,
	}
}

// Functions

// INSTANCE METHODS

// Target

type syntax_ struct {
	headers_ col.ListLike[HeaderLike]
	definitions_ col.ListLike[DefinitionLike]
}

// Attributes

func (v *syntax_) GetHeaders() col.ListLike[HeaderLike] {
	return v.headers_
}

func (v *syntax_) GetDefinitions() col.ListLike[DefinitionLike] {
	return v.definitions_
}

// Public

// Private
