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

package grammar

import (
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	stc "strconv"
	sts "strings"
)

// CLASS ACCESS

// Reference

var analyzerClass = &analyzerClass_{
	// Initialize the class constants.
}

// Function

func Analyzer() AnalyzerClassLike {
	return analyzerClass
}

// CLASS METHODS

// Target

type analyzerClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *analyzerClass_) Make() AnalyzerLike {
	var analyzer = &analyzer_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	analyzer.visitor_ = Visitor().Make(analyzer)
	return analyzer
}

// INSTANCE METHODS

// Target

type analyzer_ struct {
	// Define the instance attributes.
	class_        AnalyzerClassLike
	visitor_      VisitorLike
	isGreedy_     bool
	inDefinition_ bool
	inPattern_    bool
	hasLiteral_   bool
	syntaxName_   string
	notice_       string
	ruleName_     string
	regexp_       string
	ruleNames_    abs.SetLike[string]
	tokenNames_   abs.SetLike[string]
	pluralNames_  abs.SetLike[string]
	delimited_    abs.SetLike[string]
	ignored_      abs.SetLike[string]
	delimiters_   abs.SetLike[string]
	regexps_      abs.CatalogLike[string, string]
	terms_        abs.CatalogLike[string, abs.ListLike[ast.TermLike]]
	references_   abs.CatalogLike[string, abs.ListLike[ast.ReferenceLike]]
	identifiers_  abs.CatalogLike[string, abs.ListLike[ast.IdentifierLike]]

	// Define the inherited aspects.
	Methodical
}

// Attributes

func (v *analyzer_) GetClass() AnalyzerClassLike {
	return v.class_
}

// Methodical

func (v *analyzer_) ProcessExcluded(excluded string) {
	v.regexp_ += "^"
}

func (v *analyzer_) ProcessGlyph(glyph string) {
	var character = glyph[1:2] //Remove the single quotes.
	character = v.escapeText(character)
	v.regexp_ += character
}

func (v *analyzer_) ProcessIntrinsic(intrinsic string) {
	intrinsic = sts.ToLower(intrinsic)
	if intrinsic == "any" {
		v.isGreedy_ = false // Turn off "greedy" for expressions containing ANY.
	}
	v.regexp_ += `" + ` + intrinsic + `_ + "`
}

func (v *analyzer_) ProcessLiteral(literal string) {
	v.hasLiteral_ = true
	var delimiter, err = stc.Unquote(literal) // Remove the double quotes.
	if err != nil {
		panic(err)
	}
	delimiter = v.escapeText(delimiter)
	if v.inDefinition_ {
		v.delimiters_.AddValue(delimiter)
	}
	v.regexp_ += delimiter
}

func (v *analyzer_) ProcessLowercase(lowercase string) {
	if v.inDefinition_ {
		v.tokenNames_.AddValue(lowercase)
	}
	if v.inPattern_ {
		v.regexp_ += `(?:" + ` + lowercase + `_ + ")`
	}
}

func (v *analyzer_) ProcessNumber(number string) {
	v.regexp_ += number
}

func (v *analyzer_) ProcessOptional(optional string) {
	v.regexp_ += optional
}

func (v *analyzer_) ProcessRepeated(repeated string) {
	v.regexp_ += repeated
}

func (v *analyzer_) PreprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
	size uint,
) {
	v.regexp_ += "|"
}

func (v *analyzer_) PostprocessConstrained(constrained ast.ConstrainedLike) {
	if !v.isGreedy_ {
		v.regexp_ += "?"
		v.isGreedy_ = true // Reset scanning back to "greedy".
	}
}

func (v *analyzer_) PreprocessDefinition(definition ast.DefinitionLike) {
	v.inDefinition_ = true
}

func (v *analyzer_) PostprocessDefinition(definition ast.DefinitionLike) {
	v.inDefinition_ = false
}

func (v *analyzer_) PreprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
	v.regexp_ = `"(?:`
}

func (v *analyzer_) PostprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
	v.regexp_ += `)"`
	var name = expression.GetLowercase()
	v.regexps_.SetValue(name, v.regexp_)
}

func (v *analyzer_) PreprocessExtent(extent ast.ExtentLike) {
	v.regexp_ += "-"
}

func (v *analyzer_) PreprocessFilter(filter ast.FilterLike) {
	v.regexp_ += "["
}

func (v *analyzer_) PostprocessFilter(filter ast.FilterLike) {
	v.regexp_ += "]"
}

func (v *analyzer_) PreprocessGroup(group ast.GroupLike) {
	v.regexp_ += "("
}

func (v *analyzer_) PostprocessGroup(group ast.GroupLike) {
	v.regexp_ += ")"
}

func (v *analyzer_) PreprocessIdentifier(identifier ast.IdentifierLike) {
	var name = identifier.GetAny().(string)
	if Scanner().MatchesType(name, LowercaseToken) {
		v.tokenNames_.AddValue(name)
	}
}

func (v *analyzer_) PreprocessLimit(limit ast.LimitLike) {
	v.regexp_ += ","
}

func (v *analyzer_) PreprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
	var identifier = line.GetIdentifier()
	var identifiers = v.identifiers_.GetValue(v.ruleName_)
	identifiers.AppendValue(identifier)
}

func (v *analyzer_) PreprocessPattern(definition ast.PatternLike) {
	v.inPattern_ = true
}

func (v *analyzer_) PostprocessPattern(definition ast.PatternLike) {
	v.inPattern_ = false
}

func (v *analyzer_) PreprocessQuantified(quantified ast.QuantifiedLike) {
	v.regexp_ += "{"
}

func (v *analyzer_) PostprocessQuantified(quantified ast.QuantifiedLike) {
	v.regexp_ += "}"
	if !v.isGreedy_ {
		v.regexp_ += "?"
		v.isGreedy_ = true // Reset scanning back to "greedy".
	}
}

func (v *analyzer_) PreprocessReference(reference ast.ReferenceLike) {
	var identifier = reference.GetIdentifier()
	var cardinality = reference.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		var name = identifier.GetAny().(string)
		v.checkPlurality(name, cardinality)
	}
	var references = v.references_.GetValue(v.ruleName_)
	references.AppendValue(reference)
}

func (v *analyzer_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	v.hasLiteral_ = false
	var ruleName = rule.GetUppercase()
	v.ruleName_ = ruleName
	v.ruleNames_.AddValue(ruleName)
	var definition = rule.GetDefinition()
	switch definition.GetAny().(type) {
	case ast.InlineLike:
		var terms = col.List[ast.TermLike]()
		v.terms_.SetValue(ruleName, terms)
		var references = col.List[ast.ReferenceLike]()
		v.references_.SetValue(ruleName, references)
	case ast.MultilineLike:
		var identifiers = col.List[ast.IdentifierLike]()
		v.identifiers_.SetValue(ruleName, identifiers)
	}
}

func (v *analyzer_) PostprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var ruleName = rule.GetUppercase()
	v.ruleName_ = ruleName
	v.ruleNames_.AddValue(ruleName)
	if v.hasLiteral_ {
		v.delimited_.AddValue(ruleName)
	}
}

func (v *analyzer_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.isGreedy_ = true // The default is "greedy" scanning.
	v.syntaxName_ = v.extractSyntaxName(syntax)
	v.notice_ = v.extractNotice(syntax)
	v.ruleNames_ = col.Set[string]()
	v.tokenNames_ = col.Set[string]([]string{"delimiter"})
	v.pluralNames_ = col.Set[string]()
	v.delimited_ = col.Set[string]()
	v.ignored_ = col.Set[string]([]string{"newline", "space"})
	v.delimiters_ = col.Set[string]()
	var implicit = map[string]string{"space": `"(?:[ \\t]+)"`}
	v.regexps_ = col.Catalog[string, string](implicit)
	v.terms_ = col.Catalog[string, abs.ListLike[ast.TermLike]]()
	v.references_ = col.Catalog[string, abs.ListLike[ast.ReferenceLike]]()
	v.identifiers_ = col.Catalog[string, abs.ListLike[ast.IdentifierLike]]()
}

func (v *analyzer_) PostprocessSyntax(syntax ast.SyntaxLike) {
	v.ignored_ = v.ignored_.GetClass().Sans(v.ignored_, v.tokenNames_)
	v.tokenNames_.AddValues(v.ignored_)
	var delimiters = `"(?:`
	if !v.delimiters_.IsEmpty() {
		var iterator = v.delimiters_.GetIterator()
		delimiters += iterator.GetNext()
		for iterator.HasNext() {
			delimiters += "|" + iterator.GetNext()
		}
	}
	delimiters += `)"`
	v.regexps_.SetValue("delimiter", delimiters)
	v.regexps_.SortValues()
}

// Public

func (v *analyzer_) AnalyzeSyntax(syntax ast.SyntaxLike) {
	v.visitor_.VisitSyntax(syntax)
}

func (v *analyzer_) GetExpressions() abs.Sequential[abs.AssociationLike[string, string]] {
	return v.regexps_
}

func (v *analyzer_) GetIdentifiers(ruleName string) abs.Sequential[ast.IdentifierLike] {
	return v.identifiers_.GetValue(ruleName)
}

func (v *analyzer_) GetIgnored() abs.Sequential[string] {
	return v.ignored_
}

func (v *analyzer_) GetNotice() string {
	return v.notice_
}

func (v *analyzer_) GetReferences(ruleName string) abs.Sequential[ast.ReferenceLike] {
	return v.references_.GetValue(ruleName)
}

func (v *analyzer_) GetRuleNames() abs.Sequential[string] {
	return v.ruleNames_
}

func (v *analyzer_) GetSyntaxName() string {
	return v.syntaxName_
}

func (v *analyzer_) GetTerms(ruleName string) abs.Sequential[ast.TermLike] {
	return v.terms_.GetValue(ruleName)
}

func (v *analyzer_) GetTokenNames() abs.Sequential[string] {
	return v.tokenNames_
}

func (v *analyzer_) IsDelimited(ruleName string) bool {
	return v.delimited_.ContainsValue(ruleName)
}

func (v *analyzer_) IsIgnored(tokenName string) bool {
	return v.ignored_.ContainsValue(tokenName)
}

func (v *analyzer_) IsPlural(name string) bool {
	return v.pluralNames_.ContainsValue(name)
}

// Private

func (v *analyzer_) checkPlurality(
	name string,
	cardinality ast.CardinalityLike,
) {
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		switch actual.GetAny().(string) {
		case "*", "+":
			v.pluralNames_.AddValue(name)
		}
	case ast.QuantifiedLike:
		v.pluralNames_.AddValue(name)
	}
}

func (v *analyzer_) escapeText(text string) string {
	var escaped string
	for _, character := range text {
		switch character {
		case '"':
			escaped += `\`
		case '.', '|', '^', '$', '+', '*', '?', '(', ')', '[', ']', '{', '}':
			escaped += `\\`
		case '\\':
			escaped += `\\\`
		}
		escaped += string(character)
	}
	return escaped
}

func (v *analyzer_) extractNotice(syntax ast.SyntaxLike) string {
	var comment = syntax.GetNotice().GetComment()

	// Strip off the syntax style comment delimiters.
	comment = comment[2 : len(comment)-3]

	// Add in the go style comment delimiters.
	var notice = "/*" + comment + "*/"

	return notice
}

func (v *analyzer_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rules = syntax.GetRules().GetIterator()
	// The first rule name is the name of the syntax.
	var rule = rules.GetNext()
	var name = rule.GetUppercase()
	return name
}
