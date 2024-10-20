package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Alignof(struct {
		bool
		float64
		int16
	}{}))

	fmt.Println(unsafe.Alignof(struct {
		float64
		int16
		bool
	}{}))

	fmt.Println(unsafe.Alignof(struct {
		bool
		int16
		float64
	}{}))
}
