package prolog

import (
	"bufio"
	"io"
	"os"
	"unicode"

	log "github.com/sirupsen/logrus"
)

type Lexer struct {
	reader *bufio.Reader
	ch     rune
	file   *os.File
}

func NewLexer(filename string) (*Lexer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	lexer := &Lexer{
		reader: reader,
		file:   file,
	}
	lexer.readChar()
	return lexer, nil
}

func (l *Lexer) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

// Lex Called by goyacc
func (l *Lexer) Lex(lval *yySymType) int {
	var id int
	id, lval.tok = l.NextToken()
	log.Debugf("Lexer: Lex: %d: Token: %+v", id, lval.tok)
	return id
}

// Error Called by goyacc
func (l *Lexer) Error(e string) {
	log.Fatalf("Lexer: Error: %s", e)
}

// ========================================
type tokenType int

const (
	tokenTypeNone tokenType = iota
	tokenTypeEOF
	tokenTypeNumberLiteral
	tokenTypeStringLiteral
	tokenTypeIdent
	tokenTypeKeyword
	tokenTypeOp
)

func (l *Lexer) readChar() {
	ch, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			l.ch = 0
		} else {
			panic(err)
		}
	} else {
		l.ch = ch
	}
}

func (l *Lexer) NextToken() (int, token) {
	l.skipWhitespace()

	var id int
	var tok token

	switch {
	case l.ch == 0:
		id = 0
		tok = token{Type: tokenTypeEOF, Value: ""}
	case unicode.IsLetter(l.ch):
		id = IDENT
		tok = token{Type: tokenTypeIdent, Value: l.readIdentifier()}
	case unicode.IsNumber(l.ch):
		id = NUMBER
		tok = token{Type: tokenTypeNumberLiteral, Value: l.readNumber()}
	default:
		switch l.ch {
		case -1:
			id = 0
			tok = token{Type: tokenTypeEOF, Value: ""}
		case ',', '.', '[', ']', '(', ')', '{', '}', ';', '+', '-', '*', '/', '%', '=':
			id = int(l.ch)
			tok = token{Type: tokenTypeOp, Value: string(l.ch)}
			l.readChar()
		default:
			// TODO:
			panic("unexpected token")
		}
	}

	return id, tok
}

func (l *Lexer) readIdentifier() string {
	var result []rune
	for unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) {
		result = append(result, l.ch)
		l.readChar()
	}
	return string(result)
}

func (l *Lexer) readNumber() string {
	var result []rune
	for unicode.IsDigit(l.ch) {
		result = append(result, l.ch)
		l.readChar()
	}
	return string(result)
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}
