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

var formatterClass = &formatterClass_{
	// Initialize the class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	var formatter = &formatter_{
		// Initialize the instance attributes.
		class_:    c,
		analyzer_: gra.Analyzer().Make(),
	}
	return formatter
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_    FormatterClassLike
	analyzer_ gra.AnalyzerLike
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

// Public

func (v *formatter_) GenerateFormatterClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = formatterTemplate_
	implementation = replaceAll(implementation, "module", module)
	var notice = v.analyzer_.GetNotice()
	implementation = replaceAll(implementation, "notice", notice)
	var tokenFormatters = v.generateTokenFormatters()
	implementation = replaceAll(implementation, "tokenFormatters", tokenFormatters)
	var ruleFormatters = v.generateRuleFormatters()
	implementation = replaceAll(implementation, "ruleFormatters", ruleFormatters)
	var name = v.analyzer_.GetName()
	implementation = replaceAll(implementation, "name", name)
	return implementation
}

// Private

func (v *formatter_) generateRuleFormatters() string {
	var ruleFormatters string
	var iterator = v.analyzer_.GetRules().GetIterator()
	for iterator.HasNext() {
		var ruleName = iterator.GetNext()
		if !v.analyzer_.IsDelimited(ruleName) {
			continue
		}
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
		var ruleFormatter = formatRuleTemplate_
		ruleFormatter = replaceAll(ruleFormatter, "ruleName", ruleName)
		ruleFormatter = replaceAll(ruleFormatter, "parameters", parameters)
		ruleFormatters += ruleFormatter
	}
	return ruleFormatters
}

func (v *formatter_) generateTokenFormatters() string {
	var tokenFormatters string
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
		var tokenProcessor = formatTokenTemplate_
		tokenProcessor = replaceAll(tokenProcessor, "tokenName", tokenName)
		tokenProcessor = replaceAll(tokenProcessor, "parameters", parameters)
		tokenFormatters += tokenProcessor
	}
	return tokenFormatters
}

const formatRuleTemplate_ = `
func (v *formatter_) Preprocess<RuleName>(<parameters>) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) Postprocess<RuleName>(<parameters>) {
	// TBD - Add formatting of the delimited rule.
}
`

const formatTokenTemplate_ = `
func (v *formatter_) Process<TokenName>(<parameters>) {
	v.appendString(<tokenName>)
}
`

const formatterTemplate_ = `<Notice>

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "<module>/ast"
	stc "strconv"
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// Initialize the class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	var formatter = &formatter_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	formatter.visitor_ = Visitor().Make(formatter)
	return formatter
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_   FormatterClassLike
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Define the inherited aspects.
	Methodical
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) GetDepth() uint {
	return v.depth_
}

// Methodical
<TokenFormatters><RuleFormatters>
// Public

func (v *formatter_) Format<Name>(<name> ast.<Name>Like) string {
	v.visitor_.Visit<Name>(<name>)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "    "
	var level uint
	for ; level < v.depth_; level++ {
		newline += indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
`
