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
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
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
	tokens_    abs.SetLike[string]
	regexps_   abs.CatalogLike[string, string]
	modules_   abs.CatalogLike[string, mod.ModuleLike]
	classes_   abs.CatalogLike[string, mod.ClassLike]
	instances_ abs.CatalogLike[string, mod.InstanceLike]
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
	var template = v.generateSyntaxTemplate()
	var source = v.populateSyntaxTemplate(template, name, copyright)
	var parser = gra.Parser().Make()
	var syntax = parser.ParseSource(source)
	return syntax
}

func (v *generator_) GenerateGrammar(
	module string,
	syntax ast.SyntaxLike,
) mod.ModelLike {
	v.analyzeSyntax(syntax)
	var template = v.generateModelTemplate("grammar", syntax)
	var source = v.populateModelTemplate(template, module, syntax)
	var parser = mod.Parser()
	var model = parser.ParseSource(source)
	return model
}

func (v *generator_) GenerateAST(
	module string,
	syntax ast.SyntaxLike,
) mod.ModelLike {
	v.analyzeSyntax(syntax)
	var template = v.generateModelTemplate("ast", syntax)
	var source = v.populateModelTemplate(template, module, syntax)
	var parser = mod.Parser()
	var model = parser.ParseSource(source)
	model = v.augmentAstModel(model)
	return model
}

func (v *generator_) GenerateFormatter(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzeSyntax(syntax)
	var template = v.generateClassTemplate("formatter", syntax)
	implementation = v.populateClassTemplate(template, module, syntax)
	return implementation
}

func (v *generator_) GenerateParser(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzeSyntax(syntax)
	var template = v.generateClassTemplate("parser", syntax)
	implementation = v.populateClassTemplate(template, module, syntax)
	return implementation
}

func (v *generator_) GenerateScanner(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzeSyntax(syntax)
	var template = v.generateClassTemplate("scanner", syntax)
	implementation = v.populateScannerTemplate(template, syntax)
	return implementation
}

func (v *generator_) GenerateToken(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzeSyntax(syntax)
	var template = v.generateClassTemplate("token", syntax)
	implementation = v.populateClassTemplate(template, module, syntax)
	return implementation
}

func (v *generator_) GenerateValidator(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzeSyntax(syntax)
	var template = v.generateClassTemplate("validator", syntax)
	implementation = v.populateClassTemplate(template, module, syntax)
	return implementation
}

// Private

func (v *generator_) analyzeSyntax(syntax ast.SyntaxLike) {
	// Define the regular expressions for each intrinsic.
	var map_ = map[string]string{
		"any":     ` ".|" + eol_`,
		"base16":  ` "[0-9a-f]"`,
		"control": ` "\\p{Cc}"`,
		"digit":   ` "\\p{Nd}"`,
		"eof":     ` "\\z"`,
		"eol":     ` "\\n"`,
		"escape":  ` "\\\\(?:(?:" + unicode_ + ")|[abfnrtv'\"\\\\])"`,
		"lower":   ` "\\p{Ll}"`,
		"space":   ` "[ \\t]+"`,
		"unicode": ` "x" + base16_ + "{2}|u" + base16_ + "{4}|U" + base16_ + "{8}"`,
		"upper":   ` "\\p{Lu}"`,
	}

	// Process the syntax rule definitions.
	v.tokens_ = col.Set[string]([]string{"delimiter", "eol", "eof", "space"})
	v.modules_ = col.Catalog[string, mod.ModuleLike]()
	v.classes_ = col.Catalog[string, mod.ClassLike]()
	v.instances_ = col.Catalog[string, mod.InstanceLike]()
	var rulesIterator = syntax.GetRules().GetIterator()
	for rulesIterator.HasNext() {
		var rule = rulesIterator.GetNext()
		v.processRule(rule)
	}
	v.modules_.SortValues()
	v.classes_.SortValues()
	v.instances_.SortValues()

	// Process the syntax expression definitions.
	v.regexps_ = col.Catalog[string, string](map_)
	var expressionIterator = syntax.GetExpressions().GetIterator()
	for expressionIterator.HasNext() {
		var expression = expressionIterator.GetNext()
		v.processExpression(expression)
	}
	v.regexps_.SortValues()
}

func (v *generator_) augmentAstModel(model mod.ModelLike) mod.ModelLike {
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

func (v *generator_) consolidateAttributes(
	attributes abs.ListLike[mod.AttributeLike],
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

func (v *generator_) escapeString(string_ string) string {
	var escaped string
	for _, character := range string_ {
		switch character {
		case '"', '\\':
			escaped += "\\"
		case '.', '|', '^', '$', '+', '*', '?',
			'(', ')', '[', ']', '{', '}':
			escaped += "\\\\"
		}
		escaped += string(character)
	}
	return escaped
}

func (v *generator_) expandCopyright(copyright string) string {
	var limit = 78
	var length = len(copyright)
	if length > limit {
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
	var padding = (limit - length) / 2
	for range padding {
		copyright = " " + copyright + " "
	}
	if len(copyright) < limit {
		copyright = " " + copyright
	}
	copyright = "." + copyright + "."
	return copyright
}

func (v *generator_) extractAttribute(name string) mod.AttributeLike {
	var attributeType mod.AbstractionLike
	switch {
	case !gra.Scanner().MatchToken(gra.UppercaseToken, name).IsEmpty():
		// The attribute type is the (non-generic) abstract instance type.
		attributeType = mod.Abstraction(name + "Like")
	case !gra.Scanner().MatchToken(gra.LowercaseToken, name).IsEmpty():
		// The attribute type is simply the Go intrinsic "string" type.
		attributeType = mod.Abstraction("string")
		v.tokens_.AddValue(name)
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

func (v *generator_) extractExpressions() string {
	var expressions = `error_ = "x^"`
	var iterator = v.regexps_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var name = association.GetKey()
		var regexp = association.GetValue()
		expressions += "\n\t" + name + "_ = " + regexp
	}
	return expressions
}

func (v *generator_) extractFoundCases() string {
	var foundCases = "case v.foundToken(ErrorToken):"
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var tokenType = v.makeUppercase(tokenName) + "Token"
		foundCases += "\n\t\tcase v.foundToken(" + tokenType + "):"
	}
	return foundCases
}

func (v *generator_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = gra.Scanner().MatchToken(gra.CommentToken, comment).GetValue(2)
	notice = "\n" + notice + "\n"

	return notice
}

func (v *generator_) extractParameters(
	attributes abs.ListLike[mod.AttributeLike],
) abs.Sequential[mod.ParameterLike] {
	var parameters = col.List[mod.ParameterLike]()
	var iterator = attributes.GetIterator()
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

func (v *generator_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *generator_) extractTokenMatchers() string {
	var tokenMatchers = `ErrorToken: reg.MustCompile("x^"),`
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var tokenType = v.makeUppercase(tokenName) + "Token"
		tokenMatchers += "\n\t\t" + tokenType +
			`: reg.MustCompile("^(?:" + ` + tokenName + `_ + ")"),`
	}
	return tokenMatchers
}

func (v *generator_) extractTokenNames() string {
	var tokenNames = `ErrorToken: "error",`
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var tokenType = v.makeUppercase(tokenName) + "Token"
		tokenNames += "\n\t\t" + tokenType + `: "` + tokenName + `",`
	}
	return tokenNames
}

func (v *generator_) extractTokenTypes() string {
	var tokenTypes = "ErrorToken TokenType = iota"
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var name = iterator.GetNext()
		var tokenType = v.makeUppercase(name) + "Token"
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
	var list = col.List[mod.ConstructorLike]()
	list.AppendValue(constructor)
	var constructors = mod.Constructors(list)
	var class = mod.Class(
		declaration,
		constructors,
	)
	return class
}

func (v *generator_) generateClassTemplate(
	class string,
	syntax ast.SyntaxLike,
) (
	template string,
) {
	template = templates_[class]["notice"]
	template += templates_[class]["header"]
	template += templates_[class]["imports"]
	template += templates_[class]["access"]
	template += templates_[class]["class"]
	template += templates_[class]["instance"]
	return template
}

func (v *generator_) generateInstance(
	name string,
	attributes abs.ListLike[mod.AttributeLike],
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

/*
func (v *generator_) generateMethod(
	class string,
	rule ast.RuleLike,
) (
	implementation string,
) {
	var definition = rule.GetDefinition()
	switch actual := definition.GetAny().(type) {
	case ast.InlinedLike:
		var iterator = actual.GetFactors().GetIterator()
		for iterator.HasNext() {
		}
	case ast.MultilinedLike:
		var iterator = actual.GetLines().GetIterator()
		for iterator.HasNext() {
		}
	default:
		panic("Found an empty definition.")
	}
	return implementation
}

func (v *generator_) generateMethods(
	class string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	var iterator = syntax.GetRules().GetIterator()
	for iterator.HasNext() {
		var rule = iterator.GetNext()
		implementation += v.generateMethod(class, rule)
	}
	return implementation
}
*/

func (v *generator_) generateModelTemplate(
	model string,
	syntax ast.SyntaxLike,
) (
	template string,
) {
	template = templates_[model]["notice"]
	template += templates_[model]["header"]
	template += templates_[model]["imports"]
	template += templates_[model]["types"]
	template += templates_[model]["functionals"]
	template += templates_[model]["classes"]
	template += templates_[model]["instances"]
	template += templates_[model]["aspects"]
	return template
}

func (v *generator_) generateSyntaxTemplate() (
	template string,
) {
	template = templates_["syntax"]["notice"]
	template += templates_["syntax"]["header"]
	template += templates_["syntax"]["rules"]
	template += templates_["syntax"]["expressions"]
	return template
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

func (v *generator_) makeUppercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

func (v *generator_) optionalizeAttribute(
	attribute mod.AttributeLike,
) mod.AttributeLike {
	var name = attribute.GetName()
	name = "GetOptional" + sts.TrimPrefix(name, "Get")
	var attributeType = attribute.GetOptionalAbstraction()
	attribute = mod.Attribute(name, attributeType)
	return attribute
}

func (v *generator_) pluralizeAttribute(
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

func (v *generator_) populateClassTemplate(
	template string,
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	var name = v.extractSyntaxName(syntax)
	var uppercase = v.makeUppercase(name)
	var lowercase = v.makeLowercase(name)
	var notice = v.extractNotice(syntax)
	implementation = sts.ReplaceAll(template, "<Notice>", notice)
	implementation = sts.ReplaceAll(implementation, "<module>", module)
	implementation = sts.ReplaceAll(implementation, "<Name>", uppercase)
	implementation = sts.ReplaceAll(implementation, "<name>", lowercase)
	return implementation
}

func (v *generator_) populateModelTemplate(
	template string,
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	var name = v.extractSyntaxName(syntax)
	implementation = sts.ReplaceAll(template, "<module>", module)
	var notice = v.extractNotice(syntax)
	implementation = sts.ReplaceAll(implementation, "<Notice>", notice)
	var uppercase = v.makeUppercase(name)
	implementation = sts.ReplaceAll(implementation, "<Name>", uppercase)
	var lowercase = v.makeLowercase(name)
	implementation = sts.ReplaceAll(implementation, "<name>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<parameter>", lowercase)
	var tokenTypes = v.extractTokenTypes()
	implementation = sts.ReplaceAll(implementation, "<TokenTypes>", tokenTypes)
	return implementation
}

func (v *generator_) populateScannerTemplate(
	template string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	var notice = v.extractNotice(syntax)
	implementation = sts.ReplaceAll(template, "<Notice>", notice)
	var tokenNames = v.extractTokenNames()
	implementation = sts.ReplaceAll(implementation, "<TokenNames>", tokenNames)
	var tokenMatchers = v.extractTokenMatchers()
	implementation = sts.ReplaceAll(implementation, "<TokenMatchers>", tokenMatchers)
	var foundCases = v.extractFoundCases()
	implementation = sts.ReplaceAll(implementation, "<FoundCases>", foundCases)
	var expressions = v.extractExpressions()
	implementation = sts.ReplaceAll(implementation, "<Expressions>", expressions)
	return implementation
}

func (v *generator_) populateSyntaxTemplate(
	template string,
	syntax string,
	copyright string,
) (
	implementation string,
) {
	implementation = sts.ReplaceAll(template, "<Notice>", noticeTemplate_)
	copyright = v.expandCopyright(copyright)
	implementation = sts.ReplaceAll(implementation, "<Copyright>", copyright)
	var allCaps = sts.ToUpper(syntax)
	implementation = sts.ReplaceAll(implementation, "<SYNTAX>", allCaps)
	var uppercase = v.makeUppercase(syntax)
	implementation = sts.ReplaceAll(implementation, "<Syntax>", uppercase)
	var lowercase = v.makeLowercase(syntax)
	implementation = sts.ReplaceAll(implementation, "<syntax>", lowercase)
	return implementation
}

func (v *generator_) processAlternative(
	alternative ast.AlternativeLike,
) (
	regexp string,
) {
	regexp += "|"
	var parts = alternative.GetParts().GetIterator()
	for parts.HasNext() {
		var part = parts.GetNext()
		regexp += v.processPart(part)
	}
	return regexp
}

func (v *generator_) processBounded(
	bounded ast.BoundedLike,
) (
	regexp string,
) {
	var rune_ = bounded.GetRune()
	rune_ = v.processRune(rune_)
	if rune_ == "-" {
		rune_ = "\\-"
	}
	regexp += rune_
	var extent = bounded.GetOptionalExtent()
	if col.IsDefined(extent) {
		regexp += v.processExtent(extent)
	}
	return regexp
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
			// This attribute is optional.
			attribute = v.optionalizeAttribute(attribute)
		case "*", "+":
			// Turn the attribute into a sequence of that type attribute.
			attribute = v.pluralizeAttribute(attribute)
		}
	}
	return attribute
}

func (v *generator_) processCharacter(
	character ast.CharacterLike,
) (
	regexp string,
) {
	switch actual := character.GetAny().(type) {
	case ast.BoundedLike:
		regexp += v.processBounded(actual)
	case string:
		regexp += `" + ` + sts.ToLower(actual) + `_ + "`
	default:
		var message = fmt.Sprintf(
			"Found an invalid character type: %T",
			actual,
		)
		panic(message)
	}
	return regexp
}

func (v *generator_) processDefinition(
	name string,
	definition ast.DefinitionLike,
) (
	constructor mod.ConstructorLike,
	attributes abs.ListLike[mod.AttributeLike],
) {
	// Extract the attributes.
	attributes = col.List[mod.AttributeLike]()
	switch actual := definition.GetAny().(type) {
	case ast.InlinedLike:
		v.processInlined(actual, attributes)
	case ast.MultilinedLike:
		v.processMultilined(actual, attributes)
	default:
		panic("Found an empty definition.")
	}

	// Create the constructor.
	var abstraction = mod.Abstraction(name + "Like")
	name = "Make"
	var parameters = v.extractParameters(attributes)
	var iterator = parameters.GetIterator()
	var parameter = iterator.GetNext()
	var additionalParameters = col.List[mod.AdditionalParameterLike]()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var additionalParameter = mod.AdditionalParameter(parameter)
		additionalParameters.AppendValue(additionalParameter)
	}
	constructor = mod.Constructor(
		name,
		mod.Parameters(
			parameter,
			additionalParameters.(abs.Sequential[mod.AdditionalParameterLike]),
		),
		abstraction,
	)

	return constructor, attributes
}

func (v *generator_) processElement(
	element ast.ElementLike,
) (
	regexp string,
) {
	switch actual := element.GetAny().(type) {
	case ast.GroupedLike:
		regexp += v.processGrouped(actual)
	case ast.FilteredLike:
		regexp += v.processFiltered(actual)
	case ast.StringLike:
		regexp += v.processString(actual)
	default:
		var message = fmt.Sprintf(
			"Found an invalid element type: %T",
			actual,
		)
		panic(message)
	}
	return regexp
}

func (v *generator_) processExpression(
	expression ast.ExpressionLike,
) {
	var name = expression.GetLowercase()
	var pattern = expression.GetPattern()
	var regexp = `"`
	regexp += v.processPattern(pattern)
	regexp += `"`
	v.regexps_.SetValue(name, regexp)
}

func (v *generator_) processExtent(
	extent ast.ExtentLike,
) (
	regexp string,
) {
	regexp += "-"
	var rune_ = extent.GetRune()
	regexp += v.processRune(rune_)
	return regexp
}

func (v *generator_) processFactor(
	factor ast.FactorLike,
	attributes abs.ListLike[mod.AttributeLike],
) {
	// Attempt to extract the attribute definitions from the predicate string.
	var predicate = factor.GetPredicate()
	var attribute = v.processPredicate(predicate)
	if col.IsUndefined(attribute) {
		// The predicate does not correspond to an attribute.
		return
	}

	// Take into account any cardinality of the predicate.
	var cardinality = factor.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		// The attribute type may need to be "pluralized".
		attribute = v.processCardinality(attribute, cardinality)
	}

	// Add the attribute definition to our list.
	attributes.AppendValue(attribute)
}

func (v *generator_) processFiltered(
	filtered ast.FilteredLike,
) (
	regexp string,
) {
	var negation = filtered.GetOptionalNegation()
	regexp += "["
	if col.IsDefined(negation) {
		regexp += "^"
	}
	var characters = filtered.GetCharacters().GetIterator()
	for characters.HasNext() {
		var character = characters.GetNext()
		regexp += v.processCharacter(character)
	}
	regexp += "]"
	return regexp
}

func (v *generator_) processGrouped(
	grouped ast.GroupedLike,
) (
	regexp string,
) {
	var pattern = grouped.GetPattern()
	regexp += v.processPattern(pattern)
	regexp += "(" + regexp + ")"
	return regexp
}

func (v *generator_) processIdentifier(
	identifier ast.IdentifierLike,
) {
	var name = identifier.GetAny().(string)
	if !gra.Scanner().MatchToken(gra.LowercaseToken, name).IsEmpty() {
		v.tokens_.AddValue(name)
	}
}

func (v *generator_) processInlined(
	inlined ast.InlinedLike,
	attributes abs.ListLike[mod.AttributeLike],
) {
	// Extract the attributes.
	var iterator = inlined.GetFactors().GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		v.processFactor(factor, attributes)
	}
	v.consolidateAttributes(attributes)
}

func (v *generator_) processLine(
	line ast.LineLike,
) {
	var identifier = line.GetIdentifier()
	v.processIdentifier(identifier)
}

func (v *generator_) processMultilined(
	multilined ast.MultilinedLike,
	attributes abs.ListLike[mod.AttributeLike],
) {
	var lines = multilined.GetLines().GetIterator()
	for lines.HasNext() {
		var line = lines.GetNext()
		v.processLine(line)
	}
	var abstraction = mod.Abstraction("any")
	var attribute = mod.Attribute(
		"GetAny",
		abstraction,
	)
	attributes.AppendValue(attribute)
}

func (v *generator_) processPart(
	part ast.PartLike,
) (
	regexp string,
) {
	var element = part.GetElement()
	regexp += v.processElement(element)
	var cardinality = part.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case ast.ConstrainedLike:
			var number = actual.GetNumber()
			regexp += "{" + number
			var limit = actual.GetOptionalLimit()
			if col.IsDefined(limit) {
				regexp += ","
				number = limit.GetOptionalNumber()
				if col.IsDefined(number) {
					regexp += number
				}
			}
			regexp += "}"
		case string:
			regexp += actual
		default:
			var message = fmt.Sprintf(
				"Found an invalid cardinality type: %T",
				actual,
			)
			panic(message)
		}
	}
	return regexp
}

func (v *generator_) processPattern(
	pattern ast.PatternLike,
) (
	regexp string,
) {
	var parts = pattern.GetParts().GetIterator()
	var part = parts.GetNext()
	regexp += v.processPart(part)
	for parts.HasNext() {
		part = parts.GetNext()
		regexp += v.processPart(part)
	}
	var alternatives = pattern.GetAlternatives().GetIterator()
	for alternatives.HasNext() {
		var alternative = alternatives.GetNext()
		regexp += v.processAlternative(alternative)
	}
	return regexp
}

func (v *generator_) processPredicate(
	predicate ast.PredicateLike,
) (
	attribute mod.AttributeLike,
) {
	var actual = predicate.GetAny().(string)
	switch {
	case !gra.Scanner().MatchToken(gra.LiteralToken, actual).IsEmpty():
	case !gra.Scanner().MatchToken(gra.IntrinsicToken, actual).IsEmpty():
	default:
		// We know it is a rule or expression name which corresponds to an attribute
		// with a (non-generic) instance type, or a Go intrinsic "string" type
		// respectively.
		attribute = v.extractAttribute(actual)
	}
	return attribute
}

func (v *generator_) processRule(rule ast.RuleLike) {
	// Process the definition.
	var name = rule.GetUppercase()
	var definition = rule.GetDefinition()
	var constructor, attributes = v.processDefinition(name, definition)

	// Create the class interface.
	var class = v.generateClass(name, constructor)
	v.classes_.SetValue(name, class)

	// Create the instance interface.
	var instance = v.generateInstance(name, attributes)
	v.instances_.SetValue(name, instance)
}

func (v *generator_) processRune(
	rune_ string,
) (
	regexp string,
) {
	var character = rune_[1:2] //Remove the single quotes.
	character = v.escapeString(character)
	regexp += character
	return regexp
}

func (v *generator_) processString(
	string_ ast.StringLike,
) (
	regexp string,
) {
	var actual = string_.GetAny().(string)
	switch {
	case !gra.Scanner().MatchToken(gra.RuneToken, actual).IsEmpty():
		var literal = actual[1:2] // Remove the single quotes.
		regexp += v.escapeString(literal)
	case !gra.Scanner().MatchToken(gra.LiteralToken, actual).IsEmpty():
		var literal = actual[1 : len(actual)-1] // Remove the double quotes.
		regexp += v.escapeString(literal)
	case !gra.Scanner().MatchToken(gra.LowercaseToken, actual).IsEmpty():
		regexp += `" + ` + actual + `_ + "`
	case !gra.Scanner().MatchToken(gra.IntrinsicToken, actual).IsEmpty():
		regexp += `" + ` + sts.ToLower(actual) + `_ + "`
	default:
		var message = fmt.Sprintf(
			"Found an invalid element string: %q",
			actual,
		)
		panic(message)
	}
	return regexp
}

var reserved_ = map[string]bool{
	"any":       true,
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
