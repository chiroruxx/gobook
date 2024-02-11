package practice9_4

import (
	"sync"
)

var done <-chan struct{}

func SetDone(ch <-chan struct{}) {
	done = ch
}

type chain struct {
	id   int
	prev <-chan string
	next chan<- string
}

func newChain(id int, prev, next chan string) *chain {
	return &chain{
		id:   id,
		prev: prev,
		next: next,
	}
}

func MakeGoroutines(num int, wakeupGroup, closeGroup *sync.WaitGroup) (first, last chan string) {
	var prev chan string
	var next chan string

	for i := 0; i < num; i++ {
		if prev == nil {
			prev = make(chan string)
		}

		next = make(chan string)

		c := newChain(i+1, prev, next)
		makeGoroutine(c, wakeupGroup, closeGroup)

		if first == nil {
			first = prev
		}

		prev = next
	}

	return first, next
}

func makeGoroutine(c *chain, wakeupGroup, closeGroup *sync.WaitGroup) {
	wakeupGroup.Add(1)
	closeGroup.Add(1)
	go routine(c, wakeupGroup, closeGroup)
}

func routine(c *chain, wakeupGroup, closeGroup *sync.WaitGroup) {
	if c.prev == nil {
		panic("prev is nil!!")
	}

	printf("Start goroutine #%d\n", c.id)
	wakeupGroup.Done()

	for {
		var finish bool
		select {
		case <-done:
			finish = true
		case s := <-c.prev:
			if c.next != nil {
				printf("#%d: Receive value %s\n", c.id, s)
				c.next <- s
			}
		}
		if finish {
			break
		}
	}

	printf("Finish goroutine #%d\n", c.id)
	closeGroup.Done()
}
