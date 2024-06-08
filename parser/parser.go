package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/firesquid6/do-shell/token"
	"github.com/firesquid6/do-shell/tree"
)

type (
	prefixParseFn func() (tree.Expression, error)
	infixParseFn  func(tree.Expression) (tree.Expression, error)
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

var precedences = map[token.TokenType]int{
	token.EQUAL:              EQUALS,
	token.NOT_EQUAL:          EQUALS,
	token.LESS_THAN:          LESSGREATER,
	token.GREATER_THAN:       LESSGREATER,
	token.LESS_THAN_EQUAL:    LESSGREATER,
	token.GREATER_THAN_EQUAL: LESSGREATER,
	token.PLUS:               SUM,
	token.MINUS:              SUM,
	token.MULTIPLY:           PRODUCT,
	token.DIVIDE:             PRODUCT,
}

type Parser struct {
	tokens   []token.Token
	token    token.Token
	position int
	Program  *tree.Program
	Errors   []error

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(tokens []token.Token) *Parser {
	p := &Parser{tokens: tokens}
	p.position = 0

  if len(p.tokens) > 0 {
    p.token = p.tokens[p.position]
  }

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(token.NUMBER, p.parseNumberLiteral)
	p.registerPrefix(token.NOT, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.DIVIDE, p.parseInfixExpression)
	p.registerInfix(token.MULTIPLY, p.parseInfixExpression)
	p.registerInfix(token.EQUAL, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQUAL, p.parseInfixExpression)
	p.registerInfix(token.LESS_THAN, p.parseInfixExpression)
	p.registerInfix(token.LESS_THAN_EQUAL, p.parseInfixExpression)
	p.registerInfix(token.GREATER_THAN, p.parseInfixExpression)
	p.registerInfix(token.GREATER_THAN_EQUAL, p.parseInfixExpression)

	return p
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseGroupedExpression() (tree.Expression, error) {
	p.nextToken()

	expression, err := p.parseExpression(LOWEST)
  if err != nil {
    return nil, err
  }

	if !p.peekFor(token.RPAREN) {
		// this happens when an expression isn't closed properly
    return nil, errors.New("Expression was not closed properly")
    
	}

	return expression, nil
}

func (p *Parser) parseBoolean() tree.Expression {
	return &tree.Boolean{Token: p.token, Value: p.token.Type == token.TRUE}
}

func (p *Parser) parseInfixExpression(left tree.Expression) (tree.Expression, error) {
	expression := &tree.InfixExpression{
		Token:    p.token,
		Operator: p.token.Literal,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	p.nextToken()
  exp, err := p.parseExpression(precedence)
  if err != nil {
    return nil, errors.Join(errors.New("Error parsing infix expression:"), err)
  }


  expression.Right = exp

	return expression, nil
}

func (p *Parser) peekPrecedence() int {
	precedence, ok := precedences[p.peek().Type]
	if !ok {
		return LOWEST
	}
	return precedence
}

func (p *Parser) currentPrecedence() int {
	precedence, ok := precedences[p.token.Type]
	if !ok {
		return LOWEST
	}
	return precedence
}

func (p *Parser) parsePrefixExpression() (tree.Expression, error) {
	expression := &tree.PrefixExpression{
		Token:    p.token,
		Operator: p.token.Literal,
	}
	p.nextToken()
  exp, err := p.parseExpression(PREFIX)

  if err != nil {
    return nil, errors.Join(errors.New("error parsing prefix expression"), err)

  }
  expression.Right = exp

	return expression, nil
}

func (p *Parser) parseIdentifier() (tree.Expression, error) {
	return &tree.Identifier{Token: p.token, Value: p.token.Literal}, nil
}

func (p *Parser) parseNumberLiteral() (tree.Expression, error) {
	// TODO: handle it being a float and not an integer
	literal := &tree.IntegerLiteral{Token: p.token}
	value, err := strconv.ParseInt(string(p.token.Literal), 0, 64)

	if err != nil {
    return nil, errors.Join(errors.New("Error parsing number literal:"), err) 
	}

	literal.Value = value
	return literal, nil
}

func (p *Parser) nextToken() {
	p.position += 1

  if p.position >= len(p.tokens) {
    p.token = token.Token{Type: token.EOF, Literal: []rune("")}
    return
  }
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
	if len(p.tokens) < 1 {
		return &tree.Program{}
	}

	program := &tree.Program{}
	statements := []tree.Statement{}

	for p.token.Type != token.EOF {
		statement, err := p.parseStatement()
    if err != nil {
      p.Errors = append(p.Errors, err)
    }
    if statement != nil {
      statements = append(statements, statement)
    }

		p.nextToken()
	}

	program.Statements = statements
	p.Program = program
	return program
}

func (p *Parser) parseStatement() (tree.Statement, error) {
  var statement tree.Statement
  var err error

	switch p.token.Type {
	case token.LET:
    statement, err = p.parseLetStatement()
	case token.RETURN:
		statement = p.parseReturnStatement()
	default:
		statement, err = p.parseExpressionStatement()
	}

  return statement, err
}

func (p *Parser) parseLetStatement() (*tree.LetStatement, error) {
	statement := &tree.LetStatement{Token: p.token}

	if !p.peekFor(token.IDENTIFIER) {
    return nil, errors.New("let statement is not followed by identifier")
	}
	p.nextToken()

	statement.Name = &tree.Identifier{Token: p.token, Value: p.token.Literal}

	if !p.peekFor(token.ASSIGN) {
		return nil, errors.New("identifier in let statement is not followed by assignment")
	}
	p.nextToken()

	// TODO: parse the expression

	for p.token.Type != token.SEMICOLON {
		p.nextToken()
	}

	return statement, nil
}

func (p *Parser) parseReturnStatement() *tree.ReturnStatement {
	statement := &tree.ReturnStatement{Token: p.token}

	p.nextToken()

  // TODO: parse expression

	// it will need to be parsed correctly
	for p.token.Type != token.SEMICOLON {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpressionStatement() (*tree.ExpressionStatement, error) {
	statement := &tree.ExpressionStatement{Token: p.token}

  exp, err := p.parseExpression(LOWEST)
  statement.Expression = exp

  if err != nil {
    return nil, errors.Join(errors.New("Error parsing expression statement: "), err)
  }

	if p.peekFor(token.SEMICOLON) {
		p.nextToken()
	}

	return statement, nil
}

func (p *Parser) parseExpression(precedence int) (tree.Expression, error) {
	prefix := p.prefixParseFns[p.token.Type]
	if prefix == nil {
		return nil, errors.New(fmt.Sprintf("Failed to find a prefix parse function for %s", token.ReadableTokenName(p.token)))
	}
	leftExp := prefix()

	for !p.peekFor(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peek().Type]
		if infix == nil {
			return leftExp, nil
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp, nil
}
