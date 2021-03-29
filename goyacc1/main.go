package main

import (
	"bufio"
	"fmt"
	"os"
)

type lexer struct {
	src    string
	index  int
	result int
}

// Lex Called by goyacc
func (l *lexer) Lex(lval *yySymType) int {
	for l.index < len(l.src) {
		c := l.src[l.index]
		l.index++
		if c == '+' {
			return int(c)
		}
		if c == '-' {
			return int(c)
		}
		if c == '*' {
			return int(c)
		}
		if c == '/' {
			return int(c)
		}
		if '0' <= c && c <= '9' {
			lval.num = int(c - '0')
			return NUMBER
		}
	}
	return -1
}

// Error Called by goyacc
func (l *lexer) Error(e string) {
	fmt.Println("[error] " + e)
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var source string
	for s.Scan() {
		source += s.Text()
	}
	lexer := &lexer{src: source, index: 0}
	yyParse(lexer)
	println("result:", lexer.result)
}
