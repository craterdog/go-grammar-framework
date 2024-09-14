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
	var formatter = &formatter_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	formatter.visitor_ = Visitor().Make(formatter)
	return formatter
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_   FormatterClassLike
	visitor_ VisitorLike
	inGroup_ bool
	result_  sts.Builder

	// Define the inherited aspects.
	Methodical
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

// Methodical

func (v *formatter_) ProcessComment(comment string) {
	v.appendString(comment)
}

func (v *formatter_) ProcessGlyph(glyph string) {
	v.appendString(glyph)
}

func (v *formatter_) ProcessIntrinsic(intrinsic string) {
	v.appendString(intrinsic)
}

func (v *formatter_) ProcessLiteral(literal string) {
	v.appendString(literal)
}

func (v *formatter_) ProcessLowercase(lowercase string) {
	v.appendString(lowercase)
}

func (v *formatter_) ProcessNote(note string) {
	v.appendString("  ")
	v.appendString(note)
}

func (v *formatter_) ProcessNumber(number string) {
	v.appendString(number)
}

func (v *formatter_) ProcessOptional(optional string) {
	v.appendString(optional)
}

func (v *formatter_) ProcessRepeated(repeated string) {
	v.appendString(repeated)
}

func (v *formatter_) ProcessUppercase(uppercase string) {
	v.appendString(uppercase)
}

func (v *formatter_) PreprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
	size uint,
) {
	v.appendString(" | ")
}

func (v *formatter_) PreprocessCharacter(
	character ast.CharacterLike,
	index uint,
	size uint,
) {
	if index > 1 {
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessDefinition(definition ast.DefinitionLike) {
	v.appendString(":")
}

func (v *formatter_) PostprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
	v.appendNewline()
}

func (v *formatter_) PreprocessExtent(extent ast.ExtentLike) {
	v.appendString("..")
}

func (v *formatter_) PreprocessFilter(filter ast.FilterLike) {
	var excluded = filter.GetOptionalExcluded()
	if col.IsDefined(excluded) {
		v.appendString(excluded)
	}
	v.appendString("[")
}

func (v *formatter_) PostprocessFilter(filter ast.FilterLike) {
	v.appendString("]")
}

func (v *formatter_) PreprocessGroup(group ast.GroupLike) {
	v.appendString("(")
	v.inGroup_ = true
}

func (v *formatter_) PostprocessGroup(group ast.GroupLike) {
	v.appendString(")")
	v.inGroup_ = false
}

func (v *formatter_) PostprocessInline(inline ast.InlineLike) {
	v.appendNewline()
}

func (v *formatter_) PreprocessLimit(limit ast.LimitLike) {
	v.appendString("..")
}

func (v *formatter_) PreprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
	v.appendString("  - ")
}

func (v *formatter_) PostprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessMultiline(multiline ast.MultilineLike) {
	v.appendNewline()
}

func (v *formatter_) PostprocessNotice(notice ast.NoticeLike) {
	v.appendNewline()
}

func (v *formatter_) PreprocessPattern(pattern ast.PatternLike) {
	if !v.inGroup_ {
		v.appendString(": ")
	}
}

func (v *formatter_) PreprocessQuantified(quantified ast.QuantifiedLike) {
	v.appendString("{")
}

func (v *formatter_) PostprocessQuantified(quantified ast.QuantifiedLike) {
	v.appendString("}")
}

func (v *formatter_) PreprocessRepetition(
	repetition ast.RepetitionLike,
	index uint,
	size uint,
) {
	if index > 1 {
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessTerm(
	term ast.TermLike,
	index uint,
	size uint,
) {
	v.appendString(" ")
}

// Public

func (v *formatter_) FormatSyntax(syntax ast.SyntaxLike) string {
	v.visitor_.VisitSyntax(syntax)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var newline = "\n"
	v.appendString(newline)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
