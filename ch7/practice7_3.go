package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type tree2 struct {
	value       int
	isSet       bool
	left, right *tree2
}

func (t *tree2) add(value int) {
	if !t.isSet {
		t.setValue(value)
		return
	}

	if value < t.value {
		if t.left == nil {
			t.left = new(tree2)
		}
		t.left.add(value)
	} else {
		if t.right == nil {
			t.right = new(tree2)
		}
		t.right.add(value)
	}
}

func (t *tree2) setValue(value int) {
	t.value = value
	t.isSet = true
}

func (t *tree2) String() string {
	if t == nil {
		return ""
	}

	buffer := bytes.Buffer{}
	if t.left != nil {
		buffer.WriteString(t.left.String())
		buffer.WriteString(" - ")
	}

	buffer.WriteString(strconv.Itoa(t.value))

	if t.right != nil {
		buffer.WriteString(" - ")
		buffer.WriteString(t.right.String())
	}

	return buffer.String()
}

func main() {
	t := new(tree2)
	fmt.Println(t)
	t.add(5)
	fmt.Println(t)
	t.add(0)
	fmt.Println(t)
	t.add(1)
	fmt.Println(t)
	t.add(4)
	fmt.Println(t)
	t.add(3)
	fmt.Println(t)
}
