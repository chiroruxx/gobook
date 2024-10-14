package parser

import (
	"fmt"

	"gobook/ch12/practice12_7/lexer"
)

type state interface {
	next(t lexer.Token) (state, error)
	node() Node
}

type initialState struct{}

func (s *initialState) next(t lexer.Token) (state, error) {
	switch x := t.(type) {
	case *lexer.NumberToken:
		n := numberState{
			n: &NumberNode{x},
		}
		return &n, nil
	case *lexer.StringToken:
		n := symbolState{
			tokens: []lexer.Token{t},
		}
		return &n, nil
	case *lexer.QuoteToken:
		n := stringState{
			startToken: x,
		}
		return &n, nil
	case *lexer.ListStartToken:
		n := listState{
			startToken:    x,
			internalState: &initialState{},
		}
		return &n, nil
	default:
		return nil, fmt.Errorf("unexpected token: %v", t)
	}
}

func (s *initialState) node() Node {
	return nil
}

type numberState struct {
	n *NumberNode
}

func (s *numberState) next(_ lexer.Token) (state, error) {
	return nil, fmt.Errorf("cannot move any states")
}

func (s *numberState) node() Node {
	return s.n
}

type stringState struct {
	startToken *lexer.QuoteToken
	tokens     []lexer.Token
	endToken   *lexer.QuoteToken
}

func (s *stringState) next(t lexer.Token) (state, error) {
	if s.valid() {
		return nil, fmt.Errorf("unexpected token for vali: %v", t)
	}

	switch x := t.(type) {
	case *lexer.QuoteToken:
		s.endToken = x
		return s, nil
	case *lexer.NumberToken, *lexer.StringToken, *lexer.EscapeToken, *lexer.EscapedStringToken, *lexer.ListSeparatorToken:
		s.tokens = append(s.tokens, x)
		return s, nil
	default:
		return nil, fmt.Errorf("unexpected token: %v", t)
	}
}

func (s *stringState) valid() bool {
	return s.endToken != nil
}

func (s *stringState) node() Node {
	if !s.valid() {
		return nil
	}

	return &StringNode{
		tokens: s.tokens,
	}
}

type symbolState struct {
	tokens []lexer.Token
}

func (s *symbolState) next(t lexer.Token) (state, error) {
	switch x := t.(type) {
	case *lexer.NumberToken, *lexer.StringToken, *lexer.EscapeToken, *lexer.EscapedStringToken:
		s.tokens = append(s.tokens, x)
		return s, nil
	default:
		return nil, fmt.Errorf("unexpected token: %v", t)
	}
}

func (s *symbolState) node() Node {
	return &SymbolNode{
		tokens: s.tokens,
	}
}

type listState struct {
	startToken    *lexer.ListStartToken
	endToken      *lexer.ListEndToken
	internalState state
	nodes         []Node
}

func (s *listState) next(t lexer.Token) (state, error) {
	if s.valid() {
		return nil, fmt.Errorf("unexpected token for valid: %v", t)
	}

	if err := s.nextInternal(t); err != nil {
		switch s.internalState.(type) {
		case *initialState:
			return nil, err
		}

		switch x := t.(type) {
		case *lexer.ListSeparatorToken:
			s.flushInternal()
			return s, nil
		case *lexer.ListEndToken:
			s.endToken = x
			s.flushInternal()
			return s, nil
		default:
			return nil, err
		}
	}
	return s, nil
}

func (s *listState) valid() bool {
	return s.endToken != nil
}

func (s *listState) node() Node {
	if !s.valid() {
		return nil
	}

	return &ListNode{
		nodes: s.nodes,
	}
}

func (s *listState) nextInternal(t lexer.Token) error {
	nextState, err := s.internalState.next(t)
	if err != nil {
		return err
	}
	s.internalState = nextState
	return nil
}

func (s *listState) flushInternal() {
	n := s.internalState.node()
	s.nodes = append(s.nodes, n)
	s.internalState = &initialState{}
}
