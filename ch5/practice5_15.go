package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func max(items ...int) *int {
	var value int
	isInitialized := false

	for _, item := range items {
		if !isInitialized {
			value = item
			isInitialized = true
			continue
		}

		if value < item {
			value = item
		}
	}

	if !isInitialized {
		return nil
	}

	return &value
}

func min(items ...int) *int {
	var value int
	isInitialized := false

	for _, item := range items {
		if !isInitialized {
			value = item
			isInitialized = true
			continue
		}

		if value > item {
			value = item
		}
	}

	if !isInitialized {
		return nil
	}

	return &value
}

func main() {
	var values []int
	for _, value := range os.Args[1:] {
		intValue, _ := strconv.Atoi(value)
		values = append(values, intValue)
	}

	maxPointer := max(values...)
	if maxPointer == nil {
		log.Fatalf("Max is nil")
	}

	minPointer := min(values...)
	if minPointer == nil {
		log.Fatalf("Min is nil")
	}

	fmt.Printf("Max: %d\n", *maxPointer)
	fmt.Printf("Min: %d\n", *minPointer)
}
