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
%token '(' ')' '=' '{' '}'

%left '+','-'
%left '*','/','%'

%type<val> expr

%start program

%%
program: stmts
	;

stmts: /* empty */
	| stmts stmt
	;

stmt: expr {
		fmt.Printf("Result: %d\n", $1)
	}

expr: NUMBER
	| expr '+' expr { $$ = $1 + $3 }
	| expr '-' expr { $$ = $1 - $3 }
	| expr '*' expr { $$ =  $1 * $3 }
	| expr '/' expr { $$ = $1 / $3 }
	| '(' expr ')' { $$ = $2 }
%%
