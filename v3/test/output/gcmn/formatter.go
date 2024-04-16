
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
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// This class does not initialize any private class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// This class does not define any private class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	depth_  int
	result_ sts.Builder
}

// Public

func (v *formatter_) FormatModel(model ModelLike) string {
	v.formatModel(model)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var separator = "\n"
	var indentation = "\t"
	for level := 0; level < v.depth_; level++ {
		separator += indentation
	}
	v.appendString(separator)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) formatModel(model ModelLike) {
	// TBA - Add real method implementation.
	v.depth_++
	v.appendString("test")
	v.appendNewline()
	v.depth_--
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
