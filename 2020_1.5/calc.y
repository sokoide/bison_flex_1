%{
#include <stdio.h>
#include <stdlib.h>

extern int yylex (void);
extern int yyerror (const char* str);
%}
%union {
    int int_value;
}

%token LF
%token<int_value> INTEGER
%token<int_value> SUBJECT VERB OBJECT

%left '+' '-'

%type<int_value> obj

%start program

%%
program: line
    | program line
    ;

line:  LF { exit(0); }
    | SUBJECT VERB obj LF { printf("result: %d\n", $3 ); }
    ;

obj  : OBJECT {$$=$1;}
	 | INTEGER { $$=$1;}
	 ;

%%
#include "lex.yy.c"

int yyerror (const char* str){
    fprintf(stderr, "parse error near %s\n", yytext);
    return 0;
}

int main() {
    /* yydebug = 1; */
    yyparse();
    return 0;
}
