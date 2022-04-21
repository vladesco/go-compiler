package object

type ReturnValue struct {
	Value Object
}

func (returnValue *ReturnValue) Inspect() string {
	return returnValue.Value.Inspect()
}

func (returnValue *ReturnValue) GetObjectType() ObjectType {
	return RETURN_VALUE_OBJ
}
