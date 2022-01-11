package literals

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type BooleanLiteral struct {
	value bool
}

func NewBooleanLiteral(value bool) api.IBooleanLiteral {
	s := new(BooleanLiteral)
	s.value = value
	return s
}

func (b BooleanLiteral) String() string {
	return fmt.Sprintf("%v", b.value)
}

func (b *BooleanLiteral) BoolValue() bool {
	return b.value
}
