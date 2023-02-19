package main

import "fmt"

func main() {
	f1(3)
}

func f1(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f1(x - 1)
}
