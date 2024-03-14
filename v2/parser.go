/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
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
		next_: col.Stack[TokenLike]().MakeWithCapacity(c.stackSize_),
	}
}

// INSTANCE METHODS

// Target

type parser_ struct {
	next_   col.StackLike[TokenLike] // A stack of unprocessed retrieved tokens.
	source_ string                   // The original source code.
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
}

// Public

func (v *parser_) ParseSource(source string) GrammarLike {
	// The scanner runs in a separate Go routine.
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike]().MakeWithCapacity(parserClass.queueSize_)
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse a grammar.
	var grammar, token, ok = v.parseGrammar()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("grammar",
			"$source",
			"$grammar",
		)
		panic(message)
	}

	// Attempt to parse the end-of-file marker.
	_, token, ok = v.parseToken(EOFToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("EOF",
			"$source",
			"$grammar",
		)
		panic(message)
	}

	// Found a grammar.
	return grammar
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
		token,
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
func (v *parser_) generateGrammar(expected string, symbols ...string) string {
	var message = "Was expecting '" + expected + "' from:\n"
	for _, symbol := range symbols {
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			symbol,
			grammar[symbol],
		)
	}
	return message
}

/*
This private instance method attempts to read the next token from the token
stream and return it.
*/
func (v *parser_) getNextToken() TokenLike {
	// Check for any unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveTop()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveHead() // This will block if the queue is empty.
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
	var factor FactorLike
	var factors = col.List[FactorLike]().Make()
	var note string

	// Attempt to parse the first factor.
	factor, token, ok = v.parseFactor()
	if !ok {
		// This is not an alternative.
		return alternative, token, false
	}

	// Parse any additional factors.
	for ok {
		factors.AppendValue(factor)
		factor, _, ok = v.parseFactor()
	}

	// Attempt to parse an optional note.
	note, token, _ = v.parseToken(NoteToken, "")

	// Found an alternative.
	alternative = Alternative().MakeWithAttributes(factors, note)
	return alternative, token, true
}

func (v *parser_) parseAssertion() (
	assertion AssertionLike,
	token TokenLike,
	ok bool,
) {
	var element ElementLike
	var glyph GlyphLike
	var precedence PrecedenceLike

	// Attempt to parse an element assertion.
	element, token, ok = v.parseElement()
	if ok {
		assertion = Assertion().MakeWithElement(element)
		return assertion, token, true
	}

	// Attempt to parse a glyph assertion.
	glyph, token, ok = v.parseGlyph()
	if ok {
		assertion = Assertion().MakeWithGlyph(glyph)
		return assertion, token, true
	}

	// Attempt to parse a precedence assertion.
	precedence, token, ok = v.parsePrecedence()
	if ok {
		assertion = Assertion().MakeWithPrecedence(precedence)
		return assertion, token, true
	}

	// This is not an assertion.
	return assertion, token, false
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
		cardinality = Cardinality().MakeWithAttributes(constraint)
		return cardinality, token, true
	}

	// Attempt to parse a zero-or-more cardinality.
	_, token, ok = v.parseToken(DelimiterToken, "*")
	if ok {
		constraint = Constraint().MakeWithAttributes("0", "")
		cardinality = Cardinality().MakeWithAttributes(constraint)
		return cardinality, token, true
	}

	// Attempt to parse a one-or-more cardinality.
	_, token, ok = v.parseToken(DelimiterToken, "+")
	if ok {
		constraint = Constraint().MakeWithAttributes("1", "")
		cardinality = Cardinality().MakeWithAttributes(constraint)
		return cardinality, token, true
	}

	// Attempt to parse an explicit constrained cardinality.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if ok {
		constraint, token, ok = v.parseConstraint()
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("constraint",
				"$cardinality",
				"$constraint",
			)
			panic(message)
		}
		_, token, ok = v.parseToken(DelimiterToken, "}")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("}",
				"$cardinality",
				"$constraint",
			)
			panic(message)
		}

		// Found a cardinality.
		cardinality = Cardinality().MakeWithAttributes(constraint)
		return cardinality, token, true
	}
	// This is not a cardinality.
	return cardinality, token, false
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
		// Attempt to parse the last number in the range of numbers.
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
	var expression ExpressionLike
	var symbol string

	// Attempt to parse a symbol.
	symbol, token, ok = v.parseToken(SymbolToken, "")
	if !ok {
		// This is not a definition.
		return definition, token, false
	}

	// Attempt to parse a separator delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ":")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(":",
			"$definition",
			"$expression",
		)
		panic(message)
	}

	// Attempt to parse an expression.
	expression, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("expression",
			"$definition",
			"$expression",
		)
		panic(message)
	}

	// Found a definition.
	definition = Definition().MakeWithAttributes(symbol, expression)
	return definition, token, true
}

func (v *parser_) parseElement() (
	element ElementLike,
	token TokenLike,
	ok bool,
) {
	var intrinsic string
	var literal string
	var name string

	// Attempt to parse an intrinsic element.
	intrinsic, token, ok = v.parseToken(IntrinsicToken, "")
	if ok {
		element = Element().MakeWithIntrinsic(intrinsic)
		return element, token, true
	}

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
	var alternative AlternativeLike
	var alternatives = col.List[AlternativeLike]().Make()

	// Attempt to parse a multi-line expression.
	_, _, ok = v.parseToken(EOLToken, "")
	if ok {
		// Attempt to parse the first alternative.
		alternative, token, ok = v.parseAlternative()
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("alternative",
				"$expression",
				"$alternative",
			)
			panic(message)
		}

		// Parse any additional alternatives.
		for ok {
			alternatives.AppendValue(alternative)
			_, token, ok = v.parseToken(EOLToken, "")
			if !ok {
				break
			}
			alternative, _, ok = v.parseAlternative()
			if !ok {
				v.putBack(token)
				break
			}
		}

		// Attempt to parse an optional end-of-line character.
		_, token, _ = v.parseToken(EOLToken, "")

		// Found a multi-line expression.
		expression = Expression().MakeWithAttributes(alternatives, true)
		return expression, token, true
	}

	// Attempt to parse the first alternative in an in-line expression.
	alternative, token, ok = v.parseAlternative()
	if !ok {
		return expression, token, false
	}

	// Parse any additional alternatives.
	for ok {
		alternatives.AppendValue(alternative)
		_, token, ok = v.parseToken(DelimiterToken, "|")
		if ok {
			// Attempt to parse an alternative.
			alternative, token, ok = v.parseAlternative()
			if !ok {
				var message = v.formatError(token)
				message += v.generateGrammar("alternative",
					"$expression",
					"$alternative",
				)
				panic(message)
			}
		}
	}

	// Found an in-line expression.
	expression = Expression().MakeWithAttributes(alternatives, false)
	return expression, token, true
}

func (v *parser_) parseFactor() (
	factor FactorLike,
	token TokenLike,
	ok bool,
) {
	var cardinality CardinalityLike
	var predicate PredicateLike

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
			message += v.generateGrammar("CHARACTER",
				"$glyph",
			)
			panic(message)
		}
	}

	// Found a glyph.
	glyph = Glyph().MakeWithAttributes(first, last)
	return glyph, token, true
}

func (v *parser_) parseGrammar() (
	grammar GrammarLike,
	token TokenLike,
	ok bool,
) {
	var comment string
	var statement StatementLike
	var statements = col.List[StatementLike]().Make()

	// Attempt to parse a comment.
	comment, token, ok = v.parseToken(CommentToken, "")
	if !ok {
		return grammar, token, false
	}

	// Attempt to parse a statement.
	statement, token, ok = v.parseStatement()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("statement",
			"$grammar",
			"$copyright",
			"$statement",
		)
		panic(message)
	}

	// Parse any additional statements.
	for ok {
		statements.AppendValue(statement)
		statement, token, ok = v.parseStatement()
	}

	// Found a grammar.
	grammar = Grammar().MakeWithAttributes(comment, statements)
	return grammar, token, true
}

func (v *parser_) parsePrecedence() (
	precedence PrecedenceLike,
	token TokenLike,
	ok bool,
) {
	var expression ExpressionLike

	// Attempt to parse the opening delimiter for a precedence.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		// This is not a precedence.
		return precedence, token, false
	}

	// Attempt to parse an expression.
	expression, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("expression",
			"$precedence",
			"$expression",
		)
		panic(message)
	}

	// Attempt to parse the closing delimiter for the precedence.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"$precedence",
			"$expression",
		)
		panic(message)
	}

	// Found a precedence.
	precedence = Precedence().MakeWithAttributes(expression)
	return precedence, token, true
}

func (v *parser_) parsePredicate() (
	predicate PredicateLike,
	token TokenLike,
	ok bool,
) {
	var assertion AssertionLike
	var isInverted bool

	// Check to see if the assertion is inverted.
	_, _, isInverted = v.parseToken(DelimiterToken, "~")

	// Attempt to parse an assertion.
	assertion, token, ok = v.parseAssertion()
	if !ok {
		if isInverted {
			var message = v.formatError(token)
			message += v.generateGrammar("assertion",
				"$predicate",
				"$assertion",
			)
			panic(message)
		} else {
			// This is not a predicate.
			return predicate, token, false
		}
	}

	// Found a predicate.
	predicate = Predicate().MakeWithAttributes(assertion, isInverted)
	return predicate, token, true
}

func (v *parser_) parseStatement() (
	statement StatementLike,
	token TokenLike,
	ok bool,
) {
	var comment string
	var definition DefinitionLike

	// Attempt to parse an optional comment.
	comment, _, _ = v.parseToken(CommentToken, "")

	// Attempt to parse a definition.
	definition, token, ok = v.parseDefinition()
	if !ok {
		if len(comment) > 0 {
			var message = v.formatError(token)
			message += v.generateGrammar("definition",
				"$statement",
				"$definition",
			)
			panic(message)
		} else {
			return statement, token, false
		}
	}

	// Attempt to parse an end-of-line character.
	_, token, ok = v.parseToken(EOLToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("EOL",
			"$statement",
			"$definition",
		)
		panic(message)
	}

	// Parse any additional end-of-line characters.
	for ok {
		_, _, ok = v.parseToken(EOLToken, "")
	}

	// Found a statement.
	statement = Statement().MakeWithAttributes(comment, definition)
	return statement, token, true
}

func (v *parser_) parseToken(expectedType TokenType, expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token.
	token = v.getNextToken()
	value = token.GetValue()
	if token.GetType() == expectedType {
		var constrained = len(expectedValue) > 0
		if !constrained || value == expectedValue {
			// Found the expected token.
			return value, token, true
		}
	}

	// This is not the right token.
	v.putBack(token)
	return "", token, false
}

func (v *parser_) putBack(token TokenLike) {
	v.next_.AddValue(token)
}

var grammar = map[string]string{
	"$source":     `grammar EOF  ! Terminated with an end-of-file marker.`,
	"$grammar":    `COMMENT statement+`,
	"$statement":  `COMMENT? definition EOL+`,
	"$definition": `SYMBOL ":" expression  ! This works for tokens and rules.`,
	"$expression": `
    alternative ("|" alternative)*
    (EOL alternative)+`,
	"$alternative": `factor+ NOTE?`,
	"$factor":      `predicate cardinality?  ! The default cardinality is one.`,
	"$predicate":   `"~"? assertion`,
	"$assertion":   `element | glyph | precedence`,
	"$element":     `INTRINSIC | LITERAL | NAME`,
	"$glyph":       `CHARACTER (".." CHARACTER)?  ! The range of characters is inclusive.`,
	"$precedence":  `"(" expression ")"`,
	"$cardinality": `
    "?"  ! Zero or one instance of a predicate.
    "*"  ! Zero or more instances of a predicate.
    "+"  ! One or more instances of a predicate.
    "{" constraint "}"  ! Constrains the number of instances of a predicate.`,
	"$constraint": `NUMBER (".." NUMBER?)?  ! The range of numbers is inclusive.`,
}
