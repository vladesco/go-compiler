package object

import "fmt"

type Boolean struct {
	Value bool
}

func (boolean *Boolean) Inspect() string {
	return fmt.Sprintf("%t", boolean.Value)
}

func (boolean *Boolean) GetObjectType() ObjectType {
	return BOOLEAN_OBJ
}
