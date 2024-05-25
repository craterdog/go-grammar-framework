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

package agent

import (
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	cds "github.com/craterdog/go-grammar-framework/v4/cdsn/ast"
	mod "github.com/craterdog/go-model-framework/v4"
	sts "strings"
	tim "time"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var generatorClass = &generatorClass_{
	// This class does not initialize any class constants.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
	// This class does not define any class constants.
}

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	return &generator_{
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	class_     GeneratorClassLike
	tokens_    col.SetLike[string]
	modules_   col.CatalogLike[string, mod.ModuleLike]
	classes_   col.CatalogLike[string, mod.ClassLike]
	instances_ col.CatalogLike[string, mod.InstanceLike]
}

// Attributes

func (v *generator_) GetClass() GeneratorClassLike {
	return v.class_
}

// Public

func (v *generator_) CreateSyntax(
	name string,
	copyright string,
) cds.SyntaxLike {
	// Center and insert the copyright notice into the syntax template.
	copyright = v.expandCopyright(copyright)
	var source = sts.ReplaceAll(syntaxTemplate_, "<Copyright>", copyright)
	source = sts.ReplaceAll(source, "<NAME>", sts.ToUpper(name))
	source = sts.ReplaceAll(source, "<Name>", name)
	source = source[1:] // Strip off the leading "\n".

	// Parse the syntax.
	var parser = Parser().Make()
	var syntax = parser.ParseSource(source)
	return syntax
}

func (v *generator_) GenerateAST(syntax cds.SyntaxLike) mod.ModelLike {
	// Create the AST class model template.
	v.processSyntax(syntax)
	var source = modelTemplate_ + astTemplate_
	var notice = v.extractNotice(syntax)
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = "<module>"
	source = sts.ReplaceAll(source, "<module>", module)
	source = sts.ReplaceAll(source, "<package>", "ast")

	// Parse the AST class model template.
	var parser = mod.Parser()
	var model = parser.ParseSource(source)

	// Add additional definitions to the AST class model.
	v.addModules(model)
	v.addClasses(model)
	v.addInstances(model)

	return model
}

func (v *generator_) GenerateAgent(syntax cds.SyntaxLike) mod.ModelLike {
	// Create the agent class model template.
	v.processSyntax(syntax)
	var source = modelTemplate_ + agentTemplate_
	var notice = v.extractNotice(syntax)
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = "<module>"
	source = sts.ReplaceAll(source, "<module>", module)
	source = sts.ReplaceAll(source, "<package>", "agent")
	var class = v.extractClassName(syntax)
	source = sts.ReplaceAll(source, "<Class>", class)
	var parameter = v.makeLowercase(class)
	source = sts.ReplaceAll(source, "<parameter>", parameter)

	// Parse the agent class model template.
	var parser = mod.Parser()
	var model = parser.ParseSource(source)

	// Add additional definitions to the class model.
	v.addTokens(model)

	return model
}

func (v *generator_) GenerateFormatter(model mod.ModelLike) string {
	var source = formatterTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = "<module>"
	source = sts.ReplaceAll(source, "<module>", module)
	var class = "<class>"
	source = sts.ReplaceAll(source, "<Class>", v.makeUppercase(class))
	source = sts.ReplaceAll(source, "<class>", v.makeLowercase(class))
	return source
}

func (v *generator_) GenerateParser(model mod.ModelLike) string {
	var source = parserTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = "<module>"
	source = sts.ReplaceAll(source, "<module>", module)
	var packageName = model.GetHeader().GetIdentifier()
	source = sts.ReplaceAll(source, "<Package>", v.makeUppercase(packageName))
	var class = "<class>"
	source = sts.ReplaceAll(source, "<Class>", v.makeUppercase(class))
	source = sts.ReplaceAll(source, "<class>", v.makeLowercase(class))
	return source
}

func (v *generator_) GenerateScanner(model mod.ModelLike) string {
	var source = scannerTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = "<module>"
	source = sts.ReplaceAll(source, "<module>", module)
	return source
}

func (v *generator_) GenerateToken(model mod.ModelLike) string {
	var source = tokenTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = "<module>"
	source = sts.ReplaceAll(source, "<module>", module)
	return source
}

func (v *generator_) GenerateValidator(model mod.ModelLike) string {
	var source = validatorTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = "<module>"
	source = sts.ReplaceAll(source, "<module>", module)
	var class = "<class>"
	source = sts.ReplaceAll(source, "<Class>", v.makeUppercase(class))
	source = sts.ReplaceAll(source, "<class>", v.makeLowercase(class))
	return source
}

// Private

func (v *generator_) addClasses(model mod.ModelLike) {
	var classes = model.GetClasses()
	classes.RemoveAll() // Remove the dummy placeholder.
	classes.AppendValues(v.classes_.GetValues(v.classes_.GetKeys()))
}

func (v *generator_) addModules(model mod.ModelLike) {
	var modules = model.GetModules()
	modules.RemoveAll() // Remove the dummy placeholder.
	modules.AppendValues(v.modules_.GetValues(v.modules_.GetKeys()))
}

func (v *generator_) addInstances(model mod.ModelLike) {
	var instances = model.GetInstances()
	instances.RemoveAll() // Remove the dummy placeholder.
	instances.AppendValues(v.instances_.GetValues(v.instances_.GetKeys()))
}

func (v *generator_) addTokens(model mod.ModelLike) {
	var types = model.GetTypes()
	var enumeration = types.GetValue(1).GetEnumeration()
	var identifiers = enumeration.GetIdentifiers()
	identifiers.AppendValues(v.tokens_)
}

func (v *generator_) consolidateAttributes(
	attributes col.ListLike[mod.AttributeLike],
) {
	// Compare each attribute and remove duplicates.
	for i := 1; i <= attributes.GetSize(); i++ {
		var attribute = attributes.GetValue(i)
		var first = attribute.GetIdentifier()
		for j := i + 1; j <= attributes.GetSize(); {
			var second = attributes.GetValue(j).GetIdentifier()
			switch {
			case first == second:
				attributes.RemoveValue(j)
			case first == second[:len(second)-1] && sts.HasSuffix(second, "s"):
				attribute = attributes.GetValue(j)
				attributes.SetValue(i, attribute)
				attributes.RemoveValue(j)
			case first == second[:len(second)-2] && sts.HasSuffix(second, "es"):
				attribute = attributes.GetValue(j)
				attributes.SetValue(i, attribute)
				attributes.RemoveValue(j)
			case second == first[:len(first)-1] && sts.HasSuffix(first, "s"):
				attributes.RemoveValue(j)
			case second == first[:len(first)-2] && sts.HasSuffix(first, "es"):
				attributes.RemoveValue(j)
			default:
				// We only increment the index j if we didn't remove anything.
				j++
			}
		}
	}
}

func (v *generator_) consolidateConstructors(
	constructors col.ListLike[mod.ConstructorLike],
) {
	// Compare each constructor and remove duplicates.
	for i := 1; i <= constructors.GetSize(); i++ {
		var constructor = constructors.GetValue(i)
		var first = constructor.GetIdentifier()
		for j := i + 1; j <= constructors.GetSize(); {
			var second = constructors.GetValue(j).GetIdentifier()
			switch {
			case first == second:
				constructors.RemoveValue(j)
			case first == second[:len(second)-1] && sts.HasSuffix(second, "s"):
				constructor = constructors.GetValue(j)
				constructors.SetValue(i, constructor)
				constructors.RemoveValue(j)
			case first == second[:len(second)-2] && sts.HasSuffix(second, "es"):
				constructor = constructors.GetValue(j)
				constructors.SetValue(i, constructor)
				constructors.RemoveValue(j)
			case second == first[:len(first)-1] && sts.HasSuffix(first, "s"):
				constructors.RemoveValue(j)
			case second == first[:len(first)-2] && sts.HasSuffix(first, "es"):
				constructors.RemoveValue(j)
			default:
				// We only increment the index j if we didn't remove anything.
				j++
			}
		}
	}
}

func (v *generator_) consolidateLists(
	attributes col.ListLike[mod.AttributeLike],
) {
	// Compare each attribute and make lists out of duplicates.
	for i := 1; i <= attributes.GetSize(); i++ {
		var attribute = attributes.GetValue(i)
		var first = attribute.GetIdentifier()
		for j := i + 1; j <= attributes.GetSize(); {
			var second = attributes.GetValue(j).GetIdentifier()
			switch {
			case first == second:
				attribute = attributes.GetValue(i)
				attributes.SetValue(i, v.makeList(attribute))
				attributes.RemoveValue(j)
			case first == second[:len(second)-1] && sts.HasSuffix(second, "s"):
				attribute = attributes.GetValue(j)
				attributes.SetValue(i, attribute)
				attributes.RemoveValue(j)
			case first == second[:len(second)-2] && sts.HasSuffix(second, "es"):
				attribute = attributes.GetValue(j)
				attributes.SetValue(i, attribute)
				attributes.RemoveValue(j)
			case second == first[:len(first)-1] && sts.HasSuffix(first, "s"):
				attributes.RemoveValue(j)
			case second == first[:len(first)-2] && sts.HasSuffix(first, "es"):
				attributes.RemoveValue(j)
			default:
				// We only increment the index j if we didn't remove anything.
				j++
			}
		}
	}
}

func (v *generator_) expandCopyright(copyright string) string {
	var maximum = 78
	var length = len(copyright)
	if length > maximum {
		var message = fmt.Sprintf(
			"The copyright notice cannot be longer than 78 characters: %v",
			copyright,
		)
		panic(message)
	}
	if length == 0 {
		copyright = fmt.Sprintf(
			"Copyright (c) %v.  All Rights Reserved.",
			tim.Now().Year(),
		)
		length = len(copyright)
	}
	var padding = (maximum - length) / 2
	for range padding {
		copyright = " " + copyright + " "
	}
	if len(copyright) < maximum {
		copyright = " " + copyright
	}
	copyright = "." + copyright + "."
	return copyright
}

func (v *generator_) extractAlternatives(expression cds.ExpressionLike) col.ListLike[cds.AlternativeLike] {
	var notation = cdc.Notation().Make()
	var alternatives = col.List[cds.AlternativeLike](notation).Make()
	var inline = expression.GetInline()
	if inline != nil {
		var iterator = inline.GetAlternatives().GetIterator()
		for iterator.HasNext() {
			var alternative = iterator.GetNext()
			alternatives.AppendValue(alternative)
		}
	}
	var multiline = expression.GetMultiline()
	if multiline != nil {
		var iterator = multiline.GetLines().GetIterator()
		for iterator.HasNext() {
			var line = iterator.GetNext()
			var alternative = line.GetAlternative()
			alternatives.AppendValue(alternative)
		}
	}
	return alternatives
}

func (v *generator_) extractClassName(syntax cds.SyntaxLike) string {
	var definition = syntax.GetDefinitions().GetValue(1)
	var expression = definition.GetExpression()
	var alternatives = v.extractAlternatives(expression)
	var alternative = alternatives.GetIterator().GetNext()
	var factors = alternative.GetFactors()
	var factor = factors.GetIterator().GetNext()
	var predicate = factor.GetPredicate()
	var element = predicate.GetElement()
	var name = element.GetName()
	return name
}

func (v *generator_) extractNotice(syntax cds.SyntaxLike) string {
	var header = syntax.GetHeaders().GetValue(1)
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = Scanner().MatchToken(CommentToken, comment).GetValue(2)

	// Add the Go style comment delimiters.
	notice = "/*\n" + notice + "\n*/"

	return notice
}

func (v *generator_) extractParameters(
	attributes col.ListLike[mod.AttributeLike],
) col.ListLike[mod.ParameterLike] {
	var notation = cdc.Notation().Make()
	var parameters = col.List[mod.ParameterLike](notation).Make()
	var iterator = attributes.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var identifier = sts.TrimPrefix(attribute.GetIdentifier(), "Get")
		var abstraction = attribute.GetAbstraction()
		var parameter = mod.Parameter(
			v.makeLowercase(identifier),
			abstraction,
		)
		parameters.AppendValue(parameter)
	}
	return parameters
}

func (v *generator_) generateClass(
	name string,
	constructors col.ListLike[mod.ConstructorLike],
) mod.ClassLike {
	var comment = classCommentTemplate_[1:] // Strip off leading newline.
	comment = sts.ReplaceAll(comment, "<Class>", name)
	comment = sts.ReplaceAll(comment, "<class>", sts.ToLower(name))
	var declaration = mod.Declaration(
		comment,
		name+"ClassLike",
	)
	var class = mod.Class(
		declaration,
		constructors,
	)
	return class
}

func (v *generator_) generateInstance(
	name string,
	attributes col.ListLike[mod.AttributeLike],
) mod.InstanceLike {
	var comment = instanceCommentTemplate_[1:] // Strip off leading newline.
	comment = sts.ReplaceAll(comment, "<Class>", name)
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
	attributes.InsertValue(0, attribute)
	var instance = mod.Instance(
		declaration,
		attributes,
	)
	return instance
}

func (v *generator_) isLowercase(identifier string) bool {
	return uni.IsLower([]rune(identifier)[0])
}

func (v *generator_) isUppercase(identifier string) bool {
	return uni.IsUpper([]rune(identifier)[0])
}

func (v *generator_) makeList(attribute mod.AttributeLike) mod.AttributeLike {
	var prefix = "col"
	var path = `"github.com/craterdog/go-collection-framework/v4/collection"`
	var module = mod.Module(
		prefix,
		path,
	)
	v.modules_.SetValue(path, module)
	var identifier = attribute.GetIdentifier()
	identifier = v.makePlural(identifier)
	var abstraction = attribute.GetAbstraction()
	var notation = cdc.Notation().Make()
	var arguments = col.List[mod.AbstractionLike](notation).Make()
	arguments.AppendValue(abstraction)
	abstraction = mod.Abstraction(
		mod.Prefix(prefix, mod.AliasPrefix),
		"ListLike",
		arguments,
	)
	attribute = mod.Attribute(identifier, abstraction)
	return attribute
}

func (v *generator_) makeLowercase(identifier string) string {
	runes := []rune(identifier)
	runes[0] = uni.ToLower(runes[0])
	identifier = string(runes)
	if reserved_[identifier] {
		identifier += "_"
	}
	return identifier
}

func (v *generator_) makePlural(identifier string) string {
	if sts.HasSuffix(identifier, "s") {
		identifier += "es"
	} else {
		identifier += "s"
	}
	return identifier
}

func (v *generator_) makeUppercase(identifier string) string {
	runes := []rune(identifier)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

func (v *generator_) processAlternative(
	name string,
	alternative cds.AlternativeLike,
) (
	constructor mod.ConstructorLike,
	attributes col.ListLike[mod.AttributeLike],
) {
	// Extract the attributes.
	var notation = cdc.Notation().Make()
	attributes = col.List[mod.AttributeLike](notation).Make()
	var iterator = alternative.GetFactors().GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		var values = v.processFactor(name, factor)
		attributes.AppendValues(values)
	}
	v.consolidateLists(attributes)

	// Extract the constructor.
	var abstraction = mod.Abstraction(name + "Like")
	if !attributes.IsEmpty() {
		var identifier = "MakeWithAttributes"
		if attributes.GetSize() == 1 {
			identifier = attributes.GetValue(1).GetIdentifier()
			identifier = sts.TrimPrefix(identifier, "Get")
			identifier = "MakeWith" + identifier
		}
		var parameters = v.extractParameters(attributes)
		constructor = mod.Constructor(
			identifier,
			parameters,
			abstraction,
		)
	}

	return constructor, attributes
}

func (v *generator_) processDefinition(
	definition cds.DefinitionLike,
) {
	var name = definition.GetName()
	var expression = definition.GetExpression()
	if v.isLowercase(name) {
		v.processToken(name, expression)
	} else {
		v.processRule(name, expression)
	}
}

func (v *generator_) processExpression(
	name string,
	expression cds.ExpressionLike,
) (
	constructors col.ListLike[mod.ConstructorLike],
	attributes col.ListLike[mod.AttributeLike],
) {
	// Process the expression alternatives.
	var notation = cdc.Notation().Make()
	constructors = col.List[mod.ConstructorLike](notation).Make()
	attributes = col.List[mod.AttributeLike](notation).Make()
	var alternatives = v.extractAlternatives(expression)
	var iterator = alternatives.GetIterator()
	for iterator.HasNext() {
		var alternative = iterator.GetNext()
		var constructor, values = v.processAlternative(name, alternative)
		if constructor != nil {
			constructors.AppendValue(constructor)
		}
		attributes.AppendValues(values)
	}

	// Add a default constructor if necessary.
	if constructors.IsEmpty() {
		var abstraction = mod.Abstraction(
			name + "Like",
		)
		var constructor = mod.Constructor(
			"Make",
			abstraction,
		)
		constructors.AppendValue(constructor)
	}

	// Consolidate any duplicate methods.
	v.consolidateConstructors(constructors)
	v.consolidateAttributes(attributes)

	return constructors, attributes
}

func (v *generator_) processFactor(
	name string,
	factor cds.FactorLike,
) (attributes col.ListLike[mod.AttributeLike]) {
	var isSequential bool
	var cardinality = factor.GetCardinality()
	if cardinality != nil {
		var constraint = cardinality.GetConstraint()
		isSequential = constraint.GetFirst() != "0" || constraint.GetLast() != "1"
	}
	var identifier string
	var abstraction mod.AbstractionLike
	var attribute mod.AttributeLike
	var notation = cdc.Notation().Make()
	attributes = col.List[mod.AttributeLike](notation).Make()
	var predicate = factor.GetPredicate()
	var atom = predicate.GetAtom()
	var element = predicate.GetElement()
	var filter = predicate.GetFilter()
	var precedence = predicate.GetPrecedence()
	switch {
	case atom != nil:
	case element != nil:
		identifier = element.GetName()
		if len(identifier) > 0 {
			if v.isUppercase(identifier) {
				abstraction = mod.Abstraction(identifier + "Like")
			} else {
				var tokenType = v.makeUppercase(identifier) + "Token"
				v.tokens_.AddValue(tokenType)
				abstraction = mod.Abstraction("string")
			}
			attribute = mod.Attribute(
				"Get"+v.makeUppercase(identifier),
				abstraction,
			)
			if isSequential {
				attribute = v.makeList(attribute)
			}
			attributes.AppendValue(attribute)
		}
	case filter != nil:
	case precedence != nil:
		var expression = precedence.GetExpression()
		var _, values = v.processExpression(name, expression)
		attributes.AppendValues(values)
	default:
		panic("Found an empty predicate.")
	}
	return attributes
}

func (v *generator_) processRule(
	name string,
	expression cds.ExpressionLike,
) {
	// Process the full expression first.
	var constructors, attributes = v.processExpression(name, expression)

	// Create the class interface.
	var class = v.generateClass(name, constructors)
	v.classes_.SetValue(name, class)

	// Create the instance interface.
	var instance = v.generateInstance(name, attributes)
	v.instances_.SetValue(name, instance)
}

func (v *generator_) processSyntax(syntax cds.SyntaxLike) {
	// Initialize the collections.
	var array = []string{
		"DelimiterToken",
		"EOFToken",
		"EOLToken",
		"SpaceToken",
	}
	var notation = cdc.Notation().Make()
	v.tokens_ = col.Set[string](notation).MakeFromArray(array)
	v.modules_ = col.Catalog[string, mod.ModuleLike](notation).Make()
	v.classes_ = col.Catalog[string, mod.ClassLike](notation).Make()
	v.instances_ = col.Catalog[string, mod.InstanceLike](notation).Make()

	// Process the syntax definitions.
	var iterator = syntax.GetDefinitions().GetIterator()
	iterator.GetNext() // Skip the first rule.
	for iterator.HasNext() {
		var definition = iterator.GetNext()
		v.processDefinition(definition)
	}
}

func (v *generator_) processToken(
	name string,
	expression cds.ExpressionLike,
) {
	// Ignore token definitions for now.
}

var reserved_ = map[string]bool{
	"byte":      true,
	"case":      true,
	"complex":   true,
	"copy":      true,
	"default":   true,
	"error":     true,
	"false":     true,
	"import":    true,
	"interface": true,
	"map":       true,
	"nil":       true,
	"package":   true,
	"range":     true,
	"real":      true,
	"return":    true,
	"rune":      true,
	"string":    true,
	"switch":    true,
	"true":      true,
	"type":      true,
}
