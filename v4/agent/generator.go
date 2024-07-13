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
	fwk "github.com/craterdog/go-collection-framework/v4"
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
	// Initialize the class constants.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	return &generator_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	// Define the instance attributes.
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
	var tokenTypes = v.extractTokenTypes(syntax)
	source = sts.ReplaceAll(source, "<TokenTypes>", tokenTypes)

	// Parse the agent generated class model.
	var parser = mod.Parser()
	var model = parser.ParseSource(source)
	return model
}

func (v *generator_) GenerateAST(
	module string,
	syntax ast.SyntaxLike,
) mod.ModelLike {
	// Create the AST class model template.
	var source = modelTemplate_ + astTemplate_
	source = sts.ReplaceAll(source, "<Notice>", v.extractNotice(syntax))
	source = sts.ReplaceAll(source, "<module>", module)
	source = sts.ReplaceAll(source, "<package>", "ast")

	// Parse the AST class model template.
	var parser = mod.Parser()
	var model = parser.ParseSource(source)

	// Add additional definitions to the AST class model.
	v.processSyntax(syntax)
	var notice = model.GetNotice()
	var header = model.GetHeader()
	var modules = mod.Modules(v.modules_.GetValues(v.modules_.GetKeys()))
	var imports = mod.Imports(modules)
	var types = model.GetOptionalTypes()
	var functionals = model.GetOptionalFunctionals()
	var classes = mod.Classes(v.classes_.GetValues(v.classes_.GetKeys()))
	var instances = mod.Instances(v.instances_.GetValues(v.instances_.GetKeys()))
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

func (v *generator_) consolidateAttributes(
	attributes col.ListLike[mod.AttributeLike],
) {
	// Compare each attribute and make lists out of duplicates.
	for i := 1; i <= attributes.GetSize(); i++ {
		var attribute = attributes.GetValue(i)
		var first = attribute.GetName()
		for j := i + 1; j <= attributes.GetSize(); {
			var second = attributes.GetValue(j).GetName()
			switch {
			case first == second:
				attribute = attributes.GetValue(i)
				attribute = v.pluralizeAttribute(attribute)
				attributes.SetValue(i, attribute)
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
	var attributeType mod.AbstractionLike
	switch {
	case !Scanner().MatchToken(UppercaseToken, name).IsEmpty():
		// The attribute type is the (non-generic) abstract instance type.
		attributeType = mod.Abstraction(name + "Like")
	case !Scanner().MatchToken(LowercaseToken, name).IsEmpty():
		// The attribute type is simply the Go primitive "string" type.
		var tokenType = v.makeUppercase(name) + "Token"
		v.lexigrams_.AddValue(tokenType)
		attributeType = mod.Abstraction("string")
	default:
		var message = fmt.Sprintf(
			"Found an invalid attribute name: %q",
			name,
		)
		panic(message)
	}
	var attributeName = "Get" + v.makeUppercase(name)
	var attribute = mod.Attribute(attributeName, attributeType)
	return attribute
}

func (v *generator_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = Scanner().MatchToken(CommentToken, comment).GetValue(2)

	// Add the Go style comment delimiters.
	notice = "/*\n" + notice + "\n*/"

	return notice
}

func (v *generator_) extractParameters(
	attributes col.ListLike[mod.AttributeLike],
) col.Sequential[mod.ParameterLike] {
	var notation = fwk.CDCN()
	var parameters = col.List[mod.ParameterLike](notation).Make()
	var iterator = attributes.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var name = sts.TrimPrefix(attribute.GetName(), "Get")
		var abstraction = attribute.GetOptionalAbstraction()
		if fwk.IsUndefined(abstraction) {
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

func (v *generator_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *generator_) extractTokenTypes(syntax ast.SyntaxLike) string {
	var tokenTypes = "ErrorToken TokenType = iota"
	var iterator = v.lexigrams_.GetIterator()
	for iterator.HasNext() {
		var tokenType = iterator.GetNext()
		tokenTypes += "\n\t" + tokenType
	}
	return tokenTypes
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
	var notation = fwk.CDCN()
	var list = col.List[mod.ConstructorLike](notation).Make()
	list.AppendValue(constructor)
	var constructors = mod.Constructors(list)
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
		mod.Attributes(attributes),
	)
	return instance
}

func (v *generator_) pluralizeAttribute(
	attribute mod.AttributeLike,
) mod.AttributeLike {
	// Add the collections module to the catalog of imported modules.
	var alias = "col"
	var path = `"github.com/craterdog/go-collection-framework/v4/collection"`
	var module = mod.Module(
		alias,
		path,
	)
	v.modules_.SetValue(path, module)

	// Extract the name and attribute type from the attribute.
	var name = v.makePlural(attribute.GetName())
	var attributeType = attribute.GetOptionalAbstraction() // Not optional here.

	// Create the generic arguments list for the pluralized attribute.
	var argument = mod.Argument(attributeType)
	var notation = fwk.CDCN()
	var additionalArguments = col.List[mod.AdditionalArgumentLike](notation).Make()
	var arguments = mod.Arguments(argument, additionalArguments)
	var genericArguments = mod.GenericArguments(arguments)

	// Create the result type for the pluralized attribute.
	attributeType = mod.Abstraction(
		alias,
		"Sequential",
		genericArguments,
	)
	attribute = mod.Attribute(name, attributeType)
	return attribute
}

func (v *generator_) makeLowercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	name = string(runes)
	if reserved_[name] {
		name += "_"
	}
	return name
}

func (v *generator_) makePlural(name string) string {
	if sts.HasSuffix(name, "s") {
		name += "es"
	} else {
		name += "s"
	}
	return name
}

func (v *generator_) makeUppercase(name string) string {
	runes := []rune(name)
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
	var notation = fwk.CDCN()
	attributes = col.List[mod.AttributeLike](notation).Make()
	switch actual := expression.GetAny().(type) {
	case ast.InlinedLike:
		v.processInlined(actual, attributes)
	case ast.MultilinedLike:
		v.processMultilined(actual, attributes)
	default:
		panic("Found an empty expression.")
	}

	// Create the constructor.
	var abstraction = mod.Abstraction(name + "Like")
	name = "Make"
	var parameters = v.extractParameters(attributes)
	var iterator = parameters.GetIterator()
	var parameter = iterator.GetNext()
	var additionalParameters = col.List[mod.AdditionalParameterLike](notation).Make()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var additionalParameter = mod.AdditionalParameter(parameter)
		additionalParameters.AppendValue(additionalParameter)
	}
	constructor = mod.Constructor(
		name,
		mod.Parameters(
			parameter,
			additionalParameters.(col.Sequential[mod.AdditionalParameterLike]),
		),
		abstraction,
	)

	return constructor, attributes
}

func (v *generator_) processCardinality(
	attribute mod.AttributeLike,
	cardinality ast.CardinalityLike,
) mod.AttributeLike {
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		attribute = v.pluralizeAttribute(attribute)
	case string:
		switch actual {
		case "?":
			// This attribute is optional, not plural.
		case "*", "+":
			// Turn the attribute into a sequence of that type attribute.
			attribute = v.pluralizeAttribute(attribute)
		}
	}
	return attribute
}

func (v *generator_) processPredicate(
	predicate ast.PredicateLike,
) (
	attribute mod.AttributeLike,
) {
	var actual = predicate.GetAny().(string)
	switch {
	case !Scanner().MatchToken(IntrinsicToken, actual).IsEmpty():
		// NOTE: We must check for intrinsics first and ignore them.
	case !Scanner().MatchToken(LiteralToken, actual).IsEmpty():
		// Ignore literals as well.
	default:
		// We know it is a rule or lexigram name which corresponds to an attribute
		// with a (non-generic) instance type, or a Go primitive "string" type
		// respectively.
		attribute = v.extractAttribute(actual)
	}
	return attribute
}

func (v *generator_) processFactor(
	factor ast.FactorLike,
	attributes col.ListLike[mod.AttributeLike],
) {
	// Attempt to extract the attribute definitions from the predicate string.
	var predicate = factor.GetPredicate()
	var attribute = v.processPredicate(predicate)
	if fwk.IsUndefined(attribute) {
		// The predicate does not correspond to an attribute.
		return
	}

	// Take into account any cardinality of the predicate.
	var cardinality = factor.GetOptionalCardinality()
	if fwk.IsDefined(cardinality) {
		// The attribute type may need to be "pluralized".
		attribute = v.processCardinality(attribute, cardinality)
	}

	// Add the attribute definition to our list.
	attributes.AppendValue(attribute)
}

func (v *generator_) processInlined(
	inlined ast.InlinedLike,
	attributes col.ListLike[mod.AttributeLike],
) {
	// Extract the attributes.
	var iterator = inlined.GetFactors().GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		v.processFactor(factor, attributes)
	}
	v.consolidateAttributes(attributes)
}

func (v *generator_) processLexigram(
	lexigram ast.LexigramLike,
) {
	// Ignore lexigram definitions for now.
}

func (v *generator_) processMultilined(
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
	var notation = fwk.CDCN()
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
