package literals

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

type NilLiteral struct {
	value string
}

func NewNilLiteral() api.INilLiteral {
	s := new(NilLiteral)
	s.value = "nil"
	return s
}

func (n NilLiteral) String() string {
	return n.value
}

func (n *NilLiteral) NilValue() string {
	return n.value
}
