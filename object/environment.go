package object

type Environment struct {
	store map[string]Object
}

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
	}
}

func (environment *Environment) Get(name string) (Object, bool) {
	value, exist := environment.store[name]

	return value, exist
}

func (environment *Environment) Set(name string, value Object) Object {
	environment.store[name] = value

	return value
}
