package prolog

import "testing"

func TestUnify(t *testing.T) {
	term1 := &compoundTerm{Functor: "meal", Args: []term{&variableTerm{Name: "X"}}}
	term2 := &compoundTerm{Functor: "meal", Args: []term{&constantTerm{Lit: "orange"}}}

	substitution, ok := unify(term1, term2)

	if !ok {
		t.Errorf("Unification failed, expected success")
	}

	if subst, exists := substitution["X"]; !exists || subst.(*constantTerm).Lit != "orange" {
		t.Errorf("Expected X to be bound to 'orange', got %v", subst)
	}
}
