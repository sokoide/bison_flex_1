% facts
father(jiro, saburo).
mother(becky, saburo).
father(ichiro, jiro).
mother(alice, jiro).

% rules
parent(X,Y) :- father(X,Y).
parent(X,Y) :- mother(X,Y).
grandparent(A, B) :- parent(A, C), parent(C, B).

% end of file