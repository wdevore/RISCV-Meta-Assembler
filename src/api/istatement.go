package api

type StatementType int64

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	STMT_UNKNOWN StatementType = iota

	STMT_RETURN
)

type IStatement interface {
	Accept(IVisitorStatement) (err IRuntimeError)

	Expression() IExpression

	// Var statement
	Name() IToken
	Initializer() IExpression

	// Blocks
	Statements() []IStatement

	// "If"
	Condition() IExpression
	ThenBranch() IStatement
	ElseBranch() IStatement

	// "While", Functions
	Body() []IStatement

	// "break", "continue" interrupts
	Type() InterruptType

	StmtType() StatementType

	// Functions
	Parameters() []IToken

	// "return"
	Keyword() IToken
	Value() IExpression
}
