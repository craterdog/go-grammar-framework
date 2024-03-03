/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammar

import (
	//fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v3"
	reg "regexp"
	sts "strings"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var scannerClass = &scannerClass_{
	matchers: map[TokenType]*reg.Regexp{
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
		SymbolToken:    reg.MustCompile(`^(?:` + symbol_ + `)`),
	},
}

// Function

func Scanner() ScannerClassLike {
	return scannerClass
}

// CLASS METHODS

// Target

type scannerClass_ struct {
	matchers map[TokenType]*reg.Regexp
}

// Constructors

func (c *scannerClass_) Make(
	source string,
	tokens col.QueueLike[TokenLike],
) ScannerLike {
	var scanner = &scanner_{
		line:     1,
		position: 1,
		runes:    []rune(source),
		tokens:   tokens,
	}
	go scanner.scanTokens() // Start scanning tokens in the background.
	return scanner
}

// Functions

func (c *scannerClass_) MatchToken(
	tokenType TokenType,
	text string,
) col.ListLike[string] {
	var matcher = c.matchers[tokenType]
	var matches = matcher.FindStringSubmatch(text)
	return col.List[string]().MakeFromArray(matches)
}

// INSTANCE METHODS

// Target

type scanner_ struct {
	first    int // A zero based index of the first possible rune in the next token.
	line     int // The line number in the source string of the next rune.
	next     int // A zero based index of the next possible rune in the next token.
	position int // The position in the current line of the next rune.
	runes    []rune
	tokens   col.QueueLike[TokenLike]
}

// Private

func (v *scanner_) emitToken(tokenType TokenType) {
	var tokenValue = string(v.runes[v.first:v.next])
	switch tokenValue {
	case "\x00":
		tokenValue = "<NULL>"
	case "\a":
		tokenValue = "<BELL>"
	case "\b":
		tokenValue = "<BKSP>"
	case "\t":
		tokenValue = "<HTAB>"
	case "\f":
		tokenValue = "<FMFD>"
	case "\n":
		tokenValue = "<EOLN>"
	case "\r":
		tokenValue = "<CRTN>"
	case "\v":
		tokenValue = "<VTAB>"
	}
	var token = Token().Make(v.line, v.position, tokenType, tokenValue)
	//fmt.Println(token) // Uncomment when debugging.
	v.tokens.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundEOF() {
	v.emitToken(EOFToken)
}

func (v *scanner_) foundError() {
	v.next++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(tokenType TokenType) bool {
	var text = string(v.runes[v.next:])
	var matches = Scanner().MatchToken(tokenType, text)
	if !matches.IsEmpty() {
		var match = matches.GetValue(1)
		var token = []rune(match)
		var length = len(token)
		var nextIndex = v.next + length
		if nextIndex < len(v.runes) {
			var nextRune = v.runes[v.next+length]
			if tokenType == IntrinsicToken && (uni.IsLetter(nextRune) ||
				uni.IsDigit(nextRune) || nextRune == rune('_')) {
				// This is not an intrinsic token.
				return false
			}
		}
		v.next += length
		if tokenType != SpaceToken {
			v.emitToken(tokenType)
		}
		var count = sts.Count(match, "\n")
		if count > 0 {
			v.line += count
			v.position = v.indexOfLastEOL(token)
		} else {
			v.position += v.next - v.first
		}
		v.first = v.next
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
	for v.next < len(v.runes) {
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
		case v.foundToken(SymbolToken):
		default:
			v.foundError()
			break loop
		}
	}
	v.foundEOF()
	v.tokens.CloseQueue()
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
	any_       = `.|` + eol_
	base16_    = `[0-9a-f]`
	character_ = `['][^` + control_ + `][']`
	comment_   = `!>(?:` + any_ + `)*?<!` // This returns the shortest match.
	control_   = `\p{Cc}`
	delimiter_ = `[~?*+:|(){}]|\.\.`
	digit_     = `\p{Nd}`
	eol_       = `\n`
	escape_    = `\\(?:(?:` + unicode_ + `)|[abfnrtv'"\\])`
	intrinsic_ = `ANY|LOWER|UPPER|DIGIT|ESCAPE|CONTROL|EOL|EOF`
	letter_    = lower_ + `|` + upper_
	literal_   = `["](?:` + escape_ + `|[^"` + eol_ + `])+["]`
	lower_     = `\p{Ll}`
	name_      = `(?:` + letter_ + `)(?:_?(?:` + letter_ + `|` + digit_ + `))*`
	note_      = `! [^` + control_ + `]*`
	number_    = `(?:` + digit_ + `)+`
	space_     = `[ \t]+`
	symbol_    = `[$](` + name_ + `)`
	unicode_   = `x` + base16_ + `{2}|u` + base16_ + `{4}|U` + base16_ + `{8}`
	upper_     = `\p{Lu}`
)
