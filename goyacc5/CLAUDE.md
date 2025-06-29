# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Prolog-like language interpreter implemented in Go, using goyacc (Go version of yacc) for parser generation. The interpreter can evaluate Prolog facts, rules, and queries with support for unification and backtracking.

## Core Architecture

- **Parser Generation**: Uses goyacc to generate parser from `pkg/prolog/grammar.go.y`
- **Lexer**: Hand-written lexer in `pkg/prolog/lexer.go` 
- **Evaluation Engine**: Core logic in `pkg/prolog/eval.go` with unification in `pkg/prolog/unifier.go`
- **Term System**: Various term types (constants, variables, compounds, lists) defined in `pkg/prolog/term.go`
- **Built-ins**: Minimal built-in predicates (write/1, writeln/1) loaded from `resource/builtin.pro`

## Build Commands

```bash
# Build the interpreter
make

# Generate parser from grammar (runs automatically during build)
make generate

# Clean generated files
make clean

# Run tests
make test
go test -v ./...

# Test all .pro files in test/ directory
make testparser

# Test query files against corresponding test files
make testquery

# Debug specific test/query pair
make debug
```

## Development Workflow

1. **Grammar Changes**: Edit `pkg/prolog/grammar.go.y`, then run `make generate` to regenerate parser
2. **Parser Generation**: The `//go:generate goyacc -o grammar.go grammar.go.y` directive in `pkg/prolog/generate.go` handles code generation
3. **Testing**: Use matching pairs of `test/testN.pro` and `query/queryN.pro` files for testing

## Running the Interpreter

```bash
# Basic usage
./prolog-interpreter <facts_file> <queries_file>

# With logging levels
./prolog-interpreter -logLevel DEBUG <facts_file> <queries_file>
./prolog-interpreter -logLevel INFO <facts_file> <queries_file>

# Dump parser state
./prolog-interpreter -dump <facts_file> <queries_file>
```

## Key Implementation Details

- **Unification Algorithm**: Central to Prolog evaluation, implemented in `unifier.go`
- **Clause Types**: Facts (simple assertions) and Rules (head :- body structure)
- **Term Types**: Constants, Variables, Compounds, Lists, Anonymous variables
- **Query Evaluation**: Backtracking search through rule space with substitution tracking
- **Built-in Integration**: Built-in predicates loaded first, then user program appended

## Current Limitations

- No arithmetic operators (+, -, *, /, etc.)
- No comparison operators (>, <, =, etc.) 
- No cut (!) operator
- Limited built-in predicates (only write/1, writeln/1)
- Partial list support