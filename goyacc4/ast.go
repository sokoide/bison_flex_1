package main

// expressions
type (
	expression interface {
		expression()
	}
)

type (
	numberExpression struct {
		Lit string
	}

	parenExpression struct {
		SubExpr expression
	}

	binOpExpression struct {
		LHS      expression
		Operator int
		RHS      expression
	}

	variableExpression struct {
		Lit string
	}

	stringExpression struct {
		Lit string
	}
	condExpression struct {
		LHS      expression
		Operator int
		RHS      expression
	}
)

func (x *numberExpression) expression()   {}
func (x *parenExpression) expression()    {}
func (x *binOpExpression) expression()    {}
func (x *variableExpression) expression() {}
func (x *stringExpression) expression()   {}
func (x *condExpression) expression()     {}

// statements
type (
	emptyStatement struct {
	}

	exprStatement struct {
		Expr expression
	}

	assignStatement struct {
		Name string
		Expr expression
	}

	putStatement struct {
		Exprs []expression
	}

	whileStatement struct {
		Cond expression
		Body []expression
	}
)

func (x *emptyStatement) expression()  {}
func (x *exprStatement) expression()   {}
func (x *assignStatement) expression() {}
func (x *putStatement) expression()    {}
func (x *whileStatement) expression()  {}
