%namespace netyacc2.Interp
%partial
%parsertype InterpParser
%visibility internal
%tokentype Token

%union {
			public int n;
			public string s;
	   }

%left '+' '-'
%left '*' '/'
%right MINUS

%token NUMBER IDENT
%token IF ELSE WHILE
%token EQOP GTOP GEOP LTOP LEOP NEOP
%token PRINT

%start program

%%

program: stmts
       ;

stmts: stmts stmt
       | stmt
       ;

stmt: PRINT '(' expr ')' ';' {
              Console.WriteLine($3.n);
       }
       | IDENT '=' expr ';' {
              Console.WriteLine("assign: {0}={1}", $1.s, $3.n);
       }
       | if_prefix stmt { ; }
       | if_prefix stmt ELSE { ; } stmt
       | while_prefix { ; } stmt { ; }
       | '{' stmts '}' { ; }
       | ';'
       ;

if_prefix: IF '(' cond ')' { ; }
       ;

while_prefix: WHILE '(' cond ')' { ; }
       ;

cond: expr EQOP expr { ; }
       | expr GTOP expr { ; }
       | expr GEOP expr { ; }
       | expr LTOP expr { ; }
       | expr LEOP expr { ; }
       | expr NEOP expr { ; }
       ;

expr:  expr '+' expr { $$.n = $1.n + $3.n; }
       | expr '-' expr { $$.n = $1.n - $3.n; }
       | expr '*' expr { $$.n = $1.n * $3.n; }
       | expr '/' expr { $$.n = $1.n / $3.n; }
       | '-' expr %prec MINUS {
              $$.n = -1 * $2.n;
              Console.WriteLine("%prec MINUS {0} -> {1}", $2.n, $$.n);
       }
       | '(' expr ')' { $$.n = $2.n; }
       | IDENT { /* TODO */ $$.n = 0; }
       | NUMBER { $$.n = $1.n; }
       ;

%%