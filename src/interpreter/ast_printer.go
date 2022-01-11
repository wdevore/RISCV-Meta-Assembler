package interpreter

// import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type AstPrinter struct {
}

// func NewAstPrinter() api.IVisitorExpression {
// 	o := new(AstPrinter)
// 	return o
// }

// func (a *AstPrinter) Print(expr api.IExpression) string {
// 	obj, _ := expr.Accept(a)
// 	// TODO add err check
// 	return obj.(string)
// }

// // -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
// // IVisitorExpression interface
// // -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ -- ~~ --
// func (a *AstPrinter) VisitBinaryExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
// 	expr := exprV.(*BinaryExpression)
// 	return a.parenthesize(expr.operator.Lexeme(), expr.left, expr.right), nil
// }

// func (a *AstPrinter) VisitGroupingExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
// 	expr := exprV.(*GroupingExpression)
// 	return a.parenthesize("group", expr.expression), nil
// }

// func (a *AstPrinter) VisitLiteralExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
// 	expr := exprV.(*LiteralExpression)
// 	if expr.value == nil {
// 		return "nil", nil
// 	}
// 	return expr.value.String(), nil
// }

// func (a *AstPrinter) VisitUnaryExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
// 	expr := exprV.(*UnaryExpression)
// 	return a.parenthesize(expr.operator.Lexeme(), expr.right), nil
// }

// func (a *AstPrinter) VisitVariableExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
// 	return nil, nil
// }

// func (a *AstPrinter) VisitAssignExpression(exprV api.IExpression) (obj interface{}, err api.IRuntimeError) {
// 	return nil, nil
// }

// // *************************************

// // func (a *AstPrinter) Value() interface{} {
// // 	return nil
// // }

// func (a *AstPrinter) parenthesize(name string, expr ...interface{}) string {
// 	builder := "(" + name

// 	for _, expres := range expr {
// 		builder += " "
// 		ex := expres.(api.IExpression)
// 		obj, _ := ex.Accept(a)
// 		// TODO add err check

// 		builder += obj.(string)
// 	}
// 	builder += ")"

// 	return builder
// }
