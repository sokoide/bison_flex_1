%namespace interp_lib.Interp
%scannertype InterpScanner
%visibility internal
%tokentype Token

%option stack, minimize, parser, verbose, persistbuffer, noembedbuffers

%{
    StringBuilder sb = new System.Text.StringBuilder();
%}

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
Symbol          [+\-\*\/\(\)=;,\{\}]

/* number, ident */
Number          [0-9]+
Ident           [a-zA-Z][0-9a-zA-Z_]*

/* others */
Space           [ \t]
Eol             (\r\n?|\n)

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
                    yylval.node = new Node(Token.NUMBER, int.Parse(yytext));
                    return (int)Token.NUMBER;
                }
<INITIAL>{Ident}     {
                    // Console.WriteLine("ident: {0}", yytext);
                    yylval.node = new Node(Token.IDENT, yytext);
                    return (int)Token.IDENT;
                }
<INITIAL>{Space}+    ; /* skip */
<INITIAL>{Eol}       ; /* skip */

/* string */
<INITIAL>{Dq}       { sb.Clear(); BEGIN(STR);}
<STR>{Dq}           { yylval.addr = Pool(sb.ToString()); BEGIN(INITIAL); }
<STR>\\\"           { sb.Append("\""); }
<STR>\\n            { sb.Append("\n"); }
<STR>\n             { throw new Exception("string not closed"); }
<STR>.              { sb.Append(yytext[0]); }

/* comment */
<INITIAL>{CommSt}   { BEGIN(COMMENT); }
<COMMENT>{CommEd}   { BEGIN(INITIAL); }
<COMMENT>[^\*]*     ;

%%