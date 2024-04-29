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
DefaultClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete default-like class.
*/
type DefaultClassLike interface {
	// Constructors
	Make() DefaultLike
}

/*
ExampleClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete example-like class.
*/
type ExampleClassLike interface {
	// Constructors
	MakeWithDefault(default_ DefaultLike) ExampleLike
	MakeWithPrimitive(primitive PrimitiveLike) ExampleLike
	MakeWithLists(lists col.ListLike[ListLike]) ExampleLike
}

/*
ListClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete list-like class.
*/
type ListClassLike interface {
	// Constructors
	MakeWithExamples(examples col.ListLike[ExampleLike]) ListLike
}

/*
PrimitiveClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete primitive-like class.
*/
type PrimitiveClassLike interface {
	// Constructors
	MakeWithCharacter(character string) PrimitiveLike
	MakeWithText(text string) PrimitiveLike
	MakeWithInteger(integer string) PrimitiveLike
	MakeWithAnything(anything string) PrimitiveLike
}

// Instances

/*
DefaultLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete default-like class.
*/
type DefaultLike interface {
	// Attributes
	GetClass() DefaultClassLike
}

/*
ExampleLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete example-like class.
*/
type ExampleLike interface {
	// Attributes
	GetClass() ExampleClassLike
	GetDefault() DefaultLike
	GetPrimitive() PrimitiveLike
	GetLists() col.ListLike[ListLike]
}

/*
ListLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete list-like class.
*/
type ListLike interface {
	// Attributes
	GetClass() ListClassLike
	GetExamples() col.ListLike[ExampleLike]
}

/*
PrimitiveLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete primitive-like class.
*/
type PrimitiveLike interface {
	// Attributes
	GetClass() PrimitiveClassLike
	GetCharacter() string
	GetText() string
	GetInteger() string
	GetAnything() string
}
