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

const packageCommentTemplate_ = `
/*
Package "<packagename>" provides...

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/<module-name>/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-package-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
`

const classCommentTemplate_ = `
/*
<ClassName>ClassLike is a class interface that defines the set of class
constants, constructors and functions that must be supported by each
<class-name>-like concrete class.
*/
`

const instanceCommentTemplate_ = `
/*
<ClassName>Like is an instance interface that defines the complete set of
abstractions and methods that must be supported by each instance of a
<class-name>-like concrete class.
*/
`

const grammarTemplate_ = `
!>
................................................................................
<Copyright>
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
<!

!>
<Notation>
This document contains a formal definition of ...
A language grammar consists of a set of token and rule definitions.

TOKEN DEFINITIONS
The following token definitions are used by the scanner to generate the
stream of tokens that are processed by the parser.  Each token name begins
with an upper case letter.  Unlike with rule definitions, a token definition
cannot specify the name of a rule within its definition but it can specify
the name of other tokens.

The following intrinsic token types are environment or language specific:
 * ANY - Any language specific character.
 * LOWER - Any language specific lower case character.
 * UPPER - Any language specific upper case character.
 * DIGIT - Any language specific digit.
 * ESCAPE - Any environment specific escape sequence.
 * CONTROL - Any environment specific (non-printable) control character.
 * EOL - The environment specific end-of-line character.
 * EOF - The environment specific end-of-file marker.

Token definitions cannot be recursive and the scanning of tokens is not
greedy.  Any spaces within a token definition are NOT ignored.
<!

Token: UPPER (LOWER | UPPER)*

!>
RULE DEFINITIONS
The following rules are used by the parser when parsing the stream of tokens
generated by the scanner.  Each rule name begins with a lower case letter.
The rule definitions may specify the names of tokens or other rules and are
matched by the parser in the order listed.  A rule definition may also be
directly or indirectly recursive.  The sequence of factors within in a rule
definition may be separated by spaces which are ignored by the parser.

A predicate within a factor may also be constrained by any of the following
cardinalities:
 * predicate{M} - Exactly M instances of the specified predicate.
 * predicate{M..} - M or more instances of the specified predicate.
 * predicate{M..N} - M to N instances of the specified predicate.
 * predicate? - Zero or one instances of the specified predicate.
 * predicate* - Zero or more instances of the specified predicate.
 * predicate+ - One or more instances of the specified predicate.

Inversion within a rule definition may only be applied to an assertion
resulting in a single character or glyph.
<!

source: rule EOF  ! Terminated with an end-of-file marker.

rule: Token+

`
