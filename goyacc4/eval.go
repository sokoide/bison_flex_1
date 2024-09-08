package main

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

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
	default:
		return 0, fmt.Errorf("%s not supported yet")
	}
}

func evaluateExpr(expr expression) (int, error) {
	log.Printf("expr %+v\n", expr)
	switch e := expr.(type) {

	case *numberExpression:
		v, err := strconv.Atoi(e.Lit)
		if err != nil {
			return 0, err
		}
		return v, nil
	case *parenExpression:
		v, err := evaluateExpr(e.SubExpr)
		if err != nil {
			return 0, err
		}
		return v, nil
	case *binOpExpression:
		lhsV, err := evaluateExpr(e.LHS)
		if err != nil {
			return 0, err
		}
		rhsV, err := evaluateExpr(e.RHS)
		if err != nil {
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
	case *assignExpression:
		// TODO
		return 0, nil
	default:
		panic(fmt.Sprintf("Unknown Expression type +%v", e))
	}

}
