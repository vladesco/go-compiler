package object

type Null struct{}

func (nullObj *Null) Inspect() string {
	return "null"
}

func (nullObj *Null) GetObjectType() ObjectType {
	return NULL_OBJ
}
