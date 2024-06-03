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

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() []rune { return ls.Token.Literal }
