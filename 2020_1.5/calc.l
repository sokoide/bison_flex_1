WHITESPACE [ \t]
INTEGER    [0-9]+
SYMBOL     [+\-*/()]
LF         "\n"

SUBJECT "I"
VERB "LIKE"
OBJECT "SOCCER"



%option noyywrap

%%
{WHITESPACE}
{INTEGER} { yylval = atoi(yytext);
            return(INTEGER); }
{SYMBOL}  { return(yytext[0]); }
{LF}      { return(LF); }

{SUBJECT} { return(SUBJECT); }
{VERB} { return(VERB); }
{OBJECT} { yylval = 42; return(OBJECT); }
.
%%

