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
	pac "github.com/craterdog/go-package-framework/v2"
	osx "os"
	sts "strings"
	tim "time"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var generatorClass = &generatorClass_{
	// This class does not initialize any class constants.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
	// This class does not define any class constants.
}

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	return &generator_{
		classes_:   col.Catalog[string, pac.ClassLike]().Make(),
		instances_: col.Catalog[string, pac.InstanceLike]().Make(),
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	classes_   col.CatalogLike[string, pac.ClassLike]
	instances_ col.CatalogLike[string, pac.InstanceLike]
}

// Public

func (v *generator_) CreateGrammar(directory string, copyright string) {
	// Center and insert the copyright string into the grammar template.
	var maximum = 78
	var length = len(copyright)
	if length > maximum {
		var message = fmt.Sprintf(
			"The copyright notice cannot be longer than 78 characters: %v",
			copyright,
		)
		panic(message)
	}
	if length == 0 {
		copyright = fmt.Sprintf(
			"Copyright (c) %v.  All Rights Reserved.",
			tim.Now().Year(),
		)
		length = len(copyright)
	}
	var padding = (maximum - length) / 2
	for range padding {
		copyright = " " + copyright + " "
	}
	if len(copyright) < maximum {
		copyright = " " + copyright
	}
	copyright = "." + copyright + "."
	var template = sts.ReplaceAll(grammarTemplate_, "<Copyright>", copyright)
	var bytes = []byte(template[1:]) // Remove leading "\n".

	// Save the new grammar template.
	v.createDirectory(directory)
	var grammarFile = directory + "Grammar.cdsn"
	fmt.Printf(
		"The grammar file %q does not exist, creating a template for it.\n",
		grammarFile,
	)
	var err = osx.WriteFile(grammarFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) GeneratePackage(directory string) {
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var grammar = v.parseGrammar(directory)
	if grammar == nil {
		return
	}
	var package_ = v.processGrammar(grammar)
	v.generatePackage(directory, package_)
}

// Private

func (v *generator_) createDirectory(directory string) {
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) generateClassComment(className string) string {
	var comment = classCommentTemplate_
	comment = sts.ReplaceAll(comment, "<ClassName>", className)
	comment = sts.ReplaceAll(comment, "<class-name>", sts.ToLower(className))
	return comment
}

func (v *generator_) generateNotice(grammar GrammarLike) pac.NoticeLike {
	return nil
}

func (v *generator_) generateHeader() pac.HeaderLike {
	return nil
}

func (v *generator_) generateImports() pac.ImportsLike {
	return nil
}

func (v *generator_) generateInstanceComment(className string) string {
	var comment = instanceCommentTemplate_
	comment = sts.ReplaceAll(comment, "<ClassName>", className)
	comment = sts.ReplaceAll(comment, "<class-name>", sts.ToLower(className))
	return comment
}

func (v *generator_) generateInterfaces(
	classMethods col.Sequential[pac.ClassLike],
	instanceMethods col.Sequential[pac.InstanceLike],
) pac.InterfacesLike {
	var classes = pac.Classes().MakeWithAttributes(classMethods)
	var instances = pac.Instances().MakeWithAttributes(instanceMethods)
	var interfaces = pac.Interfaces().MakeWithAttributes(
		nil,
		classes,
		instances,
	)
	return interfaces
}

func (v *generator_) generatePackage(directory string, package_ pac.PackageLike) {
	var validator = pac.Validator().Make()
	validator.ValidatePackage(package_)
	var formatter = pac.Formatter().Make()
	var source = formatter.FormatPackage(package_)
	var bytes = []byte(source)
	var err = osx.WriteFile(directory+"package.go", bytes, 0644)
	if err != nil {
		panic(err)
	}
	var generator = pac.Generator().Make()
	generator.GeneratePackage(directory)
}

func (v *generator_) makePrivate(identifier string) string {
	runes := []rune(identifier)
	runes[0] = uni.ToLower(runes[0])
	return string(runes)
}

func (v *generator_) parseGrammar(directory string) GrammarLike {
	var grammarFile = directory + "Grammar.cdsn"
	var bytes, err = osx.ReadFile(grammarFile)
	if err != nil {
		var message = fmt.Sprintf(
			"The specified directory is missing a grammar file: %v",
			grammarFile,
		)
		panic(message)
	}
	var source = string(bytes)
	var parser = Parser().Make()
	var grammar = parser.ParseSource(source)
	var validator = Validator().Make()
	validator.ValidateGrammar(grammar)
	return grammar
}

func (v *generator_) processDefinition(
	definition DefinitionLike,
	classes col.ListLike[pac.ClassLike],
	instances col.ListLike[pac.InstanceLike],
) {
	var symbol = definition.GetSymbol()
	var expression = definition.GetExpression()
	if uni.IsUpper([]rune(symbol)[0]) {
		// Ignore token definitions.
		return
	}
	v.processRule(symbol, expression, classes, instances)
}

func (v *generator_) processExpression(
	expression ExpressionLike,
	constructorMethods col.ListLike[pac.ConstructorLike],
	attributeMethods col.ListLike[pac.AttributeLike],
) {
}

func (v *generator_) processGrammar(grammar GrammarLike) pac.PackageLike {
	var classes = col.List[pac.ClassLike]().Make()
	var instances = col.List[pac.InstanceLike]().Make()
	var iterator = grammar.GetStatements().GetIterator()
	for iterator.HasNext() {
		var statement = iterator.GetNext()
		v.processStatement(statement, classes, instances)
	}
	var copyright = v.generateNotice(grammar)
	var header = v.generateHeader()
	var imports = v.generateImports()
	var interfaces = v.generateInterfaces(classes, instances)
	var package_ = pac.Package().MakeWithAttributes(
		copyright,
		header,
		imports,
		nil,
		interfaces,
	)
	return package_
}

func (v *generator_) processRule(
	symbol string,
	expression ExpressionLike,
	classes col.ListLike[pac.ClassLike],
	instances col.ListLike[pac.InstanceLike],
) {
	var constructorMethods = col.List[pac.ConstructorLike]().Make()
	var attributeMethods = col.List[pac.AttributeLike]().Make()
	v.processExpression(expression, constructorMethods, attributeMethods)
	var className = v.makePrivate(symbol)
	var declaration = pac.Declaration().MakeWithAttributes(
		v.generateClassComment(className),
		className+"ClassLike",
		nil,
	)
	var constructors = pac.Constructors().MakeWithAttributes(constructorMethods)
	var class = pac.Class().MakeWithAttributes(
		declaration,
		nil,
		constructors,
		nil,
	)
	declaration = pac.Declaration().MakeWithAttributes(
		v.generateInstanceComment(className),
		className+"Like",
		nil,
	)
	var attributes = pac.Attributes().MakeWithAttributes(attributeMethods)
	var instance = pac.Instance().MakeWithAttributes(
		declaration,
		attributes,
		nil,
		nil,
	)
	classes.AppendValue(class)
	instances.AppendValue(instance)
}

func (v *generator_) processStatement(
	statement StatementLike,
	classes col.ListLike[pac.ClassLike],
	instances col.ListLike[pac.InstanceLike],
) {
	var definition = statement.GetDefinition()
	if definition == nil {
		// Ignore comments.
		return
	}
	v.processDefinition(definition, classes, instances)
}
