package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"gobook/ch9/practice9_4"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: practice9_4 num message [p]")
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	message := os.Args[2]

	if len(os.Args) > 3 {
		practice9_4.SetPrint(true)
	}

	done := make(chan struct{})
	practice9_4.SetDone(done)

	var wakeupGroup sync.WaitGroup
	var closeGroup sync.WaitGroup

	start, end := practice9_4.MakeGoroutines(num, &wakeupGroup, &closeGroup)
	if start == nil {
		fmt.Println("no goroutines found")
		return
	}

	wakeupGroup.Wait()

	fmt.Println()

	t := time.Now()
	start <- message
	<-end
	s := time.Since(t)

	fmt.Println()

	close(done)
	closeGroup.Wait()

	fmt.Printf("Time: %v", s)
}
