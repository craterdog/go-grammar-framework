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
)

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// Initialize the class constants.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// CLASS METHODS

// Target

type validatorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *validatorClass_) Make() ValidatorLike {
	return &validator_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// Define the instance attributes.
	class_ ValidatorClassLike
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

func (v *validator_) ValidateSyntax(syntax ast.SyntaxLike) {
	// Initialize the state.
	var name = syntax.GetRules().GetValue(1).GetUppercase()
	var notation = cdc.Notation().Make()
	var rules = col.Catalog[string, ast.ExpressionLike](notation).Make()
	var lexigrams = col.Catalog[string, ast.PatternLike](notation).Make()

	// Validate the syntax.
	v.validateSyntax(name, syntax, rules, lexigrams)

	// Check for missing rule definitions.
	var ruleIterator = rules.GetIterator()
	for ruleIterator.HasNext() {
		var association = ruleIterator.GetNext()
		var rule = association.GetKey()
		var expression = association.GetValue()
		if expression == nil {
			var message = fmt.Sprintf(
				"The syntax is missing a definition for the rule: %v\n",
				rule,
			)
			panic(message)
		}
	}

	// Check for missing lexigram definitions.
	var lexigramIterator = lexigrams.GetIterator()
	for lexigramIterator.HasNext() {
		var association = lexigramIterator.GetNext()
		var lexigram = association.GetKey()
		var expression = association.GetValue()
		if expression == nil {
			var message = fmt.Sprintf(
				"The syntax is missing a definition for the lexigram: %v\n",
				lexigram,
			)
			panic(message)
		}
	}
}

// Private

func (v *validator_) formatError(name, message string) string {
	message = fmt.Sprintf(
		"The definition for %v is invalid:\n%v\n",
		name,
		message,
	)
	return message
}

func (v *validator_) matchesToken(type_ TokenType, value string) bool {
	var matches = Scanner().MatchToken(type_, value)
	return !matches.IsEmpty()
}

func (v *validator_) validateAlternative(
	name string,
	alternative ast.AlternativeLike,
) {
	// Validate the parts.
	var parts = alternative.GetParts()
	if parts == nil || parts.IsEmpty() {
		var message = v.formatError(
			name,
			"Each alternative must have at least one part.",
		)
		panic(message)
	}
	var iterator = parts.GetIterator()
	for iterator.HasNext() {
		var part = iterator.GetNext()
		v.validatePart(name, part)
	}
}

func (v *validator_) validateBounded(
	name string,
	bounded ast.BoundedLike,
) {
	// Validate the initial rune.
	var initial = bounded.GetInitial()
	if initial == nil {
		var message = v.formatError(
			name,
			"A bounded must have an initial rune.",
		)
		panic(message)
	}
	v.validateInitial(name, initial)

	// Validate the optional extent rune.
	var extent = bounded.GetExtent()
	if extent != nil {
		v.validateExtent(name, extent)
		if initial.GetRune() > extent.GetRune() {
			var message = v.formatError(
				name,
				"The extent rune in a bounded cannot come before the initial.",
			)
			panic(message)
		}
	}
}

func (v *validator_) validateCardinality(
	name string,
	cardinality ast.CardinalityLike,
) {
	// Validate the possible cardinality types.
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		v.validateConstrained(name, actual)
	case string:
		switch {
		case v.matchesToken(QuantifiedToken, actual):
		default:
			panic("A cardinality must have a value.")
		}
	default:
		panic("A cardinality must have a value.")
	}
}

func (v *validator_) validateCharacter(
	name string,
	character ast.CharacterLike,
) {
	// Validate the possible character types.
	switch actual := character.GetAny().(type) {
	case ast.BoundedLike:
		v.validateBounded(name, actual)
	case string:
		switch {
		case v.matchesToken(IntrinsicToken, actual):
		default:
			panic("An character must have a value.")
		}
	default:
		panic("An character must have a value.")
	}
}

func (v *validator_) validateConstrained(
	name string,
	constrained ast.ConstrainedLike,
) {
	// Validate the minimum value.
	var minimum = constrained.GetMinimum()
	if minimum == nil {
		var message = v.formatError(
			name,
			"A constrained must have a minimum value.",
		)
		panic(message)
	}
	v.validateMinimum(name, minimum)

	// Validate the optional maximum value.
	var maximum = constrained.GetMaximum()
	if maximum != nil {
		v.validateMaximum(name, maximum)
	}
}

func (v *validator_) validateElement(
	name string,
	element ast.ElementLike,
) {
	// Validate the possible element types.
	switch actual := element.GetAny().(type) {
	case ast.GroupedLike:
		v.validateGrouped(name, actual)
	case ast.FilteredLike:
		v.validateFiltered(name, actual)
	case ast.BoundedLike:
		v.validateBounded(name, actual)
	case string:
		switch {
		case v.matchesToken(IntrinsicToken, actual):
		case v.matchesToken(LowercaseToken, actual):
		case v.matchesToken(LiteralToken, actual):
		default:
			panic("An element must have a value.")
		}
	default:
		panic("An element must have a value.")
	}
}

func (v *validator_) validateExpression(
	name string,
	expression ast.ExpressionLike,
) {
	// Validate the possible expression types.
	switch actual := expression.GetAny().(type) {
	case ast.InlinedLike:
		v.validateInlined(name, actual)
	case ast.MultilinedLike:
		v.validateMultilined(name, actual)
	default:
		panic("An expression must have a value.")
	}
}

func (v *validator_) validateExtent(
	name string,
	extent ast.ExtentLike,
) {
	// Validate the rune.
	var rune_ = extent.GetRune()
	if len(rune_) > 0 {
		v.validateToken(name, RuneToken, rune_)
	}
}

func (v *validator_) validateFactor(
	name string,
	factor ast.FactorLike,
) {
	// Validate the predicate.
	var predicate = factor.GetPredicate()
	if predicate == nil {
		var message = v.formatError(
			name,
			"A factor must contain a predicate.",
		)
		panic(message)
	}
	v.validatePredicate(name, predicate)

	// Validate the optional cardinality.
	var cardinality = factor.GetCardinality()
	if cardinality != nil {
		v.validateCardinality(name, cardinality)
	}
}

func (v *validator_) validateFiltered(
	name string,
	filtered ast.FilteredLike,
) {
	// Validate the optional negation.
	var negation = filtered.GetNegation()
	if len(negation) > 0 {
		v.validateToken(name, NegationToken, negation)
	}

	// Validate the characters.
	var characters = filtered.GetCharacters()
	if characters == nil || characters.IsEmpty() {
		var message = v.formatError(
			name,
			"A filtered element must contain at least one character.",
		)
		panic(message)
	}
	var iterator = characters.GetIterator()
	for iterator.HasNext() {
		var character = iterator.GetNext()
		v.validateCharacter(name, character)
	}
}

func (v *validator_) validateGrouped(
	name string,
	grouped ast.GroupedLike,
) {
	// Validate the pattern.
	var pattern = grouped.GetPattern()
	if grouped == nil {
		var message = v.formatError(
			name,
			"A grouped element must contain a pattern.",
		)
		panic(message)
	}
	v.validatePattern(name, pattern)
}

func (v *validator_) validateHeader(
	name string,
	header ast.HeaderLike,
) {
	// Validate the comment.
	var comment = header.GetComment()
	if len(comment) == 0 {
		var message = v.formatError(
			name,
			"A header must contain a comment.",
		)
		panic(message)
	}
	v.validateToken(name, CommentToken, comment)
}

func (v *validator_) validateIdentifier(
	name string,
	identifier ast.IdentifierLike,
) {
	// Validate the possible identifier types.
	switch actual := identifier.GetAny().(type) {
	case string:
		switch {
		case v.matchesToken(LowercaseToken, actual):
		case v.matchesToken(UppercaseToken, actual):
		default:
			panic("An identifier must have a value.")
		}
	default:
		panic("An identifier must have a value.")
	}
}

func (v *validator_) validateInitial(
	name string,
	initial ast.InitialLike,
) {
	// Validate the rune.
	var rune_ = initial.GetRune()
	if len(rune_) == 0 {
		var message = v.formatError(
			name,
			"A initial must have a rune.",
		)
		panic(message)
	}
	v.validateToken(name, RuneToken, rune_)
}

func (v *validator_) validateInlined(
	name string,
	inlined ast.InlinedLike,
) {
	// Validate the factors.
	var factors = inlined.GetFactors()
	if factors == nil || factors.IsEmpty() {
		var message = v.formatError(
			name,
			"Each inlined expression must have at least one factor.",
		)
		panic(message)
	}
	var iterator = factors.GetIterator()
	for iterator.HasNext() {
		var factor = iterator.GetNext()
		v.validateFactor(name, factor)
	}

	// Validate the optional note.
	var note = inlined.GetNote()
	if len(note) > 0 {
		v.validateToken(name, NoteToken, note)
	}
}

func (v *validator_) validateLexigram(
	name string,
	lexigram ast.LexigramLike,
	lexigrams col.CatalogLike[string, ast.PatternLike],
) {
	// Validate the optional comment.
	var comment = lexigram.GetComment()
	if len(comment) > 0 {
		v.validateToken(name, CommentToken, comment)
	}

	// Validate the lowercase identifier.
	var lowercase = lexigram.GetLowercase()
	if len(lowercase) == 0 {
		var message = v.formatError(
			name,
			"A lexigram must contain a lowercase identifier.",
		)
		panic(message)
	}
	v.validateToken(name, LowercaseToken, lowercase)

	// Validate the pattern.
	var pattern = lexigram.GetPattern()
	if pattern == nil {
		var message = v.formatError(
			name,
			"A lexigram must contain a pattern.",
		)
		panic(message)
	}
	v.validatePattern(name, pattern)

	// Check for duplicate lexigram definitions.
	if lexigrams.GetValue(lowercase) != nil {
		var message = v.formatError(
			name,
			"The lexigram is defined more than once.",
		)
		panic(message)
	}
	lexigrams.SetValue(lowercase, pattern)

	// Validate the optional note.
	var note = lexigram.GetNote()
	if len(note) > 0 {
		v.validateToken(name, NoteToken, note)
	}
}

func (v *validator_) validateLine(
	name string,
	line ast.LineLike,
) {
	// Validate the identifier.
	var identifier = line.GetIdentifier()
	if identifier == nil {
		var message = v.formatError(
			name,
			"An line must have an identifier.",
		)
		panic(message)
	}
	v.validateIdentifier(name, identifier)

	// Validate the optional note.
	var note = line.GetNote()
	if len(note) > 0 {
		v.validateToken(name, NoteToken, note)
	}
}

func (v *validator_) validateMaximum(
	name string,
	maximum ast.MaximumLike,
) {
	// Validate the number.
	var number = maximum.GetNumber()
	if len(number) > 0 {
		v.validateToken(name, NumberToken, number)
	}
}

func (v *validator_) validateMinimum(
	name string,
	minimum ast.MinimumLike,
) {
	// Validate the number.
	var number = minimum.GetNumber()
	if len(number) == 0 {
		var message = v.formatError(
			name,
			"A minimum must have a number.",
		)
		panic(message)
	}
	v.validateToken(name, NumberToken, number)
}

func (v *validator_) validateMultilined(
	name string,
	multilined ast.MultilinedLike,
) {
	// Validate the lines.
	var lines = multilined.GetLines()
	if lines == nil || lines.IsEmpty() {
		var message = v.formatError(
			name,
			"Each multi-line expression must have at least one line.",
		)
		panic(message)
	}
	var iterator = lines.GetIterator()
	for iterator.HasNext() {
		var line = iterator.GetNext()
		v.validateLine(name, line)
	}
}

func (v *validator_) validatePart(
	name string,
	part ast.PartLike,
) {
	// Validate the element.
	var element = part.GetElement()
	if element == nil {
		var message = v.formatError(
			name,
			"A part must contain an element.",
		)
		panic(message)
	}
	v.validateElement(name, element)

	// Validate the optional cardinality.
	var cardinality = part.GetCardinality()
	if cardinality != nil {
		v.validateCardinality(name, cardinality)
	}
}

func (v *validator_) validatePattern(
	name string,
	pattern ast.PatternLike,
) {
	// Validate the parts.
	var parts = pattern.GetParts()
	if parts == nil || parts.IsEmpty() {
		var message = v.formatError(
			name,
			"Each pattern must have at least one part.",
		)
		panic(message)
	}
	var iterator = parts.GetIterator()
	for iterator.HasNext() {
		var part = iterator.GetNext()
		v.validatePart(name, part)
	}

	// Validate the alternatives.
	var alternatives = pattern.GetAlternatives()
	if alternatives != nil {
		var iterator = alternatives.GetIterator()
		for iterator.HasNext() {
			var alternative = iterator.GetNext()
			v.validateAlternative(name, alternative)
		}
	}
}

func (v *validator_) validatePredicate(
	name string,
	predicate ast.PredicateLike,
) {
	// Validate the possible predicate types.
	switch actual := predicate.GetAny().(type) {
	case string:
		switch {
		case v.matchesToken(LowercaseToken, actual):
		case v.matchesToken(UppercaseToken, actual):
		case v.matchesToken(IntrinsicToken, actual):
		case v.matchesToken(LiteralToken, actual):
		default:
			panic("A predicate must have a value.")
		}
	default:
		panic("A predicate must have a value.")
	}
}

func (v *validator_) validateRule(
	name string,
	rule ast.RuleLike,
	rules col.CatalogLike[string, ast.ExpressionLike],
) {
	// Validate the optional comment.
	var comment = rule.GetComment()
	if len(comment) > 0 {
		v.validateToken(name, CommentToken, comment)
	}

	// Validate the uppercase identifier.
	var uppercase = rule.GetUppercase()
	if len(uppercase) == 0 {
		var message = v.formatError(
			name,
			"A rule must contain an uppercase identifier.",
		)
		panic(message)
	}
	v.validateToken(name, UppercaseToken, uppercase)

	// Validate the expression.
	var expression = rule.GetExpression()
	if expression == nil {
		var message = v.formatError(
			name,
			"A rule must contain an expression.",
		)
		panic(message)
	}
	v.validateExpression(name, expression)

	// Check for duplicate rule definitions.
	if rules.GetValue(uppercase) != nil {
		var message = v.formatError(
			name,
			fmt.Sprintf("The rule %q is defined more than once.", uppercase),
		)
		panic(message)
	}
	rules.SetValue(uppercase, expression)
}

func (v *validator_) validateSyntax(
	name string,
	syntax ast.SyntaxLike,
	rules col.CatalogLike[string, ast.ExpressionLike],
	lexigrams col.CatalogLike[string, ast.PatternLike],
) {
	// Validate the headers.
	var syntaxHeaders = syntax.GetHeaders()
	if syntaxHeaders == nil || syntaxHeaders.IsEmpty() {
		var message = "The syntax must contain at least one header.\n"
		panic(message)
	}
	var headerIterator = syntaxHeaders.GetIterator()
	for headerIterator.HasNext() {
		var header = headerIterator.GetNext()
		v.validateHeader(name, header)
	}

	// Validate the rule definitions.
	var syntaxRules = syntax.GetRules()
	if syntaxRules == nil || syntaxRules.IsEmpty() {
		var message = "The syntax must contain at least one rule definition.\n"
		panic(message)
	}
	var ruleIterator = syntaxRules.GetIterator()
	for ruleIterator.HasNext() {
		var rule = ruleIterator.GetNext()
		v.validateRule(name, rule, rules)
	}

	// Validate the lexigram definition.
	var syntaxLexigrams = syntax.GetLexigrams()
	if syntaxLexigrams == nil || syntaxLexigrams.IsEmpty() {
		var message = "The syntax must contain at least one lexigram definition.\n"
		panic(message)
	}
	var lexigramIterator = syntaxLexigrams.GetIterator()
	for lexigramIterator.HasNext() {
		var lexigram = lexigramIterator.GetNext()
		v.validateLexigram(name, lexigram, lexigrams)
	}
}

func (v *validator_) validateToken(
	name string,
	type_ TokenType,
	value string,
) {
	if !v.matchesToken(type_, value) {
		var message = fmt.Sprintf(
			"The following value is not of type %v: %v",
			Scanner().AsString(type_),
			value,
		)
		panic(message)
	}
}
