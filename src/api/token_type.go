package api

type TokenType int64

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	UNDEFINED TokenType = iota

	// Single-character tokens.
	LEFT_PAREN // "("
	RIGHT_PAREN
	LEFT_BRACE // "{"
	RIGHT_BRACE
	LEFT_BRACKET // "["
	RIGHT_BRACKET
	COMMA
	SEMICOLON
	DOT
	MINUS
	PLUS
	SLASH // forward slash "/"
	STAR
	TRUE
	FALSE
	NIL

	// One or two character tokens.
	BANG // "!"
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER // ">"
	GREATER_EQUAL
	LESS // "<"
	LESS_EQUAL
	PERCENT

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	PRINT
	IF
	ELSE
	VAR
	WHILE
	FOR
	BREAK
	CONTINUE
	FUN
	RETURN
	CONST
	IMPORT
	CODE
	ALIGN_TO
	GLOBAL
	AT
	AS
	USE
	READ_ONLY
	BYTE
	HALF
	WORD
	DATA
	INT
	HI
	LO

	// RISC-V real instructions
	ADD
	SUB
	XOR
	OR
	AND
	SLL
	SRL
	SRA
	SLT
	SLTU
	ADDI
	XORI
	ORI
	ANDI
	SLLI
	SRLI
	SRAI
	SLTI
	SLTIU
	LB // also a pseudo instruction
	LH // ditto
	LW
	LBU
	LHU
	SB
	SH
	SW
	BEQ
	BNE
	BLT
	BGE
	BLTU
	BGEU
	JAL  // also can be a pseudo instruction
	JALR // ditto
	LUI
	AUIPC
	ECALL
	EBREAK

	// RISC-V pseudo instructions
	LA
	NOP
	LI
	MV
	NOT
	NEG
	NEGW
	SEXT
	SEQZ
	SNEZ
	SLTZ
	SGTZ
	BEQZ
	BNEZ
	BLEZ
	BGEZ
	BLTZ
	BGTZ
	BGT
	BLE
	BGTU
	BLEU
	J
	RET
	CALL
	TAIL

	EOF
)

func (t TokenType) String() string {
	switch t {
	case UNDEFINED:
		return "-?-"
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case LEFT_BRACKET:
		return "["
	case RIGHT_BRACKET:
		return "]"
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case DOT:
		return "."
	case MINUS:
		return "-"
	case PLUS:
		return "+"
	case SLASH:
		return "/"
	case STAR:
		return "*"
	case BANG:
		return "!"
	case BANG_EQUAL:
		return "!="
	case EQUAL:
		return "="
	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="
	case PERCENT:
		return "%"
	case IDENTIFIER:
		return "identifier"
	case STRING:
		return "string"
	case NUMBER:
		return "number"
	case CONST:
		return "const"
	case IMPORT:
		return "import"
	case CODE:
		return "code"
	case ALIGN_TO:
		return "alignTo"
	case GLOBAL:
		return "global"
	case AT:
		return "at"
	case AS:
		return "as"
	case USE:
		return "use"
	case READ_ONLY:
		return "readOnly"
	case BYTE:
		return "byte"
	case HALF:
		return "half"
	case WORD:
		return "word"
	case DATA:
		return "data"
	case INT:
		return "int"
	case HI:
		return "hi"
	case LO:
		return "lo"
	case ADD:
		return "add"
	case SUB:
		return "sub"
	case XOR:
		return "xor"
	case OR:
		return "or"
	case AND:
		return "and"
	case SLL:
		return "sll"
	case SRL:
		return "srl"
	case SRA:
		return "sra"
	case SLT:
		return "slt"
	case SLTU:
		return "sltu"
	case ADDI:
		return "addi"
	case XORI:
		return "xori"
	case ORI:
		return "ori"
	case ANDI:
		return "andi"
	case SLLI:
		return "slli"
	case SRLI:
		return "srli"
	case SRAI:
		return "srai"
	case SLTI:
		return "slti"
	case SLTIU:
		return "sltiu"
	case LB:
		return "lb"
	case LH:
		return "lh"
	case LW:
		return "lw"
	case LBU:
		return "lbu"
	case LHU:
		return "lhu"
	case SB:
		return "sb"
	case SH:
		return "sh"
	case SW:
		return "sw"
	case BEQ:
		return "beq"
	case BNE:
		return "bne"
	case BLT:
		return "blt"
	case BGE:
		return "bge"
	case JAL:
		return "jal"
	case JALR:
		return "jalr"
	case LUI:
		return "lui"
	case AUIPC:
		return "auipc"
	case ECALL:
		return "ecall"
	case EBREAK:
		return "ebreak"
	case LA:
		return "la"
	case NOP:
		return "nop"
	case LI:
		return "li"
	case MV:
		return "mv"
	case NOT:
		return "not"
	case NEG:
		return "neg"
	case NEGW:
		return "negw"
	case SEXT:
		return "sext"
	case SEQZ:
		return "seqz"
	case SNEZ:
		return "snez"
	case SLTZ:
		return "sltz"
	case SGTZ:
		return "sgtz"
	case BEQZ:
		return "beqz"
	case BNEZ:
		return "bnez"
	case BLEZ:
		return "blez"
	case BGEZ:
		return "bgez"
	case BLTZ:
		return "bltz"
	case BGTZ:
		return "bgtz"
	case BGT:
		return "bgt"
	case BLE:
		return "ble"
	case BGTU:
		return "bgtu"
	case BLEU:
		return "bleu"
	case J:
		return "j"
	case RET:
		return "ret"
	case CALL:
		return "call"
	case TAIL:
		return "tail"
	case EOF:
		return "eof"
	}
	return "unknown"
}
