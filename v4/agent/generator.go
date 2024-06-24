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
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	mod "github.com/craterdog/go-model-framework/v4"
	sts "strings"
	tim "time"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var generatorClass = &generatorClass_{
	// Initialize class constants.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
	// Define class constants.
}

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	return &generator_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	// Define instance attributes.
	class_     GeneratorClassLike
	lexigrams_ col.SetLike[string]
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
) ast.SyntaxLike {
	// Center and insert the copyright notice into the syntax template.
	copyright = v.expandCopyright(copyright)
	var source = sts.ReplaceAll(syntaxTemplate_, "<Copyright>", copyright)
	source = sts.ReplaceAll(source, "<NAME>", sts.ToUpper(name))
	source = sts.ReplaceAll(source, "<Name>", name)

	// Parse the syntax.
	var parser = Parser().Make()
	var syntax = parser.ParseSource(source)
	return syntax
}

func (v *generator_) GenerateAgent(
	module string,
	syntax ast.SyntaxLike,
) mod.ModelLike {
	// Create the agent class model template.
	v.processSyntax(syntax)
	var source = modelTemplate_ + agentTemplate_
	var notice = v.extractNotice(syntax)
	source = sts.ReplaceAll(source, "<Notice>", notice)
	source = sts.ReplaceAll(source, "<module>", module)
	source = sts.ReplaceAll(source, "<package>", "agent")
	var name = v.extractSyntaxName(syntax)
	source = sts.ReplaceAll(source, "<Name>", v.makeUppercase(name))
	source = sts.ReplaceAll(source, "<name>", v.makeLowercase(name))
	var parameter = v.makeLowercase(name)
	source = sts.ReplaceAll(source, "<parameter>", parameter)

	// Parse the agent class model template.
	var parser = mod.Parser()
	var model = parser.ParseSource(source)

	// Add additional definitions to the class model.
	v.addLexigrams(model)

	return model
}

func (v *generator_) GenerateAST(
	module string,
	syntax ast.SyntaxLike,
) mod.ModelLike {
	// Create the AST class model template.
	v.processSyntax(syntax)
	var source = modelTemplate_ + astTemplate_
	var notice = v.extractNotice(syntax)
	source = sts.ReplaceAll(source, "<Notice>", notice)
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

func (v *generator_) GenerateFormatter(
	module string,
	syntax ast.SyntaxLike,
	model mod.ModelLike,
) string {
	var source = formatterTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	source = sts.ReplaceAll(source, "<module>", module)
	var name = v.extractSyntaxName(syntax)
	source = sts.ReplaceAll(source, "<Name>", v.makeUppercase(name))
	source = sts.ReplaceAll(source, "<name>", v.makeLowercase(name))
	return source
}

func (v *generator_) GenerateParser(
	module string,
	syntax ast.SyntaxLike,
	model mod.ModelLike,
) string {
	var source = parserTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	source = sts.ReplaceAll(source, "<module>", module)
	var name = v.extractSyntaxName(syntax)
	source = sts.ReplaceAll(source, "<Name>", v.makeUppercase(name))
	source = sts.ReplaceAll(source, "<name>", v.makeLowercase(name))
	return source
}

func (v *generator_) GenerateScanner(
	module string,
	syntax ast.SyntaxLike,
	model mod.ModelLike,
) string {
	var source = scannerTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	source = sts.ReplaceAll(source, "<module>", module)
	return source
}

func (v *generator_) GenerateToken(
	module string,
	syntax ast.SyntaxLike,
	model mod.ModelLike,
) string {
	var source = tokenTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	source = sts.ReplaceAll(source, "<module>", module)
	return source
}

func (v *generator_) GenerateValidator(
	module string,
	syntax ast.SyntaxLike,
	model mod.ModelLike,
) string {
	var source = validatorTemplate_
	var notice = model.GetNotice().GetComment()
	source = sts.ReplaceAll(source, "<Notice>", notice)
	source = sts.ReplaceAll(source, "<module>", module)
	var name = v.extractSyntaxName(syntax)
	source = sts.ReplaceAll(source, "<Name>", v.makeUppercase(name))
	source = sts.ReplaceAll(source, "<name>", v.makeLowercase(name))
	return source
}

// Private

func (v *generator_) addClasses(model mod.ModelLike) {
	var classes = model.GetClasses()
	classes.RemoveAll() // Remove the dummy placeholder.
	classes.AppendValues(v.classes_.GetValues(v.classes_.GetKeys()))
}

func (v *generator_) addInstances(model mod.ModelLike) {
	var instances = model.GetInstances()
	instances.RemoveAll() // Remove the dummy placeholder.
	instances.AppendValues(v.instances_.GetValues(v.instances_.GetKeys()))
}

func (v *generator_) addLexigrams(model mod.ModelLike) {
	var types = model.GetTypes()
	var enumeration = types.GetValue(1).GetEnumeration()
	var identifiers = enumeration.GetIdentifiers()
	identifiers.AppendValues(v.lexigrams_)
}

func (v *generator_) addModules(model mod.ModelLike) {
	var modules = model.GetModules()
	modules.RemoveAll() // Remove the dummy placeholder.
	modules.AppendValues(v.modules_.GetValues(v.modules_.GetKeys()))
}

func (v *generator_) consolidateAttributes(
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

func (v *generator_) extractAttribute(name string) mod.AttributeLike {
	var abstraction mod.AbstractionLike
	switch {
	case !Scanner().MatchToken(UppercaseToken, name).IsEmpty():
		abstraction = mod.Abstraction(name + "Like")
	case !Scanner().MatchToken(LowercaseToken, name).IsEmpty():
		var tokenType = v.makeUppercase(name) + "Token"
		v.lexigrams_.AddValue(tokenType)
		abstraction = mod.Abstraction("string")
	default:
		var message = fmt.Sprintf(
			"Found an invalid attribute name: %q",
			name,
		)
		panic(message)
	}
	var attribute = mod.Attribute(
		"Get"+v.makeUppercase(name),
		abstraction,
	)
	return attribute
}

func (v *generator_) extractNotice(syntax ast.SyntaxLike) string {
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

func (v *generator_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetValue(1)
	var name = rule.GetUppercase()
	return name
}

func (v *generator_) generateClass(
	name string,
	constructor mod.ConstructorLike,
) mod.ClassLike {
	var comment = sts.ReplaceAll(classCommentTemplate_, "<Class>", name)
	comment = sts.ReplaceAll(comment, "<class>", sts.ToLower(name))
	var declaration = mod.Declaration(
		comment,
		name+"ClassLike",
	)
	var notation = cdc.Notation().Make()
	var constructors = col.List[mod.ConstructorLike](notation).Make()
	constructors.AppendValue(constructor)
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
	attributes.InsertValue(0, attribute)
	var instance = mod.Instance(
		declaration,
		attributes,
	)
	return instance
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

func (v *generator_) processExpression(
	name string,
	expression ast.ExpressionLike,
) (
	constructor mod.ConstructorLike,
	attributes col.ListLike[mod.AttributeLike],
) {
	// Extract the attributes.
	var notation = cdc.Notation().Make()
	attributes = col.List[mod.AttributeLike](notation).Make()
	switch actual := expression.GetAny().(type) {
	case ast.InlinedLike:
		v.processInlined(name, actual, attributes)
	case ast.MultilinedLike:
		v.processMultilined(name, actual, attributes)
	default:
		panic("Found an empty expression.")
	}

	// Create the constructor.
	var abstraction = mod.Abstraction(name + "Like")
	var identifier = "Make"
	var parameters = v.extractParameters(attributes)
	constructor = mod.Constructor(
		identifier,
		parameters,
		abstraction,
	)

	return constructor, attributes
}

func (v *generator_) processFactor(
	name string,
	factor ast.FactorLike,
	attributes col.ListLike[mod.AttributeLike],
) {
	var predicate = factor.GetPredicate()
	var actual = predicate.GetAny().(string)
	switch {
	case !Scanner().MatchToken(IntrinsicToken, actual).IsEmpty():
		// NOTE: We must check for intrinsics first and ignore them.
	case !Scanner().MatchToken(LiteralToken, actual).IsEmpty():
		// Ignore literals as well.
	default:
		var attribute = v.extractAttribute(actual)
		var cardinality = factor.GetCardinality()
		if cardinality != nil {
			switch actual := cardinality.GetAny().(type) {
			case ast.ConstrainedLike:
				attribute = v.makeList(attribute)
			case string:
				switch actual {
				case "?":
					// Don't make a list for zero or one.
				case "*", "+":
					attribute = v.makeList(attribute)
				}
			}
		}
		attributes.AppendValue(attribute)
	}
}

func (v *generator_) processInlined(
	name string,
	inlined ast.InlinedLike,
	attributes col.ListLike[mod.AttributeLike],
) {
	// Extract the attributes.
	var iterator = inlined.GetFactors().GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		v.processFactor(name, factor, attributes)
	}
	v.consolidateAttributes(attributes)
}

func (v *generator_) processLexigram(
	lexigram ast.LexigramLike,
) {
	// Ignore lexigram definitions for now.
}

func (v *generator_) processMultilined(
	name string,
	multilined ast.MultilinedLike,
	attributes col.ListLike[mod.AttributeLike],
) {
	var abstraction = mod.Abstraction("any")
	var attribute = mod.Attribute(
		"GetAny",
		abstraction,
	)
	attributes.AppendValue(attribute)
}

func (v *generator_) processRule(rule ast.RuleLike) {
	// Process the expression.
	var name = rule.GetUppercase()
	var expression = rule.GetExpression()
	var constructor, attributes = v.processExpression(name, expression)

	// Create the class interface.
	var class = v.generateClass(name, constructor)
	v.classes_.SetValue(name, class)

	// Create the instance interface.
	var instance = v.generateInstance(name, attributes)
	v.instances_.SetValue(name, instance)
}

func (v *generator_) processSyntax(syntax ast.SyntaxLike) {
	// Initialize the collections.
	var array = []string{
		"DelimiterToken",
		"EOFToken",
		"EOLToken",
		"SpaceToken",
	}
	var notation = cdc.Notation().Make()
	v.lexigrams_ = col.Set[string](notation).MakeFromArray(array)
	v.modules_ = col.Catalog[string, mod.ModuleLike](notation).Make()
	v.classes_ = col.Catalog[string, mod.ClassLike](notation).Make()
	v.instances_ = col.Catalog[string, mod.InstanceLike](notation).Make()

	// Process the syntax rule definitions.
	var rulesIterator = syntax.GetRules().GetIterator()
	for rulesIterator.HasNext() {
		var rule = rulesIterator.GetNext()
		v.processRule(rule)
	}

	// Process the syntax lexigram definitions.
	var lexigramIterator = syntax.GetLexigrams().GetIterator()
	for lexigramIterator.HasNext() {
		var lexigram = lexigramIterator.GetNext()
		v.processLexigram(lexigram)
	}
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
