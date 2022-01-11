package api

type ExpressionType int64

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	UNDEFINED_EXPR ExpressionType = iota
	BINARY_EXPR
	GROUPING_EXPR
	LITERAL_EXPR
	UNARY_EXPR
	VAR_EXPR
	ASSIGN_EXPR
	LOGIC_EXPR
	WHILE_EXPR
	CALL_EXPR
	FUN_EXPR
)

type IExpression interface {
	Accept(IVisitorExpression) (obj interface{}, err IRuntimeError)

	// Literals
	Value() interface{}

	// Unary,Binary
	Left() IExpression
	Operator() IToken
	Right() IExpression

	// Grouping
	Expression() IExpression

	// Var
	Name() IToken

	// What type of expression
	Type() ExpressionType

	// Call
	Callee() IExpression
	Paren() IToken
	Arguments() []IExpression
}

func (e ExpressionType) String() string {
	switch e {
	case UNDEFINED_EXPR:
		return "-?-"
	case BINARY_EXPR:
		return "BinaryExpression"
	case GROUPING_EXPR:
		return "GroupingExpression"
	case LITERAL_EXPR:
		return "LiteralExpression"
	case UNARY_EXPR:
		return "UnaryExpression"
	case VAR_EXPR:
		return "VariableExpression"
	case ASSIGN_EXPR:
		return "AssignmentExpression"
	case LOGIC_EXPR:
		return "LogicExpression"
	case WHILE_EXPR:
		return "WhileExpression"
	case CALL_EXPR:
		return "CallableExpression"
	case FUN_EXPR:
		return "FunctionExpression"
	}

	return "unknown"
}
