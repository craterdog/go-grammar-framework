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

func (v *visitor_) visitAlternative(alternative ast.AlternativeLike) {
	// Visit the "|" delimiter literal.
	var delimiter = alternative.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)

	// Visit the part.
	var part = alternative.GetPart()
	v.processor_.PreprocessPart(part, 2, 1)
	v.visitPart(part)
	v.processor_.PostprocessPart(part, 2, 1)
}

func (v *visitor_) visitBounded(bounded ast.BoundedLike) {
	// Visit the runic token.
	var runic = bounded.GetRunic()
	v.processor_.ProcessRunic(runic)

	// Visit the optional extent.
	var extent = bounded.GetOptionalExtent()
	if col.IsDefined(extent) {
		v.processor_.PreprocessExtent(extent)
		v.visitExtent(extent)
		v.processor_.PostprocessExtent(extent)
	}
}

func (v *visitor_) visitCardinality(cardinality ast.CardinalityLike) {
	// Visit the possible cardinality types.
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		v.processor_.PreprocessConstrained(actual)
		v.visitConstrained(actual)
		v.processor_.PostprocessConstrained(actual)
	case ast.QuantifiedLike:
		v.processor_.PreprocessQuantified(actual)
		v.visitQuantified(actual)
		v.processor_.PostprocessQuantified(actual)
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitCharacter(character ast.CharacterLike) {
	// Visit the possible character types.
	switch actual := character.GetAny().(type) {
	case ast.BoundedLike:
		v.processor_.PreprocessBounded(actual)
		v.visitBounded(actual)
		v.processor_.PostprocessBounded(actual)
	case string:
		switch {
		case Scanner().MatchesType(actual, IntrinsicToken):
			v.processor_.ProcessIntrinsic(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}

	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitConstrained(constrained ast.ConstrainedLike) {
	// Visit the possible constrained types.
	switch actual := constrained.GetAny().(type) {
	case string:
		switch {
		case Scanner().MatchesType(actual, OptionalToken):
			v.processor_.ProcessOptional(actual)
		case Scanner().MatchesType(actual, RepeatedToken):
			v.processor_.ProcessRepeated(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}

	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitDefinition(definition ast.DefinitionLike) {
	// Visit the possible definition types.
	switch actual := definition.GetAny().(type) {
	case ast.MultilinedLike:
		v.processor_.PreprocessMultilined(actual)
		v.visitMultilined(actual)
		v.processor_.PostprocessMultilined(actual)
	case ast.InlinedLike:
		v.processor_.PreprocessInlined(actual)
		v.visitInlined(actual)
		v.processor_.PostprocessInlined(actual)
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitElement(element ast.ElementLike) {
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
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitExpression(expression ast.ExpressionLike) {
	// Visit the optional comment token.
	var comment = expression.GetOptionalComment()
	if col.IsDefined(comment) {
		v.processor_.ProcessComment(comment)
	}

	// Visit the lowercase token.
	var lowercase = expression.GetLowercase()
	v.processor_.ProcessLowercase(lowercase)

	// Visit the ":" delimiter literal.
	var delimiter = expression.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)

	// Visit the pattern.
	var pattern = expression.GetPattern()
	v.processor_.PreprocessPattern(pattern)
	v.visitPattern(pattern)
	v.processor_.PostprocessPattern(pattern)

	// Visit the optional note token.
	var note = expression.GetOptionalNote()
	if col.IsDefined(note) {
		v.processor_.ProcessNote(note)
	}

	// Visit each newline token.
	var newlineIndex uint
	var newlines = expression.GetNewlines().GetIterator()
	var newlinesSize = uint(newlines.GetSize())
	for newlines.HasNext() {
		newlineIndex++
		var newline = newlines.GetNext()
		v.processor_.ProcessNewline(newline, newlineIndex, newlinesSize)
	}
}

func (v *visitor_) visitExtent(extent ast.ExtentLike) {
	// Visit the ".." delimiter literal.
	var delimiter = extent.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)

	// Visit the runic token.
	var runic = extent.GetRunic()
	v.processor_.ProcessRunic(runic)
}

func (v *visitor_) visitFactor(factor ast.FactorLike) {
	// Visit the possible factor types.
	switch actual := factor.GetAny().(type) {
	case ast.PredicateLike:
		v.processor_.PreprocessPredicate(actual)
		v.visitPredicate(actual)
		v.processor_.PostprocessPredicate(actual)
	case string:
		switch {
		case Scanner().MatchesType(actual, LiteralToken):
			v.processor_.ProcessLiteral(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}

	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitFiltered(filtered ast.FilteredLike) {
	// Visit the optional excluded token.
	var excluded = filtered.GetOptionalExcluded()
	if col.IsDefined(excluded) {
		v.processor_.ProcessExcluded(excluded)
	}

	// Visit the "[" delimiter literal.
	var delimiter = filtered.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)

	// Visit each character.
	var characterIndex uint
	var characters = filtered.GetCharacters().GetIterator()
	var charactersSize = uint(characters.GetSize())
	for characters.HasNext() {
		characterIndex++
		var character = characters.GetNext()
		v.processor_.PreprocessCharacter(character, characterIndex, charactersSize)
		v.visitCharacter(character)
		v.processor_.PostprocessCharacter(character, characterIndex, charactersSize)
	}

	// Visit the "]" delimiter literal.
	var delimiter2 = filtered.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitGrouped(grouped ast.GroupedLike) {
	// Visit the "(" delimiter literal.
	var delimiter = grouped.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)

	// Visit the pattern.
	var pattern = grouped.GetPattern()
	v.processor_.PreprocessPattern(pattern)
	v.visitPattern(pattern)
	v.processor_.PostprocessPattern(pattern)

	// Visit the ")" delimiter literal.
	var delimiter2 = grouped.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitHeader(header ast.HeaderLike) {
	// Visit the comment token.
	var comment = header.GetComment()
	v.processor_.ProcessComment(comment)

	// Visit the newline token.
	var newline = header.GetNewline()
	v.processor_.ProcessNewline(newline, 0, 1)
}

func (v *visitor_) visitIdentifier(identifier ast.IdentifierLike) {
	// Visit the possible identifier types.
	switch actual := identifier.GetAny().(type) {
	case string:
		switch {
		case Scanner().MatchesType(actual, LowercaseToken):
			v.processor_.ProcessLowercase(actual)
		case Scanner().MatchesType(actual, UppercaseToken):
			v.processor_.ProcessUppercase(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}

	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitInlined(inlined ast.InlinedLike) {
	// Visit each factor.
	var factorIndex uint
	var factors = inlined.GetFactors().GetIterator()
	var factorsSize = uint(factors.GetSize())
	for factors.HasNext() {
		factorIndex++
		var factor = factors.GetNext()
		v.processor_.PreprocessFactor(factor, factorIndex, factorsSize)
		v.visitFactor(factor)
		v.processor_.PostprocessFactor(factor, factorIndex, factorsSize)
	}

	// Visit the optional note token.
	var note = inlined.GetOptionalNote()
	if col.IsDefined(note) {
		v.processor_.ProcessNote(note)
	}
}

func (v *visitor_) visitLimit(limit ast.LimitLike) {
	// Visit the ".." delimiter literal.
	var delimiter = limit.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)

	// Visit the optional number token.
	var number = limit.GetOptionalNumber()
	if col.IsDefined(number) {
		v.processor_.ProcessNumber(number)
	}
}

func (v *visitor_) visitLine(line ast.LineLike) {
	// Visit the newline token.
	var newline = line.GetNewline()
	v.processor_.ProcessNewline(newline, 0, 1)

	// Visit the identifier.
	var identifier = line.GetIdentifier()
	v.processor_.PreprocessIdentifier(identifier)
	v.visitIdentifier(identifier)
	v.processor_.PostprocessIdentifier(identifier)

	// Visit the optional note token.
	var note = line.GetOptionalNote()
	if col.IsDefined(note) {
		v.processor_.ProcessNote(note)
	}
}

func (v *visitor_) visitMultilined(multilined ast.MultilinedLike) {
	// Visit each line.
	var lineIndex uint
	var lines = multilined.GetLines().GetIterator()
	var linesSize = uint(lines.GetSize())
	for lines.HasNext() {
		lineIndex++
		var line = lines.GetNext()
		v.processor_.PreprocessLine(line, lineIndex, linesSize)
		v.visitLine(line)
		v.processor_.PostprocessLine(line, lineIndex, linesSize)
	}
}

func (v *visitor_) visitPart(part ast.PartLike) {
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

func (v *visitor_) visitPattern(pattern ast.PatternLike) {
	// Visit the part.
	var part = pattern.GetPart()
	v.processor_.PreprocessPart(part, 0, 1)
	v.visitPart(part)
	v.processor_.PostprocessPart(part, 0, 1)

	// Visit the optional supplement.
	var supplement = pattern.GetOptionalSupplement()
	if col.IsDefined(supplement) {
		v.processor_.PreprocessSupplement(supplement)
		v.visitSupplement(supplement)
		v.processor_.PostprocessSupplement(supplement)
	}
}

func (v *visitor_) visitPredicate(predicate ast.PredicateLike) {
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

func (v *visitor_) visitQuantified(quantified ast.QuantifiedLike) {
	// Visit the "{" delimiter literal.
	var delimiter = quantified.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)

	// Visit the number token.
	var number = quantified.GetNumber()
	v.processor_.ProcessNumber(number)

	// Visit the optional limit.
	var limit = quantified.GetOptionalLimit()
	if col.IsDefined(limit) {
		v.processor_.PreprocessLimit(limit)
		v.visitLimit(limit)
		v.processor_.PostprocessLimit(limit)
	}

	// Visit the "}" delimiter literal.
	var delimiter2 = quantified.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitRule(rule ast.RuleLike) {
	// Visit the optional comment token.
	var comment = rule.GetOptionalComment()
	if col.IsDefined(comment) {
		v.processor_.ProcessComment(comment)
	}

	// Visit the uppercase token.
	var uppercase = rule.GetUppercase()
	v.processor_.ProcessUppercase(uppercase)

	// Visit the ":" delimiter literal.
	var delimiter = rule.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)

	// Visit the definition.
	var definition = rule.GetDefinition()
	v.processor_.PreprocessDefinition(definition)
	v.visitDefinition(definition)
	v.processor_.PostprocessDefinition(definition)

	// Visit each newline token.
	var newlineIndex uint
	var newlines = rule.GetNewlines().GetIterator()
	var newlinesSize = uint(newlines.GetSize())
	for newlines.HasNext() {
		newlineIndex++
		var newline = newlines.GetNext()
		v.processor_.ProcessNewline(newline, newlineIndex, newlinesSize)
	}
}

func (v *visitor_) visitSelective(selective ast.SelectiveLike) {
	// Visit each alternative.
	var alternativeIndex uint
	var alternatives = selective.GetAlternatives().GetIterator()
	var alternativesSize = uint(alternatives.GetSize())
	for alternatives.HasNext() {
		alternativeIndex++
		var alternative = alternatives.GetNext()
		v.processor_.PreprocessAlternative(alternative, alternativeIndex, alternativesSize)
		v.visitAlternative(alternative)
		v.processor_.PostprocessAlternative(alternative, alternativeIndex, alternativesSize)
	}
}

func (v *visitor_) visitSequential(sequential ast.SequentialLike) {
	// Visit each part.
	var partIndex uint = 1
	var parts = sequential.GetParts().GetIterator()
	var partsSize = uint(parts.GetSize())
	for parts.HasNext() {
		partIndex++
		var part = parts.GetNext()
		v.processor_.PreprocessPart(part, partIndex, partsSize)
		v.visitPart(part)
		v.processor_.PostprocessPart(part, partIndex, partsSize)
	}
}

func (v *visitor_) visitSupplement(supplement ast.SupplementLike) {
	// Visit the possible supplement types.
	switch actual := supplement.GetAny().(type) {
	case ast.SequentialLike:
		v.processor_.PreprocessSequential(actual)
		v.visitSequential(actual)
		v.processor_.PostprocessSequential(actual)
	case ast.SelectiveLike:
		v.processor_.PreprocessSelective(actual)
		v.visitSelective(actual)
		v.processor_.PostprocessSelective(actual)
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitSyntax(syntax ast.SyntaxLike) {
	// Visit each header.
	var headerIndex uint
	var headers = syntax.GetHeaders().GetIterator()
	var headersSize = uint(headers.GetSize())
	for headers.HasNext() {
		headerIndex++
		var header = headers.GetNext()
		v.processor_.PreprocessHeader(header, headerIndex, headersSize)
		v.visitHeader(header)
		v.processor_.PostprocessHeader(header, headerIndex, headersSize)
	}

	// Visit each rule.
	var ruleIndex uint
	var rules = syntax.GetRules().GetIterator()
	var rulesSize = uint(rules.GetSize())
	for rules.HasNext() {
		ruleIndex++
		var rule = rules.GetNext()
		v.processor_.PreprocessRule(rule, ruleIndex, rulesSize)
		v.visitRule(rule)
		v.processor_.PostprocessRule(rule, ruleIndex, rulesSize)
	}

	// Visit each expression.
	var expressionIndex uint
	var expressions = syntax.GetExpressions().GetIterator()
	var expressionsSize = uint(expressions.GetSize())
	for expressions.HasNext() {
		expressionIndex++
		var expression = expressions.GetNext()
		v.processor_.PreprocessExpression(expression, expressionIndex, expressionsSize)
		v.visitExpression(expression)
		v.processor_.PostprocessExpression(expression, expressionIndex, expressionsSize)
	}
}

func (v *visitor_) visitTextual(textual ast.TextualLike) {
	// Visit the possible textual types.
	switch actual := textual.GetAny().(type) {
	case string:
		switch {
		case Scanner().MatchesType(actual, IntrinsicToken):
			v.processor_.ProcessIntrinsic(actual)
		case Scanner().MatchesType(actual, RunicToken):
			v.processor_.ProcessRunic(actual)
		case Scanner().MatchesType(actual, LiteralToken):
			v.processor_.ProcessLiteral(actual)
		case Scanner().MatchesType(actual, LowercaseToken):
			v.processor_.ProcessLowercase(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}

	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}
