package prolog

import (
	"reflect"

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

	for _, clause := range program.Clauses {
		switch cl := clause.(type) {
		case *factClause:
			if unification, ok := unify(query, cl.Fact); ok {
				solutions = append(solutions, unification)
			}
		case *ruleClause:
			ruleSolutions := evaluateRuleClause(program, cl, query)
			solutions = append(solutions, ruleSolutions...)
		}
	}

	return solutions
}
