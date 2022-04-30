package object

type Error struct {
	Message string
}

func (errorObj *Error) Inspect() string {
	return "Error: " + errorObj.Message
}

func (errorObj *Error) GetObjectType() ObjectType {
	return ERROR_OBJ
}
