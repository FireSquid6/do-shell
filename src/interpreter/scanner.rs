use super::error::{InterpreterErrors, ErrorKind};

#[derive(Debug, Eq, PartialEq)]
enum TokenKind {
    LPAREN, RPAREN, LBRACE, LBRACKET, RBRACKET, RBRACE, COMMA, DOT, SEMICOLON, COLON,
    MINUS, PLUS, MULTIPLY, DIVIDE, INTEGERDIVIDE, MODULO, RAISETO,
    EQUAL, NOTEQUAL, GREATER, GREATEREQUAL, LESS, LESSEQUAL,
    AND, OR, NOT, ASSIGN,
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
        self.tokens.push(Token::new(kind, lexeme, self.line, self.col));

    }

    fn add_token_char(self: &mut Scanner, kind: TokenKind, c: char) {
        self.tokens.push(Token::new(kind, c.to_string(), self.line, self.col));
    }


    fn peek_and_move(self: &mut Scanner, expected: char) -> bool {
        if self.current >= self.source.len() as u32 {
            return false;
        }

        let c = self.source.chars().nth(self.current as usize).unwrap_or('\0');
        if c == expected {
            self.advance();
            return true;
        }

        return false;
    }

    fn double_char_or_default(self: &mut Scanner, c: char, single_kind: TokenKind, expected: char, double_kind: TokenKind) {
        if self.peek_and_move(expected) {
            self.add_token(double_kind, format!("{}{}", c, expected));
        } else {
            self.add_token_char(single_kind, c)
        }
    }

    fn double_char_or_error(self: &mut Scanner, c: char, expected: char, double_kind: TokenKind) {
        if self.peek_and_move(expected) {
            self.add_token(double_kind, format!("{}{}", c, expected));
        } else {
            self.errors.add_error("Lexer error".to_string(), ErrorKind::LEXER);
        }
    }


    fn scan_token(self: &mut Scanner) {
        let c = self.advance();

        match c {
            '#' => {
                // skip until the end of the line
                while self.current < self.source.len() as u32 && self.source.chars().nth(self.current as usize).unwrap_or('\0') != '\n' {
                    self.advance();
                }
            }

            // single chars
            '(' => self.add_token_char(TokenKind::LPAREN, c),
            ')' => self.add_token_char(TokenKind::RPAREN, c),
            '[' => self.add_token_char(TokenKind::LBRACKET, c),
            ']' => self.add_token_char(TokenKind::RBRACKET, c),
            '{' => self.add_token_char(TokenKind::LBRACE, c),
            '}' => self.add_token_char(TokenKind::RBRACE, c),
            ',' => self.add_token_char(TokenKind::COMMA, c),
            '.' => self.add_token_char(TokenKind::DOT, c),
            '-' => self.add_token_char(TokenKind::MINUS, c),
            '+' => self.add_token_char(TokenKind::PLUS, c),
            ';' => self.add_token_char(TokenKind::SEMICOLON, c),
            '%' => self.add_token_char(TokenKind::MODULO, c),
            
            // double chars
            '*' => self.double_char_or_default(c, TokenKind::MULTIPLY, '*', TokenKind::RAISETO),
            '/' => self.double_char_or_default(c, TokenKind::DIVIDE, '/', TokenKind::INTEGERDIVIDE),
            '!' => self.double_char_or_default(c, TokenKind::NOT, '=', TokenKind::NOTEQUAL),
            '<' => self.double_char_or_default(c, TokenKind::LESS, '=', TokenKind::LESSEQUAL),
            '>' => self.double_char_or_default(c, TokenKind::GREATER, '=', TokenKind::GREATEREQUAL),
            '=' => self.double_char_or_default(c, TokenKind::ASSIGN, '=', TokenKind::EQUAL),

            '&' => self.double_char_or_error(c, '&', TokenKind::AND),
            '|' => self.double_char_or_error(c, '|', TokenKind::OR),

            // always skip whitespace
            ' ' | '\t' => {}
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

    fn lexer_test(source: String, expected_tokens: Vec<TokenKind>, expected_lexemes: Vec<&str>, expected_errors: Vec<&str>) {
        let mut scanner = Scanner::new(source);
        scanner.scan();

        println!("{:?}", scanner.tokens);

        assert!(scanner.errors.expect_errors(expected_errors));

        assert_eq!(scanner.tokens.len(), expected_tokens.len());
        for (i, token) in scanner.tokens.iter().enumerate() {
            assert_eq!(token.kind, expected_tokens[i]);
            assert_eq!(token.lexeme, expected_lexemes[i].to_string());
        }
    }

    // TODO - also test that the lexemes are correct
    #[test]
    fn test_basic_tokens() {
        let mut scanner = Scanner::new("()[]-+.,;*/".to_string());
        scanner.scan();

        let expected_tokens: Vec<TokenKind> = vec![
            TokenKind::LPAREN, TokenKind::RPAREN, TokenKind::LBRACKET, TokenKind::RBRACKET,
            TokenKind::MINUS, TokenKind::PLUS, TokenKind::DOT, TokenKind::COMMA, TokenKind::SEMICOLON,
            TokenKind::MULTIPLY, TokenKind::DIVIDE
        ];
        let expected_lexemes: Vec<&str> = vec![
            "(", ")", "[", "]", "-", "+", ".", ",", ";", "*", "/"
        ];

        assert_eq!(scanner.tokens.len(), expected_tokens.len());
        for (i, token) in scanner.tokens.iter().enumerate() {
            assert_eq!(token.kind, expected_tokens[i]);
            assert_eq!(token.lexeme, expected_lexemes[i].to_string());
        }
    }

    #[test]
    fn test_multi_char_tokens() {
        let mut scanner = Scanner::new("// = != == >= <= && || !".to_string());
        scanner.scan();
        let expected_tokens: Vec<TokenKind> = vec![
            TokenKind::INTEGERDIVIDE, TokenKind::ASSIGN, TokenKind::NOTEQUAL, TokenKind::EQUAL, TokenKind::GREATEREQUAL,
            TokenKind::LESSEQUAL, TokenKind::AND, TokenKind::OR, TokenKind::NOT
        ];

        let expected_lexemes: Vec<&str> = vec![
            "//", "=", "!=", "==", ">=", "<=", "&&", "||", "!"
        ];

        println!("{:?}", scanner.tokens);

        assert_eq!(scanner.tokens.len(), expected_tokens.len());
        for (i, token) in scanner.tokens.iter().enumerate() {
            assert_eq!(token.kind, expected_tokens[i]);
            assert_eq!(token.lexeme, expected_lexemes[i].to_string());
        }
    }
    // TODO - test utility function because this is some garbage code

    #[test]
    fn test_comments() {
        let mut scanner = Scanner::new("() []; # this comment keeps going and should be ignored\n ==;".to_string());
        scanner.scan();
        let expected_tokens: Vec<TokenKind> = vec![
            TokenKind::LPAREN, TokenKind::RPAREN, TokenKind::LBRACKET, TokenKind::RBRACKET, TokenKind::SEMICOLON, TokenKind::EQUAL, TokenKind::SEMICOLON,
        ];
        let expected_lexemes: Vec<&str> = vec![
            "(", ")", "[", "]", ";", "==", ";"
        ];

        println!("{:?}", scanner.errors);
        assert_eq!(scanner.errors.has_errors(), false);

        println!("{:?}", scanner.tokens);
        assert_eq!(scanner.tokens.len(), expected_tokens.len());
        for (i, token) in scanner.tokens.iter().enumerate() {
            println!("{:?}", token);
            assert_eq!(token.kind, expected_tokens[i]);
            assert_eq!(token.lexeme, expected_lexemes[i].to_string());
        }
    }
    //
    // #[test]
    // fn test_string_literals() {
    //     // "let myString = \"Hello, world!\";"
    // }
    //
    // #[test]
    // fn test_number_literals() {
    //     // "let myNumber = 1234;\nlet myFloat = 12.34;"
    // }
    //
    // #[test]
    // fn test_identifiers() {
    //     // "let myVar = 1234;\nlet myFloat = 12.34;"
    // }
    //
    // #[test]
    // fn test_keywords() {
    //     // "return let if else for while struct use identifier"
    // }

}
