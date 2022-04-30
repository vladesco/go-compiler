package ast

type StringLiteral struct {
	BaseNode
	Value string
}

func (literal *StringLiteral) ToString() string {
	return literal.Token.Literal
}

func (literal *StringLiteral) GetExpressionNode() {}
