%namespace netyacc1.Calc
%scannertype CalcScanner
%visibility internal
%tokentype Token

%option stack, minimize, parser, verbose, persistbuffer, noembedbuffers

Eol             (\r\n?|\n)
NotWh           [^ \t\r\n]
Space           [ \t]
Number          [0-9]+
Symbol     		[+\-*/()]

%{

%}

%%

/* Scanner body */

{Number}		{
					GetNumber();
					return (int)Token.NUMBER;
				}
{Space}+		/* skip */
{Symbol}  		{ return(yytext[0]); }

%%
