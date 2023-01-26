package main

import (
	"fmt"
	"math"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	JPY
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "e", GBP: "g", JPY: "¥"}
	fmt.Println(symbol)

	m := make(map[float64]bool)
	m[math.NaN()] = true
	m[math.NaN()] = false
	m[math.NaN()] = true
	fmt.Println(m)
}
