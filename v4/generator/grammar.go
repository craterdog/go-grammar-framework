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
	ast "github.com/craterdog/go-grammar-framework/v4/ast"
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
		analyzer_: Analyzer().Make(),
	}
	return grammar
}

// INSTANCE METHODS

// Target

type grammar_ struct {
	// Define the instance attributes.
	class_    *grammarClass_
	analyzer_ AnalyzerLike
}

// Public

func (v *grammar_) GetClass() GrammarClassLike {
	return v.class_
}

func (v *grammar_) GenerateGrammarModel(
	module string,
	wiki string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	var header = v.getTemplate(packageHeader)
	implementation = v.getTemplate(modelTemplate)
	implementation = replaceAll(implementation, "header", header)
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
		if isReserved(parameterName) {
			parameterName += "_"
		}
		var className = makeUpperCase(ruleName)
		var parameters = "(\n\t\t"
		parameters += parameterName + " ast." + className + "Like,"
		if v.analyzer_.IsPlural(ruleName) {
			parameters += "\n\t\tindex uint,"
			parameters += "\n\t\tsize uint,"
		}
		parameters += "\n\t)"
		processRules += "\n\tPreprocess" + className + parameters
		processRules += "\n\tProcess" + className + "Slot(\n\t\tslot uint,\n\t)"
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
		if name == "delimiter" {
			continue
		}
		var parameters = "(\n\t\t"
		var parameter = name
		if isReserved(parameter) {
			parameter += "_"
		}
		parameters += parameter + " string,"
		if v.analyzer_.IsPlural(name) {
			parameters += "\n\t\tindex uint,"
			parameters += "\n\t\tsize uint,"
		}
		parameters += "\n\t)"
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

func (v *grammar_) getTemplate(name string) string {
	var template = grammarTemplates_.GetValue(name)
	return template
}

// PRIVATE GLOBALS

// Constants

var grammarTemplates_ = col.Catalog[string, string](
	map[string]string{
		packageHeader: `
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
*/`,
		modelTemplate: `<Notice>
<Header>
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
	// Constructor
	Make() FormatterLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor
	Make() ParserLike
}

/*
ProcessorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor
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
	// Constructor
	Make(
		source string,
		tokens abs.QueueLike[TokenLike],
	) ScannerLike

	// Function
	FormatToken(
		token TokenLike,
	) string
	FormatType(
		tokenType TokenType,
	) string
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
	// Constructor
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
	// Constructor
	Make() ValidatorLike
}

/*
VisitorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor
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
	// Public
	GetClass() FormatterClassLike
	Format<Name>(
		<parameter> ast.<Name>Like,
	) string

	// Aspect
	Methodical
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Public
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.<Name>Like
}

/*
ProcessorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Public
	GetClass() ProcessorClassLike

	// Aspect
	Methodical
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Public
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Public
	GetClass() TokenClassLike

	// Attribute
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
	// Public
	GetClass() ValidatorClassLike
	Validate<Name>(
		<parameter> ast.<Name>Like,
	)

	// Aspect
	Methodical
}

/*
VisitorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Public
	GetClass() VisitorClassLike
	Visit<Name>(
		<parameter> ast.<Name>Like,
	)
}

// Aspects

/*
Methodical defines the set of method signatures that must be supported
by all methodical processors.
*/
type Methodical interface {<ProcessTokens><ProcessRules>}
`,
	},
)
