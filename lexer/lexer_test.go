package lexer_test

import (
  "github.com/firesquid6/do-shell/testcases"
  "github.com/firesquid6/do-shell/lexer"
  "github.com/firesquid6/do-shell/token"
	"testing"
)

func TestNextToken(t *testing.T) {
  tests := testcases.GetTestcases()

	for _, test := range tests {
    tokens := lexer.Tokenize(test.Text)

    if len(tokens) != len(test.ExpectedTokens) {
      t.Fatalf("Expected %d tokens, got %d", len(test.ExpectedTokens), len(tokens))
    }

    for i, tok:= range tokens {
      tokenType := token.ReadableTokenName(tok)
      expectedType := token.ReadableTokenName(test.ExpectedTokens[i])

      if tok.Type != test.ExpectedTokens[i].Type {
        t.Fatalf("Expected token type %d to be '%s', got '%s'", i, tokenType, expectedType)
      }

      if string(tok.Literal) != string(test.ExpectedTokens[i].Literal) {
        t.Fatalf("Expected token literal %d to be '%s', got '%s'", i, string(test.ExpectedTokens[i].Literal), string(tok.Literal))
      }
    }

	}
}
