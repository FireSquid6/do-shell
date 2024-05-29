package lexer

import (
	"github.com/firesquid6/do-shell/token"
)

type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
	state        LexerState
}

type LexerState int

const (
  SYMBOLS LexerState = iota
  STRING
  IDENTIFIER
  COMMAND
)

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) eatWhitespace() {
  for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
    l.readChar()
  }
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	switch l.ch {
	case '=':
		t = newToken(token.ASSIGN, l.ch)
	case '\n':
		// todo: make \n work as a linened character
		t = newToken(token.LINEBREAK, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case 0:
		t.Literal = []rune{0}
		t.Type = token.EOF
  default:
    t.Literal = []rune{}
    t.Type = token.ILLEGAL
	}

	l.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: []rune{ch}}
}
