package practice7_16

import "fmt"

type expr interface {
	evaluate() float64
}

type numberExpr struct {
	n float64
}

func (e numberExpr) evaluate() float64 {
	return e.n
}

type operationExpr struct {
	left, right expr
	op          byte
}

func (e operationExpr) evaluate() float64 {
	switch e.op {
	case '+':
		return e.left.evaluate() + e.right.evaluate()
	case '-':
		return e.left.evaluate() - e.right.evaluate()
	case '*':
		return e.left.evaluate() * e.right.evaluate()
	case '/':
		return e.left.evaluate() / e.right.evaluate()
	default:
		panic(e.op)
	}
}

type parserStore struct {
	left expr
	op   byte
}

func (s *parserStore) isEmpty() bool {
	return s.left == nil && s.op == 0
}

type parserState interface {
	validate(t token) error
	action(t token) expr
	next(t token) parserState
	store() *parserStore
}

type initState struct {
	sp *parserStore
}

func (s initState) validate(t token) error {
	if !t.isNumberToken() {
		return fmt.Errorf("init state should receive number token, got %x", t.op)
	}
	if !s.sp.isEmpty() {
		return fmt.Errorf("init state requires empty store")
	}

	return nil
}

func (s initState) action(t token) expr {
	e := numberExpr{n: t.number}
	s.sp.left = e
	return nil
}

func (s initState) store() *parserStore {
	return s.sp
}

func (s initState) next(_ token) parserState {
	return operableState{sp: s.sp}
}

type operableState struct {
	sp *parserStore
}

func (s operableState) validate(t token) error {
	if t.isNumberToken() {
		return fmt.Errorf("operable state should receive operation token, got %g", t.number)
	}
	if s.sp.left == nil {
		return fmt.Errorf("operable state requires left")
	}

	return nil
}

func (s operableState) action(t token) expr {
	s.sp.op = t.op

	return nil
}

func (s operableState) next(_ token) parserState {
	return numberableState{sp: s.sp}
}

func (s operableState) store() *parserStore {
	return s.sp
}

type numberableState struct {
	sp *parserStore
}

func (s numberableState) validate(t token) error {
	if !t.isNumberToken() {
		return fmt.Errorf("numberable state should receive number token, got %g", t.number)
	}
	if s.sp.left == nil {
		return fmt.Errorf("numberable state requires left")
	}
	if s.sp.op == 0 {
		return fmt.Errorf("numberable state requires op")
	}

	return nil
}

func (s numberableState) action(t token) expr {
	right := numberExpr{n: t.number}
	e := operationExpr{
		left:  s.sp.left,
		op:    s.sp.op,
		right: right,
	}
	s.sp.left = e

	return e
}

func (s numberableState) next(_ token) parserState {
	return operableState{sp: s.sp}
}

func (s numberableState) store() *parserStore {
	return s.sp
}

type Parser struct {
	state parserState
}

func NewParser() *Parser {
	store := new(parserStore)
	state := initState{sp: store}
	parser := Parser{state: state}
	return &parser
}

func (p Parser) parse(tokens []token) expr {
	for _, t := range tokens {
		err := p.state.validate(t)
		if err != nil {
			continue
		}

		p.state.action(t)

		p.state = p.state.next(t)
	}

	result := p.state.store().left
	if result == nil {
		result = numberExpr{n: 0}
	}

	return result
}
