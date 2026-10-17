%{
package main
%}

%union{
    expr    expression
    tok     token
}

%type<expr> program expr
// ILLEGAL is returned by the scanner for any character the language does not
// use. It must NOT be 0 (0 is the EOF token in goyacc's yyTok tables).
%token ILLEGAL
%token<tok> NUMBER

%left '+' '-'
%left '*' '/' '%'

%start program

%%
program: expr {
		$$ = $1
        yylex.(*lexer).program = $$
	}

expr: NUMBER {
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
    | expr '%' expr {
        $$ = &binOpExpression{LHS: $1, Operator: int('%'), RHS: $3}
	}
    | '(' expr ')' {
        $$ = &parenExpression{SubExpr: $2}
    }

%%