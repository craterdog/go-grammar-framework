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

package generator_test

import (
	gen "github.com/craterdog/go-grammar-framework/v4/generator"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
	//mod "github.com/craterdog/go-model-framework/v4"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

func TestLifecycle(t *tes.T) {
	var module = "github.com/craterdog/go-test-framework/v4"
	var wiki = "github.com/craterdog/go-test-framework/wiki"

	/*
		var name = "example"

		// Generate a new syntax with a default copyright.
		var copyright string
		var source = gen.Syntax().Make().GenerateSyntaxNotation(name, copyright)
		ass.Equal(t, syntaxNotation, source)
	*/

	// Read in the test syntax file.
	var filename = "../../../go-test-framework/v4/Syntax.cdsn"
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)

	// Parse the source code for the syntax.
	var parser = gra.Parser().Make()
	var syntax = parser.ParseSource(source)

	// Validate the syntax.
	var validator = gra.Validator().Make()
	validator.ValidateSyntax(syntax)

	// Format the syntax.
	var formatter = gra.Formatter().Make()
	var formatted = formatter.FormatSyntax(syntax)
	ass.Equal(t, formatted, source)

	// Generate the processor class for the syntax.
	source = gen.Processor().Make().GenerateProcessorClass(module, syntax)
	bytes = []byte(source)
	filename = "../../../go-test-framework/v4/grammar/processor.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
	//ass.Equal(t, processorClass, source)

	// Generate the visitor class for the syntax.
	source = gen.Visitor().Make().GenerateVisitorClass(module, syntax)
	bytes = []byte(source)
	filename = "../../../go-test-framework/v4/grammar/visitor.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
	//ass.Equal(t, visitorClass, source)

	// Generate the token class for the syntax.
	source = gen.Token().Make().GenerateTokenClass(module, syntax)
	bytes = []byte(source)
	filename = "../../../go-test-framework/v4/grammar/token.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
	//ass.Equal(t, tokenClass, source)

	// Generate the scanner class for the syntax.
	source = gen.Scanner().Make().GenerateScannerClass(module, syntax)
	bytes = []byte(source)
	filename = "../../../go-test-framework/v4/grammar/scanner.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
	//ass.Equal(t, scannerClass, source)

	/*
		// Generate the formatter class for the syntax.
		source = gen.Formatter().Make().GenerateFormatterClass(module, syntax)
		bytes = []byte(source)
		filename = "../../../go-test-framework/v4/grammar/formatter.go"
		err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
		//ass.Equal(t, formatterClass, source)
	*/

	// Generate the parser class for the syntax.
	source = gen.Parser().Make().GenerateParserClass(module, syntax)
	bytes = []byte(source)
	filename = "../../../go-test-framework/v4/grammar/parser.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
	//ass.Equal(t, parserClass, source)

	// Generate the validator class for the syntax.
	source = gen.Validator().Make().GenerateValidatorClass(module, syntax)
	bytes = []byte(source)
	filename = "../../../go-test-framework/v4/grammar/validator.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
	//ass.Equal(t, validatorClass, source)

	// Generate the language grammar model for the syntax.
	source = gen.Grammar().Make().GenerateGrammarModel(module, wiki, syntax)
	bytes = []byte(source)
	filename = "../../../go-test-framework/v4/grammar/Package.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	/*
		// ass.Equal(t, grammarModel, source)
		var model = mod.Parser().ParseSource(source)
		mod.Validator().ValidateModel(model)
	*/

	// Generate the abstract syntax tree model for the syntax.
	source = gen.Ast().Make().GenerateAstModel(wiki, syntax)
	bytes = []byte(source)
	filename = "../../../go-test-framework/v4/ast/Package.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
	panic("stop")
	/*
		ass.Equal(t, astModel, source)
		model = mod.Parser().ParseSource(source)
		mod.Validator().ValidateModel(model)
	*/
}

const syntaxNotation = `!>
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
<!

!>
EXAMPLE NOTATION
This document contains a formal definition of the Example Notation
using Crater Dog Syntax Notation™ (CDSN):
 * https://github.com/craterdog/go-grammar-framework/blob/main/v4/Syntax.cdsn

A language syntax consists of a set of rule definitions and regular expression
patterns.

Most terms within a rule definition can be constrained by one of the following
cardinalities:
  - term{M} - Exactly M instances of the specified term.
  - term{M..N} - M to N instances of the specified term.
  - term{M..} - M or more instances of the specified term.
  - term* - Zero or more instances of the specified term.
  - term+ - One or more instances of the specified term.
  - term? - An optional term.

The following intrinsic character types may be used within regular expression
pattern declarations:
  - ANY - Any language specific character.
  - LOWER - Any language specific lowercase character.
  - UPPER - Any language specific uppercase character.
  - DIGIT - Any language specific digit.
  - CONTROL - Any environment specific (non-printable) control character.
  - EOL - The environment specific end-of-line character.

The excluded "~" prefix within a regular expression pattern may only be applied
to a filtered set of possible characters.

RULE DEFINITIONS
The following rules are used by the parser when parsing the stream of tokens
generated by the scanner based on the expression patterns.  Each rule name
begins with an uppercase letter.  The rule definitions may specify the names of
expressions or other rules and are matched by the parser in the order listed.  A
rule definition may also be directly or indirectly recursive.  The parsing of
tokens is greedy and will match as many repeated token types as possible. The
sequence of terms within in a rule definition may be separated by spaces which
are ignored by the parser.  Newlines are also ignored unless a "newline" regular
expression pattern is defined and used in one or more rule definitions.
<!
Document: Component newline+

Component:
  - Intrinsic
  - List

Intrinsic:
  - integer
  - rune
  - text

List: "[" Component AdditionalComponent* "]"

AdditionalComponent: "," Component Component

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
integer: '0' | ('-'? ['1'..'9'] DIGIT*)

rune: "'" ~[CONTROL] "'"  ! Any single printable unicode character.

text: '"' ~['"' CONTROL]+ '"'

`

const formatterClass = `/*
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

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	stc "strconv"
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// Initialize the class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	var formatter = &formatter_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	formatter.visitor_ = Visitor().Make(formatter)
	return formatter
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_   *formatterClass_
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Define the inherited aspects.
	Methodical
}

// Public

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) FormatDocument(document ast.DocumentLike) string {
	v.visitor_.VisitDocument(document)
	return v.getResult()
}

// Methodical

func (v *formatter_) ProcessInteger(integer string) {
	v.appendString(integer)
}

func (v *formatter_) ProcessNewline(
	newline string,
	index uint,
	size uint,
) {
	v.appendString(newline)
}

func (v *formatter_) ProcessRune(rune_ string) {
	v.appendString(rune_)
}

func (v *formatter_) ProcessText(text string) {
	v.appendString(text)
}

func (v *formatter_) PreprocessAdditionalComponent(
	additionalComponent ast.AdditionalComponentLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAdditionalComponent(
	additionalComponent ast.AdditionalComponentLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessComponent(component ast.ComponentLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessComponent(component ast.ComponentLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessDocument(document ast.DocumentLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessDocument(document ast.DocumentLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessIntrinsic(intrinsic ast.IntrinsicLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessIntrinsic(intrinsic ast.IntrinsicLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessList(list ast.ListLike) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessList(list ast.ListLike) {
	// TBD - Add formatting of the delimited rule.
}

// Private

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "    "
	var level uint
	for ; level < v.depth_; level++ {
		newline += indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
`

const parserClass = `/*
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

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
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
	class_     *parserClass_
	ruleFound_ bool
	source_    string                   // The original source code.
	tokens_    abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_      abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Public

func (v *parser_) GetClass() ParserClassLike {
	return v.class_
}

func (v *parser_) ParseSource(source string) ast.DocumentLike {
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserClass.queueSize_)
	v.next_ = col.Stack[TokenLike](parserClass.stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse the document.
	var document, token, ok = v.parseDocument()
	if !ok {
		var message = v.formatError(token, "Document")
		panic(message)
	}

	// Found the document.
	return document
}

// Private

func (v *parser_) parseAdditionalComponent() (
	additionalComponent ast.AdditionalComponentLike,
	token TokenLike,
	ok bool,
) {
	v.ruleFound_ = false

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if v.ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token,"AdditionalComponent")
			panic(message)
		} else {
			// This is not a single additionalComponent rule.
			return additionalComponent, token, false
		}
	}
	v.ruleFound_ = true

	// Attempt to parse a single component rule.
	var component1 ast.ComponentLike
	component1, token, ok = v.parseComponent()
	if !ok {
		if v.ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token,"AdditionalComponent")
			panic(message)
		} else {
			// This is not a single additionalComponent rule.
			return additionalComponent, token, false
		}
	}
	v.ruleFound_ = true

	// Attempt to parse a single component rule.
	var component2 ast.ComponentLike
	component2, token, ok = v.parseComponent()
	if !ok {
		if v.ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token,"AdditionalComponent")
			panic(message)
		} else {
			// This is not a single additionalComponent rule.
			return additionalComponent, token, false
		}
	}
	v.ruleFound_ = true

	// Found a single additionalComponent rule.
	additionalComponent = ast.AdditionalComponent().Make(
		component1,
		component2,
	)
	return additionalComponent, token, true

}

func (v *parser_) parseComponent() (
	component ast.ComponentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single intrinsic rule.
	var intrinsic ast.IntrinsicLike
	intrinsic, token, ok = v.parseIntrinsic()
	if ok {
		// Found a single intrinsic component.
		component = ast.Component().Make(intrinsic)
		return component, token, true
	}

	// Attempt to parse a single list rule.
	var list ast.ListLike
	list, token, ok = v.parseList()
	if ok {
		// Found a single list component.
		component = ast.Component().Make(list)
		return component, token, true
	}

	// This is not a single component rule.
	return component, token, false

}

func (v *parser_) parseDocument() (
	document ast.DocumentLike,
	token TokenLike,
	ok bool,
) {
	v.ruleFound_ = false

	// Attempt to parse a single component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if !ok {
		if v.ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token,"Document")
			panic(message)
		} else {
			// This is not a single document rule.
			return document, token, false
		}
	}
	v.ruleFound_ = true

	// Attempt to parse 1 to unlimited newline tokens.
	var newlines = col.List[string]()
newlinesLoop:
	for i := 0; i < unlimited; i++ {
		var newline string
		newline, token, ok = v.parseToken(NewlineToken)
		if !ok {
			switch {
			case i < 1:
				if !v.ruleFound_ {
					// This is not a single document rule.
					return document, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Document")
				message += "Too few newline tokens found."
				panic(message)
			case i > unlimited:
				// Found a syntax error.
				var message = v.formatError(token, "Document")
				message += "Too many newline tokens found."
				panic(message)
			default:
				break newlinesLoop
			}
		}
		newlines.AppendValue(newline)
	}

	// Found a single document rule.
	document = ast.Document().Make(
		component,
		newlines,
	)
	return document, token, true

}

func (v *parser_) parseIntrinsic() (
	intrinsic ast.IntrinsicLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single integer token.
	var integer string
	integer, token, ok = v.parseToken(IntegerToken)
	if ok {
		// Found a single integer intrinsic.
		intrinsic = ast.Intrinsic().Make(integer)
		return intrinsic, token, true
	}

	// Attempt to parse a single rune token.
	var rune_ string
	rune_, token, ok = v.parseToken(RuneToken)
	if ok {
		// Found a single rune intrinsic.
		intrinsic = ast.Intrinsic().Make(rune_)
		return intrinsic, token, true
	}

	// Attempt to parse a single text token.
	var text string
	text, token, ok = v.parseToken(TextToken)
	if ok {
		// Found a single text intrinsic.
		intrinsic = ast.Intrinsic().Make(text)
		return intrinsic, token, true
	}

	// This is not a single intrinsic rule.
	return intrinsic, token, false

}

func (v *parser_) parseList() (
	list ast.ListLike,
	token TokenLike,
	ok bool,
) {
	v.ruleFound_ = false

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if v.ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token,"List")
			panic(message)
		} else {
			// This is not a single list rule.
			return list, token, false
		}
	}
	v.ruleFound_ = true

	// Attempt to parse a single component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if !ok {
		if v.ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token,"List")
			panic(message)
		} else {
			// This is not a single list rule.
			return list, token, false
		}
	}
	v.ruleFound_ = true

	// Attempt to parse 0 to unlimited additionalComponent rules.
	var additionalComponents = col.List[ast.AdditionalComponentLike]()
additionalComponentsLoop:
	for i := 0; i < unlimited; i++ {
		var additionalComponent ast.AdditionalComponentLike
		additionalComponent, token, ok = v.parseAdditionalComponent()
		if !ok {
			switch {
			case i < 0:
				if !v.ruleFound_ {
					// This is not a single list rule.
					return list, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "List")
				message += "Too few additionalComponent rules found."
				panic(message)
			case i > unlimited:
				// Found a syntax error.
				var message = v.formatError(token, "List")
				message += "Too many additionalComponent rules found."
				panic(message)
			default:
				break additionalComponentsLoop
			}
		}
		additionalComponents.AppendValue(additionalComponent)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if v.ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token,"List")
			panic(message)
		} else {
			// This is not a single list rule.
			return list, token, false
		}
	}
	v.ruleFound_ = true

	// Found a single list rule.
	list = ast.List().Make(
		component,
		additionalComponents,
	)
	return list, token, true

}

func (v *parser_) parseDelimiter(expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
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
	// Attempt to parse a specific token.
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
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(ruleName string) string {
	return syntax_.GetValue(ruleName)
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

// PRIVATE GLOBALS

// Constants

const unlimited = 4294967295 // Default to a reasonable value.

var syntax_ = col.Catalog[string, string](
	map[string]string{
		"Document": ` + "`" + `Component newline+` + "`" + `,
		"Component": ` + "`" + `
  - Intrinsic
  - List` + "`" + `,
		"Intrinsic": ` + "`" + `
  - integer
  - rune
  - text` + "`" + `,
		"List": ` + "`" + `"[" Component AdditionalComponent* "]"` + "`" + `,
		"AdditionalComponent": ` + "`" + `"," Component Component` + "`" + `,
	},
)
`

const scannerClass = `/*
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
		ErrorToken: "error",
		DelimiterToken: "delimiter",
		IntegerToken: "integer",
		NewlineToken: "newline",
		RuneToken: "rune",
		SpaceToken: "space",
		TextToken: "text",
	},
	matchers_: map[TokenType]*reg.Regexp{
		// Define pattern matchers for each type of token.
		DelimiterToken: reg.MustCompile("^" + delimiter_),
		IntegerToken: reg.MustCompile("^" + integer_),
		NewlineToken: reg.MustCompile("^" + newline_),
		RuneToken: reg.MustCompile("^" + rune_),
		SpaceToken: reg.MustCompile("^" + space_),
		TextToken: reg.MustCompile("^" + text_),
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
	class_    *scannerClass_
	first_    uint // A zero based index of the first possible rune in the next token.
	next_     uint // A zero based index of the next possible rune in the next token.
	line_     uint // The line number in the source string of the next rune.
	position_ uint // The position in the current line of the next rune.
	runes_    []rune
	tokens_   abs.QueueLike[TokenLike]
}

// Public

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

	// Define the regular expression patterns for each token type.
	delimiter_ = "(?:,|\\[|\\])"
	integer_ = "(?:0|(-?[1-9]" + digit_ + "*))"
	rune_ = "(?:'[^" + control_ + "]')"
	space_ = "(?:[ \\t]+)"
	text_ = "(?:\"[^\"" + control_ + "]+\")"
)

func (v *scanner_) emitToken(tokenType TokenType) {
	switch v.GetClass().FormatType(tokenType) {
	// Ignore the implicit token types.
	case "space":
		return
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
		// Find the next token type.
		case v.foundToken(DelimiterToken):
		case v.foundToken(IntegerToken):
		case v.foundToken(NewlineToken):
		case v.foundToken(RuneToken):
		case v.foundToken(SpaceToken):
		case v.foundToken(TextToken):
		default:
			v.foundError()
			break loop
		}
	}
	v.tokens_.CloseQueue()
}
`

const tokenClass = `/*
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

package grammar

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

func (c *tokenClass_) Make(
	line uint,
	position uint,
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
	class_    *tokenClass_
	line_     uint
	position_ uint
	type_     TokenType
	value_    string
}

// Public

func (v *token_) GetClass() TokenClassLike {
	return v.class_
}

// Attributes

func (v *token_) GetLine() uint {
	return v.line_
}

func (v *token_) GetPosition() uint {
	return v.position_
}

func (v *token_) GetType() TokenType {
	return v.type_
}

func (v *token_) GetValue() string {
	return v.value_
}
`

const validatorClass = `/*
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

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	stc "strconv"
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
	var validator = &validator_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	validator.visitor_ = Visitor().Make(validator)
	return validator
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// Define the instance attributes.
	class_       *validatorClass_
	visitor_     VisitorLike

	// Define the inherited aspects.
	Methodical
}

// Public

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

func (v *validator_) ValidateToken(
	tokenValue string,
	tokenType TokenType,
) {
	if !Scanner().MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			Scanner().FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}

func (v *validator_) ValidateDocument(document ast.DocumentLike) {
	v.visitor_.VisitDocument(document)
}

// Methodical

func (v *validator_) ProcessInteger(integer string) {
	v.ValidateToken(integer, IntegerToken)
}

func (v *validator_) ProcessNewline(
	newline string,
	index uint,
	size uint,
) {
	v.ValidateToken(newline, NewlineToken)
}

func (v *validator_) ProcessRune(rune_ string) {
	v.ValidateToken(rune_, RuneToken)
}

func (v *validator_) ProcessText(text string) {
	v.ValidateToken(text, TextToken)
}

func (v *validator_) PreprocessDocument(document ast.DocumentLike) {
}

func (v *validator_) PostprocessDocument(document ast.DocumentLike) {
}
`

const processorClass = `/*
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

package grammar

import (
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
)

// CLASS ACCESS

// Reference

var processorClass = &processorClass_{
	// Initialize the class constants.
}

// Function

func Processor() ProcessorClassLike {
	return processorClass
}

// CLASS METHODS

// Target

type processorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *processorClass_) Make() ProcessorLike {
	var processor = &processor_{
		// Initialize the instance attributes.
		class_: c,
	}
	return processor
}

// INSTANCE METHODS

// Target

type processor_ struct {
	// Define the instance attributes.
	class_ *processorClass_
}

// Public

func (v *processor_) GetClass() ProcessorClassLike {
	return v.class_
}

// Methodical

func (v *processor_) ProcessInteger(integer string) {
}

func (v *processor_) ProcessNewline(
	newline string,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessRune(rune_ string) {
}

func (v *processor_) ProcessText(text string) {
}

func (v *processor_) PreprocessAdditionalComponent(
	additionalComponent ast.AdditionalComponentLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PostprocessAdditionalComponent(
	additionalComponent ast.AdditionalComponentLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessComponent(component ast.ComponentLike) {
}

func (v *processor_) PostprocessComponent(component ast.ComponentLike) {
}

func (v *processor_) PreprocessDocument(document ast.DocumentLike) {
}

func (v *processor_) PostprocessDocument(document ast.DocumentLike) {
}

func (v *processor_) PreprocessIntrinsic(intrinsic ast.IntrinsicLike) {
}

func (v *processor_) PostprocessIntrinsic(intrinsic ast.IntrinsicLike) {
}

func (v *processor_) PreprocessList(list ast.ListLike) {
}

func (v *processor_) PostprocessList(list ast.ListLike) {
}
`

const visitorClass = `/*
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

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
)

// CLASS ACCESS

// Reference

var visitorClass = &visitorClass_{
	// Initialize the class constants.
}

// Function

func Visitor() VisitorClassLike {
	return visitorClass
}

// CLASS METHODS

// Target

type visitorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *visitorClass_) Make(processor Methodical) VisitorLike {
	return &visitor_{
		// Initialize the instance attributes.
		class_:     c,
		processor_: processor,
	}
}

// INSTANCE METHODS

// Target

type visitor_ struct {
	// Define the instance attributes.
	class_     *visitorClass_
	processor_ Methodical
}

// Public

func (v *visitor_) GetClass() VisitorClassLike {
	return v.class_
}

func (v *visitor_) VisitDocument(document ast.DocumentLike) {
	// Visit the document syntax.
	v.processor_.PreprocessDocument(document)
	v.visitDocument(document)
	v.processor_.PostprocessDocument(document)
}

// Private

func (v *visitor_) visitAdditionalComponent(additionalComponent ast.AdditionalComponentLike) {
	// Visit the component rule.
	var component1 = additionalComponent.GetComponent1()
	v.processor_.PreprocessComponent(component1)
	v.visitComponent(component1)
	v.processor_.PostprocessComponent(component1)

	// Visit the component rule.
	var component2 = additionalComponent.GetComponent2()
	v.processor_.PreprocessComponent(component2)
	v.visitComponent(component2)
	v.processor_.PostprocessComponent(component2)
}

func (v *visitor_) visitComponent(component ast.ComponentLike) {
	// Visit the possible component types.
	switch actual := component.GetAny().(type) {
	case ast.IntrinsicLike:
		v.processor_.PreprocessIntrinsic(actual)
		v.visitIntrinsic(actual)
		v.processor_.PostprocessIntrinsic(actual)
	case ast.ListLike:
		v.processor_.PreprocessList(actual)
		v.visitList(actual)
		v.processor_.PostprocessList(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitDocument(document ast.DocumentLike) {
	// Visit the component rule.
	var component = document.GetComponent()
	v.processor_.PreprocessComponent(component)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(component)

	// Visit each newline token.
	var newlineIndex uint
	var newlines = document.GetNewlines().GetIterator()
	var newlinesSize = uint(newlines.GetSize())
	for newlines.HasNext() {
		newlineIndex++
		var newline = newlines.GetNext()
		v.processor_.ProcessNewline(
			newline,
			newlineIndex,
			newlinesSize,
		)
	}
}

func (v *visitor_) visitIntrinsic(intrinsic ast.IntrinsicLike) {
	// Visit the possible intrinsic types.
	switch actual := intrinsic.GetAny().(type) {
	case string:
		switch {
		case Scanner().MatchesType(actual, IntegerToken):
			v.processor_.ProcessInteger(actual)
		case Scanner().MatchesType(actual, RuneToken):
			v.processor_.ProcessRune(actual)
		case Scanner().MatchesType(actual, TextToken):
			v.processor_.ProcessText(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitList(list ast.ListLike) {
	// Visit the component rule.
	var component = list.GetComponent()
	v.processor_.PreprocessComponent(component)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(component)

	// Visit each additionalComponent rule.
	var additionalComponentIndex uint
	var additionalComponents = list.GetAdditionalComponents().GetIterator()
	var additionalComponentsSize = uint(additionalComponents.GetSize())
	for additionalComponents.HasNext() {
		additionalComponentIndex++
		var additionalComponent = additionalComponents.GetNext()
		v.processor_.PreprocessAdditionalComponent(
			additionalComponent,
			additionalComponentIndex,
			additionalComponentsSize,
		)
		v.visitAdditionalComponent(additionalComponent)
		v.processor_.PostprocessAdditionalComponent(
			additionalComponent,
			additionalComponentIndex,
			additionalComponentsSize,
		)
	}
}
`

const grammarModel = `/*
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

/*
Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-grammar-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-test-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package grammar

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
)

// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	DelimiterToken
	IntegerToken
	NewlineToken
	RuneToken
	SpaceToken
	TextToken
)

// Classes

/*
FormatterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor
	Make() FormatterLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor
	Make() ParserLike
}

/*
ProcessorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor
	Make() ProcessorLike
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
	// Constructor
	Make(
		source string,
		tokens abs.QueueLike[TokenLike],
	) ScannerLike

	// Function
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
	// Constructor
	Make(
		line uint,
		position uint,
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
	// Constructor
	Make() ValidatorLike
}

/*
VisitorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor
	Make(processor Methodical) VisitorLike
}

// Instances

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Public
	GetClass() FormatterClassLike
	FormatDocument(document ast.DocumentLike) string

	// Aspect
	Methodical
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Public
	GetClass() ParserClassLike
	ParseSource(source string) ast.DocumentLike
}

/*
ProcessorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Public
	GetClass() ProcessorClassLike

	// Aspect
	Methodical
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Public
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Public
	GetClass() TokenClassLike

	// Attribute
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Public
	GetClass() ValidatorClassLike
	ValidateDocument(document ast.DocumentLike)

	// Aspect
	Methodical
}

/*
VisitorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Public
	GetClass() VisitorClassLike
	VisitDocument(document ast.DocumentLike)
}

// Aspects

/*
Methodical defines the set of method signatures that must be supported
by all methodical processors.
*/
type Methodical interface {
	ProcessInteger(integer string)
	ProcessNewline(
		newline string,
		index uint,
		size uint,
	)
	ProcessRune(rune string)
	ProcessText(text string)
	PreprocessAdditionalComponent(
		additionalComponent ast.AdditionalComponentLike,
		index uint,
		size uint,
	)
	PostprocessAdditionalComponent(
		additionalComponent ast.AdditionalComponentLike,
		index uint,
		size uint,
	)
	PreprocessComponent(component ast.ComponentLike)
	PostprocessComponent(component ast.ComponentLike)
	PreprocessDocument(document ast.DocumentLike)
	PostprocessDocument(document ast.DocumentLike)
	PreprocessIntrinsic(intrinsic ast.IntrinsicLike)
	PostprocessIntrinsic(intrinsic ast.IntrinsicLike)
	PreprocessList(list ast.ListLike)
	PostprocessList(list ast.ListLike)
}
`

const astModel = `/*
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

/*
Package "ast" provides the abstract syntax tree (AST) classes for this module.
Each AST class manages the attributes associated with the rule definition found
in the syntax grammar with the same rule name as the class.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-grammar-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-test-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// Classes

/*
AdditionalComponentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete additional-component-like class.
*/
type AdditionalComponentClassLike interface {
	// Constructor
	Make(
		component1 ComponentLike,
		component2 ComponentLike,
	) AdditionalComponentLike
}

/*
ComponentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructor
	Make(any_ any) ComponentLike
}

/*
DocumentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructor
	Make(
		component ComponentLike,
		newlines abs.Sequential[string],
	) DocumentLike
}

/*
IntrinsicClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete intrinsic-like class.
*/
type IntrinsicClassLike interface {
	// Constructor
	Make(any_ any) IntrinsicLike
}

/*
ListClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete list-like class.
*/
type ListClassLike interface {
	// Constructor
	Make(
		component ComponentLike,
		additionalComponents abs.Sequential[AdditionalComponentLike],
	) ListLike
}

// Instances

/*
AdditionalComponentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete additional-component-like class.
*/
type AdditionalComponentLike interface {
	// Public
	GetClass() AdditionalComponentClassLike

	// Attribute
	GetComponent1() ComponentLike
	GetComponent2() ComponentLike
}

/*
ComponentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Public
	GetClass() ComponentClassLike

	// Attribute
	GetAny() any
}

/*
DocumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete document-like class.
*/
type DocumentLike interface {
	// Public
	GetClass() DocumentClassLike

	// Attribute
	GetComponent() ComponentLike
	GetNewlines() abs.Sequential[string]
}

/*
IntrinsicLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete intrinsic-like class.
*/
type IntrinsicLike interface {
	// Public
	GetClass() IntrinsicClassLike

	// Attribute
	GetAny() any
}

/*
ListLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete list-like class.
*/
type ListLike interface {
	// Public
	GetClass() ListClassLike

	// Attribute
	GetComponent() ComponentLike
	GetAdditionalComponents() abs.Sequential[AdditionalComponentLike]
}
`
