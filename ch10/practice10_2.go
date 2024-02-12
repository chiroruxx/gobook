package main

import (
	"fmt"
	"os"

	"gobook/ch10/practice10_2"
	_ "gobook/ch10/practice10_2/tar"
	_ "gobook/ch10/practice10_2/zip"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		_, _ = fmt.Fprintln(os.Stderr, "usage: practice10_2 filepath")
		os.Exit(1)
	}
	path := args[1]

	archive, err := practice10_2.Read(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "practice10_2: %v\n", err)
	}

	for _, name := range archive.Names() {
		fmt.Println(name)
	}
	fmt.Println(archive.Size())
}
