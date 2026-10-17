WHITESPACE [ \t]
INTEGER    [0-9]+
SYMBOL     [+\-*/()]
LF         [\n]

SUBJECT "I"
VERB "LIKE"
OBJECT "SOCCER"



%option noyywrap

/* A rule with no action defaults to ECHO (the match is printed to
 * stdout). In this parser demo only tokens should reach the parser,
 * so WHITESPACE below has an explicit skip action, and any unknown
 * character (the catch-all "." rule at the end) is reported to
 * stderr instead of being echoed. */
%%
{WHITESPACE} { /* skip whitespace (default action would ECHO it) */ }
{INTEGER} { yylval.int_value = atoi(yytext);
            return(INTEGER); }
{SYMBOL}  { return(yytext[0]); }
{LF}      { return(LF); }

{SUBJECT} { return(SUBJECT); }
{VERB} { return(VERB); }
{OBJECT} { yylval.int_value = 42; return(OBJECT); }
.         { fprintf(stderr, "Unknown character: '%c'\n", yytext[0]); }
%%

