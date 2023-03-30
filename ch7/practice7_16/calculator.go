package practice7_16

type Calculator struct {
	t *Tokenizer
	p *Parser
}

func NewCalculator() *Calculator {
	c := Calculator{
		t: NewTokenizer(),
		p: NewParser(),
	}

	return &c
}

func (c Calculator) Calc(input string) float64 {
	tokens := c.t.tokenize(input)
	e := c.p.parse(tokens)
	return e.evaluate()
}
