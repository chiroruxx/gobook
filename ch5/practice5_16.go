package main

import (
	"fmt"
	"os"
	"strings"
)

func join(sep string, items ...string) string {
	return strings.Join(items, sep)
}

func main() {
	fmt.Printf(join(os.Args[1], os.Args[2:]...))
}
