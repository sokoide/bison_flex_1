# goyacc5

## About

* Prolog like language
* Please look at the 2 test scripts

```sh
# INFO logging (default)
make testparser
make testquery
# DEBUG logging (verbose lexer/parser traces)
LOGLEVEL=DEBUG make testparser
LOGLEVEL=DEBUG make testquery
```

## Grammar

* [Please look at the grammar file](<./pkg/prolog/grammar.go.y>)
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
