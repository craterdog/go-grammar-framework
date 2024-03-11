/*******************************************************************************
 *   Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammars

import (
	cla "github.com/craterdog/go-class-framework/v2"
	osx "os"
	sts "strings"
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
		// This class does not initialize any instance attributes.
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	copyright_  cla.CopyrightLike
	header_     cla.HeaderLike
	imports_    cla.ImportsLike
	types_      cla.TypesLike
	interfaces_ cla.InterfacesLike
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

func (v *generator_) processGrammar(grammar GrammarLike) cla.GoPNLike {
	var iterator = grammar.GetStatements().GetIterator()
	for iterator.HasNext() {
		var statement = iterator.GetNext()
		v.processStatement(statement)
	}
	var gopn = cla.GoPN().MakeWithAttributes(
		v.copyright_,
		v.header_,
		v.imports_,
		v.types_,
		v.interfaces_,
	)
	return gopn
}

func (v *generator_) processStatement(statement StatementLike) {
}
