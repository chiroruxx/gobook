package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type clock struct {
	name string
	port int
}

func main() {
	clocks, err := getClocks()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(clocks)
	for _, clock := range clocks {
		go showTime(clock)
	}
	select {}
}

func getClocks() ([]clock, error) {
	inputs := os.Args[1:]
	var clocks []clock
	for _, input := range inputs {
		split := strings.Split(input, "=")
		if len(split) != 2 {
			return nil, fmt.Errorf("invalid argument %s", input)
		}
		name := split[0]
		port, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		clocks = append(clocks, clock{
			name: name,
			port: port,
		})
	}

	return clocks, nil
}

func showTime(clock clock) {
	address := "localhost:" + strconv.Itoa(clock.port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		buffer := make([]byte, 10)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		buffer = buffer[:n]
		fmt.Print(clock.name + ": " + string(buffer))
		time.Sleep(1 * time.Second)
	}
}
