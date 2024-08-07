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
		Methodical: gra.Processor().Make(),
	}
	validator.visitor_ = gra.Visitor().Make(validator)
	return validator
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// Define the instance attributes.
	class_   ValidatorClassLike
	visitor_ gra.VisitorLike
	tokens_  abs.SetLike[string]

	// Define the inherited aspects.
	gra.Methodical
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

func (v *validator_) GenerateValidatorClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.visitor_.VisitSyntax(syntax)
	implementation = validatorTemplate_
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
	var validateTokens = v.extractValidateTokens()
	implementation = sts.ReplaceAll(
		implementation,
		"<ProcessTokens>",
		validateTokens,
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

func (v *validator_) PreprocessIdentifier(identifier ast.IdentifierLike) {
	var name = identifier.GetAny().(string)
	if gra.Scanner().MatchesType(name, gra.LowercaseToken) {
		v.tokens_.AddValue(name)
	}
}

func (v *validator_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.tokens_ = col.Set[string]([]string{"delimiter"})
}

// Private

func (v *validator_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

func (v *validator_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *validator_) extractValidateTokens() string {
	var validateTokens string
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var validateToken = validateTokenTemplate_
		var tokenName = iterator.GetNext()
		validateToken = sts.ReplaceAll(validateToken, "<tokenName>", tokenName)
		tokenName = v.makeUppercase(tokenName)
		validateToken = sts.ReplaceAll(validateToken, "<TokenName>", tokenName)
		var tokenType = tokenName + "Token"
		validateToken = sts.ReplaceAll(validateToken, "<TokenType>", tokenType)
		validateTokens += validateToken
	}
	return validateTokens
}

func (v *validator_) makeLowercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	name = string(runes)
	if reserved_[name] {
		name += "_"
	}
	return name
}

func (v *validator_) makeUppercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

const validateTokenTemplate_ = `
func (v *validator_) Process<TokenName>(<tokenName> string) {
	v.ValidateToken(<tokenName>, <TokenType>)
}
`

const validatorTemplate_ = `/*<Notice>*/

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "<module>/ast"
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
	class_       ValidatorClassLike
	visitor_     VisitorLike

	// Define the inherited aspects.
	Methodical
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

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

func (v *validator_) Validate<Name>(<name> ast.<Name>Like) {
	v.visitor_.Visit<Name>(<name>)
}

// Methodical
<ProcessTokens>
func (v *validator_) Preprocess<Name>(<name> ast.<Name>Like) {
}

func (v *validator_) Postprocess<Name>(<name> ast.<Name>Like) {
}
`
