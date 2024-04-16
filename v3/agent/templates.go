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

const modelTemplate_ = `
<Notice>
/*
Package "<packagename>" provides...

For detailed documentation on this package refer to the wiki:
  - <wiki>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package <packagename>

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
)

// Classes

/*
FormatterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructors
	Make() FormatterLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructors
	Make() ParserLike
}

/*
ScannerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

MatchToken() a list of strings representing any matches found in the specified
text of the specified token type using the regular expression defined for that
token type.  If the regular expression contains submatch patterns the matching
substrings are returned as additional values in the list.
*/
type ScannerClassLike interface {
	// Constructors
	Make(
		source string,
		tokens col.QueueLike[TokenLike],
	) ScannerLike

	// Functions
	FormatToken(token TokenLike) string
	MatchToken(
		type_ TokenType,
		text string,
	) col.ListLike[string]
}

/*
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructors
	MakeWithAttributes(
		line int,
		position int,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructors
	Make() ValidatorLike
}

// Instances

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Methods
	Format<Class>(<class> <Class>Like) string
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Methods
	ParseSource(source string) <Class>Like
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Attributes
	GetLine() int
	GetPosition() int
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Methods
	Validate<Class>(<class> <Class>Like)
}
`

const classCommentTemplate_ = `
/*
<ClassName>ClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete <class-name>-like class.
*/
`

const instanceCommentTemplate_ = `
/*
<ClassName>Like is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete <class-name>-like class.
*/
`

const grammarTemplate_ = `
!>
................................................................................
<Copyright>
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
<!

!>
<NAME> NOTATION
This document contains a formal definition of the <Name> Notation
using Crater Dog Syntax Notation™ (CDSN):
 * https://github.com/craterdog/go-grammar-framework/blob/main/v3/Grammar.cdsn

A language grammar consists of a set of rule definitions and token definitions.

The following intrinsic character types are context specific:
 * ANY - Any language specific character.
 * LOWER - Any language specific lowercase character.
 * UPPER - Any language specific uppercase character.
 * DIGIT - Any language specific digit.
 * ESCAPE - Any environment specific escape sequence.
 * CONTROL - Any environment specific (non-printable) control character.
 * EOL - The environment specific end-of-line character.
 * EOF - The environment specific end-of-file marker (pseudo character).

A predicate may be constrained by any of the following cardinalities:
 * predicate{M} - Exactly M instances of the specified predicate.
 * predicate{M..} - M or more instances of the specified predicate.
 * predicate{M..N} - M to N instances of the specified predicate.
 * predicate? - Zero or one instances of the specified predicate.
 * predicate* - Zero or more instances of the specified predicate.
 * predicate+ - One or more instances of the specified predicate.

An inversion "~" within a definition may only be applied to an intrinsic
character type or a glyph range.
<!

!>
RULE DEFINITIONS
The following rules are used by the parser when parsing the stream of tokens
generated by the scanner.  Each rule name begins with an uppercase letter.  The
rule definitions may specify the names of tokens or other rules and are matched
by the parser in the order listed.  A rule definition may also be directly or
indirectly recursive.  The sequence of factors within in a rule definition may
be separated by spaces which are ignored by the parser.
<!
Source: Rule EOF  ! Terminated with an end-of-file marker.

Rule: token+

!>
TOKEN DEFINITIONS
The following token definitions are used by the scanner to generate the stream
of tokens that are processed by the parser.  Each token name begins with a
lowercase letter.  Unlike with rule definitions, a token definition cannot
specify the name of a rule within its definition but it can specify the name of
other tokens.  Token definitions cannot be recursive and the scanning of tokens
is NOT greedy.  Any spaces within a token definition are NOT ignored.
<!
token: UPPER (LOWER | UPPER)*

`

const scannerTemplate_ = `
<Notice>
package <packagename>

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v3/collection"
	reg "regexp"
	sts "strings"
)

// CLASS ACCESS

// Reference

var scannerClass = &scannerClass_{
	tokens_: map[TokenType]string{
		ErrorToken:     "error",
		DelimiterToken: "delimiter",
		EOFToken:       "EOF",
		EOLToken:       "EOL",
		SpaceToken:     "space",
		// TBA - Add additional token types.
	},
	matchers_: map[TokenType]*reg.Regexp{
		DelimiterToken: reg.MustCompile("^(?:" + delimiter_ + ")"),
		EOLToken:       reg.MustCompile("^(?:" + eol_ + ")"),
		SpaceToken:     reg.MustCompile("^(?:" + space_ + ")"),
	},
}

// Function

func Scanner() ScannerClassLike {
	return scannerClass
}

// CLASS METHODS

// Target

type scannerClass_ struct {
	tokens_   map[TokenType]string
	matchers_ map[TokenType]*reg.Regexp
}

// Constructors

func (c *scannerClass_) Make(
	source string,
	tokens col.QueueLike[TokenLike],
) ScannerLike {
	var scanner = &scanner_{
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

func (c *scannerClass_) MatchToken(
	type_ TokenType,
	text string,
) col.ListLike[string] {
	var matcher = c.matchers_[type_]
	var matches = matcher.FindStringSubmatch(text)
	return col.List[string]().MakeFromArray(matches)
}

// INSTANCE METHODS

// Target

type scanner_ struct {
	first_    int // A zero based index of the first possible rune in the next token.
	next_     int // A zero based index of the next possible rune in the next token.
	line_     int // The line number in the source string of the next rune.
	position_ int // The position in the current line of the next rune.
	runes_    []rune
	tokens_   col.QueueLike[TokenLike]
}

// Private

func (v *scanner_) emitToken(type_ TokenType) {
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
	var token = Token().MakeWithAttributes(v.line_, v.position_, type_, value)
	//fmt.Println(Scanner().FormatToken(token)) // Uncomment when debugging.
	v.tokens_.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundEOF() {
	v.emitToken(EOFToken)
}

func (v *scanner_) foundError() {
	v.next_++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(type_ TokenType) bool {
	var text = string(v.runes_[v.next_:])
	var matches = Scanner().MatchToken(type_, text)
	if !matches.IsEmpty() {
		var match = matches.GetValue(1)
		var token = []rune(match)
		var length = len(token)
		v.next_ += length
		if type_ != SpaceToken {
			v.emitToken(type_)
		}
		var count = sts.Count(match, "\n")
		if count > 0 {
			v.line_ += count
			v.position_ = v.indexOfLastEOL(token)
		} else {
			v.position_ += v.next_ - v.first_
		}
		v.first_ = v.next_
		return true
	}
	return false
}

func (v *scanner_) indexOfLastEOL(runes []rune) int {
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
		// TBA - Add additional token types.
		case v.foundToken(DelimiterToken):
		case v.foundToken(EOLToken):
		case v.foundToken(SpaceToken):
		default:
			v.foundError()
			break loop
		}
	}
	v.foundEOF()
}

/*
NOTE:
These private constants define the regular expression sub-patterns that make up
all token types.  Unfortunately there is no way to make them private to the
scanner class since they must be TRUE Go constants to be initialized in this
way.  We append an underscore to each name to lessen the chance of a name
collision with other private Go class constants in this package.
*/
const (
	any_       = ` + "`" + `.|` + "`" + ` + eol_
	base16_    = ` + "`" + `[0-9a-f]` + "`" + `
	control_   = ` + "`" + `\p{Cc}` + "`" + `
	delimiter_ = ` + "`" + `[:;,\.=]` + "`" + `
	digit_     = ` + "`" + `\p{Nd}` + "`" + `
	eof_       = ` + "`" + `\z` + "`" + `
	eol_       = ` + "`" + `\n` + "`" + `
	escape_    = ` + "`" + `\\(?:(?:` + "`" + ` + unicode_ + ` + "`" + `)|[abfnrtv'"\\])` + "`" + `
	lower_     = ` + "`" + `\p{Ll}` + "`" + `
	space_     = ` + "`" + `[ \t]+` + "`" + `
	unicode_   = ` + "`" + `x` + "`" + ` + base16_ + ` + "`" + `{2}|u` + "`" + ` + base16_ + ` + "`" + `{4}|U` + "`" + ` + base16_ + ` + "`" + `{8}` + "`" + `
	upper_     = ` + "`" + `\p{Lu}` + "`" + `
	// TBA - Add additional regular expression definitions.
)
`

const parserTemplate_ = `
<Notice>
package <packagename>

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

func (v *parser_) ParseSource(source string) <ClassName>Like {
	// The scanner runs in a separate Go routine.
	v.source_ = source
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse a model.
	var model, token, ok = v.parse<ClassName>()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("<ClassName>",
			"<PackageName>",
			"<ClassName>",
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
		message += v.generateGrammar("EOF",
			"<PackageName>",
			"<ClassName>",
		)
		panic(message)
	}

	// Found a model.
	return model
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
func (v *parser_) generateGrammar(expected string, names ...string) string {
	var message = "Was expecting '" + expected + "' from:\n"
	for _, name := range names {
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			name,
			grammar[name],
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

func (v *parser_) parse<ClassName>() (
	<className> <ClassName>Like,
	token TokenLike,
	ok bool,
) {
	// TBA - Add real method implementation.
	return <className>, token, ok
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
	//fmt.Printf("Put Back %v\n", token)
	v.next_.AddValue(token)
}

var grammar = map[string]string{
	"<PackageName>": "<ClassName> EOL* EOF  ! Terminated with an end-of-file marker.",
}
`

const formatterTemplate_ = `
<Notice>
package <packagename>

import (
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// This class does not initialize any private class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// This class does not define any private class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	depth_  int
	result_ sts.Builder
}

// Public

func (v *formatter_) Format<ClassName>(<className> <ClassName>Like) string {
	v.format<ClassName>(<className>)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var separator = "\n"
	var indentation = "\t"
	for level := 0; level < v.depth_; level++ {
		separator += indentation
	}
	v.appendString(separator)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) format<ClassName>(<className> <ClassName>Like) {
	// TBA - Add real method implementation.
	v.depth_++
	v.appendString("test")
	v.appendNewline()
	v.depth_--
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
`

const validatorTemplate_ = `
<Notice>
package <packagename>

import ()

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// This class does not initialize any private class constants.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// CLASS METHODS

// Target

type validatorClass_ struct {
	// This class does not define any private class constants.
}

// Constructors

func (c *validatorClass_) Make() ValidatorLike {
	return &validator_{
		// TBA - Initialize private instance attributes.
	}
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// TBA - Add private instance attributes.
}

// Public

func (v *validator_) Validate<ClassName>(<className> <ClassName>Like) {
	// TBA - Add method implementation.
}

// Private
`
