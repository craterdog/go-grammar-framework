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
		class_:    c,
		analyzer_: gra.Analyzer().Make(),
	}
	return processor
}

// INSTANCE METHODS

// Target

type processor_ struct {
	// Define the instance attributes.
	class_    ProcessorClassLike
	analyzer_ gra.AnalyzerLike
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
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = v.getTemplate(classTemplate)
	implementation = replaceAll(implementation, "module", module)
	var notice = v.analyzer_.GetNotice()
	implementation = replaceAll(implementation, "notice", notice)
	var tokenProcessors = v.generateTokenProcessors()
	implementation = replaceAll(implementation, "tokenProcessors", tokenProcessors)
	var ruleProcessors = v.generateRuleProcessors()
	implementation = replaceAll(implementation, "ruleProcessors", ruleProcessors)
	var name = v.analyzer_.GetSyntaxName()
	implementation = replaceAll(implementation, "name", name)
	return implementation
}

// Private

func (v *processor_) generateRuleProcessors() string {
	var ruleProcessors string
	var iterator = v.analyzer_.GetRuleNames().GetIterator()
	for iterator.HasNext() {
		var ruleName = iterator.GetNext()
		var isPlural = v.analyzer_.IsPlural(ruleName)
		var parameters = v.getTemplate(ruleParameter)
		if isPlural {
			parameters = v.getTemplate(ruleParameters)
		}
		var ruleProcessor = v.getTemplate(processRule)
		ruleProcessor = replaceAll(ruleProcessor, "parameters", parameters)
		ruleProcessor = replaceAll(ruleProcessor, "ruleName", ruleName)
		ruleProcessors += ruleProcessor
	}
	return ruleProcessors
}

func (v *processor_) generateTokenProcessors() string {
	var tokenProcessors string
	var iterator = v.analyzer_.GetTokenNames().GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		if v.analyzer_.IsIgnored(tokenName) || tokenName == "delimiter" {
			continue
		}
		var isPlural = v.analyzer_.IsPlural(tokenName)
		var parameters = v.getTemplate(tokenParameter)
		if isPlural {
			parameters = v.getTemplate(tokenParameters)
		}
		var tokenProcessor = v.getTemplate(processToken)
		tokenProcessor = replaceAll(tokenProcessor, "parameters", parameters)
		tokenProcessor = replaceAll(tokenProcessor, "tokenName", tokenName)
		tokenProcessors += tokenProcessor
	}
	return tokenProcessors
}

func (v *processor_) getTemplate(name string) string {
	var template = processorTemplates_.GetValue(name)
	return template
}

// PRIVATE GLOBALS

// Constants

const (
	processRule     = "processRule"
	ruleParameter   = "ruleParameter"
	ruleParameters  = "ruleParameters"
	processToken    = "processToken"
	tokenParameter  = "tokenParameter"
	tokenParameters = "tokenParameters"
)

var processorTemplates_ = col.Catalog[string, string](
	map[string]string{
		processRule: `
func (v *processor_) Preprocess<RuleName>(<parameters>) {
}

func (v *processor_) Postprocess<RuleName>(<parameters>) {
}
`,
		ruleParameter: `<ruleName_> ast.<RuleName>Like`,
		ruleParameters: `
	<ruleName_> ast.<RuleName>Like,
	index uint,
	size uint,
`,
		processToken: `
func (v *processor_) Process<TokenName>(<parameters>) {
}
`,
		tokenParameter: `<tokenName_> string`,
		tokenParameters: `
	<tokenName_> string,
	index uint,
	size uint,
`,
		classTemplate: `<Notice>

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
<TokenProcessors><RuleProcessors>`,
	},
)
