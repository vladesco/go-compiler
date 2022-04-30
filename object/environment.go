package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
	}
}

func (environmentObj *Environment) Get(name string) (Object, bool) {
	value, exist := environmentObj.store[name]

	if !exist && environmentObj.outer != nil {
		return environmentObj.outer.Get(name)
	}

	return value, exist
}

func (environmentObj *Environment) Set(name string, value Object) Object {
	environmentObj.store[name] = value

	return value
}

func (environmentObj *Environment) Extend() *Environment {
	extendedEnvironemtObj := NewEnvironment()
	extendedEnvironemtObj.outer = environmentObj

	return extendedEnvironemtObj
}
