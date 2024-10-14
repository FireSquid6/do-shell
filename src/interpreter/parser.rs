use super::scanner::*;
use super::tree::*;
use super::error::*;

type Program = Vec<Box<dyn Statement>>;

fn parse_into_tree(tokens: &Vec<Token>) -> Program {
    let program: Program = Vec::new();
    let errors: Vec<InterpreterError> = Vec::new();

    return program;

}
