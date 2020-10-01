%{
#include <stdio.h>

extern int yylex (void);
extern int yyerror (const char* str);
%}
%token LF
%token INTEGER

%left '+' '-'
%left '*' '/'

%start program
%%
program : /* empty */
     | program LF
     | program expr LF { printf("%d\n",$2); }
     ;
expr : INTEGER
     | expr '+' expr { $$ = $1 + $3; }
     | expr '-' expr { $$ = $1 - $3; }
     | expr '*' expr { $$ = $1 * $3; }
     | expr '/' expr { $$ = $1 / $3; }
     | '(' expr ')'  { $$ = $2;      }
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
