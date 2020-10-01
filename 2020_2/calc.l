WHITESPACE [ \t]
INTEGER    [0-9]+
SYMBOL     [+\-*/()]
LF         "\n"

%option noyywrap

%%
{WHITESPACE}
{INTEGER} { yylval = atoi(yytext);
            return(INTEGER); }
{SYMBOL}  { return(yytext[0]); }
{LF}      { return(LF); }
.
%%

