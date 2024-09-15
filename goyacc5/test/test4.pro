print_list([]).

print_list([H|T]) :-
    writeln(H),
    print_list(T).

% end of file