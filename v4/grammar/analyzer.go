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
	depth_        uint
	name_         string
	notice_       string
	regexp_       string
	rules_        abs.SetLike[string]
	tokens_       abs.SetLike[string]
	plurals_      abs.SetLike[string]
	ignored_      abs.SetLike[string]
	literals_     abs.SetLike[string]
	identifiers_  abs.ListLike[ast.IdentifierLike]
	references_   abs.ListLike[ast.ReferenceLike]
	regexps_      abs.CatalogLike[string, string]
	inlines_      abs.CatalogLike[string, abs.ListLike[ast.ReferenceLike]]
	multilines_   abs.CatalogLike[string, abs.ListLike[ast.IdentifierLike]]
	cardinality_  ast.CardinalityLike

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

func (v *analyzer_) ProcessIntrinsic(intrinsic string) {
	intrinsic = sts.ToLower(intrinsic)
	if intrinsic == "any" {
		v.isGreedy_ = false // Turn off "greedy" for expressions containing ANY.
	}
	v.regexp_ += `" + ` + intrinsic + `_ + "`
}

func (v *analyzer_) ProcessLiteral(literal string) {
	literal = literal[1 : len(literal)-1] // Remove the double quotes.
	literal = v.escapeText(literal)
	if v.inDefinition_ {
		v.literals_.AddValue(literal)
	}
	v.regexp_ += literal
}

func (v *analyzer_) ProcessLowercase(lowercase string) {
	if v.inDefinition_ {
		v.tokens_.AddValue(lowercase)
	}
	if v.depth_ > 0 {
		v.regexp_ += `(?:" + ` + lowercase + `_ + ")`
	}
}

func (v *analyzer_) ProcessNumber(
	number string,
	index uint,
	size uint,
) {
	if index == 2 {
		v.regexp_ += ","
	}
	v.regexp_ += number
}

func (v *analyzer_) ProcessOptional(optional string) {
	v.regexp_ += optional
}

func (v *analyzer_) ProcessRepeated(repeated string) {
	v.regexp_ += repeated
}

func (v *analyzer_) ProcessRunic(
	runic string,
	index uint,
	size uint,
) {
	var character = runic[1:2] //Remove the single quotes.
	character = v.escapeText(character)
	if index == 2 {
		v.regexp_ += "-"
	}
	v.regexp_ += character
}

func (v *analyzer_) PreprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
	size uint,
) {
	if index > 1 {
		v.regexp_ += "|"
	}
}

func (v *analyzer_) PreprocessBracket(bracket ast.BracketLike) {
	v.cardinality_ = bracket.GetCardinality()
}

func (v *analyzer_) PostprocessBracket(bracket ast.BracketLike) {
	v.cardinality_ = nil
}

func (v *analyzer_) PostprocessConstraint(constraint ast.ConstraintLike) {
	if !v.isGreedy_ {
		v.regexp_ += "?"
		v.isGreedy_ = true // Reset scanning back to "greedy".
	}
}

func (v *analyzer_) PreprocessCount(count ast.CountLike) {
	v.regexp_ += "{"
}

func (v *analyzer_) PostprocessCount(constraint ast.CountLike) {
	v.regexp_ += "}"
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
		v.tokens_.AddValue(name)
	}
}

func (v *analyzer_) PostprocessInline(inline ast.InlineLike) {
	v.consolidateReferences()
}

func (v *analyzer_) PreprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
	var identifier = line.GetIdentifier()
	v.identifiers_.AppendValue(identifier)
}

func (v *analyzer_) PreprocessPattern(definition ast.PatternLike) {
	v.depth_++
}

func (v *analyzer_) PostprocessPattern(definition ast.PatternLike) {
	v.depth_--
}

func (v *analyzer_) PreprocessReference(reference ast.ReferenceLike) {
	reference = v.augmentCardinality(reference)
	v.references_.AppendValue(reference)
}

func (v *analyzer_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var identifier = rule.GetUppercase()
	v.rules_.AddValue(identifier)
	var definition = rule.GetDefinition()
	switch definition.GetAny().(type) {
	case ast.InlineLike:
		v.references_ = col.List[ast.ReferenceLike]()
		v.inlines_.SetValue(identifier, v.references_)
	case ast.MultilineLike:
		v.identifiers_ = col.List[ast.IdentifierLike]()
		v.multilines_.SetValue(identifier, v.identifiers_)
	}
}

func (v *analyzer_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.isGreedy_ = true // The default is "greedy" scanning.
	v.name_ = v.extractName(syntax)
	v.notice_ = v.extractNotice(syntax)
	v.rules_ = col.Set[string]()
	v.tokens_ = col.Set[string]([]string{"delimiter"})
	v.plurals_ = col.Set[string]()
	v.ignored_ = col.Set[string]([]string{"newline", "space"})
	v.literals_ = col.Set[string]()
	var implicit = map[string]string{"space": `"(?:[ \\t]+)"`}
	v.regexps_ = col.Catalog[string, string](implicit)
	v.inlines_ = col.Catalog[string, abs.ListLike[ast.ReferenceLike]]()
	v.multilines_ = col.Catalog[string, abs.ListLike[ast.IdentifierLike]]()
}

func (v *analyzer_) PostprocessSyntax(syntax ast.SyntaxLike) {
	v.ignored_ = v.ignored_.GetClass().Sans(v.ignored_, v.tokens_)
	v.tokens_.AddValues(v.ignored_)
	var literals = `"(?:`
	if !v.literals_.IsEmpty() {
		var iterator = v.literals_.GetIterator()
		literals += iterator.GetNext()
		for iterator.HasNext() {
			literals += "|" + iterator.GetNext()
		}
	}
	literals += `)"`
	v.regexps_.SetValue("delimiter", literals)
	v.regexps_.SortValues()
}

// Public

func (v *analyzer_) AnalyzeSyntax(syntax ast.SyntaxLike) {
	v.visitor_.VisitSyntax(syntax)
}

func (v *analyzer_) GetName() string {
	return v.name_
}

func (v *analyzer_) GetNotice() string {
	return v.notice_
}

func (v *analyzer_) GetTokens() abs.Sequential[string] {
	return v.tokens_
}

func (v *analyzer_) GetIgnored() abs.Sequential[string] {
	return v.ignored_
}

func (v *analyzer_) IsIgnored(token string) bool {
	return v.ignored_.ContainsValue(token)
}

func (v *analyzer_) GetRules() abs.Sequential[string] {
	return v.rules_
}

func (v *analyzer_) IsPlural(rule string) bool {
	return v.plurals_.ContainsValue(rule)
}

func (v *analyzer_) GetReferences(rule string) abs.Sequential[ast.ReferenceLike] {
	return v.inlines_.GetValue(rule)
}

func (v *analyzer_) GetIdentifiers(rule string) abs.Sequential[ast.IdentifierLike] {
	return v.multilines_.GetValue(rule)
}

func (v *analyzer_) GetExpressions() abs.Sequential[abs.AssociationLike[string, string]] {
	return v.regexps_
}

// Private

func (v *analyzer_) augmentCardinality(reference ast.ReferenceLike) ast.ReferenceLike {
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

func (v *analyzer_) consolidateReferences() {
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

func (v *analyzer_) escapeText(text string) string {
	var escaped string
	for _, character := range text {
		switch character {
		case '"':
			escaped += `\`
		case '.', '|', '^', '$', '+', '*', '?',
			'(', ')', '[', ']', '{', '}':
			escaped += `\\`
		case '\\':
			escaped += `\\\`
		}
		escaped += string(character)
	}
	return escaped
}

func (v *analyzer_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	comment = comment[2 : len(comment)-3]

	// Add in the go style comment delimiters.
	var notice = "/*" + comment + "*/"

	return notice
}

func (v *analyzer_) extractName(syntax ast.SyntaxLike) string {
	var rules = syntax.GetRules().GetIterator()
	// The first rule name is the name of the syntax.
	var rule = rules.GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *analyzer_) pluralizeReference(
	reference ast.ReferenceLike,
) ast.ReferenceLike {
	// Make the identifier plural.
	var identifier = reference.GetIdentifier()
	var name = identifier.GetAny().(string)
	v.plurals_.AddValue(name)

	// Add a plural cardinality to the reference.
	var constraint = ast.Constraint().Make("*")
	var cardinality = ast.Cardinality().Make(constraint)
	reference = ast.Reference().Make(identifier, cardinality)
	return reference
}
