package api

type IResolver interface {
	Resolve(statements []IStatement) (err IRuntimeError)
}
