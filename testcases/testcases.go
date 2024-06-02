package testcases

import (
	"os"
	"path"
  "fmt"

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

// this will fail if not run from the root of the project
func GetTestcases() []Testcase {
	// all text starts out as empty and is read into the testcases
	cases := []Testcase{}
  fmt.Println("Getting testcases")

	expectations := []Expectation{
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
