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
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v3"
	stc "strconv"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// This class does not initialize any class constants.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// CLASS METHODS

// Target

type validatorClass_ struct {
	// This class does not define any class constants.
}

// Constructors

func (c *validatorClass_) Make() ValidatorLike {
	return &validator_{}
}

// INSTANCE METHODS

// Target

type validator_ struct {
	inInversion_ bool
	isToken_     bool
	stack_       col.StackLike[DefinitionLike]
	names_       col.CatalogLike[string, ExpressionLike]
}

// Public

func (v *validator_) ValidateGrammar(grammar GrammarLike) {
	v.stack_ = col.Stack[DefinitionLike]().Make()
	v.names_ = col.Catalog[string, ExpressionLike]().Make()
	v.validateGrammar(grammar)
	var iterator = v.names_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var symbol = association.GetKey()
		var expression = association.GetValue()
		if expression == nil {
			var message = fmt.Sprintf(
				"The grammar is missing a definition for the symbol: %v\n",
				symbol,
			)
			panic(message)
		}
	}
}

// Private

func (v *validator_) formatError(message string) string {
	var definition = v.stack_.RemoveTop()
	message = fmt.Sprintf(
		"The definition for %v is invalid:\n%v\n",
		definition.GetName(),
		message,
	)
	return message
}

func (v *validator_) validateAlternative(alternative AlternativeLike) {
	var factors = alternative.GetFactors()
	if factors == nil || factors.IsEmpty() {
		var message = v.formatError(
			"Each alternative must have at least one factor.",
		)
		panic(message)
	}
	var iterator = factors.GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		v.validateFactor(factor)
	}
}

func (v *validator_) validateCardinality(cardinality CardinalityLike) {
	var constraint = cardinality.GetConstraint()
	if constraint == nil {
		var message = v.formatError(
			"A cardinality must have a constraint.",
		)
		panic(message)
	}
	v.validateConstraint(constraint)
}

func (v *validator_) validateCharacter(character string) {
	var matches = Scanner().MatchToken(CharacterToken, character)
	if matches.IsEmpty() {
		var message = v.formatError(
			"Found an invalid character.",
		)
		panic(message)
	}
}

func (v *validator_) validateComment(comment string) {
	var matches = Scanner().MatchToken(CommentToken, comment)
	if matches.IsEmpty() {
		var message = v.formatError(
			"Found an invalid comment.",
		)
		panic(message)
	}
}

func (v *validator_) validateConstraint(constraint ConstraintLike) {
	var first = constraint.GetFirst()
	v.validateNumber(first)
	var last = constraint.GetLast()
	if len(last) > 0 {
		v.validateNumber(last)
		var firstNumber, _ = stc.ParseInt(first, 10, 64)
		var lastNumber, _ = stc.ParseInt(last, 10, 64)
		if firstNumber > lastNumber {
			var message = v.formatError(
				"The first number in a constraint cannot be greater than the last.",
			)
			panic(message)
		}
	}
}

func (v *validator_) validateHeader(header HeaderLike) {
	var comment = header.GetComment()
	if len(comment) == 0 {
		var message = v.formatError(
			"A header must contain a comment.",
		)
		panic(message)
	}
	v.validateComment(comment)
}

func (v *validator_) validateDefinition(definition DefinitionLike) {
	v.stack_.AddValue(definition)
	var comment = definition.GetComment()
	if len(comment) > 0 {
		v.validateComment(comment)
	}
	var name = definition.GetName()
	v.validateName(name)
	if uni.IsUpper([]rune(name)[0]) {
		v.isToken_ = true
	}
	var expression = definition.GetExpression()
	if expression == nil {
		var message = v.formatError(
			"A definition must contain an expression.",
		)
		panic(message)
	}
	if v.names_.GetValue(name) != nil {
		var message = v.formatError(
			fmt.Sprintf(
				"The name %s is defined more than once.",
				name,
			),
		)
		panic(message)
	}
	v.names_.SetValue(name, expression)
	v.validateExpression(expression)
	v.isToken_ = false
	var _ = v.stack_.RemoveTop()
}

func (v *validator_) validateElement(element ElementLike) {
	var literal = element.GetLiteral()
	var name = element.GetName()
	switch {
	case len(literal) > 0 && len(name) == 0:
		v.validateLiteral(literal)
	case len(literal) == 0 && len(name) > 0:
		v.validateName(name)
	default:
		var message = v.formatError(
			"An element must contain a literal or a name but not both.",
		)
		panic(message)
	}
}

func (v *validator_) validateInline(inline InlineLike) {
	var alternatives = inline.GetAlternatives()
	if alternatives == nil || alternatives.IsEmpty() {
		var message = v.formatError(
			"Each inline expression must have at least one alternative.",
		)
		panic(message)
	}
	var iterator = alternatives.GetIterator()
	for iterator.HasNext() {
		var alternative = iterator.GetNext()
		v.validateAlternative(alternative)
	}
	var note = inline.GetNote()
	if len(note) > 0 {
		v.validateNote(note)
	}
}

func (v *validator_) validateInversion(inversion InversionLike) {
	var filter = inversion.GetFilter()
	if filter == nil {
		var message = v.formatError(
			"An inversion must contain a filter.",
		)
		panic(message)
	}
	v.validateFilter(filter)
}

func (v *validator_) validateLine(line LineLike) {
	var alternative = line.GetAlternative()
	if alternative == nil {
		var message = v.formatError(
			"A line must contain an alternative.",
		)
		panic(message)
	}
	v.validateAlternative(alternative)
	var note = line.GetNote()
	if len(note) > 0 {
		v.validateNote(note)
	}
}

func (v *validator_) validateMultiline(multiline MultilineLike) {
	var lines = multiline.GetLines()
	if lines == nil || lines.IsEmpty() {
		var message = v.formatError(
			"Each multi-line expression must have at least one line.",
		)
		panic(message)
	}
	var iterator = lines.GetIterator()
	for iterator.HasNext() {
		var line = iterator.GetNext()
		v.validateLine(line)
	}
}

func (v *validator_) validateExpression(expression ExpressionLike) {
	var inline = expression.GetInline()
	var multiline = expression.GetMultiline()
	switch {
	case inline != nil && multiline == nil:
		v.validateInline(inline)
	case inline == nil && multiline != nil:
		v.validateMultiline(multiline)
	default:
		var message = v.formatError(
			"An expression must be inline or multiline but not both.",
		)
		panic(message)
	}
}

func (v *validator_) validateFactor(factor FactorLike) {
	var predicate = factor.GetPredicate()
	if predicate == nil {
		var message = v.formatError(
			"A factor must contain a predicate.",
		)
		panic(message)
	}
	v.validatePredicate(predicate)
	var cardinality = factor.GetCardinality()
	if cardinality != nil {
		v.validateCardinality(cardinality)
	}
}

func (v *validator_) validateFilter(filter FilterLike) {
	var intrinsic = filter.GetIntrinsic()
	var glyph = filter.GetGlyph()
	switch {
	case len(intrinsic) > 0 && glyph == nil:
		v.validateIntrinsic(intrinsic)
	case len(intrinsic) == 0 && glyph != nil:
		v.validateGlyph(glyph)
	default:
		var message = v.formatError(
			"A filter must contain an intrinsic or a glyph but not both.",
		)
		panic(message)
	}
}

func (v *validator_) validateGlyph(glyph GlyphLike) {
	var first = glyph.GetFirst()
	v.validateCharacter(first)
	var last = glyph.GetLast()
	if len(last) > 0 {
		v.validateCharacter(last)
		if first > last {
			var message = v.formatError(
				"The first character in a glyph cannot come later than the last.",
			)
			panic(message)
		}
	}
}

func (v *validator_) validateGrammar(grammar GrammarLike) {
	var headers = grammar.GetHeaders()
	if headers == nil || headers.IsEmpty() {
		var message = "The grammar must contain at least one header.\n"
		panic(message)
	}
	var headerIterator = headers.GetIterator()
	for headerIterator.HasNext() {
		var header = headerIterator.GetNext()
		v.validateHeader(header)
	}
	var definitions = grammar.GetDefinitions()
	if definitions == nil || definitions.IsEmpty() {
		var message = "The grammar must contain at least one definition.\n"
		panic(message)
	}
	var definitionIterator = definitions.GetIterator()
	for definitionIterator.HasNext() {
		var definition = definitionIterator.GetNext()
		v.validateDefinition(definition)
	}
}

func (v *validator_) validateIntrinsic(intrinsic string) {
	var matches = Scanner().MatchToken(IntrinsicToken, intrinsic)
	if matches.IsEmpty() {
		var message = v.formatError(
			"Found an invalid intrinsic.",
		)
		panic(message)
	}
}

func (v *validator_) validateLiteral(literal string) {
	var matches = Scanner().MatchToken(LiteralToken, literal)
	if matches.IsEmpty() {
		var message = v.formatError(
			"Found an invalid literal.",
		)
		panic(message)
	}
	if v.inInversion_ && len([]rune(literal)) > 3 {
		var message = v.formatError(
			"A multi-character literal is not allowed in an inversion.",
		)
		panic(message)
	}
}

func (v *validator_) validateName(name string) {
	var matches = Scanner().MatchToken(NameToken, name)
	if matches.IsEmpty() {
		var message = v.formatError(
			"Found an invalid name.",
		)
		panic(message)
	}
	if uni.IsLower([]rune(name)[0]) {
		if v.isToken_ {
			var message = v.formatError(
				"A token definition cannot contain a rule name.",
			)
			panic(message)
		}
	}
	var expression = v.names_.GetValue(name)
	v.names_.SetValue(name, expression)
}

func (v *validator_) validateNote(note string) {
	var matches = Scanner().MatchToken(NoteToken, note)
	if matches.IsEmpty() {
		var message = v.formatError(
			"Found an invalid note.",
		)
		panic(message)
	}
}

func (v *validator_) validateNumber(number string) {
	var matches = Scanner().MatchToken(NumberToken, number)
	if matches.IsEmpty() {
		var message = v.formatError(
			"Found an invalid number.",
		)
		panic(message)
	}
}

func (v *validator_) validatePrecedence(precedence PrecedenceLike) {
	var expression = precedence.GetExpression()
	if precedence == nil {
		var message = v.formatError(
			"A precedence must contain an expression.",
		)
		panic(message)
	}
	v.validateExpression(expression)
}

func (v *validator_) validatePredicate(predicate PredicateLike) {
	var element = predicate.GetElement()
	var inversion = predicate.GetInversion()
	var precedence = predicate.GetPrecedence()
	switch {
	case element != nil && inversion == nil && precedence == nil:
		v.validateElement(element)
	case element == nil && inversion != nil && precedence == nil:
		v.validateInversion(inversion)
	case element == nil && inversion == nil && precedence != nil:
		v.validatePrecedence(precedence)
	default:
		var message = v.formatError(
			"A predicate must contain exactly one element, inversion, or precedence.",
		)
		panic(message)
	}
}
