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

const modelTemplate_ = `<Notice>

/*
Package "<package>" provides...

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
package <package>
`

const astTemplate_ = `

import (
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// Classes

/*
This is a dummy class placeholder.
*/
type DummyClassLike interface {}

// Instances

/*
This is a dummy instance placeholder.
*/
type DummyLike interface {}
`

const agentTemplate_ = `
import (
	col "github.com/craterdog/go-collection-framework/v4/collection"
	ast "<module>/ast"
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
	// Constants
	DefaultMaximum() uint

	// Constructors
	Make() FormatterLike
	MakeWithMaximum(maximum uint) FormatterLike
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
	// Attributes
	GetClass() FormatterClassLike
	GetDepth() uint
	GetMaximum() uint

	// Methods
	Format<Name>(<parameter> ast.<Name>Like) string
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Attributes
	GetClass() ParserClassLike

	// Methods
	ParseSource(source string) ast.<Name>Like
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Attributes
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Attributes
	GetClass() TokenClassLike
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
	// Attributes
	GetClass() ValidatorClassLike

	// Methods
	Validate<Name>(<parameter> ast.<Name>Like)
}
`

const classCommentTemplate_ = `/*
<Class>ClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete <class>-like class.
*/`

const instanceCommentTemplate_ = `/*
<Class>Like is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete <class>-like class.
*/`

const syntaxTemplate_ = `!>
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
 * https://github.com/craterdog/go-grammar-framework/blob/main/v4/Syntax.cdsn

A language syntax consists of a set of rule definitions and lexigram
definitions.

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
 * predicate{M..N} - M to N instances of the specified predicate.
 * predicate{M..} - M or more instances of the specified predicate.
 * predicate? - Zero or one instances of the specified predicate.
 * predicate* - Zero or more instances of the specified predicate.
 * predicate+ - One or more instances of the specified predicate.

A negation "~" within a lexigram definition may only be applied to a bounded
range of possible intrinsic character types or printable unicode characters
called runes.
<!

!>
RULE DEFINITIONS
The following rules are used by the parser when parsing the stream of tokens
generated by the scanner based on the lexigram definitions.  Each rule name
begins with an uppercase letter.  The rule definitions may specify the names of
lexigrams or other rules and are matched by the parser in the order listed.  A
rule definition may also be directly or indirectly recursive.  The parsing of
tokens is greedy and will match as many repeated token types as possible. The
sequence of factors within in a rule definition may be separated by spaces which
are ignored by the parser.
<!
Document: Component EOL* EOF  ! Terminated with an end-of-file marker.

Component:
    Default
    Primitive
    List

Default: "default"

Primitive:
    rune
    text
    integer
    anything

List: "[" Component{3..5} Additional* "]"

Additional: "," Component

!>
LEXIGRAM DEFINITIONS
The following lexigram definitions are used by the scanner to generate the
stream of tokens—each an instance of a lexigram type—that are to be processed by
the parser.  Each lexigram name begins with a lowercase letter.  Unlike with
rule definitions, a lexigram definition cannot specify the name of a rule within
its definition, but it may specify the name of another lexigram.  Lexigram
definitions cannot be recursive and the scanning of lexigrams is NOT greedy.
Any spaces within a lexigram definition are part of the lexigram and are NOT
ignored.
<!
rune: "'" ~[CONTROL] "'"  ! Any single printable unicode character.

text: '"' (ESCAPE ~['"' CONTROL]{2..})+ '"'

integer: '0'{4} | '-'? '1'..'9' DIGIT*

anything: ANY

`

const tokenTemplate_ = `<Notice>

package agent

// CLASS ACCESS

// Reference

var tokenClass = &tokenClass_{
	// Initialize the class constants.
}

// Function

func Token() TokenClassLike {
	return tokenClass
}

// CLASS METHODS

// Target

type tokenClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *tokenClass_) MakeWithAttributes(
	line int,
	position int,
	type_ TokenType,
	value string,
) TokenLike {
	return &token_{
		// Initialize the instance attributes.
		class_:    c,
		line_:     line,
		position_: position,
		type_:     type_,
		value_:    value,
	}
}

// INSTANCE METHODS

// Target

type token_ struct {
	// Define the instance attributes.
	class_    TokenClassLike
	line_     int
	position_ int
	type_     TokenType
	value_    string
}

// Attributes

func (v *token_) GetClass() TokenClassLike {
	return v.class_
}

func (v *token_) GetLine() int {
	return v.line_
}

func (v *token_) GetPosition() int {
	return v.position_
}

func (v *token_) GetType() TokenType {
	return v.type_
}

func (v *token_) GetValue() string {
	return v.value_
}
`

const scannerTemplate_ = `<Notice>
package agent

import (
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	reg "regexp"
	sts "strings"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var scannerClass = &scannerClass_{
	// Initialize the class constants.
	tokens_: map[TokenType]string{
		// TBA - Add additional token types.
		ErrorToken:     "error",
		DelimiterToken: "delimiter",
		EOFToken:       "EOF",
		EOLToken:       "EOL",
		SpaceToken:     "space",
	},
	matchers_: map[TokenType]*reg.Regexp{
		// TBA - Add additional token types.
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
	// Define the class constants.
	tokens_   map[TokenType]string
	matchers_ map[TokenType]*reg.Regexp
}

// Constructors

func (c *scannerClass_) Make(
	source string,
	tokens col.QueueLike[TokenLike],
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

func (c *scannerClass_) AsString(type_ TokenType) string {
	return c.tokens_[type_]
}

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
	var notation = cdc.Notation().Make()
	return col.List[string](notation).MakeFromArray(matches)
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
	tokens_   col.QueueLike[TokenLike]
}

// Attributes

func (v *scanner_) GetClass() ScannerClassLike {
	return v.class_
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

		// Check for false intrinsic match.
		var nextIndex = v.next_ + length
		if nextIndex < len(v.runes_) {
			var nextRune = v.runes_[v.next_+length]
			if type_ == IntrinsicToken && (uni.IsLetter(nextRune) ||
				uni.IsDigit(nextRune) || nextRune == rune('_')) {
				// This is not an intrinsic token.
				return false
			}
		}

		// Found the requested token type.
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

	// The next token is not the requested token type.
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
	// TBA - Add additional token types.
	any_       = ` + "`" + `.|` + "`" + ` + eol_
	base16_    = ` + "`" + `[0-9a-f]` + "`" + `
	control_   = ` + "`" + `\p{Cc}` + "`" + `
	delimiter_ = ` + "`" + `[:;,\.=]` + "`" + ` // TBA - Replace with the actual delimeters.
	digit_     = ` + "`" + `\p{Nd}` + "`" + `
	eof_       = ` + "`" + `\z` + "`" + `
	eol_       = ` + "`" + `\n` + "`" + `
	escape_    = ` + "`" + `\\(?:(?:` + "`" + ` + unicode_ + ` + "`" + `)|[abfnrtv'"\\])` + "`" + `
	letter_    = lower_ + ` + "`" + `|` + "`" + ` + upper_
	lower_     = ` + "`" + `\p{Ll}` + "`" + `
	number_    = ` + "`" + `(?:` + "`" + ` + digit_ + ` + "`" + `)+` + "`" + `
	rune_      = ` + "`" + `['][^` + "`" + ` + control_ + ` + "`" + `][']` + "`" + `
	space_     = ` + "`" + `[ \t]+` + "`" + `
	string_    = ` + "`" + `["](?:` + "`" + ` + escape_ + ` + "`" + `|[^"` + "`" + ` + control_ + ` + "`" + `])+?["]` + "`" + `
	unicode_   = ` + "`" + `x` + "`" + ` + base16_ + ` + "`" + `{2}|u` + "`" + ` + base16_ + ` + "`" + `{4}|U` + "`" + ` + base16_ + ` + "`" + `{8}` + "`" + `
	upper_     = ` + "`" + `\p{Lu}` + "`" + `
)
`

const parserTemplate_ = `<Notice>

package agent

import (
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
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
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Attributes

func (v *parser_) GetClass() ParserClassLike {
	return v.class_
}

// Public

func (v *parser_) ParseSource(source string) ast.<Name>Like {
	v.source_ = source
	var notation = cdc.Notation().Make()
	v.tokens_ = col.Queue[TokenLike](notation).MakeWithCapacity(parserClass.queueSize_)
	v.next_ = col.Stack[TokenLike](notation).MakeWithCapacity(parserClass.stackSize_)

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

	// Attempt to parse optional end-of-line characters.
	for ok {
		_, _, ok = v.parseToken(EOLToken, "")
	}

	// Attempt to parse the end-of-file marker.
	_, token, ok = v.parseToken(EOFToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("EOF",
			"<Name>",
		)
		panic(message)
	}

	// Found the <name>.
	return <name>
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

func (v *parser_) parse<Name>() (
	<name> ast.<Name>Like,
	token TokenLike,
	ok bool,
) {
	// TBA - Add real method implementation.
	return <name>, token, ok
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
		var notConstrained = len(expectedValue) == 0
		if notConstrained || value == expectedValue {
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
	"<Name>": "uppercase+ EOL* EOF  ! Terminated with an end-of-file marker.",
}
`

const formatterTemplate_ = `<Notice>

package agent

import (
	ast "<module>/ast"
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// Initialize the class constants.
	defaultMaximum_: 8,
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// Define the class constants.
	defaultMaximum_ uint
}

// Constants

func (c *formatterClass_) DefaultMaximum() uint {
	return c.defaultMaximum_
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{
		// Initialize the instance attributes.
		class_:   c,
		maximum_: c.defaultMaximum_,
	}
}

func (c *formatterClass_) MakeWithMaximum(maximum uint) FormatterLike {
	if maximum == 0 {
		maximum = c.defaultMaximum_
	}
	return &formatter_{
		// Initialize the instance attributes.
		class_:   c,
		maximum_: maximum,
	}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_   FormatterClassLike
	depth_   uint
	maximum_ uint
	result_  sts.Builder
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) GetDepth() uint {
	return v.depth_
}

func (v *formatter_) GetMaximum() uint {
	return v.maximum_
}

// Public

func (v *formatter_) Format<Name>(<name> ast.<Name>Like) string {
	v.format<Name>(<name>)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "\t"
	var level uint
	for ; level < v.depth_; level++ {
		newline += indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) format<Name>(<name> ast.<Name>Like) {
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

const validatorTemplate_ = `<Notice>

package agent

import (
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
)

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// Initialize the class constants.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// CLASS METHODS

// Target

type validatorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *validatorClass_) Make() ValidatorLike {
	return &validator_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// Define the instance attributes.
	class_    ValidatorClassLike
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

func (v *validator_) Validate<Name>(<name> ast.<Name>Like) {
	// TBA - Add method implementation.
}

// Private

func (v *validator_) formatError(name, message string) string {
	message = fmt.Sprintf(
		"The definition for %v is invalid:\n%v\n",
		name,
		message,
	)
	return message
}

func (v *validator_) matchesToken(type_ TokenType, value string) bool {
	var matches = Scanner().MatchToken(type_, value)
	return !matches.IsEmpty()
}
`
