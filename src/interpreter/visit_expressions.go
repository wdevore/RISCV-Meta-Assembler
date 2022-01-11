package interpreter

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/errors"
	"github.com/wdevore/RISCV-Meta-Assembler/src/scanner/literals"
)

// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
// IVisitorExpression interface
// -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
func (i *Interpreter) VisitLiteralExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	return exprV.Value(), nil
}

func (i *Interpreter) VisitLogicalExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	left, err := i.evaluate(exprV.Left())
	if err != nil {
		return nil, err
	}

	if exprV.Operator().Type() == api.OR {
		if i.isTruthy(left) {
			return left, nil
		}
	} else {
		if !i.isTruthy(left) {
			return left, nil
		}
	}

	return i.evaluate(exprV.Right())
}

func (i *Interpreter) VisitGroupingExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	return i.evaluate(exprV.Expression())
}

func (i *Interpreter) VisitUnaryExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	right, err := i.evaluate(exprV.Right())
	if err != nil {
		return nil, err
	}

	switch exprV.Operator().Type() {
	case api.MINUS:
		v, err := i.extractNumber(right, exprV.Operator())
		if err == nil {
			nl := literals.NewNumberLiteralVal(-v)
			return nl, nil
		}

		iv, err := i.extractInteger(right, exprV.Operator())
		if err == nil {
			nl := literals.NewIntegerLiteralVal(-iv)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "Minus expression invalid operand.")
	case api.BANG:
		return !i.isTruthy(right), nil
	}

	// Unreachable
	return nil, errors.NewRuntimeError(exprV.Operator(), "Unary expression hit unreachable code.")
}

func (i *Interpreter) VisitBinaryExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	left, errl := i.evaluate(exprV.Left())
	if errl != nil {
		return nil, errl
	}
	right, errr := i.evaluate(exprV.Right())
	if errr != nil {
		return nil, errr
	}

	switch exprV.Operator().Type() {
	case api.GREATER:
		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l > r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il > ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'>' Unexpected reachable code.")
	case api.GREATER_EQUAL:
		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l >= r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il >= ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'>=' Unexpected reachable code.")
	case api.LESS:
		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l < r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il < ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'<' Unexpected reachable code.")
	case api.LESS_EQUAL:
		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l <= r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il <= ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'<=' Unexpected reachable code.")
	case api.BANG_EQUAL:
		// if !i.isEqual(left, right) {
		// 	return literals.NewBooleanLiteral(true), nil
		// }

		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l != r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il != ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'!=' Unexpected reachable code.")
	case api.EQUAL_EQUAL:
		// if i.isEqual(left, right) {
		// 	return literals.NewBooleanLiteral(true), nil
		// }

		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(l == r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewBooleanLiteral(il == ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'==' Unexpected reachable code.")
	case api.MINUS:
		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewNumberLiteralVal(l - r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewIntegerLiteralVal(il - ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'"+exprV.Operator().Lexeme()+"' Operands must be two numbers.")
	case api.PLUS:
		// Numbers(floats or ints) or Strings

		// If one is a string then both must be strings
		sl, sr, err := i.extractStrings(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewStringLiteral(sl + sr)
			return nl, nil
		}

		l, r, err := i.extractNumbers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewNumberLiteralVal(l + r)
			return nl, nil
		}

		il, ir, err := i.extractIntegers(left, right, exprV.Operator())
		if err == nil {
			nl := literals.NewIntegerLiteralVal(il + ir)
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "'+' Operands must be two numbers or two strings.")
	case api.SLASH:
		// Both operands will be converted to floats before division
		lfv, isNumL := left.(api.INumberLiteral)
		rfv, isNumR := right.(api.INumberLiteral)

		var l, r float64
		if !isNumL {
			v, _ := left.(api.IIntegerLiteral)
			l = float64(v.IntValue())
		} else {
			l = lfv.NumValue()
		}

		if !isNumR {
			v, _ := right.(api.IIntegerLiteral)
			r = float64(v.IntValue())
		} else {
			r = rfv.NumValue()
		}

		nl := literals.NewNumberLiteralVal(l / r)
		return nl, nil
	case api.STAR:
		lfv, isNumL := left.(api.INumberLiteral)
		rfv, isNumR := right.(api.INumberLiteral)
		if isNumL || isNumR {
			// One maybe an integer literal
			var l, r float64
			if !isNumL {
				v, _ := left.(api.IIntegerLiteral)
				l = float64(v.IntValue())
			} else {
				l = lfv.NumValue()
			}

			if !isNumR {
				v, _ := right.(api.IIntegerLiteral)
				r = float64(v.IntValue())
			} else {
				r = rfv.NumValue()
			}

			nl := literals.NewNumberLiteralVal(l * r)
			return nl, nil
		}

		liv, liok := left.(api.IIntegerLiteral)
		riv, riok := right.(api.IIntegerLiteral)
		if liok && riok {
			nl := literals.NewIntegerLiteralVal(liv.IntValue() * riv.IntValue())
			return nl, nil
		}

		return nil, errors.NewRuntimeError(exprV.Operator(), "At least one operand must be a Number or both Integers.")
	}

	// Unreachable
	return nil, errors.NewRuntimeError(exprV.Operator(), "Binary expression hit unreachable code.")
}

func (i *Interpreter) VisitVariableExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	obj, err = i.lookUpVariable(exprV)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (i *Interpreter) VisitAssignExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	value, err := i.evaluate(exprV.Expression()) // i.e. exprV.Value()
	if err != nil {
		return nil, err
	}

	if distance, ok := i.locals[exprV]; ok {
		err = i.environment.AssignAt(distance, exprV.Name(), value)
		if err != nil {
			return nil, err
		}
	} else {
		err = i.globals.Assign(exprV.Name(), value)
		if err != nil {
			return nil, err
		}
	}

	return value, nil

}

func (i *Interpreter) VisitCallExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
	// The expression is just an identifier that looks up the
	// function by its name, but it could be anything.
	callee, err := i.evaluate(exprV.Callee())
	if err != nil {
		return nil, err
	}

	arguments := make([]interface{}, 0) // = []interface{}{}

	for _, argument := range exprV.Arguments() {
		value, err := i.evaluate(argument)
		if err != nil {
			return nil, err
		}
		arguments = append(arguments, value)
	}

	function, ok := callee.(api.ICallable)
	if !ok {
		return nil, errors.NewRuntimeError(exprV.Name(), "Can only call functions and classes.")
	}

	if len(arguments) != function.Arity() {
		msg := fmt.Sprintf("Expected %d arguments but got %d.", function.Arity(), len(arguments))
		return nil, errors.NewRuntimeError(exprV.Paren(), msg)
	}

	// The implementer’s job is
	// to return the value that the call expression produces.
	return function.Call(i, arguments)
}
