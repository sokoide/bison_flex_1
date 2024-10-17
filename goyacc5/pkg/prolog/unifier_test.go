package prolog

import (
	"reflect"
	"testing"
)

func TestUnifyConstant(t *testing.T) {
	term1 := &constantTerm{Lit: "hello"}
	term2 := &constantTerm{Lit: "hello"}
	term3 := &constantTerm{Lit: "world"}

	substitution, ok := unify(term1, term2)
	if !ok {
		t.Errorf("Unification of identical constants failed")
	}
	if len(substitution) != 0 {
		t.Errorf("Unexpected substitution for constant unification: %v", substitution)
	}

	_, ok = unify(term1, term3)
	if ok {
		t.Errorf("Unification of different constants succeeded unexpectedly")
	}
}

func TestUnifyVariable(t *testing.T) {
	term1 := &variableTerm{Name: "X"}
	term2 := &constantTerm{Lit: "hello"}

	substitution, ok := unify(term1, term2)
	if !ok {
		t.Errorf("Unification of variable with constant failed")
	}
	if subst, exists := substitution["X"]; !exists || subst.(*constantTerm).Lit != "hello" {
		t.Errorf("Expected X to be bound to 'hello', got %v", subst)
	}

	term3 := &variableTerm{Name: "Y"}
	substitution, ok = unify(term1, term3)
	if !ok {
		t.Errorf("Unification of two variables failed")
	}
	if len(substitution) != 1 {
		t.Errorf("Expected one substitution for variable unification, got %d", len(substitution))
	}
}

func TestUnifyCompound(t *testing.T) {
	term1 := &compoundTerm{Functor: "father", Args: []term{&variableTerm{Name: "X"}, &constantTerm{Lit: "john"}}}
	term2 := &compoundTerm{Functor: "father", Args: []term{&constantTerm{Lit: "bob"}, &constantTerm{Lit: "john"}}}

	substitution, ok := unify(term1, term2)
	if !ok {
		t.Errorf("Unification of compound terms failed")
	}
	if subst, exists := substitution["X"]; !exists || subst.(*constantTerm).Lit != "bob" {
		t.Errorf("Expected X to be bound to 'bob', got %v", subst)
	}

	term3 := &compoundTerm{Functor: "mother", Args: []term{&constantTerm{Lit: "alice"}, &constantTerm{Lit: "john"}}}
	_, ok = unify(term1, term3)
	if ok {
		t.Errorf("Unification of compound terms with different functors succeeded unexpectedly")
	}
}

func TestUnifyList(t *testing.T) {
	term1 := &listTerm{Args: []term{&constantTerm{Lit: "a"}, &constantTerm{Lit: "b"}, &variableTerm{Name: "X"}}}
	term2 := &listTerm{Args: []term{&constantTerm{Lit: "a"}, &constantTerm{Lit: "b"}, &constantTerm{Lit: "c"}}}

	substitution, ok := unify(term1, term2)
	if !ok {
		t.Errorf("Unification of lists failed")
	}
	if subst, exists := substitution["X"]; !exists || subst.(*constantTerm).Lit != "c" {
		t.Errorf("Expected X to be bound to 'c', got %v", subst)
	}

	emptyList1 := &listTerm{IsEmpty: true}
	emptyList2 := &listTerm{IsEmpty: true}
	substitution, ok = unify(emptyList1, emptyList2)
	if !ok {
		t.Errorf("Unification of empty lists failed")
	}
	if len(substitution) != 0 {
		t.Errorf("Unexpected substitution for empty list unification: %v", substitution)
	}

	headTailList1 := &listTerm{Head: []term{&variableTerm{Name: "H"}}, Tail: &variableTerm{Name: "T"}}
	headTailList2 := &listTerm{Args: []term{&constantTerm{Lit: "a"}, &constantTerm{Lit: "b"}, &constantTerm{Lit: "c"}}}

	substitution, ok = unify(headTailList1, headTailList2)
	if !ok {
		t.Errorf("Unification of head-tail list with regular list failed")
	}
	if subst, exists := substitution["H"]; !exists || subst.(*constantTerm).Lit != "a" {
		t.Errorf("Expected H to be bound to 'a', got %v", subst)
	}
	expectedTail := &listTerm{Args: []term{&constantTerm{Lit: "b"}, &constantTerm{Lit: "c"}}}
	if subst, exists := substitution["T"]; !exists || !reflect.DeepEqual(subst, expectedTail) {
		t.Errorf("Expected T to be bound to [b, c], got %v", subst)
	}
}

func TestUnifyAnonymousVariable(t *testing.T) {
	term1 := &anonymousVarTerm{}
	term2 := &constantTerm{Lit: "hello"}
	term3 := &variableTerm{Name: "X"}
	term4 := &compoundTerm{Functor: "test", Args: []term{&constantTerm{Lit: "a"}}}

	testCases := []term{term2, term3, term4}

	for _, tc := range testCases {
		substitution, ok := unify(term1, tc)
		if !ok {
			t.Errorf("Unification of anonymous variable with %v failed", tc)
		}
		if len(substitution) != 0 {
			t.Errorf("Unexpected substitution for anonymous variable unification: %v", substitution)
		}
	}
}

func TestUnifyComplexScenario(t *testing.T) {
	term1 := &compoundTerm{
		Functor: "family",
		Args: []term{
			&variableTerm{Name: "Parent"},
			&listTerm{
				Args: []term{
					&variableTerm{Name: "Child1"},
					&variableTerm{Name: "Child2"},
				},
			},
		},
	}

	term2 := &compoundTerm{
		Functor: "family",
		Args: []term{
			&constantTerm{Lit: "john"},
			&listTerm{
				Args: []term{
					&constantTerm{Lit: "alice"},
					&variableTerm{Name: "Y"},
				},
			},
		},
	}

	substitution, ok := unify(term1, term2)
	if !ok {
		t.Errorf("Unification in complex scenario failed")
	}

	expectedSubstitution := map[string]term{
		"Parent": &constantTerm{Lit: "john"},
		"Child1": &constantTerm{Lit: "alice"},
		"Child2": &variableTerm{Name: "Y"},
	}

	if !reflect.DeepEqual(substitution, expectedSubstitution) {
		t.Errorf("Unexpected substitution in complex scenario. Expected %v, got %v", expectedSubstitution, substitution)
	}
}
