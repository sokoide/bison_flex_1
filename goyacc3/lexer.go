package main

const (
	// EOF end of file
	EOF = -1
)

// Token-value contract for the parser:
//   - Lex() returns 0 or a negative value to mean EOF (goyacc's $end,
//     stored in yyTok1[0]). It must NEVER return 0 for a real token,
//     otherwise the parser silently stops as if the input was over.
// - Any token number returned must be > 0 and must be a token declared
//   in calc.go.y. Unrecognized characters therefore return ILLEGAL.
//   - EOF (-1 here) is converted to 0 in Lex().

type token struct {
	tok int
	lit string
	pos position
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

func (s *scanner) Scan() (tok int, lit string, pos position) {
	s.skipWhiteSpace()
	pos = s.position()
	switch ch := s.peek(); {
	case isDigit(ch):
		tok, lit = NUMBER, s.scanNumber()
	default:
		switch ch {
		case -1:
			tok = EOF
		case '(', ')', ';', '+', '-', '*', '/', '%', '=':
			tok = int(ch)
			lit = string(ch)
		default:
			// illegal character: report ILLEGAL (>0) so the parser
			// fails with a syntax error instead of silently
			// treating it as end-of-input (tok 0 is EOF for goyacc).
			tok = ILLEGAL
			lit = string(ch)
		}
		s.next()
	}
	return
}

// ========================================

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (s *scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	} else {
		return -1
	}
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

func (s *scanner) scanNumber() string {
	var ret []rune
	for isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}
