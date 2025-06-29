%{
#include <stdio.h>
#include <stdlib.h>

/* Forward declarations */
extern int yylex(void);
extern int yyerror(const char* str);
extern char* yytext;

/* Global variables */
int division_by_zero_error = 0;
%}

%union {
    int int_value;
}

%token LF
%token<int_value> INTEGER

%left '+' '-'
%left '*' '/'
%right UMINUS

%type<int_value> expr
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
     | expr '/' expr { 
         if ($3 == 0) {
             division_by_zero_error = 1;
             yyerror("division by zero");
             YYERROR;
         }
         $$ = $1 / $3; 
     }
     | '(' expr ')'  { $$ = $2;      }
     | '-' expr %prec UMINUS { $$ = -$2; }
     ;

%%

/* Include the lexer implementation */
#include "lex.yy.c"

int yyerror (const char* str){
    if (division_by_zero_error) {
        fprintf(stderr, "Error: %s\n", str);
        division_by_zero_error = 0;
    } else {
        fprintf(stderr, "Parse error: %s near '%s'\n", str, yytext);
    }
    return 0;
}

int main(void) {
    printf("Simple Calculator\n");
    printf("Enter expressions (Ctrl+D to exit):\n");
    
    /* Uncomment for debugging: yydebug = 1; */
    int result = yyparse();
    
    printf("\nGoodbye!\n");
    return result;
}
