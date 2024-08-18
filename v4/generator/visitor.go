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
	class_      VisitorClassLike
	visitor_    gra.VisitorLike
	attributes_ abs.ListLike[[2]string]
	rules_      abs.CatalogLike[string, abs.ListLike[[2]string]]
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
}

func (v *visitor_) PostprocessSyntax(syntax ast.SyntaxLike) {
	v.rules_.SortValues()
	var methods = v.extractMethods()
	v.result_ = sts.ReplaceAll(v.result_, "<Methods>", methods)
}

// Private

func (v *visitor_) extractInlineAttribute(attribute [2]string) string {
	var implementation string
	var identifier = attribute[0]
	var lowercase = v.makeLowercase(identifier)
	var uppercase = v.makeUppercase(identifier)
	var plurality = attribute[1]
	switch {
	case gra.Scanner().MatchesType(identifier, gra.LowercaseToken):
		switch plurality {
		case "optional":
			implementation = visitOptionalTokenTemplate_
		case "repeated":
			implementation = visitRepeatedTokenTemplate_
		default:
			implementation = visitTokenTemplate_
		}
		implementation = sts.ReplaceAll(implementation, "<tokenName>", lowercase)
		implementation = sts.ReplaceAll(implementation, "<TokenName>", uppercase)
		lowercase = v.makePlural(lowercase)
		uppercase = v.makePlural(uppercase)
		implementation = sts.ReplaceAll(implementation, "<tokensName>", lowercase)
		implementation = sts.ReplaceAll(implementation, "<TokensName>", uppercase)
	case gra.Scanner().MatchesType(identifier, gra.UppercaseToken):
		switch plurality {
		case "optional":
			implementation = visitOptionalRuleTemplate_
		case "repeated":
			implementation = visitRepeatedRuleTemplate_
		default:
			implementation = visitRuleTemplate_
		}
		implementation = sts.ReplaceAll(implementation, "<ruleName>", lowercase)
		implementation = sts.ReplaceAll(implementation, "<RuleName>", uppercase)
		lowercase = v.makePlural(lowercase)
		uppercase = v.makePlural(uppercase)
		implementation = sts.ReplaceAll(implementation, "<rulesName>", lowercase)
		implementation = sts.ReplaceAll(implementation, "<RulesName>", uppercase)
	}
	return implementation
}

func (v *visitor_) extractTokenAttribute(attribute [2]string) string {
	var implementation string
	var identifier = attribute[0]
	var lowercase = v.makeLowercase(identifier)
	var uppercase = v.makeUppercase(identifier)
	implementation = visitTokenCaseTemplate_
	implementation = sts.ReplaceAll(implementation, "<tokenName>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<TokenName>", uppercase)
	return implementation
}

func (v *visitor_) extractRuleAttribute(attribute [2]string) string {
	var implementation string
	var identifier = attribute[0]
	var lowercase = v.makeLowercase(identifier)
	var uppercase = v.makeUppercase(identifier)
	implementation = visitRuleCaseTemplate_
	implementation = sts.ReplaceAll(implementation, "<ruleName>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<RuleName>", uppercase)
	return implementation
}

func (v *visitor_) extractMethods() string {
	var methods string
	var names = v.rules_.GetKeys().GetIterator()
	for names.HasNext() {
		var name = names.GetNext()
		var implementation string
		var ruleCases string
		var tokenCases string
		var method = methodTemplate_
		var list = v.rules_.GetValue(name)
		var attributes = list.GetIterator()
		var plurality = list.GetValue(1)[1]
		switch plurality {
		case "singular", "optional", "repeated":
			for attributes.HasNext() {
				var attribute = attributes.GetNext()
				implementation += v.extractInlineAttribute(attribute)
			}
			method = sts.ReplaceAll(method, "<Implementation>", implementation)
		case "alternative":
			for attributes.HasNext() {
				var attribute = attributes.GetNext()
				switch {
				case gra.Scanner().MatchesType(attribute[0], gra.LowercaseToken):
					tokenCases += v.extractTokenAttribute(attribute)
				case gra.Scanner().MatchesType(attribute[0], gra.UppercaseToken):
					ruleCases += v.extractRuleAttribute(attribute)
				}
			}
			var implementation = visitAnyTemplate_
			method = sts.ReplaceAll(method, "<Implementation>", implementation)
			method = sts.ReplaceAll(method, "<RuleCases>", ruleCases)
			if len(tokenCases) > 0 {
				tokenCases = sts.ReplaceAll(visitMatchesTemplate_, "<TokenCases>", tokenCases)
			}
			method = sts.ReplaceAll(method, "<TokenCases>", tokenCases)
		}
		var lowercase = v.makeLowercase(name)
		method = sts.ReplaceAll(method, "<rule>", lowercase)
		var uppercase = v.makeUppercase(name)
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

const visitTokenCaseTemplate_ = `
		case gra.Scanner().MatchesType(actual, gra.<TokenName>Token):
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
