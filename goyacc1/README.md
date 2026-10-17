# Goyacc1

* Simplest Goyacc example: an integer expression evaluator

## Build / Run / Test

```
make          # runs goyacc on hoge.go.y -> y.go, then builds ./hoge
make test     # exact assertions, exits non-zero on mismatch
make run      # read an expression from stdin
make clean    # removes ./hoge y.go y.output (make regenerates them)
```

Example:

```
$ printf '4+1*4/2\n' | ./hoge
Result: 6
$ printf '1+\n' | ./hoge
error: syntax error, at token: +
parse failed     (exit 1)
```

## Files

* `hoge.go.y` ... grammar file (the input to goyacc)
* `y.go` ... generated from `hoge.go.y` by `goyacc` (gitignored; `make`
  regenerates it)
* `y.output` ... generated parser's state table / conflict report
  (gitignored)
* `lexer.go` ... hand-written lexer
  * `Lex()` is called by the parser to get the next token.
    Returning 0 (or any negative number, here -1) signals EOF to goyacc.
  * `Error()` is called by goyacc for SYNTAX errors reported by the parser
    (for example `1+`). Note that this is NOT called for lexical problems:
    what happens on a bad character is decided by the lexer author, and this
    lexer simply panics on one.
  * `yyErrorVerbose = true` (a generated global) makes the parser include
    the unexpected/expected token names in the error string; `yyDebug = 1..4`
    traces each parser step on stderr. Both can be set at runtime.
    This example leaves `yyErrorVerbose` at its default (`false`), which is
    why the message above is the plain `syntax error` form;
    [goyacc2](../goyacc2) switches it on so you can compare the two.
* `main.go` ... reads stdin, feeds it to the lexer, runs `yyParse`

## Grammar in one line

`program: expr` — one expression per run.

## `%token` and `%union`

* `%union{ val int }` defines the Go type that replaces yacc's `YYSTYPE`.
* `%token<val> NUMBER` declares a token and says its semantic value lives in
  `lval.val` (the `%union` member `val`).
* `%left '+' '-'` and `%left '*' '/'` give the '*' '/' group higher
  precedence, so `4+1*4/2` parses as `4+((1*4)/2)`.

## Walking through `hoge.go.y`

1. Package and imports:
   the `%{ ... %}` block at the top of the `.y` file is pasted verbatim
   into the generated `y.go`, so the generated parser lands in `package main`
   with the imports it needs (`fmt` for the action below).

2. Union declaration:
   `%union` defines the types that can be associated with grammar symbols.
   Here we only have `val` of type `int`.

3. Token and type declarations:
   `%token <val> NUMBER` declares a token for numeric values;
   `%type <val> expr` says that the `expr` non-terminal has an integer value.

4. Operator precedence:
   `%left '+' '-'` and `%left '*' '/'` — later declarations bind tighter,
   which is what makes `4+1*4/2` come out as 6.

5. Grammar rules:
   the rules are similar to the C version (see `2020_2`), but the actions
   are Go: each assigns to `$$` and references `$1`, `$3`, ..., as in C yacc.
   Actions run DURING parsing (see [goyacc2c](../goyacc2c)'s README for an
   AST alternative).

6. Lexer implementation:
   the `lexer` struct and its methods do the tokenization — `Lex()` reads
   the input and returns tokens, `Error()` reports syntax errors.

7. Entry point (`main.go`): reads all of stdin into the lexer, runs
   `yyParse`, and exits 1 with `parse failed` if parsing failed. Note that
   `main` never sees the value — the `program: expr` action prints
   `Result: %d` itself, DURING parsing. Compare [goyacc2c](../goyacc2c),
   which instead hands the finished tree back to `main` and evaluates it
   after `yyParse` returns.
