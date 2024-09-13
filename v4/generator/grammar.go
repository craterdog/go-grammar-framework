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
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
	gra "github.com/craterdog/go-grammar-framework/v4/grammar"
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
		class_:    c,
		analyzer_: gra.Analyzer().Make(),
	}
	return grammar
}

// INSTANCE METHODS

// Target

type grammar_ struct {
	// Define the instance attributes.
	class_    GrammarClassLike
	analyzer_ gra.AnalyzerLike
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
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = grammarTemplate_
	implementation = replaceAll(implementation, "module", module)
	implementation = replaceAll(implementation, "wiki", wiki)
	var notice = v.analyzer_.GetNotice()
	implementation = replaceAll(implementation, "notice", notice)
	var name = v.analyzer_.GetSyntaxName()
	implementation = replaceAll(implementation, "name", name)
	implementation = replaceAll(implementation, "parameter", name)
	var tokenTypes = v.generateTokenTypes()
	implementation = replaceAll(implementation, "tokenTypes", tokenTypes)
	var processTokens = v.generateProcessTokens()
	implementation = replaceAll(implementation, "processTokens", processTokens)
	var processRules = v.generateProcessRules()
	implementation = replaceAll(implementation, "processRules", processRules)
	return implementation
}

// Private

func (v *grammar_) generateProcessRules() string {
	var processRules string
	var iterator = v.analyzer_.GetRuleNames().GetIterator()
	for iterator.HasNext() {
		var ruleName = iterator.GetNext()
		var parameterName = makeLowerCase(ruleName)
		var className = makeUpperCase(ruleName)
		var isPlural = v.analyzer_.IsPlural(ruleName)
		var parameters = "("
		if isPlural {
			parameters += "\n\t\t"
		}
		parameters += parameterName + " ast." + className + "Like"
		if isPlural {
			parameters += ",\n\t\tindex uint"
			parameters += ",\n\t\tsize uint,\n\t"
		}
		parameters += ")"
		processRules += "\n\tPreprocess" + className + parameters
		processRules += "\n\tPostprocess" + className + parameters
	}
	processRules += "\n"
	return processRules
}

func (v *grammar_) generateProcessTokens() string {
	var processTokens string
	var iterator = v.analyzer_.GetTokenNames().GetIterator()
	for iterator.HasNext() {
		var name = iterator.GetNext()
		if v.analyzer_.IsIgnored(name) || name == "delimiter" {
			continue
		}
		var isPlural = v.analyzer_.IsPlural(name)
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
		processTokens += "\n\tProcess" + makeUpperCase(name) + parameters
	}
	return processTokens
}

func (v *grammar_) generateTokenTypes() string {
	var tokenTypes = "ErrorToken TokenType = iota"
	var iterator = v.analyzer_.GetTokenNames().GetIterator()
	for iterator.HasNext() {
		var name = iterator.GetNext()
		var tokenType = makeUpperCase(name) + "Token"
		tokenTypes += "\n\t" + tokenType
	}
	return tokenTypes
}

const grammarTemplate_ = `<Notice>

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
AnalyzerClassLike defines the set of class constants, constructors and
functions that must be supported by all analyzer-class-like classes.
*/
type AnalyzerClassLike interface {
	// Constructors
	Make() AnalyzerLike
}

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
	Make(processor Methodical) VisitorLike
}

// Instances

/*
AnalyzerLike defines the set of aspects and methods that must be supported by
all analyzer-like instances.
*/
type AnalyzerLike interface {
	// Attributes
	GetClass() AnalyzerClassLike

	// Abstractions
	Methodical

	// Methods
	AnalyzeSyntax(syntax ast.SyntaxLike)
	GetExpressions() abs.Sequential[abs.AssociationLike[string, string]]
	GetIdentifiers(ruleName string) abs.Sequential[ast.IdentifierLike]
	GetIgnored() abs.Sequential[string]
	GetNotice() string
	GetReferences(ruleName string) abs.Sequential[ast.ReferenceLike]
	GetRuleNames() abs.Sequential[string]
	GetSyntaxName() string
	GetTerms(ruleName string) abs.Sequential[ast.TermLike]
	GetTokenNames() abs.Sequential[string]
	IsDelimited(ruleName string) bool
	IsIgnored(tokenName string) bool
	IsPlural(name string) bool
}

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
