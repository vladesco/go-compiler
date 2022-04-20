package ast

import "bytes"

type InfixExpression struct {
	BaseNode
	Left     Expression
	Operator string
	Right    Expression
}

func (expression *InfixExpression) ToString() string {
	var output bytes.Buffer

	output.WriteString("(")
	output.WriteString(expression.Left.ToString())
	output.WriteString(expression.Operator)
	output.WriteString(expression.Right.ToString())
	output.WriteString(")")

	return output.String()
}

func (expression *InfixExpression) GetExpressionNode() {}
