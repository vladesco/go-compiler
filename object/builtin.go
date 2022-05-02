package object

type BuiltinFn func(arguments ...Object) Object

type Builtin struct {
	Fn BuiltinFn
}

func (builtinObj *Builtin) Inspect() string {
	return "builtin function"
}

func (builtinObj *Builtin) GetObjectType() ObjectType {
	return BUILTIN_OBJ
}
