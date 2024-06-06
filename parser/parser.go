package parser

import (
	"github.com/firesquid6/do-shell/token"
	"github.com/firesquid6/do-shell/tree"
	"strconv"
)

type (
	prefixParseFn func() tree.Expression
	infixParseFn  func(tree.Expression) tree.Expression
)

// order of operations
const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type Parser struct {
	tokens   []token.Token
	token    token.Token
	position int
	Tree     *tree.Program
	Errors   []error

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func Parse(tokens []token.Token) *tree.Program {
	p := New(tokens)
	return p.ParseProgram()

}

// the lexer should have already tokenized the text
func New(tokens []token.Token) *Parser {
	p := &Parser{tokens: tokens}
	p.position = 0
	p.token = p.tokens[p.position]

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(token.NUMBER, p.parseNumberLiteral)
  p.registerPrefix(token.NOT, p.parsePrefixExpression)
  p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	return p
}

func (p *Parser) parsePrefixExpression() tree.Expression {
  expression := &tree.PrefixExpression{
    Token: p.token,
    Operator: p.token.Literal,
  }
  p.nextToken()
  expression.Right = p.parseExpression(PREFIX)

  return expression
}

func (p *Parser) parseIdentifier() tree.Expression {
	return &tree.Identifier{Token: p.token, Value: p.token.Literal}
}

func (p *Parser) parseNumberLiteral() tree.Expression {
	// TODO: handle it being a float and not an integer
	literal := &tree.IntegerLiteral{Token: p.token}
	value, err := strconv.ParseInt(string(p.token.Literal), 0, 64)

	if err != nil {
		// todo: handle this error
	}

	literal.Value = value
	return literal
}

func (p *Parser) nextToken() {
	p.position += 1
	p.token = p.tokens[p.position]
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.position+1]
}

func (p *Parser) peekFar(i int) token.Token {
	return p.tokens[p.position+i]
}

func (p *Parser) peekFor(t token.TokenType) bool {
	return p.peek().Type == t
}

func (p *Parser) ParseProgram() *tree.Program {
	program := &tree.Program{}
	statements := []tree.Statement{}

	for p.token.Type != token.EOF {
		statement := p.parseStatement()
		p.nextToken()
		statements = append(statements, statement)
	}

	program.Statements = statements
	return program
}

func (p *Parser) parseStatement() tree.Statement {
	switch p.token.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *tree.LetStatement {
	statement := &tree.LetStatement{Token: p.token}

	if !p.peekFor(token.IDENTIFIER) {
		// TODO: handle these errors
		return nil
	}
	p.nextToken()

	statement.Name = &tree.Identifier{Token: p.token, Value: p.token.Literal}

	if !p.peekFor(token.ASSIGN) {
		// TODO: make an error handler
		return nil
	}
	p.nextToken()

	// TODO: parse the expression
	// TODO: handle hitting EOF
	for p.token.Type != token.SEMICOLON {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseReturnStatement() *tree.ReturnStatement {
	statement := &tree.ReturnStatement{Token: p.token}

	p.nextToken()

	// again, we're skipping the expression for now
	// it will need to be parsed correctly
	for p.token.Type != token.SEMICOLON {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpressionStatement() *tree.ExpressionStatement {
	statement := &tree.ExpressionStatement{Token: p.token}

	statement.Expression = p.parseExpression(LOWEST)

	if p.peekFor(token.SEMICOLON) {
		p.nextToken()
	}

	return nil
}

func (p *Parser) parseExpression(precedence int) tree.Expression {
	prefix := p.prefixParseFns[p.token.Type]
	if prefix == nil {
		// TODO: throw an error
		return nil
	}
	leftExp := prefix()

	return leftExp
}
