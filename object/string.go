package object

type String struct {
	Value string
}

func (stringObj *String) Inspect() string {
	return stringObj.Value
}

func (stringObj *String) GetObjectType() ObjectType {
	return STRING_OBJ
}
