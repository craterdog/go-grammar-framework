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
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	stc "strconv"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// Initialize class constants.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// CLASS METHODS

// Target

type validatorClass_ struct {
	// Define class constants.
}

// Constructors

func (c *validatorClass_) Make() ValidatorLike {
	return &validator_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// Define instance attributes.
	class_   ValidatorClassLike
	isToken_ bool
	stack_   col.StackLike[ast.DefinitionLike]
	names_   col.CatalogLike[string, ast.ExpressionLike]
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

func (v *validator_) ValidateSyntax(syntax ast.SyntaxLike) {
	var notation = cdc.Notation().Make()
	v.stack_ = col.Stack[ast.DefinitionLike](notation).Make()
	v.names_ = col.Catalog[string, ast.ExpressionLike](notation).Make()
	v.validateSyntax(syntax)
	var iterator = v.names_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var symbol = association.GetKey()
		var expression = association.GetValue()
		if expression == nil {
			var message = fmt.Sprintf(
				"The syntax is missing a definition for the symbol: %v\n",
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

func (v *validator_) validateAlternative(alternative ast.AlternativeLike) {
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

func (v *validator_) validateAtom(atom ast.AtomLike) {
	var intrinsic = atom.GetIntrinsic()
	var glyph = atom.GetGlyph()
	switch {
	case len(intrinsic) > 0 && glyph == nil:
		v.validateIntrinsic(intrinsic)
	case len(intrinsic) == 0 && glyph != nil:
		v.validateGlyph(glyph)
	default:
		var message = v.formatError(
			"An atom must contain an intrinsic or a glyph but not both.",
		)
		panic(message)
	}
}

func (v *validator_) validateCardinality(cardinality ast.CardinalityLike) {
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

func (v *validator_) validateConstraint(constraint ast.ConstraintLike) {
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

func (v *validator_) validateDefinition(definition ast.DefinitionLike) {
	v.stack_.AddValue(definition)
	var comment = definition.GetComment()
	if len(comment) > 0 {
		v.validateComment(comment)
	}
	var name = definition.GetName()
	v.validateName(name)
	if uni.IsLower([]rune(name)[0]) {
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

func (v *validator_) validateElement(element ast.ElementLike) {
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

func (v *validator_) validateExpression(expression ast.ExpressionLike) {
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

func (v *validator_) validateFactor(factor ast.FactorLike) {
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

func (v *validator_) validateFilter(filter ast.FilterLike) {
	var atoms = filter.GetAtoms()
	if atoms == nil || atoms.IsEmpty() {
		var message = v.formatError(
			"A filter must contain at least one atom.",
		)
		panic(message)
	}
	var iterator = atoms.GetIterator()
	for iterator.HasNext() {
		var atom = iterator.GetNext()
		v.validateAtom(atom)
	}
}

func (v *validator_) validateGlyph(glyph ast.GlyphLike) {
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

func (v *validator_) validateSyntax(syntax ast.SyntaxLike) {
	var headers = syntax.GetHeaders()
	if headers == nil || headers.IsEmpty() {
		var message = "The syntax must contain at least one header.\n"
		panic(message)
	}
	var headerIterator = headers.GetIterator()
	for headerIterator.HasNext() {
		var header = headerIterator.GetNext()
		v.validateHeader(header)
	}
	var definitions = syntax.GetDefinitions()
	if definitions == nil || definitions.IsEmpty() {
		var message = "The syntax must contain at least one definition.\n"
		panic(message)
	}
	var definitionIterator = definitions.GetIterator()
	for definitionIterator.HasNext() {
		var definition = definitionIterator.GetNext()
		v.validateDefinition(definition)
	}
}

func (v *validator_) validateHeader(header ast.HeaderLike) {
	var comment = header.GetComment()
	if len(comment) == 0 {
		var message = v.formatError(
			"A header must contain a comment.",
		)
		panic(message)
	}
	v.validateComment(comment)
}

func (v *validator_) validateInline(inline ast.InlineLike) {
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

func (v *validator_) validateIntrinsic(intrinsic string) {
	var matches = Scanner().MatchToken(IntrinsicToken, intrinsic)
	if matches.IsEmpty() {
		var message = v.formatError(
			"Found an invalid intrinsic.",
		)
		panic(message)
	}
}

func (v *validator_) validateLine(line ast.LineLike) {
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

func (v *validator_) validateLiteral(literal string) {
	var matches = Scanner().MatchToken(LiteralToken, literal)
	if matches.IsEmpty() {
		var message = v.formatError(
			"Found an invalid literal.",
		)
		panic(message)
	}
}

func (v *validator_) validateMultiline(multiline ast.MultilineLike) {
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

func (v *validator_) validateName(name string) {
	var matches = Scanner().MatchToken(NameToken, name)
	if matches.IsEmpty() {
		var message = v.formatError(
			"Found an invalid name.",
		)
		panic(message)
	}
	if uni.IsUpper([]rune(name)[0]) {
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

func (v *validator_) validatePrecedence(precedence ast.PrecedenceLike) {
	var expression = precedence.GetExpression()
	if precedence == nil {
		var message = v.formatError(
			"A precedence must contain an expression.",
		)
		panic(message)
	}
	v.validateExpression(expression)
}

func (v *validator_) validatePredicate(predicate ast.PredicateLike) {
	var atom = predicate.GetAtom()
	var element = predicate.GetElement()
	var filter = predicate.GetFilter()
	var precedence = predicate.GetPrecedence()
	switch {
	case atom != nil && element == nil && filter == nil && precedence == nil:
		v.validateAtom(atom)
	case atom == nil && element != nil && filter == nil && precedence == nil:
		v.validateElement(element)
	case atom == nil && element == nil && filter != nil && precedence == nil:
		v.validateFilter(filter)
	case atom == nil && element == nil && filter == nil && precedence != nil:
		v.validatePrecedence(precedence)
	default:
		var message = v.formatError(
			"A predicate must contain exactly one atom, element, filter, or precedence.",
		)
		panic(message)
	}
}
