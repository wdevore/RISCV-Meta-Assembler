package api

type IVisitorStatement interface {
	VisitExpressionStatement(IStatement) (err IRuntimeError)
	VisitPrintStatement(IStatement) (err IRuntimeError)
	VisitVariableStatement(IStatement) (err IRuntimeError)
	VisitBlockStatement(IStatement) (err IRuntimeError)
	VisitIfStatement(IStatement) (err IRuntimeError)
	VisitWhileStatement(IStatement) (err IRuntimeError)
	VisitInterruptStatement(IStatement) (err IRuntimeError)
	VisitFunctionStatement(IStatement) (err IRuntimeError)
	VisitReturnStatement(IStatement) (err IRuntimeError)
}
