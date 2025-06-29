# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a Go-based mini-compiler that implements a simple programming language with variables, arithmetic expressions, while loops, and a `put` statement for output. The compiler uses goyacc (Go's yacc implementation) for parsing and includes a hand-written lexer.

## Architecture

- **Parser Generation**: The grammar is defined in `grammar.go.y` and compiled to `grammar.go` using goyacc
- **Lexical Analysis**: Hand-written scanner in `lexer.go` that tokenizes the input
- **AST**: Expression and statement types defined in `ast.go`
- **Evaluation**: Tree-walking interpreter in `eval.go`
- **Scoping**: Variable scoping system in `scope.go` with stack-based scope management
- **Entry Point**: `main.go` coordinates parsing and evaluation, reads from stdin

## Build Commands

```bash
# Generate parser from grammar
make generate
# or directly:
go generate ./

# Build the compiler
make all
# or:
go build -o mini-compiler

# Clean build artifacts
make clean

# Run tests
make test
# or:
go test -v ./...

# Run CLI tests with sample programs
make testcli
```

## Development Workflow

1. Modify `grammar.go.y` for grammar changes
2. Run `make generate` to regenerate `grammar.go`
3. Build with `make all`
4. Test with `make test` and `make testcli`

The language supports:
- Variable assignment: `a=42;`
- Arithmetic: `+`, `-`, `*`, `/` with parentheses
- Conditionals: `==`, `!=`, `>=`, `>`, `<=`, `<`
- While loops: `while(condition){statements}`
- Output: `put("text", variable);`

## Testing

Uses testify/assert for unit tests. Tests cover parsing, evaluation, while loops, and variable scoping behavior.