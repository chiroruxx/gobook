package practice8_8

import (
	"bufio"
	"gobook/ch8/reverb1"
	"net"
	"time"
)

func HandleConn(c net.Conn) {
	ic := make(chan string)
	input := bufio.NewScanner(c)
	abort := make(chan struct{})

	go func(input *bufio.Scanner) {
		for input.Scan() {
			ic <- input.Text()
		}
		abort <- struct{}{}
	}(input)

	for {
		select {
		case text := <-ic:
			go reverb1.Echo(c, text, 1*time.Second)
		case <-time.After(10 * time.Second):
			c.Close()
			return
		case <-abort:
			c.Close()
		}
	}
}
