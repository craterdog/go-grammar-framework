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
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Attributes

func (v *parser_) GetClass() ParserClassLike {
	return v.class_
}

// Public

func (v *parser_) ParseSource(source string) ast.SyntaxLike {
	v.source_ = source
	var notation = fwk.CDCN()
	v.tokens_ = col.Queue[TokenLike](notation).MakeWithCapacity(parserClass.queueSize_)
	v.next_ = col.Stack[TokenLike](notation).MakeWithCapacity(parserClass.stackSize_)

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
	var count = 0
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < len(lines) {
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
			syntax[name],
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
		panic("The token channel terminated without an EOF token.")
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError(token)
		panic(message)
	}

	return token
}

func (v *parser_) parseAlternative() (
	alternative ast.AlternativeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the "|" delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "|")
	if !ok {
		// This is not the alternative.
		return alternative, token, false
	}

	// Attempt to parse one or more parts.
	var part ast.PartLike
	part, token, ok = v.parsePart()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Part",
			"Alternative",
			"Part",
		)
		panic(message)
	}
	var notation = fwk.CDCN()
	var parts = col.List[ast.PartLike](notation).Make()
	for ok {
		parts.AppendValue(part)
		part, token, ok = v.parsePart()
	}

	// Found the alternative.
	alternative = ast.Alternative().Make(parts)
	return alternative, token, true
}

func (v *parser_) parseBounded() (
	bounded ast.BoundedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the initial rune.
	var initial ast.InitialLike
	initial, token, ok = v.parseInitial()
	if !ok {
		// This is not the bounded.
		return bounded, token, false
	}

	// Attempt to parse the optional extent rune.
	var extent ast.ExtentLike
	extent, token, _ = v.parseExtent()

	// Found the bounded.
	bounded = ast.Bounded().Make(initial, extent)
	return bounded, token, true
}

func (v *parser_) parseCardinality() (
	cardinality ast.CardinalityLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the constrained cardinality.
	var constrained ast.ConstrainedLike
	constrained, token, ok = v.parseConstrained()
	if ok {
		// Found the constrained cardinality.
		cardinality = ast.Cardinality().Make(constrained)
		return cardinality, token, true
	}

	// Attempt to parse the quantified cardinality.
	var quantified string
	quantified, token, ok = v.parseToken(QuantifiedToken, "")
	if ok {
		// Found the quantified cardinality.
		cardinality = ast.Cardinality().Make(quantified)
		return cardinality, token, true
	}

	// This is not the cardinality.
	return cardinality, token, false
}

func (v *parser_) parseCharacter() (
	character ast.CharacterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the bounded character.
	var bounded ast.BoundedLike
	bounded, token, ok = v.parseBounded()
	if ok {
		// Found the bounded character.
		character = ast.Character().Make(bounded)
		return character, token, true
	}

	// Attempt to parse the intrinsic character.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken, "")
	if ok {
		// Found the intrinsic character.
		character = ast.Character().Make(intrinsic)
		return character, token, true
	}

	// This is not the character.
	return character, token, false
}

func (v *parser_) parseConstrained() (
	constrained ast.ConstrainedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the opening bracket for the constrained.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		// This is not the constrained.
		return constrained, token, false
	}

	// Attempt to parse the minimum number for the constrained.
	var minimum ast.MinimumLike
	minimum, token, ok = v.parseMinimum()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Minimum",
			"Constrained",
			"Minimum",
			"Maximum",
		)
		panic(message)
	}

	// Attempt to parse the optional maximum number for the constrained.
	var maximum ast.MaximumLike
	maximum, _, _ = v.parseMaximum()

	// Attempt to parse the closing bracket for the constrained.
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("}",
			"Constrained",
			"Minimum",
			"Maximum",
		)
		panic(message)
	}

	// Found the constrained.
	constrained = ast.Constrained().Make(minimum, maximum)
	return constrained, token, true
}

func (v *parser_) parseElement() (
	element ast.ElementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the grouped element.
	var grouped ast.GroupedLike
	grouped, token, ok = v.parseGrouped()
	if ok {
		// Found the grouped element.
		element = ast.Element().Make(grouped)
		return element, token, true
	}

	// Attempt to parse the filtered element.
	var filtered ast.FilteredLike
	filtered, token, ok = v.parseFiltered()
	if ok {
		// Found the filtered element.
		element = ast.Element().Make(filtered)
		return element, token, true
	}

	// Attempt to parse the bounded element.
	var bounded ast.BoundedLike
	bounded, token, ok = v.parseBounded()
	if ok {
		// Found the character element.
		element = ast.Element().Make(bounded)
		return element, token, true
	}

	// Attempt to parse the intrinsic element.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken, "")
	if ok {
		// Found the intrinsic element.
		element = ast.Element().Make(intrinsic)
		return element, token, true
	}

	// Attempt to parse the lowercase element.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken, "")
	if ok {
		// Found the lowercase element.
		element = ast.Element().Make(lowercase)
		return element, token, true
	}

	// Attempt to parse the literal element.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken, "")
	if ok {
		// Found the literal element.
		element = ast.Element().Make(literal)
		return element, token, true
	}

	// This is not the element.
	return element, token, false
}

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the in-line expression.
	var inlined ast.InlinedLike
	inlined, token, ok = v.parseInlined()
	if ok {
		// Found the in-line expression.
		expression = ast.Expression().Make(inlined)
		return expression, token, true
	}

	// Attempt to parse the multi-line expression.
	var multilined ast.MultilinedLike
	multilined, token, ok = v.parseMultilined()
	if ok {
		// Found the multi-line expression.
		expression = ast.Expression().Make(multilined)
		return expression, token, true
	}

	// This is not the expression.
	return expression, token, false
}

func (v *parser_) parseExtent() (
	extent ast.ExtentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the dot-dot delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "..")
	if !ok {
		// This is not the extent rune.
		return extent, token, false
	}

	// Attempt to parse the extent rune.
	var rune_ string
	rune_, token, ok = v.parseToken(RuneToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("rune",
			"Extent",
		)
		panic(message)
	}

	// Found the extent rune.
	extent = ast.Extent().Make(rune_)
	return extent, token, true
}

func (v *parser_) parseFactor() (
	factor ast.FactorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the predicate.
	var predicate ast.PredicateLike
	predicate, token, ok = v.parsePredicate()
	if !ok {
		// This is not the factor.
		return factor, token, false
	}

	// Attempt to parse the optional cardinality.
	var cardinality ast.CardinalityLike
	cardinality, token, _ = v.parseCardinality()

	// Found the factor.
	factor = ast.Factor().Make(predicate, cardinality)
	return factor, token, true
}

func (v *parser_) parseFiltered() (
	filtered ast.FilteredLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the optional negation for the filtered element.
	var negation string
	var negationToken TokenLike
	negation, negationToken, _ = v.parseToken(NegationToken, "")

	// Attempt to parse the opening bracket for the filtered element.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	if !ok {
		// This is not the filtered element, put back any negation token.
		if fwk.IsDefined(negation) {
			v.putBack(negationToken)
		}
		return filtered, token, false
	}

	// Attempt to parse one or more characters.
	var character ast.CharacterLike
	character, token, ok = v.parseCharacter()
	if !ok {
		// This is not the filtered element.
		return filtered, token, false
	}
	var notation = fwk.CDCN()
	var characters = col.List[ast.CharacterLike](notation).Make()
	for ok {
		characters.AppendValue(character)
		character, _, ok = v.parseCharacter()
	}

	// Attempt to parse the closing bracket for the filtered element.
	_, token, ok = v.parseToken(DelimiterToken, "]")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("]",
			"Filtered",
			"Character",
		)
		panic(message)
	}

	// Found the filtered element.
	filtered = ast.Filtered().Make(negation, characters)
	return filtered, token, true
}

func (v *parser_) parseGrouped() (
	grouped ast.GroupedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the opening delimiter for the grouped.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		// This is not the grouped.
		return grouped, token, false
	}

	// Attempt to parse the pattern.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Pattern",
			"Grouped",
			"Pattern",
		)
		panic(message)
	}

	// Attempt to parse the closing delimiter for the grouped.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Grouped",
			"Pattern",
		)
		panic(message)
	}

	// Found the grouped.
	grouped = ast.Grouped().Make(pattern)
	return grouped, token, true
}

func (v *parser_) parseHeader() (
	header ast.HeaderLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the comment.
	var comment string
	var commentToken TokenLike
	comment, commentToken, ok = v.parseToken(CommentToken, "")
	if !ok {
		// This is not the header.
		return header, commentToken, false
	}

	// Attempt to parse the end-of-line character.
	_, token, ok = v.parseToken(EOLToken, "")
	if !ok {
		// This is not the header, put back the comment token.
		v.putBack(commentToken)
		return header, token, false
	}

	// Found the header.
	header = ast.Header().Make(comment)
	return header, token, true
}

func (v *parser_) parseIdentifier() (
	identifier ast.IdentifierLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the lowercase identifier.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken, "")
	if ok {
		identifier = ast.Identifier().Make(lowercase)
		return identifier, token, true
	}

	// Attempt to parse the uppercase identifier.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken, "")
	if ok {
		identifier = ast.Identifier().Make(uppercase)
		return identifier, token, true
	}

	// This is not the identifier.
	return identifier, token, false
}

func (v *parser_) parseInitial() (
	initial ast.InitialLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the initial rune.
	var rune_ string
	rune_, token, ok = v.parseToken(RuneToken, "")
	if !ok {
		// This is not the initial rune.
		return initial, token, false
	}

	// Found the initial rune.
	initial = ast.Initial().Make(rune_)
	return initial, token, true
}

func (v *parser_) parseInlined() (
	inlined ast.InlinedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more factors.
	var factor ast.FactorLike
	factor, token, ok = v.parseFactor()
	if !ok {
		// This is not the inlined expression.
		return inlined, token, false
	}
	var notation = fwk.CDCN()
	var factors = col.List[ast.FactorLike](notation).Make()
	for ok {
		factors.AppendValue(factor)
		factor, _, ok = v.parseFactor()
	}

	// Attempt to parse the optional note.
	var note string
	note, token, _ = v.parseToken(NoteToken, "")

	// Found the in-line expression.
	inlined = ast.Inlined().Make(factors, note)
	return inlined, token, true
}

func (v *parser_) parseLexigram() (
	lexigram ast.LexigramLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the optional comment.
	var comment string
	var commentToken TokenLike
	comment, commentToken, _ = v.parseToken(CommentToken, "")

	// Attempt to parse the lowercase identifier.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken, "")
	if !ok {
		// This is not the lexigram, put back any comment token that was received.
		if fwk.IsDefined(comment) {
			v.putBack(commentToken)
		}
		return lexigram, token, false
	}

	// Attempt to parse the separator delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ":")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(":",
			"Lexigram",
			"Pattern",
		)
		panic(message)
	}

	// Attempt to parse the pattern.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Pattern",
			"Lexigram",
			"Pattern",
		)
		panic(message)
	}

	// Attempt to parse the optional note.
	var note string
	note, _, _ = v.parseToken(NoteToken, "")

	// Attempt to parse one or more end-of-line characters.
	_, token, ok = v.parseToken(EOLToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("EOL",
			"Lexigram",
			"Pattern",
		)
		panic(message)
	}
	for ok {
		_, token, ok = v.parseToken(EOLToken, "")
	}

	// Found the lexigram.
	lexigram = ast.Lexigram().Make(comment, lowercase, pattern, note)
	return lexigram, token, true
}

func (v *parser_) parseLine() (
	line ast.LineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the end-of-line character.
	var eolToken TokenLike
	_, eolToken, ok = v.parseToken(EOLToken, "")
	if !ok {
		// This is not the line.
		return line, eolToken, false
	}

	// Attempt to parse the identifier.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	if !ok {
		// This is not the line, put back the end-of-line token.
		v.putBack(eolToken)
		return line, token, false
	}

	// Attempt to parse the optional note.
	var note string
	note, token, _ = v.parseToken(NoteToken, "")

	// Found the line.
	line = ast.Line().Make(identifier, note)
	return line, token, true
}

func (v *parser_) parseMaximum() (
	maximum ast.MaximumLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the dot-dot delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "..")
	if !ok {
		// This is not the maximum number.
		return maximum, token, false
	}

	// Attempt to parse the optional maximum number.
	var number string
	number, token, _ = v.parseToken(NumberToken, "")

	// Found the maximum number.
	maximum = ast.Maximum().Make(number)
	return maximum, token, true
}

func (v *parser_) parseMinimum() (
	minimum ast.MinimumLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the minimum number.
	var number string
	number, token, ok = v.parseToken(NumberToken, "")
	if !ok {
		// This is not the minimum number.
		return minimum, token, false
	}

	// Found the minimum number.
	minimum = ast.Minimum().Make(number)
	return minimum, token, true
}

func (v *parser_) parseMultilined() (
	multilined ast.MultilinedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more lines.
	var line ast.LineLike
	line, token, ok = v.parseLine()
	if !ok {
		// This is not the multilined expression.
		return multilined, token, false
	}
	var notation = fwk.CDCN()
	var lines = col.List[ast.LineLike](notation).Make()
	for ok {
		lines.AppendValue(line)
		line, _, ok = v.parseLine()
	}

	// Found the multi-line expression.
	multilined = ast.Multilined().Make(lines)
	return multilined, token, true
}

func (v *parser_) parsePart() (
	part ast.PartLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the element.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	if !ok {
		// This is not the part.
		return part, token, false
	}

	// Attempt to parse the optional cardinality.
	var cardinality ast.CardinalityLike
	cardinality, token, _ = v.parseCardinality()

	// Found the part.
	part = ast.Part().Make(element, cardinality)
	return part, token, true
}

func (v *parser_) parsePattern() (
	pattern ast.PatternLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more parts.
	var part ast.PartLike
	part, token, ok = v.parsePart()
	if !ok {
		// This is not the pattern.
		return pattern, token, false
	}
	var notation = fwk.CDCN()
	var parts = col.List[ast.PartLike](notation).Make()
	for ok {
		parts.AppendValue(part)
		part, _, ok = v.parsePart()
	}

	// Attempt to parse zero or more alternatives.
	var alternative ast.AlternativeLike
	alternative, token, ok = v.parseAlternative()
	var alternatives = col.List[ast.AlternativeLike](notation).Make()
	for ok {
		alternatives.AppendValue(alternative)
		alternative, token, ok = v.parseAlternative()
	}

	// Found the pattern.
	pattern = ast.Pattern().Make(parts, alternatives)
	return pattern, token, true
}

func (v *parser_) parsePredicate() (
	predicate ast.PredicateLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the lowercase predicate.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken, "")
	if ok {
		predicate = ast.Predicate().Make(lowercase)
		return predicate, token, true
	}

	// Attempt to parse the uppercase predicate.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken, "")
	if ok {
		predicate = ast.Predicate().Make(uppercase)
		return predicate, token, true
	}

	// Attempt to parse the intrinsic predicate.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken, "")
	if ok {
		predicate = ast.Predicate().Make(intrinsic)
		return predicate, token, true
	}

	// Attempt to parse the literal predicate.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken, "")
	if ok {
		predicate = ast.Predicate().Make(literal)
		return predicate, token, true
	}

	// This is not the predicate.
	return predicate, token, false
}

func (v *parser_) parseRule() (
	rule ast.RuleLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the optional comment.
	var comment string
	var commentToken TokenLike
	comment, commentToken, _ = v.parseToken(CommentToken, "")

	// Attempt to parse the uppercase identifier.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken, "")
	if !ok {
		// This is not the rule, put back any comment token that was received.
		if fwk.IsDefined(comment) {
			v.putBack(commentToken)
		}
		return rule, token, false
	}

	// Attempt to parse the separator delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ":")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(":",
			"Rule",
			"Expression",
		)
		panic(message)
	}

	// Attempt to parse the expression.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Expression",
			"Rule",
			"Expression",
		)
		panic(message)
	}

	// Attempt to parse one or more end-of-line characters.
	_, token, ok = v.parseToken(EOLToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("EOL",
			"Rule",
			"Expression",
		)
		panic(message)
	}
	for ok {
		_, token, ok = v.parseToken(EOLToken, "")
	}

	// Found the rule.
	rule = ast.Rule().Make(comment, uppercase, expression)
	return rule, token, true
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
		// This is not the syntax.
		return syntax, token, false
	}
	var notation = fwk.CDCN()
	var headers = col.List[ast.HeaderLike](notation).Make()
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
			"Lexigram",
		)
		panic(message)
	}
	var rules = col.List[ast.RuleLike](notation).Make()
	for ok {
		rules.AppendValue(rule)
		rule, _, ok = v.parseRule()
	}

	// Attempt to parse one or more lexigrams.
	var lexigram ast.LexigramLike
	lexigram, token, ok = v.parseLexigram()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Lexigram",
			"Syntax",
			"Header",
			"Rule",
			"Lexigram",
		)
		panic(message)
	}
	var lexigrams = col.List[ast.LexigramLike](notation).Make()
	for ok {
		lexigrams.AppendValue(lexigram)
		lexigram, _, ok = v.parseLexigram()
	}

	// Attempt to parse optional end-of-line characters.
	for ok {
		_, _, ok = v.parseToken(EOLToken, "")
	}

	// Attempt to parse the end-of-file marker.
	_, token, ok = v.parseToken(EOFToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("EOF",
			"Syntax",
			"Header",
			"Rule",
			"Lexigram",
		)
		panic(message)
	}

	// Found the syntax.
	syntax = ast.Syntax().Make(headers, rules, lexigrams)
	return syntax, token, true
}

func (v *parser_) parseToken(expectedType TokenType, expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the specific token.
	token = v.getNextToken()
	if token.GetType() == expectedType {
		value = token.GetValue()
		if fwk.IsUndefined(expectedValue) || value == expectedValue {
			// Found the right token.
			return value, token, true
		}
	}

	// This is not the right token.
	v.putBack(token)
	return value, token, false
}

func (v *parser_) putBack(token TokenLike) {
	v.next_.AddValue(token)
}

var syntax = map[string]string{
	"Syntax": `Header+ Rule+ Lexigram+ EOL* EOF  ! Terminated with an end-of-file marker.`,
	"Header": `comment EOL`,
	"Rule":   `comment? uppercase ":" Expression EOL+`,
	"Expression": `,
    "Inlined
    "Multilined`,
	"Inlined":    `Factor+ note?`,
	"Multilined": `Line+`,
	"Line":       `EOL Identifier note?`,
	"Identifier": `,
	"lowercase
	"uppercase`,
	"Factor": `Predicate Cardinality?  ! The default cardinality is one.`,
	"Predicate": `,
	"Identifier
    "intrinsic
    "literal`,
	"Cardinality": `,
    "Constrained
    "quantified`,
	"Constrained": `"{" Minimum Maximum? "}"  ! A range of numbers is inclusive.`,
	"Minimum":     `number`,
	"Maximum":     `".." number?`,
	"Lexigram":    `comment? lowercase ":" Pattern note? EOL+`,
	"Pattern":     `Part+ Alternative*`,
	"Part":        `Element Cardinality?  ! The default cardinality is one.`,
	"Element": `,
	"Group
	"Filter
    "Character
    "lowercase
    "literal`,
	"Alternative": `"|" Part+`,
	"Grouped":     `"(" Pattern ")"`,
	"Filtered":    `negation? "[" Character+ "]"`,
	"Character": `,
    "Bounded
    "intrinsic`,
	"Bounded": `Initial Extent?  ! A range of runes is inclusive.`,
	"Initial": `rune`,
	"Extent":  `".." rune`,
}
