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
	mod "github.com/craterdog/go-collection-framework/v4"
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var syntaxClass = &syntaxClass_{
	// Initialize class constants.
}

// Function

func Syntax() SyntaxClassLike {
	return syntaxClass
}

// CLASS METHODS

// Target

type syntaxClass_ struct {
	// Define class constants.
}

// Constructors

func (c *syntaxClass_) Make(
	headers col.Sequential[HeaderLike],
	rules col.Sequential[RuleLike],
	lexigrams col.Sequential[LexigramLike],
) SyntaxLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(headers):
		panic("The headers attribute is required for each Syntax.")
	case mod.IsUndefined(rules):
		panic("The rules attribute is required for each Syntax.")
	case mod.IsUndefined(lexigrams):
		panic("The lexigrams attribute is required for each Syntax.")
	default:
		return &syntax_{
			// Initialize instance attributes.
			class_:     c,
			headers_:   headers,
			rules_:     rules,
			lexigrams_: lexigrams,
		}
	}
}

// INSTANCE METHODS

// Target

type syntax_ struct {
	// Define instance attributes.
	class_     SyntaxClassLike
	headers_   col.Sequential[HeaderLike]
	rules_     col.Sequential[RuleLike]
	lexigrams_ col.Sequential[LexigramLike]
}

// Attributes

func (v *syntax_) GetClass() SyntaxClassLike {
	return v.class_
}

func (v *syntax_) GetHeaders() col.Sequential[HeaderLike] {
	return v.headers_
}

func (v *syntax_) GetRules() col.Sequential[RuleLike] {
	return v.rules_
}

func (v *syntax_) GetLexigrams() col.Sequential[LexigramLike] {
	return v.lexigrams_
}

// Private
