# Goyacc3

* Goyacc calculator: expression AST + evaluation in Go

## Files

* calc.go.y ... grammar file; generates y.go (parser) and y.output (state machine dump)
* y.go, y.output ... generated parser and its conflict/status report
* lexer.go ... hand-written scanner. `Lex()` is called by the parser to get
  the next token and `Error()` is called on a syntax problem
* ast.go ... AST node types (number / paren / binary operator)
* eval.go ... tree-walking evaluator
* main.go ... reads stdin as one expression, parses it, then evaluates it

## Build / Run / Test

```
make          # build ./calc (runs goyacc first if y.go is out of date)
make run      # read a program from stdin
make test     # exact assertions, exits non-zero on mismatch
make clean    # removes ./calc y.go y.output (make regenerates them)
```

Example:

```
$ printf '(1*2)+10/3\n' | ./calc
5
$ printf '1+@\n' | ./calc
2026/10/17 12:45:37 parse error: Line 1, Column 3: "@" syntax error: unexpected ILLEGAL, expecting NUMBER or '('
(exit 1; the date/time prefix comes from Go's default logger)
```

goyacc3 is the Go counterpart of the C calculator in [2020_2](../2020_2);
reading the two side by side shows the same precedence/AST ideas in both
toolchains.

## What is this?

1. `%union` carries values between the parser and the semantic actions.
   Here it holds `expr of type expression` (an AST node) and `tok of type token`.

2. `program: expr { $$ = $1; yylex.(*lexer).program = $$ }` hands the finished
   AST back to the caller: goyacc only exposes reductions through the
   `yySymType` union, so the action stores the root node on the lexer, from
   where `parse()` returns it to `main`.

3. `%left '+' '-'` / `%left '*' '/' '%'` resolve the ambiguity of the natural
   `expr op expr` rules. Without them grammar is ambiguous and goyacc reports
   shift/reduce conflicts; later precedence declarations bind more tightly,
   and operators at the same level associate left. See `y.output` for the
   generated state machine and the line
   `0 shift/reduce, 0 reduce/reduce conflicts reported`.

4. `yyErrorVerbose = true` makes the parser say `unexpected ILLEGAL` instead
   of `syntax error`, which is a great learning tool together with `yyDebug`.

5. Lexer contract: `Lex()` returning 0 (or a negative value) means end-of-input
   to goyacc, so an unknown character must NOT return token number 0. This
   example returns the `ILLEGAL` token for unknown characters, which makes the
   parser fail with a visible error instead of silently dropping the rest of
   the line.

Progression: goyacc3 evaluates a single expression and prints its value, e.g.
`(1*2)+10/3` prints `5`, `(40+1)+2/2` prints `42`, `3%4` prints `3`. Variables
and statements come next in goyacc4.
