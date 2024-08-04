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
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
)

// CLASS ACCESS

// Reference

var visitorClass = &visitorClass_{
	// Initialize the class constants.
}

// Function

func Visitor() VisitorClassLike {
	return visitorClass
}

// CLASS METHODS

// Target

type visitorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *visitorClass_) Make(
	processor Methodical,
) VisitorLike {
	return &visitor_{
		// Initialize the instance attributes.
		class_:     c,
		processor_: processor,
	}
}

// INSTANCE METHODS

// Target

type visitor_ struct {
	// Define the instance attributes.
	class_     VisitorClassLike
	processor_ Methodical
}

// Attributes

func (v *visitor_) GetClass() VisitorClassLike {
	return v.class_
}

func (v *visitor_) GetProcessor() Methodical {
	return v.processor_
}

// Public

func (v *visitor_) VisitSyntax(syntax ast.SyntaxLike) {
	// Visit the syntax.
	v.processor_.PreprocessSyntax(syntax)
	v.visitSyntax(syntax)
	v.processor_.PostprocessSyntax(syntax)
}

// Private

func (v *visitor_) visitAlternative(
	alternative ast.AlternativeLike,
) {
	// Visit the separator.
	var separator = alternative.GetSeparator()
	v.processor_.ProcessSeparator(separator)

	// Visit the part.
	var part = alternative.GetPart()
	v.processor_.PreprocessPart(part, 2)
	v.visitPart(part)
	v.processor_.PostprocessPart(part, 2)
}

func (v *visitor_) visitBounded(
	bounded ast.BoundedLike,
) {
	// Visit the initial glyph.
	var glyph = bounded.GetGlyph()
	v.processor_.ProcessGlyph(glyph)

	// Visit the optional extent glyph.
	var extent = bounded.GetOptionalExtent()
	if col.IsDefined(extent) {
		v.processor_.PreprocessExtent(extent)
		v.visitExtent(extent)
		v.processor_.PostprocessExtent(extent)
	}
}

func (v *visitor_) visitCardinality(
	cardinality ast.CardinalityLike,
) {
	// Visit the possible cardinality types.
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		v.processor_.PreprocessConstrained(actual)
		v.visitConstrained(actual)
		v.processor_.PostprocessConstrained(actual)
	case string:
		v.processor_.ProcessQuantified(actual)
	default:
		panic("A cardinality must have a constrained or quantified value.")
	}
}

func (v *visitor_) visitCharacter(
	character ast.CharacterLike,
) {
	// Visit the possible character types.
	switch actual := character.GetAny().(type) {
	case ast.BoundedLike:
		v.processor_.PreprocessBounded(actual)
		v.visitBounded(actual)
		v.processor_.PostprocessBounded(actual)
	case string:
		v.processor_.ProcessIntrinsic(actual)
	default:
		panic("An character must have a bounded or intrinsic value.")
	}
}

func (v *visitor_) visitConstrained(
	constrained ast.ConstrainedLike,
) {
	// Visit the opening separator.
	var separator = constrained.GetSeparator()
	v.processor_.ProcessSeparator(separator)

	// Visit the minimum value.
	var number = constrained.GetNumber()
	v.processor_.ProcessNumber(number)

	// Visit the optional limit value.
	var limit = constrained.GetOptionalLimit()
	if col.IsDefined(limit) {
		v.processor_.PreprocessLimit(limit)
		v.visitLimit(limit)
		v.processor_.PostprocessLimit(limit)
	}

	// Visit the closing separator.
	separator = constrained.GetSeparator2()
	v.processor_.ProcessSeparator(separator)
}

func (v *visitor_) visitDefinition(
	definition ast.DefinitionLike,
) {
	// Visit the possible definition types.
	switch actual := definition.GetAny().(type) {
	case ast.InlinedLike:
		v.processor_.PreprocessInlined(actual)
		v.visitInlined(actual)
		v.processor_.PostprocessInlined(actual)
	case ast.MultilinedLike:
		v.processor_.PreprocessMultilined(actual)
		v.visitMultilined(actual)
		v.processor_.PostprocessMultilined(actual)
	default:
		panic("An definition must have an inline or multiline value.")
	}
}

func (v *visitor_) visitElement(
	element ast.ElementLike,
) {
	// Visit the possible element types.
	switch actual := element.GetAny().(type) {
	case ast.GroupedLike:
		v.processor_.PreprocessGrouped(actual)
		v.visitGrouped(actual)
		v.processor_.PostprocessGrouped(actual)
	case ast.FilteredLike:
		v.processor_.PreprocessFiltered(actual)
		v.visitFiltered(actual)
		v.processor_.PostprocessFiltered(actual)
	case ast.TextualLike:
		v.processor_.PreprocessTextual(actual)
		v.visitTextual(actual)
		v.processor_.PostprocessTextual(actual)
	default:
		panic("An element must have a grouped, filtered or textual value.")
	}
}

func (v *visitor_) visitExpression(
	expression ast.ExpressionLike,
) {
	// Visit the optional comment.
	var comment = expression.GetOptionalComment()
	if col.IsDefined(comment) {
		v.processor_.ProcessComment(comment)
	}

	// Visit the lowercase identifier.
	var lowercase = expression.GetLowercase()
	v.processor_.ProcessLowercase(lowercase)

	// Visit the separator.
	var separator = expression.GetSeparator()
	v.processor_.ProcessSeparator(separator)

	// Visit the pattern.
	var pattern = expression.GetPattern()
	v.processor_.PreprocessPattern(pattern)
	v.visitPattern(pattern)
	v.processor_.PostprocessPattern(pattern)

	// Visit the optional note.
	var note = expression.GetOptionalNote()
	if col.IsDefined(note) {
		v.processor_.ProcessNote(note)
	}

	// Visit each newline.
	var index uint
	var newlines = expression.GetNewlines().GetIterator()
	for newlines.HasNext() {
		index++
		var newline = newlines.GetNext()
		v.processor_.ProcessNewline(newline, index)
	}
}

func (v *visitor_) visitExtent(
	extent ast.ExtentLike,
) {
	// Visit the separator.
	var separator = extent.GetSeparator()
	v.processor_.ProcessSeparator(separator)

	// Visit the glyph.
	var glyph = extent.GetGlyph()
	v.processor_.ProcessGlyph(glyph)
}

func (v *visitor_) visitFactor(
	factor ast.FactorLike,
) {
	// Visit the possible factor types.
	switch actual := factor.GetAny().(type) {
	case ast.PredicateLike:
		v.processor_.PreprocessPredicate(actual)
		v.visitPredicate(actual)
		v.processor_.PostprocessPredicate(actual)
	case string:
		var string_ = factor.GetAny().(string)
		v.processor_.ProcessLiteral(string_)
	default:
		var message = fmt.Sprintf(
			"An invalid factor type was found: %T",
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitFiltered(
	filtered ast.FilteredLike,
) {
	// Visit the optional negation.
	var negation = filtered.GetOptionalNegation()
	if col.IsDefined(negation) {
		v.processor_.ProcessNegation(negation)
	}

	// Visit the opening separator.
	var separator = filtered.GetSeparator()
	v.processor_.ProcessSeparator(separator)

	// Visit each character.
	var index uint
	var characters = filtered.GetCharacters().GetIterator()
	for characters.HasNext() {
		index++
		var character = characters.GetNext()
		v.processor_.PreprocessCharacter(character, index)
		v.visitCharacter(character)
		v.processor_.PostprocessCharacter(character, index)
	}

	// Visit the closing separator.
	separator = filtered.GetSeparator2()
	v.processor_.ProcessSeparator(separator)
}

func (v *visitor_) visitGrouped(
	grouped ast.GroupedLike,
) {
	// Visit the opening separator.
	var separator = grouped.GetSeparator()
	v.processor_.ProcessSeparator(separator)

	// Visit the pattern.
	var pattern = grouped.GetPattern()
	v.processor_.PreprocessPattern(pattern)
	v.visitPattern(pattern)
	v.processor_.PostprocessPattern(pattern)

	// Visit the closing separator.
	separator = grouped.GetSeparator2()
	v.processor_.ProcessSeparator(separator)
}

func (v *visitor_) visitHeader(
	header ast.HeaderLike,
) {
	// Visit the comment.
	var comment = header.GetComment()
	v.processor_.ProcessComment(comment)

	// Visit the newline.
	var newline = header.GetNewline()
	v.processor_.ProcessNewline(newline, 1)
}

func (v *visitor_) visitIdentifier(
	identifier ast.IdentifierLike,
) {
	// Visit the possible identifier types.
	var string_ = identifier.GetAny().(string)
	switch {
	case Scanner().MatchesType(string_, LowercaseToken):
		v.processor_.ProcessLowercase(string_)
	case Scanner().MatchesType(string_, UppercaseToken):
		v.processor_.ProcessUppercase(string_)
	default:
		var message = fmt.Sprintf(
			"An invalid identifier was found: %v",
			string_,
		)
		panic(message)
	}
}

func (v *visitor_) visitInlined(
	inlined ast.InlinedLike,
) {
	// Visit each factor.
	var index uint
	var factors = inlined.GetFactors().GetIterator()
	for factors.HasNext() {
		index++
		var factor = factors.GetNext()
		v.processor_.PreprocessFactor(factor, index)
		v.visitFactor(factor)
		v.processor_.PostprocessFactor(factor, index)
	}

	// Visit the optional note.
	var note = inlined.GetOptionalNote()
	if col.IsDefined(note) {
		v.processor_.ProcessNote(note)
	}
}

func (v *visitor_) visitLine(
	line ast.LineLike,
) {
	// Visit the newline.
	var newline = line.GetNewline()
	v.processor_.ProcessNewline(newline, 1)

	// Visit the identifier.
	var identifier = line.GetIdentifier()
	v.processor_.PreprocessIdentifier(identifier)
	v.visitIdentifier(identifier)
	v.processor_.PostprocessIdentifier(identifier)

	// Visit the optional note.
	var note = line.GetOptionalNote()
	if col.IsDefined(note) {
		v.processor_.ProcessNote(note)
	}
}

func (v *visitor_) visitLimit(
	limit ast.LimitLike,
) {
	// Visit the separator.
	var separator = limit.GetSeparator()
	v.processor_.ProcessSeparator(separator)

	// Visit the optional number.
	var number = limit.GetOptionalNumber()
	if col.IsDefined(number) {
		v.processor_.ProcessNumber(number)
	}
}

func (v *visitor_) visitMultilined(
	multilined ast.MultilinedLike,
) {
	// Visit each line.
	var index uint
	var lines = multilined.GetLines().GetIterator()
	for lines.HasNext() {
		index++
		var line = lines.GetNext()
		v.processor_.PreprocessLine(line, index)
		v.visitLine(line)
		v.processor_.PostprocessLine(line, index)
	}
}

func (v *visitor_) visitPart(
	part ast.PartLike,
) {
	// Visit the element.
	var element = part.GetElement()
	v.processor_.PreprocessElement(element)
	v.visitElement(element)
	v.processor_.PostprocessElement(element)

	// Visit the optional cardinality.
	var cardinality = part.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		v.processor_.PreprocessCardinality(cardinality)
		v.visitCardinality(cardinality)
		v.processor_.PostprocessCardinality(cardinality)
	}
}

func (v *visitor_) visitPattern(
	pattern ast.PatternLike,
) {
	// Visit the part.
	var part = pattern.GetPart()
	v.processor_.PreprocessPart(part, 1)
	v.visitPart(part)
	v.processor_.PostprocessPart(part, 1)

	// Visit the optional supplement.
	var supplement = pattern.GetOptionalSupplement()
	if col.IsDefined(supplement) {
		v.processor_.PreprocessSupplement(supplement)
		v.visitSupplement(supplement)
		v.processor_.PostprocessSupplement(supplement)
	}
}

func (v *visitor_) visitPredicate(
	predicate ast.PredicateLike,
) {
	// Visit the identifier.
	var identifier = predicate.GetIdentifier()
	v.processor_.PreprocessIdentifier(identifier)
	v.visitIdentifier(identifier)
	v.processor_.PostprocessIdentifier(identifier)

	// Visit the optional cardinality.
	var cardinality = predicate.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		v.processor_.PreprocessCardinality(cardinality)
		v.visitCardinality(cardinality)
		v.processor_.PostprocessCardinality(cardinality)
	}
}

func (v *visitor_) visitRule(
	rule ast.RuleLike,
) {
	// Visit the optional comment.
	var comment = rule.GetOptionalComment()
	if col.IsDefined(comment) {
		v.processor_.ProcessComment(comment)
	}

	// Visit the uppercase identifier.
	var uppercase = rule.GetUppercase()
	v.processor_.ProcessUppercase(uppercase)

	// Visit the separator.
	var separator = rule.GetSeparator()
	v.processor_.ProcessSeparator(separator)

	// Visit the definition.
	var definition = rule.GetDefinition()
	v.processor_.PreprocessDefinition(definition)
	v.visitDefinition(definition)
	v.processor_.PostprocessDefinition(definition)

	// Visit each newline.
	var index uint
	var newlines = rule.GetNewlines().GetIterator()
	for newlines.HasNext() {
		index++
		var newline = newlines.GetNext()
		v.processor_.ProcessNewline(newline, index)
	}
}

func (v *visitor_) visitSelective(
	selective ast.SelectiveLike,
) {
	// Visit each alternative.
	var index uint
	var alternatives = selective.GetAlternatives().GetIterator()
	for alternatives.HasNext() {
		index++
		var alternative = alternatives.GetNext()
		v.processor_.PreprocessAlternative(alternative, index)
		v.visitAlternative(alternative)
		v.processor_.PostprocessAlternative(alternative, index)
	}
}

func (v *visitor_) visitSequential(
	sequential ast.SequentialLike,
) {
	// Visit each part.
	var index uint = 1
	var parts = sequential.GetParts().GetIterator()
	for parts.HasNext() {
		index++
		var part = parts.GetNext()
		v.processor_.PreprocessPart(part, index)
		v.visitPart(part)
		v.processor_.PostprocessPart(part, index)
	}
}

func (v *visitor_) visitSupplement(
	pattern ast.SupplementLike,
) {
	// Visit the possible pattern types.
	switch actual := pattern.GetAny().(type) {
	case ast.SequentialLike:
		v.processor_.PreprocessSequential(actual)
		v.visitSequential(actual)
		v.processor_.PostprocessSequential(actual)
	case ast.SelectiveLike:
		v.processor_.PreprocessSelective(actual)
		v.visitSelective(actual)
		v.processor_.PostprocessSelective(actual)
	default:
		panic("An supplement must have a sequential or selective  value.")
	}
}

func (v *visitor_) visitSyntax(
	syntax ast.SyntaxLike,
) {
	// Visit each header.
	var index uint
	var headers = syntax.GetHeaders().GetIterator()
	for headers.HasNext() {
		index++
		var header = headers.GetNext()
		v.processor_.PreprocessHeader(header, index)
		v.visitHeader(header)
		v.processor_.PostprocessHeader(header, index)
	}

	// Visit each rule.
	index = 0
	var rules = syntax.GetRules().GetIterator()
	for rules.HasNext() {
		index++
		var rule = rules.GetNext()
		v.processor_.PreprocessRule(rule, index)
		v.visitRule(rule)
		v.processor_.PostprocessRule(rule, index)
	}

	// Visit each expression.
	index = 0
	var expressions = syntax.GetExpressions().GetIterator()
	for expressions.HasNext() {
		index++
		var expression = expressions.GetNext()
		v.processor_.PreprocessExpression(expression, index)
		v.visitExpression(expression)
		v.processor_.PostprocessExpression(expression, index)
	}
}

func (v *visitor_) visitTextual(
	textual ast.TextualLike,
) {
	// Visit the possible textual element types.
	var string_ = textual.GetAny().(string)
	switch {
	case Scanner().MatchesType(string_, IntrinsicToken):
		v.processor_.ProcessIntrinsic(string_)
	case Scanner().MatchesType(string_, GlyphToken):
		v.processor_.ProcessGlyph(string_)
	case Scanner().MatchesType(string_, LiteralToken):
		v.processor_.ProcessLiteral(string_)
	case Scanner().MatchesType(string_, LowercaseToken):
		v.processor_.ProcessLowercase(string_)
	default:
		var message = fmt.Sprintf(
			"An invalid textual element was found: %v",
			string_,
		)
		panic(message)
	}
}
