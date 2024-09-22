%{
package main

import "fmt"
%}

// type of $$ (return value of AST)
%union{
	val int
}

// $$ of program and expr are both 'val'
%type<val> program expr

// define all supported tokens
%token<val> NUMBER

// lower line (* and /) takes precedence
%left '+' '-'
%left '*' '/'

%start program

%%
program: expr {
        fmt.Printf("Result: %d\n", $1)
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
