%{
package main

import (
    "bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/macrat/simplexer"
)

type Expression interface {
	Eval() int
}

type Number int

func (n Number) Eval() int {
	return int(n)
}
%}

%union{
	token *simplexer.Token
	expr  Expression
}

%type<expr> program expr
%token<token> NUMBER LP RP

%left ADD SUB
%left MUL DIV

%start program

%%
program: expr {
		$$ = $1
		yylex.(*Lexer).result = $$
        // fmt.Printf("%v\n", $$)
        fmt.Printf("%v\n", $$.Eval())
	}

expr: NUMBER {
		f, _ := strconv.ParseFloat($1.Literal, 64)
		$$ = Number(f)
	}
    | expr ADD expr {
        $$ = Number($1.Eval() + $3.Eval())
	}
    | expr SUB expr {
        $$ = Number($1.Eval() - $3.Eval())
	}
    | expr MUL expr {
        $$ = Number($1.Eval() * $3.Eval())
	}
    | expr DIV expr {
        $$ = Number($1.Eval() / $3.Eval())
	}
    | LP expr RP {
        $$ = $2
    }

%%

type Lexer struct {
	lexer        *simplexer.Lexer
	lastToken    *simplexer.Token
	result       Expression
}

func NewLexer(reader io.Reader) *Lexer {
	l := simplexer.NewLexer(reader)

	l.TokenTypes = []simplexer.TokenType{
		simplexer.NewRegexpTokenType(NUMBER, `[0-9]+`),
		simplexer.NewRegexpTokenType(ADD, `\+`),
		simplexer.NewRegexpTokenType(SUB, `\-`),
		simplexer.NewRegexpTokenType(MUL, `\*`),
		simplexer.NewRegexpTokenType(DIV, `/`),
		simplexer.NewRegexpTokenType(LP, `\(`),
		simplexer.NewRegexpTokenType(RP, `\)`),
	}

	return &Lexer{ lexer: l }
}

func (l *Lexer) Lex(lval *yySymType) int {
	token, err := l.lexer.Scan()
	if err != nil {
		if e, ok := err.(simplexer.UnknownTokenError); ok {
			fmt.Fprintln(os.Stderr, e.Error() + ":")
			fmt.Fprintln(os.Stderr, l.lexer.GetLastLine())
			fmt.Fprintln(os.Stderr, strings.Repeat(" ", e.Position.Column) + strings.Repeat("^", len(e.Literal)))
		} else {
			l.Error(err.Error())
		}
		os.Exit(1)
	}
	if token == nil {
		return -1
	}

	lval.token = token

	l.lastToken = token

	return int(token.Type.GetID())
}

func (l *Lexer) Error(e string) {
	fmt.Fprintln(os.Stderr, e + ":")
	fmt.Fprintln(os.Stderr, l.lexer.GetLastLine())
	fmt.Fprintln(os.Stderr, strings.Repeat(" ", l.lastToken.Position.Column) + strings.Repeat("^", len(l.lastToken.Literal)))
}

func main() {
	/* lexer := NewLexer(strings.NewReader("(1 + 2) * 3 / 4"))
    yyParse(lexer) */
    
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
	    lexer := NewLexer(strings.NewReader(s.Text()))
        yyParse(lexer)
    }
}
