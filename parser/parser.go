package parser

import (
	"github.com/firesquid6/do-shell/token"
	"github.com/firesquid6/do-shell/tree"
)

type Parser struct {
	tokens   []token.Token
	token    token.Token
	position int
	Tree     *tree.Program
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
	return p
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
  // TODO
  return nil
}

func (p *Parser) parseExpression() tree.Expression {
  // TODO
  return nil
}
