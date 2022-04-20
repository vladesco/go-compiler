package ast

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
