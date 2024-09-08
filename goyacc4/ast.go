package main

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

	assignExpression struct {
	}
)

func (x *numberExpression) expression() {}
func (x *parenExpression) expression()  {}
func (x *binOpExpression) expression()  {}
func (x *assignExpression) expression() {}

type (
	statement interface {
		statement()
	}
)

type (
	nullStatement struct {
	}
	exprStatement struct {
		Expr expression
	}
)

func (x *nullStatement) statement() {}

func (x *exprStatement) statement() {}
