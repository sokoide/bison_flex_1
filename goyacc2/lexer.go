package main

import (
	"fmt"
)

type lexer struct {
	src   string
	index int
}

var recentToken string

// Lex Called by goyacc
func (l *lexer) Lex(lval *yySymType) int {
	for l.index < len(l.src) {
		c := l.src[l.index]
		if isWhiteSpace(c) {
			l.index++
			continue
		}

		l.index++
		switch c {
		case '+', '-', '*', '/', '(', ')', '%', '=':
			recentToken = string(c)
			return int(c)
		}

		switch {
		case '0' <= c && c <= '9':
			// left value is a number
			lval.val = l.getNumber()
			recentToken = string(lval.val)
			return NUMBER
		case 'a' <= c && c <= 'z':
			lval.ident = l.getString()
			recentToken = lval.ident
			return IDENT
		default:
			panic(fmt.Sprintf("invalid character %s\n", string(c)))
		}
	}
	return -1
}

// Error Called by goyacc
func (l *lexer) Error(e string) {
	fmt.Printf("error: %s, at token: %s", e, recentToken)

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

func (l *lexer) getString() string {
	s := string(l.src[l.index-1])
	for l.index < len(l.src) && 'a' <= l.src[l.index] && l.src[l.index] <= 'z' {
		s += string(l.src[l.index])
		l.index++
	}
	return s
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
