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
  COMMAND_SEGMENT

	// operators
	ASSIGN

	PLUS
  MINUS
  DIVIDE
  MULTIPLY
  MOD

  GREATER_THAN
  LESS_THAN
  GREATER_THAN_EQUAL
  LESS_THAN_EQUAL
  EQUAL

	// delimiters
	COMMA
  SEMICOLON
  DOT
  GRAVE

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// keywords
	FUNCTION
	LET
  FOR
  IN
  ELSE
  IF
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
    SEMICOLON: "SEMICOLON",
    FUNCTION: "FUNCTION",
    DOT: "DOT",
    LET: "LET",
    NUMBER: "NUMBER",
    STRING: "STRING",
    COMMAND_SEGMENT: "COMMAND_SEGMENT",
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
