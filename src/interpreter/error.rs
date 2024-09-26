#[derive(Debug)]
pub enum ErrorKind { LEXER, PARSER, RUNTIME
}

#[derive(Debug)]
struct InterpreterError {
    kind: ErrorKind,
    message: String,


}


impl InterpreterError {
    fn new(kind: ErrorKind, message: String) -> InterpreterError {
        return InterpreterError {
            kind,
            message,
        }

    }
}

#[derive(Debug)]
pub struct InterpreterErrors {
    errors: Vec<InterpreterError>
}

impl InterpreterErrors {
    pub fn new() -> InterpreterErrors {
        return InterpreterErrors {
            errors: vec![]
        }
    }

    pub fn add_error(self: &mut InterpreterErrors, error: String, kind: ErrorKind) {
        self.errors.push(InterpreterError::new(kind, error))
    }

    pub fn has_errors(self: &InterpreterErrors) -> bool {
        return self.errors.len() > 0
    }
}
