package parser

import "github.com/wdevore/RISCV-Meta-Assembler/src/api"

// It discards tokens until it thinks it found a statement boundary.
func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type() == api.RIGHT_BRACE {
			return
		}

		switch p.peek().Type() {
		case api.CONST,
			api.IMPORT,
			api.CODE,
			api.ALIGN_TO,
			api.GLOBAL,
			api.AT,
			api.AS,
			api.USE,
			api.READ_ONLY,
			api.BYTE,
			api.HALF,
			api.WORD,
			api.DATA,
			api.INT,
			api.HI,
			api.LO,
			api.ADD,
			api.SUB,
			api.XOR,
			api.OR,
			api.AND,
			api.SLL,
			api.SRL,
			api.SRA,
			api.SLT,
			api.SLTU,
			api.ADDI,
			api.XORI,
			api.ORI,
			api.ANDI,
			api.SLLI,
			api.SRLI,
			api.SRAI,
			api.SLTI,
			api.SLTIU,
			api.LB,
			api.LH,
			api.LW,
			api.LBU,
			api.LHU,
			api.SB,
			api.SH,
			api.SW,
			api.BEQ,
			api.BNE,
			api.BLT,
			api.BGE,
			api.BLTU,
			api.BGEU,
			api.JAL,
			api.JALR,
			api.LUI,
			api.AUIPC,
			api.ECALL,
			api.EBREAK,
			api.LA,
			api.NOP,
			api.LI,
			api.MV,
			api.NOT,
			api.NEG,
			api.NEGW,
			api.SEXT,
			api.SEQZ,
			api.SNEZ,
			api.SLTZ,
			api.SGTZ,
			api.BEQZ,
			api.BNEZ,
			api.BLEZ,
			api.BGEZ,
			api.BLTZ,
			api.BGTZ,
			api.BGT,
			api.BLE,
			api.BGTU,
			api.BLEU,
			api.J,
			api.RET,
			api.CALL,
			api.TAIL:
			return
		}
	}

	p.advance()
}
