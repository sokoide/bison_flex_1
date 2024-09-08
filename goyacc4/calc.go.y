%{
package main
%}

%union{
    stmt    statement
    expr    expression
    tok     token
}

%type<stmt> program stmts stmt
%type<expr> expr
%token<tok> NUMBER IDENT

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
    stmt { $$ = $1; }
    | stmts stmt { $$ = $2; }
    ;

stmt:
    expr ';' { $$ = &exprStatement{Expr: $1} }
    ;

expr:
    NUMBER {
        $$ = &numberExpression{Lit: $1.lit}
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
    | IDENT '=' expr {
        $$ = &assignExpression{}
    }

%%
