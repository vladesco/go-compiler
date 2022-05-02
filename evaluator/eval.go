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

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.Boolean:
		return convertBoolToBooleanObject(node.Value)

	case *ast.FunctionLiteral:
		return &object.Function{
			Parameters:  node.Parameters,
			Body:        node.Body,
			Environment: environment,
		}

	case *ast.CallExpression:
		fn := Eval(node.Function, environment)

		if isError(fn) {
			return fn
		}

		switch fn := fn.(type) {
		case *object.Function:
			{
				extendedEnvironment, environmentError := createFunctionEnvironment(node.Arguments, fn.Parameters, environment)
				if environmentError != nil {
					return environmentError
				}

				return unwrapReturnValue(evalBlockStatements(fn.Body, extendedEnvironment))
			}

		case *object.Builtin:
			{
				return fn.Fn(evalArguments(node.Arguments, environment)...)
			}

		default:
			return newError(fmt.Sprintf("not a function: %s", fn.GetObjectType()))
		}

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

		if builtIn, exist := builtins[node.Value]; exist {
			return builtIn
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

func evalArguments(arguments []ast.Expression, environment *object.Environment) []object.Object {
	var result []object.Object

	for _, argument := range arguments {
		resultArgument := Eval(argument, environment)

		if isError(resultArgument) {
			return []object.Object{resultArgument}
		}

		result = append(result, resultArgument)
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
	firstArgumentType := firstArgument.GetObjectType()
	secondArgumentType := secondArgument.GetObjectType()

	switch {

	case firstArgumentType == object.INTEGER_OBJ && secondArgumentType == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, firstArgument, secondArgument)

	case firstArgumentType == object.STRING_OBJ && secondArgumentType == object.STRING_OBJ:
		return evalStringInfixExpression(operator, firstArgument, secondArgument)

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

func evalStringInfixExpression(operator string, firstArgument, secondArgument object.Object) object.Object {
	firstValue := firstArgument.(*object.String).Value
	secondValue := secondArgument.(*object.String).Value

	if operator == "+" {
		return &object.String{Value: firstValue + secondValue}
	}

	return newError(fmt.Sprintf("unknown infix operator %s", operator))
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

func createFunctionEnvironment(arguments []ast.Expression, parameters []*ast.Identifier, environment *object.Environment) (*object.Environment, object.Object) {
	extendedEnvironment := environment.Extend()
	evaledArguments := evalArguments(arguments, extendedEnvironment)

	if len(evaledArguments) == 1 && isError(evaledArguments[0]) {
		return extendedEnvironment, evaledArguments[0]
	}

	for index, argument := range evaledArguments {
		extendedEnvironment.Set(parameters[index].Value, argument)
	}

	return extendedEnvironment, nil
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

func unwrapReturnValue(returnObject object.Object) object.Object {
	if returnValue, ok := returnObject.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return returnObject
}
