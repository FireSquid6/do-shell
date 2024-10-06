use super::error::{InterpreterErrors, ErrorKind};

#[derive(Debug, Eq, PartialEq)]
pub enum TokenKind {
    LPAREN, RPAREN, LBRACE, LBRACKET, RBRACKET, RBRACE, COMMA, DOT, SEMICOLON, COLON,
    MINUS, PLUS, MULTIPLY, DIVIDE, INTEGERDIVIDE, MODULO, RAISETO,
    EQUAL, NOTEQUAL, GREATER, GREATEREQUAL, LESS, LESSEQUAL,
    AND, OR, NOT, ASSIGN,
    IDENTIFIER, STRING, NUMBER,
    IF, ELIF, ELSE, LET, FOR, USE, STRUCT, WHILE, RETURN, // MATCH, CASE
    EOF, UNKNOWN,

    NEWLINE
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

fn is_just_a_normal_number_no_bullshit(c: char) -> bool {
    return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9';
}

fn identifier_token_type(lexeme: String) -> TokenKind {
    match lexeme.as_str() {
        "if" => TokenKind::IF,
        "elif" => TokenKind::ELIF,
        "else" => TokenKind::ELSE,
        "let" => TokenKind::LET,
        "for" => TokenKind::FOR,
        "use" => TokenKind::USE,
        "struct" => TokenKind::STRUCT,
        "while" => TokenKind::WHILE,
        "return" => TokenKind::RETURN,
        _ => TokenKind::IDENTIFIER
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

    pub fn comb(self: &mut Scanner) {
        // We definitely DO NOT want to insert a semicolon if:
        // - the previous token is a comma or a semicolon
        // - the next valid character is a dot (probably chaining methods or properties)
        // - the next valid character is a comma (probably a list)
        // - the next valid character is a || or && (probably a logical operator)
        // - otherwise, we insert
        let mut i = 0;


        // ensure that the first token is not a NEWLINE
        if self.tokens.len() > 0 && self.tokens[0].kind == TokenKind::NEWLINE {
            self.tokens.remove(0);
        }

        while i < self.tokens.len() - 1 {
            if self.tokens[i].kind == TokenKind::NEWLINE {
                let prev = &self.tokens[i - 1];
                let next = &self.tokens[i + 1];

                // if the previous token is a comma or semicolon, then we just remove the newline 
                if prev.kind == TokenKind::COMMA || prev.kind == TokenKind::SEMICOLON {
                    self.tokens.remove(i);
                    continue;
                }

                // if the next token is a dot, comma, or logical operator, then we just remove the newline
                if next.kind == TokenKind::DOT || next.kind == TokenKind::COMMA || next.kind == TokenKind::AND || next.kind == TokenKind::OR {
                    self.tokens.remove(i);
                    continue;
                }
                

                self.tokens[i].kind = TokenKind::SEMICOLON;
            }

            i += 1;
        }

        // ensure that the last token is not a NEWLINE
        while self.tokens.len() > 0 && self.tokens[self.tokens.len() - 1].kind == TokenKind::NEWLINE {
            self.tokens.pop();
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


    fn scan_identifier(self: &mut Scanner, c: char) {
        let mut lexeme = c.to_string();

        while self.current < self.source.len() as u32 {
            let c = self.source.chars().nth(self.current as usize).unwrap_or('\0');
            if c.is_alphanumeric() || c == '_' {
                lexeme.push(c);
                self.advance();
            } else {
                break;
            }
        }

        let lexeme_clone = lexeme.clone();
        let kind = identifier_token_type(lexeme_clone);
        self.add_token(kind, lexeme);
    }

    fn scan_number(self: &mut Scanner, c: char) {
        let mut lexeme = c.to_string();
        let mut seen_dot = false;

        while self.current < self.source.len() as u32 {
            let c = self.source.chars().nth(self.current as usize).unwrap_or('\0');
            if c == '.' {
                if seen_dot {
                    self.errors.add_error("Invalid number contains more than one decimal".to_string(), ErrorKind::LEXER);
                    break;
                } else {
                    seen_dot = true;
                    lexeme.push(c);
                    self.advance();
                }
            }
            else if is_just_a_normal_number_no_bullshit(c) {
                lexeme.push(c);
                self.advance();
            } else {
                break;
            }
        }

        self.add_token(TokenKind::NUMBER, lexeme);
    }
    
    fn scan_string(self: &mut Scanner, c: char) {
        // c could be ' or "
        let start = c;
        let mut lexeme = String::new();
        let mut escape = false;

        while self.current < self.source.len() as u32 {
            let c = self.source.chars().nth(self.current as usize).unwrap_or('\0');

            if c == start {
                self.advance();
                if escape {
                    lexeme.push(c);
                } else {
                    break;
                }
            } else if c == '\n' {
                self.errors.add_error("Unterminated string".to_string(), ErrorKind::LEXER);
                break;
            } else {
                lexeme.push(c);
                self.advance();
            }

            if c == '\\' {
                escape = true;
            } else {
                escape = false;
            }
        }

        self.add_token(TokenKind::STRING, lexeme);
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

            '"' | '\'' => {
                self.scan_string(c);
            }

            // always skip whitespace
            ' ' | '\t' => {}
            '\n' => {

                //
                // Other valid way:
                // - convert all \ns in a row to a single NEWLINE token
                // - after we scan, we can insert semicolons where needed and remove all NEWLINE
                // tokens

                self.add_token(TokenKind::NEWLINE, "\n".to_string());
            }

            _ => {
                if is_just_a_normal_number_no_bullshit(c) {
                    self.scan_number(c);
                } else if c.is_alphabetic() {
                    self.scan_identifier(c);
                } else {
                    self.errors.add_error("Unexpected token".to_string(), ErrorKind::LEXER);
                }
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
        scanner.comb();

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
        lexer_test("()[]-+.,;*/".to_string(), vec![
            TokenKind::LPAREN, TokenKind::RPAREN, TokenKind::LBRACKET, TokenKind::RBRACKET,
            TokenKind::MINUS, TokenKind::PLUS, TokenKind::DOT, TokenKind::COMMA, TokenKind::SEMICOLON,
            TokenKind::MULTIPLY, TokenKind::DIVIDE
        ], vec![
            "(", ")", "[", "]", "-", "+", ".", ",", ";", "*", "/"
        ], vec![]);
    }

    #[test]
    fn test_multi_char_tokens() {
        lexer_test("== != >= <= && ||".to_string(), vec![
            TokenKind::EQUAL, TokenKind::NOTEQUAL, TokenKind::GREATEREQUAL, TokenKind::LESSEQUAL, TokenKind::AND, TokenKind::OR
        ], vec![
            "==", "!=", ">=", "<=", "&&", "||"
        ], vec![]);
    }

    #[test]
    fn test_comments() {
        lexer_test("() []; # this comment keeps going and should be ignored\n ==;".to_string(), vec![
            TokenKind::LPAREN, TokenKind::RPAREN, TokenKind::LBRACKET, TokenKind::RBRACKET, TokenKind::SEMICOLON, TokenKind::EQUAL, TokenKind::SEMICOLON,
        ], vec![
            "(", ")", "[", "]", ";", "==", ";"
        ], vec![]);
    }
    
    #[test]
    fn test_string_literals() {
        lexer_test("let myString = \"Hello, world!\";".to_string(), vec![
            TokenKind::LET, TokenKind::IDENTIFIER, TokenKind::ASSIGN, TokenKind::STRING, TokenKind::SEMICOLON
        ], vec![
            "let", "myString", "=", "Hello, world!", ";"
        ], vec![]);

        lexer_test("let myString = 'Hello, world!';".to_string(), vec![
            TokenKind::LET, TokenKind::IDENTIFIER, TokenKind::ASSIGN, TokenKind::STRING, TokenKind::SEMICOLON
        ], vec![
            "let", "myString", "=", "Hello, world!", ";"
        ], vec![]);
    }

    #[test]
    fn test_string_literal_with_escape() {
        lexer_test("let myString = \"Hello, \\n world!\";".to_string(), vec![
            TokenKind::LET, TokenKind::IDENTIFIER, TokenKind::ASSIGN, TokenKind::STRING, TokenKind::SEMICOLON
        ], vec![
            "let", "myString", "=", "Hello, \\n world!", ";"
        ], vec![]);

        let raw_string = r#"let myString = "Hello, \" world!";"#;

        lexer_test(raw_string.to_string(), vec![
            TokenKind::LET, TokenKind::IDENTIFIER, TokenKind::ASSIGN, TokenKind::STRING, TokenKind::SEMICOLON
        ], vec![
            "let", "myString", "=", "Hello, \\\" world!", ";"
        ], vec![]);
    }


    #[test]
    fn test_identifiers() {
        lexer_test("let myVar = 1234; let myFloat = 12.34;".to_string(), vec![
            TokenKind::LET, TokenKind::IDENTIFIER, TokenKind::ASSIGN, TokenKind::NUMBER, TokenKind::SEMICOLON,
            TokenKind::LET, TokenKind::IDENTIFIER, TokenKind::ASSIGN, TokenKind::NUMBER, TokenKind::SEMICOLON,
        ], vec![
            "let", "myVar", "=", "1234", ";",
            "let", "myFloat", "=", "12.34", ";"
        ], vec![]);
    }

    #[test]
    fn test_keywords() {
        lexer_test("return let if else for while struct use identifier".to_string(), vec![
            TokenKind::RETURN, TokenKind::LET, TokenKind::IF, TokenKind::ELSE, TokenKind::FOR, TokenKind::WHILE, TokenKind::STRUCT, TokenKind::USE, TokenKind::IDENTIFIER
        ], vec![
            "return", "let", "if", "else", "for", "while", "struct", "use", "identifier"
        ], vec![]);
    }

    #[test]
    fn test_combing() {
        lexer_test("let a = 5\n let b = 6\n".to_string(), vec![
            TokenKind::LET, TokenKind::IDENTIFIER, TokenKind::ASSIGN, TokenKind::NUMBER, TokenKind::SEMICOLON,
            TokenKind::LET, TokenKind::IDENTIFIER, TokenKind::ASSIGN, TokenKind::NUMBER, TokenKind::SEMICOLON,
        ], vec![
            "let", "a", "=", "5", "\n",
            "let", "b", "=", "6", "\n",
        ], vec![]);

    }
}
