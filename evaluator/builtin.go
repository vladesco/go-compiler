package evaluator

import (
	"compiler/object"
	"fmt"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: builtLen,
	},
}

var builtLen object.BuiltinFn = func(args ...object.Object) object.Object {
	if countOfArguments := len(args); countOfArguments != 1 {
		return newError(fmt.Sprintf("wrong number of arguments: want 1, but get %d", countOfArguments))
	}

	if stringObj, ok := args[0].(*object.String); ok {
		return &object.Integer{Value: int64(len(stringObj.Value))}

	} else {
		return newError(fmt.Sprintf("len supports only string but get %s", args[0].GetObjectType()))
	}
}
