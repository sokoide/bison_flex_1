%{
package prolog

import (
    log "github.com/sirupsen/logrus"
)

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
%token<tok> IDENT NUMBER STRING VAR PERIOD COMMA COLON_DASH

// Operator precedence rules
%left '+' '-'
%left '*' '/'

%start input

%%

// Grammar rules
input:
    clause_list PERIOD {
        log.Debugf("Parsed clauses: %v", $1)
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
    term PERIOD {
        $$ = &factClause{Fact: $1}
    }
    ;

rule_clause:
    term COLON_DASH term_list PERIOD {
        $$ = &ruleClause{HeadTerm: $1, BodyTerms: $3}
    }
    ;

term_list:
    term {
        $$ = append($$, $1)
    }
    | term_list COMMA term {
        $$ = append($1, $3)
    }
    ;

term:
    IDENT {
        $$ = &compoundTerm{Functor: $1.Value, Args: nil}
    }
    | IDENT '(' term_list ')' {
        $$ = &compoundTerm{Functor: $1.Value, Args: $3}
    }
    | VAR {
        $$ = &variableTerm{Name: $1.Value}
    }
    | NUMBER {
        $$ = &constantTerm{Lit: $1.Value}
    }
    | STRING {
        $$ = &constantTerm{Lit: $1.Value}
    }
    ;

%%