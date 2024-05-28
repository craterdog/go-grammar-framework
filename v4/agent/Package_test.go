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

package agent_test

import (
	age "github.com/craterdog/go-grammar-framework/v4/agent"
	mod "github.com/craterdog/go-model-framework/v4"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestLifecycle(t *tes.T) {
	var generator = age.Generator().Make()
	var module = "github.com/craterdog/go-grammar-framework/v4"
	var name = "example"

	// Generate a new syntax with a default copyright.
	var copyright string
	var syntax = generator.CreateSyntax(name, copyright)

	// Format the syntax.
	var formatter = age.Formatter().Make()
	var source = formatter.FormatSyntax(syntax)

	// Parse the source code for the syntax.
	var parser = age.Parser().Make()
	syntax = parser.ParseSource(source)

	// Validate the syntax.
	var validator = age.Validator().Make()
	validator.ValidateSyntax(syntax)

	// Generate the AST model for the syntax.
	var model = generator.GenerateAST(module, syntax)
	mod.Validator().ValidateModel(model)

	// Generate the agent model for the syntax.
	model = generator.GenerateAgent(module, syntax)
	mod.Validator().ValidateModel(model)

	// Generate the formatter class for the syntax.
	source = generator.GenerateFormatter(module, syntax, model)
	ass.Equal(t, modelFormatter, source)

	// Generate the parser class for the syntax.
	source = generator.GenerateParser(module, syntax, model)
	ass.Equal(t, modelParser, source)

	// Generate the scanner class for the syntax.
	source = generator.GenerateScanner(module, syntax, model)
	ass.Equal(t, modelScanner, source)

	// Generate the token class for the syntax.
	source = generator.GenerateToken(module, syntax, model)
	ass.Equal(t, modelToken, source)

	// Generate the validator class for the syntax.
	source = generator.GenerateValidator(module, syntax, model)
	ass.Equal(t, modelValidator, source)
}

const modelFormatter = `/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent

import (
	sts "strings"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	defaultMaximum_: 8,
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	defaultMaximum_ int
}

// Constants

func (c *formatterClass_) DefaultMaximum() int {
	return c.defaultMaximum_
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{
		maximum_: c.defaultMaximum_,
	}
}

func (c *formatterClass_) MakeWithMaximum(maximum int) FormatterLike {
	if maximum < 0 {
		maximum = c.defaultMaximum_
	}
	return &formatter_{
		class_:   c,
		maximum_: maximum,
	}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	class_   FormatterClassLike
	depth_   int
	maximum_ int
	result_  sts.Builder
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) GetDepth() int {
	return v.depth_
}

func (v *formatter_) GetMaximum() int {
	return v.maximum_
}

// Public

func (v *formatter_) FormatComponent(component ast.ComponentLike) string {
	v.formatComponent(component)
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

func (v *formatter_) formatComponent(component ast.ComponentLike) {
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

const modelParser = `/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent

import (
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
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
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type parser_ struct {
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

func (v *parser_) ParseSource(source string) ast.ComponentLike {
	v.source_ = source
	var notation = cdc.Notation().Make()
	v.tokens_ = col.Queue[TokenLike](notation).MakeWithCapacity(parserClass.queueSize_)
	v.next_ = col.Stack[TokenLike](notation).MakeWithCapacity(parserClass.stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse the component.
	var component, token, ok = v.parseComponent()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Component",
			"AST",
			"Component",
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
			"AST",
			"Component",
		)
		panic(message)
	}

	// Found the component.
	return component
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

func (v *parser_) parseComponent() (
	component ast.ComponentLike,
	token TokenLike,
	ok bool,
) {
	// TBA - Add real method implementation.
	return component, token, ok
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
	"AST": "Component EOL* EOF  ! Terminated with an end-of-file marker.",
}
`

const modelScanner = `/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent

import (
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
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

const modelToken = `/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent

// CLASS ACCESS

// Reference

var tokenClass = &tokenClass_{
	// This class has no private constants to initialize.
}

// Function

func Token() TokenClassLike {
	return tokenClass
}

// CLASS METHODS

// Target

type tokenClass_ struct {
	// This class has no private constants.
}

// Constructors

func (c *tokenClass_) MakeWithAttributes(
	line int,
	position int,
	type_ TokenType,
	value string,
) TokenLike {
	return &token_{
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

const modelValidator = `/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent

import (
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
)

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
	class_    ValidatorClassLike
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

func (v *validator_) ValidateComponent(component ast.ComponentLike) {
	// TBA - Add method implementation.
}

// Private
`
