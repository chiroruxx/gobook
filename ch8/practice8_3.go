package main

import (
	"gobook/ch8/netcat1"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	tcpConn := conn.(*net.TCPConn)

	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, tcpConn)
		log.Println("done")
		done <- struct{}{}
	}()
	netcat1.MustCopy(tcpConn, os.Stdin)
	tcpConn.CloseWrite()
	<-done
	tcpConn.Close()
}
