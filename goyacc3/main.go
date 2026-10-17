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

	program, err := parse(scanner)
	if err != nil {
		log.Fatalf("parse error: %v", err)
	}

	v, err := evaluate(program)
	if err != nil {
		log.Fatalf("eval error: %v", err)
	}
	fmt.Println(v)
}

type lexer struct {
	s         *scanner
	recentLit string
	recentPos position
	program   expression
	err       error
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

// Error Called by goyacc. Stores the message instead of exiting, so that
// parse() can return it and the caller decides what to do.
func (l *lexer) Error(e string) {
	if l.err == nil {
		l.err = fmt.Errorf("Line %d, Column %d: %q %s",
			l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
	}
}

func parse(s *scanner) (expression, error) {
	l := lexer{s: s}
	yyErrorVerbose = true
	if yyParse(&l) != 0 {
		err := l.err
		if err == nil {
			err = fmt.Errorf("syntax error")
		}
		return nil, err
	}
	return l.program, nil
}
