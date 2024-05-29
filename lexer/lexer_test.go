package lexer

import (
	"github.com/firesquid6/do-shell/token"
  "github.com/firesquid6/do-shell/testcases"
	"testing"
)

func TestNextToken(t *testing.T) {
  tests := testcases.GetTestcases()

	for _, test := range tests {
		l := New(test.Text)

		for _, expected := range test.ExpectedTokens{
			tok := l.NextToken()

			expectedType := token.ReadableTokenName(expected)
			tokType := token.ReadableTokenName(tok)
			name := test.Filename

			if tok.Type != expected.Type {
				t.Fatalf("tests[%s] - tokentype wrong. expected=%q, got=%q", name, expectedType, tokType)
			}

			if string(tok.Literal) != string(expected.Literal) {
				t.Fatalf("tests[%s] - literal wrong. expected=%q, got=%q", name, expected.Literal, tok.Literal)
			}
		}
	}
}
