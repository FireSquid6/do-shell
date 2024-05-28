package lexer

import (
	"github.com/firesquid6/do-shell/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		input  string
		output []token.Token
	}{
		{
			input: "+=;(){}",
			output: []token.Token{
				{Type: token.PLUS, Literal: []rune{'+'}},
				{Type: token.ASSIGN, Literal: []rune{'='}},
				{Type: token.LINEEND, Literal: []rune{';'}},
				{Type: token.LPAREN, Literal: []rune{'('}},
				{Type: token.RPAREN, Literal: []rune{')'}},
				{Type: token.LBRACE, Literal: []rune{'{'}},
				{Type: token.RBRACE, Literal: []rune{'}'}},
			},
		},
		// {
		// 	input: "let five = 5\n",
		// 	output: []token.Token{
		// 		{Type: token.LET, Literal: "let"},
		// 		{Type: token.IDENTIFIER, Literal: "five"},
		// 		{Type: token.ASSIGN, Literal: "="},
		// 		{Type: token.INT, Literal: "5"},
		// 		{Type: token.LINEEND, Literal: ";"},
		// 	},
		// },
	}

	for _, test := range tests {
		l := New(test.input)

		for i, expected := range test.output {
			tok := l.NextToken()

      expectedType := token.ReadableTokenName(expected)
      tokType := token.ReadableTokenName(tok)

			if tok.Type != expected.Type {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, expectedType, tokType)
			}

			if string(tok.Literal) != string(expected.Literal) {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expected.Literal, tok.Literal)
			}
		}
	}
}
