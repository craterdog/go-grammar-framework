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

func (v *parser_) parseAlternative() (
	alternative ast.AlternativeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the bar reserved string.
	var bar string
	bar, token, ok = v.parseToken(ReservedToken, "|")
	if !ok {
		// This is not the alternative.
		return alternative, token, false
	}

	// Attempt to parse the part.
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

	// Found the alternative.
	alternative = ast.Alternative().Make(bar, part)
	return alternative, token, true
}

func (v *parser_) parseBounded() (
	bounded ast.BoundedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the initial glyph.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken, "")
	if !ok {
		// This is not the bounded.
		return bounded, token, false
	}

	// Attempt to parse the optional extent glyph.
	var extent ast.ExtentLike
	extent, token, _ = v.parseExtent()

	// Found the bounded.
	bounded = ast.Bounded().Make(glyph, extent)
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
	var left string
	left, token, ok = v.parseToken(ReservedToken, "{")
	if !ok {
		// This is not the constrained.
		return constrained, token, false
	}

	// Attempt to parse the minimum number.
	var number string
	number, token, ok = v.parseToken(NumberToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("number",
			"Constrained",
			"Limit",
		)
		panic(message)
	}

	// Attempt to parse the optional limit number for the constrained.
	var limit ast.LimitLike
	limit, _, _ = v.parseLimit()

	// Attempt to parse the closing bracket for the constrained.
	var right string
	right, token, ok = v.parseToken(ReservedToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("}",
			"Constrained",
			"Limit",
		)
		panic(message)
	}

	// Found the constrained.
	constrained = ast.Constrained().Make(left, number, limit, right)
	return constrained, token, true
}

func (v *parser_) parseDefinition() (
	definition ast.DefinitionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the inlined definition.
	var inlined ast.InlinedLike
	inlined, token, ok = v.parseInlined()
	if ok {
		// Found the inlined definition.
		definition = ast.Definition().Make(inlined)
		return definition, token, true
	}

	// Attempt to parse the multilined definition.
	var multilined ast.MultilinedLike
	multilined, token, ok = v.parseMultilined()
	if ok {
		// Found the multilined definition.
		definition = ast.Definition().Make(multilined)
		return definition, token, true
	}

	// This is not the definition.
	return definition, token, false
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

	// Attempt to parse the string element.
	var textual ast.TextualLike
	textual, token, ok = v.parseTextual()
	if ok {
		// Found the string element.
		element = ast.Element().Make(textual)
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
	// Attempt to parse the optional comment.
	var comment string
	var commentToken TokenLike
	comment, commentToken, _ = v.parseToken(CommentToken, "")

	// Attempt to parse the lowercase identifier.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken, "")
	if !ok {
		// This is not the expression, put back any comment token that was received.
		if col.IsDefined(comment) {
			v.putBack(commentToken)
		}
		return expression, token, false
	}

	// Attempt to parse the colon reserved string.
	var colon string
	colon, token, ok = v.parseToken(ReservedToken, ":")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(":",
			"Expression",
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
			"Expression",
			"Pattern",
		)
		panic(message)
	}

	// Attempt to parse the optional note.
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

	// Found the expression.
	expression = ast.Expression().Make(
		comment,
		lowercase,
		colon,
		pattern,
		note,
		newlines,
	)
	return expression, token, true
}

func (v *parser_) parseExtent() (
	extent ast.ExtentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the dot-dot reserved string.
	var dotdot string
	dotdot, token, ok = v.parseToken(ReservedToken, "..")
	if !ok {
		// This is not the extent glyph.
		return extent, token, false
	}

	// Attempt to parse the extent for the bounded character.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("glyph",
			"Extent",
		)
		panic(message)
	}

	// Found the extent for the bounded character.
	extent = ast.Extent().Make(dotdot, glyph)
	return extent, token, true
}

func (v *parser_) parseFactor() (
	factor ast.FactorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the predicate factor.
	var predicate ast.PredicateLike
	predicate, token, ok = v.parsePredicate()
	if ok {
		factor = ast.Factor().Make(predicate)
		return factor, token, true
	}

	// Attempt to parse the literal factor.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken, "")
	if ok {
		factor = ast.Factor().Make(literal)
		return factor, token, true
	}

	// This is not the factor.
	return factor, token, false
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
	var left string
	left, token, ok = v.parseToken(ReservedToken, "[")
	if !ok {
		// This is not the filtered element, put back any negation token.
		if col.IsDefined(negation) {
			v.putBack(negationToken)
		}
		return filtered, token, false
	}

	// Attempt to parse one or more characters.
	var character ast.CharacterLike
	character, token, ok = v.parseCharacter()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Character",
			"Filtered",
			"Character",
		)
		panic(message)
	}
	var characters = col.List[ast.CharacterLike]()
	for ok {
		characters.AppendValue(character)
		character, _, ok = v.parseCharacter()
	}

	// Attempt to parse the closing bracket for the filtered element.
	var right string
	right, token, ok = v.parseToken(ReservedToken, "]")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("]",
			"Filtered",
			"Character",
		)
		panic(message)
	}

	// Found the filtered element.
	filtered = ast.Filtered().Make(negation, left, characters, right)
	return filtered, token, true
}

func (v *parser_) parseGrouped() (
	grouped ast.GroupedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the opening bracket for the grouped pattern.
	var left string
	left, token, ok = v.parseToken(ReservedToken, "(")
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

	// Attempt to parse the closing bracket for the grouped pattern.
	var right string
	right, token, ok = v.parseToken(ReservedToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Grouped",
			"Pattern",
		)
		panic(message)
	}

	// Found the grouped.
	grouped = ast.Grouped().Make(left, pattern, right)
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

	// Attempt to parse the newline character.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken, "")
	if !ok {
		// This is not the header, put back the comment token.
		v.putBack(commentToken)
		return header, token, false
	}

	// Found the header.
	header = ast.Header().Make(comment, newline)
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

func (v *parser_) parseInlined() (
	inlined ast.InlinedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more factors.
	var factor ast.FactorLike
	factor, token, ok = v.parseFactor()
	if !ok {
		// This is not the inlined definition.
		return inlined, token, false
	}
	var factors = col.List[ast.FactorLike]()
	for ok {
		factors.AppendValue(factor)
		factor, _, ok = v.parseFactor()
	}

	// Attempt to parse the optional note.
	var note string
	note, token, _ = v.parseToken(NoteToken, "")

	// Found the inlined definition.
	inlined = ast.Inlined().Make(factors, note)
	return inlined, token, true
}

func (v *parser_) parseLine() (
	line ast.LineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the newline character.
	var newline string
	var newlineToken TokenLike
	newline, newlineToken, ok = v.parseToken(NewlineToken, "")
	if !ok {
		// This is not the line.
		return line, newlineToken, false
	}

	// Attempt to parse the identifier.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	if !ok {
		// This is not the line, put back the newline token.
		v.putBack(newlineToken)
		return line, token, false
	}

	// Attempt to parse the optional note.
	var note string
	note, token, _ = v.parseToken(NoteToken, "")

	// Found the line.
	line = ast.Line().Make(newline, identifier, note)
	return line, token, true
}

func (v *parser_) parseLimit() (
	limit ast.LimitLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the dot-dot reserved string.
	var dotdot string
	dotdot, token, ok = v.parseToken(ReservedToken, "..")
	if !ok {
		// This is not the limit number.
		return limit, token, false
	}

	// Attempt to parse the optional limit number.
	var number string
	number, token, _ = v.parseToken(NumberToken, "")

	// Found the constrained limit.
	limit = ast.Limit().Make(dotdot, number)
	return limit, token, true
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
		// This is not the multilined definition.
		return multilined, token, false
	}
	var lines = col.List[ast.LineLike]()
	for ok {
		lines.AppendValue(line)
		line, _, ok = v.parseLine()
	}

	// Found the multilined definition.
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
	// Attempt to parse the part.
	var part ast.PartLike
	part, token, ok = v.parsePart()
	if !ok {
		// This is not the pattern.
		return pattern, token, false
	}

	// Attempt to parse the optional supplement.
	var supplement ast.SupplementLike
	supplement, token, _ = v.parseSupplement()

	// Found the pattern.
	pattern = ast.Pattern().Make(part, supplement)
	return pattern, token, true
}

func (v *parser_) parsePredicate() (
	predicate ast.PredicateLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the identifier.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	if !ok {
		// This is not the predicate.
		return predicate, token, false
	}

	// Attempt to parse the optional cardinality.
	var cardinality ast.CardinalityLike
	cardinality, token, _ = v.parseCardinality()

	// Found the predicate.
	predicate = ast.Predicate().Make(identifier, cardinality)
	return predicate, token, true
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
		if col.IsDefined(comment) {
			v.putBack(commentToken)
		}
		return rule, token, false
	}

	// Attempt to parse the colon reserved string.
	var colon string
	colon, token, ok = v.parseToken(ReservedToken, ":")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(":",
			"Rule",
			"Definition",
		)
		panic(message)
	}

	// Attempt to parse the definition.
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

	// Found the rule.
	rule = ast.Rule().Make(comment, uppercase, colon, definition, newlines)
	return rule, token, true
}

func (v *parser_) parseSelective() (
	selective ast.SelectiveLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more alternatives.
	var alternative ast.AlternativeLike
	alternative, token, ok = v.parseAlternative()
	if !ok {
		// This is not the selective.
		return selective, token, false
	}
	var alternatives = col.List[ast.AlternativeLike]()
	for ok {
		alternatives.AppendValue(alternative)
		alternative, _, ok = v.parseAlternative()
	}

	// Found the selective.
	selective = ast.Selective().Make(alternatives)
	return selective, token, true
}

func (v *parser_) parseSequential() (
	sequential ast.SequentialLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more parts.
	var part ast.PartLike
	part, token, ok = v.parsePart()
	if !ok {
		// This is not the sequential.
		return sequential, token, false
	}
	var parts = col.List[ast.PartLike]()
	for ok {
		parts.AppendValue(part)
		part, _, ok = v.parsePart()
	}

	// Found the sequential.
	sequential = ast.Sequential().Make(parts)
	return sequential, token, true
}

func (v *parser_) parseSupplement() (
	supplement ast.SupplementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the sequential supplement.
	var sequential ast.SequentialLike
	sequential, token, ok = v.parseSequential()
	if ok {
		// Found the sequential supplement.
		supplement = ast.Supplement().Make(sequential)
		return supplement, token, true
	}

	// Attempt to parse the selective supplement.
	var selective ast.SelectiveLike
	selective, token, ok = v.parseSelective()
	if ok {
		// Found the selective supplement.
		supplement = ast.Supplement().Make(selective)
		return supplement, token, true
	}

	// This is not the supplement.
	return supplement, token, false
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

	// Found the syntax.
	syntax = ast.Syntax().Make(headers, rules, expressions)
	return syntax, token, true
}

func (v *parser_) parseTextual() (
	textual ast.TextualLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the intrinsic textual element.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken, "")
	if ok {
		// Found the intrinsic textual element.
		textual = ast.Textual().Make(intrinsic)
		return textual, token, true
	}

	// Attempt to parse the glyph textual element.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken, "")
	if ok {
		// Found the glyph textual element.
		textual = ast.Textual().Make(glyph)
		return textual, token, true
	}

	// Attempt to parse the literal textual element.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken, "")
	if ok {
		// Found the literal textual element.
		textual = ast.Textual().Make(literal)
		return textual, token, true
	}

	// Attempt to parse the lowercase textual element.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken, "")
	if ok {
		// Found the lowercase text.
		textual = ast.Textual().Make(lowercase)
		return textual, token, true
	}

	// This is not the textual element.
	return textual, token, false
}

func (v *parser_) parseToken(expectedType TokenType, expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the specific token.
	token = v.getNextToken()
	if token == nil {
		// We are at the end-of-file marker.
		return value, token, false
	}
	if token.GetType() == expectedType {
		value = token.GetValue()
		if col.IsUndefined(expectedValue) || value == expectedValue {
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
	"Syntax": `Header+ Rule+ Expression+`,
	"Header": `comment newline`,
	"Rule":   `comment? uppercase ":" Definition newline+`,
	"Definition": `
    Inlined
    Multilined`,
	"Inlined":    `Factor+ note?`,
	"Multilined": `Line+`,
	"Line":       `newline Identifier note?`,
	"Identifier": `
    lowercase
    uppercase`,
	"Factor": `Predicate Cardinality?  ! The default cardinality is one.`,
	"Predicate": `
    literal
    lowercase
    uppercase`,
	"Cardinality": `
    Constrained
    quantified`,
	"Constrained": `"{" number Limit? "}"  ! A constrained range of numbers is inclusive.`,
	"Limit":       `".." number?`,
	"Expression":  `comment? lowercase ":" Pattern note? newline+`,
	"Pattern":     `Part Supplement?`,
	"Supplement": `
    Sequential
    Selective`,
	"Sequential":  `Part+`,
	"Selective":   `Alternative+`,
	"Alternative": `"|" Part`,
	"Part":        `Element Cardinality?  ! The default cardinality is one.`,
	"Element": `
    Grouped
    Filtered
    Textual`,
	"Grouped":  `"(" Pattern ")"`,
	"Filtered": `negation? "[" Character+ "]"`,
	"Textual": `
    intrinsic
    glyph
    literal
    lowercase`,
	"Character": `
    Bounded
    intrinsic`,
	"Bounded": `glyph Extent?  ! A bounded range of glyphs is inclusive.`,
	"Extent":  `".." glyph`,
}
