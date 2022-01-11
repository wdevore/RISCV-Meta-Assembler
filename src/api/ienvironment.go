package api

type IEnvironment interface {
	Assign(name IToken, obj interface{}) IRuntimeError
	AssignAt(distance int, name IToken, obj interface{}) IRuntimeError
	Define(name string, obj interface{}) IRuntimeError
	Get(name IToken) (obj interface{}, err IRuntimeError)
	GetAt(distance int, name IToken) (obj interface{}, err IRuntimeError)
	Enclosing() IEnvironment
	Values() map[string]interface{}
}
