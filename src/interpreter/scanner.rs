use super::error::{InterpreterErrors, ErrorKind};

#[derive(Debug, Eq, PartialEq)]
enum TokenKind {
    LPAREN, RPAREN, LBRACE, LBRACKET, RBRACKET, RBRACE, COMMA, DOT, SEMICOLON, COLON,
    MINUS, PLUS, MULTIPLY, DIVIDE, INTEGERDIVIDE, MODULO, RAISETO,
    EQUAL, NOTEQUAL, GREATER, GREATEREQUAL, LESS, LESSEQUAL,
    AND, OR, NOT,
    IDENTIFIER, STING, NUMBER,
    IF, ELIF, ELSE, LET, FOR, USE, STRUCT, WHILE, RETURN, // MATCH, CASE
    EOF, UNKNOWN,
}


#[derive(Debug)]
pub struct Token {
    kind: TokenKind,
    lexeme: String,
    line: u32,
    column: u32,
}

impl Token {
    fn to_string(self: Token) {
        println!("Token: {:?} {} {} {}", self.kind, self.lexeme, self.line, self.column);
    }

    fn new(kind: TokenKind, lexeme: String, line: u32, column: u32) -> Token {
        Token {
            kind,
            lexeme,
            line,
            column,
        }
    }

}


pub struct Scanner {
    source: String,
    tokens: Vec<Token>,

    current: u32,
    errors: InterpreterErrors,

    line: u32,
    col: u32
}

impl Scanner {
    pub fn scan(self: &mut Scanner) {
        while self.current < self.source.len() as u32 {
            self.scan_token();
        }
    }

    fn new(source: String) -> Scanner {
        Scanner {
            source,
            tokens: vec![],

            line: 1,
            col: 1,

            errors: InterpreterErrors::new(),

            current: 0,
        }
    }

    fn advance(self: &mut Scanner) -> char {
        let c = self.source.chars().nth(self.current as usize).unwrap_or('\0');
        
        self.current += 1;
        self.col += 1;

        if c == '\n' {
            self.line += 1;
            self.col = 1;
        }

        return c
    }

    fn add_token(self: &mut Scanner, kind: TokenKind, lexeme: String) {
        self.tokens.push(Token::new(kind, lexeme, self.line, self.col))

    }

    fn add_token_at_current(self: &mut Scanner, kind: TokenKind) {
        println!("Adding token: {:?}", kind);
        let c = self.source.chars().nth(self.current as usize).unwrap_or('\0');
        self.tokens.push(Token::new(kind, c.to_string(), self.line, self.col))
    }


    fn scan_token(self: &mut Scanner) {
        let c = self.advance();
        println!("Scanning: {}", c);

        match c {
            '(' => self.add_token_at_current(TokenKind::LPAREN),
            ')' => self.add_token_at_current(TokenKind::RPAREN),
            '[' => self.add_token_at_current(TokenKind::LBRACKET),
            ']' => self.add_token_at_current(TokenKind::RBRACKET),
            '{' => self.add_token_at_current(TokenKind::LBRACE),
            '}' => self.add_token_at_current(TokenKind::RBRACE),
            ',' => self.add_token_at_current(TokenKind::COMMA),
            '.' => self.add_token_at_current(TokenKind::DOT),
            '-' => self.add_token_at_current(TokenKind::MINUS),
            '+' => self.add_token_at_current(TokenKind::PLUS),
            ';' => self.add_token_at_current(TokenKind::SEMICOLON),
            '\n' => {
                // TODO - when we see \n, look back. If it's something that implies the start of
                // a new line, then we need to insert a semicolon. A line probably ends if:
                // - the previous token is string literal, identifier, ), or ]
                //
                // We definitely DO NOT want to insert a semicolon if:
                // - the previous token is a comma
                // - the next valid character is a dot (probably chaining methods or properties)
                // - the next valid character is a comma (probably a list)
                // - the next valid character is a || or && (probably a logical operator)
                //
                // Other valid way:
                // - convert all \ns in a row to a single NEWLINE token
                // - after we scan, we can insert semicolons where needed and remove all NEWLINE
                // tokens

            }

            _ => {
                self.errors.add_error("Unexpected token".to_string(), ErrorKind::LEXER);
            }

        }

    }
}


#[cfg(test)] 
mod tests {
    use super::*;

    #[test]
    fn test_basic_tokens() {
        let mut scanner = Scanner::new("()[]-+.,;*/".to_string());
        scanner.scan();

        let expected_tokens: Vec<TokenKind> = vec![
            TokenKind::LPAREN, TokenKind::RPAREN, TokenKind::LBRACKET, TokenKind::RBRACKET,
            TokenKind::MINUS, TokenKind::PLUS, TokenKind::DOT, TokenKind::COMMA, TokenKind::SEMICOLON,
            TokenKind::MULTIPLY, TokenKind::DIVIDE
        ];

        assert_eq!(scanner.tokens.len(), expected_tokens.len());
        for (i, token) in scanner.tokens.iter().enumerate() {
            assert_eq!(token.kind, expected_tokens[i]);
        }
    }

    #[test]
    fn test_multi_char_tokens() {
        let mut scanner = Scanner::new("// != == >= <= && || !".to_string());
        scanner.scan();
        let expected_tokens: Vec<TokenKind> = vec![
            TokenKind::INTEGERDIVIDE, TokenKind::NOTEQUAL, TokenKind::EQUAL, TokenKind::GREATEREQUAL,
            TokenKind::LESSEQUAL, TokenKind::AND, TokenKind::OR, TokenKind::NOT
        ];

        assert_eq!(scanner.tokens.len(), expected_tokens.len());
        for (i, token) in scanner.tokens.iter().enumerate() {
            assert_eq!(token.kind, expected_tokens[i]);
        }

        assert_eq!(0, 1);
    }

}
