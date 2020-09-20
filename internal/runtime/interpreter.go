package runtime

// Interpreter interprets lux code
type Interpreter struct {
	scope *Scope
}

// NewInterpreter return a new interpreter
func NewInterpreter() *Interpreter {
	return &Interpreter{newScope(nil)}
}

// GetScope returns the interpreter's current scope
func (i *Interpreter) GetScope() *Scope {
	return i.scope
}

func (i *Interpreter) enterScope(scope *Scope) {
	scope.parent = i.scope
	i.scope = scope
}

func (i *Interpreter) exitScope() {
	i.scope = i.scope.parent
}

// GetVariable returns a variable's value
func (i *Interpreter) GetVariable(name string) Value {
	return i.scope.getVariable(name)
}

// SetVariable sets a variable's value
func (i *Interpreter) SetVariable(name string, value Value) {
	i.scope.setVariable(name, value)
}

// DeclareVariable declares a variable in the current scope
func (i *Interpreter) DeclareVariable(name string) {
	i.scope.declareVariable(name)
}

// CallFunction calls a function by name with given parameters
func (i *Interpreter) CallFunction(name string, parameters []Value) Value {
	value := i.GetVariable(name)

	if function, ok := (value).(Function); ok {
		originalScope := i.scope
		i.scope = function.Closure

		newScope := *newScope(nil)
		for idx, parameter := range parameters {
			newScope.setVariable(function.Arguments[idx], parameter)
		}
		i.enterScope(&newScope)

		result := function.Call(i)

		i.exitScope()
		i.scope = originalScope

		return result
	}

	panic("ATTEMPTED TO CALL A NON FUNCTION")
}
