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

package agent

import (
	col "github.com/craterdog/go-collection-framework/v4"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
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
	return &formatter_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_  FormatterClassLike
	depth_  int
	result_ sts.Builder
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) GetDepth() int {
	return v.depth_
}

// Public

func (v *formatter_) FormatSyntax(syntax ast.SyntaxLike) string {
	v.formatSyntax(syntax)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "    "
	for level := 0; level < v.depth_; level++ {
		newline += indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) formatAlternative(alternative ast.AlternativeLike) {
	v.appendString("|")
	var iterator = alternative.GetParts().GetIterator()
	for iterator.HasNext() {
		var part = iterator.GetNext()
		v.appendString(" ")
		v.formatPart(part)
	}
}

func (v *formatter_) formatBounded(bounded ast.BoundedLike) {
	var rune_ = bounded.GetRune()
	v.appendString(rune_)
	var extent = bounded.GetOptionalExtent()
	if col.IsDefined(extent) {
		v.formatExtent(extent)
	}
}

func (v *formatter_) formatCardinality(cardinality ast.CardinalityLike) {
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		v.formatConstrained(actual)
	case string:
		v.appendString(actual)
	default:
		panic("Attempted to format an empty cardinality.")
	}
}

func (v *formatter_) formatCharacter(character ast.CharacterLike) {
	switch actual := character.GetAny().(type) {
	case ast.BoundedLike:
		v.formatBounded(actual)
	case string:
		v.appendString(actual)
	default:
		panic("Attempted to format an empty character.")
	}
}

func (v *formatter_) formatConstrained(constrained ast.ConstrainedLike) {
	v.appendString("{")
	var number = constrained.GetNumber()
	v.appendString(number)
	var limit = constrained.GetOptionalLimit()
	if col.IsDefined(limit) {
		v.formatLimit(limit)
	}
	v.appendString("}")
}

func (v *formatter_) formatDefinition(definition ast.DefinitionLike) {
	switch actual := definition.GetAny().(type) {
	case ast.InlinedLike:
		v.formatInlined(actual)
	case ast.MultilinedLike:
		v.formatMultilined(actual)
	default:
		panic("Attempted to format an empty definition.")
	}
}

func (v *formatter_) formatElement(element ast.ElementLike) {
	switch actual := element.GetAny().(type) {
	case ast.GroupedLike:
		v.formatGrouped(actual)
	case ast.FilteredLike:
		v.formatFiltered(actual)
	case ast.StringLike:
		v.formatString(actual)
	default:
		panic("Attempted to format an empty element.")
	}
}

func (v *formatter_) formatExpression(expression ast.ExpressionLike) {
	var comment = expression.GetOptionalComment()
	if col.IsDefined(comment) {
		v.appendString(comment)
	}
	var lower = expression.GetLowercase()
	v.appendString(lower)
	v.appendString(": ")
	var pattern = expression.GetPattern()
	v.formatPattern(pattern)
	var note = expression.GetOptionalNote()
	if col.IsDefined(note) {
		v.appendString("  ")
		v.appendString(note)
	}
	v.appendNewline()
	v.appendNewline()
}

func (v *formatter_) formatExtent(extent ast.ExtentLike) {
	v.appendString("..")
	var rune_ = extent.GetRune()
	v.appendString(rune_)
}

func (v *formatter_) formatFactor(factor ast.FactorLike) {
	var predicate = factor.GetPredicate()
	v.formatPredicate(predicate)
	var cardinality = factor.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		v.formatCardinality(cardinality)
	}
}

func (v *formatter_) formatFiltered(filtered ast.FilteredLike) {
	var negation = filtered.GetOptionalNegation()
	if col.IsDefined(negation) {
		v.appendString(negation)
	}
	v.appendString("[")
	var iterator = filtered.GetCharacters().GetIterator()
	var character = iterator.GetNext()
	v.formatCharacter(character) // The first one is not prepended with a space.
	for iterator.HasNext() {
		character = iterator.GetNext()
		v.appendString(" ")
		v.formatCharacter(character)
	}
	v.appendString("]")
}

func (v *formatter_) formatGrouped(grouped ast.GroupedLike) {
	v.appendString("(")
	var pattern = grouped.GetPattern()
	v.formatPattern(pattern)
	v.appendString(")")
}

func (v *formatter_) formatHeader(header ast.HeaderLike) {
	var comment = header.GetComment()
	v.appendString(comment)
	v.appendNewline()
}

func (v *formatter_) formatIdentifier(identifier ast.IdentifierLike) {
	switch actual := identifier.GetAny().(type) {
	case string:
		v.appendString(actual)
	default:
		panic("Attempted to format an empty identifier.")
	}
}

func (v *formatter_) formatInlined(inlined ast.InlinedLike) {
	var iterator = inlined.GetFactors().GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		v.appendString(" ")
		v.formatFactor(factor)
	}
	var note = inlined.GetOptionalNote()
	if col.IsDefined(note) {
		v.appendString("  ")
		v.appendString(note)
	}
}

func (v *formatter_) formatLine(line ast.LineLike) {
	v.appendNewline()
	var identifier = line.GetIdentifier()
	v.formatIdentifier(identifier)
	var note = line.GetOptionalNote()
	if col.IsDefined(note) {
		v.appendString("  ")
		v.appendString(note)
	}
}

func (v *formatter_) formatLimit(limit ast.LimitLike) {
	v.appendString("..")
	var number = limit.GetOptionalNumber()
	if col.IsDefined(number) {
		v.appendString(number)
	}
}

func (v *formatter_) formatMultilined(multilined ast.MultilinedLike) {
	var iterator = multilined.GetLines().GetIterator()
	v.depth_++
	for iterator.HasNext() {
		var line = iterator.GetNext()
		v.formatLine(line)
	}
	v.depth_--
}

func (v *formatter_) formatPart(part ast.PartLike) {
	var element = part.GetElement()
	v.formatElement(element)
	var cardinality = part.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		v.formatCardinality(cardinality)
	}
}

func (v *formatter_) formatPattern(pattern ast.PatternLike) {
	var partIterator = pattern.GetParts().GetIterator()
	var part = partIterator.GetNext()
	v.formatPart(part) // The first one is not prepended with a space.
	for partIterator.HasNext() {
		part = partIterator.GetNext()
		v.appendString(" ")
		v.formatPart(part)
	}
	var alternativeIterator = pattern.GetAlternatives().GetIterator()
	for alternativeIterator.HasNext() {
		var alternative = alternativeIterator.GetNext()
		v.appendString(" ")
		v.formatAlternative(alternative)
	}
}

func (v *formatter_) formatPredicate(predicate ast.PredicateLike) {
	switch actual := predicate.GetAny().(type) {
	case string:
		v.appendString(actual)
	default:
		panic("Attempted to format an empty predicate.")
	}
}

func (v *formatter_) formatRule(rule ast.RuleLike) {
	var comment = rule.GetOptionalComment()
	if col.IsDefined(comment) {
		v.appendString(comment)
	}
	var upper = rule.GetUppercase()
	v.appendString(upper)
	v.appendString(":")
	var definition = rule.GetDefinition()
	v.formatDefinition(definition)
	v.appendNewline()
	v.appendNewline()
}

func (v *formatter_) formatString(string_ ast.StringLike) {
	switch actual := string_.GetAny().(type) {
	case string:
		v.appendString(actual)
	default:
		panic("Attempted to format an empty string.")
	}
}

func (v *formatter_) formatSyntax(syntax ast.SyntaxLike) {
	// Format the headers.
	var headerIterator = syntax.GetHeaders().GetIterator()
	for headerIterator.HasNext() {
		var header = headerIterator.GetNext()
		v.formatHeader(header)
	}

	// Format the rules.
	var ruleIterator = syntax.GetRules().GetIterator()
	for ruleIterator.HasNext() {
		var rule = ruleIterator.GetNext()
		v.formatRule(rule)
	}

	// Format the expressions.
	var expressionIterator = syntax.GetExpressions().GetIterator()
	for expressionIterator.HasNext() {
		var expression = expressionIterator.GetNext()
		v.formatExpression(expression)
	}
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
