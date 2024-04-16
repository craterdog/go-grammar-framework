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

var noticeClass = &noticeClass_{
	// This class has no private constants to initialize.
}

// Function

func Notice() NoticeClassLike {
	return noticeClass
}

// CLASS METHODS

// Target

type noticeClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *noticeClass_) MakeWithComment(comment string) NoticeLike {
	return &notice_{
		comment_: comment,
	}
}

// Functions

// INSTANCE METHODS

// Target

type notice_ struct {
	comment_ string
}

// Attributes

func (v *notice_) GetComment() string {
	return v.comment_
}

// Public

// Private
