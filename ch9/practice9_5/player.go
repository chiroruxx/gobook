package practice9_5

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var done chan struct{}

type message string

const (
	ping message = "ping"
	pong message = "pong"
)

type player struct {
	name      string
	receiver  chan message
	sender    chan message
	sendCount int
}

func newPlayers() (player1, player2 *player) {
	ch1 := make(chan message)
	ch2 := make(chan message)

	player1 = &player{
		name:     "player1",
		receiver: ch1,
		sender:   ch2,
	}
	player2 = &player{
		name:     "player2",
		receiver: ch2,
		sender:   ch1,
	}
	return
}

func (p *player) fire() {
	p.send(ping)
}

func (p *player) receive() {
	wg.Done()
	for {
		select {
		case <-done:
			return
		case s := <-p.receiver:
			m := ping
			if s == ping {
				m = pong
			}
			p.send(m)
		}
	}
}

func (p *player) send(m message) {
	p.sendCount++
	p.sender <- m
}

func Start() {
	done = make(chan struct{})
	player1, player2 := newPlayers()

	wg.Add(1)
	go player1.receive()
	wg.Add(1)
	go player2.receive()

	wg.Wait()

	player1.fire()
	time.Sleep(1 * time.Second)
	close(done)
	fmt.Printf("player1: %d\n", player1.sendCount)
	fmt.Printf("player2: %d\n", player2.sendCount)
	fmt.Printf("total: %d\n", player1.sendCount+player2.sendCount)
}
