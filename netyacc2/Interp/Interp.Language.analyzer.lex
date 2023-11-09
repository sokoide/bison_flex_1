%namespace netyacc2.Interp
%scannertype InterpScanner
%visibility internal
%tokentype Token

%option stack, minimize, parser, verbose, persistbuffer, noembedbuffers

%x STR
%x COMMENT

If              if
Else            else
While           while
Print           print


Eqop            ==
Gtop            >
Geop            >=
Ltop            <
Leop            <=
Neop            !=
Dq              \"
Sq              \'
CommSt          \/\*
CommEd          \*\/
Symbol          [+\-\*\/\(\)=;]
Number          [0-9]+
Ident           [a-zA-Z][0-9a-zA-Z_]*

Space           [ \t]
Eol             (\r\n?|\n)
/* what is NotWh? */
NotWh           [^ \t\r\n]

%{
int strcnt = 0;
%}

%%

/* Scanner body */

<INITIAL>{If}        { return(int)Token.IF; }
<INITIAL>{Else}      { return(int)Token.ELSE; }
<INITIAL>{While}     { return(int)Token.WHILE; }
<INITIAL>{Print}     { return(int)Token.PRINT; }

<INITIAL>{Symbol}    { return(yytext[0]); }
<INITIAL>{Number}    {
                    Console.WriteLine("number: {0}", yytext);
                    GetNumber();
                    return (int)Token.NUMBER;
                }
<INITIAL>{Ident}     {
                    Console.WriteLine("ident: {0}", yytext);
                    GetString();
                    return (int)Token.IDENT;
                }
<INITIAL>{Space}+    ; /* skip */
<INITIAL>{Eol}       ; /* skip */
<INITIAL>{Dq}       { strcnt=0; BEGIN(STR); }
<INITIAL>{CommSt}   { BEGIN(COMMENT); }
<STR>{Dq}           { BEGIN(INITIAL); }
<COMMENT>{CommEd}   { BEGIN(INITIAL); }
<COMMENT>[^\*]*     ;

%%