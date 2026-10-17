package main

import (
	"fmt"
	"strconv"
)

func evaluate(expr expression) (int, error) {
	switch e := expr.(type) {

	case *numberExpression:
		v, err := strconv.Atoi(e.Lit)
		if err != nil {
			return 0, err
		}
		return v, nil
	case *parenExpression:
		v, err := evaluate(e.SubExpr)
		if err != nil {
			return 0, err
		}
		return v, nil
	case *binOpExpression:
		lhsV, err := evaluate(e.LHS)
		if err != nil {
			return 0, err
		}
		rhsV, err := evaluate(e.RHS)
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
			if rhsV == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return lhsV / rhsV, nil
		case '%':
			if rhsV == 0 {
				return 0, fmt.Errorf("modulo by zero")
			}
			return lhsV % rhsV, nil
		default:
			return 0, fmt.Errorf("unknown operator: %d", e.Operator)
		}
	default:
		return 0, fmt.Errorf("unknown expression type: %T", expr)
	}

}
