package ast

import "bytes"

type LetStatement struct {
	BaseNode
	Name  *Identifier
	Value Expression
}

func (statement *LetStatement) ToString() string {
	var output bytes.Buffer

	output.WriteString(statement.GetTokenLiteral())
	output.WriteString(" ")
	output.WriteString(statement.Name.ToString())
	output.WriteString(" = ")

	if statement.Value != nil {
		output.WriteString(statement.Value.ToString())
	}

	output.WriteString(";")

	return output.String()
}

func (statement *LetStatement) GetStatementNode() {}
