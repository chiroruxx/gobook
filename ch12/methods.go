package main

import (
	"strings"
	"time"

	"gobook/ch12/methods"
)

func main() {
	methods.Print(time.Hour)

	methods.Print(new(strings.Replacer))
}
