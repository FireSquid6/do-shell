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
  NUMBER
  STRING
  COMMAND

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
    NUMBER: "NUMBER",
    STRING: "STRING",
    COMMAND: "COMMAND",
    PLUS: "PLUS",
  }

  return tokenMap[t.Type]
}


func LookupIdentifier(literal []rune) TokenType {
  tokenMap := map[string]TokenType {
    "let": LET,
    "fn": FUNCTION,
  }

  if token, ok := tokenMap[string(literal)]; ok {
    return token
  }

  return IDENTIFIER
}
