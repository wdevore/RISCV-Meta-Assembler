package interpreter

import (
	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/errors"
)

type Environment struct {
	enclosing api.IEnvironment
	values    map[string]interface{}
}

// for the global scope’s environment
func NewEnvironment() api.IEnvironment {
	o := new(Environment)
	o.values = make(map[string]interface{})
	o.enclosing = nil
	return o
}

// local scope nested inside the given outer one
func NewEnvironmentEnclosing(enclosing api.IEnvironment) api.IEnvironment {
	o := new(Environment)
	o.values = make(map[string]interface{})
	o.enclosing = enclosing
	return o
}

func (e *Environment) Define(name string, obj interface{}) (err api.IRuntimeError) {
	_, ok := e.values[name]
	if !ok {
		e.values[name] = obj
		return nil
	}

	return errors.NewRuntimeError(nil, "Variable '"+name+"' already defined.")
}

func (e *Environment) Enclosing() api.IEnvironment {
	return e.enclosing
}

func (e *Environment) Values() map[string]interface{} {
	return e.values
}

func (e *Environment) Get(name api.IToken) (value interface{}, err api.IRuntimeError) {
	value, ok := e.values[name.Lexeme()]
	if ok {
		return value, nil
	}

	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}

	return nil, errors.NewRuntimeError(name, "Undefined variable '"+name.Lexeme()+"'.")
}

func (e *Environment) GetAt(distance int, name api.IToken) (obj interface{}, err api.IRuntimeError) {
	ancestor := e.ancestor(distance)
	if value, ok := ancestor.Values()[name.Lexeme()]; ok {
		return value, nil
	}

	// Should never actually reach this unless there is a coding
	// error or resolver error.
	return nil, errors.NewRuntimeError(name, "Undefined variable '"+name.Lexeme()+"'.")
}

func (e *Environment) ancestor(distance int) api.IEnvironment {
	environment := e

	for i := 0; i < distance; i++ {
		environment = environment.enclosing.(*Environment)
	}

	return environment
}

func (e *Environment) Assign(name api.IToken, value interface{}) (err api.IRuntimeError) {
	_, ok := e.values[name.Lexeme()]
	if ok {
		e.values[name.Lexeme()] = value
		return nil
	}

	if e.enclosing != nil {
		e.enclosing.Assign(name, value)
		return nil
	}

	return errors.NewRuntimeError(name, "Undefined variable '"+name.Lexeme()+"'.")
}

func (e *Environment) AssignAt(distance int, name api.IToken, value interface{}) (err api.IRuntimeError) {
	e.ancestor(distance).Values()[name.Lexeme()] = value
	return nil
}
