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

package agent_test

import (
	fmt "fmt"
	gra "github.com/craterdog/go-grammar-framework/v3/agent"
	osx "os"
	sts "strings"
	tes "testing"
)

const inputDirectory = "../test/input/"
const outputDirectory = "../test/output/"
const grammarName = "Example"

func TestInitialization(t *tes.T) {
	var generator = gra.Generator().Make()

	var err = osx.RemoveAll(outputDirectory)
	if err != nil {
		panic(err)
	}

	err = osx.MkdirAll(outputDirectory, 0755)
	if err != nil {
		panic(err)
	}

	var copyright string
	generator.CreateGrammar(outputDirectory, grammarName, copyright)
}

func TestGeneration(t *tes.T) {
	var generator = gra.Generator().Make()

	var files, err = osx.ReadDir(inputDirectory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var fileSuffix = ".cdsn"
		var fileName = sts.TrimSuffix(file.Name(), fileSuffix)
		fmt.Println(fileName)
		var bytes, err = osx.ReadFile(inputDirectory + file.Name())
		if err != nil {
			panic(err)
		}
		var directoryName = outputDirectory + fileName + "/"
		err = osx.RemoveAll(directoryName)
		if err != nil {
			panic(err)
		}
		err = osx.MkdirAll(directoryName, 0755)
		if err != nil {
			panic(err)
		}
		err = osx.WriteFile(directoryName+"Grammar.cdsn", bytes, 0644)
		if err != nil {
			panic(err)
		}
		generator.GenerateModel(directoryName)
	}
}
