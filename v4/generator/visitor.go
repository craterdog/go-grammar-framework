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
		analyzer_: gra.Analyzer().Make(),
	}
	return visitor
}

// INSTANCE METHODS

// Target

type visitor_ struct {
	// Define the instance attributes.
	class_    VisitorClassLike
	analyzer_ gra.AnalyzerLike
	result_   string
}

// Attributes

func (v *visitor_) GetClass() VisitorClassLike {
	return v.class_
}

// Public

func (v *visitor_) GenerateVisitorClass(
	module string,
	syntax ast.SyntaxLike,
) string {
	var syntaxName = v.generateSyntaxName(syntax)
	v.result_ = visitorTemplate_
	v.result_ = replaceAll(v.result_, "module", module)
	var notice = v.generateNotice(syntax)
	v.result_ = replaceAll(v.result_, "notice", notice)
	v.result_ = replaceAll(v.result_, "syntaxName", syntaxName)
	v.analyzer_.AnalyzeSyntax(syntax)
	var methods = v.generateMethods()
	v.result_ = replaceAll(v.result_, "methods", methods)
	return v.result_
}

// Private

func (v *visitor_) generateInlineMethod(name string) string {
	var implementation string
	var references = v.analyzer_.GetReferences(name).GetIterator()
	for references.HasNext() {
		var reference = references.GetNext()
		implementation += v.generateInlineReference(reference)
	}
	var method = methodTemplate_
	method = replaceAll(method, "implementation", implementation)
	return method
}

func (v *visitor_) generateInlineReference(reference ast.ReferenceLike) string {
	var name = reference.GetIdentifier().GetAny().(string)
	var cardinality = reference.GetOptionalCardinality()
	var implementation string
	switch {
	case gra.Scanner().MatchesType(name, gra.LowercaseToken):
		implementation = v.generateInlineToken(name, cardinality)
	case gra.Scanner().MatchesType(name, gra.UppercaseToken):
		implementation = v.generateInlineRule(name, cardinality)
	}
	return implementation
}

func (v *visitor_) generateInlineRule(
	ruleName string,
	cardinality ast.CardinalityLike,
) string {
	var implementation string
	switch v.generatePlurality(ruleName, cardinality) {
	case "singular":
		implementation = visitSingularRuleTemplate_
	case "optional":
		implementation = visitOptionalRuleTemplate_
	case "repeated":
		implementation = visitRepeatedRuleTemplate_
	default:
		implementation = visitRuleTemplate_
	}
	implementation = replaceAll(implementation, "ruleName", ruleName)
	var pluralName = makePlural(ruleName)
	implementation = replaceAll(implementation, "pluralName", pluralName)
	return implementation
}

func (v *visitor_) generateInlineToken(
	tokenName string,
	cardinality ast.CardinalityLike,
) string {
	var implementation string
	switch v.generatePlurality(tokenName, cardinality) {
	case "singular":
		implementation = visitSingularTokenTemplate_
	case "optional":
		implementation = visitOptionalTokenTemplate_
	case "repeated":
		implementation = visitRepeatedTokenTemplate_
	default:
		implementation = visitTokenTemplate_
	}
	implementation = replaceAll(implementation, "tokenName", tokenName)
	var pluralName = makePlural(tokenName)
	implementation = replaceAll(implementation, "pluralName", pluralName)
	return implementation
}

func (v *visitor_) generateMethods() string {
	var methods string
	var rules = v.analyzer_.GetRules().GetIterator()
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

func (v *visitor_) generateMultilineMethod(name string) string {
	var tokenCases, ruleCases string
	var identifiers = v.analyzer_.GetIdentifiers(name).GetIterator()

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
	var implementation = replaceAll(visitAnyTemplate_, "ruleCases", ruleCases)
	if len(tokenCases) > 0 {
		tokenCases = replaceAll(visitMatchesTemplate_, "tokenCases", tokenCases)
	}
	implementation = replaceAll(implementation, "tokenCases", tokenCases)
	return replaceAll(methodTemplate_, "implementation", implementation)
}

func (v *visitor_) generateMultilineRule(ruleName string) string {
	var template = visitRuleCaseTemplate_
	if v.analyzer_.IsPlural(ruleName) {
		template = visitSingularRuleCaseTemplate_
	}
	return replaceAll(template, "ruleName", ruleName)
}

func (v *visitor_) generateMultilineToken(tokenName string) string {
	var template = visitTokenCaseTemplate_
	if v.analyzer_.IsPlural(tokenName) {
		template = visitSingularTokenCaseTemplate_
	}
	return replaceAll(template, "tokenName", tokenName)
}

func (v *visitor_) generateNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

func (v *visitor_) generatePlurality(
	name string,
	cardinality ast.CardinalityLike,
) string {
	var plurality string
	if col.IsUndefined(cardinality) {
		if v.analyzer_.IsPlural(name) {
			plurality = "singular"
		}
		return plurality
	}
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstraintLike:
		var token = actual.GetAny().(string)
		switch {
		case gra.Scanner().MatchesType(token, gra.OptionalToken):
			plurality = "optional"
		case gra.Scanner().MatchesType(token, gra.RepeatedToken):
			plurality = "repeated"
		}
	case ast.CountLike:
		plurality = "repeated"
	}
	return plurality
}

func (v *visitor_) generateSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

const methodTemplate_ = `
func (v *visitor_) visit<Rule>(<rule> ast.<Rule>Like) {<Implementation>}
`

const visitAnyTemplate_ = `
	// Visit the possible <rule> types.
	switch actual := <rule>.GetAny().(type) {<RuleCases><TokenCases>
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
`

const visitMatchesTemplate_ = `
	case string:
		switch {<TokenCases>
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
`

const visitTokenTemplate_ = `
	// Visit the <tokenName> token.
	var <tokenName> = <rule>.Get<TokenName>()
	v.processor_.Process<TokenName>(<tokenName>)
`

const visitOptionalTokenTemplate_ = `
	// Visit the optional <tokenName> token.
	var <tokenName> = <rule>.GetOptional<TokenName>()
	if col.IsDefined(<tokenName>) {
		v.processor_.Process<TokenName>(<tokenName>)
	}
`

const visitSingularTokenTemplate_ = `
	// Visit the <tokenName> token.
	var <tokenName> = <rule>.Get<TokenName>()
	v.processor_.Process<TokenName>(<tokenName>, 1, 1)
`

const visitRepeatedTokenTemplate_ = `
	// Visit each <tokenName> token.
	var <tokenName>Index uint
	var <pluralName> = <rule>.Get<PluralName>().GetIterator()
	var <pluralName>Size = uint(<pluralName>.GetSize())
	for <pluralName>.HasNext() {
		<tokenName>Index++
		var <tokenName> = <pluralName>.GetNext()
		v.processor_.Process<TokenName>(
			<tokenName>,
			<tokenName>Index,
			<pluralName>Size,
		)
	}
`

const visitTokenCaseTemplate_ = `
		case Scanner().MatchesType(actual, <TokenName>Token):
			v.processor_.Process<TokenName>(actual)`

const visitSingularTokenCaseTemplate_ = `
		case Scanner().MatchesType(actual, <TokenName>Token):
			v.processor_.Process<TokenName>(actual, 1, 1)`

const visitRuleTemplate_ = `
	// Visit the <ruleName> rule.
	var <ruleName> = <rule>.Get<RuleName>()
	v.processor_.Preprocess<RuleName>(<ruleName>)
	v.visit<RuleName>(<ruleName>)
	v.processor_.Postprocess<RuleName>(<ruleName>)
`

const visitOptionalRuleTemplate_ = `
	// Visit the optional <ruleName> rule.
	var <ruleName> = <rule>.GetOptional<RuleName>()
	if col.IsDefined(<ruleName>) {
		v.processor_.Preprocess<RuleName>(<ruleName>)
		v.visit<RuleName>(<ruleName>)
		v.processor_.Postprocess<RuleName>(<ruleName>)
	}
`

const visitSingularRuleTemplate_ = `
	// Visit the <ruleName> rule.
	var <ruleName> = <rule>.Get<RuleName>()
	v.processor_.Preprocess<RuleName>(<ruleName>, 1, 1)
	v.visit<RuleName>(<ruleName>)
	v.processor_.Postprocess<RuleName>(<ruleName>, 1, 1)
`

const visitRepeatedRuleTemplate_ = `
	// Visit each <ruleName> rule.
	var <ruleName>Index uint
	var <pluralName> = <rule>.Get<PluralName>().GetIterator()
	var <pluralName>Size = uint(<pluralName>.GetSize())
	for <pluralName>.HasNext() {
		<ruleName>Index++
		var <ruleName> = <pluralName>.GetNext()
		v.processor_.Preprocess<RuleName>(
			<ruleName>,
			<ruleName>Index,
			<pluralName>Size,
		)
		v.visit<RuleName>(<ruleName>)
		v.processor_.Postprocess<RuleName>(
			<ruleName>,
			<ruleName>Index,
			<pluralName>Size,
		)
	}
`

const visitRuleCaseTemplate_ = `
	case ast.<RuleName>Like:
		v.processor_.Preprocess<RuleName>(actual)
		v.visit<RuleName>(actual)
		v.processor_.Postprocess<RuleName>(actual)`

const visitSingularRuleCaseTemplate_ = `
	case ast.<RuleName>Like:
		v.processor_.Preprocess<RuleName>(actual, 1, 1)
		v.visit<RuleName>(actual)
		v.processor_.Postprocess<RuleName>(actual, 1, 1)`

const visitorTemplate_ = `/*<Notice>*/

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
	class_     VisitorClassLike
	processor_ Methodical
}

// Attributes

func (v *visitor_) GetClass() VisitorClassLike {
	return v.class_
}

func (v *visitor_) GetProcessor() Methodical {
	return v.processor_
}

// Public

func (v *visitor_) Visit<SyntaxName>(<syntaxName> ast.<SyntaxName>Like) {
	// Visit the <syntaxName> syntax.
	v.processor_.Preprocess<SyntaxName>(<syntaxName>)
	v.visit<SyntaxName>(<syntaxName>)
	v.processor_.Postprocess<SyntaxName>(<syntaxName>)
}

// Private
<Methods>`
