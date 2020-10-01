%option noyywrap

WHITESPACE [ \t]
INTEGER    [0-9]+
SYMBOL     [+\-*/()]
LF         "\n"

%%
{WHITESPACE}
{INTEGER} {
			printf("INTEGER: %s\n", yytext);
		  }
{SYMBOL}  {
			printf("SYMBOL: %s\n", yytext);
		  }
{LF}
.
%%

int main(){ return yylex();}
