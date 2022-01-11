package api

type IToken interface {
	Lexeme() string
	Type() TokenType
	Literal() ILiteral
	Line() int
	String() string
}
