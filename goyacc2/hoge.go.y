%{
package main
import (
	"fmt"
	)
%}

%union{
	num int
	ident string
	token int
}

%type<num> program line expr let
%token<num> NUMBER LF
%token<ident> IDENT
%token<token> '(',')','='

%left '+','-'
%left '*','/','%'

%start program

%%
program: line
	| program line

line: let LF {$$ = $1}
	 | expr LF {
		$$ = $1
		fmt.Println("Result: ", $1)
	}

let: IDENT '=' expr { vars[$1] = $3 }

expr: NUMBER
	| IDENT { $$ = vars[$1] }
	| expr '+' expr { $$ = $1 + $3 }
	| expr '-' expr { $$ = $1 - $3 }
	| expr '*' expr { $$ =  $1 * $3 }
	| expr '/' expr { $$ = $1 / $3 }
	| expr '%' expr { $$ = int($1) % int($3) }
	| '(' expr ')' { $$ = $2 }

%%

// global vars
var vars = map[string]int{}
