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
	fwk "github.com/craterdog/go-collection-framework/v4"
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
	var name = syntax.GetRules().GetIterator().GetNext().GetUppercase()
	var notation = fwk.CDCN()
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
		if fwk.IsUndefined(expression) {
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
		if fwk.IsUndefined(expression) {
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
	if parts.IsEmpty() {
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
	v.validateInitial(name, initial)

	// Validate the optional extent rune.
	var extent = bounded.GetOptionalExtent()
	if fwk.IsDefined(extent) {
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
	v.validateMinimum(name, minimum)

	// Validate the optional maximum value.
	var maximum = constrained.GetOptionalMaximum()
	if fwk.IsDefined(maximum) {
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
	v.validateToken(name, RuneToken, rune_)
}

func (v *validator_) validateFactor(
	name string,
	factor ast.FactorLike,
) {
	// Validate the predicate.
	var predicate = factor.GetPredicate()
	v.validatePredicate(name, predicate)

	// Validate the optional cardinality.
	var cardinality = factor.GetOptionalCardinality()
	if fwk.IsDefined(cardinality) {
		v.validateCardinality(name, cardinality)
	}
}

func (v *validator_) validateFiltered(
	name string,
	filtered ast.FilteredLike,
) {
	// Validate the optional negation.
	var negation = filtered.GetOptionalNegation()
	if fwk.IsDefined(negation) {
		v.validateToken(name, NegationToken, negation)
	}

	// Validate the characters.
	var characters = filtered.GetCharacters()
	if characters.IsEmpty() {
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
	v.validatePattern(name, pattern)
}

func (v *validator_) validateHeader(
	name string,
	header ast.HeaderLike,
) {
	// Validate the comment.
	var comment = header.GetComment()
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
	v.validateToken(name, RuneToken, rune_)
}

func (v *validator_) validateInlined(
	name string,
	inlined ast.InlinedLike,
) {
	// Validate the factors.
	var factors = inlined.GetFactors()
	if factors.IsEmpty() {
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
	var note = inlined.GetOptionalNote()
	if fwk.IsDefined(note) {
		v.validateToken(name, NoteToken, note)
	}
}

func (v *validator_) validateLexigram(
	name string,
	lexigram ast.LexigramLike,
	lexigrams col.CatalogLike[string, ast.PatternLike],
) {
	// Validate the optional comment.
	var comment = lexigram.GetOptionalComment()
	if fwk.IsDefined(comment) {
		v.validateToken(name, CommentToken, comment)
	}

	// Validate the lowercase identifier.
	var lowercase = lexigram.GetLowercase()
	v.validateToken(name, LowercaseToken, lowercase)

	// Validate the pattern.
	var pattern = lexigram.GetPattern()
	v.validatePattern(name, pattern)

	// Check for duplicate lexigram definitions.
	var duplicate = lexigrams.GetValue(lowercase)
	if fwk.IsDefined(duplicate) {
		var message = v.formatError(
			name,
			"The lexigram is defined more than once.",
		)
		panic(message)
	}
	lexigrams.SetValue(lowercase, pattern)

	// Validate the optional note.
	var note = lexigram.GetOptionalNote()
	if fwk.IsDefined(note) {
		v.validateToken(name, NoteToken, note)
	}
}

func (v *validator_) validateLine(
	name string,
	line ast.LineLike,
) {
	// Validate the identifier.
	var identifier = line.GetIdentifier()
	v.validateIdentifier(name, identifier)

	// Validate the optional note.
	var note = line.GetOptionalNote()
	if fwk.IsDefined(note) {
		v.validateToken(name, NoteToken, note)
	}
}

func (v *validator_) validateMaximum(
	name string,
	maximum ast.MaximumLike,
) {
	// Validate the optional number.
	var number = maximum.GetOptionalNumber()
	if fwk.IsDefined(number) {
		v.validateToken(name, NumberToken, number)
	}
}

func (v *validator_) validateMinimum(
	name string,
	minimum ast.MinimumLike,
) {
	// Validate the number.
	var number = minimum.GetNumber()
	v.validateToken(name, NumberToken, number)
}

func (v *validator_) validateMultilined(
	name string,
	multilined ast.MultilinedLike,
) {
	// Validate the lines.
	var lines = multilined.GetLines()
	if lines.IsEmpty() {
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
	v.validateElement(name, element)

	// Validate the optional cardinality.
	var cardinality = part.GetOptionalCardinality()
	if fwk.IsDefined(cardinality) {
		v.validateCardinality(name, cardinality)
	}
}

func (v *validator_) validatePattern(
	name string,
	pattern ast.PatternLike,
) {
	// Validate the parts.
	var parts = pattern.GetParts()
	if parts.IsEmpty() {
		var message = v.formatError(
			name,
			"Each pattern must have at least one part.",
		)
		panic(message)
	}
	var partIterator = parts.GetIterator()
	for partIterator.HasNext() {
		var part = partIterator.GetNext()
		v.validatePart(name, part)
	}

	// Validate the alternatives.
	var alternatives = pattern.GetAlternatives()
	var alternativeIterator = alternatives.GetIterator()
	for alternativeIterator.HasNext() {
		var alternative = alternativeIterator.GetNext()
		v.validateAlternative(name, alternative)
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
	var comment = rule.GetOptionalComment()
	if fwk.IsDefined(comment) {
		v.validateToken(name, CommentToken, comment)
	}

	// Validate the uppercase identifier.
	var uppercase = rule.GetUppercase()
	v.validateToken(name, UppercaseToken, uppercase)

	// Validate the expression.
	var expression = rule.GetExpression()
	v.validateExpression(name, expression)

	// Check for duplicate rule definitions.
	var duplicate = rules.GetValue(uppercase)
	if fwk.IsDefined(duplicate) {
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
	if syntaxHeaders.IsEmpty() {
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
	if syntaxRules.IsEmpty() {
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
	if syntaxLexigrams.IsEmpty() {
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
