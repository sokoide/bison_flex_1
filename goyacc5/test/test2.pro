% facts
mammal(dog).
mammal(cat).
mammal(human).
fish(cod).
fish(swordfish).
fish(mackerel).

father(jiro, saburo).
mother(becky, saburo).
father(ichiro, jiro).
mother(alice, jiro).

% rules
animal(X) :- mammal(X).
animal(X) :- fish(X).

parent(X,Y) :- father(X,Y).
parent(X,Y) :- mother(X,Y).
grandparent(A, B) :- parent(A, C), parent(C, B).


% end of file