WHITESPACE [ \t]
INTEGER    [0-9]+
SYMBOL     [+\-*/()]
LF         [\n]

SUBJECT "I"
VERB "LIKE"
OBJECT "SOCCER"



%option noyywrap

%%
{WHITESPACE}
{INTEGER} { yylval.int_value = atoi(yytext);
            return(INTEGER); }
{SYMBOL}  { return(yytext[0]); }
{LF}      { return(LF); }

{SUBJECT} { return(SUBJECT); }
{VERB} { return(VERB); }
{OBJECT} { yylval.int_value = 42; return(OBJECT); }
.
%%

