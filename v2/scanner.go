/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package grammars

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v3"
	reg "regexp"
	sts "strings"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var scannerClass = &scannerClass_{
	tokens_: map[TokenType]string{
		ErrorToken:     "error",
		CharacterToken: "character",
		CommentToken:   "comment",
		DelimiterToken: "delimiter",
		EOFToken:       "EOF",
		EOLToken:       "EOL",
		IntrinsicToken: "intrinsic",
		LiteralToken:   "literal",
		NameToken:      "name",
		NoteToken:      "note",
		NumberToken:    "number",
		SpaceToken:     "space",
	},
	matchers_: map[TokenType]*reg.Regexp{
		CharacterToken: reg.MustCompile(`^(?:` + character_ + `)`),
		CommentToken:   reg.MustCompile(`^(?:` + comment_ + `)`),
		DelimiterToken: reg.MustCompile(`^(?:` + delimiter_ + `)`),
		EOLToken:       reg.MustCompile(`^(?:` + eol_ + `)`),
		IntrinsicToken: reg.MustCompile(`^(?:` + intrinsic_ + `)`),
		LiteralToken:   reg.MustCompile(`^(?:` + literal_ + `)`),
		NameToken:      reg.MustCompile(`^(?:` + name_ + `)`),
		NoteToken:      reg.MustCompile(`^(?:` + note_ + `)`),
		NumberToken:    reg.MustCompile(`^(?:` + number_ + `)`),
		SpaceToken:     reg.MustCompile(`^(?:` + space_ + `)`),
	},
}

// Function

func Scanner() ScannerClassLike {
	return scannerClass
}

// CLASS METHODS

// Target

type scannerClass_ struct {
	tokens_   map[TokenType]string
	matchers_ map[TokenType]*reg.Regexp
}

// Constructors

func (c *scannerClass_) Make(
	source string,
	tokens col.QueueLike[TokenLike],
) ScannerLike {
	var scanner = &scanner_{
		line_:     1,
		position_: 1,
		runes_:    []rune(source),
		tokens_:   tokens,
	}
	go scanner.scanTokens() // Start scanning tokens in the background.
	return scanner
}

// Functions

func (c *scannerClass_) FormatToken(token TokenLike) string {
	var value = token.GetValue()
	var s = fmt.Sprintf("%q", value)
	if len(s) > 40 {
		s = fmt.Sprintf("%.40q...", value)
	}
	return fmt.Sprintf(
		"Token [type: %s, line: %d, position: %d]: %s",
		c.tokens_[token.GetType()],
		token.GetLine(),
		token.GetPosition(),
		s,
	)
}

func (c *scannerClass_) MatchToken(
	type_ TokenType,
	text string,
) col.ListLike[string] {
	var matcher = c.matchers_[type_]
	var matches = matcher.FindStringSubmatch(text)
	return col.List[string]().MakeFromArray(matches)
}

// INSTANCE METHODS

// Target

type scanner_ struct {
	first_    int // A zero based index of the first possible rune in the next token.
	next_     int // A zero based index of the next possible rune in the next token.
	line_     int // The line number in the source string of the next rune.
	position_ int // The position in the current line of the next rune.
	runes_    []rune
	tokens_   col.QueueLike[TokenLike]
}

// Private

func (v *scanner_) emitToken(type_ TokenType) {
	var value = string(v.runes_[v.first_:v.next_])
	switch value {
	case "\x00":
		value = "<NULL>"
	case "\a":
		value = "<BELL>"
	case "\b":
		value = "<BKSP>"
	case "\t":
		value = "<HTAB>"
	case "\f":
		value = "<FMFD>"
	case "\n":
		value = "<EOLN>"
	case "\r":
		value = "<CRTN>"
	case "\v":
		value = "<VTAB>"
	}
	var token = Token().MakeWithAttributes(v.line_, v.position_, type_, value)
	//fmt.Println(Scanner().FormatToken(token)) // Uncomment when debugging.
	v.tokens_.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundEOF() {
	v.emitToken(EOFToken)
}

func (v *scanner_) foundError() {
	v.next_++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(type_ TokenType) bool {
	var text = string(v.runes_[v.next_:])
	var matches = Scanner().MatchToken(type_, text)
	if !matches.IsEmpty() {
		var match = matches.GetValue(1)
		var token = []rune(match)
		var length = len(token)
		var nextIndex = v.next_ + length
		if nextIndex < len(v.runes_) {
			var nextRune = v.runes_[v.next_+length]
			if type_ == IntrinsicToken && (uni.IsLetter(nextRune) ||
				uni.IsDigit(nextRune) || nextRune == rune('_')) {
				// This is not an intrinsic token.
				return false
			}
		}
		v.next_ += length
		if type_ != SpaceToken {
			v.emitToken(type_)
		}
		var count = sts.Count(match, "\n")
		if count > 0 {
			v.line_ += count
			v.position_ = v.indexOfLastEOL(token)
		} else {
			v.position_ += v.next_ - v.first_
		}
		v.first_ = v.next_
		return true
	}
	return false
}

func (v *scanner_) indexOfLastEOL(runes []rune) int {
	var length = len(runes)
	for index := length; index > 0; index-- {
		if runes[index-1] == '\n' {
			return length - index + 1
		}
	}
	return 0
}

func (v *scanner_) scanTokens() {
loop:
	for v.next_ < len(v.runes_) {
		switch {
		case v.foundToken(CharacterToken):
		case v.foundToken(CommentToken):
		case v.foundToken(DelimiterToken):
		case v.foundToken(EOLToken):
		case v.foundToken(IntrinsicToken):
		case v.foundToken(LiteralToken):
		case v.foundToken(NameToken):
		case v.foundToken(NoteToken):
		case v.foundToken(NumberToken):
		case v.foundToken(SpaceToken):
		default:
			v.foundError()
			break loop
		}
	}
	v.foundEOF()
}

/*
NOTE:
These private constants define the regular expression sub-patterns that make up
all token types.  Unfortunately there is no way to make them private to the
scanner class since they must be TRUE Go constants to be initialized in this
way.  We append an underscore to each name to lessen the chance of a name
collision with other private Go class constants in this package.
*/
const (
	_any_      = `a^`
	_control_  = `\P{Cc}`
	_digit_    = `\P{Nd}`
	_eof_      = `[^\z]`
	_eol_      = `[^\n]`
	_escape_   = `[^\\]`
	_lower_    = `\P{Ll}`
	_upper_    = `\P{Lu}`
	any_       = `.|` + eol_
	base16_    = `[0-9a-f]`
	character_ = `['][^` + control_ + `][']`
	comment_   = `!>` + eol_ + `((?:` + any_ + `)*?)` + eol_ + `<!` + eol_
	control_   = `\p{Cc}`
	delimiter_ = `[~?*+:|(){}]|\.\.`
	digit_     = `\p{Nd}`
	eof_       = `\z`
	eol_       = `\n`
	escape_    = `\\(?:(?:` + unicode_ + `)|[abfnrtv'"\\])`
	intrinsic_ = `ANY|LOWER|UPPER|DIGIT|ESCAPE|CONTROL|EOL|EOF`
	letter_    = lower_ + `|` + upper_
	literal_   = `["](?:` + escape_ + `|[^` + control_ + `])+?["]`
	lower_     = `\p{Ll}`
	name_      = `(?:` + letter_ + `)(?:` + letter_ + `|` + digit_ + `)*`
	note_      = `! [^` + control_ + `]*`
	number_    = `(?:` + digit_ + `)+`
	space_     = `[ \t]+`
	unicode_   = `x` + base16_ + `{2}|u` + base16_ + `{4}|U` + base16_ + `{8}`
	upper_     = `\p{Lu}`
)
