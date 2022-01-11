package api

type ICallable interface {
	Arity() int
	Call(interpreter IInterpreter, arguments []interface{}) (obj interface{}, err IRuntimeError)
}
