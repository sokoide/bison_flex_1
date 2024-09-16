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