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
	SelectiveLike   = ast.SelectiveLike
	SequentialLike  = ast.SequentialLike
	SupplementLike  = ast.SupplementLike
	SyntaxLike      = ast.SyntaxLike
	TextualLike     = ast.TextualLike
)

// Grammar

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ScannerLike   = gra.ScannerLike
	TokenType     = gra.TokenType
	ValidatorLike = gra.ValidatorLike
)

const (
	ErrorToken      = gra.ErrorToken
	CommentToken    = gra.CommentToken
	GlyphToken      = gra.GlyphToken
	IntrinsicToken  = gra.IntrinsicToken
	LiteralToken    = gra.LiteralToken
	LowercaseToken  = gra.LowercaseToken
	NegationToken   = gra.NegationToken
	NewlineToken    = gra.NewlineToken
	NoteToken       = gra.NoteToken
	NumberToken     = gra.NumberToken
	QuantifiedToken = gra.QuantifiedToken
	ReservedToken   = gra.ReservedToken
	SpaceToken      = gra.SpaceToken
	UppercaseToken  = gra.UppercaseToken
)

// UNIVERSAL CONSTRUCTORS

// AST

func Alternative(arguments ...any) AlternativeLike {
	// Initialize the possible arguments.
	var part PartLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PartLike:
			part = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the alternative constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var bar = "|"
	var alternative = ast.Alternative().Make(
		bar,
		part,
	)
	return alternative
}

func Bounded(arguments ...any) BoundedLike {
	// Initialize the possible arguments.
	var glyph string
	var extent ExtentLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			glyph = actual
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
		glyph,
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
	var left = "{"
	var right = "}"
	var constrained = ast.Constrained().Make(
		left,
		number,
		limit,
		right,
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
	var textual TextualLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case GroupedLike:
			grouped = actual
		case FilteredLike:
			filtered = actual
		case TextualLike:
			textual = actual
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
	case col.IsDefined(textual):
		element = ast.Element().Make(textual)
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
			case MatchesType(actual, CommentToken):
				comment = actual
			case MatchesType(actual, NoteToken):
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
	var colon = ":"
	var newlines = col.List[string]([]string{"\n", "\n"})
	var expression = ast.Expression().Make(
		comment,
		lowercase,
		colon,
		pattern,
		note,
		newlines,
	)
	return expression
}

func Extent(arguments ...any) ExtentLike {
	// Initialize the possible arguments.
	var glyph string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			glyph = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the extent constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var dotdot = ".."
	var extent = ast.Extent().Make(
		dotdot,
		glyph,
	)
	return extent
}

func Factor(arguments ...any) FactorLike {
	// Initialize the possible arguments.
	var predicate PredicateLike
	var literal string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PredicateLike:
			predicate = actual
		case string:
			literal = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the factor constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var factor FactorLike
	switch {
	case col.IsDefined(predicate):
		factor = ast.Factor().Make(predicate)
	case col.IsDefined(literal):
		factor = ast.Factor().Make(literal)
	default:
		panic("The constructor for a factor requires an argument.")
	}
	return factor
}

func Filtered(arguments ...any) FilteredLike {
	// Initialize the possible arguments.
	var negation string
	var characters abs.Sequential[CharacterLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			negation = actual
		case abs.Sequential[CharacterLike]:
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
	var left = "["
	var right = "]"
	var filtered = ast.Filtered().Make(
		left,
		negation,
		characters,
		right,
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
	var left = "("
	var right = ")"
	var grouped = ast.Grouped().Make(
		left,
		pattern,
		right,
	)
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
	var newline = "\n"
	var header = ast.Header().Make(
		comment,
		newline,
	)
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
			switch {
			case MatchesType(actual, LowercaseToken):
				lowercase = actual
			case MatchesType(actual, UppercaseToken):
				uppercase = actual
			default:
				var message = fmt.Sprintf(
					"An invalid string was passed into the identifier constructor: %v\n",
					actual,
				)
				panic(message)
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
	var factors abs.Sequential[FactorLike]
	var note string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[FactorLike]:
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
	var newline = "\n"
	var line = ast.Line().Make(
		newline,
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
	var dotdot = ".."
	var limit = ast.Limit().Make(
		dotdot,
		number,
	)
	return limit
}

func Multilined(arguments ...any) MultilinedLike {
	// Initialize the possible arguments.
	var lines abs.Sequential[LineLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[LineLike]:
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
	var part PartLike
	var supplement SupplementLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PartLike:
			part = actual
		case SupplementLike:
			supplement = actual
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
		part,
		supplement,
	)
	return pattern
}

func Predicate(arguments ...any) PredicateLike {
	// Initialize the possible arguments.
	var identifier IdentifierLike
	var cardinality CardinalityLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case IdentifierLike:
			identifier = actual
		case CardinalityLike:
			cardinality = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the predicate constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var predicate = ast.Predicate().Make(
		identifier,
		cardinality,
	)
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
			case MatchesType(actual, CommentToken):
				comment = actual
			case MatchesType(actual, UppercaseToken):
				uppercase = actual
			default:
				var message = fmt.Sprintf(
					"An invalid string was passed into the rule constructor: %v\n",
					actual,
				)
				panic(message)
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
	var colon = ":"
	var newlines = col.List[string]([]string{"\n", "\n"})
	var rule = ast.Rule().Make(
		comment,
		uppercase,
		colon,
		definition,
		newlines,
	)
	return rule
}

func Selective(arguments ...any) SelectiveLike {
	// Initialize the possible arguments.
	var alternatives abs.Sequential[AlternativeLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[AlternativeLike]:
			alternatives = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the selective constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var selective = ast.Selective().Make(alternatives)
	return selective
}

func Sequential(arguments ...any) SequentialLike {
	// Initialize the possible arguments.
	var parts abs.Sequential[PartLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[PartLike]:
			parts = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the sequential constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var sequential = ast.Sequential().Make(parts)
	return sequential
}

func Supplement(arguments ...any) SupplementLike {
	// Initialize the possible arguments.
	var sequential SequentialLike
	var selective SelectiveLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case SequentialLike:
			sequential = actual
		case SelectiveLike:
			selective = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the supplement constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var supplement SupplementLike
	switch {
	case col.IsDefined(sequential):
		supplement = ast.Supplement().Make(sequential)
	case col.IsDefined(selective):
		supplement = ast.Supplement().Make(selective)
	default:
		panic("The constructor for a supplement requires an argument.")
	}
	return supplement
}

func Syntax(arguments ...any) SyntaxLike {
	// Initialize the possible arguments.
	var headers abs.Sequential[HeaderLike]
	var rules abs.Sequential[RuleLike]
	var expressions abs.Sequential[ExpressionLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[HeaderLike]:
			headers = actual
		case abs.Sequential[RuleLike]:
			rules = actual
		case abs.Sequential[ExpressionLike]:
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

func Textual(arguments ...any) TextualLike {
	// Initialize the possible arguments.
	var intrinsic string
	var glyph string
	var literal string
	var lowercase string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case MatchesType(actual, IntrinsicToken):
				intrinsic = actual
			case MatchesType(actual, GlyphToken):
				glyph = actual
			case MatchesType(actual, LiteralToken):
				literal = actual
			case MatchesType(actual, LowercaseToken):
				lowercase = actual
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
	var text TextualLike
	switch {
	case col.IsDefined(intrinsic):
		text = Textual(intrinsic)
	case col.IsDefined(glyph):
		text = Textual(glyph)
	case col.IsDefined(literal):
		text = Textual(literal)
	case col.IsDefined(lowercase):
		text = Textual(lowercase)
	default:
		panic("The constructor for an string requires an argument.")
	}
	return text
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

// GLOBAL FUNCTIONS

// Grammar

func FormatSyntax(syntax SyntaxLike) string {
	var formatter = gra.Formatter().Make()
	var source = formatter.FormatSyntax(syntax)
	return source
}

func MatchesType(tokenValue string, tokenType TokenType) bool {
	var scannerClass = gra.Scanner()
	return scannerClass.MatchesType(tokenValue, tokenType)
}

func ParseSource(source string) SyntaxLike {
	var parser = gra.Parser().Make()
	var syntax = parser.ParseSource(source)
	return syntax
}

func ValidateSyntax(syntax SyntaxLike) {
	var validator = gra.Validator().Make()
	validator.ValidateSyntax(syntax)
}

// Generator

func GenerateFormatterClass(
	module string,
	syntax SyntaxLike,
) (
	implementation string,
) {
	var generator = gen.Formatter().Make()
	implementation = generator.GenerateFormatterClass(module, syntax)
	return implementation
}

func GenerateGrammarModel(
	module string,
	wiki string,
	syntax SyntaxLike,
) (
	implementation string,
) {
	var generator = gen.Grammar().Make()
	implementation = generator.GenerateGrammarModel(module, wiki, syntax)
	return implementation
}

func GenerateAstModel(
	module string,
	wiki string,
	syntax SyntaxLike,
) (
	implementation string,
) {
	var generator = gen.Ast().Make()
	implementation = generator.GenerateAstModel(module, wiki, syntax)
	return implementation
}

func GenerateParserClass(
	module string,
	syntax SyntaxLike,
) (
	implementation string,
) {
	var generator = gen.Parser().Make()
	implementation = generator.GenerateParserClass(module, syntax)
	return implementation
}

func GenerateScannerClass(
	module string,
	syntax SyntaxLike,
) (
	implementation string,
) {
	var generator = gen.Scanner().Make()
	implementation = generator.GenerateScannerClass(module, syntax)
	return implementation
}

func GenerateSyntaxNotation(
	syntax string,
	copyright string,
) (
	implementation string,
) {
	var generator = gen.Syntax().Make()
	implementation = generator.GenerateSyntaxNotation(syntax, copyright)
	return implementation
}

func GenerateTokenClass(
	module string,
	syntax SyntaxLike,
) (
	implementation string,
) {
	var generator = gen.Token().Make()
	implementation = generator.GenerateTokenClass(module, syntax)
	return implementation
}

func GenerateValidatorClass(
	module string,
	syntax SyntaxLike,
) (
	implementation string,
) {
	var generator = gen.Validator().Make()
	implementation = generator.GenerateValidatorClass(module, syntax)
	return implementation
}
