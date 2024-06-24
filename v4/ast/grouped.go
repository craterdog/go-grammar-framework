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

var groupedClass = &groupedClass_{
	// Initialize class constants.
}

// Function

func Grouped() GroupedClassLike {
	return groupedClass
}

// CLASS METHODS

// Target

type groupedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *groupedClass_) Make(pattern PatternLike) GroupedLike {
	return &grouped_{
		// Initialize instance attributes.
		class_:   c,
		pattern_: pattern,
	}
}

// INSTANCE METHODS

// Target

type grouped_ struct {
	// Define instance attributes.
	class_   GroupedClassLike
	pattern_ PatternLike
}

// Attributes

func (v *grouped_) GetClass() GroupedClassLike {
	return v.class_
}

func (v *grouped_) GetPattern() PatternLike {
	return v.pattern_
}

// Private
