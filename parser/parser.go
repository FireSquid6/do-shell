package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/firesquid6/do-shell/token"
	"github.com/firesquid6/do-shell/tree"
	// "github.com/firesquid6/do-shell/option"
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
	// p.registerPrefix(token.TRUE, p.parseBoolean)
	// p.registerPrefix(token.FALSE, p.parseBoolean)

	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)

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

func (p *Parser) parseIfExpression() (tree.Expression, error) {
	expression := &tree.IfExpression{Token: p.token}

	if !p.peekFor(token.LPAREN) {
		return nil, errors.New("If statement is not followed by a left parenthesis")
	}

	p.nextToken()
	p.nextToken()

	cond, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, errors.Join(errors.New("Error parsing condition in if statement:"), err)
	}

	expression.Condition = cond

	if !p.peekFor(token.LBRACE) {
		return nil, errors.New("If statement condition is not followed by a block statement")
	}
	p.nextToken()
	p.nextToken()

	consequence, err := p.parseBlockStatement()
	if err != nil {
		return nil, errors.Join(errors.New("Error parsing consequence in if statement:"), err)
	}
	expression.Consequence = consequence

	return expression, nil
}

func (p *Parser) parseBlockStatement() (*tree.BlockStatement, error) {
	block := &tree.BlockStatement{Token: p.token}
	block.Statements = []tree.Statement{}

	for p.token.Type != token.RBRACE && p.token.Type != token.EOF {
		fmt.Println("Parsing block statement: ", token.ReadableTokenName(p.token))
		statement, err := p.parseStatement()
		p.nextToken()

		if err != nil {
			return nil, errors.Join(errors.New("Error parsing statement in block statement:"), err)
		}
		block.Statements = append(block.Statements, statement)
	}
  p.previousToken()

	fmt.Println("done parsing block statement: ", token.ReadableTokenName(p.token))

	return block, nil
}

func (p *Parser) parseFunctionLiteral() (tree.Expression, error) {
	literal := &tree.FunctionLiteral{Token: p.token}

	if !p.peekFor(token.LPAREN) {
		return nil, errors.New("Function literal is not followed by a left parenthesis")
	}
	p.nextToken()

	params, err := p.parseFunctionParameters()

	if err != nil {
		return nil, errors.Join(errors.New("Error parsing function parameters:"), err)
	}
	literal.Parameters = params

	fmt.Println("Function literal parsing block statement: ", token.ReadableTokenName(p.token))

	if p.token.Type != token.LBRACE {
		return nil, errors.New("Function literal parameters are not followed by a block statement")
	}
	p.nextToken()

	statements, err := p.parseBlockStatement()
	if err != nil {
		return nil, errors.Join(errors.New("Error parsing statements in function literal:"), err)
	}

	literal.Statements = statements
	return literal, nil
}

func (p *Parser) parseFunctionParameters() ([]*tree.Identifier, error) {
	identifiers := []*tree.Identifier{}

	if p.token.Type != token.LPAREN {
		return nil, errors.New("Function parameters do not start with a left parenthesis")
	}

	p.nextToken()
	if p.token.Type == token.RPAREN {
		return identifiers, nil
	}

	// scary! this could infinite loop
	// todo: fix this crap
	for {
		if p.token.Type != token.IDENTIFIER {
			return nil, errors.New("Function parameters are not identifiers")
		}

		identifiers = append(identifiers, &tree.Identifier{Token: p.token, Value: p.token.Literal})
		p.nextToken()
		fmt.Println("Parsing function parameters 2: ", token.ReadableTokenName(p.token))

		if p.token.Type == token.RPAREN || p.token.Type == token.EOF {
			p.nextToken()
			break
		}

		if p.token.Type != token.COMMA {
			return nil, errors.New("Function parameters are not separated by commas")
		}
		p.nextToken()
	}

	return identifiers, nil
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
  fmt.Println("!!!! Parsing identifier: ", token.ReadableTokenName(p.token), p.position)
  if p.peekFor(token.LPAREN) {
    fmt.Println("Parsing call expression: ", token.ReadableTokenName(p.token), p.position)
    p.nextToken()
    return p.parseCallExpression()
  }
	return &tree.Identifier{Token: p.token, Value: p.token.Literal}, nil
}

func (p *Parser) parseCallExpression() (tree.Expression, error) {
  p.nextToken()
  fmt.Println("??? Parsing call expression: ", token.ReadableTokenName(p.token), p.position)

  callExpression := &tree.CallExpression{Token: p.token}
  callExpression.Function = tree.Identifier{Token: p.token, Value: p.token.Literal}
  callExpression.Arguments = []tree.Expression{}

  for p.token.Type != token.RPAREN {
    expression, err := p.parseExpression(LOWEST)

    if err != nil {
      return nil, errors.Join(errors.New("Error parsing call expression:"), err)
    }

    callExpression.Arguments = append(callExpression.Arguments, expression)
    p.nextToken()
    p.nextToken()
    fmt.Println("Just parsed an expression in call expression: ", token.ReadableTokenName(p.token), p.position)
  }

  return callExpression, nil
} 

func (p *Parser) parseNumberLiteral() (tree.Expression, error) {
	literal := &tree.NumberLiteral{Token: p.token}
	value, err := strconv.ParseFloat(string(p.token.Literal), 64)

	if err != nil {
		return nil, errors.Join(errors.New("Error parsing number literal:"), err)
	}

	literal.Value = value
	return literal, nil
}

func (p *Parser) previousToken() {
  if p.position > 0 {
    p.position -= 1
    p.token = p.tokens[p.position]
  }
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
    fmt.Println("Just finished a statement. Now looking at: ", token.ReadableTokenName(p.token), p.position)

		p.nextToken()
    if p.token.Type == token.EOF {
      break
    }
	}

	program.Statements = statements
	p.Program = program
	return program
}

func (p *Parser) parseStatement() (tree.Statement, error) {
  fmt.Println("Parsing statement: ", token.ReadableTokenName(p.token), p.position)
	var statement tree.Statement
	var err error

	switch p.token.Type {
	case token.LET:
		statement, err = p.parseLetStatement()
    fmt.Println("Done parsing let statement and now in the parseStatement world: ", token.ReadableTokenName(p.token), p.position)
	case token.RETURN:
		statement, err = p.parseReturnStatement()
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

	// we have to call nextToken twice to make sure we aren't looking at the assignment token
	// this smells like bad programming
	p.nextToken()
	expression, err := p.parseExpression(LOWEST)

	if err != nil {
		return nil, errors.Join(errors.New("Error parsing expression in let statement:"), err)
	}
	statement.Expression = expression

  p.nextToken()
	// if there's errors it also could be here that I need to call nextToken
  fmt.Println("Done parsing let statement: ", token.ReadableTokenName(p.token), p.position)

	return statement, nil
}

func (p *Parser) parseReturnStatement() (*tree.ReturnStatement, error) {
	statement := &tree.ReturnStatement{Token: p.token}

	p.nextToken()

	// TODO: parse expression
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, errors.Join(errors.New("Error parsing expression in return statement:"), err)
	}

  p.nextToken()

	statement.Expression = expression

	return statement, nil
}

func (p *Parser) parseExpressionStatement() (*tree.ExpressionStatement, error) {
	fmt.Println("Parsing expression statement: ", token.ReadableTokenName(p.token), p.position)
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
  fmt.Println("Parsing expression: ", token.ReadableTokenName(p.token), p.position)
	prefix, ok := p.prefixParseFns[p.token.Type]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Failed to find a prefix parse function for %s", token.ReadableTokenName(p.token)))
	}
	leftExp, err := prefix()

	if err != nil {
		return nil, errors.Join(errors.New("Error parsing expression:"), err)
	}

	for !(p.peekFor(token.SEMICOLON) || p.peekFor(token.COMMA)) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peek().Type]
		if infix == nil {
			return leftExp, nil
		}

		p.nextToken()
		leftExp, err = infix(leftExp)

		if err != nil {
			return nil, errors.Join(errors.New("Error parsing infix expression:"), err)
		}
	}

	// note to future me: this is probably where the issue is
	// you're not going to the next token and are left looking at the token right before the semicolon

	return leftExp, nil
}
