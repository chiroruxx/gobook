package main

import (
	"fmt"
	"gobook/ch5/practice5_11"
	"gobook/ch5/toposort"
	"os"
)

func main() {
	preReqs := toposort.PreReqs
	preReqs["linear algebra"] = []string{"calculus"}

	sorted, err := practice5_11.TopoSort(preReqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
