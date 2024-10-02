// metaprogramming the abstract syntax tree
// this is hacky and bad, but that is ok!
import fs from "fs";

const outputFile = "src/interpreter/ast.rs"
const types = [
  "Binary:Expr left,Token operator,Expr right",
  "Grouping:Expr expression",
  "Literal:String value,Token token",
  "Unary:Token operator,Expr right",
]

const str = `
use super::error::{InterpreterErrors, ErrorKind};
use super::scanner::*;

struct Expr
`





