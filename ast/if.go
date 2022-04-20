package ast

import "bytes"

type IfExpression struct {
	BaseNode
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (expression *IfExpression) ToString() string {
	var output bytes.Buffer

	output.WriteString("if")
	output.WriteString(expression.Condition.ToString())
	output.WriteString(" ")
	output.WriteString(expression.Consequence.ToString())

	if expression.Alternative != nil {
		output.WriteString("else ")
		output.WriteString((expression.Alternative.ToString()))
	}

	return output.String()
}

func (expression *IfExpression) GetExpressionNode() {}
