package lexer_test

import (
	"fmt"
	"testing"

	"github.com/firesquid6/do-shell/lexer"
	"github.com/firesquid6/do-shell/testcases"
	"github.com/firesquid6/do-shell/token"
)

func TestNextToken(t *testing.T) {
  fmt.Println("Hello world")
	tests := testcases.GetTestcases()


	for _, test := range tests {
		tokens := lexer.Tokenize(test.Text)
		expectation := test.Expectation

		if len(tokens) != len(expectation.Tokens) {
			printInformation(t, tokens, expectation.Tokens)
			t.Fatalf("Expected %d tokens, got %d", len(expectation.Tokens), len(tokens))
		}

		for i, tok := range tokens {
			tokenType := token.ReadableTokenName(tok)
			expectedType := token.ReadableTokenName(expectation.Tokens[i])

			if tok.Type != expectation.Tokens[i].Type {
				printInformation(t, tokens, expectation.Tokens)
				t.Fatalf("Expected token type %d to be '%s', got '%s'", i, tokenType, expectedType)
			}

			if string(tok.Literal) != string(expectation.Tokens[i].Literal) {
				printInformation(t, tokens, expectation.Tokens)
				t.Fatalf("Expected token literal %d to be '%s', got '%s'", i, string(expectation.Tokens[i].Literal), string(tok.Literal))
			}
		}

	}
}

func printInformation(t *testing.T, tokens []token.Token, expectedTokens []token.Token) {
	for i, tok := range tokens {
		tokenType := token.ReadableTokenName(tok)
		var expectedType string

		if i < len(expectedTokens) {
			expectedType = token.ReadableTokenName(expectedTokens[i])
		} else {
			expectedType = "OUT_OF_BOUNDS"
		}

		t.Logf("Token %d: Expected '%s', got '%s'", i, expectedType, tokenType)
	}

}
