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

package generator

import (
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
	mod "github.com/craterdog/go-model-framework/v4"
)

// CLASS ACCESS

// Reference

var astClass = &astClass_{
	// Initialize the class constants.
}

// Function

func Ast() AstClassLike {
	return astClass
}

// CLASS METHODS

// Target

type astClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *astClass_) Make() AstLike {
	var ast = &ast_{
		// Initialize the instance attributes.
		class_:    c,
		analyzer_: gra.Analyzer().Make(),
	}
	return ast
}

// INSTANCE METHODS

// Target

type ast_ struct {
	// Define the instance attributes.
	class_    AstClassLike
	analyzer_ gra.AnalyzerLike
	modules_  abs.CatalogLike[string, string]
}

// Attributes

func (v *ast_) GetClass() AstClassLike {
	return v.class_
}

// Public

func (v *ast_) GenerateAstModel(
	module string,
	wiki string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	v.modules_ = col.Catalog[string, string]()
	var notice = mod.Notice(v.analyzer_.GetNotice() + "\n")
	var header = v.generateHeader(wiki)
	var classes = v.generateClasses()
	var instances = v.generateInstances()
	var imports = v.generateImports(module) // This must be last.
	var model = mod.Model(
		notice,
		header,
		imports,
		classes,
		instances,
	)
	implementation = mod.Formatter().FormatModel(model)
	return implementation
}

// Private

func (v *ast_) generateAttribute(reference ast.ReferenceLike) mod.AttributeLike {
	var identifier = reference.GetIdentifier().GetAny().(string)

	// Determine the attribute type.
	var attributeType mod.AbstractionLike
	switch {
	case gra.Scanner().MatchesType(identifier, gra.LowercaseToken):
		attributeType = mod.Abstraction("string")
	case gra.Scanner().MatchesType(identifier, gra.UppercaseToken):
		attributeType = mod.Abstraction(makeUpperCase(identifier) + "Like")
	}

	// Determine the attribute name.
	var attributeName = makeLowerCase(identifier)
	var cardinality = reference.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case ast.ConstrainedLike:
			var constrained = actual.GetAny().(string)
			switch constrained {
			case "?":
				attributeName = makeOptional(attributeName)
			case "*", "+":
				attributeName = makePlural(attributeName)
				attributeType = v.pluralizeType(attributeType)
			}
		case ast.QuantifiedLike:
			attributeName = makePlural(attributeName)
			attributeType = v.pluralizeType(attributeType)
		}
	}
	attributeName = "Get" + makeUpperCase(attributeName)

	// Define the attribute.
	var attribute = mod.Attribute(
		attributeName,
		attributeType,
	)
	return attribute
}

func (v *ast_) generateClassDeclaration(name string) mod.DeclarationLike {
	var comment = replaceAll(classCommentTemplate_, "className", name)
	var declaration = mod.Declaration(
		comment,
		makeUpperCase(name)+"ClassLike",
	)
	return declaration
}

func (v *ast_) generateClasses() mod.ClassesLike {
	var classes = col.List[mod.ClassLike]()
	var rules = v.analyzer_.GetRuleNames().GetIterator()
	for rules.HasNext() {
		var rule = rules.GetNext()
		var declaration = v.generateClassDeclaration(rule)
		var constructor mod.ConstructorLike
		if col.IsDefined(v.analyzer_.GetIdentifiers(rule)) {
			constructor = v.generateMultilineConstructor(rule)
		} else {
			constructor = v.generateInlineConstructor(rule)
		}
		var list = col.List[mod.ConstructorLike]()
		list.AppendValue(constructor)
		var constructors = mod.Constructors(list)
		var class = mod.Class(declaration, constructors)
		classes.AppendValue(class)
	}
	return mod.Classes(classes)
}

func (v *ast_) generateHeader(wiki string) mod.HeaderLike {
	var comment = replaceAll(packageHeaderTemplate_, "wiki", wiki)
	var header = mod.Header(comment, "ast")
	return header
}

func (v *ast_) generateImports(module string) mod.ImportsLike {
	var imports mod.ImportsLike
	if v.modules_.IsEmpty() {
		// There are no modules that are imported.
		return imports
	}
	v.modules_.SortValues() // Modules are sorted by path, not by alias.
	var list = col.List[mod.ModuleLike]()
	var iterator = v.modules_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var alias = association.GetValue()
		var path = association.GetKey()
		var module = mod.Module(alias, path)
		list.AppendValue(module)
	}
	var modules = mod.Modules(list)
	imports = mod.Imports(modules)
	return imports
}

func (v *ast_) generateInlineAttributes(name string) mod.AttributesLike {
	// Define the first attribute.
	var uppercase = makeUpperCase(name)
	var abstraction = mod.Abstraction(uppercase + "ClassLike")
	var attribute = mod.Attribute(
		"GetClass",
		abstraction,
	)
	var attributes = col.List[mod.AttributeLike]()
	attributes.AppendValue(attribute)

	// Define any additional attributes.
	var references = v.analyzer_.GetReferences(name).GetIterator()
	for references.HasNext() {
		var reference = references.GetNext()
		attribute = v.generateAttribute(reference)
		attributes.AppendValue(attribute)
	}

	return mod.Attributes(attributes)
}

func (v *ast_) generateInlineConstructor(name string) mod.ConstructorLike {
	// Define the parameters.
	var parameters = v.generateInlineParameters(name)

	// Define the return type.
	var uppercase = makeUpperCase(name)
	var abstraction = mod.Abstraction(uppercase + "Like")

	var constructor = mod.Constructor(
		"Make",
		parameters,
		abstraction,
	)
	return constructor
}

func (v *ast_) generateInlineParameters(name string) mod.ParametersLike {
	// Define the first parameter.
	var references = v.analyzer_.GetReferences(name).GetIterator()
	var reference = references.GetNext()
	var parameter = v.generateParameter(reference)

	// Define any additional parameters.
	var additionalParameters = col.List[mod.AdditionalParameterLike]()
	for references.HasNext() {
		var reference = references.GetNext()
		var additional = v.generateParameter(reference)
		additionalParameters.AppendValue(mod.AdditionalParameter(additional))
	}
	return mod.Parameters(parameter, additionalParameters)
}

func (v *ast_) generateInstanceDeclaration(name string) mod.DeclarationLike {
	var comment = replaceAll(instanceCommentTemplate_, "className", name)
	var declaration = mod.Declaration(
		comment,
		makeUpperCase(name)+"Like",
	)
	return declaration
}

func (v *ast_) generateInstances() mod.InstancesLike {
	var instances = col.List[mod.InstanceLike]()
	var rules = v.analyzer_.GetRuleNames().GetIterator()
	for rules.HasNext() {
		var rule = rules.GetNext()
		var declaration = v.generateInstanceDeclaration(rule)
		var attributes mod.AttributesLike
		if col.IsDefined(v.analyzer_.GetIdentifiers(rule)) {
			attributes = v.generateMultilineAttributes(rule)
		} else {
			attributes = v.generateInlineAttributes(rule)
		}
		var instance = mod.Instance(declaration, attributes)
		instances.AppendValue(instance)
	}
	return mod.Instances(instances)
}

func (v *ast_) generateMultilineAttributes(name string) mod.AttributesLike {
	// Define the first attribute.
	var uppercase = makeUpperCase(name)
	var abstraction = mod.Abstraction(uppercase + "ClassLike")
	var attribute = mod.Attribute(
		"GetClass",
		abstraction,
	)
	var attributes = col.List[mod.AttributeLike]()
	attributes.AppendValue(attribute)

	// Define the second attribute.
	abstraction = mod.Abstraction("any")
	attribute = mod.Attribute(
		"GetAny",
		abstraction,
	)
	attributes.AppendValue(attribute)

	return mod.Attributes(attributes)
}

func (v *ast_) generateMultilineConstructor(name string) mod.ConstructorLike {
	// Create the parameter for the constructor.
	var parameter = mod.Parameter(
		"any_",
		mod.Abstraction("any"),
	)
	var additionalParameters = col.List[mod.AdditionalParameterLike]()
	var parameters = mod.Parameters(
		parameter,
		additionalParameters.(abs.Sequential[mod.AdditionalParameterLike]),
	)

	// Create the return type.
	var uppercase = makeUpperCase(name)
	var abstraction = mod.Abstraction(uppercase + "Like")

	// Create the constructor.
	var constructor = mod.Constructor(
		"Make",
		parameters,
		abstraction,
	)
	return constructor
}

func (v *ast_) generateParameter(reference ast.ReferenceLike) mod.ParameterLike {
	var identifier = reference.GetIdentifier().GetAny().(string)

	// Determine the parameter type.
	var parameterType mod.AbstractionLike
	switch {
	case gra.Scanner().MatchesType(identifier, gra.LowercaseToken):
		parameterType = mod.Abstraction("string")
	case gra.Scanner().MatchesType(identifier, gra.UppercaseToken):
		parameterType = mod.Abstraction(makeUpperCase(identifier) + "Like")
	}

	// Determine the parameter name.
	var parameterName = makeLowerCase(identifier)
	var cardinality = reference.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case ast.ConstrainedLike:
			var constrained = actual.GetAny().(string)
			switch constrained {
			case "?":
				parameterName = makeOptional(parameterName)
			case "*", "+":
				parameterName = makePlural(parameterName)
				parameterType = v.pluralizeType(parameterType)
			}
		case ast.QuantifiedLike:
			parameterName = makePlural(parameterName)
			parameterType = v.pluralizeType(parameterType)
		}
	}

	// Define the parameter.
	var parameter = mod.Parameter(
		parameterName,
		parameterType,
	)
	return parameter
}

func (v *ast_) pluralizeType(abstraction mod.AbstractionLike) mod.AbstractionLike {
	// Add the collections module to the catalog of imported modules.
	var alias = "abs"
	var path = `"github.com/craterdog/go-collection-framework/v4/collection"`
	v.modules_.SetValue(path, alias) // Modules are sorted by path.

	// Create the generic arguments list for the pluralized abstraction.
	var argument = mod.Argument(abstraction)
	var additionalArguments = col.List[mod.AdditionalArgumentLike]()
	var arguments = mod.Arguments(argument, additionalArguments)
	var genericArguments = mod.GenericArguments(arguments)

	// Create the result type for the pluralized abstraction.
	abstraction = mod.Abstraction(
		mod.Alias(alias),
		"Sequential",
		genericArguments,
	)
	return abstraction
}

const classCommentTemplate_ = `/*
<ClassName>ClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete <class-name>-like class.
*/
`

const instanceCommentTemplate_ = `/*
<ClassName>Like is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete <class-name>-like class.
*/
`

const packageHeaderTemplate_ = `/*
Package "ast" provides the abstract syntax tree (AST) classes for this module.
Each AST class manages the attributes associated with the rule definition found
in the syntax grammar with the same rule name as the class.

For detailed documentation on this package refer to the wiki:
  - https://<wiki>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
`
