# Goyacc4

* Goyacc mini-compiler: statement AST, while loops and a stack of scopes

## Files

* grammar.go.y ... grammar file; generates grammar.go and y.output via `go generate`
* grammar.go, y.output ... generated parser and its state machine dump
* lexer.go ... hand-written scanner. `Lex()` is called by the parser to get
  the next token and `Error()` is called on a syntax problem
* ast.go ... expression AND statement AST node types
* eval.go ... tree-walking interpreter for statements and expressions
* scope.go ... stack of scopes used when evaluating while-loop bodies
* main.go ... reads a program from stdin, parses and evaluates it

## Build / Run / Test

```
make generate  # run goyacc on grammar.go.y -> grammar.go
make           # build ./mini-compiler (grammar.go must exist before go build)
make test      # go test -v ./...
make testcli   # exact assertions against the compiled binary
make clean     # removes ./mini-compiler grammar.go y.output
```

The binary takes `-logLevel (TRACE|DEBUG|INFO|WARN|ERROR)`: `put()` output
goes to stdout, all log lines to stderr, so `-logLevel ERROR` keeps
transcripts clean.

Example (note the scope rule: `b` is created inside the loop and vanishes
after it, so `a+b` is `5+0`):

```
$ printf 'a=0;while(a<5){a=a+1;b=b+1;};put(a+b);\n' | ./mini-compiler -logLevel ERROR
5
$ printf 'put(1/0);\n' | ./mini-compiler -logLevel ERROR
time="..." level=error msg="[stmt 0] evaluation error: division by zero"
time="..." level=error msg="Evaluation error: division by zero"
(exit 1)
```

## What is this?

Similar to goyacc3, but each grammar symbol carries whole statements,
not a single expression:

1. `%union` holds `exprs []expression` (lists), `expr expression` (one node)
   and `tok token`. `%type<exprs> program stmts put_list` types the
   non-terminals.

2. `program: stmts { $$ = $1; yylex.(*lexer).program = $$ }` hands the
   finished statement list back to the caller through the lexer, from where
   `parse() <- ([]expression, error)` returns it to `main`.

3. `%left '+' '-'` / `%left '*' '/' '%'` keep the ambiguous `expr op expr`
   grammar conflict-free (`0 shift/reduce, 0 reduce/reduce conflicts
   reported` in y.output).

4. The language: assignment `a=1;`, output `put("x=",expr,...);`,
   `while(cond){...}` with `{...}` introducing a new scope, `;` as an empty
   statement, and comparisons `== != >= > <= <` that only appear inside
   `while(...)` (`cond`).

5. `scope.go` implements a stack of scopes: `while` bodies push a scope and
   pop it when the loop ends, so variables created inside the loop
   disappear afterwards (`b` in `while(a<5){a=a+1;b=b+1;}`).

6. Error handling: `yyErrorVerbose = true` gives messages such as
   `unexpected ILLEGAL`. Checking an unknown character must not end the
   parse is a classic goyacc pitfall: `Lex()` returning 0 or a negative
   value means end-of-input, so the scanner returns the `ILLEGAL` token
   instead. Lexer and parser errors are returned as Go errors (the lexer
   stores the message instead of calling `os.Exit`), and evaluation errors
   (division by zero, strings used outside `put()`) abort the program with a
   non-zero exit status.

7. Known limitations: no unary minus (`put(-1);` is a syntax error), strings
   are only valid as direct `put()` literals (so `"a"+"b"`, `a="x";` and
   `"a"=="b"` are evaluation errors), and reading an undefined variable yields
   0 (auto-created in the current scope).

`yyDebug` levels (set `yyDebug = N` in the generated parser) print the
parser's state transitions, another useful learning tool together with
`y.output`.
