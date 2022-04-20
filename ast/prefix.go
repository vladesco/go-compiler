package ast

import "bytes"

type PrefixExpression struct {
	BaseNode
	Operator string
	Right    Expression
}

func (expression *PrefixExpression) ToString() string {
	var output bytes.Buffer

	output.WriteString("(")
	output.WriteString(expression.Operator)
	output.WriteString(expression.Right.ToString())
	output.WriteString(")")

	return output.String()
}

func (expression *PrefixExpression) GetExpressionNode() {}
