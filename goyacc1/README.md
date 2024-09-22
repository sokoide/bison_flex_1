# Goyacc1

* Simple Goyacc example

## Files

* hoge.go.y ... grammar file
* hoge.go ... generated from hoge.go.y
* lexer.y ... lexer
  * `Lex()` is called by the parser to get the next token
  * `Error()` is called if three is a lexical problem (tokenize error), for example, if you type "3 * ten"
* main.go ... main

## What is this?

1. Package and imports:

The generated parser will be part of the main package.
We import necessary Go packages.

2. Union declaration:

%union defines the types that can be associated with grammar symbols.
Here, we only have val of type int.

3. Token and type declarations:

%token <val> NUM declares a token for numeric values.
%type <val> expr specifies that expressions will have an integer value.

4. Operator precedence:

We define precedence for operators.

5. Grammar rules:

The rules are similar to the C version, but use Go syntax for actions.
Each action now assigns to $$ and references $1, $3, etc., as in C Yacc.

6. Lexer implementation:

The Lexer struct and its methods handle tokenization.
Lex method reads the input and returns tokens.
Error method handles syntax errors.

Main function:

Creates a new lexer and starts parsing.
