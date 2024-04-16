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
Package "gcmn" provides...

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
package gcmn

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	CommentToken
	DelimiterToken
	EOFToken
	EOLToken
	IdentifierToken
	SpaceToken
	TextToken
)

// Classes

/*
AbstractionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstraction-like class.
*/
type AbstractionClassLike interface {
	// Constructors
	MakeWithAttributes(
		prefix PrefixLike,
		identifier string,
		arguments ArgumentsLike,
	) AbstractionLike
}

/*
AbstractionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstractions-like class.
*/
type AbstractionsClassLike interface {
	// Constructors
	MakeWithAbstractions(abstractions col.ListLike[AbstractionLike]) AbstractionsLike
}

/*
ArgumentsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete arguments-like class.
*/
type ArgumentsClassLike interface {
	// Constructors
	MakeWithAbstractions(abstractions col.ListLike[AbstractionLike]) ArgumentsLike
}

/*
AspectClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-like class.
*/
type AspectClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		methods MethodsLike,
	) AspectLike
}

/*
AspectsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspects-like class.
*/
type AspectsClassLike interface {
	// Constructors
	MakeWithAspects(aspects col.ListLike[AspectLike]) AspectsLike
}

/*
AttributeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attribute-like class.
*/
type AttributeClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameter ParameterLike,
		abstraction AbstractionLike,
	) AttributeLike
}

/*
AttributesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attributes-like class.
*/
type AttributesClassLike interface {
	// Constructors
	MakeWithAttributes(attributes col.ListLike[AttributeLike]) AttributesLike
}

/*
ClassClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-like class.
*/
type ClassClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		constants ConstantsLike,
		constructors ConstructorsLike,
		functions FunctionsLike,
	) ClassLike
}

/*
ClassesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete classes-like class.
*/
type ClassesClassLike interface {
	// Constructors
	MakeWithClasses(classes col.ListLike[ClassLike]) ClassesLike
}

/*
ConstantClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constant-like class.
*/
type ConstantClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		abstraction AbstractionLike,
	) ConstantLike
}

/*
ConstantsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constants-like class.
*/
type ConstantsClassLike interface {
	// Constructors
	MakeWithConstants(constants col.ListLike[ConstantLike]) ConstantsLike
}

/*
ConstructorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructor-like class.
*/
type ConstructorClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters ParametersLike,
		abstraction AbstractionLike,
	) ConstructorLike
}

/*
ConstructorsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructors-like class.
*/
type ConstructorsClassLike interface {
	// Constructors
	MakeWithConstructors(constructors col.ListLike[ConstructorLike]) ConstructorsLike
}

/*
DeclarationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete declaration-like class.
*/
type DeclarationClassLike interface {
	// Constructors
	MakeWithAttributes(
		comment string,
		identifier string,
		parameters ParametersLike,
	) DeclarationLike
}

/*
EnumerationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete enumeration-like class.
*/
type EnumerationClassLike interface {
	// Constructors
	MakeWithAttributes(
		parameter ParameterLike,
		identifier string,
	) EnumerationLike
}

/*
FormatterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructors
	Make() FormatterLike
}

/*
FunctionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete function-like class.
*/
type FunctionClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters ParametersLike,
		result ResultLike,
	) FunctionLike
}

/*
FunctionalClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functional-like class.
*/
type FunctionalClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		parameters ParametersLike,
		result ResultLike,
	) FunctionalLike
}

/*
FunctionalsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functionals-like class.
*/
type FunctionalsClassLike interface {
	// Constructors
	MakeWithFunctionals(functionals col.ListLike[FunctionalLike]) FunctionalsLike
}

/*
FunctionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functions-like class.
*/
type FunctionsClassLike interface {
	// Constructors
	MakeWithFunctions(functions col.ListLike[FunctionLike]) FunctionsLike
}

/*
HeaderClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete header-like class.
*/
type HeaderClassLike interface {
	// Constructors
	MakeWithAttributes(
		comment string,
		identifier string,
	) HeaderLike
}

/*
InstanceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-like class.
*/
type InstanceClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		attributes AttributesLike,
		abstractions AbstractionsLike,
		methods MethodsLike,
	) InstanceLike
}

/*
InstancesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instances-like class.
*/
type InstancesClassLike interface {
	// Constructors
	MakeWithInstances(instances col.ListLike[InstanceLike]) InstancesLike
}

/*
MethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete method-like class.
*/
type MethodClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters ParametersLike,
		result ResultLike,
	) MethodLike
}

/*
MethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete methods-like class.
*/
type MethodsClassLike interface {
	// Constructors
	MakeWithMethods(methods col.ListLike[MethodLike]) MethodsLike
}

/*
ModelClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete model-like class.
*/
type ModelClassLike interface {
	// Constructors
	MakeWithAttributes(
		notice NoticeLike,
		header HeaderLike,
		modules ModulesLike,
		types TypesLike,
		functionals FunctionalsLike,
		aspects AspectsLike,
		classes ClassesLike,
		instances InstancesLike,
	) ModelLike
}

/*
ModuleClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete module-like class.
*/
type ModuleClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		text string,
	) ModuleLike
}

/*
ModulesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete modules-like class.
*/
type ModulesClassLike interface {
	// Constructors
	MakeWithModules(modules col.ListLike[ModuleLike]) ModulesLike
}

/*
NoticeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete notice-like class.
*/
type NoticeClassLike interface {
	// Constructors
	MakeWithComment(comment string) NoticeLike
}

/*
ParameterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameter-like class.
*/
type ParameterClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		abstraction AbstractionLike,
	) ParameterLike
}

/*
ParametersClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameters-like class.
*/
type ParametersClassLike interface {
	// Constructors
	MakeWithParameters(parameters col.ListLike[ParameterLike]) ParametersLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructors
	Make() ParserLike
}

/*
PrefixClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete prefix-like class.
*/
type PrefixClassLike interface {
	// Constructors
	MakeWithIdentifier(identifier string) PrefixLike
}

/*
ResultClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete result-like class.
*/
type ResultClassLike interface {
	// Constructors
	MakeWithAbstraction(abstraction AbstractionLike) ResultLike
	MakeWithParameters(parameters ParametersLike) ResultLike
}

/*
ScannerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

MatchToken() a list of strings representing any matches found in the specified
text of the specified token type using the regular expression defined for that
token type.  If the regular expression contains submatch patterns the matching
substrings are returned as additional values in the list.
*/
type ScannerClassLike interface {
	// Constructors
	Make(
		source string,
		tokens col.QueueLike[TokenLike],
	) ScannerLike

	// Functions
	FormatToken(token TokenLike) string
	MatchToken(
		type_ TokenType,
		text string,
	) col.ListLike[string]
}

/*
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructors
	MakeWithAttributes(
		line int,
		position int,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
TypeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete type-like class.
*/
type TypeClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		abstraction AbstractionLike,
		enumeration EnumerationLike,
	) TypeLike
}

/*
TypesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete types-like class.
*/
type TypesClassLike interface {
	// Constructors
	MakeWithTypes(types col.ListLike[TypeLike]) TypesLike
}

/*
ValidatorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructors
	Make() ValidatorLike
}

// Instances

/*
AbstractionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete abstraction-like class.
*/
type AbstractionLike interface {
	// Attributes
	GetPrefix() PrefixLike
	GetIdentifier() string
	GetArguments() ArgumentsLike
}

/*
AbstractionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete abstractions-like class.
*/
type AbstractionsLike interface {
	// Attributes
	GetAbstractions() col.ListLike[AbstractionLike]
}

/*
ArgumentsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete arguments-like class.
*/
type ArgumentsLike interface {
	// Attributes
	GetAbstractions() col.ListLike[AbstractionLike]
}

/*
AspectLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-like class.
*/
type AspectLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetMethods() MethodsLike
}

/*
AspectsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspects-like class.
*/
type AspectsLike interface {
	// Attributes
	GetAspects() col.ListLike[AspectLike]
}

/*
AttributeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attribute-like class.
*/
type AttributeLike interface {
	// Attributes
	GetIdentifier() string
	GetParameter() ParameterLike
	GetAbstraction() AbstractionLike
}

/*
AttributesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attributes-like class.
*/
type AttributesLike interface {
	// Attributes
	GetAttributes() col.ListLike[AttributeLike]
}

/*
ClassLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete class-like class.
*/
type ClassLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetConstants() ConstantsLike
	GetConstructors() ConstructorsLike
	GetFunctions() FunctionsLike
}

/*
ClassesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete classes-like class.
*/
type ClassesLike interface {
	// Attributes
	GetClasses() col.ListLike[ClassLike]
}

/*
ConstantLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constant-like class.
*/
type ConstantLike interface {
	// Attributes
	GetIdentifier() string
	GetAbstraction() AbstractionLike
}

/*
ConstantsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constants-like class.
*/
type ConstantsLike interface {
	// Attributes
	GetConstants() col.ListLike[ConstantLike]
}

/*
ConstructorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructor-like class.
*/
type ConstructorLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() ParametersLike
	GetAbstraction() AbstractionLike
}

/*
ConstructorsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructors-like class.
*/
type ConstructorsLike interface {
	// Attributes
	GetConstructors() col.ListLike[ConstructorLike]
}

/*
DeclarationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete declaration-like class.
*/
type DeclarationLike interface {
	// Attributes
	GetComment() string
	GetIdentifier() string
	GetParameters() ParametersLike
}

/*
EnumerationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete enumeration-like class.
*/
type EnumerationLike interface {
	// Attributes
	GetParameter() ParameterLike
	GetIdentifier() string
}

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Methods
	FormatModel(model ModelLike) string
}

/*
FunctionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete function-like class.
*/
type FunctionLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
FunctionalLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functional-like class.
*/
type FunctionalLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
FunctionalsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functionals-like class.
*/
type FunctionalsLike interface {
	// Attributes
	GetFunctionals() col.ListLike[FunctionalLike]
}

/*
FunctionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functions-like class.
*/
type FunctionsLike interface {
	// Attributes
	GetFunctions() col.ListLike[FunctionLike]
}

/*
HeaderLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete header-like class.
*/
type HeaderLike interface {
	// Attributes
	GetComment() string
	GetIdentifier() string
}

/*
InstanceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instance-like class.
*/
type InstanceLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetAttributes() AttributesLike
	GetAbstractions() AbstractionsLike
	GetMethods() MethodsLike
}

/*
InstancesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instances-like class.
*/
type InstancesLike interface {
	// Attributes
	GetInstances() col.ListLike[InstanceLike]
}

/*
MethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete method-like class.
*/
type MethodLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
MethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete methods-like class.
*/
type MethodsLike interface {
	// Attributes
	GetMethods() col.ListLike[MethodLike]
}

/*
ModelLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete model-like class.
*/
type ModelLike interface {
	// Attributes
	GetNotice() NoticeLike
	GetHeader() HeaderLike
	GetModules() ModulesLike
	GetTypes() TypesLike
	GetFunctionals() FunctionalsLike
	GetAspects() AspectsLike
	GetClasses() ClassesLike
	GetInstances() InstancesLike
}

/*
ModuleLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete module-like class.
*/
type ModuleLike interface {
	// Attributes
	GetIdentifier() string
	GetText() string
}

/*
ModulesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete modules-like class.
*/
type ModulesLike interface {
	// Attributes
	GetModules() col.ListLike[ModuleLike]
}

/*
NoticeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete notice-like class.
*/
type NoticeLike interface {
	// Attributes
	GetComment() string
}

/*
ParameterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Attributes
	GetIdentifier() string
	GetAbstraction() AbstractionLike
}

/*
ParametersLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameters-like class.
*/
type ParametersLike interface {
	// Attributes
	GetParameters() col.ListLike[ParameterLike]
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Methods
	ParseSource(source string) ModelLike
}

/*
PrefixLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete prefix-like class.
*/
type PrefixLike interface {
	// Attributes
	GetIdentifier() string
}

/*
ResultLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete result-like class.
*/
type ResultLike interface {
	// Attributes
	GetAbstraction() AbstractionLike
	GetParameters() ParametersLike
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Attributes
	GetLine() int
	GetPosition() int
	GetType() TokenType
	GetValue() string
}

/*
TypeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete type-like class.
*/
type TypeLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetAbstraction() AbstractionLike
	GetEnumeration() EnumerationLike
}

/*
TypesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete types-like class.
*/
type TypesLike interface {
	// Attributes
	GetTypes() col.ListLike[TypeLike]
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Methods
	ValidateModel(model ModelLike)
}
