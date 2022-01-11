package api

type IVisitorExpression interface {
	VisitBinaryExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitGroupingExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitLiteralExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitUnaryExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitVariableExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitAssignExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitLogicalExpression(IExpression) (obj interface{}, err IRuntimeError)
	VisitCallExpression(IExpression) (obj interface{}, err IRuntimeError)
}
