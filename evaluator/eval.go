package evaluator

import (
	"compiler/ast"
	"compiler/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalProgram(node)

	case *ast.BlockStatement:
		return evalStatements(node)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return convertBoolToBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		right := Eval(node.Right)
		left := Eval(node.Left)
		return evalInfixExpression(node.Operator, left, right)

	case *ast.IfExpression:
		return evalIfExpression(node)

	case *ast.ReturnStatement:
		value := Eval(node.Value)
		return &object.ReturnValue{Value: value}
	}

	return nil
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement)

		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}

func evalStatements(blockStatement *ast.BlockStatement) object.Object {
	var result object.Object

	for _, statement := range blockStatement.Statements {
		result = Eval(statement)

		if result != nil && result.GetObjectType() == object.RETURN_VALUE_OBJ {
			return result
		}
	}

	return result
}

func convertBoolToBooleanObject(argument bool) *object.Boolean {
	if argument {
		return TRUE
	}

	return FALSE
}

func evalPrefixExpression(operator string, argument object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(argument)

	case "-":
		return evalMinusPrefixOperatorExpression(argument)

	default:
		return NULL
	}
}

func evalBangOperatorExpression(argument object.Object) object.Object {
	switch argument.GetObjectType() {
	case object.BOOLEAN_OBJ:
		return convertBoolToBooleanObject(!argument.(*object.Boolean).Value)

	case object.INTEGER_OBJ:
		booleanValue := argument.(*object.Integer).Value != 0
		return convertBoolToBooleanObject(!booleanValue)

	case object.NULL_OBJ:
		return FALSE

	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(argument object.Object) object.Object {
	if argument.GetObjectType() != object.INTEGER_OBJ {
		return NULL
	}

	value := argument.(*object.Integer).Value

	return &object.Integer{Value: -value}
}

func evalInfixExpression(operator string, firstArgument, secondArgument object.Object) object.Object {
	switch {

	case firstArgument.GetObjectType() == object.INTEGER_OBJ && secondArgument.GetObjectType() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, firstArgument, secondArgument)

	case operator == "==":
		return convertBoolToBooleanObject(firstArgument == secondArgument)

	case operator == "!=":
		return convertBoolToBooleanObject(firstArgument != secondArgument)

	default:
		return NULL
	}
}

func evalIntegerInfixExpression(operator string, firstArgument, secondArgument object.Object) object.Object {
	firstValue := firstArgument.(*object.Integer).Value
	secondValue := secondArgument.(*object.Integer).Value

	switch operator {
	case "/":
		return &object.Integer{Value: firstValue / secondValue}

	case "*":
		return &object.Integer{Value: firstValue * secondValue}

	case "+":
		return &object.Integer{Value: firstValue + secondValue}

	case "-":
		return &object.Integer{Value: firstValue - secondValue}

	case ">":
		return convertBoolToBooleanObject(firstValue > secondValue)

	case "<":
		return convertBoolToBooleanObject(firstValue < secondValue)

	case "==":
		return convertBoolToBooleanObject(firstValue == secondValue)

	case "!=":
		return convertBoolToBooleanObject(firstValue != secondValue)

	default:
		return NULL
	}

}

func evalIfExpression(argument *ast.IfExpression) object.Object {
	condition := Eval(argument.Condition)

	if isTruthy(condition) {
		return Eval(argument.Consequence)
	}

	if argument.Alternative != nil {
		return Eval(argument.Alternative)
	}

	return NULL
}

func isTruthy(argument object.Object) bool {
	switch argument.GetObjectType() {
	case object.BOOLEAN_OBJ:
		return argument.(*object.Boolean).Value

	case object.INTEGER_OBJ:
		return argument.(*object.Integer).Value != 0

	default:
		return false
	}
}
