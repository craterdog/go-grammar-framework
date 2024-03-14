/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
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
	cla "github.com/craterdog/go-class-framework/v2"
	col "github.com/craterdog/go-collection-framework/v3"
	osx "os"
	sts "strings"
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
		imports_:   col.Catalog[string, cla.ModuleLike]().Make(),
		classes_:   col.Catalog[string, cla.ClassLike]().Make(),
		instances_: col.Catalog[string, cla.InstanceLike]().Make(),
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	imports_   col.CatalogLike[string, cla.ModuleLike]
	classes_   col.CatalogLike[string, cla.ClassLike]
	instances_ col.CatalogLike[string, cla.InstanceLike]
}

// Public

func (v *generator_) GeneratePackage(directory string) {
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	v.createDirectory(directory)
	var grammar = v.parseGrammar(directory)
	if grammar == nil {
		return
	}
	var gopn = v.processGrammar(grammar)
	v.generatePackage(directory, gopn)
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

func (v *generator_) generateCopyright() cla.CopyrightLike {
	return nil
}

func (v *generator_) generateHeader() cla.HeaderLike {
	return nil
}

func (v *generator_) generateImports() cla.ImportsLike {
	return nil
}

func (v *generator_) generateInstanceComment(className string) string {
	var comment = instanceCommentTemplate_
	comment = sts.ReplaceAll(comment, "<ClassName>", className)
	comment = sts.ReplaceAll(comment, "<class-name>", sts.ToLower(className))
	return comment
}

func (v *generator_) generateInterfaces(
	classMethods col.Sequential[cla.ClassLike],
	instanceMethods col.Sequential[cla.InstanceLike],
) cla.InterfacesLike {
	var classes = cla.Classes().MakeWithAttributes(classMethods)
	var instances = cla.Instances().MakeWithAttributes(instanceMethods)
	var interfaces = cla.Interfaces().MakeWithAttributes(
		nil,
		classes,
		instances,
	)
	return interfaces
}

func (v *generator_) generatePackage(directory string, gopn cla.GoPNLike) {
	var validator = cla.Validator().Make()
	validator.ValidatePackage(gopn)
	var formatter = cla.Formatter().Make()
	var source = formatter.FormatGoPN(gopn)
	var bytes = []byte(source)
	var err = osx.WriteFile(directory+"package.go", bytes, 0644)
	if err != nil {
		panic(err)
	}
	var generator = cla.Generator().Make()
	generator.GeneratePackage(directory)
}

func (v *generator_) makePrivate(identifier string) string {
	runes := []rune(identifier)
	runes[0] = uni.ToLower(runes[0])
	return string(runes)
}

func (v *generator_) parseGrammar(directory string) GrammarLike {
	var filename = directory + "grammar.cdsn"
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
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
	classes col.ListLike[cla.ClassLike],
	instances col.ListLike[cla.InstanceLike],
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
	constructorMethods col.ListLike[cla.ConstructorLike],
	attributeMethods col.ListLike[cla.AttributeLike],
) {
}

func (v *generator_) processGrammar(grammar GrammarLike) cla.GoPNLike {
	var classes = col.List[cla.ClassLike]().Make()
	var instances = col.List[cla.InstanceLike]().Make()
	var iterator = grammar.GetStatements().GetIterator()
	for iterator.HasNext() {
		var statement = iterator.GetNext()
		v.processStatement(statement, classes, instances)
	}
	var copyright = v.generateCopyright()
	var header = v.generateHeader()
	var imports = v.generateImports()
	var interfaces = v.generateInterfaces(classes, instances)
	var gopn = cla.GoPN().MakeWithAttributes(
		copyright,
		header,
		imports,
		nil,
		interfaces,
	)
	return gopn
}

func (v *generator_) processRule(
	symbol string,
	expression ExpressionLike,
	classes col.ListLike[cla.ClassLike],
	instances col.ListLike[cla.InstanceLike],
) {
	var constructorMethods = col.List[cla.ConstructorLike]().Make()
	var attributeMethods = col.List[cla.AttributeLike]().Make()
	v.processExpression(expression, constructorMethods, attributeMethods)
	var className = v.makePrivate(symbol)
	var declaration = cla.Declaration().MakeWithAttributes(
		v.generateClassComment(className),
		className+"ClassLike",
		nil,
	)
	var constructors = cla.Constructors().MakeWithAttributes(constructorMethods)
	var class = cla.Class().MakeWithAttributes(
		declaration,
		nil,
		constructors,
		nil,
	)
	declaration = cla.Declaration().MakeWithAttributes(
		v.generateInstanceComment(className),
		className+"Like",
		nil,
	)
	var attributes = cla.Attributes().MakeWithAttributes(attributeMethods)
	var instance = cla.Instance().MakeWithAttributes(
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
	classes col.ListLike[cla.ClassLike],
	instances col.ListLike[cla.InstanceLike],
) {
	var definition = statement.GetDefinition()
	if definition == nil {
		// Ignore comments.
		return
	}
	v.processDefinition(definition, classes, instances)
}
