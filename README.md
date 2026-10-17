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

## Learning path

Each track is meant to be read in order:

* **flex/bison (C)**: `2020_1` (lexer only, what `yytext` and default ECHO
  mean) → `2020_1.5` (first grammar, `yylval`/`%union`) → `2020_2`
  (precedence, unary minus, error handling, a REPL).
* **goyacc (Go)**: `goyacc1` (grammar + lexer contract) → `goyacc2`
  (multiple statements) → `goyacc2b` (variables) → `goyacc2c` (AST +
  deferred evaluation) → `goyacc3` (typed AST via `%union`) → `goyacc4`
  (scopes, `while`, tests) → `goyacc5` (a real language on top of goyacc).
* **.NET**: `netyacc1` (`goyacc1` port) → `netyacc2` (compiler + VM).

## Prereqs

* C
  * bison & flex (on macOS, the Xcode command line tools provide them)
  * a C compiler (`clang` or `gcc`)
* Go
  * `go install golang.org/x/tools/cmd/goyacc@latest`
* .NET <https://github.com/ernstc/YaccLexTools>
  * `dotnet tool install dotnet-ylt --global`
    (only needed to regenerate the `.y`/`.lex` inputs; generated sources
    are committed)

## How to run

* Go (each directory is its own module)
  * `cd goyaccN && make` — generate the parser (goyacc) and build
  * `make test` — falsifiable assertions (exit non-zero on mismatch);
    `goyacc4`/`goyacc5` also run `go test`
  * `make run` — feed stdin to the binary (goyacc1–4)
  * goyacc5 needs file arguments and must run inside `goyacc5/`:
    `cd goyacc5 && make && ./prolog-interpreter -logLevel INFO test/test1.pro query/query1.pro`
* C
  * `cd 2020_2 && make && ./calc` (or `make run`; `2020_1` and `2020_1.5`
    also have `make run`)
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

## Learning notes (flex/bison)

* **The default action is ECHO**: a rule without `{...}` prints its match to
  stdout. Give every rule an explicit action (see `2020_1`).
* **`yylval` is the lexer→parser value channel**: the lexer stores a value in
  `yylval.<member>` and the grammar reads it as `$1`/`$2`... (see
  `2020_1.5`'s `%union { int int_value; }`).
* **`%prec UMINUS`** gives unary minus its own precedence: with
  `%right UMINUS` declared after `'*'`, `-2 + 3` evaluates as `(-2)+3 = 1`
  instead of `-(2+3) = -5`, and the natural `expr: '-' expr` rule stays
  conflict-free (`2020_2`).
* **`YYERROR` aborts the current parse**: without an `error` recovery rule,
  one division by zero ends the whole `yyparse()` — and with it the REPL
  (`2020_2` demonstrates this behavior).
* **`bison -v` writes `*.output`**: the C counterpart of goyacc's `y.output`,
  with the same shift/reduce and reduce/reduce conflict report.

## Useful commands

* Regenerate a parser by hand: `goyacc hoge.go.y` (the Makefiles do this
  for you), `bison -v -t calc.y`, `flex calc.l`
* .NET <https://github.com/ernstc/YaccLexTools>
  * `dotnet ylt add-parser -p <parserName>`
