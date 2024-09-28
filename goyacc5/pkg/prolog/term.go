package prolog

import (
	"fmt"
	"strings"
)

// term interface
type term interface {
	String() string
	Evaluate(context map[string]term) term
	GetFunctor() string
	GetArgs() []string
}

// constantTerm: represents a constant (like "scott" or "taro").
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

func (c *constantTerm) GetFunctor() string {
	return c.Lit
}

func (c *constantTerm) GetArgs() []string {
	return []string{}
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

func (v *variableTerm) GetFunctor() string {
	return v.Name
}

func (v *variableTerm) GetArgs() []string {
	return []string{}
}

// anonymousVarTerm: represents '_'
type anonymousVarTerm struct{}

func (a *anonymousVarTerm) String() string {
	return "_"
}

func (a *anonymousVarTerm) Evaluate(env map[string]term) term {
	// Anonymous variable evaluates to itself
	return a
}

func (a *anonymousVarTerm) GetFunctor() string {
	//  TODO
	return ""
}

func (a *anonymousVarTerm) GetArgs() []string {
	// TODOk
	return []string{}
}

// compoundTerm: represents compound terms like functors with arguments.
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

func (ct *compoundTerm) GetFunctor() string {
	return ct.Functor
}

func (ct *compoundTerm) GetArgs() []string {
	args := []string{}
	for _, arg := range ct.Args {
		args = append(args, arg.GetFunctor())
	}
	return args
}

// listTerm: represents Lists of terms
type listTerm struct {
	Args []term
	Head term
	Tail term
}

func (l *listTerm) String() string {
	argsStr := []string{}
	for _, arg := range l.Args {
		argsStr = append(argsStr, arg.String())
	}
	return "[" + strings.Join(argsStr, ", ") + "]"
}

func (l *listTerm) Evaluate(context map[string]term) term {
	// listTerm evaluates to itself
	return l
}

func (l *listTerm) GetFunctor() string {
	// TODO
	return ""
}

func (l *listTerm) GetArgs() []string {
	argsStr := []string{}
	for _, arg := range l.Args {
		argsStr = append(argsStr, arg.String())
	}
	return argsStr
}
