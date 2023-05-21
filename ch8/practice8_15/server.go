package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

type client struct {
	who string
	ch  chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			list := whos(clients)
			if len(list) > 0 {
				cli.ch <- strings.Join(list, ", ") + " are here!"
			}
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func whos(clients map[client]bool) []string {
	var list []string
	for cli := range clients {
		list = append(list, cli.who)
	}
	return list
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 10)
	go clientWriter(conn, ch)

	ch <- "What is your name?"
	input := bufio.NewScanner(conn)
	input.Scan()
	who := input.Text()

	ch <- "You are " + who
	messages <- who + " has arrived"
	cli := client{
		who,
		ch,
	}
	entering <- cli

	updated := make(chan struct{})
	done := make(chan struct{})
	go timeout(conn, updated, done)

	for input.Scan() {
		messages <- who + ": " + input.Text()
		updated <- struct{}{}
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
	close(done)
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func timeout(conn net.Conn, updated, done chan struct{}) {
	const d = 5 * time.Minute
	ticker := time.NewTicker(d)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			conn.Close()
			ticker.Stop()
			return
		case <-updated:
			ticker.Reset(d)
		}
	}
}
