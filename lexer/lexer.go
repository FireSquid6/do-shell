package lexer

import "golang.org/x/exp/constraints"

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

	SEMICOLON = ";"
	ASSIGN    = "="
	COMMENT   = "#"

	NIL = "NIL"
)

type Token struct {
	kind string
	data map[string]string
}

func Lex(source string) ([]Token, error) {
	return []Token{}, nil
}
