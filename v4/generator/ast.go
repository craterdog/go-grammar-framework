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
		analyzer_: Analyzer().Make(),
		modules_:  col.Catalog[string, string](),
	}
	return ast
}

// INSTANCE METHODS

// Target

type ast_ struct {
	// Define the instance attributes.
	class_    *astClass_
	analyzer_ AnalyzerLike
	modules_  abs.CatalogLike[string, string]
}

// Public

func (v *ast_) GetClass() AstClassLike {
	return v.class_
}

func (v *ast_) GenerateAstModel(
	wiki string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = v.getTemplate(modelTemplate)
	var notice = v.generateNotice()
	implementation = replaceAll(implementation, "notice", notice)
	var header = v.generateHeader(wiki)
	implementation = replaceAll(implementation, "header", header)
	var imports = v.generateImports()
	implementation = replaceAll(implementation, "imports", imports)
	var classes = v.generateClasses()
	implementation = replaceAll(implementation, "classes", classes)
	var instances = v.generateInstances()
	implementation = replaceAll(implementation, "instances", instances)
	return implementation
}

// Private

func (v *ast_) generateParameter(
	isPlural bool,
	attributeName string,
	attributeType string,
) (
	parameter string,
) {
	parameter = v.getTemplate(singularRuleParameter)
	if attributeType == "string" {
		parameter = v.getTemplate(singularTokenParameter)
		if isPlural {
			parameter = v.getTemplate(pluralTokenParameter)
		}
	} else {
		if isPlural {
			parameter = v.getTemplate(pluralRuleParameter)
		}
	}
	parameter = replaceAll(parameter, "attributeName", attributeName)
	parameter = replaceAll(parameter, "attributeType", attributeType)
	return parameter
}

func (v *ast_) generateClass(
	className string,
) (
	class string,
) {
	var parameters string
	var attributes = v.analyzer_.GetReferences(className)
	if col.IsDefined(attributes) {
		// This class represents an inline rule.
		var references = attributes.GetIterator()
		for references.HasNext() {
			var reference = references.GetNext()
			var isPlural = v.isPlural(reference)
			var attributeName = generateVariableName(reference)
			var attributeType = generateVariableType(reference)
			parameters += v.generateParameter(isPlural, attributeName, attributeType)
		}
		if attributes.GetSize() > 0 {
			parameters += "\n\t"
		}
	} else {
		// This class represents a multiline rule.
		parameters += "\n\t\tany_ any,\n\t"
	}
	class = v.getTemplate(classDeclaration)
	class = replaceAll(class, "parameters", parameters)
	class = replaceAll(class, "className", className)
	return class
}

func (v *ast_) generateClasses() (
	classes string,
) {
	var rules = v.analyzer_.GetRuleNames().GetIterator()
	for rules.HasNext() {
		var className = rules.GetNext()
		classes += v.generateClass(className)
	}
	return classes
}

func (v *ast_) generateGetter(
	isPlural bool,
	attributeName string,
	attributeType string,
) (
	getter string,
) {
	if isPlural {
		// Add the collections module to the imports list.
		var path = `"github.com/craterdog/go-collection-framework/v4/collection"`
		var alias = "abs"
		v.modules_.SetValue(path, alias) // Modules are sorted by path.
	}
	getter = v.getTemplate(ruleGetterMethod)
	if attributeType == "string" {
		getter = v.getTemplate(tokenGetterMethod)
		if isPlural {
			getter = v.getTemplate(pluralTokenGetterMethod)
		}
	} else {
		if isPlural {
			getter = v.getTemplate(pluralRuleGetterMethod)
		}
	}
	getter = replaceAll(getter, "attributeName", attributeName)
	getter = replaceAll(getter, "attributeType", attributeType)
	return getter
}

func (v *ast_) generateHeader(
	wiki string,
) (
	header string,
) {
	header = v.getTemplate(packageHeader)
	header = replaceAll(header, "wiki", wiki)
	return header
}

func (v *ast_) generateImports() (
	imports string,
) {
	if v.modules_.IsEmpty() {
		// There are no modules that are imported.
		return imports
	}
	v.modules_.SortValues() // Modules are sorted by path, not by alias.
	var iterator = v.modules_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var alias = association.GetValue()
		var path = association.GetKey()
		imports += "\n\t" + alias + " " + path
	}
	imports += "\n"
	return imports
}

func (v *ast_) generateInstance(
	className string,
) (
	instance string,
) {
	var getters string
	var attributes = v.analyzer_.GetReferences(className)
	if col.IsDefined(attributes) {
		// This instance represents an inline rule.
		var references = attributes.GetIterator()
		for references.HasNext() {
			var reference = references.GetNext()
			var isPlural = v.isPlural(reference)
			var attributeName = generateVariableName(reference)
			var attributeType = generateVariableType(reference)
			getters += v.generateGetter(isPlural, attributeName, attributeType)
		}
	} else {
		// This instance represents a multiline rule.
		getters += "\n\tGetAny() any"
	}
	instance = v.getTemplate(instanceDeclaration)
	instance = replaceAll(instance, "publicMethods", v.getTemplate(publicMethods))
	var template string
	if col.IsDefined(getters) {
		template = v.getTemplate(attributeMethods)
		template = replaceAll(template, "getters", getters)
	}
	instance = replaceAll(instance, "attributeMethods", template)
	instance = replaceAll(instance, "className", className)
	return instance
}

func (v *ast_) generateInstances() (
	classes string,
) {
	var rules = v.analyzer_.GetRuleNames().GetIterator()
	for rules.HasNext() {
		var className = rules.GetNext()
		classes += v.generateInstance(className)
	}
	return classes
}

func (v *ast_) generateNotice() string {
	var notice = v.analyzer_.GetNotice()
	return notice
}

func (v *ast_) getTemplate(name string) string {
	var template = astTemplates_.GetValue(name)
	return template
}

func (v *ast_) isPlural(reference ast.ReferenceLike) bool {
	var cardinality = reference.GetOptionalCardinality()
	if col.IsUndefined(cardinality) {
		return false
	}
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		if actual.GetAny().(string) == "?" {
			return false
		}
	}
	return true
}

// PRIVATE GLOBALS

// Constants

const (
	modelTemplate           = "modelTemplate"
	packageHeader           = "packageHeader"
	classDeclaration        = "classDeclaration"
	singularRuleParameter   = "singularRuleParameter"
	pluralRuleParameter     = "pluralRuleParameter"
	singularTokenParameter  = "singularTokenParameter"
	pluralTokenParameter    = "pluralTokenParameter"
	instanceDeclaration     = "instanceDeclaration"
	publicMethods           = "publicMethods"
	attributeMethods        = "attributeMethods"
	ruleGetterMethod        = "ruleGetterMethod"
	pluralRuleGetterMethod  = "pluralRuleGetterMethod"
	tokenGetterMethod       = "tokenGetterMethod"
	pluralTokenGetterMethod = "pluralTokenGetterMethod"
)

var astTemplates_ = col.Catalog[string, string](
	map[string]string{
		packageHeader: `/*
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
*/`,
		classDeclaration: `
/*
<ClassName>ClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete <class-name>-like class.
*/
type <ClassName>ClassLike interface {
	// Constructor
	Make(<parameters>) <ClassName>Like
}
`,
		singularRuleParameter: `
		<attributeName_> <AttributeType>,`,
		pluralRuleParameter: `
		<attributeName_> abs.Sequential[<AttributeType>],`,
		singularTokenParameter: `
		<attributeName_> string,`,
		pluralTokenParameter: `
		<attributeName_> abs.Sequential[string],`,
		instanceDeclaration: `
/*
<ClassName>Like is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete <class-name>-like class.
*/
type <ClassName>Like interface {<PublicMethods><AttributeMethods>}
`,
		publicMethods: `
	// Public
	GetClass() <ClassName>ClassLike
`,
		attributeMethods: `
	// Attribute<Getters>
`,
		ruleGetterMethod: `
	Get<AttributeName>() <AttributeType>`,
		pluralRuleGetterMethod: `
	Get<AttributeName>() abs.Sequential[<AttributeType>]`,
		tokenGetterMethod: `
	Get<AttributeName>() string`,
		pluralTokenGetterMethod: `
	Get<AttributeName>() abs.Sequential[string]`,
		modelTemplate: `<Notice>

<Header>
package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// Classes
<Classes>
// Instances
<Instances>`,
	},
)
