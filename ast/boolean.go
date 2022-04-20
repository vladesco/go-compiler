package ast

type Boolean struct {
	BaseNode
	Value bool
}

func (boolean *Boolean) ToString() string {
	return boolean.Token.Literal
}

func (boolean *Boolean) GetExpressionNode() {}
