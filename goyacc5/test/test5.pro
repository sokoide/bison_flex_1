first([X|_], X).
% second([_,X|_], X).

print_list([]).

print_list([H|T]) :-
    writeln(H),
    print_list(T).
