# bison_flex_1

Progressive, minimal examples for learning lexer/parser generators:
**flex + bison** (C), **goyacc** (Go), and **YaccLexTools** (C#).

Each example is deliberately tiny and builds on the previous one, so you can
read them in order and watch a lexer/parser pair grow into a small interpreter.

## Examples

| Directory  | Tool             | What it shows                                                        |
| ---------- | ---------------- | -------------------------------------------------------------------- |
| `2020_1`   | flex             | flex only: tokenize integers/operators and print the tokens           |
| `2020_1.5` | flex + bison     | first grammar: a tiny "SUBJECT VERB OBJECT" language                  |
| `2020_2`   | flex + bison     | integer calculator: precedence, `%prec UMINUS`, `YYERROR`, REPL       |
| `goyacc1`  | goyacc           | smallest goyacc parser: an arithmetic expression evaluator            |
| `goyacc2`  | goyacc           | statements, whitespace handling, `%`                                  |
| `goyacc2b` | goyacc           | variables and a `put(...)` statement                                 |
| `goyacc2c` | goyacc           | builds an AST and adds `if` (evaluation happens after parsing)        |
| `goyacc3`  | goyacc           | expression AST carried through `%union` (interface type)              |
| `goyacc4`  | goyacc           | statements, `while` loops, scoped variables, unit tests               |
| `goyacc5`  | goyacc           | Prolog-like interpreter (unification + backtracking)                  |
| `netyacc1` | YaccLexTools     | .NET port of `goyacc1`                                                |
| `netyacc2` | YaccLexTools     | .NET compiler/interpreter for a small language                        |

Every example directory (C, Go, .NET) has its own `README.md` with build/run
instructions.

## Prereqs

* C
  * bison & flex (on macOS, the Xcode command line tools provide them)
* Go
  * `go install golang.org/x/tools/cmd/goyacc@latest`
* .NET <https://github.com/ernstc/YaccLexTools>
  * `dotnet tool install dotnet-ylt --global`

## How to run

* Go (each directory is its own module)
  * `cd goyaccN && make test` (assertions) and `make run`
* C
  * `cd 2020_2 && make && ./calc`
* .NET
  * Please refer to [netyacc1/README.md](netyacc1/README.md) and
    [netyacc2/README.md](netyacc2/README.md)

## Learning notes (goyacc)

* **The `Lex()` contract**: return the token number, and return `0` (or a
  negative value) at end of input. goyacc treats `0`/negative as the EOF token
  (`yyTok1[0]`), so *never* return `0` for a real token — an "unknown" character
  must be given a declared token (see `ILLEGAL` in `goyacc3`/`goyacc4`), or the
  parser will silently stop there.
* **`Error()`** is called by goyacc for *syntax* errors. What happens on a
  *lexical* error is up to the lexer (panic, skip, or produce an error token).
* **`yyErrorVerbose = true`** makes error messages name the unexpected and
  expected tokens, e.g. `syntax error: unexpected $end, expecting NUMBER or '('`.
  This is the single most useful switch while learning.
* **`yyDebug = 1..4`** prints the parser's shift/reduce trace.
* **`y.output`** is the generated parser's state table. Look at the
  `N shift/reduce, M reduce/reduce conflicts reported` line: an ambiguous
  grammar like `expr '/' expr` is made conflict-free with `%left`/`%right`
  precedence declarations.
* Generated parser files are produced by `make` from the `.y` file. Some of
  them (`goyacc4/grammar.go`, `goyacc5/pkg/prolog/grammar.go`) are committed so
  you can read the generated code; the others (`y.go`, `*.tab.c`, `y.output`)
  are gitignored and regenerated on demand. `make clean` removes them.

## Useful commands

* .NET <https://github.com/ernstc/YaccLexTools>
  * `dotnet ylt add-parser -p <parserName>`
