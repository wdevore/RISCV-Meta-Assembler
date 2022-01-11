package resolver

import (
	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/errors"
)

// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
// IVisitorStatement implementations
// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
func (r *Resolver) VisitExpressionStatement(statement api.IStatement) (err api.IRuntimeError) {
	r.resolveExpression(statement.Expression())
	return nil
}

func (r *Resolver) VisitBlockStatement(statement api.IStatement) (err api.IRuntimeError) {
	r.beginScope()

	err = r.resolveStatements(statement.Statements())
	if err != nil {
		return err
	}

	r.endScope()

	return nil
}

func (r *Resolver) VisitFunctionStatement(statement api.IStatement) (err api.IRuntimeError) {
	err = r.declare(statement.Name())
	if err != nil {
		return err
	}
	r.define(statement.Name())

	err = r.resolveFunction(statement, FTYPE_FUNCTION)
	if err != nil {
		return err
	}

	return nil
}

func (r *Resolver) VisitIfStatement(statement api.IStatement) (err api.IRuntimeError) {
	_, err = r.resolveExpression(statement.Condition())
	if err != nil {
		return err
	}

	err = r.resolveStatement(statement.ThenBranch())
	if err != nil {
		return err
	}

	if statement.ElseBranch() != nil {
		err = r.resolveStatement(statement.ElseBranch())
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Resolver) VisitInterruptStatement(statement api.IStatement) (err api.IRuntimeError) {
	if !r.inALoop && statement.Type() == api.INTERRUPT_BREAK {
		return errors.NewRuntimeError(statement.Name(), "Can't break outside of a loop.")
	}

	return nil
}

func (r *Resolver) VisitPrintStatement(statement api.IStatement) (err api.IRuntimeError) {
	_, err = r.resolveExpression(statement.Expression())
	if err != nil {
		return err
	}
	return nil
}

func (r *Resolver) VisitReturnStatement(statement api.IStatement) (err api.IRuntimeError) {
	if r.currentFunction == FTYPE_NONE {
		return errors.NewRuntimeError(statement.Keyword(), "Can't return from top-level code.")
	}

	if statement.Value() != nil {
		_, err = r.resolveExpression(statement.Value())
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Resolver) VisitVariableStatement(statement api.IStatement) (err api.IRuntimeError) {
	err = r.declare(statement.Name())
	if err != nil {
		return err
	}

	if statement.Initializer() != nil {
		r.resolveExpression(statement.Initializer())
	}

	r.define(statement.Name())

	return nil
}

func (r *Resolver) VisitWhileStatement(statement api.IStatement) (err api.IRuntimeError) {
	err = r.resolveLoop(statement, true)
	if err != nil {
		return err
	}

	return nil
}
