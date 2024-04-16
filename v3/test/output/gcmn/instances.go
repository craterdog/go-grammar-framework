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
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var instancesClass = &instancesClass_{
	// This class has no private constants to initialize.
}

// Function

func Instances() InstancesClassLike {
	return instancesClass
}

// CLASS METHODS

// Target

type instancesClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *instancesClass_) MakeWithInstances(instances col.ListLike[InstanceLike]) InstancesLike {
	return &instances_{
		instances_: instances,
	}
}

// Functions

// INSTANCE METHODS

// Target

type instances_ struct {
	instances_ col.ListLike[InstanceLike]
}

// Attributes

func (v *instances_) GetInstances() col.ListLike[InstanceLike] {
	return v.instances_
}

// Public

// Private
