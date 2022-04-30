package object

type ReturnValue struct {
	Value Object
}

func (returnValueObj *ReturnValue) Inspect() string {
	return returnValueObj.Value.Inspect()
}

func (returnValueObj *ReturnValue) GetObjectType() ObjectType {
	return RETURN_VALUE_OBJ
}
