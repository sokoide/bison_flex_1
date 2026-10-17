/* Flex-only example: a token printer.
 *
 * This lexer recognizes three kinds of input:
 *   - INTEGER: numbers like 12, 345
 *   - SYMBOL : single characters from + - * / ( )
 *   - everything else (spaces, tabs, newlines, other chars)
 *
 * INTEGER and SYMBOL rules print what was matched (yytext).
 * Everything else is skipped: the WHITESPACE, LF, and catch-all
 * "." rules have explicit EMPTY actions, so flex's default action
 * (ECHO) never fires and only tokens are printed. (A rule with no
 * action at all would silently echo its match to stdout.)
 *
 * Note: flex 2.6.4 rejects C-style comments placed inside the
 * rules section (between the two %% lines), so the inline notes
 * below live in this header instead.
 */
%option noyywrap

WHITESPACE [ \t]
INTEGER    [0-9]+
SYMBOL     [+\-*/()]
LF         "\n"

%%
{WHITESPACE} { }
{INTEGER} {
			printf("INTEGER: %s\n", yytext);
		  }
{SYMBOL}  {
			printf("SYMBOL: %s\n", yytext);
		  }
{LF}      { }
.         { }
%%

int main(){ return yylex();}
