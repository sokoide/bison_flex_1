package main

import (
	"fmt"
	"reflect"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func evaluateStmts(stmts []expression) (int, error) {
	var ret int
	var err error

	for line, stmt := range stmts {
		ret, err = evaluateStmt(stmt)
		if err != nil {
			log.Errorf("[line: %d] Parser error: syntax error %v", line, err)
			return 0, err
		}
	}
	return ret, err
}

func evaluateStmt(stmt expression) (int, error) {
	var ret int
	var err error

	switch s := stmt.(type) {
	case *exprStatement:
		ret, err = evaluateExpr(s.Expr)
		if err == nil {
			log.Debugf("expr result: %d\n", ret)
			return 0, nil
		}
		return 0, err
	case *assignStatement:
		log.Debugf("assgin %s = %d\n", s.Name, ret)
		ret, err = evaluateExpr(s.Expr)
		if err == nil {
			log.Debugf("%s = %d\n", s.Name, ret)
			vars[s.Name] = ret
			return 0, nil
		}
		return 0, err
	case *putStatement:
		for _, expr := range s.Exprs {
			printExpr(expr)
		}
		fmt.Println()
		return 0, err
	case *whileStatement:
		log.Debug("whileStatement")
		for {
			ret, err = evaluateExpr(s.Cond)
			if err != nil {
				return 0, err
			}
			if ret > 0 {
				_, err = evaluateStmts(s.Body)
				if err != nil {
					return 0, err
				}
			} else {
				break
			}
		}
		return 0, err
	case *emptyStatement:
		log.Debug("emptyStatement")
		return 0, nil
	default:
		return 0, fmt.Errorf("%s not supported yet", reflect.TypeOf(stmt))
	}
}

func printExpr(expr expression) {
	log.Debugf("printExpr %+v\n", expr)

	var ret int
	var err error

	switch e := expr.(type) {
	case *numberExpression:
		fmt.Printf("%s", e.Lit)
	case *variableExpression:
		fmt.Printf("%d", vars[e.Lit])
	case *stringExpression:
		fmt.Printf("%s", e.Lit)
	default:
		ret, err = evaluateExpr(expr)
		if err == nil {
			fmt.Printf("%d", ret)
			return
		}
		log.Errorf("expr: %v failed to print", expr)
	}
}

func evaluateExpr(expr expression) (int, error) {
	log.Debugf("evaluateExpr %+v\n", expr)
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
	case *stringExpression:
		return 0, nil
	case *condExpression:
		// TODO: redundant
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
		case EQOP:
			if lhsV == rhsV {
				return 1, nil
			} else {
				return 0, nil
			}
		case NEOP:
			if lhsV != rhsV {
				return 1, nil
			} else {
				return 0, nil
			}
		case GEOP:
			if lhsV >= rhsV {
				return 1, nil
			} else {
				return 0, nil
			}
		case GTOP:
			if lhsV > rhsV {
				return 1, nil
			} else {
				return 0, nil
			}
		case LEOP:
			if lhsV <= rhsV {
				return 1, nil
			} else {
				return 0, nil
			}
		case LTOP:
			if lhsV < rhsV {
				return 1, nil
			} else {
				return 0, nil
			}
		default:
			return 0, fmt.Errorf("invalid Operator %v", e.Operator)
		}
	default:
		panic(fmt.Sprintf("Unknown Expression type %s for %+v", reflect.TypeOf(e), e))
	}
}
