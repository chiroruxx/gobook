package practice8_4

import (
	"bufio"
	"gobook/ch8/reverb1"
	"net"
	"sync"
	"time"
)

func HandleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup

	for input.Scan() {
		wg.Add(1)
		go func(input *bufio.Scanner) {
			reverb1.Echo(c, input.Text(), 1*time.Second)
			wg.Done()
		}(input)
	}

	wg.Wait()
}
