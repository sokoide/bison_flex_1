%{
package main
import (
	"fmt"
	)
%}

%union{
	val int
	ident string
	token int
}

%type<val> program line expr let
%token<val> NUMBER LF PUT
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
	}
	| PUT '(' expr ')' LF {
		fmt.Printf("%d\n", $3)
	}

let: IDENT '=' expr { vars[$1] = $3 }

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
