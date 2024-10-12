package prolog

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

func evaluateRuleClause(program *Program, ruleClause *ruleClause, query term) []map[string]term {
	// when ````
	// write(X) :- builtin_write(X).
	// write(hello).```
	// is given, headUnification will have
	// map[string]term{"X": "hello"}.
	headUnification, ok := unify(ruleClause.HeadTerm, query)
	if !ok {
		return nil
	}

	return evaluateBodyTerms(program, ruleClause.BodyTerms, headUnification)
}

func evaluateBodyTerms(program *Program, bodyTerms []term, currentSubstitution map[string]term) []map[string]term {
	if len(bodyTerms) == 0 {
		return []map[string]term{currentSubstitution}
	}

	var solutions []map[string]term

	firstTerm := applySubstitution(bodyTerms[0], currentSubstitution)

	// builtin function handling
	firstTermSolutions := evaluateBuiltIns(program, firstTerm)

	for _, solution := range firstTermSolutions {
		combinedSubstitution := combineSubstitutions(currentSubstitution, solution)
		subSolutions := evaluateBodyTerms(program, bodyTerms[1:], combinedSubstitution)
		solutions = append(solutions, subSolutions...)
	}

	return solutions
}

func evaluateBuiltIns(program *Program, term1 term) []map[string]term {
	var solutions []map[string]term

	switch term1.GetFunctor() {
	case "builtin_write":
		fmt.Print(strings.Join(term1.GetArgs(), ""))
		solutions = []map[string]term{{}}
	case "builtin_nl":
		fmt.Println("")
		solutions = []map[string]term{{}}
	default:
		solutions = evaluateQuery(program, term1)
	}
	return solutions
}

func combineSubstitutions(sub1, sub2 map[string]term) map[string]term {
	result := make(map[string]term)
	for k, v := range sub1 {
		result[k] = v
	}
	for k, v := range sub2 {
		result[k] = v
	}
	return result
}

func unify(term1, term2 term) (map[string]term, bool) {
	log.Debugf("Unifying: %v with %v", term1, term2)
	substitution := make(map[string]term)
	if unifyHelper(term1, term2, substitution) {
		log.Debugf("Unification successful: %v", substitution)
		return substitution, true
	}
	log.Debugf("Unification failed")
	return nil, false
}

func unifyHelper(term1, term2 term, substitution map[string]term) bool {
	log.Debugf("UnifyHelper: %v with %v", term1, term2)
	term1 = deref(term1, substitution)
	term2 = deref(term2, substitution)
	log.Debugf("After deref: %v with %v", term1, term2)

	switch t1 := term1.(type) {
	case *constantTerm:
		switch t2 := term2.(type) {
		case *constantTerm:
			return t1.Lit == t2.Lit
		case *variableTerm:
			substitution[t2.Name] = t1
			return true
		}
	case *variableTerm:
		if t1.Name == "_" {
			return true // Anonymous variable unifies with anything
		}
		substitution[t1.Name] = term2
		return true
	case *compoundTerm:
		t2, ok := term2.(*compoundTerm)
		if !ok || t1.Functor != t2.Functor || len(t1.Args) != len(t2.Args) {
			return false
		}
		for i := range t1.Args {
			if !unifyHelper(t1.Args[i], t2.Args[i], substitution) {
				return false
			}
		}
		return true
	case *listTerm:
		return unifyLists(t1, term2, substitution)
	case *anonymousVarTerm:
		return true // Anonymous variable unifies with anything
	}
	log.Debugf("Unification failed: unknown term type")
	return false
}

func unifyLists(l1 *listTerm, term2 term, substitution map[string]term) bool {
	log.Debugf("Unifying lists: l1=%v, term2=%v", l1, term2)

	switch t2 := term2.(type) {
	case *listTerm:
		if l1.IsEmpty && t2.IsEmpty {
			return true
		}
		if l1.IsEmpty || t2.IsEmpty {
			return false
		}
		if l1.Head != nil && l1.Tail != nil {
			if t2.Head != nil && t2.Tail != nil {
				if len(l1.Head) != len(t2.Head) {
					return false
				}
				for i := range l1.Head {
					if !unifyHelper(l1.Head[i], t2.Head[i], substitution) {
						return false
					}
				}
				return unifyHelper(l1.Tail, t2.Tail, substitution)
			}
			// Handle case where l1 has Head/Tail and t2 has Args
			if len(t2.Args) < len(l1.Head) {
				return false
			}
			for i, h := range l1.Head {
				if !unifyHelper(h, t2.Args[i], substitution) {
					return false
				}
			}
			// The tail of l1 should unify with the rest of t2.Args
			return unifyHelper(l1.Tail, &listTerm{Args: t2.Args[len(l1.Head):]}, substitution)
		} else if l1.Args != nil {
			// Handle case where both have Args
			if len(l1.Args) != len(t2.Args) {
				return false
			}
			for i := range l1.Args {
				if !unifyHelper(l1.Args[i], t2.Args[i], substitution) {
					return false
				}
			}
			return true
		}
	case *variableTerm:
		substitution[t2.Name] = l1
		return true
	}
	return false
}

func deref(t term, substitution map[string]term) term {
	switch tt := t.(type) {
	case *variableTerm:
		if subst, ok := substitution[tt.Name]; ok {
			return deref(subst, substitution)
		}
	}
	return t
}

func applySubstitution(t term, substitution map[string]term) term {
	log.Debugf("Applying substitution to term: %v", t)
	switch tt := t.(type) {
	case *variableTerm:
		if subst, ok := substitution[tt.Name]; ok {
			return applySubstitution(subst, substitution)
		}
		return tt
	case *compoundTerm:
		newArgs := make([]term, len(tt.Args))
		for i, arg := range tt.Args {
			newArgs[i] = applySubstitution(arg, substitution)
		}
		return &compoundTerm{Functor: tt.Functor, Args: newArgs}
	case *listTerm:
		if tt.Head != nil && tt.Tail != nil {
			newHead := make([]term, len(tt.Head))
			for i, h := range tt.Head {
				newHead[i] = applySubstitution(h, substitution)
			}
			newTail := applySubstitution(tt.Tail, substitution)
			return &listTerm{Head: newHead, Tail: newTail}
		}
		if tt.Args != nil {
			newArgs := make([]term, len(tt.Args))
			for i, arg := range tt.Args {
				newArgs[i] = applySubstitution(arg, substitution)
			}
			return &listTerm{Args: newArgs}
		}
		// Empty list
		return tt
	case *constantTerm:
		return tt
	case *anonymousVarTerm:
		return tt
	default:
		log.Warnf("Unknown term type in applySubstitution: %T", t)
		return t
	}
}
