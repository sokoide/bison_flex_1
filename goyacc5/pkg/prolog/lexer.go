package prolog

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"

	log "github.com/sirupsen/logrus"
)

type Lexer struct {
	reader  *bufio.Reader
	ch      rune
	file    *os.File
	program []clause
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
	var err error

	id, lval.tok, err = l.NextToken()
	if err != nil {
		panic(err)
	}
	log.Debugf("Lexer: %d %s '%s'", id, tokenTypeToString[lval.tok.Type], lval.tok.Value)
	return id
}

// Error Called by goyacc
func (l *Lexer) Error(e string) {
	log.Fatalf("Lexer: Error: %s", e)
}

// ========================================
var keywords = map[string]int{
	"builtin_write": BUILTIN_WRITE,
}

type tokenType int

const (
	tokenTypeNone tokenType = iota
	tokenTypeEOF
	tokenTypeNumberLiteral
	tokenTypeStringLiteral
	tokenTypeIdent
	tokenTypeVariable
	tokenTypeKeyword
	tokenTypeOp
)

var tokenTypeToString = map[tokenType]string{
	tokenTypeNone:          "None",
	tokenTypeEOF:           "EOF",
	tokenTypeNumberLiteral: "NumberLiteral",
	tokenTypeStringLiteral: "StringLiteral",
	tokenTypeIdent:         "Ident",
	tokenTypeVariable:      "Variable",
	tokenTypeKeyword:       "Keyword",
	tokenTypeOp:            "Op",
}

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

func (l *Lexer) NextToken() (int, token, error) {
	for {
		l.skipComment()
		l.skipWhitespace()
		if l.ch != '%' {
			break
		}
	}
	var id int
	var tok token

	switch {
	case l.ch == 0:
		id = 0
		tok = token{Type: tokenTypeEOF, Value: ""}
	case l.ch >= 'a' && l.ch <= 'z':
		literal := l.readIdentifier()
		if val, ok := keywords[literal]; ok {
			id = val
			tok = token{Type: tokenTypeKeyword, Value: literal}
		} else {
			id = IDENT
			tok = token{Type: tokenTypeIdent, Value: literal}
		}
	case l.ch >= 'A' && l.ch <= 'Z':
		id = VAR
		tok = token{Type: tokenTypeVariable, Value: l.readIdentifier()}
	case l.ch == ':':
		l.readChar()
		if l.ch == '-' {
			id = COLON_DASH
			tok = token{Type: tokenTypeOp, Value: ":-"}
			l.readChar()
		} else {
			return 0, tok, fmt.Errorf("Unexpected token %c(%d)", l.ch, l.ch)
		}
	case unicode.IsNumber(l.ch):
		id = NUMBER
		tok = token{Type: tokenTypeNumberLiteral, Value: l.readNumber()}
	default:
		switch l.ch {
		case ',', '.', '[', ']', '|', '(', ')', '{', '}', ';', '+', '-', '*', '/', '%', '=':
			id = int(l.ch)
			tok = token{Type: tokenTypeOp, Value: string(l.ch)}
			l.readChar()
		default:
			return 0, tok, fmt.Errorf("Unexpected token %c(%d)", l.ch, l.ch)
		}
	}

	return id, tok, nil
}

func (l *Lexer) indentifierSupportedChar(r rune) bool {
	if unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) {
		return true
	}

	supportedChars := []rune{'_', '|', '-'}
	for _, c := range supportedChars {
		if r == c {
			return true
		}
	}
	return false
}

func (l *Lexer) readIdentifier() string {
	var result []rune

	for l.indentifierSupportedChar(l.ch) {
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

func (l *Lexer) skipComment() {
	if l.ch == '%' {
		l.readChar()
		// skip until the end of line
		for {
			if l.ch == '\n' {
				l.readChar()
				return
			}
			if l.ch == 0 {
				return
			}
			l.readChar()
		}
	}
}
