package object

import "fmt"

type Boolean struct {
	Value bool
}

func (booleanObj *Boolean) Inspect() string {
	return fmt.Sprintf("%t", booleanObj.Value)
}

func (booleanObj *Boolean) GetObjectType() ObjectType {
	return BOOLEAN_OBJ
}
