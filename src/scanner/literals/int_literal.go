package literals

import (
	"fmt"
	"strconv"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type IntegerLiteral struct {
	value int
}

func NewIntegerLiteral(value string) api.IIntegerLiteral {
	s := new(IntegerLiteral)
	pi, _ := strconv.ParseInt(value, 10, 32)
	s.value = int(pi)
	return s
}

func NewIntegerLiteralVal(value int) api.IIntegerLiteral {
	s := new(IntegerLiteral)
	s.value = value
	return s
}

func (i IntegerLiteral) String() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *IntegerLiteral) IntValue() int {
	return i.value
}
