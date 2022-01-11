package resolver

import (
	"fmt"
	"log"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/errors"
)

type FunctionType int64

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	FTYPE_NONE FunctionType = iota
	FTYPE_FUNCTION
)

type Resolver struct {
	interpreter api.IInterpreter

	// The scope stack is only used for local block scopes. Variables declared at the top
	// level in the global scope are not tracked by the resolver since they are more
	// dynamic in Lox. When resolving a variable, if we can’t find it in the stack of
	// local scopes, we assume it must be global.
	scopes *stack

	// A gate-window to detect if we are in or out of a function
	currentFunction FunctionType

	// A gate-window to detect if we are in a loop construct
	inALoop bool
}

func NewResolver(interpreter api.IInterpreter) api.IResolver {
	o := new(Resolver)
	o.interpreter = interpreter
	o.scopes = newStack()
	o.currentFunction = FTYPE_NONE
	o.inALoop = false
	return o
}

func (r *Resolver) Resolve(statements []api.IStatement) (err api.IRuntimeError) {
	return r.resolveStatements(statements)
}

// These resolve methods are similar to the evaluate() and execute() methods in
// Interpreter
func (r *Resolver) resolveStatements(statements []api.IStatement) (err api.IRuntimeError) {
	// Scan statements for a "return" statement. If found then check for any following statements.
	// If has more statements then those are unreachable.
	returnStmtFound := false

	for _, statement := range statements {
		// fmt.Println(statement)
		if statement.StmtType() == api.STMT_RETURN && !returnStmtFound {
			returnStmtFound = true
		} else if returnStmtFound {
			log.Println(fmt.Sprintf("WARNING, unreachable code: %s", statement))
		}

		err = r.resolveStatement(statement)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Resolver) resolveStatement(statement api.IStatement) (err api.IRuntimeError) {
	return statement.Accept(r)
}

func (r *Resolver) resolveExpression(expr api.IExpression) (obj interface{}, err api.IRuntimeError) {
	return expr.Accept(r)
}

func (r *Resolver) beginScope() {
	r.scopes.push(node{})
}

func (r *Resolver) endScope() {
	r.scopes.pop()
}

// Make it an error to reference a variable in its initializer.
func (r *Resolver) declare(name api.IToken) (err api.IRuntimeError) {
	if r.scopes.isEmpty() {
		return
	}

	scope := r.scopes.top()

	if _, contains := scope[name.Lexeme()]; contains {
		return errors.NewRuntimeError(name, "A variable with '"+name.Lexeme()+"' name is already in this scope.")
	}

	// Declaration adds the variable to the innermost scope so that it shadows any
	// outer one and so that we know the variable exists. We mark it as “not ready yet”
	// by binding its name to false in the scope map. The value associated with a key
	// in the scope map represents whether or not we have finished resolving that
	// variable’s initializer.
	scope[name.Lexeme()] = false

	return nil
}

func (r *Resolver) define(name api.IToken) {
	if r.scopes.isEmpty() {
		return
	}

	scope := r.scopes.top()
	// We set the variable’s value in the scope map to "true" to mark it as fully
	// initialized and available for use. It’s alive!
	scope[name.Lexeme()] = true
}

func (r *Resolver) resolveLocal(expr api.IExpression) {
	name := expr.Name().Lexeme()
	scopesSize := r.scopes.count()

	for i := scopesSize - 1; i >= 0; i-- {
		if _, ok := r.scopes.get(i)[name]; ok {
			r.interpreter.Resolve(expr, scopesSize-1-i)
			return
		}
	}
}

func (r *Resolver) resolveFunction(function api.IStatement, fType FunctionType) (err api.IRuntimeError) {
	encodingFunction := r.currentFunction

	r.currentFunction = fType

	r.beginScope()

	for _, param := range function.Parameters() {
		err = r.declare(param)
		if err != nil {
			return err
		}
		r.define(param)
	}

	err = r.resolveStatements(function.Body())
	if err != nil {
		return err
	}

	r.endScope()

	r.currentFunction = encodingFunction

	return nil
}

func (r *Resolver) resolveLoop(statement api.IStatement, inLoop bool) (err api.IRuntimeError) {
	// Capture gate
	inALoop := r.inALoop

	r.inALoop = inLoop

	_, err = r.resolveExpression(statement.Condition())
	if err != nil {
		return err
	}

	err = r.resolveStatements(statement.Body())
	if err != nil {
		return err
	}

	// Release gate
	r.inALoop = inALoop

	return nil
}
