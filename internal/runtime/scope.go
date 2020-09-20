package runtime

// Scope is
type Scope struct {
	parent    *Scope
	variables map[string]interface{}
}

// newScope creates a new scope with a given scope as its parent
func newScope(parent *Scope) *Scope {
	return &Scope{parent, make(map[string]interface{})}
}

func (s Scope) declareVariable(name string) {
	s.variables[name] = &Nil{}
}

func (s Scope) setVariable(name string, value Value) {
	if _, ok := s.variables[name]; ok || s.parent == nil {
		s.variables[name] = value
	} else {
		s.parent.setVariable(name, value)
	}
}

func (s Scope) getVariable(name string) Value {
	if value, ok := s.variables[name]; ok {
		return value.(Value)
	}
	if s.parent != nil {
		return s.parent.getVariable(name)
	}
	panic("UNDEFINED VARIABLE")
}
