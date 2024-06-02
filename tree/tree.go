package tree


type Node interface {
  TokenLiteral() string
}

type Statement interface {
  Node
  statementNode()  // dummy method to help the compiler
}

type Expression interface {
  Node
  expressionNode()  // dummy method to help the compiler
}


type Program struct {
  Statements []Statement
}


func (p *Program) TokenLiteral() string {
  if len(p.Statements) > 0 {
    return p.Statements[0].TokenLiteral()
  } else {
    return ""
  }
}
