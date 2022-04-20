package ast

type Identifier struct {
	BaseNode
	Value string
}

func (identifier *Identifier) ToString() string {
	return identifier.Value
}

func (identifier *Identifier) GetExpressionNode() {}
