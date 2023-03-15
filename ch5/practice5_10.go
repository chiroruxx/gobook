package main

import (
	"fmt"
	"gobook/ch5/practice5_10"
	"gobook/ch5/toposort"
)

func main() {
	for i, course := range practice5_10.TopoSort(toposort.PreReqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
