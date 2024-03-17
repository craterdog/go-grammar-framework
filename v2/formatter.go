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
	var note = alternative.GetNote()
	if len(note) > 0 {
		v.appendString("  ")
		v.appendString(note)
	}
}

func (v *formatter_) formatAssertion(assertion AssertionLike) {
	var element = assertion.GetElement()
	var glyph = assertion.GetGlyph()
	var precedence = assertion.GetPrecedence()
	switch {
	case element != nil:
		v.formatElement(element)
	case glyph != nil:
		v.formatGlyph(glyph)
	case precedence != nil:
		v.formatPrecedence(precedence)
	default:
		panic("Attempted to format an empty assertion.")
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
	var symbol = definition.GetSymbol()
	v.appendString(symbol)
	v.appendString(":")
	var expression = definition.GetExpression()
	if !expression.IsMultilined() {
		v.appendString(" ")
	}
	v.formatExpression(expression)
}

func (v *formatter_) formatElement(element ElementLike) {
	var intrinsic = element.GetIntrinsic()
	var name = element.GetName()
	var literal = element.GetLiteral()
	switch {
	case len(intrinsic) > 0:
		v.appendString(intrinsic)
	case len(name) > 0:
		v.appendString(name)
	case len(literal) > 0:
		v.appendString(literal)
	default:
		panic("Attempted to format an empty element.")
	}
}

func (v *formatter_) formatExpression(expression ExpressionLike) {
	var alternative AlternativeLike
	var alternatives = expression.GetAlternatives()
	var iterator = alternatives.GetIterator()
	if expression.IsMultilined() {
		v.depth_++
		for iterator.HasNext() {
			v.appendNewline()
			alternative = iterator.GetNext()
			v.formatAlternative(alternative)
		}
		v.depth_--
	} else {
		alternative = iterator.GetNext()
		v.formatAlternative(alternative)
		for iterator.HasNext() {
			v.appendString(" | ")
			alternative = iterator.GetNext()
			v.formatAlternative(alternative)
		}
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
	var comment = grammar.GetComment()
	v.appendString(comment)
	var statements = grammar.GetStatements()
	var iterator = statements.GetIterator()
	for iterator.HasNext() {
		var statement = iterator.GetNext()
		v.formatStatement(statement)
		v.appendNewline()
		v.appendNewline()
	}
}

func (v *formatter_) formatPrecedence(precedence PrecedenceLike) {
	v.appendString("(")
	var expression = precedence.GetExpression()
	v.formatExpression(expression)
	if expression.IsMultilined() {
		v.appendNewline()
	}
	v.appendString(")")
}

func (v *formatter_) formatPredicate(predicate PredicateLike) {
	var assertion = predicate.GetAssertion()
	if predicate.IsInverted() {
		v.appendString("~")
	}
	v.formatAssertion(assertion)
}

func (v *formatter_) formatStatement(statement StatementLike) {
	var comment = statement.GetComment()
	v.appendString(comment)
	var definition = statement.GetDefinition()
	v.formatDefinition(definition)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
