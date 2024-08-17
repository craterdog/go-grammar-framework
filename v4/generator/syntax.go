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

// CLASS ACCESS

import (
	fmt "fmt"
	sts "strings"
	tim "time"
	uni "unicode"
)

// Reference

var syntaxClass = &syntaxClass_{
	// Initialize the class constants.
}

// Function

func Syntax() SyntaxClassLike {
	return syntaxClass
}

// CLASS METHODS

// Target

type syntaxClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *syntaxClass_) Make() SyntaxLike {
	return &syntax_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type syntax_ struct {
	// Define the instance attributes.
	class_ SyntaxClassLike
}

// Attributes

func (v *syntax_) GetClass() SyntaxClassLike {
	return v.class_
}

// Public

func (v *syntax_) GenerateSyntaxNotation(
	syntax string,
	copyright string,
) (
	implementation string,
) {
	implementation = syntaxTemplate_
	var allCaps = sts.ToUpper(syntax)
	implementation = sts.ReplaceAll(implementation, "<SYNTAX>", allCaps)
	var uppercase = v.makeUppercase(syntax)
	implementation = sts.ReplaceAll(implementation, "<Syntax>", uppercase)
	var lowercase = v.makeLowercase(syntax)
	implementation = sts.ReplaceAll(implementation, "<syntax>", lowercase)
	copyright = v.expandCopyright(copyright)
	implementation = sts.ReplaceAll(implementation, "<Copyright>", copyright)
	return implementation
}

// Private

func (v *syntax_) expandCopyright(copyright string) string {
	var limit = 78
	var length = len(copyright)
	if length > limit {
		var message = fmt.Sprintf(
			"The copyright notice cannot be longer than 78 characters: %v",
			copyright,
		)
		panic(message)
	}
	if length == 0 {
		copyright = fmt.Sprintf(
			"Copyright (c) %v.  All Rights Reserved.",
			tim.Now().Year(),
		)
		length = len(copyright)
	}
	var padding = (limit - length) / 2
	for range padding {
		copyright = " " + copyright + " "
	}
	if len(copyright) < limit {
		copyright = " " + copyright
	}
	copyright = "." + copyright + "."
	return copyright
}

func (v *syntax_) makeLowercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	name = string(runes)
	if reserved_[name] {
		name += "_"
	}
	return name
}

func (v *syntax_) makeUppercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

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

The excluded "~" prefix within a regular expression pattern may only be applied
to a bounded range of possible character types.
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
integer: '0' | ('-'? ['1'..'9'] DIGIT*)

rune: "'" ~[CONTROL] "'"  ! Any single printable unicode character.

text: '"' ~['"' CONTROL]+ '"'

`
