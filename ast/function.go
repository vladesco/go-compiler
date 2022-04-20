package ast

import (
	"bytes"
	"strings"
)

type FunctionLiteral struct {
	BaseNode
	Parameters []*Identifier
	Body       *BlockStatement
}

func (function *FunctionLiteral) ToString() string {
	var output bytes.Buffer
	var parameters = []string{}

	for _, parameter := range function.Parameters {
		parameters = append(parameters, parameter.ToString())
	}

	output.WriteString(function.GetTokenLiteral())
	output.WriteString("(")
	output.WriteString(strings.Join(parameters, ","))
	output.WriteString(")")
	output.WriteString(function.Body.ToString())

	return output.String()
}

func (function *FunctionLiteral) GetExpressionNode() {}
