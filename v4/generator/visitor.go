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
	stc "strconv"
)

// CLASS ACCESS

// Reference

var visitorClass = &visitorClass_{
	// Initialize the class constants.
}

// Function

func Visitor() VisitorClassLike {
	return visitorClass
}

// CLASS METHODS

// Target

type visitorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *visitorClass_) Make() VisitorLike {
	var visitor = &visitor_{
		// Initialize the instance attributes.
		class_:    c,
		analyzer_: Analyzer().Make(),
	}
	return visitor
}

// INSTANCE METHODS

// Target

type visitor_ struct {
	// Define the instance attributes.
	class_    *visitorClass_
	analyzer_ AnalyzerLike
}

// Public

func (v *visitor_) GetClass() VisitorClassLike {
	return v.class_
}

func (v *visitor_) GenerateVisitorClass(
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
	var syntaxName = v.analyzer_.GetSyntaxName()
	implementation = replaceAll(implementation, "syntaxName", syntaxName)
	var methods = v.generateMethods()
	implementation = replaceAll(implementation, "methods", methods)
	return implementation
}

// Private

func (v *visitor_) generateInlineMethod(
	ruleName string,
) (
	implementation string,
) {
	var sequence = v.analyzer_.GetReferences(ruleName)
	var variableNames = generateVariableNames(sequence).GetIterator()
	var references = sequence.GetIterator()
	for references.HasNext() && variableNames.HasNext() {
		var slot uint = uint(variableNames.GetSlot())
		if slot > 0 {
			implementation += v.generateInlineSlot(ruleName, slot)
		}
		var variableName = variableNames.GetNext()
		var reference = references.GetNext()
		implementation += v.generateInlineReference(variableName, reference)
	}
	var method = v.getTemplate(visitRuleMethod)
	method = replaceAll(method, "implementation", implementation)
	return method
}

func (v *visitor_) generateInlineReference(
	variableName string,
	reference ast.ReferenceLike,
) (
	implementation string,
) {
	var identifier = reference.GetIdentifier().GetAny().(string)
	switch {
	case gra.Scanner().MatchesType(identifier, gra.LowercaseToken):
		implementation = v.generateInlineToken(variableName, reference)
	case gra.Scanner().MatchesType(identifier, gra.UppercaseToken):
		implementation = v.generateInlineRule(variableName, reference)
	}
	return implementation
}

func (v *visitor_) generateInlineSlot(
	ruleName string,
	slot uint,
) (
	implementation string,
) {
	implementation = v.getTemplate(visitSlot)
	implementation = replaceAll(implementation, "ruleName", ruleName)
	implementation = replaceAll(implementation, "slot", stc.Itoa(int(slot)))
	return implementation
}

func (v *visitor_) generateInlineRule(
	variableName string,
	reference ast.ReferenceLike,
) (
	implementation string,
) {
	switch v.generatePlurality(reference) {
	case "singular":
		implementation = v.getTemplate(visitSingularRule)
	case "optional":
		implementation = v.getTemplate(visitOptionalRule)
	case "repeated":
		implementation = v.getTemplate(visitRepeatedRule)
	default:
		implementation = v.getTemplate(visitRule)
	}
	implementation = replaceAll(implementation, "variableName", variableName)
	var ruleName = reference.GetIdentifier().GetAny().(string)
	implementation = replaceAll(implementation, "ruleName", ruleName)
	return implementation
}

func (v *visitor_) generateInlineToken(
	variableName string,
	reference ast.ReferenceLike,
) (
	implementation string,
) {
	switch v.generatePlurality(reference) {
	case "singular":
		implementation = v.getTemplate(visitSingularToken)
	case "optional":
		implementation = v.getTemplate(visitOptionalToken)
	case "repeated":
		implementation = v.getTemplate(visitRepeatedToken)
	default:
		implementation = v.getTemplate(visitToken)
	}
	implementation = replaceAll(implementation, "variableName", variableName)
	var tokenName = reference.GetIdentifier().GetAny().(string)
	implementation = replaceAll(implementation, "tokenName", tokenName)
	return implementation
}

func (v *visitor_) generateMethods() (
	implementation string,
) {
	var methods string
	var rules = v.analyzer_.GetRuleNames().GetIterator()
	for rules.HasNext() {
		var method string
		var rule = rules.GetNext()
		switch {
		case col.IsDefined(v.analyzer_.GetIdentifiers(rule)):
			method = v.generateMultilineMethod(rule)
		case col.IsDefined(v.analyzer_.GetReferences(rule)):
			method = v.generateInlineMethod(rule)
		}
		method = replaceAll(method, "rule", rule)
		methods += method
	}
	return methods
}

func (v *visitor_) generateMultilineMethod(
	ruleName string,
) (
	method string,
) {
	var tokenCases, ruleCases string
	var identifiers = v.analyzer_.GetIdentifiers(ruleName).GetIterator()
	for identifiers.HasNext() {
		var identifier = identifiers.GetNext()
		var name = identifier.GetAny().(string)
		switch {
		case gra.Scanner().MatchesType(name, gra.LowercaseToken):
			tokenCases += v.generateMultilineToken(name)
		case gra.Scanner().MatchesType(name, gra.UppercaseToken):
			ruleCases += v.generateMultilineRule(name)
		}
	}
	var implementation = v.getTemplate(visitCases)
	implementation = replaceAll(implementation, "ruleCases", ruleCases)
	implementation = replaceAll(implementation, "tokenCases", tokenCases)
	method = v.getTemplate(visitRuleMethod)
	method = replaceAll(method, "implementation", implementation)
	return method
}

func (v *visitor_) generateMultilineRule(
	ruleName string,
) (
	implementation string,
) {
	implementation = v.getTemplate(visitRuleCase)
	if v.analyzer_.IsPlural(ruleName) {
		implementation = v.getTemplate(visitSingularRuleCase)
	}
	implementation = replaceAll(implementation, "ruleName", ruleName)
	return implementation
}

func (v *visitor_) generateMultilineToken(
	tokenName string,
) (
	implementation string,
) {
	implementation = v.getTemplate(visitTokenCase)
	if v.analyzer_.IsPlural(tokenName) {
		implementation = v.getTemplate(visitSingularRuleCase)
	}
	implementation = replaceAll(implementation, "tokenName", tokenName)
	return implementation
}

func (v *visitor_) generatePlurality(
	reference ast.ReferenceLike,
) (
	plurality string,
) {
	var name = reference.GetIdentifier().GetAny().(string)
	var cardinality = reference.GetOptionalCardinality()
	if col.IsUndefined(cardinality) {
		if v.analyzer_.IsPlural(name) {
			plurality = "singular"
		}
		return plurality
	}
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		var token = actual.GetAny().(string)
		switch {
		case gra.Scanner().MatchesType(token, gra.OptionalToken):
			plurality = "optional"
			if v.analyzer_.IsPlural(name) {
				plurality = "singular"
			}
		case gra.Scanner().MatchesType(token, gra.RepeatedToken):
			plurality = "repeated"
		}
	case ast.QuantifiedLike:
		plurality = "repeated"
	}
	return plurality
}

// Private

func (v *visitor_) getTemplate(name string) string {
	var template = visitorTemplates_.GetValue(name)
	return template
}

// PRIVATE GLOBALS

// Constants

const (
	visitRuleMethod        = "visitRuleMethod"
	visitRule              = "visitRule"
	visitSingularRule      = "visitSingularRule"
	visitOptionalRule      = "visitOptionalRule"
	visitRepeatedRule      = "visitRepeatedRule"
	visitToken             = "visitToken"
	visitSingularToken     = "visitSingularToken"
	visitOptionalToken     = "visitOptionalToken"
	visitRepeatedToken     = "visitRepeatedToken"
	visitSlot              = "visitSlot"
	visitCases             = "visitCases"
	visitRuleCase          = "visitRuleCase"
	visitSingularRuleCase  = "visitSingularRuleCase"
	visitSingularTokenCase = "visitSingularTokenCase"
	visitTokenCase         = "visitTokenCase"
)

var visitorTemplates_ = col.Catalog[string, string](
	map[string]string{
		visitRuleMethod: `
func (v *visitor_) visit<Rule>(<rule_> ast.<Rule>Like) {<Implementation>}
`,
		visitRule: `
	// Visit the <ruleName> rule.
	var <variableName_> = <rule_>.Get<VariableName>()
	v.processor_.Preprocess<RuleName>(<variableName_>)
	v.visit<RuleName>(<variableName_>)
	v.processor_.Postprocess<RuleName>(<variableName_>)
`,
		visitSingularRule: `
	// Visit the <ruleName> rule.
	var <variableName_> = <rule_>.Get<VariableName>()
	if col.IsDefined(<variableName_>) {
		v.processor_.Preprocess<RuleName>(<variableName_>, 1, 1)
		v.visit<RuleName>(<variableName_>)
		v.processor_.Postprocess<RuleName>(<variableName_>, 1, 1)
	}
`,
		visitOptionalRule: `
	// Visit the optional <ruleName> rule.
	var <variableName_> = <rule_>.Get<VariableName>()
	if col.IsDefined(<variableName_>) {
		v.processor_.Preprocess<RuleName>(<variableName_>)
		v.visit<RuleName>(<variableName_>)
		v.processor_.Postprocess<RuleName>(<variableName_>)
	}
`,
		visitRepeatedRule: `
	// Visit each <ruleName> rule.
	var <ruleName>Index uint
	var <variableName_> = <rule_>.Get<VariableName>().GetIterator()
	var <variableName>Size = uint(<variableName_>.GetSize())
	for <variableName_>.HasNext() {
		<ruleName>Index++
		var <ruleName_> = <variableName_>.GetNext()
		v.processor_.Preprocess<RuleName>(
			<ruleName_>,
			<ruleName>Index,
			<variableName>Size,
		)
		v.visit<RuleName>(<ruleName_>)
		v.processor_.Postprocess<RuleName>(
			<ruleName_>,
			<ruleName>Index,
			<variableName>Size,
		)
	}
`,
		visitToken: `
	// Visit the <tokenName> token.
	var <variableName_> = <rule_>.Get<VariableName>()
	v.processor_.Process<TokenName>(<variableName_>)
`,
		visitSingularToken: `
	// Visit the <tokenName> token.
	var <variableName_> = <rule_>.Get<VariableName>()
	if col.IsDefined(<variableName_>) {
		v.processor_.Process<TokenName>(<variableName_>, 1, 1)
	}
`,
		visitOptionalToken: `
	// Visit the optional <tokenName> token.
	var <variableName_> = <rule_>.Get<VariableName>()
	if col.IsDefined(<variableName_>) {
		v.processor_.Process<TokenName>(<variableName_>)
	}
`,
		visitRepeatedToken: `
	// Visit each <tokenName> token.
	var <tokenName>Index uint
	var <variableName_> = <rule_>.Get<VariableName>().GetIterator()
	var <variableName>Size = uint(<variableName_>.GetSize())
	for <variableName_>.HasNext() {
		<tokenName>Index++
		var <tokenName_> = <variableName_>.GetNext()
		v.processor_.Process<TokenName>(
			<tokenName_>,
			<tokenName>Index,
			<variableName>Size,
		)
	}
`,
		visitSlot: `
	// Visit slot <slot> between references.
	v.processor_.Process<RuleName>Slot(<slot>)
`,
		visitCases: `
	// Visit the possible <rule> types.
	switch actual := <rule_>.GetAny().(type) {<RuleCases>
	case string:
		switch {<TokenCases>
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
`,
		visitRuleCase: `
	case ast.<RuleName>Like:
		v.processor_.Preprocess<RuleName>(actual)
		v.visit<RuleName>(actual)
		v.processor_.Postprocess<RuleName>(actual)`,
		visitSingularRuleCase: `
	case ast.<RuleName>Like:
		v.processor_.Preprocess<RuleName>(actual, 1, 1)
		v.visit<RuleName>(actual)
		v.processor_.Postprocess<RuleName>(actual, 1, 1)`,
		visitSingularTokenCase: `
		case Scanner().MatchesType(actual, <TokenName>Token):
			v.processor_.Process<TokenName>(actual, 1, 1)`,
		visitTokenCase: `
		case Scanner().MatchesType(actual, <TokenName>Token):
			v.processor_.Process<TokenName>(actual)`,
		classTemplate: `<Notice>

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	ast "<module>/ast"
)

// CLASS ACCESS

// Reference

var visitorClass = &visitorClass_{
	// Initialize the class constants.
}

// Function

func Visitor() VisitorClassLike {
	return visitorClass
}

// CLASS METHODS

// Target

type visitorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *visitorClass_) Make(processor Methodical) VisitorLike {
	return &visitor_{
		// Initialize the instance attributes.
		class_:     c,
		processor_: processor,
	}
}

// INSTANCE METHODS

// Target

type visitor_ struct {
	// Define the instance attributes.
	class_     *visitorClass_
	processor_ Methodical
}

// Public

func (v *visitor_) GetClass() VisitorClassLike {
	return v.class_
}

func (v *visitor_) Visit<SyntaxName>(<syntaxName> ast.<SyntaxName>Like) {
	// Visit the <syntaxName> syntax.
	v.processor_.Preprocess<SyntaxName>(<syntaxName>)
	v.visit<SyntaxName>(<syntaxName>)
	v.processor_.Postprocess<SyntaxName>(<syntaxName>)
}

// Private
<Methods>`,
	},
)
