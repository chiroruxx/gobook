package reverb2

import (
	"bufio"
	"gobook/ch8/reverb1"
	"net"
	"time"
)

func HandleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go reverb1.Echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
