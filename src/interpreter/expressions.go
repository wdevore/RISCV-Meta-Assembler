package interpreter

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type BaseExpression struct {
}

func (e *BaseExpression) Value() interface{} {
	return nil
}

func (e *BaseExpression) Left() api.IExpression {
	return nil
}

func (e *BaseExpression) Operator() api.IToken {
	return nil
}

func (e *BaseExpression) Right() api.IExpression {
	return nil
}

func (e *BaseExpression) Expression() api.IExpression {
	return nil
}

func (e *BaseExpression) Name() api.IToken {
	return nil
}

func (e *BaseExpression) Type() api.ExpressionType {
	return api.UNDEFINED_EXPR
}

func (e *BaseExpression) Callee() api.IExpression {
	return nil
}

func (e *BaseExpression) Paren() api.IToken {
	return nil
}

func (e *BaseExpression) Arguments() []api.IExpression {
	return nil
}

// ---------------------------------------------------
// Binary
// ---------------------------------------------------
type BinaryExpression struct {
	BaseExpression

	eType api.ExpressionType

	left     api.IExpression
	operator api.IToken
	right    api.IExpression
}

func NewBinaryExpression(left api.IExpression, operator api.IToken, right api.IExpression) api.IExpression {
	e := new(BinaryExpression)
	e.left = left
	e.operator = operator
	e.right = right
	e.eType = api.BINARY_EXPR
	return e
}

func (e *BinaryExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitBinaryExpression(e)
}

func (e *BinaryExpression) Left() api.IExpression {
	return e.left
}

func (e *BinaryExpression) Operator() api.IToken {
	return e.operator
}

func (e *BinaryExpression) Right() api.IExpression {
	return e.right
}

func (e *BinaryExpression) Type() api.ExpressionType {
	return e.eType
}

// ---------------------------------------------------
// Grouping
// ---------------------------------------------------
type GroupingExpression struct {
	BaseExpression

	eType api.ExpressionType

	expression api.IExpression
}

func NewGroupingExpression(expression api.IExpression) api.IExpression {
	e := new(GroupingExpression)
	e.expression = expression
	e.eType = api.GROUPING_EXPR
	return e
}

func (e *GroupingExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitGroupingExpression(e)
}

func (e *GroupingExpression) Expression() api.IExpression {
	return e.expression
}

func (e *GroupingExpression) Type() api.ExpressionType {
	return e.eType
}

// ---------------------------------------------------
// Literal
// ---------------------------------------------------
type LiteralExpression struct {
	BaseExpression

	eType api.ExpressionType

	name  api.IToken
	value api.ILiteral
}

func NewLiteralExpression(name api.IToken, value api.ILiteral) api.IExpression {
	e := new(LiteralExpression)
	e.value = value
	e.name = name
	e.eType = api.LITERAL_EXPR
	return e
}

func (e *LiteralExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitLiteralExpression(e)
}

func (e *LiteralExpression) Value() interface{} {
	return e.value
}

func (e *LiteralExpression) Type() api.ExpressionType {
	return e.eType
}

func (e *LiteralExpression) Name() api.IToken {
	return e.name
}

// ---------------------------------------------------
// Unary
// ---------------------------------------------------
type UnaryExpression struct {
	BaseExpression

	eType api.ExpressionType

	operator api.IToken
	right    api.IExpression
}

func NewUnaryExpression(operator api.IToken, right api.IExpression) api.IExpression {
	e := new(UnaryExpression)
	e.operator = operator
	e.right = right
	e.eType = api.UNARY_EXPR
	return e
}

func (e *UnaryExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitUnaryExpression(e)
}

func (e *UnaryExpression) Operator() api.IToken {
	return e.operator
}

func (e *UnaryExpression) Right() api.IExpression {
	return e.right
}

func (e *UnaryExpression) Type() api.ExpressionType {
	return e.eType
}

// ---------------------------------------------------
// Variable
// ---------------------------------------------------
type VariableExpression struct {
	BaseExpression

	eType api.ExpressionType

	name api.IToken
}

func NewVariableExpression(name api.IToken) api.IExpression {
	e := new(VariableExpression)
	e.name = name
	e.eType = api.VAR_EXPR
	return e
}

func (e *VariableExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitVariableExpression(e)
}

func (e *VariableExpression) Name() api.IToken {
	return e.name
}

func (e *VariableExpression) Type() api.ExpressionType {
	return e.eType
}

// ---------------------------------------------------
// Assignment "="
// ---------------------------------------------------
type AssignExpression struct {
	BaseExpression

	eType api.ExpressionType

	// An l-value “evaluates” to a storage location that you can
	// assign into.
	name       api.IToken      // "l-value" The token for the variable being assigned to
	expression api.IExpression // and an expression for the new value
}

func NewAssignExpression(name api.IToken, eValue api.IExpression) api.IExpression {
	e := new(AssignExpression)
	e.name = name
	e.expression = eValue // l-value
	e.eType = api.ASSIGN_EXPR
	return e
}

func (e *AssignExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitAssignExpression(e)
}

func (e *AssignExpression) Expression() api.IExpression {
	return e.expression // Assignments are l-values
}

func (e *AssignExpression) Name() api.IToken {
	return e.name
}

func (e *AssignExpression) Type() api.ExpressionType {
	return e.eType
}

// ---------------------------------------------------
// Logic "or", "and"
// ---------------------------------------------------
type LogicExpression struct {
	BaseExpression

	eType api.ExpressionType

	left     api.IExpression
	operator api.IToken
	right    api.IExpression
}

func NewLogicExpression(left api.IExpression, operator api.IToken, right api.IExpression) api.IExpression {
	e := new(LogicExpression)
	e.left = left
	e.operator = operator
	e.right = right
	e.eType = api.LOGIC_EXPR
	return e
}

func (e *LogicExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitLogicalExpression(e)
}

func (e *LogicExpression) Left() api.IExpression {
	return e.left
}

func (e *LogicExpression) Operator() api.IToken {
	return e.operator
}

func (e *LogicExpression) Right() api.IExpression {
	return e.right
}

func (e *LogicExpression) Type() api.ExpressionType {
	return e.eType
}

// ---------------------------------------------------
// "call"
// ---------------------------------------------------
type CallExpression struct {
	BaseExpression

	eType api.ExpressionType

	callee api.IExpression
	// token’s location when we report a runtime error caused by a function call
	// the token for the closing parenthesis.
	paren     api.IToken
	arguments []api.IExpression
}

func NewCallExpression(callee api.IExpression, paren api.IToken, arguments []api.IExpression) api.IExpression {
	e := new(CallExpression)
	e.callee = callee
	e.paren = paren
	e.arguments = arguments
	e.eType = api.CALL_EXPR
	return e
}

func (e *CallExpression) Accept(visitor api.IVisitorExpression) (obj interface{}, err api.IRuntimeError) {
	return visitor.VisitCallExpression(e)
}

func (e *CallExpression) Type() api.ExpressionType {
	return e.eType
}

func (e *CallExpression) Callee() api.IExpression {
	return e.callee
}

func (e *CallExpression) Paren() api.IToken {
	return e.paren
}

func (e *CallExpression) Arguments() []api.IExpression {
	return e.arguments
}
