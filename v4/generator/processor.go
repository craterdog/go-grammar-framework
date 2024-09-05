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
	implementation = processorTemplate_
	implementation = replaceAll(implementation, "module", module)
	var notice = v.analyzer_.GetNotice()
	implementation = replaceAll(implementation, "notice", notice)
	var tokenProcessors = v.generateTokenProcessors()
	implementation = replaceAll(implementation, "tokenProcessors", tokenProcessors)
	var ruleProcessors = v.generateRuleProcessors()
	implementation = replaceAll(implementation, "ruleProcessors", ruleProcessors)
	var name = v.analyzer_.GetName()
	implementation = replaceAll(implementation, "name", name)
	return implementation
}

// Private

func (v *processor_) generateRuleProcessors() string {
	var ruleProcessors string
	var iterator = v.analyzer_.GetRules().GetIterator()
	for iterator.HasNext() {
		var ruleName = iterator.GetNext()
		var className = makeUpperCase(ruleName)
		var parameterName = makeLowerCase(ruleName)
		var isPlural = v.analyzer_.IsPlural(ruleName)
		var parameters string
		if isPlural {
			parameters += "\n\t"
		}
		parameters += parameterName + " ast." + className + "Like"
		if isPlural {
			parameters += ",\n\tindex uint"
			parameters += ",\n\tsize uint,\n"
		}
		var ruleProcessor = processRuleTemplate_
		ruleProcessor = replaceAll(ruleProcessor, "ruleName", ruleName)
		ruleProcessor = replaceAll(ruleProcessor, "parameters", parameters)
		ruleProcessors += ruleProcessor
	}
	return ruleProcessors
}

func (v *processor_) generateTokenProcessors() string {
	var tokenProcessors string
	var iterator = v.analyzer_.GetTokens().GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		if v.analyzer_.IsIgnored(tokenName) || tokenName == "delimiter" {
			continue
		}
		var parameterName = makeLowerCase(tokenName)
		var isPlural = v.analyzer_.IsPlural(tokenName)
		var parameters string
		if isPlural {
			parameters += "\n\t"
		}
		parameters += parameterName + " string"
		if isPlural {
			parameters += ",\n\tindex uint"
			parameters += ",\n\tsize uint,\n"
		}
		var tokenProcessor = processTokenTemplate_
		tokenProcessor = replaceAll(tokenProcessor, "tokenName", tokenName)
		tokenProcessor = replaceAll(tokenProcessor, "parameters", parameters)
		tokenProcessors += tokenProcessor
	}
	return tokenProcessors
}

const processRuleTemplate_ = `
func (v *processor_) Preprocess<RuleName>(<parameters>) {
}

func (v *processor_) Postprocess<RuleName>(<parameters>) {
}
`

const processTokenTemplate_ = `
func (v *processor_) Process<TokenName>(<parameters>) {
}
`

const processorTemplate_ = `<Notice>

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
