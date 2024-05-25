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
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
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
	return &formatter_{
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	class_  FormatterClassLike
	depth_  int
	result_ sts.Builder
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

// Public

func (v *formatter_) FormatDefinition(definition ast.DefinitionLike) string {
	v.formatDefinition(definition)
	return v.getResult()
}

func (v *formatter_) FormatSyntax(syntax ast.SyntaxLike) string {
	v.formatSyntax(syntax)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var separator = "\n"
	var indentation = "    "
	for level := 0; level < v.depth_; level++ {
		separator += indentation
	}
	v.appendString(separator)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) formatAlternative(alternative ast.AlternativeLike) {
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

func (v *formatter_) formatAtom(atom ast.AtomLike) {
	var glyph = atom.GetGlyph()
	var intrinsic = atom.GetIntrinsic()
	switch {
	case glyph != nil:
		v.formatGlyph(glyph)
	case len(intrinsic) > 0:
		v.appendString(intrinsic)
	default:
		panic("Attempted to format an empty atom.")
	}
}

func (v *formatter_) formatCardinality(cardinality ast.CardinalityLike) {
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
		v.appendString("{")
		v.formatConstraint(constraint)
		v.appendString("}")
	default:
		panic("Attempted to format an invalid cardinality.")
	}
}

func (v *formatter_) formatConstraint(constraint ast.ConstraintLike) {
	var first = constraint.GetFirst()
	var last = constraint.GetLast()
	v.appendString(first)
	if first != last {
		v.appendString("..")
		if len(last) > 0 {
			v.appendString(last)
		}
	}
}

func (v *formatter_) formatDefinition(definition ast.DefinitionLike) {
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

func (v *formatter_) formatElement(element ast.ElementLike) {
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

func (v *formatter_) formatExpression(expression ast.ExpressionLike) {
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

func (v *formatter_) formatFactor(factor ast.FactorLike) {
	var predicate = factor.GetPredicate()
	v.formatPredicate(predicate)
	var cardinality = factor.GetCardinality()
	if cardinality != nil {
		v.formatCardinality(cardinality)
	}
}

func (v *formatter_) formatFilter(filter ast.FilterLike) {
	if filter.IsInverted() {
		v.appendString("~")
	}
	v.appendString("[")
	var atoms = filter.GetAtoms()
	var iterator = atoms.GetIterator()
	var atom = iterator.GetNext()
	v.formatAtom(atom)
	for iterator.HasNext() {
		atom = iterator.GetNext()
		v.appendString(" ")
		v.formatAtom(atom)
	}
	v.appendString("]")
}

func (v *formatter_) formatGlyph(glyph ast.GlyphLike) {
	var first = glyph.GetFirst()
	v.appendString(first)
	var last = glyph.GetLast()
	if len(last) > 0 {
		v.appendString("..")
		v.appendString(last) // The last character may be empty.
	}
}

func (v *formatter_) formatSyntax(syntax ast.SyntaxLike) {
	// Format the headers.
	var headerIterator = syntax.GetHeaders().GetIterator()
	for headerIterator.HasNext() {
		var header = headerIterator.GetNext()
		v.formatHeader(header)
	}

	// Format the definitions.
	var definitionIterator = syntax.GetDefinitions().GetIterator()
	for definitionIterator.HasNext() {
		var definition = definitionIterator.GetNext()
		v.formatDefinition(definition)
	}
}

func (v *formatter_) formatHeader(header ast.HeaderLike) {
	var comment = header.GetComment()
	v.appendString(comment)
	v.appendNewline()
}

func (v *formatter_) formatInline(inline ast.InlineLike) {
	var iterator = inline.GetAlternatives().GetIterator()
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

func (v *formatter_) formatLine(line ast.LineLike) {
	v.appendNewline()
	var alternative = line.GetAlternative()
	v.formatAlternative(alternative)
	var note = line.GetNote()
	if len(note) > 0 {
		v.appendString("  ")
		v.appendString(note)
	}
}

func (v *formatter_) formatMultiline(multiline ast.MultilineLike) {
	v.depth_++
	var iterator = multiline.GetLines().GetIterator()
	for iterator.HasNext() {
		var line = iterator.GetNext()
		v.formatLine(line)
	}
	v.depth_--
}

func (v *formatter_) formatPrecedence(precedence ast.PrecedenceLike) {
	v.appendString("(")
	var expression = precedence.GetExpression()
	v.formatExpression(expression)
	if expression.GetMultiline() != nil {
		v.appendNewline()
	}
	v.appendString(")")
}

func (v *formatter_) formatPredicate(predicate ast.PredicateLike) {
	var atom = predicate.GetAtom()
	var element = predicate.GetElement()
	var filter = predicate.GetFilter()
	var precedence = predicate.GetPrecedence()
	switch {
	case atom != nil:
		v.formatAtom(atom)
	case element != nil:
		v.formatElement(element)
	case filter != nil:
		v.formatFilter(filter)
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
