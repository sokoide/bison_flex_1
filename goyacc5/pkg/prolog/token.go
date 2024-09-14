package prolog

import (
	"fmt"
)

type token struct {
	Type  tokenType
	Value string // The literal value of the token
}

func (t *token) String() string {
	return fmt.Sprintf("%d(%s)", t.Type, t.Value)
}
