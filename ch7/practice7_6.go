package main

import (
	"flag"
	"fmt"
	"gobook/ch7/practice7_6"
)

var temp2 = practice7_6.Celsius2Flag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp2)
}
