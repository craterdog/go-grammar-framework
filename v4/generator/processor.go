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
	class_       ProcessorClassLike
	visitor_     gra.VisitorLike
	tokens_      abs.SetLike[string]
	rules_       abs.SetLike[string]
	plurals_     abs.SetLike[string]
	cardinality_ ast.CardinalityLike

	// Define the inherited aspects.
	gra.Methodical
}

// Attributes

func (v *processor_) GetClass() ProcessorClassLike {
	return v.class_
}

// Methodical

func (v *processor_) PreprocessBracket(bracket ast.BracketLike) {
	v.cardinality_ = bracket.GetCardinality()
}

func (v *processor_) PostprocessBracket(bracket ast.BracketLike) {
	v.cardinality_ = nil
}

func (v *processor_) PreprocessIdentifier(identifier ast.IdentifierLike) {
	var name = identifier.GetAny().(string)
	if gra.Scanner().MatchesType(name, gra.LowercaseToken) {
		v.tokens_.AddValue(name)
	}
}

func (v *processor_) PreprocessReference(reference ast.ReferenceLike) {
	var identifier = makeLowerCase(reference.GetIdentifier().GetAny().(string))
	var cardinality = reference.GetOptionalCardinality()
	if col.IsDefined(v.cardinality_) {
		// The cardinality of a bracket takes precedence.
		cardinality = v.cardinality_
	}
	if col.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case ast.CountLike:
			v.plurals_.AddValue(identifier)
		case ast.ConstraintLike:
			switch actual.GetAny().(string) {
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
	v.rules_.AddValue(makeLowerCase(name))
}

func (v *processor_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.tokens_ = col.Set[string]()
	v.rules_ = col.Set[string]()
	v.plurals_ = col.Set[string]()
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
	implementation = replaceAll(implementation, "module", module)
	var notice = v.generateNotice(syntax)
	implementation = replaceAll(implementation, "notice", notice)
	var tokenProcessors = v.generateTokenProcessors()
	implementation = replaceAll(implementation, "tokenProcessors", tokenProcessors)
	var ruleProcessors = v.generateRuleProcessors()
	implementation = replaceAll(implementation, "ruleProcessors", ruleProcessors)
	var name = v.generateSyntaxName(syntax)
	implementation = replaceAll(implementation, "name", name)
	return implementation
}

// Private

func (v *processor_) generateNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

func (v *processor_) generateRuleProcessors() string {
	var ruleProcessors string
	var iterator = v.rules_.GetIterator()
	for iterator.HasNext() {
		var ruleName = iterator.GetNext()
		var className = makeUpperCase(ruleName)
		var isPlural = v.plurals_.ContainsValue(ruleName)
		if col.IsDefined(v.cardinality_) {
			// The cardinality of a bracket takes precedence.
			isPlural = true
		}
		var parameters string
		if isPlural {
			parameters += "\n\t"
		}
		parameters += ruleName + " ast." + className + "Like"
		if isPlural {
			parameters += ",\n\tindex uint"
			parameters += ",\n\tsize uint,\n"
		}
		var ruleProcessor = ruleProcessorTemplate_
		ruleProcessor = replaceAll(ruleProcessor, "ruleName", ruleName)
		ruleProcessor = replaceAll(ruleProcessor, "parameters", parameters)
		ruleProcessors += ruleProcessor
	}
	return ruleProcessors
}

func (v *processor_) generateSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *processor_) generateTokenProcessors() string {
	var tokenProcessors string
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var isPlural = v.plurals_.ContainsValue(tokenName)
		var parameters string
		if isPlural {
			parameters += "\n\t"
		}
		parameters += tokenName + " string"
		if isPlural {
			parameters += ",\n\tindex uint"
			parameters += ",\n\tsize uint,\n"
		}
		var tokenProcessor = tokenProcessorTemplate_
		tokenProcessor = replaceAll(tokenProcessor, "tokenName", tokenName)
		tokenProcessor = replaceAll(tokenProcessor, "parameters", parameters)
		tokenProcessors += tokenProcessor
	}
	return tokenProcessors
}

const tokenProcessorTemplate_ = `
func (v *processor_) Process<TokenName>(<parameters>) {
}
`

const ruleProcessorTemplate_ = `
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
