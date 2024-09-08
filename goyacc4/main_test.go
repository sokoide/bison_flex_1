package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	source := `a=1;b=1+5*2+9/3;`
	scanner := new(scanner)
	scanner.Init(source)

	var prog []statement = parse(scanner)
	_, err := evaluateStmts(prog)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, vars["a"])
	assert.Equal(t, 14, vars["b"])
}
