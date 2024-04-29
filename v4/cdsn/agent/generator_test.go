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
	age "github.com/craterdog/go-grammar-framework/v4/cdsn/agent"
	osx "os"
	sts "strings"
	tes "testing"
)

const testDirectory = "../../test/"

func TestInitialization(t *tes.T) {
	var syntaxName = "example"
	var copyright string
	var directory = testDirectory + syntaxName + "/"
	var err = osx.RemoveAll(directory)
	if err != nil {
		panic(err)
	}
	var generator = age.Generator().Make()
	generator.CreateSyntax(testDirectory, syntaxName, copyright)
}

func TestGeneration(t *tes.T) {
	var generator = age.Generator().Make()

	var directories, err = osx.ReadDir(testDirectory)
	if err != nil {
		panic(err)
	}

	for _, directory := range directories {
		var syntaxName = directory.Name()
		if sts.HasPrefix(syntaxName, "go.") {
			continue // This is not a syntax directory.
		}
		var syntaxDirectory = testDirectory + syntaxName
		fmt.Println(syntaxDirectory)

		var astDirectory = syntaxDirectory + "/ast/"
		err = osx.RemoveAll(astDirectory)
		if err != nil {
			panic(err)
		}
		err = osx.MkdirAll(astDirectory, 0755)
		if err != nil {
			panic(err)
		}
		generator.GenerateAST(testDirectory, syntaxName)

		var agentDirectory = syntaxDirectory + "/agent/"
		err = osx.RemoveAll(agentDirectory)
		if err != nil {
			panic(err)
		}
		err = osx.MkdirAll(agentDirectory, 0755)
		if err != nil {
			panic(err)
		}
		generator.GenerateAgents(testDirectory, syntaxName)
	}
}
