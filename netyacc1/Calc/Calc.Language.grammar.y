%namespace netyacc1.Calc
%partial
%parsertype CalcParser
%visibility internal
%tokentype Token

%union {
	public int n;
	public string s;
}


%token<n> NUMBER

%left '+' '-'
%left '*' '/'

%start program

%%

program	: expr {
			Console.WriteLine("{0}\n", $1.n);
		}
		;

expr: NUMBER
	| expr '+' expr {
		$$.n = $1.n + $3.n;
	}
	| expr '-' expr {
		$$.n = $1.n - $3.n;
	}
	| expr '*' expr {
		$$.n = $1.n * $3.n;
	}
	| expr '/' expr {
		$$.n = $1.n / $3.n;
	}
	| '(' expr ')' {
		$$.n = $2.n;
	}
	;
%%
