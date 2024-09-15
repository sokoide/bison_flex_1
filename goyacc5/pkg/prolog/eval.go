package prolog

import (
	log "github.com/sirupsen/logrus"
)

type Program struct {
	Clauses []clause
	// Facts   []*factClause
	// Rules   []*ruleClause
}

func Load(lexer *Lexer) (*Program, error) {
	yyErrorVerbose = true
	parseResult := yyNewParser().Parse(lexer)
	log.Debugf("parseResult: %d", parseResult)

	// rules := make([]*ruleClause, 0)
	// facts := make([]*factClause, 0)

	// for _, c := range lexer.program {
	// 	switch cl := c.(type) {
	// 	case *factClause:
	// 		facts = append(facts, cl)
	// 	case *ruleClause:
	// 		rules = append(rules, cl)
	// 	default:
	// 		err := fmt.Errorf("Only fact or rule clause is expected. %+v given", cl)
	// 		log.Error(err)
	// 		return nil, err
	// 	}
	// }

	// return &Program{
	// 	Facts: facts,
	// 	Rules: rules,
	// }, nil
	return &Program{Clauses: lexer.program}, nil
}

func Dump(program *Program) {
	for _, c := range program.Clauses {
		c.Dump()
	}
}

func Evaluate(program *Program) error {
	return nil
}
