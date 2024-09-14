package prolog

import (
	"fmt"
)

type token struct {
	Type  string // Token type (e.g., IDENT, NUMBER,...)
	Value string // The literal value of the token
}

func (t *token) String() string {
	return fmt.Sprintf("%s(%s)", t.Type, t.Value)
}
