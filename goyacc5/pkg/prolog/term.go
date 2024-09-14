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

// func main() {
// 	// Example of a fact: likes(mary, pizza).
// 	fact := &factClause{
// 		Fact: &compoundTerm{
// 			Functor: "likes",
// 			Args: []term{
// 				&constantTerm{Lit: "mary"},
// 				&constantTerm{Lit: "pizza"},
// 			},
// 		},
// 	}
// 	fmt.Println(fact.String()) // Output: likes(mary, pizza).

// 	// Example of a rule: likes(X, Y) :- friend(X, Y).
// 	rule := &ruleClause{
// 		HeadTerm: &compoundTerm{
// 			Functor: "likes",
// 			Args: []term{
// 				&variableTerm{Name: "X"},
// 				&variableTerm{Name: "Y"},
// 			},
// 		},
// 		BodyTerms: []term{
// 			&compoundTerm{
// 				Functor: "friend",
// 				Args: []term{
// 					&variableTerm{Name: "X"},
// 					&variableTerm{Name: "Y"},
// 				},
// 			},
// 		},
// 	}
// 	fmt.Println(rule.String()) // Output: likes(X, Y) :- friend(X, Y).

// 	// Token example
// 	tok := &token{Type: "IDENT", Value: "likes"}
// 	fmt.Println(tok.String()) // Output: IDENT(likes)
// }
