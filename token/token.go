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
  LINEBREAK

	LPAREN
	RPAREN
	LBRACE
	RBRACE
  LBRACKET
  RBRACKET

	// keywords
	FUNCTION
	LET
)

func ReadableTokenName(t Token) string {
  tokenMap := map[TokenType]string{
    ILLEGAL: "ILLEGAL",
    EOF: "EOF",
    IDENTIFIER: "IDENTIFIER",
    ASSIGN: "ASSIGN",
    LPAREN: "LPAREN",
    RPAREN: "RPAREN",
    LBRACE: "LBRACE",
    RBRACE: "RBRACE",
    COMMA: "COMMA",
    LINEBREAK: "LINEBREAK",
    FUNCTION: "FUNCTION",
    LET: "LET",
    INT: "INT",
    RBRACKET: "RBRACKET",
    LBRACKET: "LBRACKET",
  }

  return tokenMap[t.Type]
}
