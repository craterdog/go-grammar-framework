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
	stc "strconv"
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
	class_      VisitorClassLike
	visitor_    gra.VisitorLike
	attributes_ abs.ListLike[[2]string]
	rules_      abs.CatalogLike[string, abs.ListLike[[2]string]]
	plurals_    abs.SetLike[string]
	result_     string

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
) string {
	var name = v.extractSyntaxName(syntax)
	v.result_ = visitorTemplate_
	v.result_ = sts.ReplaceAll(v.result_, "<module>", module)
	var notice = v.extractNotice(syntax)
	v.result_ = sts.ReplaceAll(v.result_, "<Notice>", notice)
	var uppercase = v.makeUppercase(name)
	v.result_ = sts.ReplaceAll(v.result_, "<Name>", uppercase)
	var lowercase = v.makeLowercase(name)
	v.result_ = sts.ReplaceAll(v.result_, "<name>", lowercase)
	v.visitor_.VisitSyntax(syntax)
	return v.result_
}

// Methodical

func (v *visitor_) PreprocessFactor(
	factor ast.FactorLike,
	index uint,
	size uint,
) {
	switch actual := factor.GetAny().(type) {
	case string:
		// Use the actual literal as the plurality.
		v.attributes_.AppendValue([2]string{"delimiter", actual})
	}
}

func (v *visitor_) PostprocessInlined(inlined ast.InlinedLike) {
	v.consolidateAttributes()
}

func (v *visitor_) PreprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
	var plurality = "alternative"
	var identifier = line.GetIdentifier().GetAny().(string)
	v.attributes_.AppendValue([2]string{identifier, plurality})
}

func (v *visitor_) PreprocessPredicate(
	predicate ast.PredicateLike,
) {
	var plurality = "singular"
	var identifier = predicate.GetIdentifier().GetAny().(string)
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
	if plurality == "repeated" {
		v.plurals_.AddValue(identifier)
	}
	v.attributes_.AppendValue([2]string{identifier, plurality})
}

func (v *visitor_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var name = rule.GetUppercase()
	v.attributes_ = col.List[[2]string]()
	v.rules_.SetValue(name, v.attributes_)
}

func (v *visitor_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.rules_ = col.Catalog[string, abs.ListLike[[2]string]]()
	v.plurals_ = col.Set[string]()
}

func (v *visitor_) PostprocessSyntax(syntax ast.SyntaxLike) {
	v.rules_.SortValues()
	var methods = v.extractMethods()
	v.result_ = sts.ReplaceAll(v.result_, "<Methods>", methods)
}

// Private

func (v *visitor_) consolidateAttributes() {
	// Compare each attribute type and rename duplicates.
	for i := 1; i <= v.attributes_.GetSize(); i++ {
		var attribute = v.attributes_.GetValue(i)
		var first = attribute[0]
		for j := i + 1; j <= v.attributes_.GetSize(); j++ {
			var count = 1
			attribute = v.attributes_.GetValue(j)
			var second = attribute[0]
			if first == second {
				count++
				attribute[0] = second + stc.Itoa(count)
				v.attributes_.SetValue(j, attribute)
			}
		}
	}
}

func (v *visitor_) extractMethods() string {
	var methods string
	var names = v.rules_.GetKeys().GetIterator()
	for names.HasNext() {
		var method string
		var name = names.GetNext()
		var attributes = v.rules_.GetValue(name)
		var plurality = attributes.GetValue(1)[1]
		switch plurality {
		case "alternative":
			method = v.extractMultilineMethod(name)
		default:
			method = v.extractInlineMethod(name)
		}
		var lowercase = v.makeLowercase(name)
		method = sts.ReplaceAll(method, "<rule>", lowercase)
		var uppercase = v.makeUppercase(name)
		method = sts.ReplaceAll(method, "<Rule>", uppercase)
		methods += method
	}
	return methods
}

func (v *visitor_) extractInlineMethod(name string) string {
	var implementation string
	var attributes = v.rules_.GetValue(name).GetIterator()
	for attributes.HasNext() {
		var attribute = attributes.GetNext()
		implementation += v.extractInlineAttribute(attribute)
	}
	var method = methodTemplate_
	method = sts.ReplaceAll(method, "<Implementation>", implementation)
	return method
}

func (v *visitor_) extractInlineAttribute(attribute [2]string) string {
	var implementation string
	var identifier = attribute[0]
	var plurality = attribute[1]
	switch {
	case gra.Scanner().MatchesType(plurality, gra.LiteralToken):
		implementation = v.extractInlineLiteralAttribute(attribute)
	case gra.Scanner().MatchesType(identifier, gra.LowercaseToken):
		implementation = v.extractInlineTokenAttribute(attribute)
	case gra.Scanner().MatchesType(identifier, gra.UppercaseToken):
		implementation = v.extractInlineRuleAttribute(attribute)
	}
	return implementation
}

func (v *visitor_) extractInlineLiteralAttribute(attribute [2]string) string {
	var implementation string
	var identifier = attribute[0]
	var literal = attribute[1]
	var lowercase = v.makeLowercase(identifier)
	var uppercase = v.makeUppercase(identifier)
	implementation = visitLiteralTemplate_
	implementation = sts.ReplaceAll(implementation, "<delimiterName>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<DelimiterName>", uppercase)
	implementation = sts.ReplaceAll(implementation, "<literal>", literal)
	return implementation
}

func (v *visitor_) extractInlineTokenAttribute(attribute [2]string) string {
	var implementation string
	var identifier = attribute[0]
	var lowercase = v.makeLowercase(identifier)
	var uppercase = v.makeUppercase(identifier)
	var plurality = attribute[1]
	switch plurality {
	case "optional":
		implementation = visitOptionalTokenTemplate_
	case "repeated":
		implementation = visitRepeatedTokenTemplate_
	default:
		implementation = visitTokenTemplate_
		if v.plurals_.ContainsValue(identifier) {
			implementation = visitSingleTokenTemplate_
		}
	}
	implementation = sts.ReplaceAll(implementation, "<tokenName>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<TokenName>", uppercase)
	lowercase = v.makePlural(lowercase)
	uppercase = v.makePlural(uppercase)
	implementation = sts.ReplaceAll(implementation, "<tokensName>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<TokensName>", uppercase)
	return implementation
}

func (v *visitor_) extractInlineRuleAttribute(attribute [2]string) string {
	var implementation string
	var identifier = attribute[0]
	var lowercase = v.makeLowercase(identifier)
	var uppercase = v.makeUppercase(identifier)
	var plurality = attribute[1]
	switch plurality {
	case "optional":
		implementation = visitOptionalRuleTemplate_
	case "repeated":
		implementation = visitRepeatedRuleTemplate_
	default:
		implementation = visitRuleTemplate_
		if v.plurals_.ContainsValue(identifier) {
			implementation = visitSingleRuleTemplate_
		}
	}
	implementation = sts.ReplaceAll(implementation, "<ruleName>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<RuleName>", uppercase)
	lowercase = v.makePlural(lowercase)
	uppercase = v.makePlural(uppercase)
	implementation = sts.ReplaceAll(implementation, "<rulesName>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<RulesName>", uppercase)
	return implementation
}

func (v *visitor_) extractMultilineMethod(name string) string {
	var tokenCases string
	var ruleCases string
	var attributes = v.rules_.GetValue(name).GetIterator()
	for attributes.HasNext() {
		var attribute = attributes.GetNext()
		var identifier = attribute[0]
		switch {
		case gra.Scanner().MatchesType(identifier, gra.LowercaseToken):
			tokenCases += v.extractMultilineTokenAttribute(attribute)
		case gra.Scanner().MatchesType(identifier, gra.UppercaseToken):
			ruleCases += v.extractMultilineRuleAttribute(attribute)
		}
	}
	var implementation = visitAnyTemplate_
	implementation = sts.ReplaceAll(implementation, "<RuleCases>", ruleCases)
	if len(tokenCases) > 0 {
		tokenCases = sts.ReplaceAll(visitMatchesTemplate_, "<TokenCases>", tokenCases)
	}
	implementation = sts.ReplaceAll(implementation, "<TokenCases>", tokenCases)
	var method = methodTemplate_
	method = sts.ReplaceAll(method, "<Implementation>", implementation)
	return method
}

func (v *visitor_) extractMultilineTokenAttribute(attribute [2]string) string {
	var implementation string
	var identifier = attribute[0]
	var lowercase = v.makeLowercase(identifier)
	var uppercase = v.makeUppercase(identifier)
	implementation = visitTokenCaseTemplate_
	implementation = sts.ReplaceAll(implementation, "<tokenName>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<TokenName>", uppercase)
	return implementation
}

func (v *visitor_) extractMultilineRuleAttribute(attribute [2]string) string {
	var implementation string
	var identifier = attribute[0]
	var lowercase = v.makeLowercase(identifier)
	var uppercase = v.makeUppercase(identifier)
	implementation = visitRuleCaseTemplate_
	implementation = sts.ReplaceAll(implementation, "<ruleName>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<RuleName>", uppercase)
	return implementation
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

const visitSingleTokenTemplate_ = `
	// Visit the <tokenName> token.
	var <tokenName> = <rule>.Get<TokenName>()
	v.processor_.Process<TokenName>(<tokenName>, 1, 1)
`

const visitRepeatedTokenTemplate_ = `
	// Visit each <tokenName> token.
	var <tokenName>Index uint
	var <tokensName> = <rule>.Get<TokensName>().GetIterator()
	var <tokensName>Size = uint(<tokensName>.GetSize())
	for <tokensName>.HasNext() {
		<tokenName>Index++
		var <tokenName> = <tokensName>.GetNext()
		v.processor_.Process<TokenName>(<tokenName>, <tokenName>Index, <tokensName>Size)
	}
`

const visitTokenCaseTemplate_ = `
		case Scanner().MatchesType(actual, <TokenName>Token):
			v.processor_.Process<TokenName>(actual)`

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

const visitSingleRuleTemplate_ = `
	// Visit the <ruleName>.
	var <ruleName> = <rule>.Get<RuleName>()
	v.processor_.Preprocess<RuleName>(<ruleName>, 1, 1)
	v.visit<RuleName>(<ruleName>)
	v.processor_.Postprocess<RuleName>(<ruleName>, 1, 1)
`

const visitRepeatedRuleTemplate_ = `
	// Visit each <ruleName>.
	var <ruleName>Index uint
	var <rulesName> = <rule>.Get<RulesName>().GetIterator()
	var <rulesName>Size = uint(<rulesName>.GetSize())
	for <rulesName>.HasNext() {
		<ruleName>Index++
		var <ruleName> = <rulesName>.GetNext()
		v.processor_.Preprocess<RuleName>(<ruleName>, <ruleName>Index, <rulesName>Size)
		v.visit<RuleName>(<ruleName>)
		v.processor_.Postprocess<RuleName>(<ruleName>, <ruleName>Index, <rulesName>Size)
	}
`

const visitRuleCaseTemplate_ = `
	case ast.<RuleName>Like:
		v.processor_.Preprocess<RuleName>(actual)
		v.visit<RuleName>(actual)
		v.processor_.Postprocess<RuleName>(actual)`

const visitLiteralTemplate_ = `
	// Visit the <literal> delimiter literal.
	var <delimiterName> = <rule>.Get<DelimiterName>()
	v.processor_.ProcessDelimiter(<delimiterName>)
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
