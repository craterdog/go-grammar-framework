/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package grammars

import (
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// This class does not initialize any class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// This class does not define any class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	depth_  int
	result_ sts.Builder
}

// Public

func (v *formatter_) FormatDefinition(definition DefinitionLike) string {
	v.formatDefinition(definition)
	return v.getResult()
}

func (v *formatter_) FormatGrammar(grammar GrammarLike) string {
	v.formatGrammar(grammar)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var separator = "\n"
	for level := 0; level < v.depth_; level++ {
		separator += "    "
	}
	v.appendString(separator)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) formatAlternative(alternative AlternativeLike) {
	var factors = alternative.GetFactors()
	var iterator = factors.GetIterator()
	var factor = iterator.GetNext()
	v.formatFactor(factor)
	for iterator.HasNext() {
		v.appendString(" ")
		factor = iterator.GetNext()
		v.formatFactor(factor)
	}
}

func (v *formatter_) formatCardinality(cardinality CardinalityLike) {
	var constraint = cardinality.GetConstraint()
	var first = constraint.GetFirst()
	var last = constraint.GetLast()
	switch {
	case first == "1" && last == "1":
		// This is the default case so do nothing.
	case first == "0" && last == "1":
		v.appendString("?")
	case first == "0" && len(last) == 0:
		v.appendString("*")
	case first == "1" && len(last) == 0:
		v.appendString("+")
	case len(first) > 0:
		v.formatConstraint(constraint)
	default:
		panic("Attempted to format an invalid cardinality.")
	}
}

func (v *formatter_) formatConstraint(constraint ConstraintLike) {
	var first = constraint.GetFirst()
	var last = constraint.GetLast()
	v.appendString("{")
	v.appendString(first)
	if first != last {
		v.appendString("..")
		if len(last) > 0 {
			v.appendString(last)
		}
	}
	v.appendString("}")
}

func (v *formatter_) formatDefinition(definition DefinitionLike) {
	var comment = definition.GetComment()
	if len(comment) > 0 {
		v.appendString(comment)
	}
	var name = definition.GetName()
	v.appendString(name)
	v.appendString(":")
	var expression = definition.GetExpression()
	if expression.GetInline() != nil {
		v.appendString(" ")
	}
	v.formatExpression(expression)
	v.appendNewline()
	v.appendNewline()
}

func (v *formatter_) formatElement(element ElementLike) {
	var literal = element.GetLiteral()
	var name = element.GetName()
	switch {
	case len(literal) > 0:
		v.appendString(literal)
	case len(name) > 0:
		v.appendString(name)
	default:
		panic("Attempted to format an empty element.")
	}
}

func (v *formatter_) formatExpression(expression ExpressionLike) {
	var inline = expression.GetInline()
	var multiline = expression.GetMultiline()
	switch {
	case inline != nil:
		v.formatInline(inline)
	case multiline != nil:
		v.formatMultiline(multiline)
	default:
		panic("Attempted to format an empty expression.")
	}
}

func (v *formatter_) formatFactor(factor FactorLike) {
	var predicate = factor.GetPredicate()
	v.formatPredicate(predicate)
	var cardinality = factor.GetCardinality()
	if cardinality != nil {
		v.formatCardinality(cardinality)
	}
}

func (v *formatter_) formatFilter(filter FilterLike) {
	var intrinsic = filter.GetIntrinsic()
	var glyph = filter.GetGlyph()
	switch {
	case len(intrinsic) > 0:
		v.appendString(intrinsic)
	case glyph != nil:
		v.formatGlyph(glyph)
	default:
		panic("Attempted to format an empty filter.")
	}
}

func (v *formatter_) formatGlyph(glyph GlyphLike) {
	var first = glyph.GetFirst()
	v.appendString(first)
	var last = glyph.GetLast()
	if len(last) > 0 {
		v.appendString("..")
		v.appendString(last)
	}
}

func (v *formatter_) formatGrammar(grammar GrammarLike) {
	// Format the headers.
	var headers = grammar.GetHeaders()
	var headerIterator = headers.GetIterator()
	var header = headerIterator.GetNext()
	v.formatHeader(header)
	for headerIterator.HasNext() {
		header = headerIterator.GetNext()
		v.formatHeader(header)
	}

	// Format the definitions.
	var definitions = grammar.GetDefinitions()
	var definitionIterator = definitions.GetIterator()
	var definition = definitionIterator.GetNext()
	v.formatDefinition(definition)
	for definitionIterator.HasNext() {
		definition = definitionIterator.GetNext()
		v.formatDefinition(definition)
	}
}

func (v *formatter_) formatHeader(header HeaderLike) {
	var comment = header.GetComment()
	v.appendString(comment)
	v.appendNewline()
}

func (v *formatter_) formatInline(inline InlineLike) {
	var alternatives = inline.GetAlternatives()
	var iterator = alternatives.GetIterator()
	var alternative = iterator.GetNext()
	v.formatAlternative(alternative)
	for iterator.HasNext() {
		v.appendString(" | ")
		alternative = iterator.GetNext()
		v.formatAlternative(alternative)
	}
	var note = inline.GetNote()
	if len(note) > 0 {
		v.appendString("  ")
		v.appendString(note)
	}
}

func (v *formatter_) formatInversion(inversion InversionLike) {
	if inversion.IsInverted() {
		v.appendString("~")
	}
	var filter = inversion.GetFilter()
	v.formatFilter(filter)
}

func (v *formatter_) formatLine(line LineLike) {
	v.appendNewline()
	var alternative = line.GetAlternative()
	v.formatAlternative(alternative)
	var note = line.GetNote()
	if len(note) > 0 {
		v.appendString("  ")
		v.appendString(note)
	}
}

func (v *formatter_) formatMultiline(multiline MultilineLike) {
	v.depth_++
	var lines = multiline.GetLines()
	var iterator = lines.GetIterator()
	for iterator.HasNext() {
		var line = iterator.GetNext()
		v.formatLine(line)
	}
	v.depth_--
}

func (v *formatter_) formatPrecedence(precedence PrecedenceLike) {
	v.appendString("(")
	var expression = precedence.GetExpression()
	v.formatExpression(expression)
	if expression.GetMultiline() != nil {
		v.appendNewline()
	}
	v.appendString(")")
}

func (v *formatter_) formatPredicate(predicate PredicateLike) {
	var element = predicate.GetElement()
	var inversion = predicate.GetInversion()
	var precedence = predicate.GetPrecedence()
	switch {
	case element != nil:
		v.formatElement(element)
	case inversion != nil:
		v.formatInversion(inversion)
	case precedence != nil:
		v.formatPrecedence(precedence)
	default:
		panic("Attempted to format an empty predicate.")
	}
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
