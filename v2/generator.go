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
	return &generator_{
		modules_:   col.Catalog[string, pac.ModuleLike]().Make(),
		classes_:   col.Catalog[string, pac.ClassLike]().Make(),
		instances_: col.Catalog[string, pac.InstanceLike]().Make(),
	}
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
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var grammar = v.parseGrammar(directory)
	if grammar == nil {
		return
	}
	var model = v.processGrammar(grammar)
	v.generateModel(directory, model)
}

// Private

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
	var multiline = expression.GetInline()
	if multiline != nil {
		var iterator = multiline.GetAlternatives().GetIterator()
		for iterator.HasNext() {
			var alternative = iterator.GetNext()
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
	if sts.HasSuffix(name, "s") {
		name += "es"
	} else {
		name += "s"
	}
	return name
}

func (v *generator_) generateClass(
	name string,
	constructors col.Sequential[pac.ConstructorLike],
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
		pac.Constructors().MakeWithAttributes(constructors),
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
	attributes col.Sequential[pac.AttributeLike],
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
		pac.Attributes().MakeWithAttributes(attributes),
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
	attributes col.Sequential[pac.AttributeLike],
) {
	// Extract the constructor.
	var prefix pac.PrefixLike
	var arguments pac.ArgumentsLike
	var abstraction = pac.Abstraction().MakeWithAttributes(
		prefix,
		name+"Like",
		arguments,
	)
	var parameters pac.ParametersLike
	constructor = pac.Constructor().MakeWithAttributes(
		"MakeWith"+name,
		parameters,
		abstraction,
	)

	// Extract the attributes.
	var attributeList = col.List[pac.AttributeLike]().Make()
	var iterator = alternative.GetFactors().GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		attributes = v.processFactor(name, factor)
		attributeList.AppendValues(attributes)
	}

	return constructor, attributeList
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
	constructors col.Sequential[pac.ConstructorLike],
	attributes col.Sequential[pac.AttributeLike],
) {
	var constructorList = col.List[pac.ConstructorLike]().Make()
	var attributeList = col.List[pac.AttributeLike]().Make()
	var alternatives = v.extractAlternatives(expression)
	var iterator = alternatives.GetIterator()
	for iterator.HasNext() {
		var alternative = iterator.GetNext()
		var constructor, attributes = v.processAlternative(name, alternative)
		constructorList.AppendValue(constructor)
		attributeList.AppendValues(attributes)
	}
	return constructorList, attributeList
}

func (v *generator_) processFactor(
	name string,
	factor FactorLike,
) col.Sequential[pac.AttributeLike] {
	var cardinality = factor.GetCardinality()
	if cardinality != nil {
		var prefix = "col"
		var repository = `"github.com/craterdog/go-collection-framework/v3"`
		var module = pac.Module().MakeWithAttributes(
			prefix,
			repository,
		)
		// Must be sorted by the repository name NOT the prefix.
		v.modules_.SetValue(repository, module)
	}
	var attributes = col.List[pac.AttributeLike]().Make()
	var predicate = factor.GetPredicate()
	var element = predicate.GetElement()
	var inversion = predicate.GetInversion()
	var precedence = predicate.GetPrecedence()
	switch {
	case element != nil:
	case inversion != nil:
	case precedence != nil:
		var expression = precedence.GetExpression()
		var _, sequence = v.processExpression(name, expression)
		attributes.AppendValues(sequence)
	default:
		panic("Found an empty predicate.")
	}
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
