%{
package main
// import (
// 	"fmt"
// 	)
%}

%union{
	val int
	ident string
	nodes []*Node
	node *Node
}

%token<val> NUMBER
%token<ident> IDENT
%token '(' ')' '=' '{' '}' ';' PUT IF

%left '+','-'
%left '*','/','%'

%type<node> expr stmt
%type<nodes> program stmts

%start program

%%
program: stmts {
	$$ = $1
	yylex.(*lexer).program = $$
	};

stmts: /* empty */ {
		$$ = nil
	}
	| stmts stmt {
		$$ = append($1, $2)
	}
	;

stmt: IDENT '=' expr ';' {
		$$ = newNode("assign", $1, 0, $3, nil, nil)
	}
	| PUT '(' expr ')' ';'{
		$$ = newNode("put", "", 0, $3, nil, nil)
	}
	| IF '(' expr ')' '{' stmts '}' {
		$$ = newNode("if", "", 0, $3, nil, $6)
	}
	;

expr: NUMBER { $$ = newNode("int", "", $1, nil, nil, nil) }
	| IDENT { $$ = newNode("ident", $1, 0, nil, nil, nil) }
	| expr '+' expr { $$ = newNode("+", "", 0, $1, $3, nil) }
	| expr '-' expr { $$ = newNode("-", "", 0, $1, $3, nil) }
	| expr '*' expr { $$ = newNode("*", "", 0, $1, $3, nil) }
	| expr '/' expr { $$ = newNode("/", "", 0, $1, $3, nil) }
	| expr '%' expr { $$ = newNode("%", "", 0, $1, $3, nil) }
	| '(' expr ')' { $$ = $2 }
%%

// global vars
var vars = map[string]int{}
