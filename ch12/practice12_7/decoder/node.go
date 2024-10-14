package decoder

import (
	"errors"
	"strconv"

	"gobook/ch12/practice12_7/parser"
)

type listNode struct {
	*parser.ListNode
}

func (n *listNode) Nodes() []*node {
	var nodes []*node
	for _, on := range n.ListNode.Nodes() {
		nodes = append(nodes, newNode(on))
	}
	return nodes
}

func (n *listNode) Len() int {
	return len(n.Nodes())
}

type symbolNode struct {
	*parser.SymbolNode
}

type numberNode struct {
	*parser.NumberNode
}

type stringNode struct {
	*parser.StringNode
}

type node struct {
	parser.Node
}

func newNode(origin parser.Node) *node {
	return &node{
		Node: origin,
	}
}

func (n *node) listNode() (*listNode, error) {
	ln, ok := n.Node.(*parser.ListNode)
	if !ok {
		return nil, errors.New("expected list")
	}
	return &listNode{ln}, nil
}

func (n *node) symbolNode() (*symbolNode, error) {
	sn, ok := n.Node.(*parser.SymbolNode)
	if !ok {
		return nil, errors.New("expected symbol")
	}
	return &symbolNode{sn}, nil
}

func (n *node) numberNode() (*numberNode, error) {
	nn, ok := n.Node.(*parser.NumberNode)
	if !ok {
		return nil, errors.New("expected number")
	}
	return &numberNode{nn}, nil
}

func (n *node) stringNode() (*stringNode, error) {
	sn, ok := n.Node.(*parser.StringNode)
	if !ok {
		return nil, errors.New("expected string")
	}
	return &stringNode{sn}, nil
}

func (n *node) Number() (int, error) {
	nn, err := n.numberNode()
	if err != nil {
		return 0, err
	}

	number, err := strconv.Atoi(nn.Value())
	if err != nil {
		return 0, err
	}
	return number, nil
}

func (n *node) String() (string, error) {
	sn, err := n.stringNode()
	if err != nil {
		return "", err
	}

	return sn.Value(), nil
}

type nodeStruct struct {
	fields map[string]*node
}

type structFieldItem struct {
	name  *symbolNode
	value *node
}

func (n *node) newStructFieldItem() (*structFieldItem, error) {
	ln, err := n.listNode()
	if err != nil {
		return nil, err
	}
	if ln.Len() != 2 {
		return nil, errors.New("invalid item length")
	}

	nameNode, err := ln.Nodes()[0].symbolNode()
	if err != nil {
		return nil, err
	}

	valueNode := ln.Nodes()[1]
	if !valueNode.canConvertFieldValue() {
		return nil, errors.New("cannot convert struct field")
	}

	return &structFieldItem{
		name:  nameNode,
		value: valueNode,
	}, nil
}

func (n *node) Struct() (*nodeStruct, error) {
	ln, err := n.listNode()
	if err != nil {
		return nil, err
	}

	var res nodeStruct
	res.fields = make(map[string]*node)

	children := ln.Nodes()
	for _, child := range children {
		item, err := child.newStructFieldItem()
		if err != nil {
			return nil, err
		}
		name := item.name.Value()
		res.fields[name] = item.value
	}
	return &res, nil
}

func (n *node) canConvertFieldValue() bool {
	if _, err := n.numberNode(); err == nil {
		return true
	}
	if _, err := n.stringNode(); err == nil {
		return true
	}
	return false
}
