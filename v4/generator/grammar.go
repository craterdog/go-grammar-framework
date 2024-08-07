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

import (
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
	sts "strings"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var grammarClass = &grammarClass_{
	// Initialize the class constants.
}

// Function

func Grammar() GrammarClassLike {
	return grammarClass
}

// CLASS METHODS

// Target

type grammarClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *grammarClass_) Make() GrammarLike {
	var grammar = &grammar_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: gra.Processor().Make(),
	}
	grammar.visitor_ = gra.Visitor().Make(grammar)
	return grammar
}

// INSTANCE METHODS

// Target

type grammar_ struct {
	// Define the instance attributes.
	class_   GrammarClassLike
	visitor_ gra.VisitorLike
	tokens_  abs.SetLike[string]
	rules_   abs.SetLike[string]
	plurals_ abs.SetLike[string]

	// Define the inherited aspects.
	gra.Methodical
}

// Attributes

func (v *grammar_) GetClass() GrammarClassLike {
	return v.class_
}

// Public

func (v *grammar_) GenerateGrammarModel(
	module string,
	wiki string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.visitor_.VisitSyntax(syntax)
	implementation = grammarTemplate_
	implementation = sts.ReplaceAll(implementation, "<wiki>", wiki)
	var name = v.extractSyntaxName(syntax)
	implementation = sts.ReplaceAll(implementation, "<module>", module)
	var notice = v.extractNotice(syntax)
	implementation = sts.ReplaceAll(implementation, "<Notice>", notice)
	var uppercase = v.makeUppercase(name)
	implementation = sts.ReplaceAll(implementation, "<Name>", uppercase)
	var lowercase = v.makeLowercase(name)
	implementation = sts.ReplaceAll(implementation, "<name>", lowercase)
	implementation = sts.ReplaceAll(implementation, "<parameter>", lowercase)
	var tokenTypes = v.extractTokenTypes()
	implementation = sts.ReplaceAll(implementation, "<TokenTypes>", tokenTypes)
	var processTokens = v.extractProcessTokens()
	implementation = sts.ReplaceAll(implementation, "<ProcessTokens>", processTokens)
	var processRules = v.extractProcessRules()
	implementation = sts.ReplaceAll(implementation, "<ProcessRules>", processRules)
	return implementation
}

// Methodical

func (v *grammar_) PreprocessIdentifier(
	identifier ast.IdentifierLike,
) {
	var name = identifier.GetAny().(string)
	switch {
	case gra.Scanner().MatchesType(name, gra.LowercaseToken):
		v.tokens_.AddValue(name)
	}
}

func (v *grammar_) PreprocessPredicate(
	predicate ast.PredicateLike,
) {
	var identifier = v.makeLowercase(predicate.GetIdentifier().GetAny().(string))
	var cardinality = predicate.GetOptionalCardinality()
	if col.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case ast.ConstrainedLike:
			v.plurals_.AddValue(identifier)
		case string:
			switch actual {
			case "*", "+":
				v.plurals_.AddValue(identifier)
			}
		}
	}
}

func (v *grammar_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	var name = rule.GetUppercase()
	v.rules_.AddValue(v.makeLowercase(name))
}

func (v *grammar_) PreprocessSyntax(syntax ast.SyntaxLike) {
	v.tokens_ = col.Set[string]([]string{"reserved"})
	v.rules_ = col.Set[string]()
	v.plurals_ = col.Set[string]()
}

// Private

func (v *grammar_) extractNotice(syntax ast.SyntaxLike) string {
	var header = syntax.GetHeaders().GetIterator().GetNext()
	var comment = header.GetComment()

	// Strip off the syntax style comment delimiters.
	var notice = comment[2 : len(comment)-3]

	return notice
}

func (v *grammar_) extractProcessRules() string {
	var processRules string
	var iterator = v.rules_.GetIterator()
	for iterator.HasNext() {
		var lowercase = iterator.GetNext()
		var isPlural = v.plurals_.ContainsValue(lowercase)
		var uppercase = v.makeUppercase(lowercase)
		var parameters = "("
		if isPlural {
			parameters += "\n\t\t"
		}
		parameters += lowercase + " ast." + uppercase + "Like"
		if isPlural {
			parameters += ",\n\t\tindex uint"
			parameters += ",\n\t\tsize uint,\n\t"
		}
		parameters += ")"
		processRules += "\n\tPreprocess" + uppercase + parameters
		processRules += "\n\tPostprocess" + uppercase + parameters
	}
	processRules += "\n"
	return processRules
}

func (v *grammar_) extractProcessTokens() string {
	var processTokens string
	var iterator = v.tokens_.GetIterator()
	for iterator.HasNext() {
		var name = iterator.GetNext()
		var isPlural = v.plurals_.ContainsValue(name)
		var parameters = "("
		if isPlural {
			parameters += "\n\t\t"
		}
		parameters += name + " string"
		if isPlural {
			parameters += ",\n\t\tindex uint"
			parameters += ",\n\t\tsize uint,\n\t"
		}
		parameters += ")"
		processTokens += "\n\tProcess" + v.makeUppercase(name) + parameters
	}
	return processTokens
}

func (v *grammar_) extractTokenTypes() string {
	var tokenTypes = "ErrorToken TokenType = iota"
	var tokens = col.Set[string](v.tokens_)
	tokens.AddValue("space")
	tokens.AddValue("newline")
	var iterator = tokens.GetIterator()
	for iterator.HasNext() {
		var name = iterator.GetNext()
		var tokenType = v.makeUppercase(name) + "Token"
		tokenTypes += "\n\t" + tokenType
	}
	return tokenTypes
}

func (v *grammar_) extractSyntaxName(syntax ast.SyntaxLike) string {
	var rule = syntax.GetRules().GetIterator().GetNext()
	var name = rule.GetUppercase()
	return name
}

func (v *grammar_) makeLowercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	name = string(runes)
	if reserved_[name] {
		name += "_"
	}
	return name
}

func (v *grammar_) makeUppercase(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToUpper(runes[0])
	return string(runes)
}

const grammarTemplate_ = `/*<Notice>*/

/*
Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

For detailed documentation on this package refer to the wiki:
  - https://<wiki>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package grammar

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "<module>/ast"
)

// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	<TokenTypes>
)

// Classes

/*
FormatterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructors
	Make() FormatterLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructors
	Make() ParserLike
}

/*
ProcessorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructors
	Make() ProcessorLike
}

/*
ScannerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructors
	Make(
		source string,
		tokens abs.QueueLike[TokenLike],
	) ScannerLike

	// Functions
	FormatToken(token TokenLike) string
	FormatType(tokenType TokenType) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructors
	Make(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructors
	Make() ValidatorLike
}

/*
VisitorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructors
	Make(
		processor Methodical,
	) VisitorLike
}

// Instances

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Attributes
	GetClass() FormatterClassLike
	GetDepth() uint

	// Abstractions
	Methodical

	// Methods
	Format<Name>(<parameter> ast.<Name>Like) string
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Attributes
	GetClass() ParserClassLike

	// Methods
	ParseSource(source string) ast.<Name>Like
}

/*
ProcessorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Attributes
	GetClass() ProcessorClassLike

	// Abstractions
	Methodical
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Attributes
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Attributes
	GetClass() TokenClassLike
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Attributes
	GetClass() ValidatorClassLike

	// Abstractions
	Methodical

	// Methods
	Validate<Name>(<parameter> ast.<Name>Like)
}

/*
VisitorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Attributes
	GetClass() VisitorClassLike

	// Methods
	Visit<Name>(<parameter> ast.<Name>Like)
}

// Aspects

/*
Methodical defines the set of method signatures that must be supported
by all methodical processors.
*/
type Methodical interface {<ProcessTokens><ProcessRules>}
`
