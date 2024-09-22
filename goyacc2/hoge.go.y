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
%token<val> NUMBER LF
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
	| expr '%' expr { $$ = $1 % $3 }
	| '(' expr ')' { $$ = $2 }

	// | expr '%' expr { $$ = int($1) % int($3) }
%%

// global vars
var vars = map[string]int{}
