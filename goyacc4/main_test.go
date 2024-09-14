package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	source := `a=1;b=1+5*2+9/3;a=a+1;`
	scanner := new(scanner)
	scanner.Init(source)

	var prog []expression = parse(scanner)
	_, err := evaluateStmts(prog)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 2, vars["a"])
	assert.Equal(t, 14, vars["b"])
}

func TestWhile(t *testing.T) {
	source := `a=0;while(a<5){a=a+1;}`
	scanner := new(scanner)
	scanner.Init(source)

	var prog []expression = parse(scanner)
	_, err := evaluateStmts(prog)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 5, vars["a"])
}
