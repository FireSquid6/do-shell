#[derive(Debug)]
enum TokenKind {
    LPAREN, RPAREN, LBRACE, LBRACKET, RBRACKET, RBRACE, COMMA, DOT, SEMICOLON, COLON,
    MINUS, PLUS, MULTIPLY, DIVIDE, INTEGERDIVIDE, MODULO, RAISETO,
    EQUAL, NOTEQUAL, GREATER, GREATEREQUAL, LESS, LESSEQUAL,
    AND, OR, NOT,
    IDENTIFIER, STING, NUMBER,
    IF, ELIF, ELSE, LET, FOR, USE, STRUCT, WHILE, RETURN, // MATCH, CASE
    EOF, UNKNOWN,
}


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

    line: u32,
    col: u32
}

impl Scanner {
    fn new(source: String) -> Scanner {
        Scanner {
            source,
            tokens: vec![],

            line: 1,
            col: 1,

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
        let c = self.source.chars().nth(self.current as usize).unwrap_or('\0');
        self.tokens.push(Token::new(kind, c.to_string(), self.line, self.col))
    }


    fn scan_token(self: &mut Scanner) {
        let c = self.advance();

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
            '*' => self.add_token_at_current(TokenKind::MULTIPLY),
            ';' => self.add_token_at_current(TokenKind::SEMICOLON),
            '\n' => {
                // TODO - when we see \n, look back. If it's something that implies the start of
                // a new line, then we need to insert a semicolon. A line probably ends if:
                // - the previous token is string literal, identifier, ), or ]
                //
                // We definitely DO NOT want to insert a semicolon if:
                // - the previous token is a comma
                // - the next valid character is a dot (probably chaining methods or properties)

            }

            _ => self.add_token(TokenKind::UNKNOWN, c.to_string())

        }

    }
}


#[cfg(test)]
fn test_scanner() {
    let scanner = Scanner::new("let x = 10;".to_string());

}



