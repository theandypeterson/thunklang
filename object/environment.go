package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	env.fnCache = outer.fnCache
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	f := make(map[string]Object)
	return &Environment{store: s, outer: nil, fnCache: f}
}

type Environment struct {
	store   map[string]Object
	outer   *Environment
	fnCache map[string]Object
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (e *Environment) GetFnCache(key string) (Object, bool) {
	obj, ok := e.fnCache[key]
	return obj, ok
}

func (e *Environment) SetFnCache(key string, val Object) Object {
	e.fnCache[key] = val
	return val
}
