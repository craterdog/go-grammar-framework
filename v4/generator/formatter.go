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
	var processor = gra.Processor().Make()
	var formatter = &formatter_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: processor,
	}
	formatter.visitor_ = gra.Visitor().Make(formatter)
	return formatter
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_   FormatterClassLike
	visitor_ gra.VisitorLike
	tokens_  abs.SetLike[string]

	// Define the inherited aspects.
	gra.Methodical
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

// Public

func (v *formatter_) GenerateFormatterClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.visitor_.VisitSyntax(syntax)
	implementation = formatterTemplate_
	var name = v.extractSyntaxName(syntax)
	implementation = sts.ReplaceAll(
		implementation,
		"<module>",
		module,
	)
	var notice = v.extractNotice(syntax)
	implementation = sts.ReplaceAll(
		implementation,
		"<Notice>",
		notice,
	)
	var tokenProcessors = v.extractTokenProcessors()
	implementation = sts.ReplaceAll(
		implementation,
		"<TokenProcessors>",
		tokenProcessors,
	)
	var uppercase = v.makeUppercase(name)
	implementation = sts.ReplaceAll(
		implementation,
		"<Name>",
		uppercase,
	)
	var lowercase = v.makeLowercase(name)
	implementation = sts.ReplaceAll(
		implementation,
		"<name>",
		lowercase,
	)
	return implementation
}

// Methodical

func (v *formatter_) PreprocessIdentifier(identifier ast.IdentifierLike) {
	var name = identifier.GetAny().(string)
	if gra.Scanner().MatchesType(name, gra.LowercaseToken) {
		v.tokens_.AddValue(name)
	}
}

func (v *formatter_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.tokens_ = col.Set[string]([]string{"reserved"})
}

// Private

func (v *formatter_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

func (v *formatter_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *formatter_) extractTokenProcessors() string {
	var tokenProcessors string
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var tokenProcessor = formatTemplate_
		var tokenName = iterator.GetNext()
		tokenProcessor = sts.ReplaceAll(tokenProcessor, "<tokenName>", tokenName)
		tokenName = v.makeUppercase(tokenName)
		tokenProcessor = sts.ReplaceAll(tokenProcessor, "<TokenName>", tokenName)
		var tokenType = tokenName + "Token"
		tokenProcessor = sts.ReplaceAll(tokenProcessor, "<TokenType>", tokenType)
		tokenProcessors += tokenProcessor
	}
	return tokenProcessors
}

func (v *formatter_) makeLowercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	name = string(runes)
	if reserved_[name] {
		name += "_"
	}
	return name
}

func (v *formatter_) makeUppercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

const formatTemplate_ = `
func (v *formatter_) Process<TokenName>(<tokenName> string) {
	v.appendString(<tokenName>)
}
`

const formatterTemplate_ = `/*<Notice>*/

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "<module>/ast"
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
	var processor = Processor().Make()
	var formatter = &formatter_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: processor,
	}
	formatter.visitor_ = Visitor().Make(formatter)
	return formatter
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_   FormatterClassLike
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Define the inherited aspects.
	Methodical
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
	v.visitor_.Visit<Name>(<name>)
	return v.getResult()
}

// Methodical
<TokenProcessors>
func (v *formatter_) Preprocess<Name>(<name> ast.<Name>Like) {
}

func (v *formatter_) Postprocess<Name>(<name> ast.<Name>Like) {
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
