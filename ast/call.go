package ast

import (
	"bytes"
	"strings"
)

type CallExpression struct {
	BaseNode
	Function  Expression
	Arguments []Expression
}

func (expression *CallExpression) ToString() string {
	var output bytes.Buffer
	var arguments = []string{}

	for _, argument := range expression.Arguments {
		arguments = append(arguments, argument.ToString())
	}

	output.WriteString(expression.Function.ToString())
	output.WriteString("(")
	output.WriteString(strings.Join(arguments, ","))
	output.WriteString(")")

	return output.String()
}

func (expression *CallExpression) GetExpressionNode() {}
