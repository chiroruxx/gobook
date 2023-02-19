package main

import "fmt"

func getNonZeroValue() (result int) {
	defer func() {
		recover()
		result++
	}()

	panic(nil)
}

func main() {
	fmt.Println(getNonZeroValue())
}
