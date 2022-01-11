package literals

import (
	"fmt"
	"strconv"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type NumberLiteral struct {
	value float64
}

func NewNumberLiteral(value string) api.INumberLiteral {
	s := new(NumberLiteral)
	s.value, _ = strconv.ParseFloat(value, 32)
	return s
}

func NewNumberLiteralVal(value float64) api.INumberLiteral {
	s := new(NumberLiteral)
	s.value = value
	return s
}

func (n NumberLiteral) String() string {
	return fmt.Sprintf("%f", n.value)
}

func (n *NumberLiteral) NumValue() float64 {
	return n.value
}
