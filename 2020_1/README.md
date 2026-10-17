# 2020_1 — flex-only token printer

The simplest example in this repository: a **pure flex** lexer (no
bison/yacc involved). It recognizes input tokens and prints what it
matched.

## What it shows

- Defining named patterns in flex (`WHITESPACE`, `INTEGER`, `SYMBOL`, `LF`)
- Rule actions printing `yytext`, the text of the current match
- Why every rule needs an explicit action: a rule with no action gets
  flex's **default action (ECHO)**, which would print the match to
  stdout. Here the whitespace, newline, and catch-all `.` rules all
  have explicit empty actions so only tokens are printed.

## Pieces

- `calc.l` — the flex lexer specification (contains `main()`)
- `Makefile` — runs `flex calc.l` and compiles `lex.yy.c` with `-ly -lm`

## Build and run

    make
    ./calc

## Example

    $ printf '12 + 3\n' | ./calc
    INTEGER: 12
    SYMBOL: +
    INTEGER: 3

Spaces, newlines, and unknown characters are skipped silently.
`make clean` removes `calc` and `lex.yy.c`.
