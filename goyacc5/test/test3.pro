father(ichiro, jiro).
father(jiro, saburo).
mother(alice, jiro).
mother(becky, saburo).

parent(X,Y) :- father(X,Y).
parent(X,Y) :- mother(X,Y).
grandparent(A, B) :- parent(A, C), parent(C, B).

% end of file