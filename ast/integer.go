package ast

type IntegerLiteral struct {
	BaseNode
	Value int64
}

func (literal *IntegerLiteral) ToString() string {
	return literal.Token.Literal
}

func (literal *IntegerLiteral) GetExpressionNode() {}
