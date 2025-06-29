# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the `2020_2` example within the `bison_flex_1` repository - a simple integer calculator implementation using Bison (yacc) and Flex (lex) in C. The calculator supports basic arithmetic operations (+, -, *, /) and parentheses for grouping.

## Architecture

The calculator is built using the traditional lexer/parser approach:
- `calc.l`: Flex lexer specification that tokenizes input (integers, operators, whitespace, newlines)
- `calc.y`: Bison parser specification that defines grammar rules and semantic actions
- Generated files: `lex.yy.c` (lexer) and `calc.tab.c` (parser) are auto-generated

The parser includes the lexer directly via `#include "lex.yy.c"` and implements a simple REPL that evaluates expressions line by line.

## Build Commands

- `make` or `make all`: Build the calculator executable
- `make clean`: Remove generated files and executable
- `make run`: Build and run the calculator
- `make calc.tab.c`: Generate parser from grammar file
- `make lex.yy.c`: Generate lexer from lexer specification

## Development Workflow

1. Modify `calc.y` for grammar changes or `calc.l` for lexer changes
2. Run `make` to regenerate and compile
3. Test with `make run` or run `./calc` directly
4. The calculator reads from stdin and evaluates expressions line by line

## Key Files

- `calc.y`: Bison grammar with arithmetic expression rules and precedence
- `calc.l`: Flex lexer with token definitions
- `Makefile`: Build configuration using gcc with `-ly -lm` flags