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
%token '(' ')' '=' '{' '}' ';' PUT

%left '+','-'
%left '*','/','%'

%type<val> expr

%start program

%%
program: stmts
	;

stmts: /* empty */
	| stmts stmt ';'
	;

stmt: IDENT '=' expr {
		vars[$1] = $3
	}
	| PUT '(' expr ')' {
		fmt.Printf("%d\n", $3)
	}
	;

expr: NUMBER
	| IDENT { $$ = vars[$1] }
	| expr '+' expr { $$ = $1 + $3 }
	| expr '-' expr { $$ = $1 - $3 }
	| expr '*' expr { $$ =  $1 * $3 }
	| expr '/' expr { $$ = $1 / $3 }
	| expr '%' expr { $$ = $1 % $3 }
	| '(' expr ')' { $$ = $2 }
%%

// global vars
var vars = map[string]int{}
