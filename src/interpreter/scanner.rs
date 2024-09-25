enum TokenKind {
    LPAREN, RPAREN, LBRACE, RBRACE, COMMA, DOT, SEMICOLON,
    MINUS, PLUS, MULTIPLY, DIVIDE, INTEGERDIVIDE, MODULO, RAISETO,
    EQUAL, NOTEQUAL, GREATER, GREATEREQUAL, LESS, LESSEQUAL,
    AND, OR, NOT,
    IDENTIFIER, STING, NUMBER,
    IF, ELIF, ELSE, LET, FOR, WHILE, RETURN, // MATCH, CASE
    EOF,
}


pub struct Token {
    kind: TokenKind,
    lexeme: String,
    line: u32,
    column: u32,
}


pub struct Scanner {


}


fn new_scanner() -> Scanner {
    Scanner {}
}


#[cfg(test)]
fn test_scanner() {
    let scanner = new_scanner();

}



