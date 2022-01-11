package resolver

import (
	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/errors"
)

// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
// IVisitorExpression interface
// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
func (r *Resolver) VisitAssignExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	// First, we resolve the expression for the assigned value in case it also containsreferences to other variables.
	_, err = r.resolveExpression(exprV.Expression()) // i.e. exprV.Value()
	if err != nil {
		return nil, err
	}

	// Then we use our existing resolveLocal()
	// method to resolve the variable that’s being assigned to.
	r.resolveLocal(exprV)

	return nil, nil
}

func (r *Resolver) VisitBinaryExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	_, err = r.resolveExpression(exprV.Left())
	if err != nil {
		return nil, err
	}

	_, err = r.resolveExpression(exprV.Right())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *Resolver) VisitCallExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	_, err = r.resolveExpression(exprV.Callee())
	if err != nil {
		return nil, err
	}

	for _, argument := range exprV.Arguments() {
		_, err = r.resolveExpression(argument)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (r *Resolver) VisitGroupingExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	_, err = r.resolveExpression(exprV.Expression())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *Resolver) VisitLiteralExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	return nil, nil
}

func (r *Resolver) VisitLogicalExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	_, err = r.resolveExpression(exprV.Left())
	if err != nil {
		return nil, err
	}

	_, err = r.resolveExpression(exprV.Right())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *Resolver) VisitUnaryExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	_, err = r.resolveExpression(exprV.Right())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *Resolver) VisitVariableExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	if !r.scopes.isEmpty() {
		name := exprV.Name()
		if alive, ok := r.scopes.top()[name.Lexeme()]; ok {
			// If the variable exists in the current scope but its
			// value is "false" , that means we have declared it but
			// not yet defined it.
			if !alive {
				return nil, errors.NewRuntimeError(name, "Can't read '"+name.Lexeme()+"' local variable in its own initializer.")
			}
		}
		// else {
		// When resolving a variable, if we can’t find it in the stack of
		// local scopes, we assume it must be global.
		// See resolver.go "struct"
		// Hence, we don't return an err
		// return nil, errors.NewRuntimeError(name, "Can't find '"+name.Lexeme()+"' local variable in current scope.")
		// }
	}

	r.resolveLocal(exprV)

	return nil, nil
}
