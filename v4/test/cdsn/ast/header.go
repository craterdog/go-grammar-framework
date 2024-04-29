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

var headerClass = &headerClass_{
	// Any private class constants should be initialized here.
}

// Function

func Header() HeaderClassLike {
	return headerClass
}

// CLASS METHODS

// Target

type headerClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *headerClass_) MakeWithComment(comment string) HeaderLike {
	return &header_{
		comment_: comment,
	}
}

// Functions

// INSTANCE METHODS

// Target

type header_ struct {
	class_ HeaderClassLike
	comment_ string
}

// Attributes

func (v *header_) GetClass() HeaderClassLike {
	return v.class_
}

func (v *header_) GetComment() string {
	return v.comment_
}

// Public

// Private
