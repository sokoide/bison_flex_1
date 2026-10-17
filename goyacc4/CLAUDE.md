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
- Arithmetic: `+`, `-`, `*`, `/`, `%` with parentheses
- Conditionals: `==`, `!=`, `>=`, `>`, `<=`, `<`
- While loops: `while(condition){statements}`
- Output: `put("text", variable);`

Known limitations:
- No unary minus (`put(-1);` is a syntax error)
- Strings are only valid as direct `put()` literals; anywhere else
  (`"a"+"b"`, `a="x";`, `"a"=="b"`) is an evaluation error
- Reading an undefined variable yields 0 (auto-created in the current scope)

`make clean` deletes `grammar.go`; regenerate it with `make generate`
before a plain `go build` (make does this automatically via its
`grammar.go:` rule). Parse and lex errors are returned as Go errors from
`parse()` — the lexer stores the message instead of exiting — so they can
be unit-tested; evaluation errors are handled by `main` with a non-zero
exit status.

## Testing

Uses testify/assert for unit tests. Tests cover parsing, evaluation, while loops, and variable scoping behavior.