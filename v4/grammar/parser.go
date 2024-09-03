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
	mat "math"
	stc "strconv"
	sts "strings"
)

// CLASS ACCESS

// Reference

var parserClass = &parserClass_{
	// Initialize the class constants.
	queueSize_: 16,
	stackSize_: 4,
}

// Function

func Parser() ParserClassLike {
	return parserClass
}

// CLASS METHODS

// Target

type parserClass_ struct {
	// Define the class constants.
	queueSize_ uint
	stackSize_ uint
}

// Constructors

func (c *parserClass_) Make() ParserLike {
	return &parser_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type parser_ struct {
	// Define the instance attributes.
	class_  ParserClassLike
	source_ string                   // The original source code.
	tokens_ abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Attributes

func (v *parser_) GetClass() ParserClassLike {
	return v.class_
}

// Public

func (v *parser_) ParseSource(source string) ast.SyntaxLike {
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserClass.queueSize_)
	v.next_ = col.Stack[TokenLike](parserClass.stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse the syntax.
	var syntax, token, ok = v.parseSyntax()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Syntax",
			"Syntax",
		)
		panic(message)
	}

	// Found the syntax.
	return syntax
}

// Private

func (v *parser_) parseAlternative() (
	alternative ast.AlternativeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more repetitions.
	var repetition ast.RepetitionLike
	repetition, token, ok = v.parseRepetition()
	if !ok {
		// This is not an alternative.
		return alternative, token, false
	}
	var repetitions = col.List[ast.RepetitionLike]()
	for ok {
		repetitions.AppendValue(repetition)
		repetition, token, ok = v.parseRepetition()
	}

	// Found an alternative.
	alternative = ast.Alternative().Make(repetitions)
	return alternative, token, true
}

func (v *parser_) parseBracket() (
	bracket ast.BracketLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a "(" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		// This is not a bracket.
		return bracket, token, false
	}

	// Attempt to parse one or more factors.
	var factor ast.FactorLike
	factor, token, ok = v.parseFactor()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Factor",
			"Bracket",
			"Factor",
			"Cardinality",
		)
		panic(message)
	}
	var factors = col.List[ast.FactorLike]()
	for ok {
		factors.AppendValue(factor)
		factor, _, ok = v.parseFactor()
	}

	// Attempt to parse a ")" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Bracket",
			"Factor",
			"Cardinality",
		)
		panic(message)
	}

	// Attempt to parse a cardinality.
	var cardinality ast.CardinalityLike
	cardinality, token, _ = v.parseCardinality()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Cardinality",
			"Bracket",
			"Factor",
			"Cardinality",
		)
		panic(message)
	}

	// Found a bracket term.
	bracket = ast.Bracket().Make(factors, cardinality)
	return bracket, token, true
}

func (v *parser_) parseCardinality() (
	cardinality ast.CardinalityLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a constraint cardinality.
	var constraint ast.ConstraintLike
	constraint, token, ok = v.parseConstraint()
	if ok {
		// Found a constraint cardinality.
		cardinality = ast.Cardinality().Make(constraint)
		return cardinality, token, true
	}

	// Attempt to parse a count cardinality.
	var count ast.CountLike
	count, token, ok = v.parseCount()
	if ok {
		// Found a count cardinality.
		cardinality = ast.Cardinality().Make(count)
		return cardinality, token, true
	}

	// This is not a cardinality.
	return cardinality, token, false
}

func (v *parser_) parseCharacter() (
	character ast.CharacterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific character.
	var specific ast.SpecificLike
	specific, token, ok = v.parseSpecific()
	if ok {
		// Found a specific character.
		character = ast.Character().Make(specific)
		return character, token, true
	}

	// Attempt to parse an intrinsic character.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken, "")
	if ok {
		// Found an intrinsic character.
		character = ast.Character().Make(intrinsic)
		return character, token, true
	}

	// This is not a character.
	return character, token, false
}

func (v *parser_) parseConstraint() (
	constraint ast.ConstraintLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an optional constraint cardinality.
	var optional string
	optional, token, ok = v.parseToken(OptionalToken, "")
	if ok {
		// Found an optional constraint cardinality.
		constraint = ast.Constraint().Make(optional)
		return constraint, token, true
	}

	// Attempt to parse a repeated constraint cardinality.
	var repeated string
	repeated, token, ok = v.parseToken(RepeatedToken, "")
	if ok {
		// Found a repeated constraint cardinality.
		constraint = ast.Constraint().Make(repeated)
		return constraint, token, true
	}

	// This is not a constraint cardinality.
	return constraint, token, false
}

func (v *parser_) parseCount() (
	count ast.CountLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a "{" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		// This is not a count cardinality.
		return count, token, false
	}

	// Attempt to parse the first number.
	var number string
	number, token, ok = v.parseToken(NumberToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("number",
			"Count",
		)
		panic(message)
	}
	var numbers = col.List[string]()
	numbers.AppendValue(number)

	// Attempt to parse an optional additional number range.
	_, _, ok = v.parseToken(DelimiterToken, "..")
	if ok {
		number, _, ok = v.parseToken(NumberToken, "")
		if !ok {
			number = stc.Itoa(mat.MaxInt)
		}
		numbers.AppendValue(number)
	}

	// Attempt to parse a "}" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("}",
			"Count",
		)
		panic(message)
	}

	// Found a count cardinality.
	count = ast.Count().Make(numbers)
	return count, token, true
}

func (v *parser_) parseDefinition() (
	definition ast.DefinitionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a multiline definition.
	var multiline ast.MultilineLike
	multiline, token, ok = v.parseMultiline()
	if ok {
		// Found a multiline definition.
		definition = ast.Definition().Make(multiline)
		return definition, token, true
	}

	// Attempt to parse an inline definition.
	var inline ast.InlineLike
	inline, token, ok = v.parseInline()
	if ok {
		// Found an inline definition.
		definition = ast.Definition().Make(inline)
		return definition, token, true
	}

	// This is not a definition.
	return definition, token, false
}

func (v *parser_) parseElement() (
	element ast.ElementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a group element.
	var group ast.GroupLike
	group, token, ok = v.parseGroup()
	if ok {
		// Found a group element.
		element = ast.Element().Make(group)
		return element, token, true
	}

	// Attempt to parse a filter element.
	var filter ast.FilterLike
	filter, token, ok = v.parseFilter()
	if ok {
		// Found a filter element.
		element = ast.Element().Make(filter)
		return element, token, true
	}

	// Attempt to parse a text element.
	var text ast.TextLike
	text, token, ok = v.parseText()
	if ok {
		// Found a text element.
		element = ast.Element().Make(text)
		return element, token, true
	}

	// This is not an element.
	return element, token, false
}

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an optional comment.
	var comment string
	var commentToken TokenLike
	comment, commentToken, _ = v.parseToken(CommentToken, "")

	// Attempt to parse a lowercase identifier.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken, "")
	if !ok {
		// This is not an expression, so put back any comment token.
		if col.IsDefined(comment) {
			v.putBack(commentToken)
		}
		return expression, token, false
	}

	// Attempt to parse a ":" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ":")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(":",
			"Expression",
			"Pattern",
		)
		panic(message)
	}

	// Attempt to parse a pattern.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Pattern",
			"Expression",
			"Pattern",
		)
		panic(message)
	}

	// Attempt to parse an optional note.
	var note string
	note, _, _ = v.parseToken(NoteToken, "")

	// Attempt to parse one or more newline characters.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("newline",
			"Expression",
			"Pattern",
		)
		panic(message)
	}
	var newlines = col.List[string]()
	for ok {
		newlines.AppendValue(newline)
		newline, token, ok = v.parseToken(NewlineToken, "")
	}

	// Found an expression.
	expression = ast.Expression().Make(
		comment,
		lowercase,
		pattern,
		note,
		newlines,
	)
	return expression, token, true
}

func (v *parser_) parseFactor() (
	factor ast.FactorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a reference factor.
	var reference ast.ReferenceLike
	reference, token, ok = v.parseReference()
	if ok {
		// Found a reference factor.
		factor = ast.Factor().Make(reference)
		return factor, token, true
	}

	// Attempt to parse a literal factor.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken, "")
	if ok {
		// Found a literal factor.
		factor = ast.Factor().Make(literal)
		return factor, token, true
	}

	// This is not a factor.
	return factor, token, false
}

func (v *parser_) parseFilter() (
	filter ast.FilterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an optional excluded character token.
	var excluded string
	var excludedToken TokenLike
	excluded, excludedToken, _ = v.parseToken(ExcludedToken, "")

	// Attempt to parse a "[" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	if !ok {
		// This is not a filter element, so put back any excluded character token.
		if col.IsDefined(excluded) {
			v.putBack(excludedToken)
		}
		return filter, token, false
	}

	// Attempt to parse one or more filter characters.
	var character ast.CharacterLike
	character, token, ok = v.parseCharacter()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Character",
			"Filter",
			"Character",
		)
		panic(message)
	}
	var characters = col.List[ast.CharacterLike]()
	for ok {
		characters.AppendValue(character)
		character, _, ok = v.parseCharacter()
	}

	// Attempt to parse a "]" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "]")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("]",
			"Filter",
			"Character",
		)
		panic(message)
	}

	// Found a filter element.
	filter = ast.Filter().Make(excluded, characters)
	return filter, token, true
}

func (v *parser_) parseGroup() (
	group ast.GroupLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a "(" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		// This is not a group element.
		return group, token, false
	}

	// Attempt to parse a pattern.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Pattern",
			"Group",
			"Pattern",
		)
		panic(message)
	}

	// Attempt to parse a ")" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Group",
			"Pattern",
		)
		panic(message)
	}

	// Found a group element.
	group = ast.Group().Make(pattern)
	return group, token, true
}

func (v *parser_) parseHeader() (
	header ast.HeaderLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a comment.
	var comment string
	var commentToken TokenLike
	comment, commentToken, ok = v.parseToken(CommentToken, "")
	if !ok {
		// This is not a header.
		return header, commentToken, false
	}

	// Attempt to parse a newline character.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken, "")
	if !ok {
		// This is not a header, so put back the comment token.
		v.putBack(commentToken)
		return header, token, false
	}

	// Found a header.
	header = ast.Header().Make(comment, newline)
	return header, token, true
}

func (v *parser_) parseIdentifier() (
	identifier ast.IdentifierLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a lowercase identifier.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken, "")
	if ok {
		// Found a lowercase identifier.
		identifier = ast.Identifier().Make(lowercase)
		return identifier, token, true
	}

	// Attempt to parse an uppercase identifier.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken, "")
	if ok {
		// Found an uppercase identifier.
		identifier = ast.Identifier().Make(uppercase)
		return identifier, token, true
	}

	// This is not an identifier.
	return identifier, token, false
}

func (v *parser_) parseInline() (
	inline ast.InlineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more terms.
	var term ast.TermLike
	term, token, ok = v.parseTerm()
	if !ok {
		// This is not an inline definition.
		return inline, token, false
	}
	var terms = col.List[ast.TermLike]()
	for ok {
		terms.AppendValue(term)
		term, _, ok = v.parseTerm()
	}

	// Attempt to parse an optional note.
	var note string
	note, token, _ = v.parseToken(NoteToken, "")

	// Found an inline definition.
	inline = ast.Inline().Make(terms, note)
	return inline, token, true
}

func (v *parser_) parseLine() (
	line ast.LineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a newline character.
	var newline string
	var newlineToken TokenLike
	newline, newlineToken, ok = v.parseToken(NewlineToken, "")
	if !ok {
		// This is not a line.
		return line, newlineToken, false
	}

	// Attempt to parse an identifier.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	if !ok {
		// This is not a line, so put back the newline token.
		v.putBack(newlineToken)
		return line, token, false
	}

	// Attempt to parse an optional note.
	var note string
	note, token, _ = v.parseToken(NoteToken, "")

	// Found a line.
	line = ast.Line().Make(newline, identifier, note)
	return line, token, true
}

func (v *parser_) parseMultiline() (
	multiline ast.MultilineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more lines.
	var line ast.LineLike
	line, token, ok = v.parseLine()
	if !ok {
		// This is not a multiline definition.
		return multiline, token, false
	}
	var lines = col.List[ast.LineLike]()
	for ok {
		lines.AppendValue(line)
		line, _, ok = v.parseLine()
	}

	// Found a multiline definition.
	multiline = ast.Multiline().Make(lines)
	return multiline, token, true
}

func (v *parser_) parsePattern() (
	pattern ast.PatternLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an alternative.
	var alternative ast.AlternativeLike
	alternative, token, ok = v.parseAlternative()
	if !ok {
		// This is not a pattern.
		return pattern, token, false
	}

	// Attempt to parse additional alternatives.
	var alternatives = col.List[ast.AlternativeLike]()
	alternatives.AppendValue(alternative)
	_, token, ok = v.parseToken(DelimiterToken, "|")
	for ok {
		alternative, token, ok = v.parseAlternative()
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("Alternative",
				"Pattern",
				"Alternative",
			)
			panic(message)
		}
		alternatives.AppendValue(alternative)
		_, _, ok = v.parseToken(DelimiterToken, "|")
	}

	// Found a pattern.
	pattern = ast.Pattern().Make(alternatives)
	return pattern, token, true
}

func (v *parser_) parseReference() (
	reference ast.ReferenceLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an identifier.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	if !ok {
		// This is not a reference.
		return reference, token, false
	}

	// Attempt to parse an optional cardinality.
	var cardinality ast.CardinalityLike
	cardinality, token, _ = v.parseCardinality()

	// Found a reference.
	reference = ast.Reference().Make(identifier, cardinality)
	return reference, token, true
}

func (v *parser_) parseRepetition() (
	repetition ast.RepetitionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an element.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	if !ok {
		// This is not a repetition.
		return repetition, token, false
	}

	// Attempt to parse an optional cardinality.
	var cardinality ast.CardinalityLike
	cardinality, token, _ = v.parseCardinality()

	// Found a repetition.
	repetition = ast.Repetition().Make(element, cardinality)
	return repetition, token, true
}

func (v *parser_) parseRule() (
	rule ast.RuleLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an optional comment.
	var comment string
	var commentToken TokenLike
	comment, commentToken, _ = v.parseToken(CommentToken, "")

	// Attempt to parse an uppercase identifier.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken, "")
	if !ok {
		// This is not a rule, so put back any comment token.
		if col.IsDefined(comment) {
			v.putBack(commentToken)
		}
		return rule, token, false
	}

	// Attempt to parse a ":" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ":")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(":",
			"Rule",
			"Definition",
		)
		panic(message)
	}

	// Attempt to parse a definition.
	var definition ast.DefinitionLike
	definition, token, ok = v.parseDefinition()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Definition",
			"Rule",
			"Definition",
		)
		panic(message)
	}

	// Attempt to parse one or more newline characters.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("newline",
			"Rule",
			"Definition",
		)
		panic(message)
	}
	var newlines = col.List[string]()
	for ok {
		newlines.AppendValue(newline)
		newline, token, ok = v.parseToken(NewlineToken, "")
	}

	// Found a rule.
	rule = ast.Rule().Make(comment, uppercase, definition, newlines)
	return rule, token, true
}

func (v *parser_) parseSpecific() (
	specific ast.SpecificLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the first runic.
	var runic string
	runic, token, ok = v.parseToken(RunicToken, "")
	if !ok {
		// This is not a specific character.
		return specific, token, false
	}
	var runics = col.List[string]()
	runics.AppendValue(runic)

	// Attempt to parse an optional second runic.
	_, token, ok = v.parseToken(DelimiterToken, "..")
	if ok {
		runic, token, ok = v.parseToken(RunicToken, "")
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("runic",
				"Specific",
			)
			panic(message)
		}
		runics.AppendValue(runic)
	}

	// Found a specific character.
	specific = ast.Specific().Make(runics)
	return specific, token, true
}

func (v *parser_) parseSyntax() (
	syntax ast.SyntaxLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more headers.
	var header ast.HeaderLike
	header, token, ok = v.parseHeader()
	if !ok {
		// This is not a syntax.
		return syntax, token, false
	}
	var headers = col.List[ast.HeaderLike]()
	for ok {
		headers.AppendValue(header)
		header, _, ok = v.parseHeader()
	}

	// Attempt to parse one or more rules.
	var rule ast.RuleLike
	rule, token, ok = v.parseRule()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Rule",
			"Syntax",
			"Header",
			"Rule",
			"Expression",
		)
		panic(message)
	}
	var rules = col.List[ast.RuleLike]()
	for ok {
		rules.AppendValue(rule)
		rule, _, ok = v.parseRule()
	}

	// Attempt to parse one or more expressions.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Expression",
			"Syntax",
			"Header",
			"Rule",
			"Expression",
		)
		panic(message)
	}
	var expressions = col.List[ast.ExpressionLike]()
	for ok {
		expressions.AppendValue(expression)
		expression, _, ok = v.parseExpression()
	}

	// Sort the expressions alphabetically by name.
	expressions.SortValuesWithRanker(
		func(first, second ast.ExpressionLike) col.Rank {
			var firstName = first.GetLowercase()
			var secondName = second.GetLowercase()
			switch {
			case firstName < secondName:
				return col.LesserRank
			case firstName > secondName:
				return col.GreaterRank
			default:
				return col.EqualRank
			}
		},
	)

	// Found a syntax.
	syntax = ast.Syntax().Make(headers, rules, expressions)
	return syntax, token, true
}

func (v *parser_) parseTerm() (
	term ast.TermLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a factor term.
	var factor ast.FactorLike
	factor, token, ok = v.parseFactor()
	if ok {
		// Found a factor term.
		term = ast.Term().Make(factor)
		return term, token, true
	}

	// Attempt to parse a bracket term.
	var bracket ast.BracketLike
	bracket, token, ok = v.parseBracket()
	if ok {
		// Found a bracket term.
		term = ast.Term().Make(bracket)
		return term, token, true
	}

	// This is not a term.
	return term, token, false
}

func (v *parser_) parseText() (
	text ast.TextLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an intrinsic text element.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken, "")
	if ok {
		// Found an intrinsic text element.
		text = ast.Text().Make(intrinsic)
		return text, token, true
	}

	// Attempt to parse a runic text element.
	var runic string
	runic, token, ok = v.parseToken(RunicToken, "")
	if ok {
		// Found a runic text element.
		text = ast.Text().Make(runic)
		return text, token, true
	}

	// Attempt to parse a literal text element.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken, "")
	if ok {
		// Found a literal text element.
		text = ast.Text().Make(literal)
		return text, token, true
	}

	// Attempt to parse a lowercase text element.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken, "")
	if ok {
		// Found a lowercase text element.
		text = ast.Text().Make(lowercase)
		return text, token, true
	}

	// This is not a text element.
	return text, token, false
}

func (v *parser_) parseToken(expectedType TokenType, expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token.
	token = v.getNextToken()
	if token == nil {
		// We are at the end-of-file marker.
		return value, token, false
	}
	if token.GetType() == expectedType {
		value = token.GetValue()
		if col.IsUndefined(expectedValue) || value == expectedValue {
			// Found the expected token.
			return value, token, true
		}
		value = "" // We must reset this!
	}

	// This is not the expected token.
	v.putBack(token)
	return value, token, false
}

func (v *parser_) formatError(token TokenLike) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		Scanner().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source line with the error in it.
	message += "\033[36m"
	if line > 1 {
		message += fmt.Sprintf("%04d: ", line-1) + string(lines[line-2]) + "\n"
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"

	return message
}

func (v *parser_) generateSyntax(expected string, names ...string) string {
	var message = "Was expecting '" + expected + "' from:\n"
	for _, name := range names {
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			name,
			syntax_[name],
		)
	}
	return message
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveTop()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveHead() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError(token)
		panic(message)
	}

	return token
}

func (v *parser_) putBack(token TokenLike) {
	v.next_.AddValue(token)
}

var syntax_ = map[string]string{
	"Syntax": `Header+ Rule+ Expression+`,
	"Header": `comment newline`,
	"Rule":   `comment? uppercase ":" Definition newline+`,
	"Definition": `
    Inline
    Multiline`,
	"Multiline": `Line+`,
	"Line":      `newline Identifier note?`,
	"Identifier": `
    lowercase
    uppercase`,
	"Inline": `Term+ note?`,
	"Term": `
    Factor
    Bracket`,
	"Bracket": `"(" Factor+ ")" Cardinality`,
	"Factor": `
    Reference
    literal`,
	"Reference": `Identifier Cardinality?  ! The default cardinality is one.`,
	"Cardinality": `
    Constraint
    Count`,
	"Constraint": `
    optional
    repeated`,
	"Count":       `"{" number (".." number?)? "}"  ! The range of numbers is inclusive.`,
	"Expression":  `comment? lowercase ":" Pattern note? newline+`,
	"Pattern":     `Alternative ("|" Alternative)*`,
	"Alternative": `Repetition+`,
	"Repetition":  `Element Cardinality?  ! The default cardinality is one.`,
	"Element": `
    Group
    Filter
    Text`,
	"Group":  `"(" Pattern ")"`,
	"Filter": `excluded? "[" Character+ "]"`,
	"Character": `
    Specific
    intrinsic`,
	"Specific": `runic (".." runic)?  ! The range of runic elements is inclusive.`,
	"Text": `
    intrinsic
    runic
    literal
    lowercase`,
}
