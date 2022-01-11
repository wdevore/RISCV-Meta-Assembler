package literals

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type CharLiteral struct {
	value rune
}

func NewCharLiteral(value rune) api.ICharLiteral {
	s := new(CharLiteral)
	s.value = value
	return s
}

func (c CharLiteral) String() string {
	return string(c.value)
}

func (c *CharLiteral) CharValue() rune {
	return c.value
}
