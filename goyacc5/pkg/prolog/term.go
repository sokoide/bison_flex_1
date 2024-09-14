package prolog

import (
	"fmt"
	"strings"
)

// term interface
type term interface {
	String() string
	Evaluate(context map[string]term) term
}

// constantTerm: represents a constant (like "mary" or "pizza").
type constantTerm struct {
	Lit string // Literal value of the constant
}

func (c *constantTerm) String() string {
	return c.Lit
}

func (c *constantTerm) Evaluate(context map[string]term) term {
	// Constant terms evaluate to themselves
	return c
}

// variableTerm: represents a variable (like "X" or "Y").
type variableTerm struct {
	Name string // Name of the variable
}

func (v *variableTerm) String() string {
	return v.Name
}

func (v *variableTerm) Evaluate(context map[string]term) term {
	// Look up the variable in the evaluation context
	if val, ok := context[v.Name]; ok {
		return val
	}
	// If variable is unbound, it evaluates to itself
	return v
}

// compoundTerm (from previous code): represents compound terms like functors with arguments.
type compoundTerm struct {
	Functor string
	Args    []term
}

func (ct *compoundTerm) String() string {
	argsStr := []string{}
	for _, arg := range ct.Args {
		argsStr = append(argsStr, arg.String())
	}
	return fmt.Sprintf("%s(%s)", ct.Functor, strings.Join(argsStr, ", "))
}

func (ct *compoundTerm) Evaluate(context map[string]term) term {
	// Evaluate each argument in the context
	evaluatedArgs := []term{}
	for _, arg := range ct.Args {
		evaluatedArgs = append(evaluatedArgs, arg.Evaluate(context))
	}
	return &compoundTerm{
		Functor: ct.Functor,
		Args:    evaluatedArgs,
	}
}