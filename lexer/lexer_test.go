package lexer_test

import (
	"os"
	"testing"
  "path"

	"github.com/firesquid6/do-shell/lexer"
	"github.com/firesquid6/do-shell/token"
)

type Expectation struct {
	Tokens   []token.Token
	Filename string
}

type Testcase struct {
	Expectation Expectation
	Text        string
}

func GetTestcases() []Testcase {
	// all text starts out as empty and is read into the testcases
	cases := []Testcase{}

	expectations := []Expectation{
		{
			Filename: "let_statements.dosh",
			Tokens: []token.Token{
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("x")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.NUMBER, Literal: []rune("5")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("y")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.NUMBER, Literal: []rune("10")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("foobar")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.NUMBER, Literal: []rune("12312")},
        {Type: token.SEMICOLON, Literal: []rune(";")},
        {Type: token.EOF, Literal: []rune{0}},
			},
		},
		{
			Filename: "basic_symbols.dosh",
			Tokens: []token.Token{
				{Type: token.PLUS, Literal: []rune{'+'}},
				{Type: token.ASSIGN, Literal: []rune{'='}},
				{Type: token.LPAREN, Literal: []rune{'('}},
				{Type: token.RPAREN, Literal: []rune{')'}},
				{Type: token.LBRACE, Literal: []rune{'{'}},
				{Type: token.RBRACE, Literal: []rune{'}'}},
				{Type: token.SEMICOLON, Literal: []rune{';'}},
				{Type: token.EOF, Literal: []rune{0}},
			},
		},
		{
			Filename: "identifiers.dosh",
			Tokens: []token.Token{
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("five")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.NUMBER, Literal: []rune("5")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("ten")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.NUMBER, Literal: []rune("10")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
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
				{Type: token.IDENTIFIER, Literal: []rune("x")},
				{Type: token.PLUS, Literal: []rune("+")},
				{Type: token.IDENTIFIER, Literal: []rune("y")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.RBRACE, Literal: []rune("}")},
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("result")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.IDENTIFIER, Literal: []rune("add")},
				{Type: token.LPAREN, Literal: []rune("(")},
				{Type: token.IDENTIFIER, Literal: []rune("five")},
				{Type: token.COMMA, Literal: []rune(",")},
				{Type: token.IDENTIFIER, Literal: []rune("ten")},
				{Type: token.RPAREN, Literal: []rune(")")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.EOF, Literal: []rune{0}},
			},
		},
		{
			Filename: "operators.dosh",
			Tokens: []token.Token{
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("four")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.NUMBER, Literal: []rune("2")},
				{Type: token.MULTIPLY, Literal: []rune("*")},
				{Type: token.NUMBER, Literal: []rune("2")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.IF, Literal: []rune("if")},
				{Type: token.IDENTIFIER, Literal: []rune("four")},
				{Type: token.EQUAL, Literal: []rune("==")},
				{Type: token.NUMBER, Literal: []rune("4")},
				{Type: token.LBRACE, Literal: []rune("{")},
				{Type: token.IDENTIFIER, Literal: []rune("print")},
				{Type: token.LPAREN, Literal: []rune("(")},
				{Type: token.STRING, Literal: []rune("four is 4")},
				{Type: token.RPAREN, Literal: []rune(")")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.RBRACE, Literal: []rune("}")},
				{Type: token.ELSE, Literal: []rune("else")},
				{Type: token.LBRACE, Literal: []rune("{")},
				{Type: token.IDENTIFIER, Literal: []rune("print")},
				{Type: token.LPAREN, Literal: []rune("(")},
				{Type: token.STRING, Literal: []rune("four is not 4")},
				{Type: token.RPAREN, Literal: []rune(")")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.RBRACE, Literal: []rune("}")},
				{Type: token.EOF, Literal: []rune{0}},
			},
		},
		{
			Filename: "comparisons.dosh",
			Tokens: []token.Token{
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("x")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.NUMBER, Literal: []rune("2")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("y")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.NUMBER, Literal: []rune("4")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.IDENTIFIER, Literal: []rune("x")},
				{Type: token.GREATER_THAN_EQUAL, Literal: []rune(">=")},
				{Type: token.NUMBER, Literal: []rune("2")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.IDENTIFIER, Literal: []rune("y")},
				{Type: token.LESS_THAN_EQUAL, Literal: []rune("<=")},
				{Type: token.NUMBER, Literal: []rune("4")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.IDENTIFIER, Literal: []rune("x")},
				{Type: token.EQUAL, Literal: []rune("==")},
				{Type: token.NUMBER, Literal: []rune("2")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.IDENTIFIER, Literal: []rune("y")},
				{Type: token.NOT_EQUAL, Literal: []rune("!=")},
				{Type: token.IDENTIFIER, Literal: []rune("x")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.IDENTIFIER, Literal: []rune("x")},
				{Type: token.GREATER_THAN, Literal: []rune(">")},
				{Type: token.NUMBER, Literal: []rune("2")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.IDENTIFIER, Literal: []rune("y")},
				{Type: token.LESS_THAN, Literal: []rune("<")},
				{Type: token.NUMBER, Literal: []rune("2")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.EOF, Literal: []rune{0}},
			},
		},
		{
			Filename: "comments.dosh",
			Tokens: []token.Token{
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("mystuff")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.STRING, Literal: []rune("whatever")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("otherstuff")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.STRING, Literal: []rune("this comes after the comment")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.EOF, Literal: []rune{0}},
			},
		},
		{
			Filename: "strings_and_commands.dosh",
			Tokens: []token.Token{
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("mystr")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.STRING, Literal: []rune("hello, world!")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.FOR, Literal: []rune("for")},
				{Type: token.IDENTIFIER, Literal: []rune("c")},
				{Type: token.IN, Literal: []rune("in")},
				{Type: token.IDENTIFIER, Literal: []rune("mystr")},
				{Type: token.LBRACE, Literal: []rune("{")},
				{Type: token.COMMAND, Literal: []rune("echo {c}")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.RBRACE, Literal: []rune("}")},
				{Type: token.EOF, Literal: []rune{0}},
			},
		},
	}

	for _, expectation := range expectations {
		file := path.Join("../testcases", expectation.Filename)
		text, err := os.ReadFile(file)

		if err != nil {
			// if this is panicking, you're probably not running from the right directory
			panic(err)
		}

		cases = append(cases, Testcase{Text: string(text), Expectation: expectation})
	}

	return cases
}

func TestNextToken(t *testing.T) {
	tests := GetTestcases()

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
