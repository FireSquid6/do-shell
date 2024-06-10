package parser_test

import (
	"testing"

	"github.com/firesquid6/do-shell/parser"
	"github.com/firesquid6/do-shell/token"
	"github.com/firesquid6/do-shell/tree"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input       []token.Token
		expectation *tree.Program
		expectedErr []error
	}{
		{
			input: []token.Token{},
			expectation: &tree.Program{
				Statements: []tree.Statement{},
			},
			expectedErr: []error{},
		},
		{
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
		p := parser.New(tt.input)
		program := p.ParseProgram()

		if len(p.Errors) != len(tt.expectedErr) {
			t.Errorf("Expected %d errors, got %d", len(tt.expectedErr), len(p.Errors))
		}

		for i, err := range p.Errors {
			if err != tt.expectedErr[i] {
				t.Errorf("Expected error %s, got %s", tt.expectedErr[i], err)
			}
		}

		programString := program.String()
		expectedProgramString := tt.expectation.String()

		if programString != expectedProgramString {
			t.Errorf("Expected program string to be %s, got %s", expectedProgramString, programString)
		}
	}
}

func failTest(t *testing.T, p *parser.Parser) {
	t.Logf("Parser errors: %v", p.Errors)
	for _, err := range p.Errors {
		t.Errorf("Parser error: %s", err)
	}
	t.Errorf("You suck. Be better.")
}
