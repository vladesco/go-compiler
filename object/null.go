package object

type Null struct{}

func (null *Null) Inspect() string {
	return "null"
}

func (null *Null) GetObjectType() ObjectType {
	return NULL_OBJ
}
