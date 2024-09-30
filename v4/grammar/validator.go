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

package grammar

import (
	fmt "fmt"
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
	class_   *validatorClass_
	visitor_ VisitorLike

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

func (v *validator_) ValidateSyntax(syntax ast.SyntaxLike) {
	v.visitor_.VisitSyntax(syntax)
}

// Methodical

func (v *validator_) ProcessComment(comment string) {
	v.ValidateToken(comment, CommentToken)
}

func (v *validator_) ProcessExcluded(excluded string) {
	v.ValidateToken(excluded, ExcludedToken)
}

func (v *validator_) ProcessGlyph(glyph string) {
	v.ValidateToken(glyph, GlyphToken)
}

func (v *validator_) ProcessIntrinsic(intrinsic string) {
	v.ValidateToken(intrinsic, IntrinsicToken)
}

func (v *validator_) ProcessLiteral(literal string) {
	v.ValidateToken(literal, LiteralToken)
}

func (v *validator_) ProcessLowercase(lowercase string) {
	v.ValidateToken(lowercase, LowercaseToken)
}

func (v *validator_) ProcessNewline(
	newline string,
	index uint,
	size uint,
) {
	v.ValidateToken(newline, NewlineToken)
}

func (v *validator_) ProcessNote(note string) {
	v.ValidateToken(note, NoteToken)
}

func (v *validator_) ProcessNumber(number string) {
	v.ValidateToken(number, NumberToken)
}

func (v *validator_) ProcessOptional(optional string) {
	v.ValidateToken(optional, OptionalToken)
}

func (v *validator_) ProcessRepeated(repeated string) {
	v.ValidateToken(repeated, RepeatedToken)
}

func (v *validator_) ProcessSpace(space string) {
	v.ValidateToken(space, SpaceToken)
}

func (v *validator_) ProcessUppercase(uppercase string) {
	v.ValidateToken(uppercase, UppercaseToken)
}

func (v *validator_) PreprocessSyntax(syntax ast.SyntaxLike) {
}

func (v *validator_) ProcessSyntaxSlot(slot uint) {
}

func (v *validator_) PostprocessSyntax(syntax ast.SyntaxLike) {
}
