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
)

// CLASS ACCESS

// Reference

var scannerClass = &scannerClass_{
	// Initialize the class constants.
}

// Function

func Scanner() ScannerClassLike {
	return scannerClass
}

// CLASS METHODS

// Target

type scannerClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *scannerClass_) Make() ScannerLike {
	var scanner = &scanner_{
		// Initialize the instance attributes.
		class_:    c,
		analyzer_: gra.Analyzer().Make(),
	}
	return scanner
}

// INSTANCE METHODS

// Target

type scanner_ struct {
	// Define the instance attributes.
	class_    ScannerClassLike
	analyzer_ gra.AnalyzerLike
}

// Attributes

func (v *scanner_) GetClass() ScannerClassLike {
	return v.class_
}

// Public

func (v *scanner_) GenerateScannerClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = v.getTemplate(classTemplate)
	var notice = v.analyzer_.GetNotice()
	implementation = replaceAll(implementation, "notice", notice)
	var tokenNames = v.generateTokenNames()
	implementation = replaceAll(implementation, "tokenNames", tokenNames)
	var tokenMatchers = v.generateTokenMatchers()
	implementation = replaceAll(implementation, "tokenMatchers", tokenMatchers)
	var foundCases = v.generateFoundCases()
	implementation = replaceAll(implementation, "foundCases", foundCases)
	var ignoredCases = v.generateIgnoredCases()
	implementation = replaceAll(implementation, "ignoredCases", ignoredCases)
	var expressions = v.generateExpressions()
	implementation = replaceAll(implementation, "expressions", expressions)
	return implementation
}

// Private

func (v *scanner_) generateExpressions() string {
	var expressions = "// Define the regular expression patterns for each token type."
	var iterator = v.analyzer_.GetExpressions().GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var name = association.GetKey()
		var regexp = association.GetValue()
		expressions += "\n\t" + name + "_ = " + regexp
	}
	return expressions
}

func (v *scanner_) generateFoundCases() string {
	var foundCases = "// Find the next token type."
	var iterator = v.analyzer_.GetTokenNames().GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var tokenType = makeUpperCase(tokenName) + "Token"
		foundCases += "\n\t\tcase v.foundToken(" + tokenType + "):"
	}
	return foundCases
}

func (v *scanner_) generateIgnoredCases() string {
	var ignoreCases = "// Ignore the implicit token types."
	var iterator = v.analyzer_.GetIgnored().GetIterator()
	for iterator.HasNext() {
		var tokenType = iterator.GetNext()
		ignoreCases += "\n\tcase \"" + tokenType + "\":"
		ignoreCases += "\n\t\treturn"
	}
	return ignoreCases
}

func (v *scanner_) generateTokenMatchers() string {
	var tokenMatchers = "// Define pattern matchers for each type of token."
	var iterator = v.analyzer_.GetTokenNames().GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var tokenType = makeUpperCase(tokenName) + "Token"
		tokenMatchers += "\n\t\t" + tokenType +
			`: reg.MustCompile("^" + ` + tokenName + `_),`
	}
	return tokenMatchers
}

func (v *scanner_) generateTokenNames() string {
	var tokenNames = `ErrorToken: "error",`
	var iterator = v.analyzer_.GetTokenNames().GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var tokenType = makeUpperCase(tokenName) + "Token"
		tokenNames += "\n\t\t" + tokenType + `: "` + tokenName + `",`
	}
	return tokenNames
}

func (v *scanner_) getTemplate(name string) string {
	var template = scannerTemplates_.GetValue(name)
	return template
}

// PRIVATE GLOBALS

// Constants

var scannerTemplates_ = col.Catalog[string, string](
	map[string]string{
		classTemplate: `<Notice>

package grammar

import (
	fmt "fmt"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	reg "regexp"
	sts "strings"
)

// CLASS ACCESS

// Reference

var scannerClass = &scannerClass_{
	// Initialize the class constants.
	tokens_: map[TokenType]string{
		<TokenNames>
	},
	matchers_: map[TokenType]*reg.Regexp{
		<TokenMatchers>
	},
}

// Function

func Scanner() ScannerClassLike {
	return scannerClass
}

// CLASS METHODS

// Target

type scannerClass_ struct {
	// Define the class constants.
	tokens_   map[TokenType]string
	matchers_ map[TokenType]*reg.Regexp
}

// Constructors

func (c *scannerClass_) Make(
	source string,
	tokens abs.QueueLike[TokenLike],
) ScannerLike {
	var scanner = &scanner_{
		// Initialize the instance attributes.
		class_:    c,
		line_:     1,
		position_: 1,
		runes_:    []rune(source),
		tokens_:   tokens,
	}
	go scanner.scanTokens() // Start scanning tokens in the background.
	return scanner
}

// Functions

func (c *scannerClass_) FormatToken(token TokenLike) string {
	var value = token.GetValue()
	var s = fmt.Sprintf("%q", value)
	if len(s) > 40 {
		s = fmt.Sprintf("%.40q...", value)
	}
	return fmt.Sprintf(
		"Token [type: %s, line: %d, position: %d]: %s",
		c.tokens_[token.GetType()],
		token.GetLine(),
		token.GetPosition(),
		s,
	)
}

func (c *scannerClass_) FormatType(tokenType TokenType) string {
	return c.tokens_[tokenType]
}

func (c *scannerClass_) MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var matcher = c.matchers_[tokenType]
	var match = matcher.FindString(tokenValue)
	return len(match) > 0
}

// INSTANCE METHODS

// Target

type scanner_ struct {
	// Define the instance attributes.
	class_    ScannerClassLike
	first_    uint // A zero based index of the first possible rune in the next token.
	next_     uint // A zero based index of the next possible rune in the next token.
	line_     uint // The line number in the source string of the next rune.
	position_ uint // The position in the current line of the next rune.
	runes_    []rune
	tokens_   abs.QueueLike[TokenLike]
}

// Attributes

func (v *scanner_) GetClass() ScannerClassLike {
	return v.class_
}

// Private

/*
NOTE:
These private constants define the regular expression sub-patterns that make up
the intrinsic types and token types.  Unfortunately there is no way to make them
private to the scanner class since they must be TRUE Go constants to be used in
this way.  We append an underscore to each name to lessen the chance of a name
collision with other private Go class constants in this package.
*/
const (
	// Define the regular expression patterns for each intrinsic type.
	any_     = "." // This does NOT include newline characters.
	control_ = "\\p{Cc}"
	digit_   = "\\p{Nd}"
	eol_     = "\\r?\\n"
	lower_   = "\\p{Ll}"
	upper_   = "\\p{Lu}"

	<Expressions>
)

func (v *scanner_) emitToken(tokenType TokenType) {
	switch v.GetClass().FormatType(tokenType) {
	<IgnoredCases>
	}
	var value = string(v.runes_[v.first_:v.next_])
	switch value {
	case "\x00":
		value = "<NULL>"
	case "\a":
		value = "<BELL>"
	case "\b":
		value = "<BKSP>"
	case "\t":
		value = "<HTAB>"
	case "\f":
		value = "<FMFD>"
	case "\n":
		value = "<EOLN>"
	case "\r":
		value = "<CRTN>"
	case "\v":
		value = "<VTAB>"
	}
	var token = Token().Make(v.line_, v.position_, tokenType, value)
	//fmt.Println(Scanner().FormatToken(token)) // Uncomment when debugging.
	v.tokens_.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundError() {
	v.next_++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(tokenType TokenType) bool {
	var text = string(v.runes_[v.next_:])
	var matcher = scannerClass.matchers_[tokenType]
	var match = matcher.FindString(text)
	if len(match) > 0 {
		var token = []rune(match)
		var length = uint(len(token))

		// Found the requested token type.
		v.next_ += length
		v.emitToken(tokenType)
		var count = uint(sts.Count(match, "\n"))
		if count > 0 {
			v.line_ += count
			v.position_ = v.indexOfLastEol(token)
		} else {
			v.position_ += v.next_ - v.first_
		}
		v.first_ = v.next_
		return true
	}

	// The next token is not the requested token type.
	return false
}

func (v *scanner_) indexOfLastEol(runes []rune) uint {
	var length = uint(len(runes))
	for index := length; index > 0; index-- {
		if runes[index-1] == '\n' {
			return length - index + 1
		}
	}
	return 0
}

func (v *scanner_) scanTokens() {
loop:
	for v.next_ < uint(len(v.runes_)) {
		switch {
		<FoundCases>
		default:
			v.foundError()
			break loop
		}
	}
	v.tokens_.CloseQueue()
}
`,
	},
)
