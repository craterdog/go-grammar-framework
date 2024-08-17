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
	class_   VisitorClassLike
	visitor_ gra.VisitorLike
	rules_   abs.SetLike[string]
	method_  sts.Builder
	methods_ abs.CatalogLike[string, string]

	// Define the inherited aspects.
	gra.Methodical
}

// Attributes

func (v *visitor_) GetClass() VisitorClassLike {
	return v.class_
}

// Public

func (v *visitor_) GenerateVisitorClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.visitor_.VisitSyntax(syntax)
	implementation = visitorTemplate_
	var name = v.extractSyntaxName(syntax)
	implementation = sts.ReplaceAll(implementation, "<module>", module)
	var notice = v.extractNotice(syntax)
	implementation = sts.ReplaceAll(implementation, "<Notice>", notice)
	var uppercase = v.makeUppercase(name)
	implementation = sts.ReplaceAll(implementation, "<Name>", uppercase)
	var lowercase = v.makeLowercase(name)
	implementation = sts.ReplaceAll(implementation, "<name>", lowercase)
	var methods = v.extractMethods()
	implementation = sts.ReplaceAll(implementation, "<Methods>", methods)
	return implementation
}

// Methodical

func (v *visitor_) PreprocessPredicate(
	predicate ast.PredicateLike,
) {
	// Check to see if the predicate has plurality.
	var plurality = "singlular"
	var cardinality = predicate.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case ast.ConstrainedLike:
			switch actual.GetAny().(string) {
			case "?":
				plurality = "optional"
			case "*", "+":
				plurality = "repeated"
			}
		case ast.QuantifiedLike:
			plurality = "repeated"
		}
	}

	// Choose the right method template.
	var template string
	var identifier = predicate.GetIdentifier().GetAny().(string)
	var lowercase = v.makeLowercase(identifier)
	var uppercase = v.makeUppercase(identifier)
	switch {
	case gra.Scanner().MatchesType(identifier, gra.LowercaseToken):
		switch plurality {
		case "optional":
			template = visitOptionalTokenTemplate_
		case "repeated":
			template = visitRepeatedTokenTemplate_
		default:
			template = visitTokenTemplate_
		}
		template = sts.ReplaceAll(template, "<tokenName>", lowercase)
		template = sts.ReplaceAll(template, "<TokenName>", uppercase)
		lowercase = v.makePlural(lowercase)
		uppercase = v.makePlural(uppercase)
		template = sts.ReplaceAll(template, "<tokensName>", lowercase)
		template = sts.ReplaceAll(template, "<TokensName>", uppercase)
	case gra.Scanner().MatchesType(identifier, gra.UppercaseToken):
		switch plurality {
		case "optional":
			template = visitOptionalRuleTemplate_
		case "repeated":
			template = visitRepeatedRuleTemplate_
		default:
			template = visitRuleTemplate_
		}
		template = sts.ReplaceAll(template, "<ruleName>", lowercase)
		template = sts.ReplaceAll(template, "<RuleName>", uppercase)
		lowercase = v.makePlural(lowercase)
		uppercase = v.makePlural(uppercase)
		template = sts.ReplaceAll(template, "<rulesName>", lowercase)
		template = sts.ReplaceAll(template, "<RulesName>", uppercase)
	}

	v.method_.WriteString(template)
}

func (v *visitor_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var name = v.makeLowercase(rule.GetUppercase())
	v.rules_.AddValue(name)
	v.method_.Reset()
}

func (v *visitor_) PostprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var name = v.makeLowercase(rule.GetUppercase())
	v.methods_.SetValue(name, v.method_.String())
}

func (v *visitor_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.rules_ = col.Set[string]()
	v.methods_ = col.Catalog[string, string]()
}

// Private

func (v *visitor_) extractMethods() string {
	var methods string
	var iterator = v.rules_.GetIterator()
	for iterator.HasNext() {
		var rule = iterator.GetNext()
		var implementation = v.methods_.GetValue(rule)
		var method = methodTemplate_
		method = sts.ReplaceAll(method, "<Implementation>", implementation)
		method = sts.ReplaceAll(method, "<rule>", rule)
		var uppercase = v.makeUppercase(rule)
		method = sts.ReplaceAll(method, "<Rule>", uppercase)
		methods += method
	}
	return methods
}

func (v *visitor_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

func (v *visitor_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *visitor_) makeLowercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	name = string(runes)
	if reserved_[name] {
		name += "_"
	}
	return name
}

func (v *visitor_) makePlural(name string) string {
	if sts.HasSuffix(name, "s") {
		name += "es"
	} else {
		name += "s"
	}
	return name
}

func (v *visitor_) makeUppercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

const methodTemplate_ = `
func (v *visitor_) visit<Rule>(<rule> ast.<Rule>Like) {<Implementation>}
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

/*
const visitSingleTokenTemplate_ = `
	// Visit the <tokenName> token.
	var <tokenName> = <rule>.Get<TokenName>()
	v.processor_.Process<TokenName>(<tokenName>, 1, 1)
`
*/

const visitRepeatedTokenTemplate_ = `
	// Visit each <tokenName> token.
	var index uint
	var <tokensName> = <rule>.Get<TokensName>().GetIterator()
	var size = uint(<tokensName>.GetSize())
	for <tokensName>.HasNext() {
		index++
		var <tokenName> = <tokensName>.GetNext()
		v.processor_.Process<TokenName>(<tokenName>, index, size)
	}
`

const visitOptionalRuleTemplate_ = `
	// Visit the optional <ruleName>.
	var <ruleName> = <rule>.GetOptional<RuleName>()
	if col.IsDefined(<ruleName>) {
		v.processor_.Preprocess<RuleName>(<ruleName>)
		v.visit<RuleName>(<ruleName>)
		v.processor_.Postprocess<RuleName>(<ruleName>)
	}
`

const visitRuleTemplate_ = `
	// Visit the <ruleName>.
	var <ruleName> = <rule>.Get<RuleName>()
	v.processor_.Preprocess<RuleName>(<ruleName>)
	v.visit<RuleName>(<ruleName>)
	v.processor_.Postprocess<RuleName>(<ruleName>)
`

/*
const visitSingleRuleTemplate_ = `
	// Visit the <ruleName>.
	var <ruleName> = <rule>.Get<RuleName>()
	v.processor_.Preprocess<RuleName>(<ruleName>, 1, 1)
	v.visit<RuleName>(<ruleName>)
	v.processor_.Postprocess<RuleName>(<ruleName>, 1, 1)
`
*/

const visitRepeatedRuleTemplate_ = `
	// Visit each <ruleName>.
	var index uint
	var <rulesName> = <rule>.Get<RulesName>().GetIterator()
	var size = uint(<rulesName>.GetSize())
	for <rulesName>.HasNext() {
		index++
		var <ruleName> = <rulesName>.GetNext()
		v.processor_.Preprocess<RuleName>(<ruleName>, index, size)
		v.visit<RuleName>(<ruleName>)
		v.processor_.Postprocess<RuleName>(<ruleName>, index, size)
	}
`

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

func (c *visitorClass_) Make(
	processor Methodical,
) VisitorLike {
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

func (v *visitor_) Visit<Name>(<name> ast.<Name>Like) {
	// Visit the <name>.
	v.processor_.Preprocess<Name>(<name>)
	v.visit<Name>(<name>)
	v.processor_.Postprocess<Name>(<name>)
}

// Private
<Methods>`
