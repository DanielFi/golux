package ast

import (
	"github.com/DanielFi/golux/internal/runtime"
)

type operator uint

// Opeators
const (
	Plus operator = iota
	Minus
	Times
	Divides
	Equals
	NotEquals
	Less
	LessEquals
	Greater
	GreaterEquals
	And
	Or
	Not
)

// Expression is
type Expression interface {
	Evaluate(*runtime.Interpreter) runtime.Value
}

// BooleanLiteral is
type BooleanLiteral bool

// Evaluate the boolean literal
func (b BooleanLiteral) Evaluate(*runtime.Interpreter) runtime.Value {
	return runtime.Boolean(bool(b))
}

// IntegerLiteral is
type IntegerLiteral int32

// Evaluate the integer literal
func (i IntegerLiteral) Evaluate(*runtime.Interpreter) runtime.Value {
	return runtime.Integer(int32(i))
}

// StringLiteral is
type StringLiteral string

// Evaluate the string literal
func (s StringLiteral) Evaluate(*runtime.Interpreter) runtime.Value {
	return runtime.String(string(s))
}

// Identifier is
type Identifier string

// Evaluate evaluates an identifier
func (i Identifier) Evaluate(interpreter *runtime.Interpreter) runtime.Value {
	return interpreter.GetVariable(string(i))
}

// BlockExpression is a collection of expressions
type BlockExpression struct {
	Expressions []Expression
}

// AppendExpression appends an expression to a block expression
func AppendExpression(b BlockExpression, e Expression) BlockExpression {
	b.Expressions = append(b.Expressions, e)
	return b
}

// Evaluate evaluates a block expression
func (e BlockExpression) Evaluate(interpreter *runtime.Interpreter) runtime.Value {
	var result runtime.Value
	result = runtime.Nil{}

	for _, expression := range e.Expressions {
		result = expression.Evaluate(interpreter)
	}

	return result
}

// BinaryOperationExpression is
type BinaryOperationExpression struct {
	Operator operator
	LHS      Expression
	RHS      Expression
}

// Evaluate the binary operation expression
func (e BinaryOperationExpression) Evaluate(interpreter *runtime.Interpreter) runtime.Value {
	lhs := e.LHS.Evaluate(interpreter)

	switch lhs.(type) {
	case runtime.Boolean:
		switch e.Operator {
		case And:
			if bool(lhs.(runtime.Boolean)) {
				return lhs
			}
			rhs := e.RHS.Evaluate(interpreter)
			_, ok := rhs.(runtime.Boolean)
			if ok {
				return rhs
			}
			panic("INVALID TYPE")
		case Or:
			if bool(lhs.(runtime.Boolean)) {
				return lhs
			}
			rhs := e.RHS.Evaluate(interpreter)
			_, ok := rhs.(runtime.Boolean)
			if ok {
				return rhs
			}
			panic("INVALID TYPE")
		default:
			panic("INVALID OPERATOR")
		}
	case runtime.Integer:
		rhs, ok := e.RHS.Evaluate(interpreter).(runtime.Integer)
		if !ok {
			panic("INVALID TYPE")
		}
		var result runtime.Value
		switch e.Operator {
		case Plus:
			result = runtime.Integer(int32(lhs.(runtime.Integer)) + int32(rhs))
		case Minus:
			result = runtime.Integer(int32(lhs.(runtime.Integer)) - int32(rhs))
		case Times:
			result = runtime.Integer(int32(lhs.(runtime.Integer)) * int32(rhs))
		case Divides:
			result = runtime.Integer(int32(lhs.(runtime.Integer)) / int32(rhs))
		case Equals:
			result = runtime.Boolean(int32(lhs.(runtime.Integer)) == int32(rhs))
		case NotEquals:
			result = runtime.Boolean(int32(lhs.(runtime.Integer)) != int32(rhs))
		case Less:
			result = runtime.Boolean(int32(lhs.(runtime.Integer)) < int32(rhs))
		case LessEquals:
			result = runtime.Boolean(int32(lhs.(runtime.Integer)) <= int32(rhs))
		case Greater:
			result = runtime.Boolean(int32(lhs.(runtime.Integer)) > int32(rhs))
		case GreaterEquals:
			result = runtime.Boolean(int32(lhs.(runtime.Integer)) >= int32(rhs))
		default:
			panic("INVALID OPERATOR")
		}
		return result
	case runtime.String:
		rhs, ok := e.RHS.Evaluate(interpreter).(runtime.String)
		if !ok {
			panic("INVALID TYPE")
		}
		switch e.Operator {
		case Plus:
			return runtime.String(string((lhs).(runtime.String)) + string(rhs))
		default:
			panic("INVALID OPERATOR")
		}
	default:
		panic("INVALID TYPE")
	}
}

// VariableDeclaration is
type VariableDeclaration struct {
	LHS Identifier
	RHS Expression
}

// Evaluate declares a new variable
func (d VariableDeclaration) Evaluate(interpreter *runtime.Interpreter) runtime.Value {
	interpreter.DeclareVariable(string(d.LHS))

	var value runtime.Value
	value = runtime.Nil{}
	if d.RHS != nil {
		value = d.RHS.Evaluate(interpreter)
	}
	interpreter.SetVariable(string(d.LHS), value)

	return value
}

// VariableAssignment is
type VariableAssignment struct {
	LHS Identifier
	RHS Expression
}

// Evaluate assign a value to a variable
func (a VariableAssignment) Evaluate(interpreter *runtime.Interpreter) runtime.Value {
	var value runtime.Value

	value = a.RHS.Evaluate(interpreter)
	interpreter.SetVariable(string(a.LHS), value)

	return value
}

// FunctionDeclaration is
type FunctionDeclaration struct {
	Name      Identifier
	Arguments []Identifier
	Body      Expression
}

// Evaluate evaluates a function declaration
func (d FunctionDeclaration) Evaluate(interpreter *runtime.Interpreter) runtime.Value {
	arguments := make([]string, len(d.Arguments))
	for i, a := range d.Arguments {
		arguments[i] = string(a)
	}

	function := runtime.Function{
		Closure:   interpreter.GetScope(),
		Arguments: arguments,
		Call: func(i *runtime.Interpreter) runtime.Value {
			return d.Body.Evaluate(i)
		},
	}

	interpreter.DeclareVariable(string(d.Name))
	interpreter.SetVariable(string(d.Name), function)

	return function
}

// CallExpression is
type CallExpression struct {
	Function   Identifier
	Parameters []Expression
}

// Evaluate evaluates a call expression
func (e CallExpression) Evaluate(interpreter *runtime.Interpreter) runtime.Value {
	parameters := make([]runtime.Value, len(e.Parameters))
	for i, p := range e.Parameters {
		parameters[i] = p.Evaluate(interpreter)
	}

	return interpreter.CallFunction(string(e.Function), parameters)
}
