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
			lval.num = l.getNumber()
			return NUMBER
		}
	}
	return -1
}

// Error Called by goyacc
func (l *lexer) Error(e string) {
	fmt.Println("[error] " + e)
}

func (l *lexer) getNumber() int {
	n := int(l.src[l.index-1] - '0')

	for l.index < len(l.src) && '0' <= l.src[l.index] && l.src[l.index] <= '9' {
		c := int(l.src[l.index] - '0')
		n = n*10 + c
		l.index++
	}
	return n
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
