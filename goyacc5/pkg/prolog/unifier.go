package prolog

func evaluateRuleClause(program *Program, ruleClause *ruleClause, query term) []map[string]term {
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
	firstTermSolutions := evaluateQuery(program, firstTerm)

	for _, solution := range firstTermSolutions {
		combinedSubstitution := combineSubstitutions(currentSubstitution, solution)
		subSolutions := evaluateBodyTerms(program, bodyTerms[1:], combinedSubstitution)
		solutions = append(solutions, subSolutions...)
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
	substitution := make(map[string]term)
	if unifyHelper(term1, term2, substitution) {
		return substitution, true
	}
	return nil, false
}

func unifyHelper(term1, term2 term, substitution map[string]term) bool {
	term1 = deref(term1, substitution)
	term2 = deref(term2, substitution)

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
		substitution[t1.Name] = term2
		return true
	case *compoundTerm:
		t2, ok := term2.(*compoundTerm)
		if !ok {
			return false
		}
		if t1.Functor != t2.Functor || len(t1.Args) != len(t2.Args) {
			return false
		}
		for i := range t1.Args {
			if !unifyHelper(t1.Args[i], t2.Args[i], substitution) {
				return false
			}
		}
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
	default:
		return t
	}
}
