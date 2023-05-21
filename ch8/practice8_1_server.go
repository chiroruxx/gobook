package main

import (
	"gobook/ch8/clock1"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("server: port is not found.")
	}

	port := os.Args[1]
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("invalid port number")
	}

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go clock1.HandleConn(conn)
	}
}
