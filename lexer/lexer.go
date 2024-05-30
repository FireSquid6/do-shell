package lexer

import (
	"fmt"

	"github.com/firesquid6/do-shell/token"
)

type StateName int

const (
	IDENTIFIER StateName = iota
	NORMAL
	NUMBER
	STRING
	COMMAND
)

func Tokenize(text string) []token.Token {
	l := newLexer()
	return l.LexText(text)
}

type Lexer struct {
	states       map[StateName]LexerState
	currentState StateName
	status       *LexerStatus
}

func newLexer() Lexer {
	return Lexer{
		states: map[StateName]LexerState{
			NORMAL: &NormalState{},
      COMMAND: &CommandState{},
      NUMBER: &NumberState{},
      STRING: &StringState{},
		},
		currentState: NORMAL,
		status: &LexerStatus{
			Position: 0,
		},
	}
}

type StringState struct{}
func (s *StringState) Process(ls *LexerStatus) {
  if ls.Ch != '"' {
    panic("Tried to parse string state but didn't start with a '\"")
  }

  ls.Advance()
  start := ls.Position
  for ls.Ch != '"' || ls.Ch == 0 {
    ls.Advance()
  }

  if ls.Ch == 0 {
    ls.AddToken(token.Token{Type: token.ILLEGAL, Literal: ls.Source[start:ls.Position]})
    return
  }

  ls.AddToken(token.Token{Type: token.STRING, Literal: ls.Source[start:ls.Position]})
}


type NumberState struct{}
func (s *NumberState) Process(ls *LexerStatus) {
  start := ls.Position

  for isDigit(ls.Ch) || ls.Ch == '.' {
    ls.Advance()
  }

  literal := ls.Source[start:ls.Position]
  var seenDecimal bool = false

  // ensures that there is only one decimal point and it is at an appropriate place
  for i, ch := range literal {
    if ch == '.' {
      if seenDecimal || i == 0 || i == len(literal) - 1 {
        ls.AddToken(token.Token{Type: token.ILLEGAL, Literal: literal})
        return
      }
      seenDecimal = true
    }
  }

  ls.AddToken(token.Token{Type: token.NUMBER, Literal: literal})
}

type NormalState struct{}

func (s *NormalState) Process(ls *LexerStatus) {
  for {
    ls.EatWhitespace()
    
    switch ls.Ch {
    case '+':
      ls.AddToken(newToken(token.PLUS, ls.Ch))
    case '=':
      // TODO: peek
      ls.AddToken(newToken(token.ASSIGN, ls.Ch))
    case '(':
      ls.AddToken(newToken(token.LPAREN, ls.Ch))
    case ')':
      ls.AddToken(newToken(token.RPAREN, ls.Ch))
    case '{':
      ls.AddToken(newToken(token.LBRACE, ls.Ch))
    case '}':
      ls.AddToken(newToken(token.RBRACE, ls.Ch))
    case '[':
      ls.CurrentState = COMMAND
      return
    case ',':
      ls.AddToken(newToken(token.COMMA, ls.Ch))
    case ';':  // lines end in semicolons
      ls.AddToken(newToken(token.SEMICOLON, ls.Ch))
    case 0:
      ls.AddToken(newToken(token.EOF, ls.Ch))
    case '"':
      ls.CurrentState = STRING
      return
    case '.':
      ls.AddToken(newToken(token.DOT, ls.Ch))
    case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
      fmt.Println("Moving to number state")
      ls.CurrentState = NUMBER
      return
    default:
      if isLetter(ls.Ch) {
        ls.CurrentState = IDENTIFIER
        return
      }
      
      ls.AddToken(newToken(token.ILLEGAL, ls.Ch))
    }

    ls.Advance()
  }

}

type IdentifierState struct{}
func (s *IdentifierState) Process(ls *LexerStatus) {
  start := ls.Position
  for isLetter(ls.Ch) {
    ls.Advance()
  }

  literal := ls.Source[start:ls.Position]
  tokenType := token.LookupIdentifier(literal)
  ls.AddToken(token.Token{Type: tokenType, Literal: literal})
}

type CommandState struct{}
func (s *CommandState) Process(ls *LexerStatus) {
  if ls.Ch != '[' {
    panic("Tried to parse command state but didn't start with a '[")
  }
  
  ls.Advance()
  start := ls.Position
  for ls.Ch != ']' || ls.Ch == 0 {
    // TODO: handle escape characters
    // TODO: handle if user
    ls.Advance()
  }

  ls.AddToken(token.Token{Type: token.COMMAND, Literal: ls.Source[start:ls.Position]})
}

func (l *Lexer) LexText(text string) []token.Token {
	l.status.Position = 0
	l.status.Source = []rune(text)
	l.status.Ch = l.status.Source[l.status.Position]

	for l.status.Position < len(l.status.Source) {
		l.Process()
	}

	return l.status.Tokens
}

func (l *Lexer) Process() {
	state, ok := l.states[l.currentState]
	if !ok {
		panic("State not found. Firesquid screwed up programming real bad.")
	}

  fmt.Println("Processing state", l.currentState)
	state.Process(l.status)
	// todo: ensure that something changed so that we don't get stuck in an infinite loop
	l.status.EatWhitespace()
}

type LexerStatus struct {
	Position     int
	Ch           rune
	Source       []rune
	Tokens       []token.Token
	CurrentState StateName
}

func (l *LexerStatus) Advance() {
	l.Position += 1

	if l.Position >= len(l.Source) {
		l.Ch = 0
		l.Position -= 1
		return
	}

	l.Ch = l.Source[l.Position]
}

func (l *LexerStatus) EatWhitespace() {
	for l.Ch == ' ' || l.Ch == '\t' || l.Ch == '\r' || l.Ch == '\n' {
		l.Advance()
	}
}

func (l *LexerStatus) PeekFor(ch rune) bool {
	if l.Position+1 >= len(l.Source) {
		return false
	}

	if l.Source[l.Position+1] == ch {
		l.Advance()
		return true
	}

	return false
}

func (l *LexerStatus) AddToken(t token.Token) {
	l.Tokens = append(l.Tokens, t)
}

type LexerState interface {
	Process(ls *LexerStatus)
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
