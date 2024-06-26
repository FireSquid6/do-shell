package tree

import (
	"github.com/firesquid6/do-shell/token"
	"strings"
)

type Node interface {
	TokenLiteral() []rune
	String() string
  // Evaluate(ctx *interpreter.InterpreterContext) 
}

type Statement interface {
	Node
	statementNode() // dummy method to help the compiler
}

type Expression interface {
	Node
	expressionNode() // dummy method to help the compiler

}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator []rune
	Right    Expression
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() []rune { return b.Token.Literal }
func (b *Boolean) String() string {
	return string(b.Token.Literal)
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() []rune { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	return "(" + ie.Left.String() + " " + string(ie.Operator) + " " + ie.Right.String() + ")"
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	builder := strings.Builder{}

	for _, s := range p.Statements {
		_, _ = builder.WriteString(s.String())
	}

	return builder.String()
}

// TODO: parse a float instaed of an int
type NumberLiteral struct {
	Token token.Token
	Value float64
}

func (il *NumberLiteral) expressionNode()      {}
func (il *NumberLiteral) TokenLiteral() []rune { return il.Token.Literal }
func (il *NumberLiteral) String() string       { return string(il.Token.Literal) }

type Identifier struct {
	Token token.Token
	Value []rune
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() []rune { return i.Token.Literal }
func (i *Identifier) String() string       { return string(i.Value) }

// TODO: string and command expressions

type PrefixExpression struct {
	Token    token.Token
	Operator []rune
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() []rune { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	return "(" + string(pe.Operator) + pe.Right.String() + ")"
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
func (ls *LetStatement) String() string {
	builder := strings.Builder{}
	builder.WriteString(string(ls.TokenLiteral()))
	builder.WriteString(" ")
	builder.WriteString(string(ls.Name.Value))
	builder.WriteString(" = ")
	builder.WriteString(ls.Expression.String())
	builder.WriteString(";")
	return builder.String()
}

type ReturnStatement struct {
	Token      token.Token
	Expression Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() []rune { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	return string(rs.TokenLiteral()) + " " + rs.Expression.String() + ";"
}

type CallExpression struct {
	Token     token.Token
	Function  Identifier
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() []rune { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	builder := strings.Builder{}
	builder.WriteString(ce.Function.String())
	builder.WriteString("(")
	for i, arg := range ce.Arguments {
		builder.WriteString(arg.String())
		if i < len(ce.Arguments)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(")")
	return builder.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (s *ExpressionStatement) statementNode()       {}
func (s *ExpressionStatement) TokenLiteral() []rune { return s.Token.Literal }
func (s *ExpressionStatement) String() string {
	return s.Expression.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() []rune { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	builder := strings.Builder{}
	for _, s := range bs.Statements {
		builder.WriteString(s.String())
	}
	return builder.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() []rune { return ie.Token.Literal }
func (ie *IfExpression) String() string {
  builder := strings.Builder{}
  builder.WriteString("if ")
  builder.WriteString(ie.Condition.String())
  builder.WriteString(" ")
  builder.WriteString(ie.Consequence.String())
  if ie.Alternative != nil {
    builder.WriteString(" else ")
    builder.WriteString(ie.Alternative.String())
  }
  return builder.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Statements *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) String() string {
	builder := strings.Builder{}
	builder.WriteString(string(fl.TokenLiteral()))
	builder.WriteString("(")
	for i, p := range fl.Parameters {
		builder.WriteString(string(p.Value))
		if i < len(fl.Parameters)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(") {\n")
	builder.WriteString(fl.Statements.String())
	builder.WriteString("}")
	return builder.String()
}
func (fl *FunctionLiteral) TokenLiteral() []rune { return fl.Token.Literal }

// TODO: if, for, while, else, elif statements
