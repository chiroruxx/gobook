package practice9_3

type entry struct {
	res   result
	ready chan struct{}
}

type Func func(key string, cancel <-chan struct{}) (any, error)

type result struct {
	value any
	err   error
}

type request struct {
	key      string
	cancel   <-chan struct{}
	response chan<- result
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{
		requests: make(chan request),
	}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, cancel <-chan struct{}) (any, error) {
	response := make(chan result)

	if canceled(cancel) {
		return nil, nil
	}
	memo.requests <- request{key, cancel, response}

	select {
	case <-cancel:
		return nil, nil
	case res := <-response:
		return res.value, res.err
	}
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	try := make(map[string]bool)

	for req := range memo.requests {
		e := cache[req.key]
		tried := try[req.key]
		if !tried {
			e = &entry{
				ready: make(chan struct{}),
			}

			try[req.key] = true
			go e.call(f, req.key, req.cancel)
			select {
			case <-req.cancel:
				// do nothing
			case <-e.ready:
				cache[req.key] = e
			}
		}
		go e.deliver(req)
	}
}

func (e *entry) call(f Func, key string, cancel <-chan struct{}) {
	e.res.value, e.res.err = f(key, cancel)
	close(e.ready)
}

func (e *entry) deliver(req request) {
	select {
	case <-req.cancel:
		req.response <- result{}
	case <-e.ready:
		req.response <- e.res
	}
}

func canceled(cancel <-chan struct{}) bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}
