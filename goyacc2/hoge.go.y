%{
package main
import (
	"fmt"
	)
%}

%union{
	val int
	ident string
}

%token<val> NUMBER
%token<ident> IDENT
%token '(' ')' '=' '{' '}' LF

%left '+','-'
%left '*','/','%'

%type<val> expr

%start program

%%
program: stmts
	;

stmts: /* empty */ {
	}
	| stmts expr {
		fmt.Printf("Result: %d\n", $2)
	}

expr: NUMBER
	| expr '+' expr { $$ = $1 + $3 }
	| expr '-' expr { $$ = $1 - $3 }
	| expr '*' expr { $$ =  $1 * $3 }
	| expr '/' expr { $$ = $1 / $3 }
	| expr '%' expr { $$ = $1 % $3 }
	| '(' expr ')' { $$ = $2 }
%%
