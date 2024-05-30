package lexer

import (
	"fmt"
	"github.com/firesquid6/do-shell/token"
)

type StateName int

const (
	IDENTIFIER StateName = iota
	NUMBER
	STRING
	COMMAND
)

func Tokenize(text string) []token.Token {
	tokens := []token.Token{}

	return tokens
}

type Lexer struct {
	states       map[StateName]LexerState
	currentState StateName
	status       LexerStatus
}

func (l *Lexer) Process() {
	state, ok := l.states[l.currentState]
	if !ok {
		panic("State not found. Firesquid screwed up programming real bad.")
	}

	newStatus := state.Process(&l.status)

	if newStatus.position == l.status.position {
		panic("Lexer did not advance position. In state " + fmt.Sprint(l.currentState))
	}

}

type LexerStatus struct {
	position     int
	ch           rune
	source       []rune
	tokens       []token.Token
	currentState StateName
}

type ProcessResult struct {
  currentState StateName
  position int
}

type LexerState interface {
	Process(ls *LexerStatus) ProcessResult
}

//	func New(input string) *Lexer {
//		if len(input) == 0 {
//			return nil
//		}
//
//		fmt.Println(input)
//		l := &Lexer{input: []rune(input)}
//
//		l.position = 0
//		l.ch = l.input[l.position]
//
//		return l
//	}
//
//	func (l *Lexer) advance() {
//		l.position += 1
//
//		if l.position >= len(l.input) {
//			l.ch = 0
//			l.position -= 1
//			return
//		}
//
//		l.ch = l.input[l.position]
//
//		fmt.Println("Advance: ", string(l.ch), " ", l.ch)
//	}
//
//	func (l *Lexer) eatWhitespace() {
//		for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
//			l.advance()
//		}
//	}
//
//	func (l *Lexer) readIdentifier() []rune {
//		position := l.position
//
//		for isLetter(l.ch) {
//			l.advance()
//		}
//
//		return l.input[position:l.position]
//	}
//
//	func (l *Lexer) NextToken() token.Token {
//		var t token.Token
//
//		fmt.Println("currently reading: ", string(l.ch), " ", l.ch)
//
//		switch l.ch {
//		// simple single character tokens
//		case '=':
//			t = newToken(token.ASSIGN, l.ch)
//		case 10:
//			fmt.Println("Linebreak")
//			t = newToken(token.LINEBREAK, l.ch)
//		case '(':
//			t = newToken(token.LPAREN, l.ch)
//		case ')':
//			t = newToken(token.RPAREN, l.ch)
//		case '{':
//			t = newToken(token.LBRACE, l.ch)
//		case '}':
//			t = newToken(token.RBRACE, l.ch)
//		case '[':
//			t = newToken(token.LBRACKET, l.ch)
//		case ']':
//			t = newToken(token.RBRACKET, l.ch)
//		case '+':
//			t = newToken(token.PLUS, l.ch)
//		case 0:
//			t.Literal = []rune{0}
//			t.Type = token.EOF
//		default:
//			if isLetter(l.ch) {
//				t.Literal = l.readIdentifier()
//				t.Type = token.LookupIdentifier(t.Literal)
//				return t
//			} else if isDigit(l.ch) {
//				t = l.readNumber()
//			} else {
//				t = newToken(token.ILLEGAL, l.ch)
//			}
//		}
//
//		l.advance()
//		l.eatWhitespace()
//		fmt.Println("NextToken: ", token.ReadableTokenName(t), " ", string(t.Literal))
//		return t
//	}
//
//	func (l *Lexer) readNumber() token.Token {
//		position := l.position
//		t := token.Token{Type: token.INT, Literal: []rune{}}
//
//		for isDigit(l.ch) {
//			l.advance()
//		}
//
//		t.Literal = l.input[position:l.position]
//		return t
//	}
func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: []rune{ch}}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
