package prolog

import (
	"fmt"
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Program struct {
	Clauses []clause
}

func Load(lexer *Lexer) (*Program, error) {
	yyErrorVerbose = true

	// parse builtin program
	builtinLexer, err := NewLexer("resource/builtin.pro")
	if err != nil {
		log.Errorf("failed to load builtin.pro, err: %s", err)
		return &Program{}, err
	}
	parseResult := yyNewParser().Parse(builtinLexer)
	log.Debugf("builtin parseResult: %d", parseResult)

	// parse user program
	parseResult = yyNewParser().Parse(lexer)
	log.Debugf("user parseResult: %d", parseResult)

	// combine the 2
	clauses := append(builtinLexer.program, lexer.program...)

	log.Debug("Loaded clauses:")
	for i, c := range clauses {
		log.Debugf("  [%d] %v", i, c)
	}

	return &Program{Clauses: clauses}, nil
}

func Dump(program *Program) {
	for _, c := range program.Clauses {
		c.Dump()
	}
}

func Query(program *Program, queryProgram *Program) error {
	log.Info("querying...")
	for _, c := range queryProgram.Clauses {
		switch cl := c.(type) {
		case *factClause:
			log.Debugf("evaluating: %s/%d, %v", cl.Fact.GetFunctor(), len(cl.Fact.GetArgs()), cl.String())
			solutions := evaluateQuery(program, cl.Fact)
			result := false
			if len(solutions) > 0 {
				result = true
			}
			PrintQueryResults(cl.Fact, solutions)

			log.Infof("%v -> %v", cl.String(), result)
			log.Debugf("results for %s:", cl.String())
			for _, solution := range solutions {
				log.Debugf("  solution: %v", solution)
			}
		default:
			log.Debugf("skipping %v, %v", reflect.TypeOf(cl), cl.String())
		}
	}
	log.Info("query completed.")
	return nil
}

func evaluateQuery(program *Program, query term) []map[string]term {
	var solutions []map[string]term

	log.Debugf("Evaluating query: %v", query)

	queryPredicate, ok := query.(*compoundTerm)
	if !ok {
		log.Errorf("Query is not a compound term: %v", query)
		return nil
	}

	for _, clause := range program.Clauses {
		var clauseHead term
		switch c := clause.(type) {
		case *factClause:
			clauseHead = c.Fact
		case *ruleClause:
			clauseHead = c.HeadTerm
		default:
			continue
		}

		clausePredicate, ok := clauseHead.(*compoundTerm)
		if !ok {
			continue
		}

		// Only attempt unification if the predicates match
		if queryPredicate.Functor == clausePredicate.Functor && len(queryPredicate.Args) == len(clausePredicate.Args) {
			log.Debugf("Trying to unify with clause: %v", clause)

			switch c := clause.(type) {
			case *factClause:
				log.Debugf("Attempting to unify %v with %v", query, c.Fact)
				if unification, ok := unify(c.Fact, query); ok {
					log.Debugf("Unification successful: %v", unification)
					solutions = append(solutions, unification)
				} else {
					log.Debugf("Unification failed")
				}
			case *ruleClause:
				log.Debugf("Evaluating rule: %v", c)
				ruleSolutions := evaluateRuleClause(program, c, query)
				log.Debugf("Rule evaluation resulted in %d solutions", len(ruleSolutions))
				solutions = append(solutions, ruleSolutions...)
			}
		}
	}

	log.Debugf("Query evaluation completed with %d solutions", len(solutions))
	return solutions
}

func PrintQueryResults(query term, solutions []map[string]term) {
	if len(solutions) == 0 {
		fmt.Println("No solutions found.")
		return
	}

	for _, solution := range solutions {
		printSolution(query, solution)
	}
}

func printSolution(query term, solution map[string]term) {
	switch q := query.(type) {
	case *compoundTerm:
		for i, arg := range q.Args {
			if v, ok := arg.(*variableTerm); ok {
				if value, exists := solution[v.Name]; exists {
					fmt.Printf("%s = %v", v.Name, formatTerm(value))
					if i < len(q.Args)-1 {
						fmt.Print(", ")
					}
				}
			}
		}
		fmt.Println()
	default:
		fmt.Printf("%v\n", formatTerm(applySubstitution(query, solution)))
	}
}

func formatTerm(t term) string {
	switch tt := t.(type) {
	case *constantTerm:
		return tt.Lit
	case *variableTerm:
		return tt.Name
	case *compoundTerm:
		args := make([]string, len(tt.Args))
		for i, arg := range tt.Args {
			args[i] = formatTerm(arg)
		}
		return fmt.Sprintf("%s(%s)", tt.Functor, strings.Join(args, ", "))
	case *listTerm:
		if tt.Head != nil && tt.Tail != nil {
			headTerms := make([]string, len(tt.Head))
			for i, h := range tt.Head {
				headTerms[i] = formatTerm(h)
			}
			return fmt.Sprintf("[%s|%s]", strings.Join(headTerms, ", "), formatTerm(tt.Tail))
		}
		args := make([]string, len(tt.Args))
		for i, arg := range tt.Args {
			args[i] = formatTerm(arg)
		}
		return fmt.Sprintf("[%s]", strings.Join(args, ", "))
	default:
		return fmt.Sprintf("%v", t)
	}
}
