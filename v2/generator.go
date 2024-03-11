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
	// This class does not define any instance attributes.
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
	v.generatePackage(directory, grammar)
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

func (v *generator_) generatePackage(directory string, grammar GrammarLike) {
	var copyright cla.CopyrightLike
	var header cla.HeaderLike
	var imports cla.ImportsLike
	var types cla.TypesLike
	var interfaces cla.InterfacesLike
	var gopn = cla.GoPN().MakeWithAttributes(
		copyright,
		header,
		imports,
		types,
		interfaces,
	)
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
	return nil
}
