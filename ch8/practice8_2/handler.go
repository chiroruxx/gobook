package practice8_2

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func HandleConn(conn net.Conn) {
	closeConnection := func(conn net.Conn) {
		log.Print("Connection closed.")
		conn.Close()
	}
	defer closeConnection(conn)

	u := newUser()

	for {
		buffer := make([]byte, 100)
		n, err := conn.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			log.Fatal(err)
		}

		if n < 2 {
			continue
		}

		buffer = buffer[:n-1]

		split := strings.Split(string(buffer), " ")
		command, err := newCommand(split[0], split[1:])
		if err != nil {
			_, err := conn.Write([]byte(err.Error()))
			if err != nil {
				log.Print(err)
			}
			continue
		}

		shouldClose := command.action(conn, u)
		if shouldClose {
			return
		}
	}
}

func mustWriteConn(message string, conn net.Conn) {
	_, err := conn.Write([]byte(message + "\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isDirectory(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}
