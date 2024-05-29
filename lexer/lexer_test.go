package lexer

import (
	"github.com/firesquid6/do-shell/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
    name string
		input  string
		output []token.Token
	}{
		{
      name: "Simple symbols",
			input: "+=(){}\n",
			output: []token.Token{
				{Type: token.PLUS, Literal: []rune{'+'}},
				{Type: token.ASSIGN, Literal: []rune{'='}},
				{Type: token.LPAREN, Literal: []rune{'('}},
				{Type: token.RPAREN, Literal: []rune{')'}},
				{Type: token.LBRACE, Literal: []rune{'{'}},
				{Type: token.RBRACE, Literal: []rune{'}'}},
				{Type: token.LINEBREAK, Literal: []rune{'\n'}},
			},
		},
		{
      name: "identifiers",
			input: `let five = 5
      let ten = 10

      let add = fn(x, y) {
        x + y
      }

      let result = add (five, ten)`,
			output: []token.Token{
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("five")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.INT, Literal: []rune("5")},
				{Type: token.LINEBREAK, Literal: []rune("\n")},
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("ten")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.INT, Literal: []rune("10")},
				{Type: token.LINEBREAK, Literal: []rune("\n")},
				{Type: token.LINEBREAK, Literal: []rune("\n")},
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("add")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.FUNCTION, Literal: []rune("fn")},
				{Type: token.LPAREN, Literal: []rune("(")},
				{Type: token.IDENTIFIER, Literal: []rune("x")},
				{Type: token.COMMA, Literal: []rune(",")},
				{Type: token.IDENTIFIER, Literal: []rune("y")},
				{Type: token.RPAREN, Literal: []rune(")")},
				{Type: token.LBRACE, Literal: []rune("{")},
				{Type: token.LINEBREAK, Literal: []rune("\n")},
				{Type: token.IDENTIFIER, Literal: []rune("x")},
				{Type: token.PLUS, Literal: []rune("+")},
				{Type: token.IDENTIFIER, Literal: []rune("y")},
				{Type: token.LINEBREAK, Literal: []rune("\n")},
				{Type: token.RBRACE, Literal: []rune("}")},
				{Type: token.LINEBREAK, Literal: []rune("\n")},
				{Type: token.LINEBREAK, Literal: []rune("\n")},
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("result")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.IDENTIFIER, Literal: []rune("add")},
				{Type: token.LPAREN, Literal: []rune("(")},
				{Type: token.IDENTIFIER, Literal: []rune("five")},
				{Type: token.COMMA, Literal: []rune(",")},
				{Type: token.IDENTIFIER, Literal: []rune("ten")},
				{Type: token.RPAREN, Literal: []rune(")")},
			},
		},
	}

	for _, test := range tests {
		l := New(test.input)

		for _, expected := range test.output {
			tok := l.NextToken()

			expectedType := token.ReadableTokenName(expected)
			tokType := token.ReadableTokenName(tok)
      name := test.name

			if tok.Type != expected.Type {
				t.Fatalf("tests[%s] - tokentype wrong. expected=%q, got=%q", name, expectedType, tokType)
			}

			if string(tok.Literal) != string(expected.Literal) {
				t.Fatalf("tests[%s] - literal wrong. expected=%q, got=%q", name, expected.Literal, tok.Literal)
			}
		}
	}
}
