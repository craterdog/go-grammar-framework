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

var noticeClass = &noticeClass_{
	// Initialize class constants.
}

// Function

func Notice() NoticeClassLike {
	return noticeClass
}

// CLASS METHODS

// Target

type noticeClass_ struct {
	// Define class constants.
}

// Constructors

func (c *noticeClass_) Make(
	comment string,
	newline string,
) NoticeLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(comment):
		panic("The comment attribute is required by this class.")
	case col.IsUndefined(newline):
		panic("The newline attribute is required by this class.")
	default:
		return &notice_{
			// Initialize instance attributes.
			class_:   c,
			comment_: comment,
			newline_: newline,
		}
	}
}

// INSTANCE METHODS

// Target

type notice_ struct {
	// Define instance attributes.
	class_   NoticeClassLike
	comment_ string
	newline_ string
}

// Attributes

func (v *notice_) GetClass() NoticeClassLike {
	return v.class_
}

func (v *notice_) GetComment() string {
	return v.comment_
}

func (v *notice_) GetNewline() string {
	return v.newline_
}

// Private
