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
	stc "strconv"
	sts "strings"
	uni "unicode"
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
	var processor = gra.Processor().Make()
	var ast = &ast_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: processor,
	}
	ast.visitor_ = gra.Visitor().Make(ast)
	return ast
}

// INSTANCE METHODS

// Target

type ast_ struct {
	// Define the instance attributes.
	class_      AstClassLike
	visitor_    gra.VisitorLike
	modules_    abs.CatalogLike[string, mod.ModuleLike]
	classes_    abs.CatalogLike[string, mod.ClassLike]
	instances_  abs.CatalogLike[string, mod.InstanceLike]
	attributes_ abs.ListLike[mod.AttributeLike]

	// Define the inherited aspects.
	gra.Methodical
}

// Attributes

func (v *ast_) GetClass() AstClassLike {
	return v.class_
}

// Methodical

func (v *ast_) PreprocessFactor(
	factor ast.FactorLike,
	index uint,
	size uint,
) {
	switch factor.GetAny().(type) {
	case string:
		var abstraction = mod.Abstraction("string")
		var attribute = mod.Attribute(
			"GetReserved",
			abstraction,
		)
		v.attributes_.AppendValue(attribute)
	}
}

func (v *ast_) PostprocessInlined(inlined ast.InlinedLike) {
	v.consolidateAttributes()
}

func (v *ast_) PostprocessMultilined(multilined ast.MultilinedLike) {
	var abstraction = mod.Abstraction("any")
	var attribute = mod.Attribute(
		"GetAny",
		abstraction,
	)
	v.attributes_ = col.List[mod.AttributeLike]()
	v.attributes_.AppendValue(attribute)
}

func (v *ast_) PreprocessPredicate(
	predicate ast.PredicateLike,
) {
	var identifier = predicate.GetIdentifier().GetAny().(string)
	var attributeName = v.makeUppercase(identifier)
	var attributeType mod.AbstractionLike
	switch {
	case gra.Scanner().MatchesType(identifier, gra.LowercaseToken):
		attributeType = mod.Abstraction("string")
	case gra.Scanner().MatchesType(identifier, gra.UppercaseToken):
		attributeType = mod.Abstraction(identifier + "Like")
	}
	var attribute = mod.Attribute(
		"Get"+attributeName,
		attributeType,
	)
	var cardinality = predicate.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case ast.ConstrainedLike:
			attribute = v.pluralizeAttribute(attribute)
		case string:
			switch actual {
			case "?":
				// This attribute is optional.
				attribute = v.optionalizeAttribute(attribute)
			case "*", "+":
				// Turn the attribute into a sequence of that type attribute.
				attribute = v.pluralizeAttribute(attribute)
			}
		}
	}
	v.attributes_.AppendValue(attribute)
}

func (v *ast_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	v.attributes_ = col.List[mod.AttributeLike]()
}

func (v *ast_) PostprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var name = rule.GetUppercase()

	// Generate the class constructor.
	var constructor = v.generateConstructor(name)

	// Generate the class interface.
	var class = v.generateClass(name, constructor)
	v.classes_.SetValue(name, class)

	// Generate the instance interface.
	var instance = v.generateInstance(name)
	v.instances_.SetValue(name, instance)
}

func (v *ast_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.modules_ = col.Catalog[string, mod.ModuleLike]()
	v.classes_ = col.Catalog[string, mod.ClassLike]()
	v.instances_ = col.Catalog[string, mod.InstanceLike]()
}

func (v *ast_) PostprocessSyntax(syntax ast.SyntaxLike) {
	v.modules_.SortValues()
	v.classes_.SortValues()
	v.instances_.SortValues()
}

func (v *ast_) GenerateAstModel(
	module string,
	wiki string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.visitor_.VisitSyntax(syntax)
	implementation = astTemplate_
	implementation = sts.ReplaceAll(implementation, "<wiki>", wiki)
	var name = v.extractSyntaxName(syntax)
	implementation = sts.ReplaceAll(implementation, "<module>", module)
	var notice = v.extractNotice(syntax)
	implementation = sts.ReplaceAll(implementation, "<Notice>", notice)
	var uppercase = v.makeUppercase(name)
	implementation = sts.ReplaceAll(implementation, "<Name>", uppercase)
	var lowercase = v.makeLowercase(name)
	implementation = sts.ReplaceAll(implementation, "<name>", lowercase)
	implementation = v.augmentModel(implementation)
	return implementation
}

// Private

func (v *ast_) augmentModel(implementation string) string {
	var model = mod.Parser().ParseSource(implementation)
	var notice = model.GetNotice()
	var header = model.GetHeader()
	var keys = v.modules_.GetKeys()
	var imports mod.ImportsLike
	if keys.GetSize() > 0 {
		var modules = mod.Modules(v.modules_.GetValues(keys))
		imports = mod.Imports(modules)
	}
	var types = model.GetOptionalTypes()
	var functionals = model.GetOptionalFunctionals()
	keys = v.classes_.GetKeys()
	var classes = mod.Classes(v.classes_.GetValues(keys))
	var instances = mod.Instances(v.instances_.GetValues(keys))
	var aspects = model.GetOptionalAspects()
	model = mod.Model(
		notice,
		header,
		imports,
		types,
		functionals,
		classes,
		instances,
		aspects,
	)
	implementation = mod.Formatter().FormatModel(model)
	return implementation
}

func (v *ast_) consolidateAttributes() {
	// Compare each attribute type and rename duplicates.
	for i := 1; i <= v.attributes_.GetSize(); i++ {
		var attribute = v.attributes_.GetValue(i)
		var first = attribute.GetName()
		for j := i + 1; j <= v.attributes_.GetSize(); j++ {
			var count = 1
			var second = v.attributes_.GetValue(j).GetName()
			if first == second {
				count++
				var attributeName = second + stc.Itoa(count)
				var attributeType = attribute.GetOptionalAbstraction()
				var newAttribute = mod.Attribute(attributeName, attributeType)
				v.attributes_.SetValue(j, newAttribute)
			}
		}
	}
}

func (v *ast_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

func (v *ast_) extractParameters() abs.Sequential[mod.ParameterLike] {
	var parameters = col.List[mod.ParameterLike]()
	var iterator = v.attributes_.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var name = sts.TrimPrefix(attribute.GetName(), "Get")
		var abstraction = attribute.GetOptionalAbstraction()
		if col.IsUndefined(abstraction) {
			var parameter = attribute.GetOptionalParameter()
			abstraction = parameter.GetAbstraction()
		}
		var parameter = mod.Parameter(
			v.makeLowercase(name),
			abstraction,
		)
		parameters.AppendValue(parameter)
	}
	return parameters
}

func (v *ast_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *ast_) generateClass(
	name string,
	constructor mod.ConstructorLike,
) mod.ClassLike {
	var comment = sts.ReplaceAll(classCommentTemplate_, "<Class>", name)
	comment = sts.ReplaceAll(comment, "<class>", sts.ToLower(name))
	var declaration = mod.Declaration(
		comment,
		name+"ClassLike",
	)
	var list = col.List[mod.ConstructorLike]()
	list.AppendValue(constructor)
	var constructors = mod.Constructors(list)
	var class = mod.Class(
		declaration,
		constructors,
	)
	return class
}

func (v *ast_) generateConstructor(name string) mod.ConstructorLike {
	var abstraction = mod.Abstraction(name + "Like")
	name = "Make"
	var parameters mod.ParametersLike
	var iterator = v.extractParameters().GetIterator()
	if iterator.HasNext() {
		var parameter = iterator.GetNext()
		var additionalParameters = col.List[mod.AdditionalParameterLike]()
		for iterator.HasNext() {
			var parameter = iterator.GetNext()
			var additionalParameter = mod.AdditionalParameter(parameter)
			additionalParameters.AppendValue(additionalParameter)
		}
		parameters = mod.Parameters(
			parameter,
			additionalParameters.(abs.Sequential[mod.AdditionalParameterLike]),
		)
	}
	var constructor = mod.Constructor(
		name,
		parameters,
		abstraction,
	)
	return constructor
}

func (v *ast_) generateInstance(
	name string,
) mod.InstanceLike {
	var comment = sts.ReplaceAll(instanceCommentTemplate_, "<Class>", name)
	comment = sts.ReplaceAll(comment, "<class>", sts.ToLower(name))
	var declaration = mod.Declaration(
		comment,
		name+"Like",
	)
	var abstraction = mod.Abstraction(
		name + "ClassLike",
	)
	var attribute = mod.Attribute(
		"GetClass",
		abstraction,
	)
	v.attributes_.InsertValue(0, attribute)
	var instance = mod.Instance(
		declaration,
		mod.Attributes(v.attributes_),
	)
	return instance
}

func (v *ast_) makeLowercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	name = string(runes)
	if reserved_[name] {
		name += "_"
	}
	return name
}

func (v *ast_) makeUppercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

func (v *ast_) optionalizeAttribute(
	attribute mod.AttributeLike,
) mod.AttributeLike {
	var name = attribute.GetName()
	name = "GetOptional" + sts.TrimPrefix(name, "Get")
	var attributeType = attribute.GetOptionalAbstraction()
	attribute = mod.Attribute(name, attributeType)
	return attribute
}

func (v *ast_) pluralizeAttribute(
	attribute mod.AttributeLike,
) mod.AttributeLike {
	// Add the collections module to the catalog of imported modules.
	var alias = "abs"
	var path = `"github.com/craterdog/go-collection-framework/v4/collection"`
	var module = mod.Module(
		alias,
		path,
	)
	v.modules_.SetValue(path, module)

	// Extract the name and attribute type from the attribute.
	var name = attribute.GetName()
	if sts.HasSuffix(name, "s") {
		name += "es"
	} else {
		name += "s"
	}
	var attributeType = attribute.GetOptionalAbstraction() // Not optional here.

	// Create the generic arguments list for the pluralized attribute.
	var argument = mod.Argument(attributeType)
	var additionalArguments = col.List[mod.AdditionalArgumentLike]()
	var arguments = mod.Arguments(argument, additionalArguments)
	var genericArguments = mod.GenericArguments(arguments)

	// Create the result type for the pluralized attribute.
	attributeType = mod.Abstraction(
		mod.Alias(alias),
		"Sequential",
		genericArguments,
	)
	attribute = mod.Attribute(name, attributeType)
	return attribute
}

const astTemplate_ = `/*<Notice>*/

/*
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
package ast

import (
	ast "github.com/craterdog/go-collection-framework/v4/collection"
)

// Classes

/*
This is a dummy class placeholder.
*/
type DummyClassLike interface {
	// Constructors
	Make() DummyLike
}

// Instances

/*
This is a dummy instance placeholder.
*/
type DummyLike interface {
	// Attributes
	GetClass() DummyClassLike
}
`

const classCommentTemplate_ = `/*
<Class>ClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete <class>-like class.
*/
`

const instanceCommentTemplate_ = `/*
<Class>Like is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete <class>-like class.
*/
`
