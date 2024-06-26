package parser_test

import (
	"testing"
  "fmt"

	"github.com/firesquid6/do-shell/parser"
	"github.com/firesquid6/do-shell/token"
	"github.com/firesquid6/do-shell/tree"
)

func TestParser(t *testing.T) {
	tests := []struct {
    name       string
		input       []token.Token
		expectation *tree.Program
		expectedErr []error
	}{
		{
      name: "Empty program",
			input: []token.Token{},
			expectation: &tree.Program{
				Statements: []tree.Statement{},
			},
			expectedErr: []error{},
		},
		{
      name: "Let statement",
			input: []token.Token{
				{Type: token.LET, Literal: []rune("let")},
				{Type: token.IDENTIFIER, Literal: []rune("x")},
				{Type: token.ASSIGN, Literal: []rune("=")},
				{Type: token.NUMBER, Literal: []rune("5")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.EOF, Literal: []rune("")},
			},
			expectation: &tree.Program{
				Statements: []tree.Statement{
					&tree.LetStatement{
						Token: token.Token{Type: token.LET, Literal: []rune("let")},
						Name: &tree.Identifier{
							Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("x")},
							Value: []rune("x"),
						},
						Expression: &tree.NumberLiteral{
							Token: token.Token{Type: token.NUMBER, Literal: []rune("5")},
							Value: 5.0,
						},
					},
				},
			},
		},
		{
      name: "Return statement",
			input: []token.Token{
				{Type: token.RETURN, Literal: []rune("return")},
				{Type: token.NUMBER, Literal: []rune("5")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.EOF, Literal: []rune("")},
			},
			expectation: &tree.Program{
				Statements: []tree.Statement{
					&tree.ReturnStatement{
						Token: token.Token{Type: token.RETURN, Literal: []rune("return")},
						Expression: &tree.NumberLiteral{
							Token: token.Token{Type: token.NUMBER, Literal: []rune("5")},
							Value: 5.0,
						},
					},
				},
			},
		},
		{
      name: "Function literal",
			input: []token.Token{
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
				{Type: token.IDENTIFIER, Literal: []rune("add")},
				{Type: token.LPAREN, Literal: []rune("(")},
				{Type: token.IDENTIFIER, Literal: []rune("five")},
				{Type: token.COMMA, Literal: []rune(",")},
				{Type: token.IDENTIFIER, Literal: []rune("ten")},
				{Type: token.RPAREN, Literal: []rune(")")},
				{Type: token.SEMICOLON, Literal: []rune(";")},
				{Type: token.EOF, Literal: []rune{0}},
			},
			expectation: &tree.Program{
				Statements: []tree.Statement{
					&tree.LetStatement{
						Token: token.Token{Type: token.LET, Literal: []rune("let")},
						Name: &tree.Identifier{
							Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("five")},
							Value: []rune("five"),
						},
						Expression: &tree.NumberLiteral{
							Token: token.Token{Type: token.NUMBER, Literal: []rune("5")},
							Value: 5.0,
						},
					},
					&tree.LetStatement{
						Token: token.Token{Type: token.LET, Literal: []rune("let")},
						Name: &tree.Identifier{
							Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("ten")},
							Value: []rune("ten"),
						},
						Expression: &tree.NumberLiteral{
							Token: token.Token{Type: token.NUMBER, Literal: []rune("10")},
							Value: 10.0,
						},
					},
					&tree.LetStatement{
						Token: token.Token{Type: token.LET, Literal: []rune("let")},
						Name: &tree.Identifier{
							Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("add")},
							Value: []rune("add"),
						},

						Expression: &tree.FunctionLiteral{
							Token: token.Token{Type: token.FUNCTION, Literal: []rune("fn")},
							Parameters: []*tree.Identifier{
								{Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("x")}, Value: []rune("x")},
								{Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("y")}, Value: []rune("y")},
							},
							Statements: &tree.BlockStatement{
								Token: token.Token{Type: token.LBRACE, Literal: []rune("{")},
								Statements: []tree.Statement{
									&tree.ExpressionStatement{
										Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("x")},
										Expression: &tree.InfixExpression{
											Token:    token.Token{Type: token.PLUS, Literal: []rune("+")},
											Operator: []rune("+"),
											Left: &tree.Identifier{
												Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("x")},
												Value: []rune("x"),
											},
											Right: &tree.Identifier{
												Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("y")},
												Value: []rune("y"),
											},
										},
									},
								},
							},
						},
					},
					&tree.ExpressionStatement{
						Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("add")},
						Expression: &tree.CallExpression{
							Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("add")},
							Function: tree.Identifier{
								Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("add")},
								Value: []rune("add"),
							},
							Arguments: []tree.Expression{
								&tree.Identifier{
									Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("five")},
									Value: []rune("five"),
								},
								&tree.Identifier{
									Token: token.Token{Type: token.IDENTIFIER, Literal: []rune("ten")},
									Value: []rune("ten"),
								},
							},
						},
					},
				},
			},
		},
		// To test:
		// - when it fails with a bunch of errors
		// - if/else statements
		// - commands
		// - while statemnts
		// - for statements
		// - strings
	}

	for _, tt := range tests {
    fmt.Println("\n\nAbout to do a new test, ", tt.name)
		p := parser.New(tt.input)
		program := p.ParseProgram()

		if len(p.Errors) != len(tt.expectedErr) {
      // print all of the errors
      for _, err := range p.Errors {
        t.Log(err)
      }
      t.Errorf("%s: Expected %d errors, got %d", tt.name, len(tt.expectedErr), len(p.Errors))
		}

		for i, err := range p.Errors {
			if err != tt.expectedErr[i] {
        t.Errorf("%s: Expected error %s, got %s", tt.name, tt.expectedErr[i], err)
			}
		}

		programString := program.String()
		expectedProgramString := tt.expectation.String()

		if programString != expectedProgramString {
      t.Errorf("%s: Expected program string to be %s, got %s", tt.name, expectedProgramString, programString)
		}
	}
}
