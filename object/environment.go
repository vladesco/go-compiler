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

func (environment *Environment) Get(name string) (Object, bool) {
	value, exist := environment.store[name]

	if !exist && environment.outer != nil {
		return environment.outer.Get(name)
	}

	return value, exist
}

func (environment *Environment) Set(name string, value Object) Object {
	environment.store[name] = value

	return value
}

func (environment *Environment) Extend() *Environment {
	extendedEnvironemt := NewEnvironment()
	extendedEnvironemt.outer = environment

	return extendedEnvironemt
}
