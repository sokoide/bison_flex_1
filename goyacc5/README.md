# goyacc5

> 日本語版: [README_ja.md](README_ja.md)

## About

* A Prolog-like language on top of goyacc: facts, rules and queries with
  unification and depth-first backtracking.
* Same shape as the earlier steps — goyacc grammar → generated parser →
  hand-written lexer — but this time the result is a real (small) language.

## Build / Run / Test

```sh
make           # goyacc on pkg/prolog/grammar.go.y, build ./prolog-interpreter
make test      # go test -v ./... (unifier + lexer unit tests)
make testparser  # parse/evaluate each test/*.pro program on its own
make testquery   # for each query/queryN.pro: run it against test/testN.pro
make clean     # removes the binary, pkg/prolog/y.output and grammar.go
               # (make regenerates grammar.go automatically)
```

Test scripts honor `LOGLEVEL` (default `INFO`):

```sh
make testparser            # INFO logging
LOGLEVEL=DEBUG make testparser   # verbose lexer/parser traces
make testquery
LOGLEVEL=DEBUG make testquery
```

`make testquery` pairs the files **by number**: `query/query5.pro` is run
against `test/test5.pro`, so a missing or misnamed file fails the target
(`set -e` aborts on the first failure).

### Run it by hand — from this directory only

```sh
./prolog-interpreter -logLevel INFO test/test1.pro query/query1.pro
```

The interpreter loads `resource/builtin.pro` with a **path relative to the
current directory**, so it must be started inside `goyacc5/` (all `make`
targets run there for you). From the repo root you would get:

```
level=fatal msg="... failed to open resource/builtin.pro: no such file or directory"
```

`make run` starts the binary without arguments, which only prints
`Usage: prolog <source-file> <query-file>` — invoke the binary directly as
above (argument 0 = the program to load, argument 1 = the queries to run).

> **Transcript format**: the examples below show the output of an
> interactive terminal, where logrus prints `INFO[0000] ...`. When you
> redirect or pipe the output, the same log lines come out as
> `time="..." level=info msg="..."` instead — the content is identical.

## Files

* `pkg/prolog/grammar.go.y` ... grammar (input to goyacc); generates
  `grammar.go` (committed, so you can read the generated parser) and
  `y.output` (state machine / conflict report)
* `pkg/prolog/lexer.go` ... tokenizer; lexer errors are stored and returned
  through `Lexer.Err()` rather than panicking
* `pkg/prolog/term.go` ... term representation (atom, number, string,
  variable, compound term, list)
* `pkg/prolog/clause.go` ... fact and rule clauses
* `pkg/prolog/unifier.go` ... `unify()` plus backtracking over rule bodies
* `pkg/prolog/eval.go` ... `Load()` (parses builtins + program), `Query()`,
  solution printing
* `pkg/prolog/lexer_test.go`, `pkg/prolog/unifier_test.go` ... unit tests
* `cmd/prolog-interpreter/main.go` ... flags, loads the two files, runs the
  query
* `resource/builtin.pro` ... built-in predicates (`write/1`, `writeln/1`)
* `test/testN.pro` ... sample programs; `query/queryN.pro` ... queries run
  against the program with the same number

## How it works

1. `Load()` parses `resource/builtin.pro` and then the program file with
   the goyacc parser; the action
   `input: clause_list { yylex.(*lexer).program = $1 }` hands the clause
   list back through the lexer — the same trick as goyacc2c/3/4.
2. The query file is parsed the same way (queries are clauses too:
   `food(orange).` ends with a dot).
3. `Query()` walks the clauses top-down for each query term. `unify()`
   matches a goal against a clause head, producing a substitution
   (variable bindings); rule bodies are solved recursively, and every
   choice point keeps its alternatives so a later goal failing backtracks
   into them (depth-first).
4. Solutions are printed back by substituting bindings into the query
   (`X = 5`), and `writeln/1` inside a rule body has the side effect of
   writing output. Quirk worth studying: in `test5.pro`,
   `print_list([5,6,7,8,9])` prints every element and still ends up
   `false` — once `[H|T]` has consumed the list, the base case
   `print_list([]).` no longer unifies with the empty tail (empty-list
   unification only works literal-to-literal; see "List is partially
   supported" below), and the side effects are not undone.

## Grammar

* [Grammar file](./pkg/prolog/grammar.go.y), in outline (actions omitted):

  ```
  input       : /* empty */ | clause_list
  clause_list : clause | clause_list clause
  clause      : fact_clause | rule_clause
  fact_clause : term '.'
  rule_clause : term COLON_DASH term_list '.'        -- "head :- body."
  term_list   : term | term_list ',' term
  term        : IDENT | VAR | '_' | NUMBER_LITERAL | STRING_LITERAL
              | IDENT '(' term_list ')'              -- compound: f(A,B)
              | '[' term_list ']'                    -- [5,6,7]
              | '[' term_list '|' term ']'           -- [H|T]
              | '[' ']'                              -- []
              | BUILTIN_WRITE '(' term_list ')' | BUILTIN_NL
  ```

  Note that a list is represented three different ways internally
  (`Args`, `Head`/`Tail`, `IsEmpty`) — that shortcut is what makes list
  support "partial" (see the `print_list` quirk under How it works).

* Supported built-in predicates
  * write/1
  * writeln/1
* List is partially supported. see `test5.pro` and `query5.pro`

## TODO

* Many items are not implemented including...
  * comparison (> >= < =< =:= =\= or etc) operators
  * arithmetic (+ - * / ** // mod) operators
  * cut (!) operator
  * other built-in predicates
* The lexer already tokenizes `+ - * / = ; { }`,
  but no grammar rules accept them yet; using them in a .pro file is a
  syntax error until the grammar grows the corresponding rules.

## Examples

### Simple fact / rule evaluation

```sh
$ cat test/test1.pro
food(orange).
food(apple).

meal(X) :- food(X).

$ cat query/query1.pro
food(orange).
meal(orange).
food(car).

$ ./prolog-interpreter -logLevel INFO  test/test1.pro  query/query1.pro
INFO[0000] querying...
INFO[0000] food(orange). -> true
INFO[0000] meal(orange). -> true
INFO[0000] food(car). -> false
INFO[0000] query completed.
```

### Simple fact / rule evaluation with recursion

```sh
$ cat test/test2.pro
% facts
mammal(dog).
mammal(cat).
mammal(human).
fish(cod).
fish(swordfish).
fish(mackerel).

% rules
animal(X) :- mammal(X).
animal(X) :- fish(X).

print_if_mammal(X) :-
    mammal(X),
    write(X),
    writeln(" is mammal").

% end of file

$ cat query/query2.pro
mammal(dog).
fish(dog).
animal(dog).

mammal(cod).
fish(cod).
animal(cod).

print_if_mammal(dog).
print_if_mammal(cod).

$ ./prolog-interpreter -logLevel INFO  test/test2.pro  query/query2.pro
INFO[0000] querying...
INFO[0000] mammal(dog). -> true
INFO[0000] fish(dog). -> false
INFO[0000] animal(dog). -> true
INFO[0000] mammal(cod). -> false
INFO[0000] fish(cod). -> true
INFO[0000] animal(cod). -> true
dog is mammal
INFO[0000] print_if_mammal(dog). -> true
INFO[0000] print_if_mammal(cod). -> false
INFO[0000] query completed.
```

### Combined rules (grandparent)

```sh
$ cat test/test3.pro
father(ichiro, jiro).
father(jiro, saburo).
mother(alice, jiro).
mother(becky, saburo).

parent(X,Y) :- father(X,Y).
parent(X,Y) :- mother(X,Y).
grandparent(A, B) :- parent(A, C), parent(C, B).

% end of file

$ cat query/query3.pro
parent(ichiro, jiro).
parent(alice, jiro).
parent(becky, jiro).
grandparent(ichiro, saburo).
grandparent(alice, saburo).
grandparent(alice, jiro).

$ ./prolog-interpreter -logLevel INFO  test/test3.pro  query/query3.pro
INFO[0000] querying...
INFO[0000] parent(ichiro, jiro). -> true
INFO[0000] parent(alice, jiro). -> true
INFO[0000] parent(becky, jiro). -> false
INFO[0000] grandparent(ichiro, saburo). -> true
INFO[0000] grandparent(alice, saburo). -> true
INFO[0000] grandparent(alice, jiro). -> false
INFO[0000] query completed.
```

### List

```sh
$ cat test/test5.pro
first([X|_], X).
second([_,X|_], X).

print_list([]).

print_list([H|T]) :-
    writeln(H),
    print_list(T).

$ cat query/query5.pro
first([5,6,7], X).
second([5,6,7], X).
first([hello,world], Y).
print_list([5,6,7,8,9]).

$ ./prolog-interpreter -logLevel INFO  test/test5.pro query/query5.pro
INFO[0000] querying...
X = 5
INFO[0000] first([5, 6, 7], X). -> true
X = 6
INFO[0000] second([5, 6, 7], X). -> true
Y = hello
INFO[0000] first([hello, world], Y). -> true
5
6
7
8
9
INFO[0000] print_list([5, 6, 7, 8, 9]). -> false
INFO[0000] query completed.
```
