package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	scopeStack = NewScopeStack()
	source := `a=1;b=1+5*2+9/3;a=a+1;`
	scanner := new(scanner)
	scanner.Init(source)

	var prog []expression = parse(scanner)
	_, err := evaluateStmts(prog)
	if err != nil {
		panic(err)
	}

	var got int
	var scope *Scope
	got, scope = scopeStack.Get("a")
	assert.NotNil(t, scope)
	assert.Equal(t, 2, got)

	got, scope = scopeStack.Get("b")
	assert.NotNil(t, scope)
	assert.Equal(t, 14, got)
}

func TestWhile(t *testing.T) {
	scopeStack = NewScopeStack()
	source := `a=0;while(a<5){a=a+1;}`
	scanner := new(scanner)
	scanner.Init(source)

	var prog []expression = parse(scanner)
	_, err := evaluateStmts(prog)
	if err != nil {
		panic(err)
	}

	var got int
	var scope *Scope
	got, scope = scopeStack.Get("a")
	assert.NotNil(t, scope)
	assert.Equal(t, 5, got)
}

func TestScope(t *testing.T) {
	scopeStack = NewScopeStack()
	// notice b doesn't exist outside of while
	source := `a=0;while(a<5){a=a+1;b=b+1;}`
	scanner := new(scanner)
	scanner.Init(source)

	var prog []expression = parse(scanner)
	_, err := evaluateStmts(prog)
	if err != nil {
		panic(err)
	}

	var got int
	var scope *Scope
	got, scope = scopeStack.Get("a")
	assert.NotNil(t, scope)
	assert.Equal(t, 5, got)
	got, scope = scopeStack.Get("b")
	assert.Nil(t, scope)
	assert.Equal(t, 0, got)
}

func TestDivisionByZero(t *testing.T) {
	scopeStack = NewScopeStack()
	source := `a=5/0;`
	scanner := new(scanner)
	scanner.Init(source)

	var prog []expression = parse(scanner)
	_, err := evaluateStmts(prog)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "division by zero")
}

func TestModuloByZero(t *testing.T) {
	scopeStack = NewScopeStack()
	source := `a=5%0;`
	scanner := new(scanner)
	scanner.Init(source)

	var prog []expression = parse(scanner)
	_, err := evaluateStmts(prog)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "modulo by zero")
}
