package parser

import (
	"errors"

	"github.com/wdevore/RISCV-Meta-Assembler/src/api"
	"github.com/wdevore/RISCV-Meta-Assembler/src/interpreter"
	"github.com/wdevore/RISCV-Meta-Assembler/src/scanner/literals"
	"github.com/wdevore/RISCV-Meta-Assembler/src/statements"
)

type Parser struct {
	assembler api.IAssembler
	tokens    []api.IToken
	current   int
}

func NewParser(assembler api.IAssembler, tokens []api.IToken) *Parser {
	o := new(Parser)
	o.tokens = tokens
	o.assembler = assembler
	return o
}

func (p *Parser) Parse() (statements []api.IStatement, err error) {
	statements = []api.IStatement{}

	for !p.isAtEnd() {
		statement, err := p.declaration()
		if err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}

	return statements, nil
}

func (p *Parser) declaration() (expr api.IStatement, err error) {
	if p.match(api.FUN) {
		return p.function("function")
	}

	if p.match(api.VAR) {
		statement, err := p.varDeclaration()
		if err != nil {
			p.synchronize()
			return nil, err
		}
		return statement, err
	}

	return p.statement()
}

func (p *Parser) statement() (expr api.IStatement, err error) {
	if p.match(api.LEFT_BRACE) {
		block, err := p.block()
		if err != nil {
			return nil, err
		}
		return statements.NewBlockStatement(block), nil
	}

	if p.match(api.FOR) {
		return p.forStatement()
	}

	if p.match(api.IF) {
		return p.ifStatement()
	}

	if p.match(api.PRINT) {
		return p.printStatement()
	}

	if p.match(api.RETURN) {
		return p.returnStatement()
	}

	if p.match(api.BREAK) {
		return p.breakStatement()
	}

	if p.match(api.CONTINUE) {
		return p.continueStatement()
	}

	if p.match(api.WHILE) {
		return p.whileStatement()
	}

	return p.expressionStatement()
}

func (p *Parser) function(kind string) (expr api.IStatement, err error) {
	funName, err := p.consume(api.IDENTIFIER, "Expect '"+kind+"' name.")
	if err != nil {
		return nil, err
	}

	_, err = p.consume(api.LEFT_PAREN, "Expect '(' after '"+kind+"' name.")
	if err != nil {
		return nil, err
	}

	parameters := []api.IToken{}

	if !p.check(api.RIGHT_PAREN) {
		for matchComma := true; matchComma; matchComma = p.match(api.COMMA) {
			if len(parameters) >= 255 {
				return nil, errors.New("'" + p.peek().Lexeme() + "' can't more than 255 parameters.")
			}

			parmName, err := p.consume(api.IDENTIFIER, "Expect parameter name.")
			if err != nil {
				return nil, err
			}

			parameters = append(parameters, parmName)
		}
	}

	_, err = p.consume(api.RIGHT_PAREN, "Expect ')' after parameters.")
	if err != nil {
		return nil, err
	}

	// We consume the { at the beginning of the body here before calling
	// block(). Because block() assumes the brace token has already been
	// matched.
	_, err = p.consume(api.LEFT_BRACE, "Expect '{' before '"+kind+"' body.")
	if err != nil {
		return nil, err
	}

	body, err := p.block()
	if err != nil {
		return nil, err
	}

	return statements.NewFunctionStatement(funName, parameters, body), nil
}

func (p *Parser) expression() (expr api.IExpression, err error) {
	return p.assignment()
}

func (p *Parser) assignment() (expr api.IExpression, err error) {
	// parse the left-hand side, which can be any
	// expression of higher precedence
	expr, err = p.or()

	if err != nil {
		return nil, err
	}

	if p.match(api.EQUAL) {
		// parse the right-hand side
		// and then wrap it all up in an assignment expression
		equals := p.previous()
		value, err := p.assignment()

		if err != nil {
			return nil, err
		}

		if expr.Type() == api.VAR_EXPR {
			name := expr.Name()
			return interpreter.NewAssignExpression(name, value), nil
		}

		// TODO create a NewParseError
		return nil, errors.New(equals.String() + " : Invalid assignment target.")
		// return nil, errors.NewRuntimeError(equals, "Invalid assignment target.")
	}

	return expr, nil
}

func (p *Parser) or() (expr api.IExpression, err error) {
	expr, err = p.and()

	if err != nil {
		return nil, err
	}

	for p.match(api.OR) {
		operator := p.previous()

		right, err := p.and()
		if err != nil {
			return nil, err
		}

		expr = interpreter.NewLogicExpression(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) and() (expr api.IExpression, err error) {
	expr, err = p.equality()

	if err != nil {
		return nil, err
	}

	for p.match(api.AND) {
		operator := p.previous()

		right, err := p.equality()
		if err != nil {
			return nil, err
		}

		expr = interpreter.NewLogicExpression(expr, operator, right)
	}

	return expr, nil
}

// The parser has already matched the var token,
// so next it requires and consumes an identifier token for the variable name.
func (p *Parser) varDeclaration() (expr api.IStatement, err error) {
	name, err := p.consume(api.IDENTIFIER, "Expect variable name.")
	if err != nil {
		return nil, err
	}

	var initializer api.IExpression

	if p.match(api.EQUAL) {
		initializer, err = p.expression()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.consume(api.SEMICOLON, "Expect ';' after variable declaration.")
	if err != nil {
		return nil, err
	}

	return statements.NewVarStatement(name, initializer), nil
}

// --------------------------------------------------------
// equality
// --------------------------------------------------------
func (p *Parser) equality() (expr api.IExpression, err error) {
	expr, err = p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(api.BANG_EQUAL, api.EQUAL_EQUAL) {
		operator := p.previous()
		right, errc := p.comparison()
		if errc != nil {
			return nil, errc
		}
		expr = interpreter.NewBinaryExpression(expr, operator, right)
	}

	return expr, nil
}

// This checks to see if the current token has any of the given types.
// If so, it consumes the token and returns true.
// Otherwise, it returns false and leaves the current token alone
func (p *Parser) match(types ...api.TokenType) bool {
	for _, ttype := range types {
		if p.check(ttype) {
			p.advance()
			return true
		}
	}

	return false
}

// returns true if the current token is of the given type
func (p *Parser) check(ttype api.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	// fmt.Println("parser check: ", p.peek().Type(), " -> ", ttype)
	return p.peek().Type() == ttype
}

// consumes the current token and returns it, similar to
// how our scanner’s corresponding method crawled through characters
func (p *Parser) advance() api.IToken {
	if !p.isAtEnd() {
		p.current++
	}

	return p.previous()
}

// checks if we’ve run out of tokens to parse
func (p *Parser) isAtEnd() bool {
	return p.peek().Type() == api.EOF
}

// returns the current token we have yet to consume
func (p *Parser) peek() api.IToken {
	return p.tokens[p.current]
}

// returns the most recently consumed token
func (p *Parser) previous() api.IToken {
	return p.tokens[p.current-1]
}

// --------------------------------------------------------
// comparison
// --------------------------------------------------------
func (p *Parser) comparison() (expr api.IExpression, err error) {
	expr, err = p.term()
	if err != nil {
		return nil, err
	}

	for p.match(api.GREATER, api.GREATER_EQUAL, api.LESS, api.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = interpreter.NewBinaryExpression(expr, operator, right)
	}

	return expr, nil
}

// --------------------------------------------------------
// term
// --------------------------------------------------------
func (p *Parser) term() (expr api.IExpression, err error) {
	expr, err = p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(api.MINUS, api.PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = interpreter.NewBinaryExpression(expr, operator, right)
	}

	return expr, nil
}

// --------------------------------------------------------
// factor
// --------------------------------------------------------
func (p *Parser) factor() (expr api.IExpression, err error) {
	expr, err = p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(api.SLASH, api.STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}

		expr = interpreter.NewBinaryExpression(expr, operator, right)
	}

	return expr, nil
}

// --------------------------------------------------------
// unary
// --------------------------------------------------------
func (p *Parser) unary() (expr api.IExpression, err error) {

	if p.match(api.BANG, api.MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return interpreter.NewUnaryExpression(operator, right), nil
	}

	return p.call()
}

func (p *Parser) call() (expr api.IExpression, err error) {
	// First, we parse a primary expression, the “left operand” to the call.
	expr, err = p.primary()
	if err != nil {
		return nil, err
	}

	for {
		// Each time we see a "("" , we call finishCall() to parse the call expression using the
		// previously parsed expression as the callee
		if p.match(api.LEFT_PAREN) {
			// The returned expression becomes the
			// new expr and we loop to see if the result is itself called.
			expr, err = p.finishCall(expr)
			if err != nil {
				return nil, err
			}
		} else {
			break
		}
	}

	return expr, nil
}
func (p *Parser) finishCall(callee api.IExpression) (expr api.IExpression, err error) {
	arguments := []api.IExpression{}

	// handle the zero-argument case. We check for that case first by
	// seeing if the next token is ")"
	if !p.check(api.RIGHT_PAREN) {
		for matchComma := true; matchComma; matchComma = p.match(api.COMMA) {
			expr, err = p.expression()
			if err != nil {
				return nil, err
			}

			if len(arguments) >= 255 {
				return nil, errors.New("'" + p.peek().Lexeme() + "' can't more than 255 arguments.")
			}
			arguments = append(arguments, expr)
		}
	}

	paren, err := p.consume(api.RIGHT_PAREN, "Expect ')' after arguments.")
	if err != nil {
		return nil, err
	}

	return interpreter.NewCallExpression(callee, paren, arguments), nil
}

func (p *Parser) primary() (expr api.IExpression, err error) {

	if p.match(api.FALSE) {
		return interpreter.NewLiteralExpression(nil, literals.NewBooleanLiteral(false)), nil
	}
	if p.match(api.TRUE) {
		return interpreter.NewLiteralExpression(nil, literals.NewBooleanLiteral(true)), nil
	}
	if p.match(api.NIL) {
		return interpreter.NewLiteralExpression(nil, literals.NewNilLiteral()), nil
	}

	if p.match(api.NUMBER, api.STRING) {
		// NOTE: may need to copy the literal!!!!
		return interpreter.NewLiteralExpression(p.previous(), p.previous().Literal()), nil
	}

	// Parsing a variable expression
	if p.match(api.IDENTIFIER) {
		return interpreter.NewVariableExpression(p.previous()), nil
	}

	if p.match(api.LEFT_PAREN) {
		expr, errc := p.expression()
		if errc != nil {
			return nil, errc
		}
		_, err = p.consume(api.RIGHT_PAREN, "Expect ')' after expression.")
		if err != nil {
			return nil, err
		}
		return interpreter.NewGroupingExpression(expr), nil
	}

	// If none of the cases in there match,
	// it means we are sitting on a token that can’t start an expression.
	return nil, p.lerror(p.previous(), "Expected expression to begin.")
}

// --------------------------------------------------------
// Blocks
// --------------------------------------------------------
func (p *Parser) block() (statements []api.IStatement, err error) {
	statements = make([]api.IStatement, 0)

	for !p.check(api.RIGHT_BRACE) && !p.isAtEnd() {
		decl, err := p.declaration()

		if err != nil {
			return nil, err
		}

		statements = append(statements, decl)
	}

	_, err = p.consume(api.RIGHT_BRACE, "Expect '}' after block.")

	if err != nil {
		return nil, err
	}

	return statements, nil
}

func (p *Parser) consume(ttype api.TokenType, message string) (token api.IToken, err error) {
	if p.check(ttype) {
		return p.advance(), nil
	}

	token = p.peek()
	return token, p.lerror(token, message)
}

func (p *Parser) lerror(ttype api.IToken, message string) error {
	p.assembler.ReportToken(ttype, message)
	return errors.New(message)
}
