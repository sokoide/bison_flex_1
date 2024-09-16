write(X) :- builtin_write(X).

writeln(X) :- builtin_write(X), builtin_nl.
