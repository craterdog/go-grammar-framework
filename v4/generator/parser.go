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
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
	stc "strconv"
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
		class_:    c,
		analyzer_: gra.Analyzer().Make(),
	}
	return parser
}

// INSTANCE METHODS

// Target

type parser_ struct {
	// Define the instance attributes.
	class_    ParserClassLike
	analyzer_ gra.AnalyzerLike
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
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = parserTemplate_
	implementation = replaceAll(implementation, "module", module)
	var notice = v.analyzer_.GetNotice()
	implementation = replaceAll(implementation, "notice", notice)
	var syntaxName = v.analyzer_.GetSyntaxName()
	implementation = replaceAll(implementation, "syntaxName", syntaxName)
	var methods = v.generateMethods()
	implementation = replaceAll(implementation, "methods", methods)
	return implementation
}

// Private

func (v *parser_) generateArguments(
	rule string,
) (
	arguments string,
) {
	var references = v.analyzer_.GetReferences(rule)
	var variableNames = generateVariableNames(references).GetIterator()

	// Define the first argument.
	if variableNames.GetSize() > 1 {
		// Use the multiline argument style.
		arguments += "\n\t\t"
	}
	var argument = variableNames.GetNext()
	arguments += replaceAll(parseArgumentTemplate_, "argument", argument)

	// Define any additional arguments.
	for variableNames.HasNext() {
		arguments += ",\n\t\t"
		argument = variableNames.GetNext()
		arguments += replaceAll(parseArgumentTemplate_, "argument", argument)
	}
	if variableNames.GetSize() > 1 {
		// Use the multiline argument style.
		arguments += ",\n\t"
	}

	return arguments
}

func (v *parser_) generateInlineRule(
	variableName string,
	reference ast.ReferenceLike,
) (
	implementation string,
) {
	implementation = parseRuleTemplate_
	var cardinality = reference.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		implementation = v.generateInlineCardinality(
			cardinality,
			parseOptionalRuleTemplate_,
			parseRepeatedRuleTemplate_,
		)
	}
	implementation = replaceAll(implementation, "variableName", variableName)
	var pluralName = makePlural(variableName)
	implementation = replaceAll(implementation, "pluralName", pluralName)
	var ruleName = reference.GetIdentifier().GetAny().(string)
	implementation = replaceAll(implementation, "ruleName", ruleName)
	return implementation
}

func (v *parser_) generateInlineToken(
	variableName string,
	reference ast.ReferenceLike,
) (
	implementation string,
) {
	implementation = parseTokenTemplate_
	var cardinality = reference.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		implementation = v.generateInlineCardinality(
			cardinality,
			parseOptionalTokenTemplate_,
			parseRepeatedTokenTemplate_,
		)
	}
	implementation = replaceAll(implementation, "variableName", variableName)
	var pluralName = makePlural(variableName)
	implementation = replaceAll(implementation, "pluralName", pluralName)
	var tokenName = reference.GetIdentifier().GetAny().(string)
	implementation = replaceAll(implementation, "tokenName", tokenName)
	return implementation
}

func (v *parser_) generateInlineCardinality(
	cardinality ast.CardinalityLike,
	optionalTemplate string,
	repeatedTemplate string,
) (
	implementation string,
) {
	var first string
	var last = "unlimited"
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		implementation = repeatedTemplate
		switch actual.GetAny().(string) {
		case "?":
			// This is the "{0..1}" case.
			first = "0"
			last = "1"
			implementation = optionalTemplate
		case "*":
			// This is the "{0..}" case.
			first = "0"
		case "+":
			// This is the "{1..}" case.
			first = "1"
		}
	case ast.QuantifiedLike:
		first = actual.GetNumber()
		var limit = actual.GetOptionalLimit()
		if col.IsUndefined(limit) {
			// This is the "{m}" case.
			last = first
		} else {
			last = limit.GetOptionalNumber()
			if col.IsUndefined(last) {
				// This is the "{m..}" case.
				last = "unlimited"
			}
			// This is the "{m..n}" case.
		}
	}
	implementation = replaceAll(implementation, "first", first)
	implementation = replaceAll(implementation, "last", last)
	return implementation
}

func (v *parser_) generateInlineMethod(
	rule string,
) (
	implementation string,
) {
	var terms = v.analyzer_.GetTerms(rule).GetIterator()
	var references = v.analyzer_.GetReferences(rule)
	var variableNames = generateVariableNames(references).GetIterator()
	var handler string
	var caseHandler string
	for terms.HasNext() {
		var term = terms.GetNext()
		switch actual := term.GetAny().(type) {
		case ast.ReferenceLike:
			var variableName = variableNames.GetNext()
			implementation += v.generateInlineReference(variableName, actual)
		case string:
			implementation += v.generateInlineLiteral(actual)
		}
		if col.IsUndefined(handler) {
			handler = parseReturnFalseTemplate_
			caseHandler = parseTooFewTemplate_
		} else {
			handler = parseReturnPanicTemplate_
			caseHandler = parseTooManyTemplate_
		}
		implementation = replaceAll(implementation, "handler", handler)
		implementation = replaceAll(implementation, "caseHandler", caseHandler)

	}
	implementation += parseRuleFoundTemplate_
	implementation = replaceAll(implementation, "rule", rule)
	var arguments = v.generateArguments(rule)
	implementation = replaceAll(implementation, "arguments", arguments)
	implementation = replaceAll(parseRuleMethodTemplate_, "implementation", implementation)
	return implementation
}

func (v *parser_) generateInlineLiteral(
	literal string,
) (
	implementation string,
) {
	var delimiter, err = stc.Unquote(literal) // Remove the double quotes.
	if err != nil {
		panic(err)
	}
	implementation = parseDelimiterTemplate_
	implementation = replaceAll(implementation, "delimiter", delimiter)
	return implementation
}

func (v *parser_) generateMethods() (
	implementation string,
) {
	var rules = v.analyzer_.GetRuleNames().GetIterator()
	for rules.HasNext() {
		var method string
		var rule = rules.GetNext()
		switch {
		case col.IsDefined(v.analyzer_.GetIdentifiers(rule)):
			method = v.generateMultilineMethod(rule)
		case col.IsDefined(v.analyzer_.GetReferences(rule)):
			method = v.generateInlineMethod(rule)
		}
		method = replaceAll(method, "rule", rule)
		implementation += method
	}
	return implementation
}

func (v *parser_) generateMultilineMethod(
	rule string,
) (
	implementation string,
) {
	var tokenCases, ruleCases string
	var identifiers = v.analyzer_.GetIdentifiers(rule).GetIterator()

	for identifiers.HasNext() {
		var identifier = identifiers.GetNext()
		var name = identifier.GetAny().(string)
		switch {
		case gra.Scanner().MatchesType(name, gra.LowercaseToken):
			tokenCases += v.generateMultilineToken(name)
		case gra.Scanner().MatchesType(name, gra.UppercaseToken):
			ruleCases += v.generateMultilineRule(name)
		}
	}
	var defaultCase = parseRuleDefaultCaseTemplate_

	implementation = parseAnyTemplate_
	implementation = replaceAll(implementation, "ruleCases", ruleCases)
	implementation = replaceAll(implementation, "tokenCases", tokenCases)
	implementation = replaceAll(implementation, "defaultCase", defaultCase)
	implementation = replaceAll(implementation, "rule", rule)
	return replaceAll(parseRuleMethodTemplate_, "implementation", implementation)
}

func (v *parser_) generateMultilineRule(
	ruleName string,
) (
	implementation string,
) {
	implementation = parseRuleCaseTemplate_
	if v.analyzer_.IsPlural(ruleName) {
		implementation = parseSingularRuleCaseTemplate_
	}
	implementation = replaceAll(implementation, "ruleName", ruleName)
	return implementation
}

func (v *parser_) generateMultilineToken(
	tokenName string,
) (
	implementation string,
) {
	implementation = parseTokenCaseTemplate_
	if v.analyzer_.IsPlural(tokenName) {
		implementation = parseSingularTokenCaseTemplate_
	}
	implementation = replaceAll(implementation, "tokenName", tokenName)
	return implementation
}

func (v *parser_) generateInlineReference(
	variableName string,
	reference ast.ReferenceLike,
) (
	implementation string,
) {
	var identifier = reference.GetIdentifier().GetAny().(string)
	switch {
	case gra.Scanner().MatchesType(identifier, gra.LowercaseToken):
		implementation = v.generateInlineToken(variableName, reference)
	case gra.Scanner().MatchesType(identifier, gra.UppercaseToken):
		implementation = v.generateInlineRule(variableName, reference)
	}
	return implementation
}

// Templates

const parseAnyTemplate_ = `<RuleCases><TokenCases><DefaultCase>`

const parseArgumentTemplate_ = `<argument_>`

const parseOptionalRuleTemplate_ = `
	// Attempt to parse an optional <ruleName> rule.
	var <variableName_> ast.<RuleName>Like
	<variableName_>, _, _ = v.parse<RuleName>()
`

const parseRepeatedRuleTemplate_ = `
	// Attempt to parse <first> to <last> <ruleName> rules.
	var <variableName> = col.List[ast.<RuleName>Like]()
	for i := 0; i < <last>; i++ {
		<ruleName_>, token, ok = v.parse<RuleName>()
		if !ok {
			switch {
			case i < <first>:<CaseHandler>
			case i > <last>:<CaseHandler>
			default:
				break;
			}
		}
		<variableName_>.AppendValue(<ruleName_>)
	}
`

const parseOptionalTokenTemplate_ = `
	// Attempt to parse an optional <ruleName> rule.
	var <variableName_> string
	<variableName_>, _, _ = v.parseToken(<TokenName>Token)
`

const parseRepeatedTokenTemplate_ = `
	// Attempt to parse <first> to <last> <tokenName> tokens.
	var <variableName_> = col.List[string]()
	for i := 0; i < <last>; i++ {
		<tokenName_>, token, ok = v.parseToken(<TokenName>Token)
		if !ok {
			switch {
			case i < <first>:<CaseHandler>
			case i > <last>:<CaseHandler>
			default:
				break;
			}
		}
		<variableName_>.AppendValue(<tokenName_>)
	}
`

const parseReturnFalseTemplate_ = `
		// This is not a <rule> rule.
		return <rule_>, token, false`

const parseReturnPanicTemplate_ = `
		// Found a syntax error.
		var message = v.formatError(token,"<Rule>")
		panic(message)`

const parseTooFewTemplate_ = `
				var message = v.formatError(token, "")
				message += "Too few <pluralName> tokens found."
				panic(message)`

const parseTooManyTemplate_ = `
				var message = v.formatError(token, "")
				message += "Too many <pluralName> tokens found."
				panic(message)`

const parseRuleFoundTemplate_ = `
	// Found a <rule> rule.
	<rule_> = ast.<Rule>().Make(<arguments>)
	return <rule_>, token, true
`

const parseRuleDefaultCaseTemplate_ = `
	// This is not a <rule> rule.
	return <rule_>, token, false
`

const parseRuleMethodTemplate_ = `
func (v *parser_) parse<Rule>() (
	<rule_> ast.<Rule>Like,
	token TokenLike,
	ok bool,
) {<Implementation>}
`

const parseDelimiterTemplate_ = `
	// Attempt to parse a "<delimiter>" delimiter.
	_, token, ok = v.parseDelimiter("<delimiter>")
	if !ok {<Handler>
	}
`

const parseRuleCaseTemplate_ = `
	// Attempt to parse a <ruleName> rule.
	var <ruleName_> ast.<RuleName>Like
	<ruleName_>, token, ok = v.parse<RuleName>()
	if ok {
		// Found a <ruleName> <rule>.
		<rule_> = ast.<Rule>().Make(<ruleName_>)
		return <rule_>, token, true
	}
`

const parseTokenCaseTemplate_ = `
	// Attempt to parse a <tokenName> token.
	var <tokenName_> string
	<tokenName_>, token, ok = v.parseToken(<TokenName>Token)
	if ok {
		// Found a <tokenName> <rule>.
		<rule_> = ast.<Rule>().Make(<tokenName_>)
		return <rule_>, token, true
	}
`

const parseSingularRuleCaseTemplate_ = `
	// Attempt to parse a <ruleName> rule.
	var <ruleName_> ast.<RuleName>Like
	<ruleName_>, token, ok = v.parse<RuleName>()
	if ok {
		// Found a <ruleName> <rule>.
		<rule_> = ast.<Rule>().Make(<ruleName_>)
		return <rule_>, token, true
	}
`

const parseSingularTokenCaseTemplate_ = `
	// Attempt to parse a <tokenName> token.
	var <tokenName_> string
	<tokenName_>, token, ok = v.parse<TokenName>()
	if ok {
		// Found a <tokenName> <rule>.
		<rule_> = ast.<Rule>().Make(<tokenName_>)
		return <rule_>, token, true
	}
`

const parseRuleTemplate_ = `
	// Attempt to parse a <ruleName> rule.
	var <variableName_> ast.<RuleName>Like
	<variableName_>, token, ok = v.parse<RuleName>()
	if !ok {<Handler>
	}
`

const parseTokenTemplate_ = `
	// Attempt to parse a <tokenName> token.
	var <tokenName_> string
	<tokenName_>, token, ok = v.parseToken(<TokenName>Token)
	if !ok {<Handler>
	}
`

const parserTemplate_ = `<Notice>

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

func (v *parser_) ParseSource(source string) ast.<SyntaxName>Like {
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserClass.queueSize_)
	v.next_ = col.Stack[TokenLike](parserClass.stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse the <syntaxName>.
	var <syntaxName>, token, ok = v.parse<SyntaxName>()
	if !ok {
		var message = v.formatError(token, "<SyntaxName>")
		panic(message)
	}

	// Found the <syntaxName>.
	return <syntaxName>
}

// Private

const unlimited = "4294967295" // Default to a reasonable value.
<Methods>
func (v *parser_) parseDelimiter(expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == expectedValue {
			// Found the right delimiter.
			return value, token, true
		}
		v.putBack(token)
	}

	// This is not the right delimiter.
	return value, token, false
}

func (v *parser_) parseToken(tokenType TokenType) (
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
	if token.GetType() == tokenType {
		// Found the right token type.
		value = token.GetValue()
		return value, token, true
	}

	// This is not the right token type.
	v.putBack(token)
	return value, token, false
}

func (v *parser_) formatError(token TokenLike, ruleName string) string {
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
	if col.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			syntax_[ruleName],
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
		var message = v.formatError(token, "")
		panic(message)
	}

	return token
}

func (v *parser_) putBack(token TokenLike) {
	v.next_.AddValue(token)
}

var syntax_ = map[string]string{
	"<SyntaxName>": "Component newline*",
}
`
