package testcases

import (
	"github.com/firesquid6/do-shell/token"
	"os"
	"path"
)

type Testcase struct {
	Text           string
	Filename       string
	ExpectedTokens []token.Token
}

// this will fail if not run from the root of the project
func GetTestcases() []Testcase {
	// all text starts out as empty and is read into the testcases
	cases := []Testcase{
		{
			Text:     "",
			Filename: "basic_symbols.do",
			ExpectedTokens: []token.Token{
				{Type: token.PLUS, Literal: []rune{'+'}},
				{Type: token.ASSIGN, Literal: []rune{'='}},
				{Type: token.LPAREN, Literal: []rune{'('}},
				{Type: token.RPAREN, Literal: []rune{')'}},
				{Type: token.LBRACE, Literal: []rune{'{'}},
				{Type: token.RBRACE, Literal: []rune{'}'}},
				{Type: token.LINEBREAK, Literal: []rune{'\n'}},
        {Type: token.EOF, Literal: []rune{0}},
			},
		},
		{
			Text:     "",
			Filename: "identifiers.do",
			ExpectedTokens: []token.Token{
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

	for _, test := range cases {
		file := path.Join("../testcases", test.Filename)
		text, err := os.ReadFile(file)

		if err != nil {
			// if this is panicking, you're probably not running from the right directory
			panic(err)
		}

		test.Text = string(text)
	}

	return cases
}
