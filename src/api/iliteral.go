package api

type ILiteral interface {
	String() string
}

type IStringLiteral interface {
	ILiteral
	StringValue() string
}

type ICharLiteral interface {
	ILiteral
	CharValue() rune
}

type IIntegerLiteral interface {
	ILiteral
	IntValue() int
}

type INumberLiteral interface {
	ILiteral
	NumValue() float64
}

type IBooleanLiteral interface {
	ILiteral
	BoolValue() bool
}

type IHexNumberLiteral interface {
	ILiteral
	HexValue() string
}

type IBinaryNumberLiteral interface {
	ILiteral
	BinValue() string
}

type INilLiteral interface {
	ILiteral
	NilValue() string
}
