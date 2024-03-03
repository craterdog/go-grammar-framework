/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammar

import (
	fmt "fmt"
)

// CLASS ACCESS

// Reference

var elementClass = &elementClass_{
	// This class does not initialize any constants.
}

// Function

func Element() ElementClassLike {
	return elementClass
}

// CLASS METHODS

// Target

type elementClass_ struct {
	// This class does not define any constants.
}

// Constructors

func (c *elementClass_) MakeFromIntrinsic(
	intrinsic string,
) ElementLike {
	var element = &element_{
		// This class does not initialize any attributes.
	}
	element.SetIntrinsic(intrinsic)
	return element
}

func (c *elementClass_) MakeFromLiteral(
	literal string,
) ElementLike {
	var element = &element_{
		// This class does not initialize any attributes.
	}
	element.SetLiteral(literal)
	return element
}

func (c *elementClass_) MakeFromName(
	name string,
) ElementLike {
	var element = &element_{
		// This class does not initialize any attributes.
	}
	element.SetName(name)
	return element
}

// INSTANCE METHODS

// Target

type element_ struct {
	intrinsic string
	literal   string
	name      string
}

// Public

func (v *element_) GetIntrinsic() string {
	return v.intrinsic
}

func (v *element_) GetLiteral() string {
	return v.literal
}

func (v *element_) GetName() string {
	return v.name
}

func (v *element_) SetIntrinsic(intrinsic string) {
	if len(intrinsic) < 1 {
		var message = fmt.Sprintf("An invalid intrinsic was found:\n    %v\n", intrinsic)
		panic(message)
	}
	v.intrinsic = intrinsic
}

func (v *element_) SetLiteral(literal string) {
	if len(literal) < 1 {
		var message = fmt.Sprintf("An invalid literal was found:\n    %v\n", literal)
		panic(message)
	}
	v.literal = literal
}

func (v *element_) SetName(name string) {
	if len(name) < 1 {
		var message = fmt.Sprintf("An invalid name was found:\n    %v\n", name)
		panic(message)
	}
	v.name = name
}
