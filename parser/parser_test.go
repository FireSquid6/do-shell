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
              Value: 5,
            },
          },
        },
      },
    },
    {
      input: []token.Token{
        {Type: token.RETURN, Literal: []rune("return")},
        {Type: token.NUMBER, Literal: []rune("5")},
      },
      expectation: &tree.Program{
        Statements: []tree.Statement{
          &tree.ReturnStatement{
            Token: token.Token{Type: token.RETURN, Literal: []rune("return")},
            ReturnValue: &tree.NumberLiteral{
              Token: token.Token{Type: token.NUMBER, Literal: []rune("5")},
              Value: 5,
            },
          },
        },
      },
    },
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
