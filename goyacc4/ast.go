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

	variableExpression struct {
		Lit string
	}
)

func (x *numberExpression) expression()   {}
func (x *parenExpression) expression()    {}
func (x *binOpExpression) expression()    {}
func (x *variableExpression) expression() {}

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

	assignStatement struct {
		Name string
		Expr expression
	}

	putStatement struct {
		Exprs []expression
	}
)

func (x *nullStatement) statement() {}

func (x *exprStatement) statement() {}

func (x *assignStatement) statement() {}

func (x *putStatement) statement() {}
