package ast

import (
	"compiler/token"
)

type Node interface {
	GetTokenLiteral() string
}

type Statement interface {
	Node
	GetStatementNode()
}

type Expression interface {
	Node
	GetExpressionNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (statement *LetStatement) GetStatementNode()       {}
func (statement *LetStatement) GetTokenLiteral() string { return statement.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (identifier *Identifier) GetExpressionNode()      {}
func (identifier *Identifier) GetTokenLiteral() string { return identifier.Token.Literal }

type Program struct {
	Statements []Statement
}

func (program *Program) GetNextTokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].GetTokenLiteral()
	}

	return ""
}
