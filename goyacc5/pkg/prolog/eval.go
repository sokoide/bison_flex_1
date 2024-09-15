package prolog

import (
	log "github.com/sirupsen/logrus"
)

type Program struct {
	Clauses []clause
}

func Load(lexer *Lexer) (*Program, error) {
	yyErrorVerbose = true
	parseResult := yyNewParser().Parse(lexer)
	log.Debugf("parseResult: %d", parseResult)
	return &Program{Clauses: lexer.program}, nil
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
			ret := evaluate(program, cl)
			log.Infof("result: %s => %v", cl.String(), ret)
		default:
			log.Debugf("skipping %v", cl.String())
		}
	}
	log.Info("query completed.")
	return nil
}

func evaluate(program *Program, fc *factClause) bool {
	for _, c := range program.Clauses {
		switch cl := c.(type) {
		case *factClause:
			log.Debugf("evaluating factClause: %s/%d, %v", cl.Fact.GetFunctor(), len(cl.Fact.GetArgs()), cl.String())
			if fc.String() == cl.String() {
				return true
			} else {
				log.Debugf("%s != %s", fc.String(), cl.String())
			}
		case *ruleClause:
			log.Debugf("evaluating ruleClause: %s/%d, %v", cl.HeadTerm.GetFunctor(), len(cl.HeadTerm.GetArgs()), cl.String())
			log.Debugf("evaluating ruleClause: %s/%d, %v", cl.HeadTerm.String(), len(cl.HeadTerm.GetArgs()), cl.String())
			if evaluateRuleClause(program, cl, fc.Head()) {
				return true
			}
		default:
			continue
		}
	}
	return false
}
