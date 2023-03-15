package main

import (
	"gobook/ch5/defer1"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	defer1.F(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
