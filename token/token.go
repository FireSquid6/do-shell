package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal []rune
	// TODO: add line and column info
}

const (
	ILLEGAL TokenType = iota
	EOF

	// identifiers and literals
	IDENTIFIER
	INT

	// operators
	ASSIGN
	PLUS

	// delimiters
	COMMA
	LINEEND

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// keywords
	FUNCTION
	LET
)

func ReadableTokenName(t Token) string {
	switch t.Type {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENTIFIER:
		return "IDENTIFIER"
	case INT:
		return "INT"
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case COMMA:
		return "COMMA"
	case LINEEND:
		return "LINEEND"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	default:
		return "UNKNOWN"
	}
}
