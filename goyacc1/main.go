package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	scanner := new(scanner)
	source := []string{}
	for s.Scan() {
		source = append(source, s.Text())
	}
	scanner.Init(strings.Join(source, "\n"))

	var program expression = parse(scanner)
	v, err := evaluate(program)
	if err != nil {
		panic(nil)
	}
	fmt.Println(v)
}

type lexer struct {
	s         *scanner
	recentLit string
	recentPos position
	program   expression
}

// Lex Called by goyacc
func (l *lexer) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.Scan()
	if tok == EOF {
		return 0
	}
	lval.tok = token{tok: tok, lit: lit, pos: pos}
	l.recentLit = lit
	l.recentPos = pos
	return tok
}

// Error Called by goyacc
func (l *lexer) Error(e string) {
	log.Fatalf("Line %d, Column %d: %q %s",
		l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func parse(s *scanner) expression {
	l := lexer{s: s}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.program
}
