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
		if isWhiteSpace(c) {
			l.index++
			continue
		}
		if '0' <= c && c <= '9' {
			lval.num = l.getNumber()
			return NUMBER
		}

		l.index++
		if c == ';' || c == '\n' {
			return LF
		}
		if c == '+' || c == '-' || c == '*' || c == '/' ||
			c == '%' || c == '(' || c == ')' || c == '=' {
			return int(c)
		}
		if 'a' <= c && c <= 'z' {
			lval.ident = string(c)
			return IDENT
		}
	}
	return -1
}

// Error Called by goyacc
func (l *lexer) Error(e string) {
	fmt.Println("[error] " + e)
}

func (l *lexer) getNumber() int {
	n := 0
	for l.index < len(l.src) && '0' <= l.src[l.index] && l.src[l.index] <= '9' {
		c := int(l.src[l.index] - '0')
		n = n*10 + c
		l.index++
	}
	return n
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var source string
	for s.Scan() {
		source += s.Text()
	}
	lexer := &lexer{src: source, index: 0}
	yyParse(lexer)
}
