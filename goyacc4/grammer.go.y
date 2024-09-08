%{
package main

import (
	log "github.com/sirupsen/logrus"
)
%}

%union{
	stmts		[]statement
	stmt		statement
	exprs		[]expression
	expr		expression
	tok	 		token
}

%type<stmts> program stmts
%type<stmt> stmt
%type<exprs> put_list
%type<expr> expr cond while_prefix

%token<tok> PUT WHILE
%token<tok> EQOP, NEOP, GEOP, GTOP, LEOP, LTOP
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
		log.Debug("stmts: empty");
		$$ = nil
	}
	| stmts stmt {
		log.Debugf("stmts: stmt %v", $2);
		$$ = append($$, $2)
	}
	;

stmt:
	IDENT '=' expr ';' {
		$$ = &assignStatement{Name: $1.lit, Expr: $3}
	}
	| while_prefix {
	} '{' stmts '}' {
		$$ = &whileStatement{
			Cond: $1,
			// $4 is `stmts`
			Body: $4,
		}
	}
	| PUT '(' put_list ')' ';' {
		$$ = &putStatement{Exprs: $3}
	}
	/* | '{' stmts '}' {
		$$ = &nullStatement{}
	} */
	| ';' {
		$$ = &nullStatement{}
	}
	;

put_list: expr {
		$$ = append($$, $1)
	}
	| put_list ',' expr {
		$$ = append($$, $3)
	}
	;

while_prefix: WHILE '(' cond ')' {
		$$ = $3
	}
	;

cond: expr EQOP expr {
		$$ = &condExpression{LHS: $1, RHS: $3, Operator: EQOP}
	}
	| expr NEOP expr {
		$$ = &condExpression{LHS: $1, RHS: $3, Operator: NEOP}
	}
	| expr GEOP expr {
		$$ = &condExpression{LHS: $1, RHS: $3, Operator: GEOP}
	}
	| expr GTOP expr {
		$$ = &condExpression{LHS: $1, RHS: $3, Operator: GTOP}
	}
	| expr LEOP expr {
		$$ = &condExpression{LHS: $1, RHS: $3, Operator: LEOP}
	}
	| expr LTOP expr {
		$$ = &condExpression{LHS: $1, RHS: $3, Operator: LTOP}
	}
	| expr {
		$$ = $1;
	}
	;
expr:
	NUMBER_LITERAL {
		$$ = &numberExpression{Lit: $1.lit}
	}
	| IDENT {
		 $$ = &variableExpression{Lit: $1.lit}
	}
	| STRING_LITERAL {
		$$ = &stringExpression{Lit: $1.lit}
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
