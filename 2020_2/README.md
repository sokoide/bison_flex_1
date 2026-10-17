# 2020_2 — flex + bison integer calculator

A REPL calculator for integer arithmetic: `+ - * /`, parentheses, and
unary minus, evaluated line by line.

## What it shows

- Operator precedence declarations in bison:
  `%left '+' '-'` / `%left '*' '/'` / `%right UMINUS`
- Unary minus with `%prec UMINUS`: `- expr %prec UMINUS { $$ = -$2; }`
- Lexer-to-parser values via `%union { int int_value; }` and
  `yylval.int_value` set in `calc.l` for every integer token
- Error handling with `YYERROR`: the division rule detects a zero
  divisor, calls `yyerror("division by zero")`, and aborts the
  current parse with `YYERROR`

## Pieces

- `calc.y` — grammar, precedence declarations, error reporting (`yyerror`)
- `calc.l` — lexer: integers, operators, `LF` tokens; whitespace is
  skipped, unknown characters are reported to stderr
- `calc.y` includes `lex.yy.c` directly via `#include "lex.yy.c"`

## Build and run

    make
    ./calc

or `make run`. Other targets: `make debug` (rebuilds with `-g -DDEBUG`)
and `make help` (lists the targets).

## Example

    $ printf '1 + 2\n(40+1)+2/2\n1/0\n' | ./calc
    Simple Calculator
    Enter expressions (Ctrl+D to exit):
    3
    42
    Error: division by zero

    Goodbye!

(The `Error: division by zero` line goes to stderr, so its position in
the transcript can differ when stdout is buffered.)

Note: `1/0` does not just skip the current line — because the grammar
has no `error` recovery rule, `YYERROR` aborts the entire `yyparse()`
and the program ends:

    $ printf '1/0\n2+2\n' | ./calc
    Simple Calculator
    Enter expressions (Ctrl+D to exit):
    Error: division by zero

    Goodbye!
    (exit code 1; the "2+2" line is never evaluated)

`make clean` removes `calc`, `calc.tab.c`, `calc.tab.h`, `lex.yy.c`,
and `calc.output`.
