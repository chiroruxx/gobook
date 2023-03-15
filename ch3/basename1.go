package main

import (
	"fmt"
	"gobook/ch3/basename1"
	"os"
)

func main() {
	fmt.Print(basename1.Basename(os.Args[1]))
}
