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
			log.Errorf("[stmt %d] evaluation error: %v", line, err)
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
		ret, err = evaluateExpr(s.Expr)
		if err == nil {
			if _, scope := scopeStack.Get(s.Name); scope != nil {
				scope.Set(s.Name, ret)
			} else {
				log.Debugf("%s doesn't exist in all scopeStack. Making it in the current scope", s.Name)
				scopeStack.Set(s.Name, ret)
			}
			log.Debugf("%s = %d\n", s.Name, ret)
			return 0, nil
		}
		return 0, err
	case *putStatement:
		for _, expr := range s.Exprs {
			if err := printExpr(expr); err != nil {
				return 0, err
			}
		}
		fmt.Println()
		return 0, nil
	case *whileStatement:
		log.Debug("whileStatement - scope pushed")
		scopeStack.PushScope()
		for {
			ret, err = evaluateExpr(s.Cond)
			if err != nil {
				log.Debug("whileStatement -scope poped")
				scopeStack.PopScope()
				return 0, err
			}
			if ret > 0 {
				_, err = evaluateStmts(s.Body)
				if err != nil {
					log.Debug("whileStatement -scope poped")
					scopeStack.PopScope()
					return 0, err
				}
			} else {
				break
			}
		}
		log.Debug("whileStatement -scope poped")
		scopeStack.PopScope()
		return 0, err
	case *emptyStatement:
		log.Debug("emptyStatement")
		return 0, nil
	default:
		return 0, fmt.Errorf("%s not supported yet", reflect.TypeOf(stmt))
	}
}

func printExpr(expr expression) error {
	log.Debugf("printExpr %+v\n", expr)

	var ret int
	var err error
	var scope *Scope

	switch e := expr.(type) {
	case *numberExpression:
		fmt.Printf("%s", e.Lit)
	case *variableExpression:
		if ret, scope = scopeStack.Get(e.Lit); scope != nil {
			fmt.Printf("%d", ret)
		} else {
			// Same semantics as evaluateExpr: reading an undefined
			// variable yields 0 (and auto-creates it in the current
			// scope) instead of failing.
			scopeStack.Set(e.Lit, 0)
			fmt.Printf("%d", 0)
		}
	case *stringExpression:
		fmt.Printf("%s", e.Lit)
	default:
		ret, err = evaluateExpr(expr)
		if err != nil {
			return err
		}
		fmt.Printf("%d", ret)
	}
	return nil
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
			return 0, fmt.Errorf("unknown operator: %v", e.Operator)
		}
	case *variableExpression:
		var v int
		var scope *Scope
		if v, scope = scopeStack.Get(e.Lit); scope != nil {
			return v, nil
		}
		// Reading an undefined variable yields 0 (and auto-creates it in
		// the current scope so assignments like `b=b+1;` inside a while
		// body still work).
		scopeStack.Set(e.Lit, 0)
		log.Warnf("variable %s not found, 0 assigned", e.Lit)
		return 0, nil
	case *stringExpression:
		return 0, fmt.Errorf("string %q not supported in arithmetic", e.Lit)
	case *condExpression:
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
			}
			return 0, nil
		case NEOP:
			if lhsV != rhsV {
				return 1, nil
			}
			return 0, nil
		case GEOP:
			if lhsV >= rhsV {
				return 1, nil
			}
			return 0, nil
		case GTOP:
			if lhsV > rhsV {
				return 1, nil
			}
			return 0, nil
		case LEOP:
			if lhsV <= rhsV {
				return 1, nil
			}
			return 0, nil
		case LTOP:
			if lhsV < rhsV {
				return 1, nil
			}
			return 0, nil
		default:
			return 0, fmt.Errorf("invalid Operator %v", e.Operator)
		}
	default:
		return 0, fmt.Errorf("unknown expression type %s for %+v", reflect.TypeOf(e), e)
	}
}
