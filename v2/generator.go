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
	mod "github.com/craterdog/go-model-framework/v2"
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
	modules_   col.CatalogLike[string, mod.ModuleLike]
	classes_   col.CatalogLike[string, mod.ClassLike]
	instances_ col.CatalogLike[string, mod.InstanceLike]
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
	v.modules_ = col.Catalog[string, mod.ModuleLike]().Make()
	v.classes_ = col.Catalog[string, mod.ClassLike]().Make()
	v.instances_ = col.Catalog[string, mod.InstanceLike]().Make()

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

func (v *generator_) extractAlternatives(expression ExpressionLike) col.ListLike[AlternativeLike] {
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
	attributes col.ListLike[mod.AttributeLike],
) col.ListLike[mod.ParameterLike] {
	var parameters = col.List[mod.ParameterLike]().Make()
	var iterator = attributes.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var identifier = sts.TrimPrefix(attribute.GetIdentifier(), "Get")
		var abstraction = attribute.GetAbstraction()
		var parameter = mod.Parameter().MakeWithAttributes(
			v.makeLowercase(identifier),
			abstraction,
		)
		parameters.AppendValue(parameter)
	}
	return parameters
}

func (v *generator_) generateAspects() col.ListLike[mod.AspectLike] {
	var aspects col.ListLike[mod.AspectLike]
	return aspects
}

func (v *generator_) generateClass(
	name string,
	constructors col.ListLike[mod.ConstructorLike],
) mod.ClassLike {
	var comment = classCommentTemplate_
	comment = sts.ReplaceAll(comment, "<ClassName>", name)
	comment = sts.ReplaceAll(comment, "<class-name>", sts.ToLower(name))
	var parameters col.ListLike[mod.ParameterLike]
	var declaration = mod.Declaration().MakeWithAttributes(
		comment,
		name+"ClassLike",
		parameters,
	)
	var constants col.ListLike[mod.ConstantLike]
	var functions col.ListLike[mod.FunctionLike]
	var class = mod.Class().MakeWithAttributes(
		declaration,
		constants,
		constructors,
		functions,
	)
	return class
}

func (v *generator_) generateClasses() col.ListLike[mod.ClassLike] {
	var classes = col.List[mod.ClassLike]().Make()
	if !v.classes_.IsEmpty() {
		v.classes_.SortValues()
		classes.AppendValues(v.classes_.GetValues(v.classes_.GetKeys()))
	}
	return classes
}

func (v *generator_) generateFunctionals() col.ListLike[mod.FunctionalLike] {
	var functionals col.ListLike[mod.FunctionalLike]
	return functionals
}

func (v *generator_) generateHeader(packageName string) mod.HeaderLike {
	var comment = v.generatePackageComment(packageName)
	var header = mod.Header().MakeWithAttributes(comment, packageName)
	return header
}

func (v *generator_) generateInstance(
	name string,
	attributes col.ListLike[mod.AttributeLike],
) mod.InstanceLike {
	var comment = instanceCommentTemplate_
	comment = sts.ReplaceAll(comment, "<ClassName>", name)
	comment = sts.ReplaceAll(comment, "<class-name>", sts.ToLower(name))
	var parameters col.ListLike[mod.ParameterLike]
	var declaration = mod.Declaration().MakeWithAttributes(
		comment,
		name+"Like",
		parameters,
	)
	var abstractions col.ListLike[mod.AbstractionLike]
	var methods col.ListLike[mod.MethodLike]
	var instance = mod.Instance().MakeWithAttributes(
		declaration,
		attributes,
		abstractions,
		methods,
	)
	return instance
}

func (v *generator_) generateInstances() col.ListLike[mod.InstanceLike] {
	var instances = col.List[mod.InstanceLike]().Make()
	if !v.instances_.IsEmpty() {
		v.instances_.SortValues()
		instances.AppendValues(v.instances_.GetValues(v.instances_.GetKeys()))
	}
	return instances
}

func (v *generator_) generateModel(directory string, model mod.ModelLike) {
	var validator = mod.Validator().Make()
	validator.ValidateModel(model)
	var formatter = mod.Formatter().Make()
	var source = formatter.FormatModel(model)
	var bytes = []byte(source)
	var err = osx.WriteFile(directory+"Package.go", bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) generateModules() col.ListLike[mod.ModuleLike] {
	var modules = col.List[mod.ModuleLike]().Make()
	if !v.modules_.IsEmpty() {
		v.modules_.SortValues()
		modules.AppendValues(v.modules_.GetValues(v.modules_.GetKeys()))
	}
	return modules
}

func (v *generator_) generateNotice(grammar GrammarLike) mod.NoticeLike {
	var header = grammar.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the grammar style comment delimiters.
	comment = Scanner().MatchToken(CommentToken, comment).GetValue(2)

	// Add the Go style comment delimiters.
	comment = "/*\n" + comment + "\n*/\n"

	var notice = mod.Notice().MakeWithComment(comment)
	return notice
}

func (v *generator_) generatePackageComment(
	packageName string,
) string {
	var comment = packageCommentTemplate_
	comment = sts.ReplaceAll(comment, "<packagename>", packageName)
	return comment
}

func (v *generator_) generateTypes() col.ListLike[mod.TypeLike] {
	var types col.ListLike[mod.TypeLike]
	return types
}

func (v *generator_) isLowercase(identifier string) bool {
	return uni.IsLower([]rune(identifier)[0])
}

func (v *generator_) isUppercase(identifier string) bool {
	return uni.IsUpper([]rune(identifier)[0])
}

func (v *generator_) makeList(attribute mod.AttributeLike) mod.AttributeLike {
	var identifier = attribute.GetIdentifier()
	identifier = v.makePlural(identifier)
	var abstraction = attribute.GetAbstraction()
	var arguments = col.List[mod.AbstractionLike]().Make()
	arguments.AppendValue(abstraction)
	abstraction = mod.Abstraction().MakeWithAttributes(
		mod.Prefix().MakeWithAttributes("col", mod.AliasPrefix),
		"ListLike",
		arguments,
	)
	var parameter mod.ParameterLike
	attribute = mod.Attribute().MakeWithAttributes(identifier, parameter, abstraction)
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
	constructor mod.ConstructorLike,
	attributes col.ListLike[mod.AttributeLike],
) {
	// Extract the attributes.
	attributes = col.List[mod.AttributeLike]().Make()
	var iterator = alternative.GetFactors().GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		var values = v.processFactor(name, factor)
		attributes.AppendValues(values)
	}
	v.consolidateLists(attributes)

	// Extract the constructor.
	var prefix mod.PrefixLike
	var arguments col.ListLike[mod.AbstractionLike]
	var abstraction = mod.Abstraction().MakeWithAttributes(
		prefix,
		name+"Like",
		arguments,
	)
	if !attributes.IsEmpty() {
		var identifier = "MakeWithAttributes"
		if attributes.GetSize() == 1 {
			identifier = attributes.GetValue(1).GetIdentifier()
			identifier = sts.TrimPrefix(identifier, "Get")
			identifier = "MakeWith" + identifier
		}
		var parameters = v.extractParameters(attributes)
		constructor = mod.Constructor().MakeWithAttributes(
			identifier,
			parameters,
			abstraction,
		)
	}

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
	constructors col.ListLike[mod.ConstructorLike],
	attributes col.ListLike[mod.AttributeLike],
) {
	constructors = col.List[mod.ConstructorLike]().Make()
	attributes = col.List[mod.AttributeLike]().Make()
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
	v.consolidateConstructors(constructors)
	v.consolidateAttributes(attributes)
	return constructors, attributes
}

func (v *generator_) processFactor(
	name string,
	factor FactorLike,
) (attributes col.ListLike[mod.AttributeLike]) {
	var isSequential bool
	var cardinality = factor.GetCardinality()
	if cardinality != nil {
		var constraint = cardinality.GetConstraint()
		if constraint.GetFirst() != "0" || constraint.GetLast() != "1" {
			isSequential = true
			var prefix = "col"
			var repository = `"github.com/craterdog/go-collection-framework/v3"`
			var module = mod.Module().MakeWithAttributes(
				prefix,
				repository,
			)
			// Must be sorted by the repository name NOT the prefix.
			v.modules_.SetValue(repository, module)
		}
	}
	var identifier string
	var abstraction mod.AbstractionLike
	var attribute mod.AttributeLike
	attributes = col.List[mod.AttributeLike]().Make()
	var predicate = factor.GetPredicate()
	var element = predicate.GetElement()
	var inversion = predicate.GetInversion()
	var precedence = predicate.GetPrecedence()
	switch {
	case element != nil:
		identifier = element.GetName()
		if len(identifier) > 0 {
			var prefix mod.PrefixLike
			var arguments col.ListLike[mod.AbstractionLike]
			if v.isUppercase(identifier) {
				abstraction = mod.Abstraction().MakeWithAttributes(
					prefix,
					identifier+"Like",
					arguments,
				)
				if isSequential {
					identifier = v.makePlural(identifier)
					var arguments = col.List[mod.AbstractionLike]().Make()
					arguments.AppendValue(abstraction)
					abstraction = mod.Abstraction().MakeWithAttributes(
						mod.Prefix().MakeWithAttributes("col", mod.AliasPrefix),
						"ListLike",
						arguments,
					)
				}
			} else {
				abstraction = mod.Abstraction().MakeWithAttributes(
					prefix,
					"string",
					arguments,
				)
			}
			var parameter mod.ParameterLike
			attribute = mod.Attribute().MakeWithAttributes(
				"Get"+v.makeUppercase(identifier),
				parameter,
				abstraction,
			)
			attributes.AppendValue(attribute)
		}
	case inversion != nil:
	case precedence != nil:
		var expression = precedence.GetExpression()
		var _, values = v.processExpression(name, expression)
		attributes.AppendValues(values)
	default:
		panic("Found an empty predicate.")
	}
	return attributes
}

func (v *generator_) processGrammar(grammar GrammarLike) mod.ModelLike {
	var iterator = grammar.GetDefinitions().GetIterator()
	var definition = iterator.GetNext()
	var packageName = v.extractName(definition)
	for iterator.HasNext() {
		definition = iterator.GetNext()
		v.processDefinition(definition)
	}
	var notice = v.generateNotice(grammar)
	var header = v.generateHeader(packageName)
	var imports = v.generateModules()
	var types = v.generateTypes()
	var functionals = v.generateFunctionals()
	var aspects = v.generateAspects()
	var classes = v.generateClasses()
	var instances = v.generateInstances()
	var model = mod.Model().MakeWithAttributes(
		notice,
		header,
		imports,
		types,
		functionals,
		aspects,
		classes,
		instances,
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
