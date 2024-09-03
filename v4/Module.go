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
	BracketLike     = ast.BracketLike
	CardinalityLike = ast.CardinalityLike
	CharacterLike   = ast.CharacterLike
	ConstraintLike  = ast.ConstraintLike
	CountLike       = ast.CountLike
	DefinitionLike  = ast.DefinitionLike
	ElementLike     = ast.ElementLike
	ExpressionLike  = ast.ExpressionLike
	FactorLike      = ast.FactorLike
	FilterLike      = ast.FilterLike
	GroupLike       = ast.GroupLike
	HeaderLike      = ast.HeaderLike
	IdentifierLike  = ast.IdentifierLike
	InlineLike      = ast.InlineLike
	LineLike        = ast.LineLike
	MultilineLike   = ast.MultilineLike
	PatternLike     = ast.PatternLike
	ReferenceLike   = ast.ReferenceLike
	RepetitionLike  = ast.RepetitionLike
	RuleLike        = ast.RuleLike
	SpecificLike    = ast.SpecificLike
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
	IntrinsicToken = gra.IntrinsicToken
	LiteralToken   = gra.LiteralToken
	LowercaseToken = gra.LowercaseToken
	NewlineToken   = gra.NewlineToken
	NoteToken      = gra.NoteToken
	NumberToken    = gra.NumberToken
	OptionalToken  = gra.OptionalToken
	RepeatedToken  = gra.RepeatedToken
	RunicToken     = gra.RunicToken
	SpaceToken     = gra.SpaceToken
	UppercaseToken = gra.UppercaseToken
)

// UNIVERSAL CONSTRUCTORS

// AST

func Alternative(arguments ...any) AlternativeLike {
	// Initialize the possible arguments.
	var repetition RepetitionLike
	var repetitions abs.Sequential[RepetitionLike]

	// Process the actual arguments.
	var list = col.List[RepetitionLike]()
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case RepetitionLike:
			repetition = actual
			list.AppendValue(repetition)
		case abs.Sequential[RepetitionLike]:
			repetitions = actual
			list.AppendValues(repetitions)
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the alternative constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}
	if list.IsEmpty() {
		panic("The alternative constructor requires at least one argument.")
	}
	repetitions = list

	// Call the constructor.
	var alternative = ast.Alternative().Make(
		repetitions,
	)
	return alternative
}

func Bracket(arguments ...any) BracketLike {
	// Initialize the possible arguments.
	var factors abs.Sequential[ast.FactorLike]
	var cardinality ast.CardinalityLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[FactorLike]:
			factors = actual
		case ast.CardinalityLike:
			cardinality = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the bracket constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var bracket = ast.Bracket().Make(
		factors,
		cardinality,
	)
	return bracket
}

func Cardinality(arguments ...any) CardinalityLike {
	// Initialize the possible arguments.
	var constraint ConstraintLike
	var count CountLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ConstraintLike:
			constraint = actual
		case CountLike:
			count = actual
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
	case col.IsDefined(constraint):
		cardinality = ast.Cardinality().Make(constraint)
	case col.IsDefined(count):
		cardinality = ast.Cardinality().Make(count)
	default:
		panic("The constructor for a cardinality requires an argument.")
	}
	return cardinality
}

func Character(arguments ...any) CharacterLike {
	// Initialize the possible arguments.
	var specific SpecificLike
	var intrinsic string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case SpecificLike:
			specific = actual
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
	case col.IsDefined(specific):
		character = ast.Character().Make(specific)
	case col.IsDefined(intrinsic):
		character = ast.Character().Make(intrinsic)
	default:
		panic("The constructor for a character requires an argument.")
	}
	return character
}

func Constraint(arguments ...any) ConstraintLike {
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
					"An invalid string was passed into the constraint constructor: %q\n",
					actual,
				)
				panic(message)
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the constraint constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var constraint ConstraintLike
	switch {
	case col.IsDefined(optional):
		constraint = ast.Constraint().Make(optional)
	case col.IsDefined(repeated):
		constraint = ast.Constraint().Make(repeated)
	default:
		panic("The constructor for an constraint requires an argument.")
	}
	return constraint
}

func Count(arguments ...any) CountLike {
	// Initialize the possible arguments.
	var numbers = col.List[string]()

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			numbers.AppendValue(actual)
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the count constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var count = ast.Count().Make(numbers)
	return count
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
	var newlines = col.List[string]([]string{"\n", "\n"})
	var expression = ast.Expression().Make(
		comment,
		lowercase,
		pattern,
		note,
		newlines,
	)
	return expression
}

func Factor(arguments ...any) FactorLike {
	// Initialize the possible arguments.
	var reference ReferenceLike
	var literal string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ReferenceLike:
			reference = actual
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
	case col.IsDefined(reference):
		factor = ast.Factor().Make(reference)
	case col.IsDefined(literal):
		factor = ast.Factor().Make(literal)
	default:
		panic("The constructor for a factor requires an argument.")
	}
	return factor
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
	var multiline = ast.Multiline().Make(lines)
	return multiline
}

func Pattern(arguments ...any) PatternLike {
	// Initialize the possible arguments.
	var alternatives = col.List[AlternativeLike]()

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
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
	if alternatives.IsEmpty() {
		panic("The pattern constructor requires at least one argument.")
	}

	// Call the constructor.
	var pattern = ast.Pattern().Make(
		alternatives,
	)
	return pattern
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
		comment,
		uppercase,
		definition,
		newlines,
	)
	return rule
}

func Specific(arguments ...any) SpecificLike {
	// Initialize the possible arguments.
	var runics = col.List[string]()

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			runics.AppendValue(actual)
		default:
			var message = fmt.Sprintf(
				"An unknown argument type passed into the specific constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var specific = ast.Specific().Make(runics)
	return specific
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

func Text(arguments ...any) TextLike {
	// Initialize the possible arguments.
	var intrinsic string
	var runic string
	var literal string
	var lowercase string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			switch {
			case MatchesType(actual, IntrinsicToken):
				intrinsic = actual
			case MatchesType(actual, RunicToken):
				runic = actual
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
	case col.IsDefined(runic):
		text = Text(runic)
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
