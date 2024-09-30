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
	class_   *formatterClass_
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Define the inherited aspects.
	Methodical
}

// Public

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) FormatSyntax(syntax ast.SyntaxLike) string {
	v.visitor_.VisitSyntax(syntax)
	return v.getResult()
}

// Methodical

func (v *formatter_) ProcessComment(comment string) {
	v.appendString(comment)
}

func (v *formatter_) ProcessExcluded(excluded string) {
	v.appendString(excluded)
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

func (v *formatter_) ProcessNewline(
	newline string,
	index uint,
	size uint,
) {
	v.appendString(newline)
}

func (v *formatter_) ProcessNote(note string) {
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

func (v *formatter_) ProcessSpace(space string) {
	v.appendString(space)
}

func (v *formatter_) ProcessUppercase(uppercase string) {
	v.appendString(uppercase)
}

func (v *formatter_) PreprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAlternativeSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessCardinality(cardinality ast.CardinalityLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessCardinalitySlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessCardinality(cardinality ast.CardinalityLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessCharacter(
	character ast.CharacterLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessCharacterSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessCharacter(
	character ast.CharacterLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessConstrained(constrained ast.ConstrainedLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessConstrainedSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessConstrained(constrained ast.ConstrainedLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessDefinition(definition ast.DefinitionLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessDefinitionSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessDefinition(definition ast.DefinitionLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessElement(element ast.ElementLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessElementSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessElement(element ast.ElementLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessExplicit(explicit ast.ExplicitLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessExplicitSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessExplicit(explicit ast.ExplicitLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessExpressionSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessExtent(extent ast.ExtentLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessExtentSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessExtent(extent ast.ExtentLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessFilter(filter ast.FilterLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessFilterSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessFilter(filter ast.FilterLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessGroup(group ast.GroupLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessGroupSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessGroup(group ast.GroupLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessIdentifier(identifier ast.IdentifierLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessIdentifierSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessIdentifier(identifier ast.IdentifierLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInline(inline ast.InlineLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInlineSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInline(inline ast.InlineLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessLimit(limit ast.LimitLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessLimitSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessLimit(limit ast.LimitLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessLineSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMultiline(multiline ast.MultilineLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMultilineSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMultiline(multiline ast.MultilineLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessNotice(notice ast.NoticeLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessNoticeSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessNotice(notice ast.NoticeLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessOption(option ast.OptionLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessOptionSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessOption(option ast.OptionLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessPattern(pattern ast.PatternLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPatternSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPattern(pattern ast.PatternLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessQuantified(quantified ast.QuantifiedLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessQuantifiedSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessQuantified(quantified ast.QuantifiedLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessReference(reference ast.ReferenceLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessReferenceSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessReference(reference ast.ReferenceLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessRepetition(
	repetition ast.RepetitionLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessRepetitionSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessRepetition(
	repetition ast.RepetitionLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessRuleSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessSyntax(syntax ast.SyntaxLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessSyntaxSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessSyntax(syntax ast.SyntaxLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessTerm(
	term ast.TermLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessTermSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessTerm(
	term ast.TermLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessText(text ast.TextLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessTextSlot(slot uint) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessText(text ast.TextLike) {
	// TBD - Add formatting of the delimited rule.
}

// Private

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "\t"
	var level uint
	for ; level < v.depth_; level++ {
		newline += indentation
	}
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
