%{
package main

import (
	log "github.com/sirupsen/logrus"
)
%}

%union{
	stmts   []statement
	stmt	statement
	exprs	[]expression
	expr	expression
	tok	 	token
}

%type<stmts> program stmts
%type<stmt> stmt
%type<exprs> put_list
%type<expr> expr
%token<tok> PUT
%token<tok> NUMBER_LITERAL IDENT STRING_LITERAL

%left '+' '-'
%left '*' '/'

%start program

%%
program:
	stmts {
		$$ = $1
		yylex.(*lexer).program = $$
	}
	;

stmts:
	/* empty */ {
		log.Info("stmts: empty");
		$$ = nil
	}
	| stmts stmt {
		log.Infof("stmts: stmt %v", $2);
		$$ = append($$, $2)
	}
	;

stmt:
	IDENT '=' expr ';' {
		$$ = &assignStatement{Name: $1.lit, Expr: $3}
	}
	| PUT '(' put_list ')' ';' {
		$$ = &putStatement{Exprs: $3}
	}
	| expr ';' {
		$$ = &exprStatement{Expr: $1}
	}
	;

put_list: expr {
		$$ = append($$, $1)
	}
	| put_list ',' expr {
		$$ = append($$, $3)
	}
	;

expr:
	NUMBER_LITERAL {
		$$ = &numberExpression{Lit: $1.lit}
	}
	| IDENT {
		 $$ = &variableExpression{Lit: $1.lit}
	}
	| expr '+' expr {
		$$ = &binOpExpression{LHS: $1, Operator: int('+'), RHS: $3}
	}
	| expr '-' expr {
		$$ = &binOpExpression{LHS: $1, Operator: int('-'), RHS: $3}
	}
	| expr '*' expr {
		$$ = &binOpExpression{LHS: $1, Operator: int('*'), RHS: $3}
	}
	| expr '/' expr {
		$$ = &binOpExpression{LHS: $1, Operator: int('/'), RHS: $3}
	}
	| '(' expr ')' {
		$$ = &parenExpression{SubExpr: $2}
	}
	;

%%
// global vars
var vars = map[string]int{}
