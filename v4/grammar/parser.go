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
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	sts "strings"
)

// CLASS ACCESS

// Reference

var parserClass = &parserClass_{
	// Initialize the class constants.
	queueSize_: 16,
	stackSize_: 4,
}

// Function

func Parser() ParserClassLike {
	return parserClass
}

// CLASS METHODS

// Target

type parserClass_ struct {
	// Define the class constants.
	queueSize_ uint
	stackSize_ uint
}

// Constructors

func (c *parserClass_) Make() ParserLike {
	return &parser_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type parser_ struct {
	// Define the instance attributes.
	class_  ParserClassLike
	source_ string                   // The original source code.
	tokens_ abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Attributes

func (v *parser_) GetClass() ParserClassLike {
	return v.class_
}

// Public

func (v *parser_) ParseSource(source string) ast.SyntaxLike {
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserClass.queueSize_)
	v.next_ = col.Stack[TokenLike](parserClass.stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse the syntax.
	var syntax, token, ok = v.parseSyntax()
	if !ok {
		var message = v.formatError(token, "Syntax")
		panic(message)
	}

	// Found the syntax.
	return syntax
}

// Private

const unlimited = 4294967295 // Default to a reasonable value.

func (v *parser_) parseAlternative() (
	alternative ast.AlternativeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a "|" delimiter.
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		// This is not a alternative rule.
		return alternative, token, false
	}

	// Attempt to parse a option rule.
	var option ast.OptionLike
	option, token, ok = v.parseOption()
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Alternative")
		panic(message)
	}

	// Found a alternative rule.
	alternative = ast.Alternative().Make(option)
	return alternative, token, true
}

func (v *parser_) parseCardinality() (
	cardinality ast.CardinalityLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a constrained rule.
	var constrained ast.ConstrainedLike
	constrained, token, ok = v.parseConstrained()
	if ok {
		// Found a constrained cardinality.
		cardinality = ast.Cardinality().Make(constrained)
		return cardinality, token, true
	}

	// Attempt to parse a quantified rule.
	var quantified ast.QuantifiedLike
	quantified, token, ok = v.parseQuantified()
	if ok {
		// Found a quantified cardinality.
		cardinality = ast.Cardinality().Make(quantified)
		return cardinality, token, true
	}

	// This is not a cardinality rule.
	return cardinality, token, false
}

func (v *parser_) parseCharacter() (
	character ast.CharacterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a explicit rule.
	var explicit ast.ExplicitLike
	explicit, token, ok = v.parseExplicit()
	if ok {
		// Found a explicit character.
		character = ast.Character().Make(explicit)
		return character, token, true
	}

	// Attempt to parse a intrinsic token.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken)
	if ok {
		// Found a intrinsic character.
		character = ast.Character().Make(intrinsic)
		return character, token, true
	}

	// This is not a character rule.
	return character, token, false
}

func (v *parser_) parseConstrained() (
	constrained ast.ConstrainedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a optional token.
	var optional string
	optional, token, ok = v.parseToken(OptionalToken)
	if ok {
		// Found a optional constrained.
		constrained = ast.Constrained().Make(optional)
		return constrained, token, true
	}

	// Attempt to parse a repeated token.
	var repeated string
	repeated, token, ok = v.parseToken(RepeatedToken)
	if ok {
		// Found a repeated constrained.
		constrained = ast.Constrained().Make(repeated)
		return constrained, token, true
	}

	// This is not a constrained rule.
	return constrained, token, false
}

func (v *parser_) parseDefinition() (
	definition ast.DefinitionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a multiline rule.
	var multiline ast.MultilineLike
	multiline, token, ok = v.parseMultiline()
	if ok {
		// Found a multiline definition.
		definition = ast.Definition().Make(multiline)
		return definition, token, true
	}

	// Attempt to parse a inline rule.
	var inline ast.InlineLike
	inline, token, ok = v.parseInline()
	if ok {
		// Found a inline definition.
		definition = ast.Definition().Make(inline)
		return definition, token, true
	}

	// This is not a definition rule.
	return definition, token, false
}

func (v *parser_) parseElement() (
	element ast.ElementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a group rule.
	var group ast.GroupLike
	group, token, ok = v.parseGroup()
	if ok {
		// Found a group element.
		element = ast.Element().Make(group)
		return element, token, true
	}

	// Attempt to parse a filter rule.
	var filter ast.FilterLike
	filter, token, ok = v.parseFilter()
	if ok {
		// Found a filter element.
		element = ast.Element().Make(filter)
		return element, token, true
	}

	// Attempt to parse a text rule.
	var text ast.TextLike
	text, token, ok = v.parseText()
	if ok {
		// Found a text element.
		element = ast.Element().Make(text)
		return element, token, true
	}

	// This is not a element rule.
	return element, token, false
}

func (v *parser_) parseExplicit() (
	explicit ast.ExplicitLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if !ok {
		// This is not a explicit rule.
		return explicit, token, false
	}

	// Attempt to parse an optional extent rule.
	var optionalExtent ast.ExtentLike
	optionalExtent, _, _ = v.parseExtent()

	// Found a explicit rule.
	explicit = ast.Explicit().Make(
		glyph,
		optionalExtent,
	)
	return explicit, token, true
}

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a lowercase token.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if !ok {
		// This is not a expression rule.
		return expression, token, false
	}

	// Attempt to parse a ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Expression")
		panic(message)
	}

	// Attempt to parse a pattern rule.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Expression")
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, _, _ = v.parseToken(NoteToken)

	// Attempt to parse 1 to unlimited newline tokens.
	var newlines = col.List[string]()
loop:
	for i := 0; i < unlimited; i++ {
		var newline string
		newline, token, ok = v.parseToken(NewlineToken)
		if !ok {
			switch {
			case i < 1:
				var message = v.formatError(token, "Expression")
				message += "Too few newline tokens found."
				panic(message)
			case i > unlimited:
				var message = v.formatError(token, "Expression")
				message += "Too many newline tokens found."
				panic(message)
			default:
				break loop
			}
		}
		newlines.AppendValue(newline)
	}

	// Found a expression rule.
	expression = ast.Expression().Make(
		lowercase,
		pattern,
		optionalNote,
		newlines,
	)
	return expression, token, true
}

func (v *parser_) parseExtent() (
	extent ast.ExtentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		// This is not a extent rule.
		return extent, token, false
	}

	// Attempt to parse a glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Extent")
		panic(message)
	}

	// Found a extent rule.
	extent = ast.Extent().Make(glyph)
	return extent, token, true
}

func (v *parser_) parseFilter() (
	filter ast.FilterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an optional excluded token.
	var optionalExcluded string
	optionalExcluded, _, _ = v.parseToken(ExcludedToken)

	// Attempt to parse a "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if col.IsDefined(optionalExcluded) {
			// Found a syntax error.
			var message = v.formatError(token, "Filter")
			panic(message)
		}
		// This is not a filter element.
		return filter, token, false
	}

	// Attempt to parse 1 to unlimited character rules.
	var characters = col.List[ast.CharacterLike]()
loop:
	for i := 0; i < unlimited; i++ {
		var character ast.CharacterLike
		character, token, ok = v.parseCharacter()
		if !ok {
			switch {
			case i < 1:
				var message = v.formatError(token, "Filter")
				message += "Too few character rules found."
				panic(message)
			case i > unlimited:
				var message = v.formatError(token, "Filter")
				message += "Too many character rules found."
				panic(message)
			default:
				break loop
			}
		}
		characters.AppendValue(character)
	}

	// Attempt to parse a "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Filter")
		panic(message)
	}

	// Found a filter rule.
	filter = ast.Filter().Make(
		optionalExcluded,
		characters,
	)
	return filter, token, true
}

func (v *parser_) parseGroup() (
	group ast.GroupLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		// This is not a group rule.
		return group, token, false
	}

	// Attempt to parse a pattern rule.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Group")
		panic(message)
	}

	// Attempt to parse a ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Group")
		panic(message)
	}

	// Found a group rule.
	group = ast.Group().Make(pattern)
	return group, token, true
}

func (v *parser_) parseIdentifier() (
	identifier ast.IdentifierLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a lowercase token.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if ok {
		// Found a lowercase identifier.
		identifier = ast.Identifier().Make(lowercase)
		return identifier, token, true
	}

	// Attempt to parse a uppercase token.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if ok {
		// Found a uppercase identifier.
		identifier = ast.Identifier().Make(uppercase)
		return identifier, token, true
	}

	// This is not a identifier rule.
	return identifier, token, false
}

func (v *parser_) parseInline() (
	inline ast.InlineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse 1 to unlimited term rules.
	var terms = col.List[ast.TermLike]()
loop:
	for i := 0; i < unlimited; i++ {
		var term ast.TermLike
		term, token, ok = v.parseTerm()
		if !ok {
			switch {
			case i < 1:
				var message = v.formatError(token, "Inline")
				message += "Too few term rules found."
				panic(message)
			case i > unlimited:
				var message = v.formatError(token, "Inline")
				message += "Too many term rules found."
				panic(message)
			default:
				break loop
			}
		}
		terms.AppendValue(term)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, _, _ = v.parseToken(NoteToken)

	// Found a inline rule.
	inline = ast.Inline().Make(
		terms,
		optionalNote,
	)
	return inline, token, true
}

func (v *parser_) parseLimit() (
	limit ast.LimitLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		// This is not a limit rule.
		return limit, token, false
	}

	// Attempt to parse an optional number token.
	var optionalNumber string
	optionalNumber, _, _ = v.parseToken(NumberToken)

	// Found a limit rule.
	limit = ast.Limit().Make(optionalNumber)
	return limit, token, true
}

func (v *parser_) parseLine() (
	line ast.LineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a "-" delimiter.
	_, token, ok = v.parseDelimiter("-")
	if !ok {
		// This is not a line rule.
		return line, token, false
	}

	// Attempt to parse a identifier rule.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Line")
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, _, _ = v.parseToken(NoteToken)

	// Attempt to parse a newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Line")
		panic(message)
	}

	// Found a line rule.
	line = ast.Line().Make(
		identifier,
		optionalNote,
		newline,
	)
	return line, token, true
}

func (v *parser_) parseMultiline() (
	multiline ast.MultilineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		// This is not a multiline rule.
		return multiline, token, false
	}

	// Attempt to parse 1 to unlimited line rules.
	var lines = col.List[ast.LineLike]()
loop:
	for i := 0; i < unlimited; i++ {
		var line ast.LineLike
		line, token, ok = v.parseLine()
		if !ok {
			switch {
			case i < 1:
				var message = v.formatError(token, "Multiline")
				message += "Too few line rules found."
				panic(message)
			case i > unlimited:
				var message = v.formatError(token, "Multiline")
				message += "Too many line rules found."
				panic(message)
			default:
				break loop
			}
		}
		lines.AppendValue(line)
	}

	// Found a multiline rule.
	multiline = ast.Multiline().Make(
		newline,
		lines,
	)
	return multiline, token, true
}

func (v *parser_) parseNotice() (
	notice ast.NoticeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		// This is not a notice rule.
		return notice, token, false
	}

	// Attempt to parse a newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Notice")
		panic(message)
	}

	// Found a notice rule.
	notice = ast.Notice().Make(
		comment,
		newline,
	)
	return notice, token, true
}

func (v *parser_) parseOption() (
	option ast.OptionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse 1 to unlimited repetition rules.
	var repetitions = col.List[ast.RepetitionLike]()
loop:
	for i := 0; i < unlimited; i++ {
		var repetition ast.RepetitionLike
		repetition, token, ok = v.parseRepetition()
		if !ok {
			switch {
			case i < 1:
				var message = v.formatError(token, "Option")
				message += "Too few repetition rules found."
				panic(message)
			case i > unlimited:
				var message = v.formatError(token, "Option")
				message += "Too many repetition rules found."
				panic(message)
			default:
				break loop
			}
		}
		repetitions.AppendValue(repetition)
	}

	// Found a option rule.
	option = ast.Option().Make(repetitions)
	return option, token, true
}

func (v *parser_) parsePattern() (
	pattern ast.PatternLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a option rule.
	var option ast.OptionLike
	option, token, ok = v.parseOption()
	if !ok {
		// This is not a pattern rule.
		return pattern, token, false
	}

	// Attempt to parse 0 to unlimited alternative rules.
	var alternatives = col.List[ast.AlternativeLike]()
loop:
	for i := 0; i < unlimited; i++ {
		var alternative ast.AlternativeLike
		alternative, token, ok = v.parseAlternative()
		if !ok {
			switch {
			case i < 0:
				var message = v.formatError(token, "Pattern")
				message += "Too few alternative rules found."
				panic(message)
			case i > unlimited:
				var message = v.formatError(token, "Pattern")
				message += "Too many alternative rules found."
				panic(message)
			default:
				break loop
			}
		}
		alternatives.AppendValue(alternative)
	}

	// Found a pattern rule.
	pattern = ast.Pattern().Make(
		option,
		alternatives,
	)
	return pattern, token, true
}

func (v *parser_) parseQuantified() (
	quantified ast.QuantifiedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		// This is not a quantified rule.
		return quantified, token, false
	}

	// Attempt to parse a number token.
	var number string
	number, token, ok = v.parseToken(NumberToken)
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Quantified")
		panic(message)
	}

	// Attempt to parse an optional limit rule.
	var optionalLimit ast.LimitLike
	optionalLimit, _, _ = v.parseLimit()

	// Attempt to parse a "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Quantified")
		panic(message)
	}

	// Found a quantified rule.
	quantified = ast.Quantified().Make(
		number,
		optionalLimit,
	)
	return quantified, token, true
}

func (v *parser_) parseReference() (
	reference ast.ReferenceLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a identifier rule.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	if !ok {
		// This is not a reference rule.
		return reference, token, false
	}

	// Attempt to parse an optional cardinality rule.
	var optionalCardinality ast.CardinalityLike
	optionalCardinality, _, _ = v.parseCardinality()

	// Found a reference rule.
	reference = ast.Reference().Make(
		identifier,
		optionalCardinality,
	)
	return reference, token, true
}

func (v *parser_) parseRepetition() (
	repetition ast.RepetitionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a element rule.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	if !ok {
		// This is not a repetition rule.
		return repetition, token, false
	}

	// Attempt to parse an optional cardinality rule.
	var optionalCardinality ast.CardinalityLike
	optionalCardinality, _, _ = v.parseCardinality()

	// Found a repetition rule.
	repetition = ast.Repetition().Make(
		element,
		optionalCardinality,
	)
	return repetition, token, true
}

func (v *parser_) parseRule() (
	rule ast.RuleLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a uppercase token.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if !ok {
		// This is not a rule rule.
		return rule, token, false
	}

	// Attempt to parse a ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Rule")
		panic(message)
	}

	// Attempt to parse a definition rule.
	var definition ast.DefinitionLike
	definition, token, ok = v.parseDefinition()
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Rule")
		panic(message)
	}

	// Attempt to parse 1 to unlimited newline tokens.
	var newlines = col.List[string]()
loop:
	for i := 0; i < unlimited; i++ {
		var newline string
		newline, token, ok = v.parseToken(NewlineToken)
		if !ok {
			switch {
			case i < 1:
				var message = v.formatError(token, "Rule")
				message += "Too few newline tokens found."
				panic(message)
			case i > unlimited:
				var message = v.formatError(token, "Rule")
				message += "Too many newline tokens found."
				panic(message)
			default:
				break loop
			}
		}
		newlines.AppendValue(newline)
	}

	// Found a rule rule.
	rule = ast.Rule().Make(
		uppercase,
		definition,
		newlines,
	)
	return rule, token, true
}

func (v *parser_) parseSyntax() (
	syntax ast.SyntaxLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a notice rule.
	var notice ast.NoticeLike
	notice, token, ok = v.parseNotice()
	if !ok {
		// This is not a syntax rule.
		return syntax, token, false
	}

	// Attempt to parse a comment token.
	var comment1 string
	comment1, token, ok = v.parseToken(CommentToken)
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Syntax")
		panic(message)
	}

	// Attempt to parse 1 to unlimited rule rules.
	var rules = col.List[ast.RuleLike]()
loop1:
	for i := 0; i < unlimited; i++ {
		var rule ast.RuleLike
		rule, token, ok = v.parseRule()
		if !ok {
			switch {
			case i < 1:
				var message = v.formatError(token, "Syntax")
				message += "Too few rule rules found."
				panic(message)
			case i > unlimited:
				var message = v.formatError(token, "Syntax")
				message += "Too many rule rules found."
				panic(message)
			default:
				break loop1
			}
		}
		rules.AppendValue(rule)
	}

	// Attempt to parse a comment token.
	var comment2 string
	comment2, token, ok = v.parseToken(CommentToken)
	if !ok {
		// Found a syntax error.
		var message = v.formatError(token, "Syntax")
		panic(message)
	}

	// Attempt to parse 1 to unlimited expression rules.
	var expressions = col.List[ast.ExpressionLike]()
loop2:
	for i := 0; i < unlimited; i++ {
		var expression ast.ExpressionLike
		expression, token, ok = v.parseExpression()
		if !ok {
			switch {
			case i < 1:
				var message = v.formatError(token, "Syntax")
				message += "Too few expression rules found."
				panic(message)
			case i > unlimited:
				var message = v.formatError(token, "Syntax")
				message += "Too many expression rules found."
				panic(message)
			default:
				break loop2
			}
		}
		expressions.AppendValue(expression)
	}

	// Found a syntax rule.
	syntax = ast.Syntax().Make(
		notice,
		comment1,
		rules,
		comment2,
		expressions,
	)
	return syntax, token, true
}

func (v *parser_) parseTerm() (
	term ast.TermLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a reference rule.
	var reference ast.ReferenceLike
	reference, token, ok = v.parseReference()
	if ok {
		// Found a reference term.
		term = ast.Term().Make(reference)
		return term, token, true
	}

	// Attempt to parse a literal token.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken)
	if ok {
		// Found a literal term.
		term = ast.Term().Make(literal)
		return term, token, true
	}

	// This is not a term rule.
	return term, token, false
}

func (v *parser_) parseText() (
	text ast.TextLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a intrinsic token.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken)
	if ok {
		// Found a intrinsic text.
		text = ast.Text().Make(intrinsic)
		return text, token, true
	}

	// Attempt to parse a glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if ok {
		// Found a glyph text.
		text = ast.Text().Make(glyph)
		return text, token, true
	}

	// Attempt to parse a literal token.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken)
	if ok {
		// Found a literal text.
		text = ast.Text().Make(literal)
		return text, token, true
	}

	// Attempt to parse a lowercase token.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if ok {
		// Found a lowercase text.
		text = ast.Text().Make(lowercase)
		return text, token, true
	}

	// This is not a text rule.
	return text, token, false
}

func (v *parser_) parseDelimiter(expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == expectedValue {
			// Found the right delimiter.
			return value, token, true
		}
		v.putBack(token)
	}

	// This is not the right delimiter.
	return value, token, false
}

func (v *parser_) parseToken(tokenType TokenType) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the specific token.
	token = v.getNextToken()
	if token == nil {
		// We are at the end-of-file marker.
		return value, token, false
	}
	if token.GetType() == tokenType {
		// Found the right token type.
		value = token.GetValue()
		return value, token, true
	}

	// This is not the right token type.
	v.putBack(token)
	return value, token, false
}

func (v *parser_) formatError(token TokenLike, ruleName string) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		Scanner().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source line with the error in it.
	message += "\033[36m"
	if line > 1 {
		message += fmt.Sprintf("%04d: ", line-1) + string(lines[line-2]) + "\n"
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if col.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			syntax_[ruleName],
		)
	}
	return message
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveTop()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveHead() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError(token, "")
		panic(message)
	}

	return token
}

func (v *parser_) putBack(token TokenLike) {
	v.next_.AddValue(token)
}

var syntax_ = map[string]string{
	"Syntax": `Notice comment Rule+ comment Expression+`,
	"Notice": `comment newline`,
	"Rule":   `uppercase ":" Definition newline+`,
	"Definition": `,
  - Inline
  - Multiline`,
	"Inline": `Term+ note?`,
	"Term": `,
  - Reference
  - literal`,
	"Reference": `Identifier Cardinality?  ! The default cardinality is one.`,
	"Identifier": `,
  - lowercase
  - uppercase`,
	"Cardinality": `,
  - Constrained
  - Quantified`,
	"Constrained": `,
  - optional
  - repeated`,
	"Quantified":  `"{" number Limit? "}"`,
	"Limit":       `".." number?  ! The limit of a range of numbers is inclusive.`,
	"Multiline":   `newline Line+`,
	"Line":        `"-" Identifier note? newline`,
	"Expression":  `lowercase ":" Pattern note? newline+`,
	"Pattern":     `Option Alternative*`,
	"Alternative": `"|" Option`,
	"Option":      `Repetition+`,
	"Repetition":  `Element Cardinality?  ! The default cardinality is one.`,
	"Element": `,
  - Group
  - Filter
  - Text`,
	"Group":  `"(" Pattern ")"`,
	"Filter": `excluded? "[" Character+ "]"`,
	"Character": `,
  - Explicit
  - intrinsic`,
	"Explicit": `glyph Extent?`,
	"Extent":   `".." glyph  ! The extent of a range of glyphs is inclusive.`,
	"Text": `,
  - intrinsic
  - glyph
  - literal
  - lowercase`,
}
