package statements

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type Statement struct {
}

func (s *Statement) Expression() api.IExpression {
	return nil
}

func (s *Statement) Initializer() api.IExpression {
	return nil
}

func (s *Statement) Name() api.IToken {
	return nil
}

func (s *Statement) Statements() []api.IStatement {
	return nil
}

func (s *Statement) Condition() api.IExpression {
	return nil
}

func (s *Statement) ThenBranch() api.IStatement {
	return nil
}

func (s *Statement) ElseBranch() api.IStatement {
	return nil
}

func (s *Statement) Body() []api.IStatement {
	return nil
}

func (s *Statement) Type() api.InterruptType {
	return api.INTERRUPT_UNKNOWN
}

func (s *Statement) StmtType() api.StatementType {
	return api.STMT_UNKNOWN
}

func (s *Statement) Parameters() []api.IToken {
	return nil
}

func (s *Statement) Keyword() api.IToken {
	return nil
}

func (s *Statement) Value() api.IExpression {
	return nil
}

func (s Statement) String() string {
	return ""
}

// ---------------------------------------------------
// Expression statement
// ---------------------------------------------------
type ExpressionStatement struct {
	Statement

	expression api.IExpression
}

func NewExpressionStatement(expression api.IExpression) api.IStatement {
	o := new(ExpressionStatement)
	o.expression = expression
	return o
}

func (s *ExpressionStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitExpressionStatement(s)
}

func (s *ExpressionStatement) Expression() api.IExpression {
	return s.expression
}

func (s ExpressionStatement) String() string {
	return "Expression"
}

// ---------------------------------------------------
// Block statement
// ---------------------------------------------------
type BlockStatement struct {
	Statement

	statements []api.IStatement
}

func NewBlockStatement(statements []api.IStatement) api.IStatement {
	o := new(BlockStatement)
	o.statements = statements
	return o
}

func (s *BlockStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitBlockStatement(s)
}

func (s *BlockStatement) Statements() []api.IStatement {
	return s.statements
}

func (s BlockStatement) String() string {
	return "BlockStatement"
}

// ---------------------------------------------------
// Print statement
// ---------------------------------------------------
type PrintStatement struct {
	Statement

	expression api.IExpression
}

func NewPrintStatement(expression api.IExpression) api.IStatement {
	o := new(PrintStatement)
	o.expression = expression
	return o
}

func (s *PrintStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitPrintStatement(s)
}

func (s *PrintStatement) Expression() api.IExpression {
	return s.expression
}

func (s PrintStatement) String() string {
	msg := ""

	if s.expression.Name() == nil {
		msg = fmt.Sprintf("PrintStatement type: [%s]", s.expression.Type())
	} else {
		msg = fmt.Sprintf("PrintStatement line: [%d]", s.expression.Name().Line())
	}
	return msg
}

// ---------------------------------------------------
// var statement
// ---------------------------------------------------
type VarStatement struct {
	Statement

	name        api.IToken
	initializer api.IExpression
}

func NewVarStatement(name api.IToken, initializer api.IExpression) api.IStatement {
	o := new(VarStatement)
	o.name = name
	o.initializer = initializer
	return o
}

func (s *VarStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitVariableStatement(s)
}

func (s *VarStatement) Initializer() api.IExpression {
	return s.initializer
}

func (s *VarStatement) Name() api.IToken {
	return s.name
}

func (s VarStatement) String() string {
	return s.name.String()
}

// ---------------------------------------------------
// "if" statement
// ---------------------------------------------------
type IfStatement struct {
	Statement

	condition  api.IExpression
	thenBranch api.IStatement
	elseBranch api.IStatement
}

func NewIfStatement(condition api.IExpression, thenBranch, elseBranch api.IStatement) api.IStatement {
	o := new(IfStatement)
	o.condition = condition
	o.thenBranch = thenBranch
	o.elseBranch = elseBranch
	return o
}

func (s *IfStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitIfStatement(s)
}

func (s *IfStatement) Condition() api.IExpression {
	return s.condition
}

func (s *IfStatement) ThenBranch() api.IStatement {
	return s.thenBranch
}

func (s *IfStatement) ElseBranch() api.IStatement {
	return s.elseBranch
}

func (s IfStatement) String() string {
	return fmt.Sprintf("IfStatement line: [%d]", (s.condition.Operator().Line()))
}

// ---------------------------------------------------
// "while" statement
// ---------------------------------------------------
type WhileStatement struct {
	Statement

	condition api.IExpression
	body      []api.IStatement
}

func NewWhileStatement(condition api.IExpression, body api.IStatement) api.IStatement {
	o := new(WhileStatement)
	o.condition = condition
	o.body = []api.IStatement{body}
	return o
}

func (s *WhileStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitWhileStatement(s)
}

func (s *WhileStatement) Condition() api.IExpression {
	return s.condition
}

func (s *WhileStatement) Body() []api.IStatement {
	return s.body
}

func (s WhileStatement) String() string {
	return "WhileStatement"
}

// ---------------------------------------------------
// "break", "continue" interrupt statements
// ---------------------------------------------------
type InterruptStatement struct {
	Statement

	name  api.IToken
	iType api.InterruptType
}

func NewInterruptStatement(name api.IToken, iType api.InterruptType) api.IStatement {
	o := new(InterruptStatement)
	o.iType = iType
	o.name = name
	return o
}

func (s *InterruptStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitInterruptStatement(s)
}

func (s *InterruptStatement) Type() api.InterruptType {
	return s.iType
}

func (s *InterruptStatement) Name() api.IToken {
	return s.name
}

func (s InterruptStatement) String() string {
	return "InterruptStatement " + s.iType.String()
}

// ---------------------------------------------------
// "fun" statement
// ---------------------------------------------------
type FunctionStatement struct {
	Statement

	name   api.IToken
	params []api.IToken
	body   []api.IStatement
}

func NewFunctionStatement(name api.IToken, params []api.IToken, body []api.IStatement) api.IStatement {
	o := new(FunctionStatement)
	o.name = name
	o.params = params
	o.body = body
	return o
}

func (s *FunctionStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitFunctionStatement(s)
}

func (s *FunctionStatement) Name() api.IToken {
	return s.name
}

func (s *FunctionStatement) Parameters() []api.IToken {
	return s.params
}
func (s *FunctionStatement) Body() []api.IStatement {
	return s.body
}

func (s FunctionStatement) String() string {
	return "FunctionStatement " + s.name.Lexeme()
}

// ---------------------------------------------------
// "return" statement
// ---------------------------------------------------
type ReturnStatement struct {
	Statement

	keyword api.IToken
	value   api.IExpression

	iType api.InterruptType
}

func NewReturnStatement(keyword api.IToken, value api.IExpression) api.IStatement {
	o := new(ReturnStatement)
	o.keyword = keyword
	o.value = value
	o.iType = api.INTERRUPT_RETURN
	return o
}

func (s *ReturnStatement) Accept(visitor api.IVisitorStatement) (err api.IRuntimeError) {
	return visitor.VisitReturnStatement(s)
}

func (s *ReturnStatement) Keyword() api.IToken {
	return s.keyword
}
func (s *ReturnStatement) Value() api.IExpression {
	return s.value
}

func (s *ReturnStatement) Type() api.InterruptType {
	return s.iType
}

func (s *ReturnStatement) StmtType() api.StatementType {
	return api.STMT_RETURN
}

func (s ReturnStatement) String() string {
	return "ReturnStatement"
}
