package main

import (
	"fmt"
	"gobook/ch5/toposort"
)

func main() {
	for i, course := range toposort.TopoSort(toposort.PreReqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
