package main

import (
	"fmt"
	"runtime"
)

func main() {
	val := runtime.GOMAXPROCS(-1)
	fmt.Printf("GOMAXPROCS is %v\n", val)
}
