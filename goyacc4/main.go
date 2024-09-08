package main

import (
	"bufio"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	scanner := new(scanner)
	source := []string{}
	for s.Scan() {
		source = append(source, s.Text())
	}

	scanner.Init(strings.Join(source, "\n"))

	var prog []statement = parse(scanner)
	_, err := evaluateStmts(prog)
	if err != nil {
		panic(err)
	}
}

type lexer struct {
	s         *scanner
	recentLit string
	recentPos position
	program   []statement
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

func parse(s *scanner) []statement {
	l := lexer{s: s}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.program
}
