package literals

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type BinaryNumberLiteral struct {
	value string
}

func NewBinaryNumberLiteral(value string) api.IBinaryNumberLiteral {
	s := new(BinaryNumberLiteral)
	s.value = value
	return s
}

func (b BinaryNumberLiteral) String() string {
	return fmt.Sprintf("0b%s", b.value)
}

func (b *BinaryNumberLiteral) BinValue() string {
	return b.value
}
