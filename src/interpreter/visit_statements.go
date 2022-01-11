package interpreter

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/errors"
)

// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
// IVisitorStatement implementations
// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
func (i *Interpreter) VisitExpressionStatement(statement api.IStatement) (err api.IRuntimeError) {
	// Simply decend
	_, err = i.evaluate(statement.Expression())
	return err
}

func (i *Interpreter) VisitPrintStatement(statement api.IStatement) (err api.IRuntimeError) {
	value, err := i.evaluate(statement.Expression())

	if err == nil {
		fmt.Println(value)
	}

	return err
}

func (i *Interpreter) VisitVariableStatement(statement api.IStatement) (err api.IRuntimeError) {
	var value interface{} = nil

	if statement.Initializer() != nil {
		value, err = i.evaluate(statement.Initializer())
		if err != nil {
			return err
		}
	}

	return i.environment.Define(statement.Name().Lexeme(), value)
}

func (i *Interpreter) VisitBlockStatement(statement api.IStatement) (err api.IRuntimeError) {
	childEnv := NewEnvironmentEnclosing(i.environment)
	return i.ExecuteBlock(statement.Statements(), childEnv)
}

func (i *Interpreter) VisitIfStatement(statement api.IStatement) (err api.IRuntimeError) {
	value, err := i.evaluate(statement.Condition())
	if err != nil {
		return err
	}

	if i.isTruthy(value) {
		err = i.execute(statement.ThenBranch())
		if err != nil {
			return err
		}
	} else if statement.ElseBranch() != nil {
		err = i.execute(statement.ElseBranch())
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Interpreter) VisitWhileStatement(statement api.IStatement) (err api.IRuntimeError) {
	value, err := i.evaluate(statement.Condition())
	if err != nil {
		return err
	}

	for i.isTruthy(value) {
		err = i.execute(statement.Body()[0])
		if err != nil {
			// "break", "continue" interrupt statements
			if err.Interrupt() == api.INTERRUPT_BREAK {
				// fmt.Println("VisitWhileStatement: breaking")
				break
			} else if err.Interrupt() == api.INTERRUPT_CONTINUE {
				// fmt.Println("VisitWhileStatement: continuing")
				// Fall through
				// (i.e.) continue
			} else {
				// An actual error
				return err
			}
		}

		value, err = i.evaluate(statement.Condition())
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Interpreter) VisitInterruptStatement(statement api.IStatement) (err api.IRuntimeError) {
	msg := fmt.Sprintf("'%s' interrupt type.", statement.Type())
	err = errors.NewRuntimeError(nil, msg)

	err.SetInterrupt(statement.Type())

	return err
}

func (i *Interpreter) VisitFunctionStatement(statement api.IStatement) (err api.IRuntimeError) {
	// This is similar to how we interpret other literal expressions. We take a function
	// syntax node--a compile time representation of the function—-and convert it to
	// its runtime representation. Here, that’s a LoxFunction that wraps the syntax
	// node.

	// The environment that is active when the function is declared not when it’s called
	// It represents the lexical scope surrounding the
	// function declaration.
	function := NewFunctionCallable(statement, i.environment)
	return i.environment.Define(statement.Name().Lexeme(), function)
}

func (i *Interpreter) VisitReturnStatement(statement api.IStatement) (err api.IRuntimeError) {
	var value interface{}

	if statement.Value() != nil {
		value, err = i.evaluate(statement.Value())
		if err != nil {
			return err
		}
	}

	msg := fmt.Sprintf("'%s' interrupt type.", statement.Type())
	err = errors.NewRuntimeError(nil, msg)

	err.SetValue(value)
	err.SetInterrupt(statement.Type())

	return err
}
