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
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
	stc "strconv"
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
	class_ *syntaxClass_
}

// Public

func (v *syntax_) GetClass() SyntaxClassLike {
	return v.class_
}

func (v *syntax_) GenerateSyntaxNotation(
	syntax string,
	copyright string,
) (
	implementation string,
) {
	var template = v.getTemplate(syntaxTemplate)
	implementation = replaceAll(template, "syntax", syntax)
	copyright = expandCopyright(copyright)
	implementation = replaceAll(implementation, "copyright", copyright)
	return implementation
}

// Private

func (v *syntax_) getTemplate(name string) string {
	var template = syntaxTemplates_.GetValue(name)
	return template
}

// PRIVATE GLOBALS

// Functions

func expandCopyright(copyright string) string {
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

func generateVariableName(reference ast.ReferenceLike) string {
	var mixedCase = reference.GetIdentifier().GetAny().(string)
	var variableName = makeLowerCase(mixedCase)
	var cardinality = reference.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case ast.ConstrainedLike:
			var constrained = actual.GetAny().(string)
			switch constrained {
			case "?":
				variableName = makeOptional(variableName)
			case "*", "+":
				variableName = makePlural(variableName)
			}
		case ast.QuantifiedLike:
			variableName = makePlural(variableName)
		}
	}
	return variableName
}

func generateVariableNames(
	references abs.Sequential[ast.ReferenceLike],
) abs.Sequential[string] {
	var variableNames = col.List[string]()

	// Extract the reference identifiers as attribute names.
	var iterator = references.GetIterator()
	for iterator.HasNext() {
		var reference = iterator.GetNext()
		var variableName = generateVariableName(reference)
		variableNames.AppendValue(variableName)
	}

	// Normalize any duplicate names.
	for i := 1; i <= variableNames.GetSize(); i++ {
		var count = 1
		var firstName = variableNames.GetValue(i)
		for j := i + 1; j <= variableNames.GetSize(); j++ {
			var secondName = variableNames.GetValue(j)
			if firstName == secondName {
				count++
				secondName += stc.Itoa(count)
				variableNames.SetValue(j, secondName)
			}
		}
		if count > 1 {
			firstName += "1"
			variableNames.SetValue(i, firstName)
		}
	}

	return variableNames
}

func generateVariableType(
	reference ast.ReferenceLike,
) (
	variableType string,
) {
	var identifier = reference.GetIdentifier().GetAny().(string)
	switch {
	case gra.Scanner().MatchesType(identifier, gra.LowercaseToken):
		variableType = "string"
	case gra.Scanner().MatchesType(identifier, gra.UppercaseToken):
		variableType = makeUpperCase(identifier) + "Like"
	}
	return variableType
}

func isReserved(name string) bool {
	return reserved_.ContainsValue(name)
}

func makeAllCaps(mixedCase string) string {
	var allCaps sts.Builder
	for _, r := range mixedCase {
		switch {
		case uni.IsLower(r):
			allCaps.WriteRune(uni.ToUpper(r))
		case uni.IsUpper(r):
			allCaps.WriteString("_")
			allCaps.WriteRune(r)
		default:
			allCaps.WriteRune(r)
		}
	}
	return allCaps.String()
}

func makeLowerCase(mixedCase string) string {
	var lowerCase string
	if len(mixedCase) > 0 {
		runes := []rune(mixedCase)
		runes[0] = uni.ToLower(runes[0])
		lowerCase = string(runes)
	}
	return lowerCase
}

func makeOptional(mixedCase string) string {
	var optional string
	if len(mixedCase) > 0 {
		optional = "optional" + makeUpperCase(mixedCase)
	}
	return optional
}

func makePlural(mixedCase string) string {
	var plural string
	if sts.HasSuffix(mixedCase, "s") {
		plural = mixedCase + "es"
	} else {
		plural = mixedCase + "s"
	}
	return plural
}

func makeSnakeCase(mixedCase string) string {
	mixedCase = makeLowerCase(mixedCase)
	var snakeCase sts.Builder
	for _, r := range mixedCase {
		switch {
		case uni.IsLower(r):
			snakeCase.WriteRune(r)
		case uni.IsUpper(r):
			snakeCase.WriteString("-")
			snakeCase.WriteRune(uni.ToLower(r))
		default:
			snakeCase.WriteRune(r)
		}
	}
	return snakeCase.String()
}

func makeUpperCase(mixedCase string) string {
	var upperCase string
	if len(mixedCase) > 0 {
		runes := []rune(mixedCase)
		runes[0] = uni.ToUpper(runes[0])
		upperCase = string(runes)
	}
	return upperCase
}

func replaceAll(template string, name string, value string) string {
	// <variableName> -> variableValue[_]
	var variableName = makeLowerCase(name) + "_"
	var variableValue = makeLowerCase(value)
	if isReserved(variableValue) {
		variableValue += "_"
	}
	template = sts.ReplaceAll(template, "<"+variableName+">", variableValue)

	// <lowerCaseName> -> lowerCaseValue
	var lowerCaseName = makeLowerCase(name)
	var lowerCaseValue = makeLowerCase(value)
	template = sts.ReplaceAll(template, "<"+lowerCaseName+">", lowerCaseValue)

	// <snake-case-name> -> snake-case-value
	var snakeCaseName = makeSnakeCase(name)
	var snakeCaseValue = makeSnakeCase(value)
	template = sts.ReplaceAll(template, "<"+snakeCaseName+">", snakeCaseValue)

	// <UpperCaseName> -> UpperCaseValue
	var upperCaseName = makeUpperCase(name)
	var upperCaseValue = makeUpperCase(value)
	template = sts.ReplaceAll(template, "<"+upperCaseName+">", upperCaseValue)

	// <ALL_CAPS_NAME> -> ALL_CAPS_VALUE
	var allCapsName = makeAllCaps(name)
	var allCapsValue = makeAllCaps(value)
	template = sts.ReplaceAll(template, "<"+allCapsName+">", allCapsValue)

	return template
}

// Constants

var reserved_ = col.Set[string](
	[]string{
		"any",
		"byte",
		"case",
		"complex",
		"copy",
		"default",
		"error",
		"false",
		"import",
		"interface",
		"map",
		"nil",
		"package",
		"range",
		"real",
		"return",
		"rune",
		"string",
		"switch",
		"true",
		"type",
	},
)

const (
	syntaxTemplate = "syntaxTemplate"
)

var syntaxTemplates_ = col.Catalog[string, string](
	map[string]string{
		syntaxTemplate: `!>
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

`,
	},
)
