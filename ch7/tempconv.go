package main

import (
	"flag"
	"fmt"
	"gobook/ch7/tempconv"
)

var temp = tempconv.SetCelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
