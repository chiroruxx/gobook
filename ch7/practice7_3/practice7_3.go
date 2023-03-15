package practice7_3

import (
	"bytes"
	"strconv"
)

type Tree struct {
	value       int
	isSet       bool
	left, right *Tree
}

func (t *Tree) Add(value int) {
	if !t.isSet {
		t.setValue(value)
		return
	}

	if value < t.value {
		if t.left == nil {
			t.left = new(Tree)
		}
		t.left.Add(value)
	} else {
		if t.right == nil {
			t.right = new(Tree)
		}
		t.right.Add(value)
	}
}

func (t *Tree) setValue(value int) {
	t.value = value
	t.isSet = true
}

func (t *Tree) String() string {
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
