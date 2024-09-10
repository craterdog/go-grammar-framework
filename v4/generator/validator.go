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
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
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
		class_:    c,
		analyzer_: gra.Analyzer().Make(),
	}
	return validator
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// Define the instance attributes.
	class_    ValidatorClassLike
	analyzer_ gra.AnalyzerLike
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
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = validatorTemplate_
	implementation = replaceAll(implementation, "module", module)
	var notice = v.analyzer_.GetNotice()
	implementation = replaceAll(implementation, "notice", notice)
	var tokenValidators = v.generateTokenValidators()
	implementation = replaceAll(implementation, "tokenValidators", tokenValidators)
	var name = v.analyzer_.GetName()
	implementation = replaceAll(implementation, "name", name)
	return implementation
}

// Private

func (v *validator_) generateTokenValidators() string {
	var tokenValidators string
	var iterator = v.analyzer_.GetTokens().GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		if v.analyzer_.IsIgnored(tokenName) || tokenName == "delimiter" {
			continue
		}
		var isPlural = v.analyzer_.IsPlural(tokenName)
		var parameters = validateTokenParameterTemplate_
		if isPlural {
			parameters = validateTokenParametersTemplate_
		}
		var tokenValidator = validateTokenTemplate_
		tokenValidator = replaceAll(tokenValidator, "parameters", parameters)
		tokenValidator = replaceAll(tokenValidator, "tokenName", tokenName)
		tokenValidators += tokenValidator
	}
	return tokenValidators
}

const validateTokenTemplate_ = `
func (v *validator_) Process<TokenName>(<parameters>) {
	v.ValidateToken(<tokenName_>, <TokenName>Token)
}
`

const validateTokenParameterTemplate_ = `<tokenName_> string`

const validateTokenParametersTemplate_ = `
	<tokenName_> string,
	index uint,
	size uint,
`

const validatorTemplate_ = `<Notice>

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

// Methodical
<TokenValidators>
func (v *validator_) Preprocess<Name>(<name> ast.<Name>Like) {
}

func (v *validator_) Postprocess<Name>(<name> ast.<Name>Like) {
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
`
