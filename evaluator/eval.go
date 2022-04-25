package evaluator

import (
	"compiler/ast"
	"compiler/object"
	"fmt"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, environment *object.Environment) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalProgram(node, environment)

	case *ast.BlockStatement:
		return evalBlockStatements(node, environment.Extend())

	case *ast.ExpressionStatement:
		return Eval(node.Expression, environment)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return convertBoolToBooleanObject(node.Value)

	case *ast.LetStatement:
		value := Eval(node.Value, environment)

		if isError(value) {
			return value
		}

		environment.Set(node.Name.Value, value)

	case *ast.Identifier:
		value, exist := environment.Get(node.Value)

		if exist {
			return value
		}

		return newError(fmt.Sprintf("variable doesn`t exist %s", node.Value))

	case *ast.PrefixExpression:
		right := Eval(node.Right, environment)

		if isError(right) {
			return right
		}

		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		right := Eval(node.Right, environment)

		if isError(right) {
			return right
		}

		left := Eval(node.Left, environment)

		if isError(left) {
			return left
		}

		return evalInfixExpression(node.Operator, left, right)

	case *ast.IfExpression:
		return evalIfExpression(node, environment)

	case *ast.ReturnStatement:
		value := Eval(node.Value, environment)

		if isError(value) {
			return value
		}

		return &object.ReturnValue{Value: value}
	}

	return NULL
}

func evalProgram(program *ast.Program, environment *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, environment)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value

		case *object.Error:
			return result
		}

	}

	return result
}

func evalBlockStatements(blockStatement *ast.BlockStatement, environment *object.Environment) object.Object {
	var result object.Object

	for _, statement := range blockStatement.Statements {
		result = Eval(statement, environment)

		if result.GetObjectType() == object.RETURN_VALUE_OBJ || result.GetObjectType() == object.ERROR_OBJ {
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
		return newError(fmt.Sprintf("unknown prefix operator %s", operator))
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
		return newError(fmt.Sprintf("%s should be a number", argument.Inspect()))
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
		return newError(fmt.Sprintf("unknown infix operator %s or operands %s %s", operator, firstArgument.Inspect(), secondArgument.Inspect()))
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
		return newError(fmt.Sprintf("unknown infix operator %s", operator))

	}

}

func evalIfExpression(argument *ast.IfExpression, environment *object.Environment) object.Object {
	condition := Eval(argument.Condition, environment)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(argument.Consequence, environment)
	}

	if argument.Alternative != nil {
		return Eval(argument.Alternative, environment)
	}

	return NULL
}

func newError(errorMessage string) object.Object {
	return &object.Error{Message: errorMessage}
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

func isError(error object.Object) bool {
	if error != nil {
		return error.GetObjectType() == object.ERROR_OBJ
	}
	return false
}
