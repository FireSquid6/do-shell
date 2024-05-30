package lexer

import (
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
		},
		currentState: NORMAL,
		status: &LexerStatus{
			Position: 0,
		},
	}
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
    case 10:  // linebreak
      ls.AddToken(newToken(token.LINEBREAK, ls.Ch))
    case 0:
      ls.AddToken(newToken(token.EOF, ls.Ch))
    case '"':
      ls.CurrentState = STRING
      return
    case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
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


type CommandState struct{}
func (s *CommandState) Process(ls *LexerStatus) {
  if ls.Ch != '[' {
    panic("Tried to parse command state but didn't start with a '[")
  }
  
  ls.Advance()
  start := ls.Position
  for ls.Ch != ']' || ls.Ch == 0 {
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
	for l.Ch == ' ' || l.Ch == '\t' || l.Ch == '\r' {
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
