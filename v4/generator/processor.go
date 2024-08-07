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
	sts "strings"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var processorClass = &processorClass_{
	// Initialize the class constants.
}

// Function

func Processor() ProcessorClassLike {
	return processorClass
}

// CLASS METHODS

// Target

type processorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *processorClass_) Make() ProcessorLike {
	var processor = &processor_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: gra.Processor().Make(),
	}
	processor.visitor_ = gra.Visitor().Make(processor)
	return processor
}

// INSTANCE METHODS

// Target

type processor_ struct {
	// Define the instance attributes.
	class_   ProcessorClassLike
	visitor_ gra.VisitorLike
	tokens_  abs.SetLike[string]
	rules_   abs.SetLike[string]
	plurals_ abs.SetLike[string]

	// Define the inherited aspects.
	gra.Methodical
}

// Attributes

func (v *processor_) GetClass() ProcessorClassLike {
	return v.class_
}

// Public

func (v *processor_) GenerateProcessorClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.visitor_.VisitSyntax(syntax)
	implementation = processorTemplate_
	var name = v.extractSyntaxName(syntax)
	implementation = sts.ReplaceAll(
		implementation,
		"<module>",
		module,
	)
	var notice = v.extractNotice(syntax)
	implementation = sts.ReplaceAll(
		implementation,
		"<Notice>",
		notice,
	)
	var tokenProcessors = v.extractProcessTokens()
	implementation = sts.ReplaceAll(
		implementation,
		"<TokenProcessors>",
		tokenProcessors,
	)
	var ruleProcessors = v.extractProcessRules()
	implementation = sts.ReplaceAll(
		implementation,
		"<RuleProcessors>",
		ruleProcessors,
	)
	var uppercase = v.makeUppercase(name)
	implementation = sts.ReplaceAll(
		implementation,
		"<Name>",
		uppercase,
	)
	var lowercase = v.makeLowercase(name)
	implementation = sts.ReplaceAll(
		implementation,
		"<name>",
		lowercase,
	)
	return implementation
}

// Methodical

func (v *processor_) PreprocessIdentifier(identifier ast.IdentifierLike) {
	var name = identifier.GetAny().(string)
	if gra.Scanner().MatchesType(name, gra.LowercaseToken) {
		v.tokens_.AddValue(name)
	}
}

func (v *processor_) PreprocessPredicate(
	predicate ast.PredicateLike,
) {
	var identifier = v.makeLowercase(predicate.GetIdentifier().GetAny().(string))
	var cardinality = predicate.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case ast.ConstrainedLike:
			v.plurals_.AddValue(identifier)
		case string:
			switch actual {
			case "*", "+":
				v.plurals_.AddValue(identifier)
			}
		}
	}
}

func (v *processor_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var name = rule.GetUppercase()
	v.rules_.AddValue(v.makeLowercase(name))
}

func (v *processor_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.tokens_ = col.Set[string]([]string{"delimiter"})
	v.rules_ = col.Set[string]()
	v.plurals_ = col.Set[string]()
}

// Private

func (v *processor_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

func (v *processor_) extractProcessRules() string {
	var processRules string
	var iterator = v.rules_.GetIterator()
	for iterator.HasNext() {
		var lowercase = iterator.GetNext()
		var isPlural = v.plurals_.ContainsValue(lowercase)
		var uppercase = v.makeUppercase(lowercase)
		var parameters string
		if isPlural {
			parameters += "\n\t"
		}
		parameters += lowercase + " ast." + uppercase + "Like"
		if isPlural {
			parameters += ",\n\tindex uint"
			parameters += ",\n\tsize uint,\n"
		}
		var processRule = processRuleTemplate_
		processRule = sts.ReplaceAll(processRule, "<RuleName>", uppercase)
		processRule = sts.ReplaceAll(processRule, "<parameters>", parameters)
		processRules += processRule
	}
	return processRules
}

func (v *processor_) extractProcessTokens() string {
	var processTokens string
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var lowercase = iterator.GetNext()
		var uppercase = v.makeUppercase(lowercase)
		var isPlural = v.plurals_.ContainsValue(lowercase)
		var parameters string
		if isPlural {
			parameters += "\n\t"
		}
		parameters += lowercase + " string"
		if isPlural {
			parameters += ",\n\tindex uint"
			parameters += ",\n\tsize uint,\n"
		}
		var processToken = processTokenTemplate_
		processToken = sts.ReplaceAll(processToken, "<TokenName>", uppercase)
		processToken = sts.ReplaceAll(processToken, "<parameters>", parameters)
		processTokens += processToken
	}
	return processTokens
}

func (v *processor_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *processor_) makeLowercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	name = string(runes)
	if reserved_[name] {
		name += "_"
	}
	return name
}

func (v *processor_) makeUppercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

const processTokenTemplate_ = `
func (v *processor_) Process<TokenName>(<parameters>) {
}
`

const processRuleTemplate_ = `
func (v *processor_) Preprocess<RuleName>(<parameters>) {
}

func (v *processor_) Postprocess<RuleName>(<parameters>) {
}
`

const processorTemplate_ = `/*<Notice>*/

package grammar

import (
	ast "<module>/ast"
)

// CLASS ACCESS

// Reference

var processorClass = &processorClass_{
	// Initialize the class constants.
}

// Function

func Processor() ProcessorClassLike {
	return processorClass
}

// CLASS METHODS

// Target

type processorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *processorClass_) Make() ProcessorLike {
	var processor = &processor_{
		// Initialize the instance attributes.
		class_: c,
	}
	return processor
}

// INSTANCE METHODS

// Target

type processor_ struct {
	// Define the instance attributes.
	class_ ProcessorClassLike
}

// Attributes

func (v *processor_) GetClass() ProcessorClassLike {
	return v.class_
}

// Methodical
<TokenProcessors><RuleProcessors>`
