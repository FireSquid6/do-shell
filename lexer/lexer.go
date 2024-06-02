package lexer

import (
	"fmt"

	"github.com/firesquid6/do-shell/token"
)

func Tokenize(text string) []token.Token {
  fmt.Println("Starting new tokenization ------------")
	l := Lexer{}
	l.LexText(text)

	return l.Tokens
}

type Lexer struct {
	Position int
	Ch       rune
	Source   []rune
	Tokens   []token.Token
}

func (l *Lexer) AddToken(t token.Token) {
	l.Tokens = append(l.Tokens, t)
  fmt.Println("Added token: ", token.ReadableTokenName(t))
}

func (l *Lexer) Advance() {
	l.Position += 1

	if l.Position >= len(l.Source) {
		l.Ch = 0
		l.Position -= 1
		return
	}

	l.Ch = l.Source[l.Position]
  fmt.Println("Now looking at: ", string(l.Ch))
}

func (l *Lexer) EatWhitespace() {
	for l.Ch == ' ' || l.Ch == '\t' || l.Ch == '\r' || l.Ch == '\n' {
		l.Advance()
	}
}

func (l *Lexer) PeekFor(ch rune) bool {
	if l.Position+1 >= len(l.Source) {
		return false
	}

	if l.Source[l.Position+1] == ch {
		l.Advance()
		return true
	}

	return false
}

func (l *Lexer) PeekForNumber() bool {
  if l.Position+1 >= len(l.Source) {
    return false
  }

  if isDigit(l.Source[l.Position+1]) {
    l.Advance()
    return true
  }

  return false
}

func (l *Lexer) PeekForLetter() bool {
  if l.Position+1 >= len(l.Source) {
    return false
  }

  if isLetter(l.Source[l.Position+1]) {
    l.Advance()
    return true
  }

  return false
}


func (l *Lexer) SkipComment() {
  for l.Ch != '\n' {
    l.Advance()
  }
}

func (l *Lexer) ProcessString() {
	if l.Ch != '"' {
		panic("Tried to parse string state but didn't start with a '\"")
	}

	l.Advance()
	start := l.Position
	for l.Ch != '"' || l.Ch == 0 {
		l.Advance()
	}

	if l.Ch == 0 {
		l.AddToken(token.Token{Type: token.ILLEGAL, Literal: l.Source[start:l.Position]})
		return
	}

	l.AddToken(token.Token{Type: token.STRING, Literal: l.Source[start:l.Position]})
}

func (l *Lexer) ProcessNumber() {
	start := l.Position

	for l.PeekForNumber() {}

	literal := l.Source[start:l.Position+1]
	var seenDecimal bool = false

	// ensures that there is only one decimal point and it is at an appropriate place
	for i, ch := range literal {
		if ch == '.' {
			if seenDecimal || i == 0 || i == len(literal)-1 {
				l.AddToken(token.Token{Type: token.ILLEGAL, Literal: literal})
				return
			}
			seenDecimal = true
		}
	}


  fmt.Println("Adding number: ", string(literal))

	l.AddToken(token.Token{Type: token.NUMBER, Literal: literal})
}

func (l *Lexer) Process() {
	finished := false
	for !finished {
		l.EatWhitespace()
    fmt.Println("Processing: ", string(l.Ch))

		switch l.Ch {
		case '+':
			l.AddToken(newToken(token.PLUS, l.Ch))
		case '=':
      if l.PeekFor('=') {
        l.AddToken(token.Token{Type: token.EQUAL, Literal: []rune{'=', '='} })
      } else {
        l.AddToken(newToken(token.ASSIGN, l.Ch))
      }
    case '<':
      if l.PeekFor('=') {
        l.AddToken(token.Token{Type: token.LESS_THAN_EQUAL, Literal: []rune{'<', '='} })
      } else {
        l.AddToken(newToken(token.LESS_THAN, l.Ch))
      }
    case '>':
      if l.PeekFor('=') {
        l.AddToken(token.Token{Type: token.LESS_THAN_EQUAL, Literal: []rune{'>', '='} })
      } else {
        l.AddToken(newToken(token.GREATER_THAN, l.Ch))
      }
    case '!':
      if l.PeekFor('=') {
        l.AddToken(token.Token{Type: token.LESS_THAN_EQUAL, Literal: []rune{'!', '='} })
      } else {
        l.AddToken(newToken(token.NOT, l.Ch))
      }
		case '(':
			l.AddToken(newToken(token.LPAREN, l.Ch))
		case ')':
			l.AddToken(newToken(token.RPAREN, l.Ch))
		case '{':
			l.AddToken(newToken(token.LBRACE, l.Ch))
		case '}':
			l.AddToken(newToken(token.RBRACE, l.Ch))
		case '`':
			l.ProcessCommand()
		case ',':
			l.AddToken(newToken(token.COMMA, l.Ch))
		case ';': // lines end in semicolons
			l.AddToken(newToken(token.SEMICOLON, l.Ch))
		case 0:
			l.AddToken(newToken(token.EOF, l.Ch))
			finished = true
		case '"':
			l.ProcessString()
    case '#':
      l.SkipComment()
    case '*':
      l.AddToken(newToken(token.MULTIPLY, l.Ch))
    case '%':
      l.AddToken(newToken(token.MOD, l.Ch))
    case '-':
      l.AddToken(newToken(token.MINUS, l.Ch))
		case '.':
			l.AddToken(newToken(token.DOT, l.Ch))
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.ProcessNumber()
		default:
			if isLetter(l.Ch) {
				l.ProcessIdentifier()
			} else {
        l.AddToken(newToken(token.ILLEGAL, l.Ch))
      }
		}

    // when this happens after a processing of a number or letter, it will skip a character that needs to be read
    // to be fixed
		l.Advance()
	}

}

func (l *Lexer) ProcessIdentifier() {
	start := l.Position

  for l.PeekForLetter() {}

	literal := l.Source[start:l.Position+1]
	tokenType := token.LookupIdentifier(literal)

	l.AddToken(token.Token{Type: tokenType, Literal: literal})
}

func (l *Lexer) ProcessCommand() {
	if l.Ch != '`' {
		panic("Tried to parse command state but didn't start with a '`")
	}

	l.Advance()
	start := l.Position
	for l.Ch != '`' || l.Ch == 0 {
		l.Advance()
	}

	l.AddToken(token.Token{Type: token.COMMAND, Literal: l.Source[start:l.Position]})
}

func (l *Lexer) LexText(text string) []token.Token {
	l.Position = 0
	l.Source = []rune(text)
	l.Ch = l.Source[l.Position]

	l.Process()

	return l.Tokens
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
