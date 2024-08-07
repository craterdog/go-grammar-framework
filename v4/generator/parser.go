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

package generator

import (
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
	sts "strings"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var parserClass = &parserClass_{
	// Initialize the class constants.
}

// Function

func Parser() ParserClassLike {
	return parserClass
}

// CLASS METHODS

// Target

type parserClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *parserClass_) Make() ParserLike {
	var parser = &parser_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: gra.Processor().Make(),
	}
	parser.visitor_ = gra.Visitor().Make(parser)
	return parser
}

// INSTANCE METHODS

// Target

type parser_ struct {
	// Define the instance attributes.
	class_   ParserClassLike
	visitor_ gra.VisitorLike
	rules_   abs.SetLike[string]

	// Define the inherited aspects.
	gra.Methodical
}

// Attributes

func (v *parser_) GetClass() ParserClassLike {
	return v.class_
}

// Public

func (v *parser_) GenerateParserClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.visitor_.VisitSyntax(syntax)
	implementation = parserTemplate_
	var name = v.extractSyntaxName(syntax)
	implementation = sts.ReplaceAll(implementation, "<module>", module)
	var notice = v.extractNotice(syntax)
	implementation = sts.ReplaceAll(implementation, "<Notice>", notice)
	var uppercase = v.makeUppercase(name)
	implementation = sts.ReplaceAll(implementation, "<Name>", uppercase)
	var lowercase = v.makeLowercase(name)
	implementation = sts.ReplaceAll(implementation, "<name>", lowercase)
	var parseRules = v.extractParseRules()
	implementation = sts.ReplaceAll(implementation, "<ParseRules>", parseRules)
	return implementation
}

// Methodical

func (v *parser_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var name = rule.GetUppercase()
	v.rules_.AddValue(v.makeLowercase(name))
}

func (v *parser_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.rules_ = col.Set[string]()
}

// Private

func (v *parser_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

func (v *parser_) extractParseRules() string {
	var parseRules string
	var iterator = v.rules_.GetIterator()
	for iterator.HasNext() {
		var parseRule = parseTemplate_
		var ruleName = iterator.GetNext()
		parseRule = sts.ReplaceAll(parseRule, "<ruleName>", ruleName)
		ruleName = v.makeUppercase(ruleName)
		parseRule = sts.ReplaceAll(parseRule, "<RuleName>", ruleName)
		parseRules += parseRule
	}
	return parseRules
}

func (v *parser_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *parser_) makeLowercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	name = string(runes)
	if reserved_[name] {
		name += "_"
	}
	return name
}

func (v *parser_) makeUppercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

const parseTemplate_ = `
func (v *parser_) parse<RuleName>() (
	<ruleName> ast.<RuleName>Like,
	token TokenLike,
	ok bool,
) {
	panic("The parse<RuleName>() method has not yet been implemented.")
}
`

const parserTemplate_ = `/*<Notice>*/

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "<module>/ast"
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

func (v *parser_) ParseSource(source string) ast.<Name>Like {
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserClass.queueSize_)
	v.next_ = col.Stack[TokenLike](parserClass.stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse the <name>.
	var <name>, token, ok = v.parse<Name>()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("<Name>",
			"<Name>",
		)
		panic(message)
	}

	// Found the <name>.
	return <name>
}

// Private

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
<ParseRules>
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
	"<Name>": "Component newline*",
}
`
