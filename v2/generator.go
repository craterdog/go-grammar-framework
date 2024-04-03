/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package grammars

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v3"
	pac "github.com/craterdog/go-package-framework/v2"
	osx "os"
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
	return &generator_{}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	modules_   col.CatalogLike[string, pac.ModuleLike]
	classes_   col.CatalogLike[string, pac.ClassLike]
	instances_ col.CatalogLike[string, pac.InstanceLike]
}

// Public

func (v *generator_) CreateGrammar(directory string, copyright string) {
	// Insert the copyright statement into a new grammar template.
	copyright = v.expandCopyright(copyright)
	var template = sts.ReplaceAll(grammarTemplate_, "<Copyright>", copyright)

	// Save the new grammar template into the directory.
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	v.createDirectory(directory)
	var grammarFile = directory + "Grammar.cdsn"
	var bytes = []byte(template[1:]) // Remove leading "\n".
	var err = osx.WriteFile(grammarFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) GenerateModel(directory string) {
	// Initialize the catalogs.
	v.modules_ = col.Catalog[string, pac.ModuleLike]().Make()
	v.classes_ = col.Catalog[string, pac.ClassLike]().Make()
	v.instances_ = col.Catalog[string, pac.InstanceLike]().Make()

	// Parse the grammar file.
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var grammar = v.parseGrammar(directory)
	if grammar == nil {
		return
	}

	// Generate the Package.go file.
	var model = v.processGrammar(grammar)
	v.generateModel(directory, model)
}

// Private

func (v *generator_) consolidateAttributes(
	attributes col.ListLike[pac.AttributeLike],
) {
	for i := 1; i <= attributes.GetSize(); i++ {
		var attribute = attributes.GetValue(i)
		var first = attribute.GetIdentifier()
		for j := i + 1; j <= attributes.GetSize(); {
			var second = attributes.GetValue(j).GetIdentifier()
			switch {
			case first == second:
				attributes.RemoveValue(j)
				var attribute = attributes.GetValue(i)
				attributes.SetValue(i, v.makeSequence(attribute))
			case first == second[:len(second)-1] && sts.HasSuffix(second, "s"):
				attribute = attributes.GetValue(j)
				attributes.RemoveValue(j)
				attributes.SetValue(i, attribute)
			case first == second[:len(second)-2] && sts.HasSuffix(second, "es"):
				attribute = attributes.GetValue(j)
				attributes.RemoveValue(j)
				attributes.SetValue(i, attribute)
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
	constructors col.ListLike[pac.ConstructorLike],
) {
	for i := 1; i <= constructors.GetSize(); i++ {
		var first = constructors.GetValue(i).GetIdentifier()
		for j := i + 1; j <= constructors.GetSize(); {
			var second = constructors.GetValue(j).GetIdentifier()
			if first == second {
				constructors.RemoveValue(j)
			} else {
				// We only increment the index j if we didn't remove anything.
				j++
			}
		}
	}
}

func (v *generator_) createDirectory(directory string) {
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
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

func (v *generator_) extractAlternatives(expression ExpressionLike) col.Sequential[AlternativeLike] {
	var alternatives = col.List[AlternativeLike]().Make()
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

func (v *generator_) extractName(definition DefinitionLike) string {
	var expression = definition.GetExpression()
	var alternatives = v.extractAlternatives(expression)
	var alternative = alternatives.GetIterator().GetNext()
	var factors = alternative.GetFactors()
	var factor = factors.GetIterator().GetNext()
	var predicate = factor.GetPredicate()
	var element = predicate.GetElement()
	var name = element.GetName()
	name = sts.ToLower(name)
	name = v.makePlural(name)
	return name
}

func (v *generator_) extractParameters(
	attributes col.Sequential[pac.AttributeLike],
) pac.ParametersLike {
	var parameterList = col.List[pac.ParameterLike]().Make()
	var iterator = attributes.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var identifier = sts.TrimPrefix(attribute.GetIdentifier(), "Get")
		var abstraction = attribute.GetAbstraction()
		var parameter = pac.Parameter().MakeWithAttributes(
			v.makeLowercase(identifier),
			abstraction,
		)
		parameterList.AppendValue(parameter)
	}
	var parameters = pac.Parameters().MakeWithAttributes(parameterList)
	return parameters
}

func (v *generator_) generateClass(
	name string,
	constructors pac.ConstructorsLike,
) pac.ClassLike {
	var comment = classCommentTemplate_
	comment = sts.ReplaceAll(comment, "<ClassName>", name)
	comment = sts.ReplaceAll(comment, "<class-name>", sts.ToLower(name))
	var parameters pac.ParametersLike
	var declaration = pac.Declaration().MakeWithAttributes(
		comment,
		name+"ClassLike",
		parameters,
	)
	var constants pac.ConstantsLike
	var functions pac.FunctionsLike
	var class = pac.Class().MakeWithAttributes(
		declaration,
		constants,
		constructors,
		functions,
	)
	return class
}

func (v *generator_) generateHeader(packageName string) pac.HeaderLike {
	var comment = v.generatePackageComment(packageName)
	var header = pac.Header().MakeWithAttributes(comment, packageName)
	return header
}

func (v *generator_) generateImports() pac.ImportsLike {
	var imports pac.ImportsLike
	if !v.modules_.IsEmpty() {
		v.modules_.SortValues()
		var values = v.modules_.GetValues(v.modules_.GetKeys())
		var modules = pac.Modules().MakeWithAttributes(values)
		imports = pac.Imports().MakeWithAttributes(modules)
	}
	return imports
}

func (v *generator_) generateInstance(
	name string,
	attributes pac.AttributesLike,
) pac.InstanceLike {
	var comment = instanceCommentTemplate_
	comment = sts.ReplaceAll(comment, "<ClassName>", name)
	comment = sts.ReplaceAll(comment, "<class-name>", sts.ToLower(name))
	var parameters pac.ParametersLike
	var declaration = pac.Declaration().MakeWithAttributes(
		comment,
		name+"Like",
		parameters,
	)
	var abstractions pac.AbstractionsLike
	var methods pac.MethodsLike
	var instance = pac.Instance().MakeWithAttributes(
		declaration,
		attributes,
		abstractions,
		methods,
	)
	return instance
}

func (v *generator_) generateInterfaces() pac.InterfacesLike {
	var aspects pac.AspectsLike

	v.classes_.SortValues()
	var classes = pac.Classes().MakeWithAttributes(
		v.classes_.GetValues(v.classes_.GetKeys()),
	)

	v.instances_.SortValues()
	var instances = pac.Instances().MakeWithAttributes(
		v.instances_.GetValues(v.instances_.GetKeys()),
	)

	var interfaces = pac.Interfaces().MakeWithAttributes(
		aspects,
		classes,
		instances,
	)
	return interfaces
}

func (v *generator_) generateModel(directory string, model pac.ModelLike) {
	var validator = pac.Validator().Make()
	validator.ValidateModel(model)
	var formatter = pac.Formatter().Make()
	var source = formatter.FormatModel(model)
	var bytes = []byte(source)
	var err = osx.WriteFile(directory+"Package.go", bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) generateNotice(grammar GrammarLike) pac.NoticeLike {
	var header = grammar.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the grammar style comment delimiters.
	comment = Scanner().MatchToken(CommentToken, comment).GetValue(2)

	// Add the Go style comment delimiters.
	comment = "/*\n" + comment + "\n*/\n"

	var notice = pac.Notice().MakeWithAttributes(comment)
	return notice
}

func (v *generator_) generatePackageComment(
	packageName string,
) string {
	var comment = packageCommentTemplate_
	comment = sts.ReplaceAll(comment, "<packagename>", packageName)
	return comment
}

func (v *generator_) isLowercase(identifier string) bool {
	return uni.IsLower([]rune(identifier)[0])
}

func (v *generator_) isUppercase(identifier string) bool {
	return uni.IsUpper([]rune(identifier)[0])
}

func (v *generator_) makeLowercase(identifier string) string {
	runes := []rune(identifier)
	runes[0] = uni.ToLower(runes[0])
	return string(runes)
}

func (v *generator_) makePlural(identifier string) string {
	if sts.HasSuffix(identifier, "s") {
		identifier += "es"
	} else {
		identifier += "s"
	}
	return identifier
}

func (v *generator_) makeSequence(attribute pac.AttributeLike) pac.AttributeLike {
	var identifier = attribute.GetIdentifier()
	identifier = v.makePlural(identifier)
	var abstraction = attribute.GetAbstraction()
	var argumentList = col.List[pac.AbstractionLike]().Make()
	argumentList.AppendValue(abstraction)
	var arguments = pac.Arguments().MakeWithAttributes(argumentList)
	abstraction = pac.Abstraction().MakeWithAttributes(
		pac.Prefix().MakeWithAttributes("col", pac.AliasPrefix),
		"Sequential",
		arguments,
	)
	var parameter pac.ParameterLike
	attribute = pac.Attribute().MakeWithAttributes(identifier, parameter, abstraction)
	return attribute
}

func (v *generator_) makeUppercase(identifier string) string {
	runes := []rune(identifier)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

func (v *generator_) parseGrammar(directory string) GrammarLike {
	var grammarFile = directory + "Grammar.cdsn"
	var bytes, err = osx.ReadFile(grammarFile)
	if err != nil {
		var message = fmt.Sprintf(
			"The specified directory is missing a grammar file: %v",
			grammarFile,
		)
		panic(message)
	}
	var source = string(bytes)
	var parser = Parser().Make()
	var grammar = parser.ParseSource(source)
	var validator = Validator().Make()
	validator.ValidateGrammar(grammar)
	return grammar
}

func (v *generator_) processAlternative(name string, alternative AlternativeLike) (
	constructor pac.ConstructorLike,
	attributes pac.AttributesLike,
) {
	// Extract the attributes.
	var attributeList = col.List[pac.AttributeLike]().Make()
	var iterator = alternative.GetFactors().GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		attributes = v.processFactor(name, factor)
		attributeList.AppendValues(attributes.GetSequence())
	}
	v.consolidateAttributes(attributeList)

	// Extract the constructor.
	var prefix pac.PrefixLike
	var arguments pac.ArgumentsLike
	var abstraction = pac.Abstraction().MakeWithAttributes(
		prefix,
		name+"Like",
		arguments,
	)
	name = "MakeWithAttributes"
	if attributeList.GetSize() == 1 {
		name = sts.TrimPrefix(attributeList.GetValue(1).GetIdentifier(), "Get")
		name = "MakeWith" + name
	}
	if !attributeList.IsEmpty() {
		var parameters = v.extractParameters(attributeList)
		constructor = pac.Constructor().MakeWithAttributes(
			name,
			parameters,
			abstraction,
		)
	}

	attributes = pac.Attributes().MakeWithAttributes(attributeList)
	return constructor, attributes
}

func (v *generator_) processDefinition(
	definition DefinitionLike,
) {
	var name = definition.GetName()
	var expression = definition.GetExpression()
	if v.isLowercase(name) {
		v.processToken(name, expression)
	} else {
		v.processRule(name, expression)
	}
}

func (v *generator_) processExpression(name string, expression ExpressionLike) (
	constructors pac.ConstructorsLike,
	attributes pac.AttributesLike,
) {
	var constructorList = col.List[pac.ConstructorLike]().Make()
	var attributeList = col.List[pac.AttributeLike]().Make()
	var alternatives = v.extractAlternatives(expression)
	var iterator = alternatives.GetIterator()
	for iterator.HasNext() {
		var alternative = iterator.GetNext()
		var constructor, attributes = v.processAlternative(name, alternative)
		if constructor != nil {
			constructorList.AppendValue(constructor)
		}
		attributeList.AppendValues(attributes.GetSequence())
	}
	v.consolidateConstructors(constructorList)
	v.consolidateAttributes(attributeList)
	constructors = pac.Constructors().MakeWithAttributes(constructorList)
	attributes = pac.Attributes().MakeWithAttributes(attributeList)
	return constructors, attributes
}

func (v *generator_) processFactor(
	name string,
	factor FactorLike,
) (attributes pac.AttributesLike) {
	var isSequential bool
	var cardinality = factor.GetCardinality()
	if cardinality != nil {
		var constraint = cardinality.GetConstraint()
		if constraint.GetFirst() != "0" || constraint.GetLast() != "1" {
			isSequential = true
			var prefix = "col"
			var repository = `"github.com/craterdog/go-collection-framework/v3"`
			var module = pac.Module().MakeWithAttributes(
				prefix,
				repository,
			)
			// Must be sorted by the repository name NOT the prefix.
			v.modules_.SetValue(repository, module)
		}
	}
	var identifier string
	var abstraction pac.AbstractionLike
	var attribute pac.AttributeLike
	var attributeList = col.List[pac.AttributeLike]().Make()
	var predicate = factor.GetPredicate()
	var element = predicate.GetElement()
	var inversion = predicate.GetInversion()
	var precedence = predicate.GetPrecedence()
	switch {
	case element != nil:
		identifier = element.GetName()
		if len(identifier) > 0 {
			var prefix pac.PrefixLike
			var arguments pac.ArgumentsLike
			if v.isUppercase(identifier) {
				abstraction = pac.Abstraction().MakeWithAttributes(
					prefix,
					identifier+"Like",
					arguments,
				)
				if isSequential {
					identifier = v.makePlural(identifier)
					var argumentList = col.List[pac.AbstractionLike]().Make()
					argumentList.AppendValue(abstraction)
					arguments = pac.Arguments().MakeWithAttributes(argumentList)
					abstraction = pac.Abstraction().MakeWithAttributes(
						pac.Prefix().MakeWithAttributes("col", pac.AliasPrefix),
						"Sequential",
						arguments,
					)
				}
			} else {
				abstraction = pac.Abstraction().MakeWithAttributes(
					prefix,
					"string",
					arguments,
				)
			}
			var parameter pac.ParameterLike
			attribute = pac.Attribute().MakeWithAttributes(
				"Get"+v.makeUppercase(identifier),
				parameter,
				abstraction,
			)
			attributeList.AppendValue(attribute)
		}
	case inversion != nil:
	case precedence != nil:
		var expression = precedence.GetExpression()
		var _, attributes = v.processExpression(name, expression)
		attributeList.AppendValues(attributes.GetSequence())
	default:
		panic("Found an empty predicate.")
	}
	attributes = pac.Attributes().MakeWithAttributes(attributeList)
	return attributes
}

func (v *generator_) processGrammar(grammar GrammarLike) pac.ModelLike {
	var iterator = grammar.GetDefinitions().GetIterator()
	var definition = iterator.GetNext()
	var packageName = v.extractName(definition)
	for iterator.HasNext() {
		definition = iterator.GetNext()
		v.processDefinition(definition)
	}
	var notice = v.generateNotice(grammar)
	var header = v.generateHeader(packageName)
	var imports = v.generateImports()
	var types pac.TypesLike
	var interfaces = v.generateInterfaces()
	var model = pac.Model().MakeWithAttributes(
		notice,
		header,
		imports,
		types,
		interfaces,
	)
	return model
}

func (v *generator_) processRule(
	name string,
	expression ExpressionLike,
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

func (v *generator_) processToken(name string, expression ExpressionLike) {
	// Ignore token definitions for now.
}
