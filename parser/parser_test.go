package parser_test

import (
  "testing"
  "reflect"
  "github.com/firesquid6/do-shell/parser"
  "github.com/firesquid6/do-shell/lexer"
  "github.com/firesquid6/do-shell/tree"
)



func TestLetStatements(t *testing.T) {
  input := `
  let x = 5;
  let y = 10;
  let foobar = 838383;
  `

  l := lexer.Lexer{}
  l.LexText(input)

  p := parser.New(&l)

  program := p.ParseProgram()

  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }

  if len(program.Statements) != 3 {
    t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
  }

  tests := []struct {expectedIdentifier string}{
    {"x"},
    {"y"},
    {"foobar"},
  }

  for i, tt := range tests {
    stmt := program.Statements[i]

    if !testLetStatement(t, stmt, tt.expectedIdentifier) {
      return
    }
    
  }
}


func testLetStatement(t *testing.T, s tree.Statement, name string) bool {
  if !reflect.DeepEqual(s.TokenLiteral(), []rune{'l', 'e', 't'}) {
    t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
    return false
  }

  letStmt, ok := s.(*tree.LetStatement)
  if !ok {
    t.Errorf("s not *tree.LetStatement. got=%T", s)
  }

  if letStmt.Name.Value != name {
    t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
  }


  return true
}
