package ast

import "bytes"

type BlockStatement struct {
	BaseNode
	Statements []Statement
}

func (statement *BlockStatement) ToString() string {
	var output bytes.Buffer

	for _, subStatement := range statement.Statements {
		output.WriteString(subStatement.ToString())
	}

	return output.String()
}
