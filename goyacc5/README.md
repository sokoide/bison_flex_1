# goyacc5

## About

* Prolog like language
* Please look at the 2 test scripts

```sh
make testparser
make testqquery
```

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
INFO[0000] result: food(orange). -> true
INFO[0000] result: meal(orange). -> true
INFO[0000] result: food(car). -> false
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

% end of file

$ cat query/query2.pro
mammal(dog).
fish(dog).
animal(dog).

mammal(cod).
fish(cod).
animal(cod).

$ ./prolog-interpreter -logLevel INFO  test/test2.pro  query/query2.pro
INFO[0000] querying...
INFO[0000] result: mammal(dog). -> true
INFO[0000] result: fish(dog). -> false
INFO[0000] result: animal(dog). -> true
INFO[0000] result: mammal(cod). -> false
INFO[0000] result: fish(cod). -> true
INFO[0000] result: animal(cod). -> true
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