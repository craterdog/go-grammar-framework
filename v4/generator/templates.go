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

var templates_ = map[string]map[string]string{
	"ast":       astTemplates_,
	"grammar":   grammarTemplates_,
	"token":     tokenTemplates_,
	"scanner":   scannerTemplates_,
	"parser":    parserTemplates_,
	"validator": validatorTemplates_,
	"formatter": formatterTemplates_,
}

var astTemplates_ = map[string]string{
	"notice":      astNoticeTemplate_,
	"header":      astHeaderTemplate_,
	"imports":     astImportsTemplate_,
	"types":       astTypesTemplate_,
	"functionals": astFunctionalsTemplate_,
	"classes":     astClassesTemplate_,
	"instances":   astInstancesTemplate_,
	"aspects":     astAspectsTemplate_,
}

var grammarTemplates_ = map[string]string{
	"notice":      grammarNoticeTemplate_,
	"header":      grammarHeaderTemplate_,
	"imports":     grammarImportsTemplate_,
	"types":       grammarTypesTemplate_,
	"functionals": grammarFunctionalsTemplate_,
	"classes":     grammarClassesTemplate_,
	"instances":   grammarInstancesTemplate_,
	"aspects":     grammarAspectsTemplate_,
}

var tokenTemplates_ = map[string]string{
	"notice":   tokenNoticeTemplate_,
	"header":   tokenHeaderTemplate_,
	"imports":  tokenImportsTemplate_,
	"access":   tokenAccessTemplate_,
	"class":    tokenClassTemplate_,
	"instance": tokenInstanceTemplate_,
}

var scannerTemplates_ = map[string]string{
	"notice":   scannerNoticeTemplate_,
	"header":   scannerHeaderTemplate_,
	"imports":  scannerImportsTemplate_,
	"access":   scannerAccessTemplate_,
	"class":    scannerClassTemplate_,
	"instance": scannerInstanceTemplate_,
}

var parserTemplates_ = map[string]string{
	"notice":   parserNoticeTemplate_,
	"header":   parserHeaderTemplate_,
	"imports":  parserImportsTemplate_,
	"access":   parserAccessTemplate_,
	"class":    parserClassTemplate_,
	"instance": parserInstanceTemplate_,
}

var validatorTemplates_ = map[string]string{
	"notice":   validatorNoticeTemplate_,
	"header":   validatorHeaderTemplate_,
	"imports":  validatorImportsTemplate_,
	"access":   validatorAccessTemplate_,
	"class":    validatorClassTemplate_,
	"instance": validatorInstanceTemplate_,
}

var formatterTemplates_ = map[string]string{
	"notice":   formatterNoticeTemplate_,
	"header":   formatterHeaderTemplate_,
	"imports":  formatterImportsTemplate_,
	"access":   formatterAccessTemplate_,
	"class":    formatterClassTemplate_,
	"instance": formatterInstanceTemplate_,
}

// GENERAL TEMPLATES

const noticeTemplate_ = `
................................................................................
<Copyright>
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
`

// SYNTAX TEMPLATES

const syntaxTemplate_ = `!><Notice><!

!>
<SYNTAX> NOTATION
This document contains a formal definition of the <Syntax> Notation
using Crater Dog Syntax Notation™ (CDSN):
 * https://github.com/craterdog/go-grammar-framework/blob/main/v4/Syntax.cdsn

A language syntax consists of a set of rule definitions and regular expression
patterns.

Each predicate within a rule definition may be constrained by one of the
following cardinalities:
 * predicate{M} - Exactly M instances of the specified predicate.
 * predicate{M..N} - M to N instances of the specified predicate.
 * predicate{M..} - M or more instances of the specified predicate.
 * predicate? - Zero or one instances of the specified predicate.
 * predicate* - Zero or more instances of the specified predicate.
 * predicate+ - One or more instances of the specified predicate.

The following intrinsic character types may be used within regular expression
pattern declarations:
 * ANY - Any language specific character.
 * LOWER - Any language specific lowercase character.
 * UPPER - Any language specific uppercase character.
 * DIGIT - Any language specific digit.
 * CONTROL - Any environment specific (non-printable) control character.
 * EOL - The environment specific end-of-line character.

The negation "~" prefix within a regular expression pattern may only be applied
to a bounded range of possible intrinsic character types or printable unicode
characters called runes.
<!

!>
RULE DEFINITIONS
The following rules are used by the parser when parsing the stream of tokens
generated by the scanner based on the expression patterns.  Each rule name
begins with an uppercase letter.  The rule definitions may specify the names of
expressions or other rules and are matched by the parser in the order listed.  A
rule definition may also be directly or indirectly recursive.  The parsing of
tokens is greedy and will match as many repeated token types as possible. The
sequence of factors within in a rule definition may be separated by spaces which
are ignored by the parser.  Newlines are also ignored unless a "newline" regular
expression pattern is defined and used in one or more rule definitions.
<!
Document: Component newline*

Component:
    Intrinsic
    List

Intrinsic:
    integer
    rune
    text

List: "[" Component Additional* "]"

Additional: "," Component

!>
EXPRESSION DEFINITIONS
The following expression definitions are used by the scanner to generate the
stream of tokens—each an instance of an expression type—that are to be processed by
the parser.  Each expression name begins with a lowercase letter.  Unlike with
rule definitions, an expression definition cannot specify the name of a rule within
its definition, but it may specify the name of another expression.  Expression
definitions cannot be recursive and the scanning of expressions is NOT greedy.
Any spaces within an expression definition are part of the expression and are NOT
ignored.
<!
integer: '0' | '-'? ['1'..'9'] DIGIT*

rune: "'" ~[CONTROL] "'"  ! Any single printable unicode character.

text: '"' ~['"' CONTROL]+ '"'

`

// AST TEMPLATES

const astNoticeTemplate_ = `/*<Notice>*/
`

const astHeaderTemplate_ = `
/*
Package "ast" provides the abstract syntax tree (AST) classes for this module.
Each AST class manages the attributes associated with the rule definition found
in the syntax grammar with the same rule name as the class.

For detailed documentation on this package refer to the wiki:
  - https://<wiki>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast
`

const astImportsTemplate_ = `
import (
	ast "github.com/craterdog/go-collection-framework/v4/collection"
)
`

const astTypesTemplate_ = ``

const astFunctionalsTemplate_ = ``

const astClassesTemplate_ = `
// Classes

/*
This is a dummy class placeholder.
*/
type DummyClassLike interface {
	// Constructors
	Make() DummyLike
}
`

const astInstancesTemplate_ = `
// Instances

/*
This is a dummy instance placeholder.
*/
type DummyLike interface {
	// Attributes
	GetClass() DummyClassLike
}
`

const astAspectsTemplate_ = ``

// AGENT TEMPLATES

const grammarNoticeTemplate_ = `/*<Notice>*/
`

const grammarHeaderTemplate_ = `
/*
Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.

For detailed documentation on this package refer to the wiki:
  - https://<wiki>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package grammar
`

const grammarImportsTemplate_ = `
import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "<module>/ast"
)
`

const grammarTypesTemplate_ = `
// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	<TokenTypes>
)`

const grammarFunctionalsTemplate_ = ``

const grammarClassesTemplate_ = `
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

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructors
	Make(
		source string,
		tokens abs.QueueLike[TokenLike],
	) ScannerLike

	// Functions
	FormatToken(token TokenLike) string
	FormatType(tokenType TokenType) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructors
	Make(
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
`

const grammarInstancesTemplate_ = `
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

const grammarAspectsTemplate_ = ``

// TOKEN TEMPLATES

const tokenNoticeTemplate_ = `/*<Notice>*/
`

const tokenHeaderTemplate_ = `
package grammar
`

const tokenImportsTemplate_ = ``

const tokenAccessTemplate_ = `
// CLASS ACCESS

// Reference

var tokenClass = &tokenClass_{
	// Initialize the class constants.
}

// Function

func Token() TokenClassLike {
	return tokenClass
}
`

const tokenClassTemplate_ = `
// CLASS METHODS

// Target

type tokenClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *tokenClass_) Make(
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
`

const tokenInstanceTemplate_ = `
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

// SCANNER TEMPLATES

const scannerNoticeTemplate_ = `/*<Notice>*/
`

const scannerHeaderTemplate_ = `
package grammar
`

const scannerImportsTemplate_ = `
import (
	fmt "fmt"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	reg "regexp"
	sts "strings"
)
`

const scannerAccessTemplate_ = `
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
`

const scannerClassTemplate_ = `
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
`

const scannerInstanceTemplate_ = `
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

// PARSER TEMPLATES

const parserNoticeTemplate_ = `/*<Notice>*/
`

const parserHeaderTemplate_ = `
package grammar
`

const parserImportsTemplate_ = `
import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "<module>/ast"
	sts "strings"
)
`

const parserAccessTemplate_ = `
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
`

const parserClassTemplate_ = `
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
`

const parserInstanceTemplate_ = `
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
	"<Name>": "Component newline*",
}
`

// VALIDATOR TEMPLATES

const validatorNoticeTemplate_ = `/*<Notice>*/
`

const validatorHeaderTemplate_ = `
package grammar
`

const validatorImportsTemplate_ = `
import (
	fmt "fmt"
	ast "<module>/ast"
)
`

const validatorAccessTemplate_ = `
// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// Initialize the class constants.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}
`

const validatorClassTemplate_ = `
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
`

const validatorInstanceTemplate_ = `
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
	// TBA - Add a real method implementation.
	var name = "foobar"
	if !v.matchesToken(ErrorToken, name) {
		var message = v.formatError(name, "Oops!")
		panic(message)
	}
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
`

// FORMATTER TEMPLATES

const formatterNoticeTemplate_ = `/*<Notice>*/
`

const formatterHeaderTemplate_ = `
package grammar
`

const formatterImportsTemplate_ = `
import (
	ast "<module>/ast"
	sts "strings"
)
`

const formatterAccessTemplate_ = `
// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// Initialize the class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}
`

const formatterClassTemplate_ = `
// CLASS METHODS

// Target

type formatterClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{
		// Initialize the instance attributes.
		class_:   c,
	}
}
`

const formatterInstanceTemplate_ = `
// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_   FormatterClassLike
	depth_   uint
	result_  sts.Builder
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) GetDepth() uint {
	return v.depth_
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

const classCommentTemplate_ = `/*
<Class>ClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete <class>-like class.
*/
`

const instanceCommentTemplate_ = `/*
<Class>Like is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete <class>-like class.
*/
`
