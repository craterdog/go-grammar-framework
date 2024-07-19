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
implementation files for each abstract class defined in the Packgra.go file.
*/
package module

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gen "github.com/craterdog/go-grammar-framework/v4/generator"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
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
	DefinitionLike  = ast.DefinitionLike
	ElementLike     = ast.ElementLike
	ExpressionLike  = ast.ExpressionLike
	ExtentLike      = ast.ExtentLike
	FactorLike      = ast.FactorLike
	FilteredLike    = ast.FilteredLike
	GroupedLike     = ast.GroupedLike
	HeaderLike      = ast.HeaderLike
	IdentifierLike  = ast.IdentifierLike
	InlinedLike     = ast.InlinedLike
	LimitLike       = ast.LimitLike
	LineLike        = ast.LineLike
	MultilinedLike  = ast.MultilinedLike
	PartLike        = ast.PartLike
	PatternLike     = ast.PatternLike
	PredicateLike   = ast.PredicateLike
	RuleLike        = ast.RuleLike
	StringLike      = ast.StringLike
	SyntaxLike      = ast.SyntaxLike
)

// Grammar

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ValidatorLike = gra.ValidatorLike
	TokenType     = gra.TokenType
)

const (
	ErrorToken      = gra.ErrorToken
	CommentToken    = gra.CommentToken
	DelimiterToken  = gra.DelimiterToken
	EofToken        = gra.EofToken
	EolToken        = gra.EolToken
	IntrinsicToken  = gra.IntrinsicToken
	LiteralToken    = gra.LiteralToken
	LowercaseToken  = gra.LowercaseToken
	NegationToken   = gra.NegationToken
	NoteToken       = gra.NoteToken
	NumberToken     = gra.NumberToken
	QuantifiedToken = gra.QuantifiedToken
	RuneToken       = gra.RuneToken
	SpaceToken      = gra.SpaceToken
	UppercaseToken  = gra.UppercaseToken
)

// Generator

type (
	GeneratorLike = gen.GeneratorLike
)

// UNIVERSAL CONSTRUCTORS

// AST

func Alternative(arguments ...any) AlternativeLike {
	// Initialize the possible arguments.
	var parts abs.ListLike[PartLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.ListLike[PartLike]:
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

func Bounded(arguments ...any) BoundedLike {
	// Initialize the possible arguments.
	var rune_ string
	var extent ExtentLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			rune_ = actual
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
		rune_,
		extent,
	)
	return bounded
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
	case col.IsDefined(constrained):
		cardinality = ast.Cardinality().Make(constrained)
	case col.IsDefined(quantified):
		cardinality = ast.Cardinality().Make(quantified)
	default:
		panic("The constructor for a cardinality requires an argument.")
	}
	return cardinality
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
	case col.IsDefined(bounded):
		character = ast.Character().Make(bounded)
	case col.IsDefined(intrinsic):
		character = ast.Character().Make(intrinsic)
	default:
		panic("The constructor for a character requires an argument.")
	}
	return character
}

func Constrained(arguments ...any) ConstrainedLike {
	// Initialize the possible arguments.
	var number string
	var limit LimitLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			number = actual
		case LimitLike:
			limit = actual
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
		number,
		limit,
	)
	return constrained
}

func Definition(arguments ...any) DefinitionLike {
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
				"An unknown argument type passed into the definition constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var definition DefinitionLike
	switch {
	case col.IsDefined(inlined):
		definition = ast.Definition().Make(inlined)
	case col.IsDefined(multilined):
		definition = ast.Definition().Make(multilined)
	default:
		panic("The constructor for a definition requires an argument.")
	}
	return definition
}

func Element(arguments ...any) ElementLike {
	// Initialize the possible arguments.
	var grouped GroupedLike
	var filtered FilteredLike
	var string_ StringLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case GroupedLike:
			grouped = actual
		case FilteredLike:
			filtered = actual
		case StringLike:
			string_ = actual
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
	case col.IsDefined(grouped):
		element = ast.Element().Make(grouped)
	case col.IsDefined(filtered):
		element = ast.Element().Make(filtered)
	case col.IsDefined(string_):
		element = ast.Element().Make(string_)
	default:
		panic("The constructor for an element requires an argument.")
	}
	return element
}

func Expression(arguments ...any) ExpressionLike {
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
				"An unknown argument type passed into the expression constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var expression = ast.Expression().Make(
		comment,
		lowercase,
		pattern,
		note,
	)
	return expression
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
	var characters abs.ListLike[CharacterLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			negation = actual
		case abs.ListLike[CharacterLike]:
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
	case col.IsDefined(lowercase):
		identifier = ast.Identifier().Make(lowercase)
	case col.IsDefined(uppercase):
		identifier = ast.Identifier().Make(uppercase)
	default:
		panic("The constructor for an identifier requires an argument.")
	}
	return identifier
}

func Inlined(arguments ...any) InlinedLike {
	// Initialize the possible arguments.
	var factors abs.ListLike[FactorLike]
	var note string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.ListLike[FactorLike]:
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

func Limit(arguments ...any) LimitLike {
	// Initialize the possible arguments.
	var number string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			number = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the limit constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var limit = ast.Limit().Make(number)
	return limit
}

func Multilined(arguments ...any) MultilinedLike {
	// Initialize the possible arguments.
	var lines abs.ListLike[LineLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.ListLike[LineLike]:
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

func Pattern(arguments ...any) PatternLike {
	// Initialize the possible arguments.
	var parts abs.ListLike[PartLike]
	var alternatives abs.ListLike[AlternativeLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.ListLike[PartLike]:
			parts = actual
		case abs.ListLike[AlternativeLike]:
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
	case col.IsDefined(lowercase):
		predicate = ast.Predicate().Make(lowercase)
	case col.IsDefined(uppercase):
		predicate = ast.Predicate().Make(uppercase)
	case col.IsDefined(intrinsic):
		predicate = ast.Predicate().Make(intrinsic)
	case col.IsDefined(literal):
		predicate = ast.Predicate().Make(literal)
	default:
		panic("The constructor for a predicate requires an argument.")
	}
	return predicate
}

func Rule(arguments ...any) RuleLike {
	// Initialize the possible arguments.
	var comment string
	var uppercase string
	var definition DefinitionLike

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
		case DefinitionLike:
			definition = actual
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
		definition,
	)
	return rule
}

func String(arguments ...any) StringLike {
	// Initialize the possible arguments.
	var rune_ string
	var literal string
	var lowercase string
	var intrinsic string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case matchesToken(RuneToken, actual):
				rune_ = actual
			case matchesToken(LiteralToken, actual):
				literal = actual
			case matchesToken(LowercaseToken, actual):
				lowercase = actual
			case matchesToken(IntrinsicToken, actual):
				intrinsic = actual
			default:
				var message = fmt.Sprintf(
					"An invalid string was passed into the string constructor: %v\n",
					actual,
				)
				panic(message)
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the string constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var string_ StringLike
	switch {
	case col.IsDefined(rune_):
		string_ = ast.String().Make(rune_)
	case col.IsDefined(literal):
		string_ = ast.String().Make(literal)
	case col.IsDefined(lowercase):
		string_ = ast.String().Make(lowercase)
	case col.IsDefined(intrinsic):
		string_ = ast.String().Make(intrinsic)
	default:
		panic("The constructor for an string requires an argument.")
	}
	return string_
}

func Syntax(arguments ...any) SyntaxLike {
	// Initialize the possible arguments.
	var headers abs.ListLike[HeaderLike]
	var rules abs.ListLike[RuleLike]
	var expressions abs.ListLike[ExpressionLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.ListLike[HeaderLike]:
			headers = actual
		case abs.ListLike[RuleLike]:
			rules = actual
		case abs.ListLike[ExpressionLike]:
			expressions = actual
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
		expressions,
	)
	return syntax
}

// Grammar

func Formatter(arguments ...any) FormatterLike {
	if len(arguments) > 0 {
		panic("The formatter constructor does not take any arguments.")
	}
	var formatter = gra.Formatter().Make()
	return formatter
}

func Parser(arguments ...any) ParserLike {
	if len(arguments) > 0 {
		panic("The parser constructor does not take any arguments.")
	}
	var parser = gra.Parser().Make()
	return parser
}

func Validator(arguments ...any) ValidatorLike {
	if len(arguments) > 0 {
		panic("The validator constructor does not take any arguments.")
	}
	var validator = gra.Validator().Make()
	return validator
}

// Generator

func Generator(arguments ...any) GeneratorLike {
	if len(arguments) > 0 {
		panic("The generator constructor does not take any arguments.")
	}
	var generator = gen.Generator().Make()
	return generator
}

// Private

func matchesToken(type_ TokenType, value string) bool {
	var matches = gra.Scanner().MatchToken(type_, value)
	return !matches.IsEmpty()
}
