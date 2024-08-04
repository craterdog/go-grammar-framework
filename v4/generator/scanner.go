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
	var processor = gra.Processor().Make()
	var scanner = &scanner_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: processor,
	}
	scanner.visitor_ = gra.Visitor().Make(scanner)
	return scanner
}

// INSTANCE METHODS

// Target

type scanner_ struct {
	// Define the instance attributes.
	class_    ScannerClassLike
	visitor_  gra.VisitorLike
	pattern_  bool
	greedy_   bool
	ignored_  abs.SetLike[string]
	tokens_   abs.SetLike[string]
	reserved_ abs.SetLike[string]
	regexp_   string
	regexps_  abs.CatalogLike[string, string]

	// Define the inherited aspects.
	gra.Methodical
}

// Attributes

func (v *scanner_) GetClass() ScannerClassLike {
	return v.class_
}

// Methodical

func (v *scanner_) ProcessGlyph(glyph string) {
	var character = glyph[1:2] //Remove the single quotes.
	character = v.escapeText(character)
	v.regexp_ += character
}

func (v *scanner_) ProcessIntrinsic(intrinsic string) {
	intrinsic = sts.ToLower(intrinsic)
	if intrinsic == "any" {
		v.greedy_ = false // Turn off "greedy" for expressions containing ANY.
	}
	v.regexp_ += `" + ` + intrinsic + `_ + "`
}

func (v *scanner_) ProcessLiteral(literal string) {
	literal = literal[1 : len(literal)-1] // Remove the double quotes.
	literal = v.escapeText(literal)
	if !v.pattern_ {
		v.reserved_.AddValue(literal)
	}
	v.regexp_ += literal
}

func (v *scanner_) ProcessLowercase(lowercase string) {
	if v.pattern_ {
		v.regexp_ += `(?:" + ` + lowercase + `_ + ")`
	} else {
		v.tokens_.AddValue(lowercase)
	}
}

func (v *scanner_) ProcessNegation(negation string) {
	v.regexp_ += "^"
}

func (v *scanner_) ProcessQuantified(quantified string) {
	v.regexp_ += quantified
}

func (v *scanner_) PreprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
) {
	v.regexp_ += "|"
}

func (v *scanner_) PreprocessConstrained(constrained ast.ConstrainedLike) {
	v.regexp_ += "{"
}

func (v *scanner_) PostprocessConstrained(constrained ast.ConstrainedLike) {
	v.regexp_ += "}"
	if !v.greedy_ {
		v.regexp_ += "?"
		v.greedy_ = true // Reset scanning back to "greedy".
	}
}

func (v *scanner_) PreprocessExpression(
	expression ast.ExpressionLike,
	index uint,
) {
	v.regexp_ = `"(?:`
}

func (v *scanner_) PostprocessExpression(
	expression ast.ExpressionLike,
	index uint,
) {
	v.regexp_ += `)"`
	var name = expression.GetLowercase()
	v.regexps_.SetValue(name, v.regexp_)
}

func (v *scanner_) PreprocessExtent(extent ast.ExtentLike) {
	v.regexp_ += "-"
}

func (v *scanner_) PreprocessFiltered(filtered ast.FilteredLike) {
	v.regexp_ += "["
}

func (v *scanner_) PostprocessFiltered(filtered ast.FilteredLike) {
	v.regexp_ += "]"
}

func (v *scanner_) PreprocessGrouped(grouped ast.GroupedLike) {
	v.regexp_ += "("
}

func (v *scanner_) PostprocessGrouped(grouped ast.GroupedLike) {
	v.regexp_ += ")"
}

func (v *scanner_) PreprocessIdentifier(identifier ast.IdentifierLike) {
	var name = identifier.GetAny().(string)
	if gra.Scanner().MatchesType(name, gra.LowercaseToken) {
		v.tokens_.AddValue(name)
	}
}

func (v *scanner_) PreprocessLimit(limit ast.LimitLike) {
	v.regexp_ += ","
	var number = limit.GetOptionalNumber()
	if col.IsDefined(number) {
		v.regexp_ += number
	}
}

func (v *scanner_) PreprocessPattern(pattern ast.PatternLike) {
	v.pattern_ = true
}

func (v *scanner_) PostprocessPattern(pattern ast.PatternLike) {
	v.pattern_ = false
}

func (v *scanner_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.greedy_ = true // The default is "greedy" scanning.
	v.ignored_ = col.Set[string]([]string{"newline", "space"})
	v.tokens_ = col.Set[string]([]string{"reserved"})
	v.reserved_ = col.Set[string]()
	var implicit = map[string]string{"space": `"(?:[ \\t]+)"`}
	v.regexps_ = col.Catalog[string, string](implicit)
}

func (v *scanner_) PostprocessSyntax(syntax ast.SyntaxLike) {
	v.ignored_ = v.ignored_.GetClass().Sans(v.ignored_, v.tokens_)
	v.tokens_.AddValues(v.ignored_)
	var reserved = v.extractReserved()
	v.regexps_.SetValue("reserved", reserved)
	v.regexps_.SortValues()
}

// Public

func (v *scanner_) GenerateScannerClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.visitor_.VisitSyntax(syntax)
	implementation = scannerTemplate_
	var notice = v.extractNotice(syntax)
	implementation = sts.ReplaceAll(implementation, "<Notice>", notice)
	var tokenNames = v.extractTokenNames()
	implementation = sts.ReplaceAll(implementation, "<TokenNames>", tokenNames)
	var tokenMatchers = v.extractTokenMatchers()
	implementation = sts.ReplaceAll(implementation, "<TokenMatchers>", tokenMatchers)
	var foundCases = v.extractFoundCases()
	implementation = sts.ReplaceAll(implementation, "<FoundCases>", foundCases)
	var ignoredCases = v.extractIgnoredCases()
	implementation = sts.ReplaceAll(implementation, "<IgnoredCases>", ignoredCases)
	var expressions = v.extractExpressions()
	implementation = sts.ReplaceAll(implementation, "<Expressions>", expressions)
	return implementation
}

// Private

func (v *scanner_) escapeText(text string) string {
	var escaped string
	for _, character := range text {
		switch character {
		case '"':
			escaped += `\`
		case '.', '|', '^', '$', '+', '*', '?',
			'(', ')', '[', ']', '{', '}':
			escaped += `\\`
		case '\\':
			escaped += `\\\`
		}
		escaped += string(character)
	}
	return escaped
}

func (v *scanner_) extractExpressions() string {
	var expressions = "// Define the regular expression patterns for each token type."
	var iterator = v.regexps_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var name = association.GetKey()
		var regexp = association.GetValue()
		expressions += "\n\t" + name + "_ = " + regexp
	}
	return expressions
}

func (v *scanner_) extractFoundCases() string {
	var foundCases = "// Find the next token type."
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var tokenType = v.makeUppercase(tokenName) + "Token"
		foundCases += "\n\t\tcase v.foundToken(" + tokenType + "):"
	}
	return foundCases
}

func (v *scanner_) extractIgnoredCases() string {
	var ignoreCases = "// Ignore the implicit token types."
	var iterator = v.ignored_.GetIterator()
	for iterator.HasNext() {
		var tokenType = iterator.GetNext()
		ignoreCases += "\n\tcase \"" + tokenType + "\":"
		ignoreCases += "\n\t\treturn"
	}
	return ignoreCases
}

func (v *scanner_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

func (v *scanner_) extractReserved() string {
	var reserved = `"(?:`
	if !v.reserved_.IsEmpty() {
		var iterator = v.reserved_.GetIterator()
		reserved += iterator.GetNext()
		for iterator.HasNext() {
			reserved += "|" + iterator.GetNext()
		}
	}
	reserved += `)"`
	return reserved
}

func (v *scanner_) extractTokenMatchers() string {
	var tokenMatchers = "// Define pattern matchers for each type of token."
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var tokenType = v.makeUppercase(tokenName) + "Token"
		tokenMatchers += "\n\t\t" + tokenType +
			`: reg.MustCompile("^" + ` + tokenName + `_),`
	}
	return tokenMatchers
}

func (v *scanner_) extractTokenNames() string {
	var tokenNames = `ErrorToken: "error",`
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var tokenType = v.makeUppercase(tokenName) + "Token"
		tokenNames += "\n\t\t" + tokenType + `: "` + tokenName + `",`
	}
	return tokenNames
}

func (v *scanner_) makeUppercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

var reserved_ = map[string]bool{
	"any":       true,
	"byte":      true,
	"case":      true,
	"complex":   true,
	"copy":      true,
	"default":   true,
	"error":     true,
	"false":     true,
	"import":    true,
	"interface": true,
	"map":       true,
	"nil":       true,
	"package":   true,
	"range":     true,
	"real":      true,
	"return":    true,
	"rune":      true,
	"string":    true,
	"switch":    true,
	"true":      true,
	"type":      true,
}

const scannerTemplate_ = `/*<Notice>*/

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
	first_    int // A zero based index of the first possible rune in the next token.
	next_     int // A zero based index of the next possible rune in the next token.
	line_     int // The line number in the source string of the next rune.
	position_ int // The position in the current line of the next rune.
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
		var length = len(token)

		// Found the requested token type.
		v.next_ += length
		v.emitToken(tokenType)
		var count = sts.Count(match, "\n")
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

func (v *scanner_) indexOfLastEol(runes []rune) int {
	var length = len(runes)
	for index := length; index > 0; index-- {
		if runes[index-1] == '\n' {
			return length - index + 1
		}
	}
	return 0
}

func (v *scanner_) scanTokens() {
loop:
	for v.next_ < len(v.runes_) {
		switch {
		<FoundCases>
		default:
			v.foundError()
			break loop
		}
	}
	v.tokens_.CloseQueue()
}
`
