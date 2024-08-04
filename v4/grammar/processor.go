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

func (v *processor_) ProcessGlyph(glyph string) {
}

func (v *processor_) ProcessIntrinsic(intrinsic string) {
}

func (v *processor_) ProcessLiteral(literal string) {
}

func (v *processor_) ProcessLowercase(lowercase string) {
}

func (v *processor_) ProcessNegation(negation string) {
}

func (v *processor_) ProcessNewline(
	newline string,
	index uint,
) {
}

func (v *processor_) ProcessNote(note string) {
}

func (v *processor_) ProcessNumber(number string) {
}

func (v *processor_) ProcessQuantified(quantified string) {
}

func (v *processor_) ProcessReserved(reserved string) {
}

func (v *processor_) ProcessUppercase(uppercase string) {
}

func (v *processor_) PreprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
) {
}

func (v *processor_) PostprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
) {
}

func (v *processor_) PreprocessBounded(bounded ast.BoundedLike) {
}

func (v *processor_) PostprocessBounded(bounded ast.BoundedLike) {
}

func (v *processor_) PreprocessCardinality(cardinality ast.CardinalityLike) {
}

func (v *processor_) PostprocessCardinality(cardinality ast.CardinalityLike) {
}

func (v *processor_) PreprocessCharacter(
	character ast.CharacterLike,
	index uint,
) {
}

func (v *processor_) PostprocessCharacter(
	character ast.CharacterLike,
	index uint,
) {
}

func (v *processor_) PreprocessConstrained(constrained ast.ConstrainedLike) {
}

func (v *processor_) PostprocessConstrained(constrained ast.ConstrainedLike) {
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
) {
}

func (v *processor_) PostprocessExpression(
	expression ast.ExpressionLike,
	index uint,
) {
}

func (v *processor_) PreprocessExtent(extent ast.ExtentLike) {
}

func (v *processor_) PostprocessExtent(extent ast.ExtentLike) {
}

func (v *processor_) PreprocessFactor(
	factor ast.FactorLike,
	index uint,
) {
}

func (v *processor_) PostprocessFactor(
	factor ast.FactorLike,
	index uint,
) {
}

func (v *processor_) PreprocessFiltered(filtered ast.FilteredLike) {
}

func (v *processor_) PostprocessFiltered(filtered ast.FilteredLike) {
}

func (v *processor_) PreprocessGrouped(grouped ast.GroupedLike) {
}

func (v *processor_) PostprocessGrouped(grouped ast.GroupedLike) {
}

func (v *processor_) PreprocessHeader(
	header ast.HeaderLike,
	index uint,
) {
}

func (v *processor_) PostprocessHeader(
	header ast.HeaderLike,
	index uint,
) {
}

func (v *processor_) PreprocessIdentifier(identifier ast.IdentifierLike) {
}

func (v *processor_) PostprocessIdentifier(identifier ast.IdentifierLike) {
}

func (v *processor_) PreprocessInlined(inlined ast.InlinedLike) {
}

func (v *processor_) PostprocessInlined(inlined ast.InlinedLike) {
}

func (v *processor_) PreprocessLimit(limit ast.LimitLike) {
}

func (v *processor_) PostprocessLimit(limit ast.LimitLike) {
}

func (v *processor_) PreprocessLine(
	line ast.LineLike,
	index uint,
) {
}

func (v *processor_) PostprocessLine(
	line ast.LineLike,
	index uint,
) {
}

func (v *processor_) PreprocessMultilined(multilined ast.MultilinedLike) {
}

func (v *processor_) PostprocessMultilined(multilined ast.MultilinedLike) {
}

func (v *processor_) PreprocessPart(
	part ast.PartLike,
	index uint,
) {
}

func (v *processor_) PostprocessPart(
	part ast.PartLike,
	index uint,
) {
}

func (v *processor_) PreprocessPattern(pattern ast.PatternLike) {
}

func (v *processor_) PostprocessPattern(pattern ast.PatternLike) {
}

func (v *processor_) PreprocessPredicate(predicate ast.PredicateLike) {
}

func (v *processor_) PostprocessPredicate(predicate ast.PredicateLike) {
}

func (v *processor_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
) {
}

func (v *processor_) PostprocessRule(
	rule ast.RuleLike,
	index uint,
) {
}

func (v *processor_) PreprocessSelective(selective ast.SelectiveLike) {
}

func (v *processor_) PostprocessSelective(selective ast.SelectiveLike) {
}

func (v *processor_) PreprocessSequential(sequential ast.SequentialLike) {
}

func (v *processor_) PostprocessSequential(sequential ast.SequentialLike) {
}

func (v *processor_) PreprocessSupplement(syntax ast.SupplementLike) {
}

func (v *processor_) PostprocessSupplement(syntax ast.SupplementLike) {
}

func (v *processor_) PreprocessSyntax(syntax ast.SyntaxLike) {
}

func (v *processor_) PostprocessSyntax(syntax ast.SyntaxLike) {
}

func (v *processor_) PreprocessTextual(textual ast.TextualLike) {
}

func (v *processor_) PostprocessTextual(textual ast.TextualLike) {
}
