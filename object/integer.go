package object

import "fmt"

type Integer struct {
	Value int64
}

func (integer *Integer) Inspect() string {
	return fmt.Sprintf("%d", integer.Value)
}

func (integer *Integer) GetObjectType() ObjectType {
	return INTEGER_OBJ
}
