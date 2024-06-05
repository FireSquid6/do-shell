package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal []rune
  // TODO: make these actually work
	// Column  int
	// Line    int
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
	MINUS
	DIVIDE
	MULTIPLY
	MOD

	NOT
	NOT_EQUAL

	GREATER_THAN
	LESS_THAN
	GREATER_THAN_EQUAL
	LESS_THAN_EQUAL
	EQUAL

	// delimiters
	COMMA
	SEMICOLON
	DOT

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
  RETURN
)

func ReadableTokenName(t Token) string {
	tokenMap := map[TokenType]string{
		ILLEGAL:            "ILLEGAL",
		EOF:                "EOF",
		IDENTIFIER:         "IDENTIFIER",
		ASSIGN:             "ASSIGN",
		LPAREN:             "LPAREN",
		RPAREN:             "RPAREN",
		LBRACE:             "LBRACE",
		RBRACE:             "RBRACE",
		COMMA:              "COMMA",
		SEMICOLON:          "SEMICOLON",
		FUNCTION:           "FUNCTION",
		DOT:                "DOT",
		LET:                "LET",
		NUMBER:             "NUMBER",
		STRING:             "STRING",
		COMMAND:            "COMMAND",
		PLUS:               "PLUS",
		IF:                 "IF",
		ELSE:               "ELSE",
		FOR:                "FOR",
		IN:                 "IN",
		MULTIPLY:           "MULTIPLY",
		EQUAL:              "EQUAL",
		GREATER_THAN:       "GREATER_THAN",
		LESS_THAN:          "LESS_THAN",
		GREATER_THAN_EQUAL: "GREATER_THAN_EQUAL",
    NOT_EQUAL:          "NOT_EQUAL",
		LESS_THAN_EQUAL:    "LESS_THAN_EQUAL",
	}

	return tokenMap[t.Type]
}

func LookupIdentifier(literal []rune) TokenType {
	tokenMap := map[string]TokenType{
		"let":  LET,
		"fn":   FUNCTION,
		"if":   IF,
		"else": ELSE,
		"for":  FOR,
		"in":   IN,
    "return": RETURN,
	}

	if token, ok := tokenMap[string(literal)]; ok {
		return token
	}

	return IDENTIFIER
}
