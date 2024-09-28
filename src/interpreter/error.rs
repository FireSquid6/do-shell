#[derive(Debug, PartialEq, Eq)]
pub enum ErrorKind { LEXER, PARSER, RUNTIME
}

#[derive(Debug, PartialEq, Eq)]
pub struct InterpreterError {
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

#[derive(Debug, PartialEq, Eq)]
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

    // testing utility funciton when we expect
    pub fn expect_errors(self: &InterpreterErrors, expected: Vec<&str>) -> bool {
        if self.errors.len() != expected.len() {
            return false
        }

        for i in 0..self.errors.len() {
            if self.errors[i].message != expected[i] {
                return false
            }
        }

        return true
    }
}
