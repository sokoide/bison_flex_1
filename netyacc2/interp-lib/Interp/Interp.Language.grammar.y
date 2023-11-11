%namespace interp_lib.Interp
%partial
%parsertype InterpParser
%visibility public
%tokentype Token

%union {
       public int label;
       public Node node;
}

%left '+' '-'
%left '*' '/'
%right MINUS

%token<node> NUMBER
%token<node> IDENT
%token IF ELSE WHILE
%token EQOP GTOP GEOP LTOP LEOP NEOP
%token ADD SUB MUL DIV
%token PUT GET
%token NULL

%type<label> if_prefix while_prefix
%type<node> expr cond
%start program

%%

program: stmts { }
       ;

stmts: /* empty  */
       | stmts stmt
       ;

stmt:  IDENT '=' expr ';' {
              GenExpr($3);
              GenCode(Op.Pop, $1);
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
              GenCode(Op.JumpF, MakeNode(Token.NULL, $<label>$ = NewLabel()));
       } stmt {
              GenCode(Op.Jump, MakeNode(Token.NULL, $1));
              GenCode(Op.Label, MakeNode(Token.NULL, $<label>2));
       }
       | PUT '(' put_list ')' ';'
       | '{' stmts '}'
       | ';'
       ;

put_list: put_id_num_str
       | put_list ',' put_id_num_str
       ;

put_id_num_str: IDENT { GenCode(Op.PutI, $1); }
       | NUMBER { GenCode(Op.PutN, $1); }
       ;

if_prefix: IF '(' cond ')' {
              GenExpr($3);
              // GenCode(Op.JumpF, $$=NewLabel());
              }
       ;

while_prefix: WHILE '(' cond ')' {
              GenCode(Op.Label, MakeNode(Token.NULL, $$=NewLabel()));
              GenExpr($3);
       }
       ;

cond:  expr EQOP expr { $$ = MakeExpr(Token.EQOP, $1, $3); }
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
       | IDENT { $$ = $1; }
       | NUMBER { $$ = $1; }
       ;

%%