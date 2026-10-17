package prolog

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempProgram(t *testing.T, src string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "prog.pro")
	if err := os.WriteFile(path, []byte(src), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

// A lexer error must be surfaced through Lexer.Err() and stop the parse,
// instead of being swallowed or silently truncating the program.
func TestUnterminatedStringIsReported(t *testing.T) {
	lexer, err := NewLexer(writeTempProgram(t, "foo(\"bar\n"))
	if err != nil {
		t.Fatal(err)
	}
	defer lexer.Close()

	yyErrorVerbose = true
	result := yyNewParser().Parse(lexer)

	if result == 0 {
		t.Errorf("parse unexpectedly succeeded on unterminated string")
	}
	if lexer.Err() == nil {
		t.Fatalf("Lexer.Err() is nil, want an unterminated-string error")
	}
	if got := lexer.Err().Error(); got != "unterminated string literal" {
		t.Errorf("Lexer.Err() = %q, want %q", got, "unterminated string literal")
	}
}

func TestIllegalCharacterIsReported(t *testing.T) {
	lexer, err := NewLexer(writeTempProgram(t, "foo(?)\n"))
	if err != nil {
		t.Fatal(err)
	}
	defer lexer.Close()

	yyErrorVerbose = true
	result := yyNewParser().Parse(lexer)

	if result == 0 {
		t.Errorf("parse unexpectedly succeeded on an illegal character")
	}
	if lexer.Err() == nil {
		t.Fatalf("Lexer.Err() is nil, want an unexpected-token error")
	}
}
