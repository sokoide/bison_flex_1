%{
package main
%}

%union{
	num int
}

%type<num> program expr
%token<num> NUMBER

%left '+' '-'
%left '*' '/'

%start program

%%
program: expr {
		$$ = $1
        yylex.(*lexer).result = $$
	}

expr: NUMBER
    | expr '+' expr {
        $$ = $1 + $3
	}
    | expr '-' expr {
        $$ = $1 - $3
	}
    | expr '*' expr {
        $$ =  $1 * $3
	}
    | expr '/' expr {
        $$ = $1 / $3
	}
    | '(' expr ')' {
        $$ = $2
    }

%%
