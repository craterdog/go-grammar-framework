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

/*
Package "module" defines a universal constructor for each class that is exported
by this module.  Each constructor delegates the actual construction process to
one of the classes defined in a subpackage for this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-model-framework/wiki

This package follows the Crater Dog Technologies™ (craterdog) Go Coding
Conventions located here:
  - https://github.com/craterdog/go-model-framework/wiki

The classes defined in this module provide the ability to parse, validate and
format Go Class Model Notation (GCMN).  They can also generate concrete class
implementation files for each abstract class defined in the Package.go file.
*/
package module

import (
	fmt "fmt"
	fwk "github.com/craterdog/go-collection-framework/v4"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	age "github.com/craterdog/go-grammar-framework/v4/agent"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	sts "strings"
	uni "unicode"
)

// TYPE ALIASES

// AST

type (
	AlternativeLike = ast.AlternativeLike
	BoundedLike     = ast.BoundedLike
	CardinalityLike = ast.CardinalityLike
	CharacterLike   = ast.CharacterLike
	ConstrainedLike = ast.ConstrainedLike
	ElementLike     = ast.ElementLike
	ExpressionLike  = ast.ExpressionLike
	FactorLike      = ast.FactorLike
	FilteredLike    = ast.FilteredLike
	InitialLike     = ast.InitialLike
	GroupedLike     = ast.GroupedLike
	HeaderLike      = ast.HeaderLike
	IdentifierLike  = ast.IdentifierLike
	InlinedLike     = ast.InlinedLike
	ExtentLike      = ast.ExtentLike
	LexigramLike    = ast.LexigramLike
	LineLike        = ast.LineLike
	MaximumLike     = ast.MaximumLike
	MinimumLike     = ast.MinimumLike
	MultilinedLike  = ast.MultilinedLike
	PartLike        = ast.PartLike
	PatternLike     = ast.PatternLike
	PredicateLike   = ast.PredicateLike
	RuleLike        = ast.RuleLike
	SyntaxLike      = ast.SyntaxLike
)

// Agents

type (
	FormatterLike = age.FormatterLike
	GeneratorLike = age.GeneratorLike
	ParserLike    = age.ParserLike
	ValidatorLike = age.ValidatorLike
	TokenType     = age.TokenType
)

const ErrorToken = age.ErrorToken
const CommentToken = age.CommentToken
const DelimiterToken = age.DelimiterToken
const EOFToken = age.EOFToken
const EOLToken = age.EOLToken
const IntrinsicToken = age.IntrinsicToken
const LiteralToken = age.LiteralToken
const LowercaseToken = age.LowercaseToken
const NegationToken = age.NegationToken
const NoteToken = age.NoteToken
const NumberToken = age.NumberToken
const QuantifiedToken = age.QuantifiedToken
const RuneToken = age.RuneToken
const SpaceToken = age.SpaceToken
const UppercaseToken = age.UppercaseToken

// UNIVERSAL CONSTRUCTORS

// AST

func Alternative(arguments ...any) AlternativeLike {
	// Initialize the possible arguments.
	var parts col.ListLike[PartLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case col.ListLike[PartLike]:
			parts = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the alternative constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var alternative = ast.Alternative().Make(parts)
	return alternative
}

func Character(arguments ...any) CharacterLike {
	// Initialize the possible arguments.
	var bounded BoundedLike
	var intrinsic string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case BoundedLike:
			bounded = actual
		case string:
			intrinsic = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the character constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var character CharacterLike
	switch {
	case fwk.IsDefined(bounded):
		character = ast.Character().Make(bounded)
	case fwk.IsDefined(intrinsic):
		character = ast.Character().Make(intrinsic)
	default:
		panic("The constructor for a character requires an argument.")
	}
	return character
}

func Cardinality(arguments ...any) CardinalityLike {
	// Initialize the possible arguments.
	var constrained ConstrainedLike
	var quantified string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ConstrainedLike:
			constrained = actual
		case string:
			quantified = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the cardinality constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var cardinality CardinalityLike
	switch {
	case fwk.IsDefined(constrained):
		cardinality = ast.Cardinality().Make(constrained)
	case fwk.IsDefined(quantified):
		cardinality = ast.Cardinality().Make(quantified)
	default:
		panic("The constructor for a cardinality requires an argument.")
	}
	return cardinality
}

func Constrained(arguments ...any) ConstrainedLike {
	// Initialize the possible arguments.
	var minimum MinimumLike
	var maximum MaximumLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case MinimumLike:
			minimum = actual
		case MaximumLike:
			maximum = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the constrained constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var constrained = ast.Constrained().Make(
		minimum,
		maximum,
	)
	return constrained
}

func Rule(arguments ...any) RuleLike {
	// Initialize the possible arguments.
	var comment string
	var uppercase string
	var expression ExpressionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case sts.HasPrefix(actual, "!>"):
				comment = actual
			default:
				uppercase = actual
			}
		case ExpressionLike:
			expression = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the rule constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var rule = ast.Rule().Make(
		comment,
		uppercase,
		expression,
	)
	return rule
}

func Element(arguments ...any) ElementLike {
	// Initialize the possible arguments.
	var grouped GroupedLike
	var filtered FilteredLike
	var bounded BoundedLike
	var intrinsic string
	var lowercase string
	var literal string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case GroupedLike:
			grouped = actual
		case FilteredLike:
			filtered = actual
		case BoundedLike:
			bounded = actual
		case string:
			switch {
			case matchesToken(IntrinsicToken, actual):
				intrinsic = actual
			case matchesToken(LowercaseToken, actual):
				lowercase = actual
			case matchesToken(LiteralToken, actual):
				literal = actual
			default:
				var message = fmt.Sprintf(
					"An invalid string was passed into the element constructor: %v\n",
					actual,
				)
				panic(message)
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the element constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var element ElementLike
	switch {
	case fwk.IsDefined(grouped):
		element = ast.Element().Make(grouped)
	case fwk.IsDefined(filtered):
		element = ast.Element().Make(filtered)
	case fwk.IsDefined(bounded):
		element = ast.Element().Make(bounded)
	case fwk.IsDefined(intrinsic):
		element = ast.Element().Make(intrinsic)
	case fwk.IsDefined(lowercase):
		element = ast.Element().Make(lowercase)
	case fwk.IsDefined(literal):
		element = ast.Element().Make(literal)
	default:
		panic("The constructor for an element requires an argument.")
	}
	return element
}

func Expression(arguments ...any) ExpressionLike {
	// Initialize the possible arguments.
	var inlined InlinedLike
	var multilined MultilinedLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case InlinedLike:
			inlined = actual
		case MultilinedLike:
			multilined = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the expression constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var expression ExpressionLike
	switch {
	case fwk.IsDefined(inlined):
		expression = ast.Expression().Make(inlined)
	case fwk.IsDefined(multilined):
		expression = ast.Expression().Make(multilined)
	default:
		panic("The constructor for an expression requires an argument.")
	}
	return expression
}

func Factor(arguments ...any) FactorLike {
	// Initialize the possible arguments.
	var predicate PredicateLike
	var cardinality CardinalityLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PredicateLike:
			predicate = actual
		case CardinalityLike:
			cardinality = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the factor constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var factor = ast.Factor().Make(
		predicate,
		cardinality,
	)
	return factor
}

func Filtered(arguments ...any) FilteredLike {
	// Initialize the possible arguments.
	var negation string
	var characters col.ListLike[CharacterLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			negation = actual
		case col.ListLike[CharacterLike]:
			characters = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the filtered constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var filtered = ast.Filtered().Make(
		negation,
		characters,
	)
	return filtered
}

func Initial(arguments ...any) InitialLike {
	// Initialize the possible arguments.
	var rune_ string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			rune_ = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the initial constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var initial = ast.Initial().Make(rune_)
	return initial
}

func Bounded(arguments ...any) BoundedLike {
	// Initialize the possible arguments.
	var initial InitialLike
	var extent ExtentLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case InitialLike:
			initial = actual
		case ExtentLike:
			extent = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the bounded constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var bounded = ast.Bounded().Make(
		initial,
		extent,
	)
	return bounded
}

func Header(arguments ...any) HeaderLike {
	// Initialize the possible arguments.
	var comment string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			comment = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the header constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var header = ast.Header().Make(comment)
	return header
}

func Identifier(arguments ...any) IdentifierLike {
	// Initialize the possible arguments.
	var lowercase string
	var uppercase string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			var runes = []rune(actual)
			switch {
			case uni.IsLower(runes[0]):
				lowercase = actual
			default:
				uppercase = actual
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the identifier constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var identifier IdentifierLike
	switch {
	case fwk.IsDefined(lowercase):
		identifier = ast.Identifier().Make(lowercase)
	case fwk.IsDefined(uppercase):
		identifier = ast.Identifier().Make(uppercase)
	default:
		panic("The constructor for an identifier requires an argument.")
	}
	return identifier
}

func Inlined(arguments ...any) InlinedLike {
	// Initialize the possible arguments.
	var factors col.ListLike[FactorLike]
	var note string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case col.ListLike[FactorLike]:
			factors = actual
		case string:
			note = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the inlined constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var inlined = ast.Inlined().Make(
		factors,
		note,
	)
	return inlined
}

func Extent(arguments ...any) ExtentLike {
	// Initialize the possible arguments.
	var rune_ string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			rune_ = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the extent constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var extent = ast.Extent().Make(rune_)
	return extent
}

func Lexigram(arguments ...any) LexigramLike {
	// Initialize the possible arguments.
	var comment string
	var lowercase string
	var pattern PatternLike
	var note string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case sts.HasPrefix(actual, "!>"):
				comment = actual
			case sts.HasPrefix(actual, "! "):
				note = actual
			default:
				lowercase = actual
			}
		case PatternLike:
			pattern = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the lexigram constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var lexigram = ast.Lexigram().Make(
		comment,
		lowercase,
		pattern,
		note,
	)
	return lexigram
}

func Line(arguments ...any) LineLike {
	// Initialize the possible arguments.
	var identifier IdentifierLike
	var note string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case IdentifierLike:
			identifier = actual
		case string:
			note = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the line constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var line = ast.Line().Make(
		identifier,
		note,
	)
	return line
}

func Maximum(arguments ...any) MaximumLike {
	// Initialize the possible arguments.
	var number string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			number = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the maximum constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var maximum = ast.Maximum().Make(number)
	return maximum
}

func Minimum(arguments ...any) MinimumLike {
	// Initialize the possible arguments.
	var number string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			number = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the minimum constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var minimum = ast.Minimum().Make(number)
	return minimum
}

func Multilined(arguments ...any) MultilinedLike {
	// Initialize the possible arguments.
	var lines col.ListLike[LineLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case col.ListLike[LineLike]:
			lines = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the multilined constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var multilined = ast.Multilined().Make(lines)
	return multilined
}

func Grouped(arguments ...any) GroupedLike {
	// Initialize the possible arguments.
	var pattern PatternLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PatternLike:
			pattern = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the grouped constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var grouped = ast.Grouped().Make(pattern)
	return grouped
}

func Part(arguments ...any) PartLike {
	// Initialize the possible arguments.
	var element ElementLike
	var cardinality CardinalityLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ElementLike:
			element = actual
		case CardinalityLike:
			cardinality = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the part constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var part = ast.Part().Make(
		element,
		cardinality,
	)
	return part
}

func Predicate(arguments ...any) PredicateLike {
	// Initialize the possible arguments.
	var lowercase string
	var uppercase string
	var intrinsic string
	var literal string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case matchesToken(LowercaseToken, actual):
				lowercase = actual
			case matchesToken(UppercaseToken, actual):
				uppercase = actual
			case matchesToken(IntrinsicToken, actual):
				intrinsic = actual
			case matchesToken(LiteralToken, actual):
				literal = actual
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the predicate constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var predicate PredicateLike
	switch {
	case fwk.IsDefined(lowercase):
		predicate = ast.Predicate().Make(lowercase)
	case fwk.IsDefined(uppercase):
		predicate = ast.Predicate().Make(uppercase)
	case fwk.IsDefined(intrinsic):
		predicate = ast.Predicate().Make(intrinsic)
	case fwk.IsDefined(literal):
		predicate = ast.Predicate().Make(literal)
	default:
		panic("The constructor for a predicate requires an argument.")
	}
	return predicate
}

func Syntax(arguments ...any) SyntaxLike {
	// Initialize the possible arguments.
	var headers col.ListLike[HeaderLike]
	var rules col.ListLike[RuleLike]
	var lexigrams col.ListLike[LexigramLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case col.ListLike[HeaderLike]:
			headers = actual
		case col.ListLike[RuleLike]:
			rules = actual
		case col.ListLike[LexigramLike]:
			lexigrams = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the syntax constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var syntax = ast.Syntax().Make(
		headers,
		rules,
		lexigrams,
	)
	return syntax
}

// Agents

func Formatter(arguments ...any) FormatterLike {
	if len(arguments) > 0 {
		panic("The formatter constructor does not take any arguments.")
	}
	var formatter = age.Formatter().Make()
	return formatter
}

func Generator(arguments ...any) GeneratorLike {
	if len(arguments) > 0 {
		panic("The generator constructor does not take any arguments.")
	}
	var generator = age.Generator().Make()
	return generator
}

func Parser(arguments ...any) ParserLike {
	if len(arguments) > 0 {
		panic("The parser constructor does not take any arguments.")
	}
	var parser = age.Parser().Make()
	return parser
}

func Validator(arguments ...any) ValidatorLike {
	if len(arguments) > 0 {
		panic("The validator constructor does not take any arguments.")
	}
	var validator = age.Validator().Make()
	return validator
}

func Pattern(arguments ...any) PatternLike {
	// Initialize the possible arguments.
	var parts col.ListLike[PartLike]
	var alternatives col.ListLike[AlternativeLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case col.ListLike[PartLike]:
			parts = actual
		case col.ListLike[AlternativeLike]:
			alternatives = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the pattern constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var pattern = ast.Pattern().Make(
		parts,
		alternatives,
	)
	return pattern
}

// Private

func matchesToken(type_ TokenType, value string) bool {
	var matches = age.Scanner().MatchToken(type_, value)
	return !matches.IsEmpty()
}
