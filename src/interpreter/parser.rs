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


        let mut token = self.advance();

        while token.kind != TokenKind::EOF {
            match token.kind {
                TokenKind::LET => {
                    // let statement = self.parse_let_statement();
                    // program.push(Box::new(statement));
                },
                TokenKind::RETURN => {
                    // let statement = self.parse_return_statement();
                    // program.push(Box::new(statement));
                },
                _ => {
                    // let statement = self.parse_expression_statement();
                    // program.push(Box::new(statement));
                }
            }
        }

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

    fn parse_let_statement() -> Box<LetStatement> {

    }

}
