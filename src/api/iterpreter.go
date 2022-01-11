package api

type IInterpreter interface {
	Interpret(statements []IStatement) IRuntimeError
	Globals() IEnvironment
	ExecuteBlock(statements []IStatement, parentEnv IEnvironment) (err IRuntimeError)
	Resolve(expression IExpression, depth int) IRuntimeError
}
