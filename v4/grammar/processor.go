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
)

// CLASS ACCESS

// Reference

var processorClass = &processorClass_{
	// Initialize the class constants.
}

// Function

func Processor() ProcessorClassLike {
	return processorClass
}

// CLASS METHODS

// Target

type processorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *processorClass_) Make() ProcessorLike {
	var processor = &processor_{
		// Initialize the instance attributes.
		class_: c,
	}
	return processor
}

// INSTANCE METHODS

// Target

type processor_ struct {
	// Define the instance attributes.
	class_ ProcessorClassLike
}

// Attributes

func (v *processor_) GetClass() ProcessorClassLike {
	return v.class_
}

// Methodical

func (v *processor_) ProcessComment(comment string) {
}

func (v *processor_) ProcessExcluded(excluded string) {
}

func (v *processor_) ProcessIntrinsic(intrinsic string) {
}

func (v *processor_) ProcessLiteral(literal string) {
}

func (v *processor_) ProcessLowercase(lowercase string) {
}

func (v *processor_) ProcessNewline(
	newline string,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessNote(note string) {
}

func (v *processor_) ProcessNumber(
	number string,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessOptional(optional string) {
}

func (v *processor_) ProcessRepeated(repeated string) {
}

func (v *processor_) ProcessRunic(
	runic string,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessUppercase(uppercase string) {
}

func (v *processor_) PreprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PostprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessBracket(bracket ast.BracketLike) {
}

func (v *processor_) PostprocessBracket(bracket ast.BracketLike) {
}

func (v *processor_) PreprocessCardinality(cardinality ast.CardinalityLike) {
}

func (v *processor_) PostprocessCardinality(cardinality ast.CardinalityLike) {
}

func (v *processor_) PreprocessCharacter(
	character ast.CharacterLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PostprocessCharacter(
	character ast.CharacterLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessConstraint(constraint ast.ConstraintLike) {
}

func (v *processor_) PostprocessConstraint(constraint ast.ConstraintLike) {
}

func (v *processor_) PreprocessCount(count ast.CountLike) {
}

func (v *processor_) PostprocessCount(count ast.CountLike) {
}

func (v *processor_) PreprocessDefinition(definition ast.DefinitionLike) {
}

func (v *processor_) PostprocessDefinition(definition ast.DefinitionLike) {
}

func (v *processor_) PreprocessElement(element ast.ElementLike) {
}

func (v *processor_) PostprocessElement(element ast.ElementLike) {
}

func (v *processor_) PreprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PostprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessFactor(
	factor ast.FactorLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PostprocessFactor(
	factor ast.FactorLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessFilter(filter ast.FilterLike) {
}

func (v *processor_) PostprocessFilter(filter ast.FilterLike) {
}

func (v *processor_) PreprocessGroup(group ast.GroupLike) {
}

func (v *processor_) PostprocessGroup(group ast.GroupLike) {
}

func (v *processor_) PreprocessHeader(
	header ast.HeaderLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PostprocessHeader(
	header ast.HeaderLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessIdentifier(identifier ast.IdentifierLike) {
}

func (v *processor_) PostprocessIdentifier(identifier ast.IdentifierLike) {
}

func (v *processor_) PreprocessInline(inline ast.InlineLike) {
}

func (v *processor_) PostprocessInline(inline ast.InlineLike) {
}

func (v *processor_) PreprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PostprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessMultiline(multiline ast.MultilineLike) {
}

func (v *processor_) PostprocessMultiline(multiline ast.MultilineLike) {
}

func (v *processor_) PreprocessPattern(pattern ast.PatternLike) {
}

func (v *processor_) PostprocessPattern(pattern ast.PatternLike) {
}

func (v *processor_) PreprocessReference(reference ast.ReferenceLike) {
}

func (v *processor_) PostprocessReference(reference ast.ReferenceLike) {
}

func (v *processor_) PreprocessRepetition(
	repetition ast.RepetitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PostprocessRepetition(
	repetition ast.RepetitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PostprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessSpecific(specific ast.SpecificLike) {
}

func (v *processor_) PostprocessSpecific(specific ast.SpecificLike) {
}

func (v *processor_) PreprocessSyntax(syntax ast.SyntaxLike) {
}

func (v *processor_) PostprocessSyntax(syntax ast.SyntaxLike) {
}

func (v *processor_) PreprocessTerm(
	term ast.TermLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PostprocessTerm(
	term ast.TermLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessText(text ast.TextLike) {
}

func (v *processor_) PostprocessText(text ast.TextLike) {
}
