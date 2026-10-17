# 2020_1.5 — flex + bison "I LIKE <object>" demo

The first combined flex + bison example: a tiny grammar that accepts
sentences of the form `SUBJECT VERB OBJECT` and prints the object's
value.

## What it shows

- The lexer returning token codes (`SUBJECT`, `VERB`, `OBJECT`,
  `INTEGER`, `LF`) that bison parses
- Passing a semantic value from lexer to parser through `yylval`
  (`%union { int int_value; }`): `OBJECT` always carries the value 42,
  `INTEGER` carries the number it matched
- A two-rule grammar with a start symbol that handles any number of
  lines

## The grammar

    line:  LF                 { exit(0); }        -- a blank line ends the demo
        | SUBJECT VERB obj LF { printf("result: %d\n", $3); }
    obj  : OBJECT   {$$ = $1;}
         | INTEGER  {$$ = $1;}

The lexer hardcodes the words: `I` = SUBJECT, `LIKE` = VERB,
`SOCCER` = OBJECT. Whitespace is skipped; any other character is
reported to stderr as `Unknown character: '<c>'`.

## Build and run

    make
    ./calc

## Example input/output

    $ printf 'I LIKE SOCCER\n' | ./calc
    result: 42
    $ printf 'I LIKE 3\n' | ./calc
    result: 3

## Invalid input

Input not matching the grammar produces a parse error on stderr, for
example:

    $ printf 'dog barks 3\n' | ./calc
    Unknown character: 'd'
    Unknown character: 'o'
    Unknown character: 'g'
    Unknown character: 'b'
    Unknown character: 'a'
    Unknown character: 'r'
    Unknown character: 'k'
    Unknown character: 's'
    parse error near 3

`make clean` removes `calc`, `calc.tab.c`, and `lex.yy.c`.
