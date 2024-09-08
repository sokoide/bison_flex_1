package main

const (
	// EOF end of file
	EOF = -1
	// UNKNOWN unknown token
	UNKNOWN = 0 * iota
)

var keywords = map[string]int{
	"put": PUT,
}

// tokenType enum
type tokenType int

const (
	tokenTypeNone tokenType = iota
	tokenTypeNumberLiteral
	tokenTypeStringLiteral
	tokenTypeIdent
	tokenTypeKeyword
	tokenTypeOp
)

type token struct {
	tok int
	lit string
	pos position
	tt  tokenType
}

type position struct {
	Line   int
	Column int
}

type scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

func (s *scanner) Init(src string) {
	s.src = []rune(src)
}

func (s *scanner) Scan() (tok int, lit string, pos position, tt tokenType) {
	s.skipWhiteSpace()
	pos = s.position()
	switch ch := s.peek(); {
	case isDigit(ch):
		tt = tokenTypeNumberLiteral
		tok, lit = NUMBER_LITERAL, s.scanNumber()
	case isLetter(ch):
		tt = tokenTypeIdent
		tok, lit = IDENT, s.scanIdentifier()
		if t, ok := keywords[lit]; ok {
			tok = t
			tt = tokenTypeKeyword
		}
	default:
		switch ch {
		case -1:
			tok = EOF
		case '.', '[', ']', '(', ')', ';', '+', '-', '*', '/', '%', '=':
			tok = int(ch)
			lit = string(ch)
			tt = tokenTypeOp
		}
		s.next()
	}
	return
}

// ========================================

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (s *scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	}
	return -1
}

func (s *scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

func (s *scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

func (s *scanner) position() position {
	return position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

func (s *scanner) skipWhiteSpace() {
	for isWhiteSpace(s.peek()) {
		s.next()
	}
}

func (s *scanner) scanIdentifier() string {
	var ret []rune
	for isLetter(s.peek()) || isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}

func (s *scanner) scanNumber() string {
	var ret []rune
	for isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}
