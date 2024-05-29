package lexer

import (
	"fmt"
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

func Tokenize(text string) []token.Token {
  fmt.Println(text)
	l := New(text)
	tokens := []token.Token{}

	for {
		t := l.NextToken()
		tokens = append(tokens, t)

		if t.Type == token.EOF {
			break
		}
	}

	return tokens

}

func New(input string) *Lexer {
  fmt.Println(input)
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
    fmt.Println("readChar: ", l.ch, " ", string(l.ch))
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() []rune {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

  l.eatWhitespace()

	switch l.ch {
  // simple single character tokens
	case '=':
		t = newToken(token.ASSIGN, l.ch)
	case  10:
    fmt.Println("Linebreak")
		t = newToken(token.LINEBREAK, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
  case '[':
    t = newToken(token.LBRACKET, l.ch)
  case ']':
    t = newToken(token.RBRACKET, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case 0:
		t.Literal = []rune{0}
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
      t.Type = token.LookupIdentifier(t.Literal)
			return t
    } else if isDigit(l.ch) {
      t = l.readNumber()
		} else {
      t = newToken(token.ILLEGAL, l.ch)
    }
	}

	l.readChar()
  fmt.Println("NextToken: ", token.ReadableTokenName(t), " ", string(t.Literal))
	return t
}

func (l *Lexer) readNumber() token.Token {
  position := l.position
  t := token.Token{Type: token.INT, Literal: []rune{}}

  for isDigit(l.ch) {
    l.readChar()
  }


  t.Literal = l.input[position:l.position]
  return t
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: []rune{ch}}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
  return '0' <= ch && ch <= '9'
}
