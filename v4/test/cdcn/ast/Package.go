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

/*
Package "ast" provides...

For detailed documentation on this package refer to the wiki:
  - <wiki>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// Classes

/*
AssociationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete association-like class.
*/
type AssociationClassLike interface {
	// Constructors
	MakeWithAttributes(
		key KeyLike,
		value ValueLike,
	) AssociationLike
}

/*
AssociationsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete associations-like class.
*/
type AssociationsClassLike interface {
	// Constructors
	MakeWithAssociations(associations col.ListLike[AssociationLike]) AssociationsLike
}

/*
CollectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete collection-like class.
*/
type CollectionClassLike interface {
	// Constructors
	MakeWithAttributes(
		associations AssociationsLike,
		values ValuesLike,
		context string,
	) CollectionLike
}

/*
KeyClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete key-like class.
*/
type KeyClassLike interface {
	// Constructors
	MakeWithPrimitive(primitive PrimitiveLike) KeyLike
}

/*
PrimitiveClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete primitive-like class.
*/
type PrimitiveClassLike interface {
	// Constructors
	MakeWithBoolean(boolean string) PrimitiveLike
	MakeWithComplex(complex_ string) PrimitiveLike
	MakeWithFloat(float string) PrimitiveLike
	MakeWithHexadecimal(hexadecimal string) PrimitiveLike
	MakeWithInteger(integer string) PrimitiveLike
	MakeWithNil(nil_ string) PrimitiveLike
	MakeWithRune(rune_ string) PrimitiveLike
	MakeWithString(string_ string) PrimitiveLike
}

/*
ValueClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete value-like class.
*/
type ValueClassLike interface {
	// Constructors
	MakeWithPrimitive(primitive PrimitiveLike) ValueLike
	MakeWithCollection(collection CollectionLike) ValueLike
}

/*
ValuesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete values-like class.
*/
type ValuesClassLike interface {
	// Constructors
	MakeWithValues(values col.ListLike[ValueLike]) ValuesLike
}

// Instances

/*
AssociationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete association-like class.
*/
type AssociationLike interface {
	// Attributes
	GetClass() AssociationClassLike
	GetKey() KeyLike
	GetValue() ValueLike
}

/*
AssociationsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete associations-like class.
*/
type AssociationsLike interface {
	// Attributes
	GetClass() AssociationsClassLike
	GetAssociations() col.ListLike[AssociationLike]
}

/*
CollectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete collection-like class.
*/
type CollectionLike interface {
	// Attributes
	GetClass() CollectionClassLike
	GetAssociations() AssociationsLike
	GetValues() ValuesLike
	GetContext() string
}

/*
KeyLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete key-like class.
*/
type KeyLike interface {
	// Attributes
	GetClass() KeyClassLike
	GetPrimitive() PrimitiveLike
}

/*
PrimitiveLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete primitive-like class.
*/
type PrimitiveLike interface {
	// Attributes
	GetClass() PrimitiveClassLike
	GetBoolean() string
	GetComplex() string
	GetFloat() string
	GetHexadecimal() string
	GetInteger() string
	GetNil() string
	GetRune() string
	GetString() string
}

/*
ValueLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete value-like class.
*/
type ValueLike interface {
	// Attributes
	GetClass() ValueClassLike
	GetPrimitive() PrimitiveLike
	GetCollection() CollectionLike
}

/*
ValuesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete values-like class.
*/
type ValuesLike interface {
	// Attributes
	GetClass() ValuesClassLike
	GetValues() col.ListLike[ValueLike]
}
