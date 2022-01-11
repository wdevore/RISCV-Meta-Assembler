package errors

import (
	"fmt"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type RuntimeError struct {
	interruptType api.InterruptType

	token   api.IToken
	message string

	value interface{}
}

func NewRuntimeError(token api.IToken, message string) api.IRuntimeError {
	o := new(RuntimeError)
	o.token = token
	o.message = message
	return o
}

func (r *RuntimeError) Token() api.IToken {
	return r.token
}

func (r *RuntimeError) Message() string {
	return r.message
}

func (r RuntimeError) String() string {
	if r.token != nil {
		return fmt.Sprintf("[line %d] %s", r.token.Line(), r.message)
	}

	return r.message
}

func (r *RuntimeError) Interrupt() api.InterruptType {
	return r.interruptType
}

func (r *RuntimeError) ClearInterrupt() {
	r.interruptType = api.INTERRUPT_UNKNOWN
}

func (r *RuntimeError) SetInterrupt(iType api.InterruptType) {
	r.interruptType = iType
}

func (r *RuntimeError) Value() interface{} {
	return r.value
}

func (r *RuntimeError) SetValue(value interface{}) {
	r.value = value
}
