use super::scanner::*;
use super::tree::*;
use super::error::*;

struct Parser {
    tokens: Vec<Token>,
    current: usize,
}


type Program = Vec<Box<dyn Statement>>;
impl Parser {
    fn parse(&mut self) -> (Program, InterpreterErrors) {
        let program: Program = Vec::new();
        let errors: InterpreterErrors = InterpreterErrors::new();


        return (program, errors);
    }

    fn advance(&mut self) -> Token {
        self.current += 1;
        let token = self.tokens.get(self.current);

        if token.is_none() {
            self.current -= 1;
            return Token::new(TokenKind::EOF, String::from(""), 0, 0);
        }

        return token.unwrap().clone();
    }

}
