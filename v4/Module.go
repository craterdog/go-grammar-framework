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
	CardinalityLike = ast.CardinalityLike
	CharacterLike   = ast.CharacterLike
	ConstrainedLike = ast.ConstrainedLike
	DefinitionLike  = ast.DefinitionLike
	ElementLike     = ast.ElementLike
	ExplicitLike    = ast.ExplicitLike
	ExpressionLike  = ast.ExpressionLike
	ExtentLike      = ast.ExtentLike
	FilterLike      = ast.FilterLike
	GroupLike       = ast.GroupLike
	IdentifierLike  = ast.IdentifierLike
	InlineLike      = ast.InlineLike
	LimitLike       = ast.LimitLike
	LineLike        = ast.LineLike
	MultilineLike   = ast.MultilineLike
	NoticeLike      = ast.NoticeLike
	OptionLike      = ast.OptionLike
	PatternLike     = ast.PatternLike
	QuantifiedLike  = ast.QuantifiedLike
	ReferenceLike   = ast.ReferenceLike
	RepetitionLike  = ast.RepetitionLike
	RuleLike        = ast.RuleLike
	SyntaxLike      = ast.SyntaxLike
	TermLike        = ast.TermLike
	TextLike        = ast.TextLike
)

// Grammar

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenType     = gra.TokenType
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
	Methodical    = gra.Methodical
)

const (
	ErrorToken     = gra.ErrorToken
	CommentToken   = gra.CommentToken
	DelimiterToken = gra.DelimiterToken
	ExcludedToken  = gra.ExcludedToken
	GlyphToken     = gra.GlyphToken
	IntrinsicToken = gra.IntrinsicToken
	LiteralToken   = gra.LiteralToken
	LowercaseToken = gra.LowercaseToken
	NewlineToken   = gra.NewlineToken
	NoteToken      = gra.NoteToken
	NumberToken    = gra.NumberToken
	OptionalToken  = gra.OptionalToken
	RepeatedToken  = gra.RepeatedToken
	SpaceToken     = gra.SpaceToken
	UppercaseToken = gra.UppercaseToken
)

// UNIVERSAL CONSTRUCTORS

// AST

func Alternative(arguments ...any) AlternativeLike {
	// Initialize the possible arguments.
	var option OptionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case OptionLike:
			option = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the alternative constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var alternative = ast.Alternative().Make(option)
	return alternative
}

func Cardinality(arguments ...any) CardinalityLike {
	// Initialize the possible arguments.
	var constrained ConstrainedLike
	var quantified QuantifiedLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ConstrainedLike:
			constrained = actual
		case QuantifiedLike:
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
	var explicit ExplicitLike
	var intrinsic string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ExplicitLike:
			explicit = actual
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
	case col.IsDefined(explicit):
		character = ast.Character().Make(explicit)
	case col.IsDefined(intrinsic):
		character = ast.Character().Make(intrinsic)
	default:
		panic("The constructor for a character requires an argument.")
	}
	return character
}

func Constrained(arguments ...any) ConstrainedLike {
	// Initialize the possible arguments.
	var optional string
	var repeated string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case MatchesType(actual, OptionalToken):
				optional = actual
			case MatchesType(actual, RepeatedToken):
				repeated = actual
			default:
				var message = fmt.Sprintf(
					"An invalid string was passed into the constrained constructor: %q\n",
					actual,
				)
				panic(message)
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the constrained constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var constrained ConstrainedLike
	switch {
	case col.IsDefined(optional):
		constrained = ast.Constrained().Make(optional)
	case col.IsDefined(repeated):
		constrained = ast.Constrained().Make(repeated)
	default:
		panic("The constructor for an constrained requires an argument.")
	}
	return constrained
}

func Definition(arguments ...any) DefinitionLike {
	// Initialize the possible arguments.
	var inline InlineLike
	var multiline MultilineLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case InlineLike:
			inline = actual
		case MultilineLike:
			multiline = actual
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
	case col.IsDefined(inline):
		definition = ast.Definition().Make(inline)
	case col.IsDefined(multiline):
		definition = ast.Definition().Make(multiline)
	default:
		panic("The constructor for a definition requires an argument.")
	}
	return definition
}

func Element(arguments ...any) ElementLike {
	// Initialize the possible arguments.
	var group GroupLike
	var filter FilterLike
	var text TextLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case GroupLike:
			group = actual
		case FilterLike:
			filter = actual
		case TextLike:
			text = actual
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
	case col.IsDefined(group):
		element = ast.Element().Make(group)
	case col.IsDefined(filter):
		element = ast.Element().Make(filter)
	case col.IsDefined(text):
		element = ast.Element().Make(text)
	default:
		panic("The constructor for an element requires an argument.")
	}
	return element
}

func Explicit(arguments ...any) ExplicitLike {
	// Initialize the possible arguments.
	var glyph string
	var extent ExtentLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case MatchesType(actual, GlyphToken):
				glyph = actual
			default:
				var message = fmt.Sprintf(
					"An unknown argument value was passed into the explicit constructor: %q\n",
					actual,
				)
				panic(message)
			}
		case ExtentLike:
			extent = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed was into the explicit constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var explicit = ast.Explicit().Make(
		glyph,
		extent,
	)
	return explicit
}

func Expression(arguments ...any) ExpressionLike {
	// Initialize the possible arguments.
	var lowercase string
	var pattern PatternLike
	var note string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case MatchesType(actual, LowercaseToken):
				lowercase = actual
			case MatchesType(actual, NoteToken):
				note = actual
			default:
				var message = fmt.Sprintf(
					"An unknown argument value was passed into the expression constructor: %q\n",
					actual,
				)
				panic(message)
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
	var newlines = col.List[string]([]string{"\n", "\n"})
	var expression = ast.Expression().Make(
		lowercase,
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
			switch {
			case MatchesType(actual, GlyphToken):
				glyph = actual
			default:
				var message = fmt.Sprintf(
					"An unknown argument value was passed into the extent constructor: %q\n",
					actual,
				)
				panic(message)
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the extent constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var extent = ast.Extent().Make(glyph)
	return extent
}

func Filter(arguments ...any) FilterLike {
	// Initialize the possible arguments.
	var excluded string
	var characters abs.Sequential[CharacterLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			excluded = actual
		case abs.Sequential[CharacterLike]:
			characters = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the filter constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var filter = ast.Filter().Make(
		excluded,
		characters,
	)
	return filter
}

func Group(arguments ...any) GroupLike {
	// Initialize the possible arguments.
	var pattern PatternLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PatternLike:
			pattern = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the group constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var group = ast.Group().Make(
		pattern,
	)
	return group
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
					"An invalid string was passed into the identifier constructor: %q\n",
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

func Inline(arguments ...any) InlineLike {
	// Initialize the possible arguments.
	var terms abs.Sequential[TermLike]
	var note string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[TermLike]:
			terms = actual
		case string:
			note = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the inline constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var inline = ast.Inline().Make(
		terms,
		note,
	)
	return inline
}

func Limit(arguments ...any) LimitLike {
	// Initialize the possible arguments.
	var number string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case MatchesType(actual, NumberToken):
				number = actual
			default:
				var message = fmt.Sprintf(
					"An unknown argument value was passed into the limit constructor: %q\n",
					actual,
				)
				panic(message)
			}
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
		identifier,
		note,
		newline,
	)
	return line
}

func Multiline(arguments ...any) MultilineLike {
	// Initialize the possible arguments.
	var lines abs.Sequential[LineLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[LineLike]:
			lines = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the multiline constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var newline = "\n"
	var multiline = ast.Multiline().Make(newline, lines)
	return multiline
}

func Notice(arguments ...any) NoticeLike {
	// Initialize the possible arguments.
	var comment string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case MatchesType(actual, CommentToken):
				comment = actual
			default:
				var message = fmt.Sprintf(
					"An unknown argument value was passed into the notice constructor: %q\n",
					actual,
				)
				panic(message)
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the notice constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var newline = "\n"
	var notice = ast.Notice().Make(comment, newline)
	return notice
}

func Option(arguments ...any) OptionLike {
	// Initialize the possible arguments.
	var repetitions = col.List[RepetitionLike]()

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case RepetitionLike:
			repetitions.AppendValue(actual)
		case abs.Sequential[RepetitionLike]:
			repetitions.AppendValues(actual)
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the option constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}
	if repetitions.IsEmpty() {
		panic("The option constructor requires at least one argument.")
	}

	// Call the constructor.
	var option = ast.Option().Make(repetitions)
	return option
}

func Pattern(arguments ...any) PatternLike {
	// Initialize the possible arguments.
	var option OptionLike
	var alternatives = col.List[AlternativeLike]()

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case OptionLike:
			option = actual
		case AlternativeLike:
			alternatives.AppendValue(actual)
		case abs.Sequential[AlternativeLike]:
			alternatives.AppendValues(actual)
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
		option,
		alternatives,
	)
	return pattern
}

func Quantified(arguments ...any) QuantifiedLike {
	// Initialize the possible arguments.
	var number string
	var limit LimitLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case MatchesType(actual, NumberToken):
				number = actual
			default:
				var message = fmt.Sprintf(
					"An unknown argument value was passed into the quantified constructor: %q\n",
					actual,
				)
				panic(message)
			}
		case LimitLike:
			limit = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed was into the quantified constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var quantified = ast.Quantified().Make(
		number,
		limit,
	)
	return quantified
}

func Reference(arguments ...any) ReferenceLike {
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
				"An unknown argument type passed into the reference constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var reference = ast.Reference().Make(
		identifier,
		cardinality,
	)
	return reference
}

func Repetition(arguments ...any) RepetitionLike {
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
				"An unknown argument type passed into the repetition constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var repetition = ast.Repetition().Make(
		element,
		cardinality,
	)
	return repetition
}

func Rule(arguments ...any) RuleLike {
	// Initialize the possible arguments.
	var uppercase string
	var definition DefinitionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case MatchesType(actual, UppercaseToken):
				uppercase = actual
			default:
				var message = fmt.Sprintf(
					"An invalid string was passed into the rule constructor: %q\n",
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
	var newlines = col.List[string]([]string{"\n", "\n"})
	var rule = ast.Rule().Make(
		uppercase,
		definition,
		newlines,
	)
	return rule
}

func Syntax(arguments ...any) SyntaxLike {
	// Initialize the possible arguments.
	var notice NoticeLike
	var comment1 string
	var rules abs.Sequential[RuleLike]
	var comment2 string
	var expressions abs.Sequential[ExpressionLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case NoticeLike:
			notice = actual
		case abs.Sequential[RuleLike]:
			rules = actual
		case abs.Sequential[ExpressionLike]:
			expressions = actual
		case string:
			if col.IsUndefined(comment1) {
				comment1 = actual
			} else {
				comment2 = actual
			}
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
		notice,
		comment1,
		rules,
		comment2,
		expressions,
	)
	return syntax
}

func Term(arguments ...any) TermLike {
	// Initialize the possible arguments.
	var reference ReferenceLike
	var literal string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ReferenceLike:
			reference = actual
		case string:
			switch {
			case MatchesType(actual, LiteralToken):
				literal = actual
			default:
				var message = fmt.Sprintf(
					"An unknown argument value was passed into the term constructor: %q\n",
					actual,
				)
				panic(message)
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed was into the term constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var term TermLike
	switch {
	case col.IsDefined(reference):
		term = ast.Term().Make(reference)
	case col.IsDefined(literal):
		term = ast.Term().Make(literal)
	default:
		panic("The constructor for a term requires an argument.")
	}
	return term
}

func Text(arguments ...any) TextLike {
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
					"An invalid string was passed into the text constructor: %q\n",
					actual,
				)
				panic(message)
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the text constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var text TextLike
	switch {
	case col.IsDefined(intrinsic):
		text = Text(intrinsic)
	case col.IsDefined(glyph):
		text = Text(glyph)
	case col.IsDefined(literal):
		text = Text(literal)
	case col.IsDefined(lowercase):
		text = Text(lowercase)
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

func Processor(arguments ...any) ProcessorLike {
	if len(arguments) > 0 {
		panic("The processor constructor does not take any arguments.")
	}
	var processor = gra.Processor().Make()
	return processor
}

func Validator(arguments ...any) ValidatorLike {
	if len(arguments) > 0 {
		panic("The validator constructor does not take any arguments.")
	}
	var validator = gra.Validator().Make()
	return validator
}

func Visitor(arguments ...any) VisitorLike {
	// Initialize the possible arguments.
	var processor Methodical

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case Methodical:
			processor = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the visitor constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var visitor = gra.Visitor().Make(processor)
	return visitor
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

func GenerateProcessorClass(
	module string,
	syntax SyntaxLike,
) (
	implementation string,
) {
	var generator = gen.Processor().Make()
	implementation = generator.GenerateProcessorClass(module, syntax)
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

func GenerateVisitorClass(
	module string,
	syntax SyntaxLike,
) (
	implementation string,
) {
	var generator = gen.Visitor().Make()
	implementation = generator.GenerateVisitorClass(module, syntax)
	return implementation
}
