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

type ExpressionStatement struct {
	BaseNode
	Expression Expression
}

func (statement *ExpressionStatement) ToString() string {
	if statement.Expression != nil {
		return statement.Expression.ToString()
	}

	return ""
}

func (statement *ExpressionStatement) GetStatementNode() {}

type Identifier struct {
	BaseNode
	Value string
}

func (identifier *Identifier) ToString() string {
	return identifier.Value
}

func (identifier *Identifier) GetExpressionNode() {}

type IntegerLiteral struct {
	BaseNode
	Value int64
}

func (literal *IntegerLiteral) ToString() string {
	return literal.Token.Literal
}

func (literal *IntegerLiteral) GetExpressionNode() {}

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

type Program struct {
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
