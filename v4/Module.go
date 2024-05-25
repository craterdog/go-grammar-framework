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
	col "github.com/craterdog/go-collection-framework/v4/collection"
	age "github.com/craterdog/go-grammar-framework/v4/agent"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	str "strings"
)

// TYPE ALIASES

// AST

type (
	AlternativeLike = ast.AlternativeLike
	AtomLike        = ast.AtomLike
	CardinalityLike = ast.CardinalityLike
	ConstraintLike  = ast.ConstraintLike
	DefinitionLike  = ast.DefinitionLike
	ElementLike     = ast.ElementLike
	ExpressionLike  = ast.ExpressionLike
	FactorLike      = ast.FactorLike
	FilterLike      = ast.FilterLike
	GlyphLike       = ast.GlyphLike
	HeaderLike      = ast.HeaderLike
	InlineLike      = ast.InlineLike
	LineLike        = ast.LineLike
	MultilineLike   = ast.MultilineLike
	PrecedenceLike  = ast.PrecedenceLike
	PredicateLike   = ast.PredicateLike
	SyntaxLike      = ast.SyntaxLike
)

// Agents

type (
	FormatterLike = age.FormatterLike
	GeneratorLike = age.GeneratorLike
	ParserLike    = age.ParserLike
	ValidatorLike = age.ValidatorLike
)

// UNIVERSAL CONSTRUCTORS

// AST

func Alternative(arguments ...any) AlternativeLike {
	// Initialize the possible arguments.
	var factors col.ListLike[FactorLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case col.ListLike[FactorLike]:
			factors = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the alternative constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var alternative = ast.Alternative().MakeWithFactors(factors)
	return alternative
}

func Atom(arguments ...any) AtomLike {
	// Initialize the possible arguments.
	var glyph GlyphLike
	var intrinsic string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case GlyphLike:
			glyph = actual
		case string:
			intrinsic = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the atom constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var atom AtomLike
	switch {
	case glyph != nil:
		atom = ast.Atom().MakeWithGlyph(glyph)
	case len(intrinsic) > 0:
		atom = ast.Atom().MakeWithIntrinsic(intrinsic)
	default:
		panic("The constructor for an atom requires an argument.")
	}
	return atom
}

func Cardinality(arguments ...any) CardinalityLike {
	// Initialize the possible arguments.
	var constraint ConstraintLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ConstraintLike:
			constraint = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the cardinality constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var cardinality = ast.Cardinality().MakeWithConstraint(constraint)
	return cardinality
}

func Constraint(arguments ...any) ConstraintLike {
	// Initialize the possible arguments.
	var first string
	var last string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			if len(first) == 0 {
				first = actual
			} else {
				last = actual
			}
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the constraint constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var constraint = ast.Constraint().MakeWithAttributes(
		first,
		last,
	)
	return constraint
}

func Definition(arguments ...any) DefinitionLike {
	// Initialize the possible arguments.
	var comment string
	var name string
	var expression ExpressionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			if len(arguments) == 3 && len(comment) == 0 {
				comment = actual
			} else {
				name = actual
			}
		case ExpressionLike:
			expression = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the definition constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var definition = ast.Definition().MakeWithAttributes(
		comment,
		name,
		expression,
	)
	return definition
}

func Element(arguments ...any) ElementLike {
	// Initialize the possible arguments.
	var literal string
	var name string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			if str.HasPrefix(actual, `"`) {
				literal = actual
			} else {
				name = actual
			}
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the element constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var element ElementLike
	switch {
	case len(literal) > 0:
		element = ast.Element().MakeWithLiteral(literal)
	case len(name) > 0:
		element = ast.Element().MakeWithName(name)
	default:
		panic("The constructor for an element requires an argument.")
	}
	return element
}

func Expression(arguments ...any) ExpressionLike {
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
				"Unknown argument type passed into the expression constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var expression ExpressionLike
	switch {
	case inline != nil:
		expression = ast.Expression().MakeWithInline(inline)
	case multiline != nil:
		expression = ast.Expression().MakeWithMultiline(multiline)
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
				"Unknown argument type passed into the factor constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var factor = ast.Factor().MakeWithAttributes(
		predicate,
		cardinality,
	)
	return factor
}

func Filter(arguments ...any) FilterLike {
	// Initialize the possible arguments.
	var inverted bool
	var atoms col.ListLike[AtomLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case bool:
			inverted = actual
		case col.ListLike[AtomLike]:
			atoms = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the filter constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var filter = ast.Filter().MakeWithAttributes(
		inverted,
		atoms,
	)
	return filter
}

func Glyph(arguments ...any) GlyphLike {
	// Initialize the possible arguments.
	var first string
	var last string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			if len(first) == 0 {
				first = actual
			} else {
				last = actual
			}
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the glyph constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var glyph = ast.Glyph().MakeWithAttributes(
		first,
		last,
	)
	return glyph
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
				"Unknown argument type passed into the header constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var header = ast.Header().MakeWithComment(comment)
	return header
}

func Inline(arguments ...any) InlineLike {
	// Initialize the possible arguments.
	var alternatives col.ListLike[AlternativeLike]
	var note string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case col.ListLike[AlternativeLike]:
			alternatives = actual
		case string:
			note = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the inline constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var inline = ast.Inline().MakeWithAttributes(
		alternatives,
		note,
	)
	return inline
}

func Line(arguments ...any) LineLike {
	// Initialize the possible arguments.
	var alternative AlternativeLike
	var note string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case AlternativeLike:
			alternative = actual
		case string:
			note = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the line constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var line = ast.Line().MakeWithAttributes(
		alternative,
		note,
	)
	return line
}

func Multiline(arguments ...any) MultilineLike {
	// Initialize the possible arguments.
	var lines col.ListLike[LineLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case col.ListLike[LineLike]:
			lines = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the multiline constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var multiline = ast.Multiline().MakeWithLines(lines)
	return multiline
}

func Precedence(arguments ...any) PrecedenceLike {
	// Initialize the possible arguments.
	var expression ExpressionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ExpressionLike:
			expression = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the precedence constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var precedence = ast.Precedence().MakeWithExpression(expression)
	return precedence
}

func Predicate(arguments ...any) PredicateLike {
	// Initialize the possible arguments.
	var atom AtomLike
	var element ElementLike
	var filter FilterLike
	var precedence PrecedenceLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case AtomLike:
			atom = actual
		case ElementLike:
			element = actual
		case FilterLike:
			filter = actual
		case PrecedenceLike:
			precedence = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the predicate constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var predicate PredicateLike
	switch {
	case atom != nil:
		predicate = ast.Predicate().MakeWithAtom(atom)
	case element != nil:
		predicate = ast.Predicate().MakeWithElement(element)
	case filter != nil:
		predicate = ast.Predicate().MakeWithFilter(filter)
	case precedence != nil:
		predicate = ast.Predicate().MakeWithPrecedence(precedence)
	default:
		panic("The constructor for a predicate requires an argument.")
	}
	return predicate
}

func Syntax(arguments ...any) SyntaxLike {
	// Initialize the possible arguments.
	var headers col.ListLike[HeaderLike]
	var definitions col.ListLike[DefinitionLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case col.ListLike[HeaderLike]:
			headers = actual
		case col.ListLike[DefinitionLike]:
			definitions = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the syntax constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var syntax = ast.Syntax().MakeWithAttributes(
		headers,
		definitions,
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
