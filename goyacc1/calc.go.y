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

type Operator struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (s Operator) Eval() int {
	switch s.Operator {
	case "+":
		return s.Left.Eval() + s.Right.Eval()
	case "-":
		return s.Left.Eval() - s.Right.Eval()
	case "*":
		return s.Left.Eval() * s.Right.Eval()
	case "/":
		return s.Left.Eval() / s.Right.Eval()
	}
	return 0
}

%}

%union{
	token *simplexer.Token
	expr  Expression
}

%type<expr> program expr operator
%token<token> NUMBER L1_OPERATOR L2_OPERATOR LP RP

%left L1_OPERATOR
%left L2_OPERATOR
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
	| operator {
		$$ = $1
	}

operator: expr L2_OPERATOR expr {
		$$ = Operator{
			Left: $1,
			Operator: $2.Literal,
			Right: $3,
		}
	}
	| expr L1_OPERATOR expr {
		$$ = Operator{
			Left: $1,
			Operator: $2.Literal,
			Right: $3,
		}
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
		simplexer.NewRegexpTokenType(L1_OPERATOR, `[-+]`),
		simplexer.NewRegexpTokenType(L2_OPERATOR, `[*/]`),
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
	// lexer := NewLexer(strings.NewReader("(1 + 2) * 3 / 4"))
    // yyParse(lexer)
    
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
	    lexer := NewLexer(strings.NewReader(s.Text()))
        yyParse(lexer)
    }
}
