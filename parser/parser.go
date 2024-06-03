package parser

import (
	"github.com/firesquid6/do-shell/lexer"
	"github.com/firesquid6/do-shell/token"
	"github.com/firesquid6/do-shell/tree"
)

type Parser struct {
	l        *lexer.Lexer
	token    token.Token
	position int
}

// the lexer should have already tokenized the text
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
  p.position = 0
  p.token = p.l.Tokens[p.position]
	return p
}

func (p *Parser) nextToken() {
  p.position += 1
  p.token = p.l.Tokens[p.position]
}

func (p *Parser) peek() token.Token {
  return p.l.Tokens[p.position + 1]
}

func (p *Parser) peekFor(t token.TokenType) bool {
  return p.peek().Type == t
}

func (p *Parser) ParseProgram() *tree.Program {
  return nil
}
