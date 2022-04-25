package object

type Error struct {
	Message string
}

func (error *Error) Inspect() string {
	return "Error: " + error.Message
}

func (error *Error) GetObjectType() ObjectType {
	return ERROR_OBJ
}
