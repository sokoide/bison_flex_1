%namespace interp_lib.Interp
%partial
%parsertype InterpParser
%visibility public
%tokentype Token

%union {
       public Node node;    // node
       public int labelno;  // label number
       public int addr;     // string address
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

%type<labelno> if_prefix while_prefix
%type<node> expr cond
%start program

%%

program: stmts { }
       ;

stmts: /* empty  */
       | stmts stmt
       ;

stmt:  IDENT '=' expr ';' {
              GenNode($3);
              GenCode(Op.Pop, $1);
       }
       | if_prefix stmt {
              GenCode(Op.Label, $1);
       }
       | if_prefix stmt ELSE {
              GenCode(Op.Jump, $<labelno>$ = NewLabel());
              GenCode(Op.Label, $1);
       } stmt {
              GenCode(Op.Label, $<labelno>4);
       }
       | while_prefix {
              // $<labelno>$ means a value of this scope which means $2 usied by the following `stmt``
              GenCode(Op.JumpF, $<labelno>$=NewLabel());
       } stmt {
              // $1 means a value of `while_prefix`
              GenCode(Op.Jump, $1);
              // $<labelno>2 means a value of $2 as `labelno` type  which is `GenCode(Op.JumpF... inside while_prefix`
              // `stmt` is $3
              GenCode(Op.Label, $<labelno>2);
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
              GenNode($3);
              GenCode(Op.JumpF, $$=NewLabel());
       }
       ;

while_prefix: WHILE '(' cond ')' {
              GenCode(Op.Label, $$=NewLabel());
              GenNode($3);
       }
       ;

cond:  expr EQOP expr { $$ = MakeNode(Token.EQOP, $1, $3); }
       | expr GTOP expr { $$ = MakeNode(Token.GTOP, $1, $3); }
       | expr GEOP expr { $$ = MakeNode(Token.GEOP, $1, $3); }
       | expr LTOP expr { $$ = MakeNode(Token.LTOP, $1, $3); }
       | expr LEOP expr { $$ = MakeNode(Token.LEOP, $1, $3); }
       | expr NEOP expr { $$ = MakeNode(Token.NEOP, $1, $3); }
       | expr { $$ = $1; }
       ;

expr:  expr '+' expr { $$ = MakeNode(Token.ADD, $1, $3); }
       | '-' expr %prec MINUS {
              $$ = MakeNode(Token.MINUS, $2, null);
              //Console.WriteLine("%prec MINUS {0}", $2);
       }
       | expr '-' expr { $$ = MakeNode(Token.SUB, $1, $3); }
       | expr '*' expr { $$ = MakeNode(Token.MUL, $1, $3); }
       | expr '/' expr { $$ = MakeNode(Token.DIV, $1, $3); }
       | '(' expr ')' { $$ = $2; }
       | IDENT { $$ = $1; }
       | NUMBER { $$ = $1; }
       ;

%%