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

package cdsn

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v3/collection"
	sts "strings"
)

// CLASS ACCESS

// Reference

var parserClass = &parserClass_{
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
	queueSize_ int
	stackSize_ int
}

// Constructors

func (c *parserClass_) Make() ParserLike {
	return &parser_{
		tokens_: col.Queue[TokenLike]().MakeWithCapacity(c.queueSize_),
		next_:   col.Stack[TokenLike]().MakeWithCapacity(c.stackSize_),
	}
}

// INSTANCE METHODS

// Target

type parser_ struct {
	source_ string                   // The original source code.
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Public

func (v *parser_) ParseSource(source string) SyntaxLike {
	// The scanner runs in a separate Go routine.
	v.source_ = source
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse a syntax.
	var syntax, token, ok = v.parseSyntax()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Syntax",
			"Cdsn",
			"Syntax",
		)
		panic(message)
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
			"Cdsn",
			"Syntax",
		)
		panic(message)
	}

	// Found a syntax.
	return syntax
}

// Private

/*
This private instance method returns an error message containing the context for
a parsing error.
*/
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

/*
This private instance method is useful when creating scanner and parser error
messages that include the required grammatical rules.
*/
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

/*
This private instance method attempts to read the next token from the token
stream and return it.
*/
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
	alternative AlternativeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more factors.
	var factor FactorLike
	factor, token, ok = v.parseFactor()
	if !ok {
		// This is not an alternative.
		return alternative, token, false
	}
	var factors = col.List[FactorLike]().Make()
	for ok {
		factors.AppendValue(factor)
		factor, token, ok = v.parseFactor()
	}

	// Found an alternative.
	alternative = Alternative().MakeWithFactors(factors)
	return alternative, token, true
}

func (v *parser_) parseAtom() (
	atom AtomLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a glyph.
	var glyph GlyphLike
	glyph, token, ok = v.parseGlyph()
	if ok {
		// Found a glyph atom.
		atom = Atom().MakeWithGlyph(glyph)
		return atom, token, true
	}

	// Attempt to parse an intrinsic.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken, "")
	if ok {
		// Found an intrinsic atom.
		atom = Atom().MakeWithIntrinsic(intrinsic)
		return atom, token, true
	}

	// This is not a atom.
	return atom, token, false
}

func (v *parser_) parseCardinality() (
	cardinality CardinalityLike,
	token TokenLike,
	ok bool,
) {
	var constraint ConstraintLike

	// Attempt to parse a zero-or-one cardinality.
	_, token, ok = v.parseToken(DelimiterToken, "?")
	if ok {
		constraint = Constraint().MakeWithAttributes("0", "1")
		cardinality = Cardinality().MakeWithConstraint(constraint)
		return cardinality, token, true
	}

	// Attempt to parse a zero-or-more cardinality.
	_, token, ok = v.parseToken(DelimiterToken, "*")
	if ok {
		constraint = Constraint().MakeWithAttributes("0", "")
		cardinality = Cardinality().MakeWithConstraint(constraint)
		return cardinality, token, true
	}

	// Attempt to parse a one-or-more cardinality.
	_, token, ok = v.parseToken(DelimiterToken, "+")
	if ok {
		constraint = Constraint().MakeWithAttributes("1", "")
		cardinality = Cardinality().MakeWithConstraint(constraint)
		return cardinality, token, true
	}

	// Attempt to parse an explicit constrained cardinality.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		// This is not a cardinality.
		return cardinality, token, false
	}

	constraint, token, ok = v.parseConstraint()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Constraint",
			"Cardinality",
			"Constraint",
		)
		panic(message)
	}
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("}",
			"Cardinality",
			"Constraint",
		)
		panic(message)
	}

	// Found a cardinality.
	cardinality = Cardinality().MakeWithConstraint(constraint)
	return cardinality, token, true
}

func (v *parser_) parseConstraint() (
	constraint ConstraintLike,
	token TokenLike,
	ok bool,
) {
	var first, last string

	// Attempt to parse the first number in a constraint.
	first, token, ok = v.parseToken(NumberToken, "")
	if !ok {
		// This is not a constraint.
		return constraint, token, false
	}

	// Attempt to parse an additional range of numbers in the constraint.
	_, _, ok = v.parseToken(DelimiterToken, "..")
	if ok {
		// Attempt to parse the optional last number in the range of numbers.
		last, token, _ = v.parseToken(NumberToken, "")
	} else {
		// This constraint is not a range of numbers.
		last = first
	}

	// Found a constraint.
	constraint = Constraint().MakeWithAttributes(first, last)
	return constraint, token, true
}

func (v *parser_) parseDefinition() (
	definition DefinitionLike,
	token TokenLike,
	ok bool,
) {
	var comment string
	var name string
	var expression ExpressionLike

	// Attempt to parse an optional comment.
	comment, _, _ = v.parseToken(CommentToken, "")

	// Attempt to parse a name.
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not a definition.
		return definition, token, false
	}

	// Attempt to parse a separator delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ":")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(":",
			"Definition",
			"Expression",
		)
		panic(message)
	}

	// Attempt to parse an expression.
	expression, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Expression",
			"Definition",
			"Expression",
		)
		panic(message)
	}

	// Attempt to parse one or more end-of-line characters.
	_, token, ok = v.parseToken(EOLToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("EOL",
			"Definition",
			"Expression",
		)
		panic(message)
	}
	for ok {
		_, token, ok = v.parseToken(EOLToken, "")
	}

	// Found a definition.
	definition = Definition().MakeWithAttributes(comment, name, expression)
	return definition, token, true
}

func (v *parser_) parseElement() (
	element ElementLike,
	token TokenLike,
	ok bool,
) {
	var literal string
	var name string

	// Attempt to parse a literal element.
	literal, token, ok = v.parseToken(LiteralToken, "")
	if ok {
		element = Element().MakeWithLiteral(literal)
		return element, token, true
	}

	// Attempt to parse a name element.
	name, token, ok = v.parseToken(NameToken, "")
	if ok {
		element = Element().MakeWithName(name)
		return element, token, true
	}

	// This is not an element.
	return element, token, false
}

func (v *parser_) parseExpression() (
	expression ExpressionLike,
	token TokenLike,
	ok bool,
) {
	var inline InlineLike
	var multiline MultilineLike

	// Attempt to parse an in-line expression.
	inline, token, ok = v.parseInline()
	if ok {
		// Found an in-line expression.
		expression = Expression().MakeWithInline(inline)
		return expression, token, true
	}

	// Attempt to parse a multi-line expression.
	multiline, token, ok = v.parseMultiline()
	if ok {
		// Found a multi-line expression.
		expression = Expression().MakeWithMultiline(multiline)
		return expression, token, true
	}

	// This is not an expression.
	return expression, token, false
}

func (v *parser_) parseFactor() (
	factor FactorLike,
	token TokenLike,
	ok bool,
) {
	var predicate PredicateLike
	var cardinality CardinalityLike

	// Attempt to parse a predicate.
	predicate, token, ok = v.parsePredicate()
	if !ok {
		// This is not a factor.
		return factor, token, false
	}

	// Attempt to parse an optional cardinality.
	cardinality, token, _ = v.parseCardinality()

	// Found a factor.
	factor = Factor().MakeWithAttributes(predicate, cardinality)
	return factor, token, true
}

func (v *parser_) parseFilter() (
	filter FilterLike,
	token TokenLike,
	ok bool,
) {
	// Check to see if the filter is inverted.
	var _, _, inverted = v.parseToken(DelimiterToken, "~")

	// Attempt to parse the opening delimiter for a filter.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	if !ok {
		if !inverted {
			// This is not a filter.
			return filter, token, false
		} else {
			var message = v.formatError(token)
			message += v.generateSyntax("[",
				"Filter",
				"Atom",
			)
			panic(message)
		}
	}

	// Attempt to parse one or more atoms.
	var atom AtomLike
	atom, token, ok = v.parseAtom()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Atom",
			"Filter",
			"Atom",
		)
		panic(message)
	}
	var atoms = col.List[AtomLike]().Make()
	for ok {
		atoms.AppendValue(atom)
		atom, _, ok = v.parseAtom()
	}

	// Attempt to parse the closing delimiter for a filter.
	_, token, ok = v.parseToken(DelimiterToken, "]")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("]",
			"Filter",
			"Atom",
		)
		panic(message)
	}

	// Found a filter.
	filter = Filter().MakeWithAttributes(inverted, atoms)
	return filter, token, true
}

func (v *parser_) parseGlyph() (
	glyph GlyphLike,
	token TokenLike,
	ok bool,
) {
	var first, last string

	// Attempt to parse the first character in a glyph.
	first, token, ok = v.parseToken(CharacterToken, "")
	if !ok {
		// This is not a glyph.
		return glyph, token, false
	}

	// Attempt to parse an additional range of characters in the glyph.
	_, _, ok = v.parseToken(DelimiterToken, "..")
	if ok {
		// Attempt to parse the last character in the range of characters.
		last, token, ok = v.parseToken(CharacterToken, "")
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("character",
				"Glyph",
			)
			panic(message)
		}
	}

	// Found a glyph.
	glyph = Glyph().MakeWithAttributes(first, last)
	return glyph, token, true
}

func (v *parser_) parseSyntax() (
	syntax SyntaxLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more headers.
	var header HeaderLike
	header, token, ok = v.parseHeader()
	if !ok {
		return syntax, token, false
	}
	var headers = col.List[HeaderLike]().Make()
	for ok {
		headers.AppendValue(header)
		header, _, ok = v.parseHeader()
	}

	// Attempt to parse one or more definitions.
	var definition DefinitionLike
	definition, token, ok = v.parseDefinition()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Definition",
			"Syntax",
			"Header",
			"Definition",
		)
		panic(message)
	}
	var definitions = col.List[DefinitionLike]().Make()
	for ok {
		definitions.AppendValue(definition)
		definition, token, ok = v.parseDefinition()
	}

	// Found a syntax.
	syntax = Syntax().MakeWithAttributes(headers, definitions)
	return syntax, token, true
}

func (v *parser_) parseHeader() (
	header HeaderLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a comment.
	var comment string
	var commentToken TokenLike
	comment, commentToken, ok = v.parseToken(CommentToken, "")
	if !ok {
		return header, commentToken, false
	}

	// Attempt to parse one or more end-of-line characters.
	_, token, ok = v.parseToken(EOLToken, "")
	if !ok {
		// This is not a header, put back the comment token.
		v.putBack(commentToken)
		return header, token, false
	}
	for ok {
		_, token, ok = v.parseToken(EOLToken, "")
	}

	// Found a header.
	header = Header().MakeWithComment(comment)
	return header, token, true
}

func (v *parser_) parseInline() (
	inline InlineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more alternatives in an in-line expression.
	var alternative AlternativeLike
	alternative, token, ok = v.parseAlternative()
	if !ok {
		return inline, token, false
	}
	var alternatives = col.List[AlternativeLike]().Make()
	for ok {
		alternatives.AppendValue(alternative)
		_, _, ok = v.parseToken(DelimiterToken, "|")
		if ok {
			// Attempt to parse an alternative.
			alternative, token, ok = v.parseAlternative()
			if !ok {
				var message = v.formatError(token)
				message += v.generateSyntax("Alternative",
					"Inline",
					"Alternative",
				)
				panic(message)
			}
		}
	}

	// Attempt to parse an optional note.
	var note string
	note, token, _ = v.parseToken(NoteToken, "")

	// Found an in-line expression.
	inline = Inline().MakeWithAttributes(alternatives, note)
	return inline, token, true
}

func (v *parser_) parseLine() (
	line LineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an end-of-line character.
	var eolToken TokenLike
	_, eolToken, ok = v.parseToken(EOLToken, "")
	if !ok {
		// This is not a line.
		return line, eolToken, false
	}

	// Attempt to parse the an alternative.
	var alternative AlternativeLike
	alternative, token, ok = v.parseAlternative()
	if !ok {
		// This is not a line, so put back the end-of-line token.
		v.putBack(eolToken)
		return line, token, false
	}

	// Attempt to parse an optional note.
	var note string
	note, token, _ = v.parseToken(NoteToken, "")

	// Found a line.
	line = Line().MakeWithAttributes(alternative, note)
	return line, token, true
}

func (v *parser_) parseMultiline() (
	multiline MultilineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse one or more lines of a multi-line expression.
	var line LineLike
	line, token, ok = v.parseLine()
	if !ok {
		// This is not a multi-line expression.
		return multiline, token, false
	}
	var lines = col.List[LineLike]().Make()
	for ok {
		lines.AppendValue(line)
		line, token, ok = v.parseLine()
	}

	// Found a multi-line expression.
	multiline = Multiline().MakeWithLines(lines)
	return multiline, token, true
}

func (v *parser_) parsePrecedence() (
	precedence PrecedenceLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the opening delimiter for a precedence.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		// This is not a precedence.
		return precedence, token, false
	}

	// Attempt to parse an expression.
	var expression ExpressionLike
	expression, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Expression",
			"Precedence",
			"Expression",
		)
		panic(message)
	}

	// Attempt to parse an optional end-of-line character.
	_, _, _ = v.parseToken(EOLToken, "")

	// Attempt to parse the closing delimiter for the precedence.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Precedence",
			"Expression",
		)
		panic(message)
	}

	// Found a precedence.
	precedence = Precedence().MakeWithExpression(expression)
	return precedence, token, true
}

func (v *parser_) parsePredicate() (
	predicate PredicateLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an atom predicate.
	var atom AtomLike
	atom, token, ok = v.parseAtom()
	if ok {
		predicate = Predicate().MakeWithAtom(atom)
		return predicate, token, true
	}

	// Attempt to parse an element predicate.
	var element ElementLike
	element, token, ok = v.parseElement()
	if ok {
		predicate = Predicate().MakeWithElement(element)
		return predicate, token, true
	}

	// Attempt to parse an filter predicate.
	var filter FilterLike
	filter, token, ok = v.parseFilter()
	if ok {
		predicate = Predicate().MakeWithFilter(filter)
		return predicate, token, true
	}

	// Attempt to parse a precedence predicate.
	var precedence PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		predicate = Predicate().MakeWithPrecedence(precedence)
		return predicate, token, true
	}

	// This is not an predicate.
	return predicate, token, false
}

func (v *parser_) parseToken(expectedType TokenType, expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token.
	token = v.getNextToken()
	if token.GetType() == expectedType {
		value = token.GetValue()
		if len(expectedValue) == 0 || expectedValue == value {
			// Found the expected token.
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
	"Cdsn":        `Syntax EOL* EOF  ! Terminated with an end-of-file marker.`,
	"Syntax":      `Header+ Definition+`,
	"Header":      `comment EOL+`,
	"Definition":  `comment? name ":" Expression EOL+`,
	"Expression":  `Inline | Multiline`,
	"Inline":      `Alternative ("|" Alternative)* note?`,
	"Multiline":   `Line+`,
	"Line":        `EOL Alternative note?`,
	"Alternative": `Factor+`,
	"Factor":      `Predicate Cardinality?  ! The default cardinality is one.`,
	"Predicate":   `Element | Filter | Precedence`,
	"Element":     `literal | name`,
	"Filter":      `"~"? Atom`,
	"Atom":        `Glyph | intrinsic`,
	"Glyph":       `character (".." character)?  ! The range of characters is inclusive.`,
	"Precedence":  `"(" Expression EOL? ")"`,
	"Cardinality": `
    "?"  ! Zero or one instance of a predicate.
    "*"  ! Zero or more instances of a predicate.
    "+"  ! One or more instances of a predicate.
    "{" Constraint "}"  ! Constrains the number of instances of a predicate.`,
	"Constraint": `number (".." number?)?  ! The range of numbers is inclusive.`,
}
