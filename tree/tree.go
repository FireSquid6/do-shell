package tree

import "github.com/firesquid6/do-shell/token"

type Node interface {
	TokenLiteral() []rune
}

type Statement interface {
	Node
	statementNode() // dummy method to help the compiler
}

type Expression interface {
	Node
	expressionNode() // dummy method to help the compiler
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

func (p *Program) TokenLiteral() []rune {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return []rune{}
	}
}

type LetStatement struct {
	Token      token.Token
	Name       *Identifier
	Expression Expression
}
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() []rune { return ls.Token.Literal }

type FnStatement struct {
  Token      token.Token
  Name       *Identifier
  Parameters []*Identifier
  Body       []*Statement
}
func (s *FnStatement) statementNode() {}
func (s *FnStatement) TokenLiteral() []rune { return s.Token.Literal }

type ExpressionStatement struct {
  Token      token.Token
  Expression Expression
}
func (s *ExpressionStatement) statementNode() {}
func (s *ExpressionStatement) TokenLiteral() []rune { return s.Token.Literal }

type ReturnStatement struct {
  Token      token.Token
  ReturnValue Expression
}
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() []rune { return rs.Token.Literal }

type IfStatement struct {
  Token       token.Token
  Expression Expression
  Body []*Statement
}
func (is *IfStatement) statementNode() {}
func (is *IfStatement) TokenLiteral() []rune { return is.Token.Literal }

type ElifStatement struct {
  Token       token.Token
  Expression Expression
  Body []*Statement
}
func (es *ElifStatement) statementNode() {}
func (es *ElifStatement) TokenLiteral() []rune { return es.Token.Literal }

type ElseStatement struct {
  Token token.Token
  Body []*Statement
}
func (es *ElseStatement) statementNode() {}
func (es *ElseStatement) TokenLiteral() []rune { return es.Token.Literal }





