package runtime

// Value is a concrete runtime value
type Value interface {
}

// Boolean is a runtime boolean value
type Boolean bool

// Integer is a runtime integer value
type Integer int32

// String is a runtime string value
type String string

// Nil is a runtime nil value
type Nil struct{}

// Function is a runtime function value
type Function struct {
	Closure   *Scope
	Arguments []string
	Call      func(*Interpreter) Value
}
