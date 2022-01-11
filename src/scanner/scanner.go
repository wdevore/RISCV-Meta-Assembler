package scanner

import (
	"io/ioutil"
	"path/filepath"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/scanner/literals"
)

type Scanner struct {
	assembler api.IAssembler

	source string
	tokens []api.IToken

	start   int
	current int
	line    int
}

func NewScanner(assembler api.IAssembler) *Scanner {
	s := new(Scanner)
	s.start = 0
	s.current = 0
	s.line = 1
	s.assembler = assembler
	return s
}

func (s *Scanner) Scan(source string) (tokens []api.IToken, err error) {
	s.source = source

	dataPath, err := filepath.Abs(s.assembler.ConfigRelPath())

	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(dataPath + "/" + source)
	if err != nil {
		return nil, err
	}

	s.source = string(bytes)
	s.scanTokens(s.source)

	// for _, token := range s.tokens {
	// 	log.Println(token)
	// }

	return s.tokens, nil
}

func (s *Scanner) scanTokens(line string) {
	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.start = s.current
		s.scanToken()
	}

	t := NewToken(api.EOF, "", literals.NewNilLiteral(), 1 /*line*/)
	s.tokens = append(s.tokens, t)
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case "(":
		s.addTokenNullLiteral(api.LEFT_PAREN)
	case ")":
		s.addTokenNullLiteral(api.RIGHT_PAREN)
	case "{":
		s.addTokenNullLiteral(api.LEFT_BRACE)
	case "}":
		s.addTokenNullLiteral(api.RIGHT_BRACE)
	case "[":
		s.addTokenNullLiteral(api.LEFT_BRACKET)
	case "]":
		s.addTokenNullLiteral(api.RIGHT_BRACKET)
	case ",":
		s.addTokenNullLiteral(api.COMMA)
	case ";":
		s.addTokenNullLiteral(api.SEMICOLON)
	case ".":
		s.addTokenNullLiteral(api.DOT)
	case "-":
		s.addTokenNullLiteral(api.MINUS)
	case "+":
		s.addTokenNullLiteral(api.PLUS)
	case "*":
		s.addTokenNullLiteral(api.STAR)
	case "%":
		s.addTokenNullLiteral(api.PERCENT)
	case "!":
		match := s.match("=")
		if match {
			s.addTokenNullLiteral(api.BANG_EQUAL)
		} else {
			s.addTokenNullLiteral(api.BANG)
		}
	case "=":
		match := s.match("=")
		if match {
			s.addTokenNullLiteral(api.EQUAL_EQUAL)
		} else {
			s.addTokenNullLiteral(api.EQUAL)
		}
	case "<":
		match := s.match("=")
		if match {
			s.addTokenNullLiteral(api.LESS_EQUAL)
		} else {
			// It could be "<42>" example or just "<"
			if !s.isDigit(s.peek()) {
				s.addTokenNullLiteral(api.LESS)
			}
		}
	case ">":
		match := s.match("=")
		if match {
			s.addTokenNullLiteral(api.GREATER_EQUAL)
		} else {
			s.addTokenNullLiteral(api.GREATER)
		}

	case "/":
		match := s.match("/")

		if match {
			// A comment goes until the end of the line.
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addTokenNullLiteral(api.SLASH)
		}
	case " ", "\r", "\t":
		// Ignore whitespace.
	case "\n":
		s.line++
	case "\"":
		s.string()
	case "'":
		s.char()
	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			s.assembler.ReportLine(s.line, "Unexpected character '"+c+"'")
		}
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() string {
	s.current++
	return string(s.source[s.current-1])
}

func (s *Scanner) addTokenNullLiteral(ttype api.TokenType) {
	s.addToken(ttype, literals.NewNilLiteral())
}

func (s *Scanner) addToken(ttype api.TokenType, literal api.ILiteral) {
	text := s.source[s.start:s.current]
	token := NewToken(ttype, text, literal, s.line)
	s.tokens = append(s.tokens, token)
}

// We only consume the current character if it’s
// what we’re looking for.
func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.source[s.current]) != expected {
		return false
	}

	s.current++
	return true
}

// It’s sort of like advance() , but doesn’t consume the character.
// This is called lookahead.
func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\n"
	}
	return string(s.source[s.current])
}

func (s *Scanner) string() {
	for s.peek() != "\"" && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		s.assembler.ReportLine(s.line, "Unterminated string.")
		return
	}
	// The closing " character
	s.advance()

	// Trim the surrounding quotes.
	value := s.source[s.start+1 : s.current-1]
	s.addToken(api.STRING, literals.NewStringLiteral(value))
}

func (s *Scanner) char() {
	for s.peek() != "'" && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.assembler.ReportLine(s.line, "Unterminated char.")
		return
	}

	// The closing "'"
	s.advance()

	// Trim the surrounding single quotes.
	value := s.source[s.start+1 : s.current-1]

	// -2 for right "'" and advanced position
	if s.current-s.start-2 > 1 {
		s.assembler.ReportLine(s.line, "To many characters for '"+value+"'")
		return
	}

	s.addToken(api.STRING, literals.NewCharLiteral([]rune(value)[0]))
}

func (s *Scanner) isDigit(c string) bool {
	return c >= "0" && c <= "9"
}

func (s *Scanner) isAlpha(c string) bool {
	return (c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || c == "_"
}

func (s *Scanner) isAlphaNumeric(c string) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	// Look for a fractional part of a decimal number
	if s.peek() == "." && s.isDigit(s.peekNext()) {
		// Consume the "."
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}

		value := s.source[s.start:s.current]
		s.addToken(api.NUMBER, literals.NewNumberLiteral(value))
		return
	}

	// Look for a base specifier "x","b".
	if s.peek() == "x" && s.isAlphaNumeric(s.peekNext()) {
		// Consume the "x"
		s.advance()

		for s.isAlphaNumeric(s.peek()) {
			s.advance()
		}

		// Trim "0x"
		value := s.source[s.start+2 : s.current]
		s.addToken(api.NUMBER, literals.NewHexNumberLiteral(value))
		return
	}

	if s.peek() == "b" && s.isDigit(s.peekNext()) {
		// Consume the "b"
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}

		// Trim "0b"
		value := s.source[s.start+2 : s.current]
		s.addToken(api.NUMBER, literals.NewBinaryNumberLiteral(value))
		return
	}

	value := s.source[s.start:s.current]
	s.addToken(api.NUMBER, literals.NewIntegerLiteral(value))
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	ttype := Keywords[text]
	if ttype == api.UNDEFINED {
		ttype = api.IDENTIFIER
	}

	s.addTokenNullLiteral(ttype)
}

func (s *Scanner) peekNext() string {
	if s.current+1 >= len(s.source) {
		return "" // "\0"
	}

	return string(s.source[s.current+1])
}
