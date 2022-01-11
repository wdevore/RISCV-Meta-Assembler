package literals

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type StringLiteral struct {
	value string
}

func NewStringLiteral(value string) api.IStringLiteral {
	s := new(StringLiteral)
	s.value = value
	return s
}

func (s StringLiteral) String() string {
	return "'" + s.value + "'"
}

func (s *StringLiteral) StringValue() string {
	return s.value
}
