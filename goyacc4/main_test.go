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

	prog, err := parse(scanner)
	if err != nil {
		t.Fatal(err)
	}
	_, err = evaluateStmts(prog)
	if err != nil {
		t.Fatal(err)
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

	prog, err := parse(scanner)
	if err != nil {
		t.Fatal(err)
	}
	_, err = evaluateStmts(prog)
	if err != nil {
		t.Fatal(err)
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

	prog, err := parse(scanner)
	if err != nil {
		t.Fatal(err)
	}
	_, err = evaluateStmts(prog)
	if err != nil {
		t.Fatal(err)
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

func TestUndefinedVariable(t *testing.T) {
	scopeStack = NewScopeStack()
	// "reading an undefined variable yields 0" applies both inside
	// expressions (a=x+1) and to put() (put(x)), so there is no way to
	// crash with an unknown identifier
	source := `a=x+1;put(x);`
	scanner := new(scanner)
	scanner.Init(source)

	prog, err := parse(scanner)
	if err != nil {
		t.Fatal(err)
	}
	_, err = evaluateStmts(prog)
	if err != nil {
		t.Fatal(err)
	}

	var got int
	var scope *Scope
	got, scope = scopeStack.Get("x")
	assert.NotNil(t, scope)
	assert.Equal(t, 0, got)

	got, scope = scopeStack.Get("a")
	assert.NotNil(t, scope)
	assert.Equal(t, 1, got)
}

func TestStringInArithmetic(t *testing.T) {
	scopeStack = NewScopeStack()
	source := `a="x"+"y";`
	scanner := new(scanner)
	scanner.Init(source)

	prog, err := parse(scanner)
	if err != nil {
		t.Fatal(err)
	}
	_, err = evaluateStmts(prog)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not supported in arithmetic")
}

func TestDivisionByZero(t *testing.T) {
	scopeStack = NewScopeStack()
	source := `a=5/0;`
	scanner := new(scanner)
	scanner.Init(source)

	prog, err := parse(scanner)
	if err != nil {
		t.Fatal(err)
	}
	_, err = evaluateStmts(prog)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "division by zero")
}

func TestModuloByZero(t *testing.T) {
	scopeStack = NewScopeStack()
	source := `a=5%0;`
	scanner := new(scanner)
	scanner.Init(source)

	prog, err := parse(scanner)
	if err != nil {
		t.Fatal(err)
	}
	_, err = evaluateStmts(prog)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "modulo by zero")
}

func TestSyntaxError(t *testing.T) {
	// no unary minus in the language: `-1` terminates the parse with an
	// error instead of exiting the process
	scopeStack = NewScopeStack()
	source := `put(-1);`
	scanner := new(scanner)
	scanner.Init(source)

	_, err := parse(scanner)
	assert.Error(t, err)
}

func TestIllegalCharacter(t *testing.T) {
	// an unknown character becomes the declared ILLEGAL token, so the parser
	// reports a syntax error instead of silently stopping at it (token 0
	// would be EOF to goyacc)
	scopeStack = NewScopeStack()
	scanner := new(scanner)
	scanner.Init(`put(1);@put(2);`)

	_, err := parse(scanner)
	assert.Error(t, err)
}
