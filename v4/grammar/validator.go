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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	stc "strconv"
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
	var processor = Processor().Make()
	var validator = &validator_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: processor,
	}
	validator.visitor_ = Visitor().Make(validator)
	return validator
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// Define the instance attributes.
	class_       ValidatorClassLike
	visitor_     VisitorLike
	rules_       abs.CatalogLike[string, ast.DefinitionLike]
	expressions_ abs.CatalogLike[string, ast.PatternLike]

	// Define the inherited aspects.
	Methodical
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

func (v *validator_) ValidateToken(
	tokenValue string,
	tokenType TokenType,
) {
	if !Scanner().MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			Scanner().FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}

func (v *validator_) ValidateSyntax(syntax ast.SyntaxLike) {
	v.rules_ = col.Catalog[string, ast.DefinitionLike]()
	v.expressions_ = col.Catalog[string, ast.PatternLike]()
	v.visitor_.VisitSyntax(syntax)
}

// Methodical

func (v *validator_) ProcessComment(comment string) {
	v.ValidateToken(comment, CommentToken)
}

func (v *validator_) ProcessDelimiter(delimiter string) {
	v.ValidateToken(delimiter, DelimiterToken)
}

func (v *validator_) ProcessGlyph(glyph string) {
	v.ValidateToken(glyph, GlyphToken)
}

func (v *validator_) ProcessIntrinsic(intrinsic string) {
	v.ValidateToken(intrinsic, IntrinsicToken)
}

func (v *validator_) ProcessLiteral(literal string) {
	v.ValidateToken(literal, LiteralToken)
}

func (v *validator_) ProcessLowercase(lowercase string) {
	v.ValidateToken(lowercase, LowercaseToken)
}

func (v *validator_) ProcessNegation(negation string) {
	v.ValidateToken(negation, NegationToken)
}

func (v *validator_) ProcessNote(note string) {
	v.ValidateToken(note, NoteToken)
}

func (v *validator_) ProcessNumber(number string) {
	v.ValidateToken(number, NumberToken)
}

func (v *validator_) ProcessQuantified(quantified string) {
	v.ValidateToken(quantified, QuantifiedToken)
}

func (v *validator_) ProcessUppercase(uppercase string) {
	v.ValidateToken(uppercase, UppercaseToken)
}

func (v *validator_) PreprocessBounded(bounded ast.BoundedLike) {
	var glyph = bounded.GetGlyph()
	var extent = bounded.GetOptionalExtent()
	if col.IsDefined(extent) {
		if glyph > extent.GetGlyph() {
			var message = "The extent glyph in a bounded character range cannot come before the initial glyph."
			panic(message)
		}
	}
}

func (v *validator_) PreprocessConstrained(constrained ast.ConstrainedLike) {
	var number = constrained.GetNumber()
	var limit = constrained.GetOptionalLimit()
	if col.IsDefined(limit) {
		var optionalNumber = limit.GetOptionalNumber()
		if col.IsDefined(optionalNumber) {
			var minimum, _ = stc.Atoi(number)
			var maximum, _ = stc.Atoi(optionalNumber)
			if minimum > maximum {
				var message = "The limit in a constrained cardinality cannot be less than the minimum."
				panic(message)
			}
		}
	}
}

func (v *validator_) PreprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
	var lowercase = expression.GetLowercase()
	var duplicate = v.expressions_.GetValue(lowercase)
	if col.IsDefined(duplicate) {
		var message = fmt.Sprintf(
			"The expression %q is defined more than once.",
			lowercase,
		)
		panic(message)
	}
	var pattern = expression.GetPattern()
	v.expressions_.SetValue(lowercase, pattern)
}

func (v *validator_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var uppercase = rule.GetUppercase()
	var duplicate = v.rules_.GetValue(uppercase)
	if col.IsDefined(duplicate) {
		var message = fmt.Sprintf(
			"The rule %q is defined more than once.",
			uppercase,
		)
		panic(message)
	}
	var definition = rule.GetDefinition()
	v.rules_.SetValue(uppercase, definition)
}

func (v *validator_) PostprocessSyntax(syntax ast.SyntaxLike) {
	// Make sure each rule is defined.
	var rules = syntax.GetRules().GetIterator()
	var rulenames = v.rules_.GetKeys().GetIterator()
ruleLoop:
	for rulenames.HasNext() {
		var name = rulenames.GetNext()
		for rules.HasNext() {
			var rule = rules.GetNext()
			if name == rule.GetUppercase() {
				// Found a matching rule name.
				continue ruleLoop
			}
		}
		var message = fmt.Sprintf(
			"The rule %q is missing a definition.",
			name,
		)
		panic(message)
	}

	// Make sure each expression is defined.
	var expressions = syntax.GetExpressions().GetIterator()
	var expressionnames = v.expressions_.GetKeys().GetIterator()
expressionLoop:
	for expressionnames.HasNext() {
		var name = expressionnames.GetNext()
		for expressions.HasNext() {
			var expression = expressions.GetNext()
			if name == expression.GetLowercase() {
				// Found a matching expression name.
				continue expressionLoop
			}
		}
		var message = fmt.Sprintf(
			"The expression %q is missing a pattern.",
			name,
		)
		panic(message)
	}
}
