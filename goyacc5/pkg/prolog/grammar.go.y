%{
package prolog
%}

// yySymType
%union {
    clauses  	[]clause
    clause  	clause
    fact    	*factClause
    rule    	*ruleClause
    terms  	 	[]term
    term    	term
    tok     	token
}

// Types
%type<clauses> clause_list
%type<clause> clause
%type<fact> fact_clause
%type<rule> rule_clause
%type<term> term
%type<terms> term_list

// Tokens
%token<tok> IDENT NUMBER VAR OP COLON_DASH

// Operator precedence rules
%left '+' '-'
%left '*' '/'

%start input

%%

// Grammar rules
input:
    clause_list {
        yylex.(*Lexer).program = $1
    }
    ;

clause_list:
    clause {
        $$ = append($$, $1)
    }
    | clause_list clause {
        $$ = append($1, $2)
    }
    ;

clause:
    fact_clause {
        $$ = $1
    }
    | rule_clause {
        $$ = $1
    }
    ;

fact_clause:
    term '.' {
        $$ = &factClause{Fact: $1}
    }
    ;

rule_clause:
    term COLON_DASH term_list '.' {
        $$ = &ruleClause{HeadTerm: $1, BodyTerms: $3}
    }
    ;

term_list:
    term {
        $$ = append($$, $1)
    }
    | term_list ',' term {
        $$ = append($1, $3)
    }
    ;

term:
    IDENT {
        $$ = &constantTerm{Lit: $1.Value}
    }
    | VAR {
        $$ = &variableTerm{Name: $1.Value}
    }
    | IDENT '(' term_list ')' {
        $$ = &compoundTerm{Functor: $1.Value, Args: $3}
    }
    | NUMBER {
        $$ = &constantTerm{Lit: $1.Value}
    }
    | '[' term_list ']' {
        $$ = &listTerm{Args: $2}
    }
    | '[' ']' {
        $$ = &listTerm{Args: []term{}}
    }
    ;

%%