WHITESPACE [ \t]
INTEGER    [0-9]+
SYMBOL     [+\-*/()]
LF         [\n]

%option noyywrap

%%
{WHITESPACE}
{INTEGER} { yylval.int_value = atoi(yytext);
            return(INTEGER); }
{SYMBOL}  { return(yytext[0]); }
{LF}      { return(LF); }
. { fprintf(stderr, "Unknown character: '%c'\\n", yytext[0]); }
%%

