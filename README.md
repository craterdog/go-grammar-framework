<img src="https://craterdog.com/images/CraterDog.png" width="50%">

## Go Language Grammar Framework

### Overview
This project provides a Go based framework for writing, validating and formatting
language grammars defined using the Crater Dog Syntax Notation™.  This notation
is based on
[Wirth Syntax Notation (WSN)](https://en.wikipedia.org/wiki/Wirth_syntax_notation)
but adds support for comments and negation.

### Getting Started
The general development process—in a nutshell—is as follows:
 1. Install the
    [go-grammar-tools](https://github.com/craterdog/go-grammar-tools) module.
 1. Run the `bin/initialize` program to create a `Grammar.cdsn` syntax notation
    template file in your package directory.
 1. Fill in the `Grammar.cdsn` template with the rule and token definitions for
    the language grammar that this package will support.
 1. Run the `bin/generate` program to generate the corresponding `Package.go`
    class model file in your package directory.
 1. Fill in specific method implementations for the generated scanner, parser,
    validator and formatter classes.

### Contributing
Project contributors are always welcome. Check out the contributing guidelines
[here](https://github.com/craterdog/go-grammar-framework/blob/main/.github/CONTRIBUTING.md).

<H5 align="center"> Copyright © 2009 - 2024  Crater Dog Technologies™. All rights reserved. </H5>
