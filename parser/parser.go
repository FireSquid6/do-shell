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
	return nil
}
