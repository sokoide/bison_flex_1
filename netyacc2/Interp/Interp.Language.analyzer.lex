%namespace netyacc2.Interp
%scannertype InterpScanner
%visibility internal
%tokentype Token

%option stack, minimize, parser, verbose, persistbuffer, noembedbuffers

%x STR
%x COMMENT

/* keywords */
If              if
Else            else
While           while
Put             put
Get             get


/* cond */
Eqop            ==
Gtop            >
Geop            >=
Ltop            <
Leop            <=
Neop            !=

/* symbols */
Dq              \"
Sq              \'
CommSt          \/\*
CommEd          \*\/
Symbol          [+\-\*\/\(\)=;,]

/* number, ident */
Number          [0-9]+
Ident           [a-zA-Z][0-9a-zA-Z_]*

/* others */
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
<INITIAL>{Put}       { return(int)Token.PUT; }
<INITIAL>{Get}       { return(int)Token.GET; }

<INITIAL>{Eqop}       { return(int)Token.EQOP; }
<INITIAL>{Gtop}       { return(int)Token.GTOP; }
<INITIAL>{Geop}       { return(int)Token.GEOP; }
<INITIAL>{Ltop}       { return(int)Token.LTOP; }
<INITIAL>{Leop}       { return(int)Token.LEOP; }
<INITIAL>{Neop}       { return(int)Token.NEOP; }

<INITIAL>{Symbol}    { return(yytext[0]); }
<INITIAL>{Number}    {
                    // Console.WriteLine("number: {0}", yytext);
                    GetNumber();
                    return (int)Token.NUMBER;
                }
<INITIAL>{Ident}     {
                    // Console.WriteLine("ident: {0}", yytext);
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