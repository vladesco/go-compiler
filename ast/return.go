package ast

import "bytes"

type ReturnStatement struct {
	BaseNode
	Value Expression
}

func (statement *ReturnStatement) ToString() string {
	var output bytes.Buffer

	output.WriteString(statement.GetTokenLiteral())
	output.WriteString(" ")

	if statement.Value != nil {
		output.WriteString(statement.Value.ToString())
	}

	output.WriteString(";")

	return output.String()
}

func (statement *ReturnStatement) GetStatementNode() {}
