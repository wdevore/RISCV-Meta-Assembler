package literals

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type HexNumberLiteral struct {
	value string
}

func NewHexNumberLiteral(value string) api.IHexNumberLiteral {
	s := new(HexNumberLiteral)
	l := len(value)
	if l < 3 {
		value = fmt.Sprintf("%02s", value)
	} else if l < 5 {
		value = fmt.Sprintf("%04s", value)
	} else {
		value = fmt.Sprintf("%08s", value)
	}

	s.value = value
	return s
}

func (h HexNumberLiteral) String() string {
	return fmt.Sprintf("0x%s", h.value)
}

func (h *HexNumberLiteral) HexValue() string {
	return h.value
}
