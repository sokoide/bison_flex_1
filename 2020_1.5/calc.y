%{
#include <stdio.h>

extern int yylex (void);
extern int yyerror (const char* str);
%}
%token LF
%token INTEGER
%token SUBJECT VERB OBJECT

%left '+' '-'

%start program
%%

program : /* empty */
     | SUBJECT VERB obj LF { printf("result: %d\n", $3 )}
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
    yyparse();
    return 0;
}
