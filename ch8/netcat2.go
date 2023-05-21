package main

import (
	"gobook/ch8/netcat1"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go netcat1.MustCopy(os.Stdout, conn)
	netcat1.MustCopy(conn, os.Stdin)
}
