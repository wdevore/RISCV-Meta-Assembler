package errors

import (
	"log"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
)

type Report struct {
}

func NewReport() api.IReporter {
	o := new(Report)
	return o
}

func (r *Report) ReportLine(line int, message string) {
	log.Printf("[line %d] Error: %s", line, message)
}

func (r *Report) ReportWhere(line int, where, message string) {
	log.Printf("[line %d] Error %s : %s", line, where, message)
}
