package parser

import (
	"bytes"

	"gobook/ch12/practice12_7/lexer"
)

type Node interface {
	Value() string
}

type NumberNode struct {
	token *lexer.NumberToken
}

func (n *NumberNode) Value() string {
	return n.token.String()
}

type SymbolNode struct {
	tokens []lexer.Token
}

func (n *SymbolNode) Value() string {
	var buf bytes.Buffer
	for _, token := range n.tokens {
		buf.WriteString(token.String())
	}
	return buf.String()
}

type StringNode struct {
	tokens []lexer.Token
}

func (n *StringNode) Value() string {
	var buf bytes.Buffer
	for _, token := range n.tokens {
		buf.WriteString(token.String())
	}
	return buf.String()
}

type ListNode struct {
	nodes []Node
}

func (n *ListNode) Value() string {
	return ""
}

func (n *ListNode) Nodes() []Node {
	return n.nodes
}
