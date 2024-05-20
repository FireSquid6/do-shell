package lexer

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	OPEN_PAREN  = "("
	CLOSE_PAREN = ")"
	OPEN_BRACE  = "{"
	CLOSE_BRACE = "}"

	COMMA = ","
	DOT   = "."

	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	NOT           = "!"
	NOT_EQUAL     = "!="
	EQUAL         = "=="
	ASSIGNMENT    = "="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	AND = "AND"
	OR  = "OR"

	ELSE  = "ELSE"
	FALSE = "FALSE"
	FUNC  = "FUNC"
	FOR   = "FOR"
	IF    = "IF"

	ASSIGN  = "="
	COMMENT = "#"

	NEWLINE = "\n"

	NIL = "NIL"
)

type Token struct {
	kind string
	data map[string]string
}

type Lexer struct {
  i int
  tokens []Token
}


func Lex(source string) ([]Token, error) {
	var i = 0
	tokens := []Token{}

	currentString := ""
	mode := "normal" // normal, string, number, identifier

	for i < len(source) {
		c := source[i]

		if mode == "string" {
			// todo: support faststrings
			if c == '"' {
				tokens = append(tokens, Token{kind: STRING, data: map[string]string{"value": currentString}})
				currentString = ""
				mode = "normal"
			} else {
				currentString += string(c)
			}

			continue
		}

		switch c {
		// single character tokens
		case '+':
			tokens = append(tokens, Token{kind: PLUS})
		case '-':
			tokens = append(tokens, Token{kind: MINUS})
		case '*':
			tokens = append(tokens, Token{kind: ASTERISK})
		case '/':
			tokens = append(tokens, Token{kind: SLASH})
		case '(':
			tokens = append(tokens, Token{kind: OPEN_PAREN})
		case ')':
			tokens = append(tokens, Token{kind: CLOSE_PAREN})
		case '{':
			tokens = append(tokens, Token{kind: OPEN_BRACE})
		case '}':
			tokens = append(tokens, Token{kind: CLOSE_BRACE})
		case ',':
			tokens = append(tokens, Token{kind: COMMA})
		case '.':
			tokens = append(tokens, Token{kind: DOT})
		case '!':
			if peekAhead(source, '=', &i) {
				tokens = append(tokens, Token{kind: NOT_EQUAL})
			} else {
				tokens = append(tokens, Token{kind: NOT})
			}
		case '=':
			if peekAhead(source, '=', &i) {
				tokens = append(tokens, Token{kind: EQUAL})
			} else {
				tokens = append(tokens, Token{kind: ASSIGNMENT})
			}
		case '>':
			if peekAhead(source, '=', &i) {
				tokens = append(tokens, Token{kind: GREATER_EQUAL})
			} else {
				tokens = append(tokens, Token{kind: GREATER})
			}
		case '<':
			if peekAhead(source, '=', &i) {
				tokens = append(tokens, Token{kind: LESS_EQUAL})
			} else {
				tokens = append(tokens, Token{kind: LESS})
			}
		case '#':
			tokens = append(tokens, Token{kind: COMMENT})
		case '\n':
			// newlines act as semicolons
			tokens = append(tokens, Token{kind: NEWLINE})
		case ' ':
			// ignore whitespace. This will kill any identifier though
		default:
			// identifiers and keywords
			// todo: match with regex. If it's

		}

	}

	return tokens, nil
}

func peekAhead(source string, match byte, i *int) bool {
	if source[*i+1] == match {
		*i++
		return true
	}

	return false
}
