%namespace netyacc2.Interp
%partial
%parsertype InterpParser
%visibility internal
%tokentype Token

%union {
       public int n;
       public string s;
       public int label;
       public Node node;
}

%left '+' '-'
%left '*' '/'
%right MINUS

%token<n> NUMBER
%token<s> IDENT
%token IF ELSE WHILE
%token EQOP GTOP GEOP LTOP LEOP NEOP
%token ADD SUB MUL DIV
%token PUT GET

%type<label> if_prefix while_prefix
%type<node> expr cond
%start program

%%

program: stmts { Dump(); }
       ;

stmts: stmts stmt
       | stmt
       ;

stmt: PUT '(' put_list ')' ';'
       | IDENT '=' expr ';' {
              GenExpr($3);
              GenCode(Op.Pop, MakeNode(Token.IDENT, $1));
       }
       | if_prefix stmt {
              // TODO:
              // GenCode(Op.Label, $1);
       }
       | if_prefix stmt ELSE {
              // TODO:
              // GenCode(Op.Jump, $<label>$ = "newlbl");
              // GenCode(Op.Label, $1);
       } stmt {
              // TODO:
              // GenCode(Op.Label, $<label>4);
       }
       | while_prefix {
              GenCode(Op.JumpF, MakeNode(Token.WHILE, $<label>$ = NewLabel()));
       } stmt {
              GenCode(Op.Jump, MakeNode(Token.WHILE, $1));
              GenCode(Op.Label, MakeNode(Token.WHILE, $<label>2));
       }
       | '{' stmts '}' { ; }
       | ';'
       ;

put_list: put_id_num_str
       | put_list ',' put_id_num_str
       ;

put_id_num_str: IDENT { GenCode(Op.PutI, MakeNode(Token.IDENT, $1)); }
       | NUMBER { GenCode(Op.PutN, MakeNode(Token.NUMBER, $1)); }
       ;

if_prefix: IF '(' cond ')' {
              GenExpr($3);
              // GenCode(Op.JumpF, $$=NewLabel());
              }
       ;

while_prefix: WHILE '(' cond ')' {
              GenCode(Op.Label, MakeNode(Token.WHILE, $$=NewLabel()));
              GenExpr($3);
       }
       ;

cond: expr EQOP expr { $$ = MakeExpr(Token.EQOP, $1, $3); }
       | expr GTOP expr { $$ = MakeExpr(Token.GTOP, $1, $3); }
       | expr GEOP expr { $$ = MakeExpr(Token.GEOP, $1, $3); }
       | expr LTOP expr { $$ = MakeExpr(Token.LTOP, $1, $3); }
       | expr LEOP expr { $$ = MakeExpr(Token.LEOP, $1, $3); }
       | expr NEOP expr { $$ = MakeExpr(Token.NEOP, $1, $3); }
       | expr { $$ = $1; }
       ;

expr:  expr '+' expr { $$ = MakeExpr(Token.ADD, $1, $3); }
       | '-' expr %prec MINUS {
              $$ = MakeExpr(Token.MINUS, $2, null);
              //Console.WriteLine("%prec MINUS {0}", $2);
       }
       | expr '-' expr { $$ = MakeExpr(Token.SUB, $1, $3); }
       | expr '*' expr { $$ = MakeExpr(Token.MUL, $1, $3); }
       | expr '/' expr { $$ = MakeExpr(Token.DIV, $1, $3); }
       | '(' expr ')' { $$ = $2; }
       | IDENT { $$ = MakeNode(Token.IDENT, $1); }
       | NUMBER { $$ = MakeNode(Token.NUMBER, $1); }
       ;

%%