package ast

import (
	"bytes"
	"compiler/token"
)

type Node interface {
	GetTokenLiteral() string
	ToString() string
}

type BaseNode struct {
	Token token.Token
}

func (node *BaseNode) GetTokenLiteral() string { return node.Token.Literal }
func (node *BaseNode) ToString() string        { return node.GetTokenLiteral() }

type Statement interface {
	Node
	GetStatementNode()
}

type Expression interface {
	Node
	GetExpressionNode()
}

type Program struct {
	Node
	Statements []Statement
}

func (program *Program) GetNextTokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].GetTokenLiteral()
	}

	return ""
}

func (program *Program) ToString() string {
	var output bytes.Buffer

	for _, statement := range program.Statements {
		output.WriteString(statement.ToString())
	}

	return output.String()
}
