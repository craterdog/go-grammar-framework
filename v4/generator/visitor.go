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
		class_: c,

		// Initialize the inherited aspects.
		Methodical: gra.Processor().Make(),
	}
	visitor.visitor_ = gra.Visitor().Make(visitor)
	return visitor
}

// INSTANCE METHODS

// Target

type visitor_ struct {
	// Define the instance attributes.
	class_       VisitorClassLike
	visitor_     gra.VisitorLike
	rules_       abs.SetLike[string]
	plurals_     abs.SetLike[string]
	identifiers_ abs.ListLike[ast.IdentifierLike]
	references_  abs.ListLike[ast.ReferenceLike]
	inline_      abs.CatalogLike[string, abs.ListLike[ast.ReferenceLike]]
	multiline_   abs.CatalogLike[string, abs.ListLike[ast.IdentifierLike]]
	cardinality_ ast.CardinalityLike
	result_      string

	// Define the inherited aspects.
	gra.Methodical
}

// Attributes

func (v *visitor_) GetClass() VisitorClassLike {
	return v.class_
}

// Methodical

func (v *visitor_) PreprocessBracket(bracket ast.BracketLike) {
	v.cardinality_ = bracket.GetCardinality()
}

func (v *visitor_) PostprocessBracket(bracket ast.BracketLike) {
	v.cardinality_ = nil
}

func (v *visitor_) PostprocessInline(inline ast.InlineLike) {
	v.consolidateReferences()
}

func (v *visitor_) PreprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
	var identifier = line.GetIdentifier()
	v.identifiers_.AppendValue(identifier)
}

func (v *visitor_) PreprocessReference(reference ast.ReferenceLike) {
	reference = v.augmentCardinality(reference)
	v.references_.AppendValue(reference)
}

func (v *visitor_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var identifier = rule.GetUppercase()
	v.rules_.AddValue(identifier)
	var definition = rule.GetDefinition()
	switch definition.GetAny().(type) {
	case ast.MultilineLike:
		v.identifiers_ = col.List[ast.IdentifierLike]()
		v.multiline_.SetValue(identifier, v.identifiers_)
	case ast.InlineLike:
		v.references_ = col.List[ast.ReferenceLike]()
		v.inline_.SetValue(identifier, v.references_)
	}
}

func (v *visitor_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.rules_ = col.Set[string]()
	v.plurals_ = col.Set[string]()
	v.multiline_ = col.Catalog[string, abs.ListLike[ast.IdentifierLike]]()
	v.inline_ = col.Catalog[string, abs.ListLike[ast.ReferenceLike]]()
}

func (v *visitor_) PostprocessSyntax(syntax ast.SyntaxLike) {
	var methods = v.generateMethods()
	v.result_ = replaceAll(v.result_, "methods", methods)
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
	v.visitor_.VisitSyntax(syntax)
	return v.result_
}

// Private

func (v *visitor_) augmentCardinality(reference ast.ReferenceLike) ast.ReferenceLike {
	var identifier = reference.GetIdentifier()
	var cardinality = reference.GetOptionalCardinality()
	if col.IsDefined(v.cardinality_) {
		// The cardinality of a bracket takes precedence.
		cardinality = v.cardinality_
		reference = ast.Reference().Make(identifier, cardinality)
	}
	if col.IsDefined(cardinality) {
		var name = identifier.GetAny().(string)
		switch actual := cardinality.GetAny().(type) {
		case ast.CountLike:
			v.plurals_.AddValue(name)
		case ast.ConstraintLike:
			switch actual.GetAny().(string) {
			case "*", "+":
				v.plurals_.AddValue(name)
			}
		}
	}
	return reference
}

func (v *visitor_) consolidateReferences() {
	// Compare each reference type and rename duplicates.
	for i := 1; i <= v.references_.GetSize(); i++ {
		var reference = v.references_.GetValue(i)
		var first = reference.GetIdentifier().GetAny().(string)
		for j := i + 1; j <= v.references_.GetSize(); j++ {
			var second = v.references_.GetValue(j).GetIdentifier().GetAny().(string)
			if first == second {
				var plural = v.pluralizeReference(reference)
				v.references_.SetValue(i, plural)
				v.references_.RemoveValue(j)
				j--
			}
		}
	}
}

func (v *visitor_) generateInlineMethod(name string) string {
	var implementation string
	var references = v.inline_.GetValue(name).GetIterator()
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
	var rules = v.rules_.GetIterator()
	for rules.HasNext() {
		var method string
		var rule = rules.GetNext()
		switch {
		case col.IsDefined(v.multiline_.GetValue(rule)):
			method = v.generateMultilineMethod(rule)
		case col.IsDefined(v.inline_.GetValue(rule)):
			method = v.generateInlineMethod(rule)
		}
		method = replaceAll(method, "rule", rule)
		methods += method
	}
	return methods
}

func (v *visitor_) generateMultilineMethod(name string) string {
	var tokenCases, ruleCases string
	var identifiers = v.multiline_.GetValue(name).GetIterator()
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
	return replaceAll(visitRuleCaseTemplate_, "ruleName", ruleName)
}

func (v *visitor_) generateMultilineToken(tokenName string) string {
	return replaceAll(visitTokenCaseTemplate_, "tokenName", tokenName)
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
		if v.plurals_.ContainsValue(name) {
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

func (v *visitor_) pluralizeReference(
	reference ast.ReferenceLike,
) ast.ReferenceLike {
	// Make the identifier plural.
	var identifier = reference.GetIdentifier()
	var name = identifier.GetAny().(string)
	name = makePlural(name)
	v.plurals_.AddValue(name)

	// Add a plural cardinality to the reference.
	var constraint = ast.Constraint().Make("*")
	var cardinality = ast.Cardinality().Make(constraint)
	reference = ast.Reference().Make(identifier, cardinality)
	return reference
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
	// Visit the optional <TokenName> token.
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

const visitRuleTemplate_ = `
	// Visit the <ruleName>.
	var <ruleName> = <rule>.Get<RuleName>()
	v.processor_.Preprocess<RuleName>(<ruleName>)
	v.visit<RuleName>(<ruleName>)
	v.processor_.Postprocess<RuleName>(<ruleName>)
`

const visitOptionalRuleTemplate_ = `
	// Visit the optional <RuleName>.
	var <ruleName> = <rule>.GetOptional<RuleName>()
	if col.IsDefined(<ruleName>) {
		v.processor_.Preprocess<RuleName>(<ruleName>)
		v.visit<RuleName>(<ruleName>)
		v.processor_.Postprocess<RuleName>(<ruleName>)
	}
`

const visitSingularRuleTemplate_ = `
	// Visit the <ruleName>.
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
	// Visit the <syntaxName>.
	v.processor_.Preprocess<SyntaxName>(<syntaxName>)
	v.visit<SyntaxName>(<syntaxName>)
	v.processor_.Postprocess<SyntaxName>(<syntaxName>)
}

// Private
<Methods>`
