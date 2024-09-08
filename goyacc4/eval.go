package main

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func evaluateStmts(stmts []statement) (int, error) {
	var ret int
	var err error

	for line, stmt := range stmts {
		ret, err = evaluateStmt(stmt)
		if err != nil {
			log.Errorf("[line: %d] syntax error %v", line, err)
			return 0, err
		}
	}
	return ret, err
}

func evaluateStmt(stmt statement) (int, error) {
	var ret int
	var err error

	switch s := stmt.(type) {
	case *exprStatement:
		ret, err = evaluateExpr(s.Expr)
		if err == nil {
			log.Printf("expr result: %d\n", ret)
			return 0, nil
		}
		return 0, err
	case *assignStatement:
		log.Printf("assgin %s = %d\n", s.Name, ret)
		ret, err = evaluateExpr(s.Expr)
		if err == nil {
			log.Printf("%s = %d\n", s.Name, ret)
			vars[s.Name] = ret
			return 0, nil
		}
		return 0, err
	default:
		return 0, fmt.Errorf("%s not supported yet")
	}
}

func evaluateExpr(expr expression) (int, error) {
	log.Debugf("expr %+v\n", expr)
	switch e := expr.(type) {

	case *numberExpression:
		v, err := strconv.Atoi(e.Lit)
		if err != nil {
			log.Warnf("err: %v", err)
			return 0, err
		}
		return v, nil
	case *parenExpression:
		v, err := evaluateExpr(e.SubExpr)
		if err != nil {
			log.Warnf("err: %v", err)
			return 0, err
		}
		return v, nil
	case *binOpExpression:
		lhsV, err := evaluateExpr(e.LHS)
		if err != nil {
			log.Warnf("err: %v", err)
			return 0, err
		}
		rhsV, err := evaluateExpr(e.RHS)
		if err != nil {
			log.Warnf("err: %v", err)
			return 0, err
		}
		switch e.Operator {
		case '+':
			return lhsV + rhsV, nil
		case '-':
			return lhsV - rhsV, nil
		case '*':
			return lhsV * rhsV, nil
		case '/':
			return lhsV / rhsV, nil
		case '%':
			return lhsV % rhsV, nil
		default:
			panic("Unknown operator")
		}
	case *variableExpression:
		if v, ok := vars[e.Lit]; ok {
			return v, nil
		}
		vars[e.Lit] = 0
		log.Warnf("err: variable %s not found", e.Lit)
		return 0, nil
	default:
		panic(fmt.Sprintf("Unknown Expression type +%v", e))
	}

}
