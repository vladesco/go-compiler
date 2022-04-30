package object

import "fmt"

type Integer struct {
	Value int64
}

func (integerObj *Integer) Inspect() string {
	return fmt.Sprintf("%d", integerObj.Value)
}

func (integer *Integer) GetObjectType() ObjectType {
	return INTEGER_OBJ
}
