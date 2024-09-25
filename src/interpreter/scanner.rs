enum TokenType {
    LPAREN,
    RPAREN,
    LBRACE,
    RBRACE,
    COMMA,
    DOT,
    SEMICOLON,

    MINUS,
    PLUS,
    MULTIPLY,
    DIVIDE,
    INTEGERDIVIDE,
    MODULO,
    RAISETO,

    EQUAL,
    NOTEQUAL,
    GREATER,
    GREATEREQUAL,
    LESS,
    LESSEQUAL,

    AND,
    OR,
    NOT,

    IDENTIFIER,
    STING,
    NUMBER,

    IF,
    ELIF,
    ELSE,
    // MATCH,
    // CASE,
    LET,
    FOR,
    WHILE,
    RETURN,

    EOF,
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



