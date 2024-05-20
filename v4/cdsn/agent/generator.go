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
	gcf "github.com/craterdog/go-collection-framework/v4"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	cds "github.com/craterdog/go-grammar-framework/v4/cdsn/ast"
	age "github.com/craterdog/go-model-framework/v4/gcmn/agent"
	gcm "github.com/craterdog/go-model-framework/v4/gcmn/ast"
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
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	class_     GeneratorClassLike
	tokens_    col.SetLike[string]
	modules_   col.CatalogLike[string, gcm.ModuleLike]
	classes_   col.CatalogLike[string, gcm.ClassLike]
	instances_ col.CatalogLike[string, gcm.InstanceLike]
}

// Attributes

func (v *generator_) GetClass() GeneratorClassLike {
	return v.class_
}

// Public

func (v *generator_) CreateSyntax(
	directory string,
	name string,
	copyright string,
) {
	// Create a new directory for the syntax.
	var syntaxDirectory = v.createDirectory(directory, sts.ToLower(name))

	// Center and insert the copyright notice into the syntax template.
	copyright = v.expandCopyright(copyright)
	var template = sts.ReplaceAll(syntaxTemplate_, "<Copyright>", copyright)
	template = sts.ReplaceAll(template, "<NAME>", sts.ToUpper(name))
	template = sts.ReplaceAll(template, "<Name>", name)

	// Save the new syntax template into the directory.
	var syntaxFile = syntaxDirectory + "Syntax.cdsn"
	v.outputFile(syntaxFile, template[1:]) // Remove leading "\n".
}

func (v *generator_) GenerateAST(directory string, name string) {
	// Create a new directory for the AST.
	var syntaxDirectory = v.createDirectory(directory, name)
	var astDirectory = v.createDirectory(syntaxDirectory, "ast")

	// Parse the syntax file.
	var syntax = v.parseSyntax(syntaxDirectory)
	if syntax == nil {
		return
	}

	// Create the ast/Package.go file model template.
	v.processSyntax(syntax)
	var source = modelTemplate_ + astTemplate_
	var notice = v.extractNotice(syntax)
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = v.extractModule(directory, name)
	source = sts.ReplaceAll(source, "<module>", module)
	source = sts.ReplaceAll(source, "<package>", "ast")
	var parser = age.Parser().Make()
	var model = parser.ParseSource(source)

	// Add additional definitions to the class model.
	v.addModules(model)
	v.addClasses(model)
	v.addInstances(model)

	// Generate the class model for the syntax directory.
	v.formatModel(astDirectory, model)
	v.generatePackage(astDirectory)
}

func (v *generator_) GenerateAgents(directory string, name string) {
	// Create a new directory for the agents.
	var syntaxDirectory = v.createDirectory(directory, name)
	var agentDirectory = v.createDirectory(syntaxDirectory, "agent")

	// Parse the syntax file.
	var syntax = v.parseSyntax(syntaxDirectory)
	if syntax == nil {
		return
	}

	// Create the agent/Package.go file model template.
	v.processSyntax(syntax)
	var source = modelTemplate_ + agentTemplate_
	var notice = v.extractNotice(syntax)
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = v.extractModule(directory, name)
	source = sts.ReplaceAll(source, "<module>", module)
	source = sts.ReplaceAll(source, "<package>", "agent")
	var class = v.extractClassName(syntax)
	source = sts.ReplaceAll(source, "<Class>", class)
	var parameter = v.makeLowercase(class)
	source = sts.ReplaceAll(source, "<parameter>", parameter)
	var parser = age.Parser().Make()
	var model = parser.ParseSource(source)

	// Add additional definitions to the class model.
	v.addTokens(model)

	// Generate the Package.go file for the agent directory.
	v.formatModel(agentDirectory, model)
	v.generateToken(directory, name, model)
	v.generateScanner(directory, name, model)
	v.generateParser(directory, name, model, class)
	v.generateValidator(directory, name, model, class)
	v.generateFormatter(directory, name, model, class)
}

// Private

func (v *generator_) addClasses(model gcm.ModelLike) {
	var classes = model.GetClasses()
	classes.RemoveAll() // Remove the dummy placeholder.
	classes.AppendValues(v.classes_.GetValues(v.classes_.GetKeys()))
}

func (v *generator_) addModules(model gcm.ModelLike) {
	var modules = model.GetModules()
	modules.RemoveAll() // Remove the dummy placeholder.
	modules.AppendValues(v.modules_.GetValues(v.modules_.GetKeys()))
}

func (v *generator_) addInstances(model gcm.ModelLike) {
	var instances = model.GetInstances()
	instances.RemoveAll() // Remove the dummy placeholder.
	instances.AppendValues(v.instances_.GetValues(v.instances_.GetKeys()))
}

func (v *generator_) addTokens(model gcm.ModelLike) {
	var types = model.GetTypes()
	var enumeration = types.GetValue(1).GetEnumeration()
	var identifiers = enumeration.GetIdentifiers()
	identifiers.AppendValues(v.tokens_)
}

func (v *generator_) consolidateAttributes(
	attributes col.ListLike[gcm.AttributeLike],
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
	constructors col.ListLike[gcm.ConstructorLike],
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
	attributes col.ListLike[gcm.AttributeLike],
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

func (v *generator_) createDirectory(directory string, name string) string {
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	directory += name + "/"
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
	return directory
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
	var alternatives = gcf.List[cds.AlternativeLike]()
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

func (v *generator_) extractModule(directory string, name string) string {
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var modFile = directory + "go.mod"
	var bytes, err = osx.ReadFile(modFile)
	if err != nil {
		panic(err)
	}
	var line = sts.Split(string(bytes), "\n")[0]
	var module = sts.TrimPrefix(line, "module ") + "/" + name
	return module
}

func (v *generator_) extractNotice(syntax cds.SyntaxLike) string {
	var header = syntax.GetHeaders().GetValue(1)
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = Scanner().MatchToken(CommentToken, comment).GetValue(2)

	// Add the Go style comment delimiters.
	notice = "/*\n" + notice + "\n*/\n"

	return notice
}

func (v *generator_) extractParameters(
	attributes col.ListLike[gcm.AttributeLike],
) col.ListLike[gcm.ParameterLike] {
	var parameters = gcf.List[gcm.ParameterLike]()
	var iterator = attributes.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var identifier = sts.TrimPrefix(attribute.GetIdentifier(), "Get")
		var abstraction = attribute.GetAbstraction()
		var parameter = gcm.Parameter().MakeWithAttributes(
			v.makeLowercase(identifier),
			abstraction,
		)
		parameters.AppendValue(parameter)
	}
	return parameters
}

func (v *generator_) formatModel(
	directory string,
	model gcm.ModelLike,
) {
	var validator = age.Validator().Make()
	validator.ValidateModel(model)
	var formatter = age.Formatter().Make()
	var source = formatter.FormatModel(model)
	v.outputFile(directory+"Package.go", source)
}

func (v *generator_) generateClass(
	name string,
	constructors col.ListLike[gcm.ConstructorLike],
) gcm.ClassLike {
	var comment = classCommentTemplate_[1:] // Strip off leading newline.
	comment = sts.ReplaceAll(comment, "<Class>", name)
	comment = sts.ReplaceAll(comment, "<class>", sts.ToLower(name))
	var parameters col.ListLike[gcm.ParameterLike]
	var declaration = gcm.Declaration().MakeWithAttributes(
		comment,
		name+"ClassLike",
		parameters,
	)
	var constants col.ListLike[gcm.ConstantLike]
	var functions col.ListLike[gcm.FunctionLike]
	var class = gcm.Class().MakeWithAttributes(
		declaration,
		constants,
		constructors,
		functions,
	)
	return class
}

func (v *generator_) generateFormatter(
	directory string,
	name string,
	model gcm.ModelLike,
	class string,
) {
	var source = formatterTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = v.extractModule(directory, name)
	source = sts.ReplaceAll(source, "<module>", module)
	source = sts.ReplaceAll(source, "<Class>", v.makeUppercase(class))
	source = sts.ReplaceAll(source, "<class>", v.makeLowercase(class))
	var file = directory + name + "/agent/formatter.go"
	v.outputFile(file, source)
}

func (v *generator_) generateInstance(
	name string,
	attributes col.ListLike[gcm.AttributeLike],
) gcm.InstanceLike {
	var comment = instanceCommentTemplate_[1:] // Strip off leading newline.
	comment = sts.ReplaceAll(comment, "<Class>", name)
	comment = sts.ReplaceAll(comment, "<class>", sts.ToLower(name))
	var parameters col.ListLike[gcm.ParameterLike]
	var declaration = gcm.Declaration().MakeWithAttributes(
		comment,
		name+"Like",
		parameters,
	)
	var parameter gcm.ParameterLike
	var prefix gcm.PrefixLike
	var arguments col.ListLike[gcm.AbstractionLike]
	var abstraction = gcm.Abstraction().MakeWithAttributes(
		prefix,
		name+"ClassLike",
		arguments,
	)
	var attribute = gcm.Attribute().MakeWithAttributes(
		"GetClass",
		parameter,
		abstraction,
	)
	attributes.InsertValue(0, attribute)
	var abstractions col.ListLike[gcm.AbstractionLike]
	var methods col.ListLike[gcm.MethodLike]
	var instance = gcm.Instance().MakeWithAttributes(
		declaration,
		attributes,
		abstractions,
		methods,
	)
	return instance
}

func (v *generator_) generatePackage(
	directory string,
) {
	var generator = age.Generator().Make()
	generator.GeneratePackage(directory)
}

func (v *generator_) generateParser(
	directory string,
	name string,
	model gcm.ModelLike,
	class string,
) {
	var source = parserTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = v.extractModule(directory, name)
	source = sts.ReplaceAll(source, "<module>", module)
	var packageName = model.GetHeader().GetIdentifier()
	source = sts.ReplaceAll(source, "<Package>", v.makeUppercase(packageName))
	source = sts.ReplaceAll(source, "<Class>", v.makeUppercase(class))
	source = sts.ReplaceAll(source, "<class>", v.makeLowercase(class))
	var file = directory + name + "/agent/parser.go"
	v.outputFile(file, source)
}

func (v *generator_) generateScanner(
	directory string,
	name string,
	model gcm.ModelLike,
) {
	var source = scannerTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = v.extractModule(directory, name)
	source = sts.ReplaceAll(source, "<module>", module)
	var file = directory + name + "/agent/scanner.go"
	v.outputFile(file, source)
}

func (v *generator_) generateToken(
	directory string,
	name string,
	model gcm.ModelLike,
) {
	var source = tokenTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = v.extractModule(directory, name)
	source = sts.ReplaceAll(source, "<module>", module)
	var file = directory + name + "/agent/token.go"
	v.outputFile(file, source)
}

func (v *generator_) generateValidator(
	directory string,
	name string,
	model gcm.ModelLike,
	class string,
) {
	var source = validatorTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	var module = v.extractModule(directory, name)
	source = sts.ReplaceAll(source, "<module>", module)
	source = sts.ReplaceAll(source, "<Class>", v.makeUppercase(class))
	source = sts.ReplaceAll(source, "<class>", v.makeLowercase(class))
	var file = directory + name + "/agent/validator.go"
	v.outputFile(file, source)
}

func (v *generator_) isLowercase(identifier string) bool {
	return uni.IsLower([]rune(identifier)[0])
}

func (v *generator_) isUppercase(identifier string) bool {
	return uni.IsUpper([]rune(identifier)[0])
}

func (v *generator_) makeList(attribute gcm.AttributeLike) gcm.AttributeLike {
	var prefix = "col"
	var path = `"github.com/craterdog/go-collection-framework/v4/collection"`
	var module = gcm.Module().MakeWithAttributes(
		prefix,
		path,
	)
	v.modules_.SetValue(path, module)
	var identifier = attribute.GetIdentifier()
	identifier = v.makePlural(identifier)
	var abstraction = attribute.GetAbstraction()
	var arguments = gcf.List[gcm.AbstractionLike]()
	arguments.AppendValue(abstraction)
	abstraction = gcm.Abstraction().MakeWithAttributes(
		gcm.Prefix().MakeWithAttributes(prefix, gcm.AliasPrefix),
		"ListLike",
		arguments,
	)
	var parameter gcm.ParameterLike
	attribute = gcm.Attribute().MakeWithAttributes(identifier, parameter, abstraction)
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

func (v *generator_) outputFile(file, source string) {
	var _, err = osx.ReadFile(file)
	if err == nil {
		// Don't overwrite an existing class file.
		fmt.Printf(
			"The file %q already exists, leaving it alone.\n",
			file,
		)
		return
	}
	err = osx.WriteFile(file, []byte(source), 0644)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) parseSyntax(directory string) cds.SyntaxLike {
	var syntaxFile = directory + "Syntax.cdsn"
	var bytes, err = osx.ReadFile(syntaxFile)
	if err != nil {
		var message = fmt.Sprintf(
			"The specified directory is missing a syntax file: %v",
			syntaxFile,
		)
		panic(message)
	}
	var source = string(bytes)
	var parser = Parser().Make()
	var syntax = parser.ParseSource(source)
	var validator = Validator().Make()
	validator.ValidateSyntax(syntax)
	return syntax
}

func (v *generator_) processAlternative(
	name string,
	alternative cds.AlternativeLike,
) (
	constructor gcm.ConstructorLike,
	attributes col.ListLike[gcm.AttributeLike],
) {
	// Extract the attributes.
	attributes = gcf.List[gcm.AttributeLike]()
	var iterator = alternative.GetFactors().GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		var values = v.processFactor(name, factor)
		attributes.AppendValues(values)
	}
	v.consolidateLists(attributes)

	// Extract the constructor.
	var prefix gcm.PrefixLike
	var arguments col.ListLike[gcm.AbstractionLike]
	var abstraction = gcm.Abstraction().MakeWithAttributes(
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
		constructor = gcm.Constructor().MakeWithAttributes(
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
	constructors col.ListLike[gcm.ConstructorLike],
	attributes col.ListLike[gcm.AttributeLike],
) {
	// Process the expression alternatives.
	constructors = gcf.List[gcm.ConstructorLike]()
	attributes = gcf.List[gcm.AttributeLike]()
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
		var prefix gcm.PrefixLike
		var arguments col.ListLike[gcm.AbstractionLike]
		var abstraction = gcm.Abstraction().MakeWithAttributes(
			prefix,
			name+"Like",
			arguments,
		)
		var parameters col.ListLike[gcm.ParameterLike]
		var constructor = gcm.Constructor().MakeWithAttributes(
			"Make",
			parameters,
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
) (attributes col.ListLike[gcm.AttributeLike]) {
	var isSequential bool
	var cardinality = factor.GetCardinality()
	if cardinality != nil {
		var constraint = cardinality.GetConstraint()
		isSequential = constraint.GetFirst() != "0" || constraint.GetLast() != "1"
	}
	var identifier string
	var abstraction gcm.AbstractionLike
	var attribute gcm.AttributeLike
	attributes = gcf.List[gcm.AttributeLike]()
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
			var prefix gcm.PrefixLike
			var arguments col.ListLike[gcm.AbstractionLike]
			if v.isUppercase(identifier) {
				abstraction = gcm.Abstraction().MakeWithAttributes(
					prefix,
					identifier+"Like",
					arguments,
				)
			} else {
				var tokenType = v.makeUppercase(identifier) + "Token"
				v.tokens_.AddValue(tokenType)
				abstraction = gcm.Abstraction().MakeWithAttributes(
					prefix,
					"string",
					arguments,
				)
			}
			var parameter gcm.ParameterLike
			attribute = gcm.Attribute().MakeWithAttributes(
				"Get"+v.makeUppercase(identifier),
				parameter,
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
	v.tokens_ = gcf.Set[string](array)
	v.modules_ = gcf.Catalog[string, gcm.ModuleLike]()
	v.classes_ = gcf.Catalog[string, gcm.ClassLike]()
	v.instances_ = gcf.Catalog[string, gcm.InstanceLike]()

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
