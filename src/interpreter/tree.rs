use super::scanner::*;

pub trait Statement {
    fn execute(&self);
}

pub struct LetStatement {
    identifier: String,
    expression: dyn Expression,
}

pub struct ReturnStatement {
    expression: dyn Expression,
}

pub struct ExpressionStatement {
    expression: dyn Expression,
}


pub trait Expression {
    fn evaluate(&self);
}


pub struct LiteralExpression {
    value: dyn Value,
}

pub struct IdentifierExpression {
    identifier: String,
}

pub struct PrefixExpression {
    operator: TokenKind,
    right: dyn Expression,
}

pub struct InfixExpression {
    left: Box<dyn Expression>,
    operator: TokenKind,
    right: dyn Expression,
}

pub trait Value {
    fn to_string(&self) -> String;
    fn add(&self, other: &dyn Value) -> dyn Value;
    fn subtract(&self, other: &dyn Value) -> dyn Value;
    fn multiply(&self, other: &dyn Value) -> dyn Value;
    fn divide(&self, other: &dyn Value) -> dyn Value;
    fn modulo(&self, other: &dyn Value) -> dyn Value;

    fn raiseto(&self, other: &dyn Value) -> dyn Value;
    fn integer_divide(&self, other: &dyn Value) -> dyn Value;
    fn negate(&self) -> dyn Value;

    fn equal(&self, other: &dyn Value) -> bool;
    fn not_equal(&self, other: &dyn Value) -> bool;
    fn greater_than(&self, other: &dyn Value) -> bool;
    fn greater_than_or_equal(&self, other: &dyn Value) -> bool;
    fn less_than(&self, other: &dyn Value) -> bool;
    fn less_than_or_equal(&self, other: &dyn Value) -> bool;
}
