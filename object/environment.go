package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

// In order to support closures and lexical scoping, we need to
// ensure we extend all outer environments with new bindings
// (rather than overwriting, since an `x` inside a function body
// shouldn't overwrite an `x` already bound in an outer scope).
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

func (e *Environment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}
