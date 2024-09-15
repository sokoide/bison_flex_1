package prolog

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

func evaluateRuleClause(program *Program, ruleClause *ruleClause, query term) bool {
	// Check if the head of the rule unifies with the query
	headUnification, ok := unify(ruleClause.HeadTerm, query)
	if !ok {
		return false
	}

	// Evaluate each term in the body of the rule
	for _, bodyTerm := range ruleClause.BodyTerms {
		substitutedBodyTerm := applySubstitution(bodyTerm, headUnification)

		if !evaluate(program, &factClause{Fact: substitutedBodyTerm}) {
			return false
		}
	}

	return true
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
