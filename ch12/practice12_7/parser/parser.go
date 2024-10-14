package parser

import (
	"io"

	"gobook/ch12/practice12_7/lexer"
)

func Parse(r io.Reader) (Node, error) {
	tokens, err := lexer.Lex(r)
	if err != nil {
		return nil, err
	}

	p := newParser()
	return p.parse(tokens)
}

type parser struct {
	current state
}

func newParser() parser {
	return parser{
		current: &initialState{},
	}
}

func (p *parser) parse(tokens []lexer.Token) (Node, error) {
	for _, t := range tokens {
		nextState, err := p.current.next(t)
		if err != nil {
			return nil, err
		}

		if nextState != nil {
			p.current = nextState
		}
	}
	return p.current.node(), nil
}
