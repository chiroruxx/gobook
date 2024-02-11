package memo4

import "sync"

type entry struct {
	res   result
	ready chan struct{}
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type Func func(key string) (any, error)

type result struct {
	value any
	err   error
}

func (memo *Memo) Get(key string) (any, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{
			ready: make(chan struct{}),
		}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready)
	} else {
		memo.mu.Unlock()

		<-e.ready
	}
	return e.res.value, e.res.err
}
